// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
	outwayapi "go.nlx.io/nlx/outway/api"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

// nolint:funlen // this is a test
func TestRequestClaim(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup   func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context
		request *external.RequestClaimRequest
		want    *external.RequestClaimResponse
		wantErr error
	}{
		"when_the_proxy_metadata_is_missing": {
			request: &external.RequestClaimRequest{},
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.InvalidArgument, "request has invalid fields"),
		},
		"when_providing_an_empty_order_reference": {
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(t, context.Background())
			},
			request: &external.RequestClaimRequest{
				OrderReference: "",
				ServiceName:    "service-name",
			},
			wantErr: status.Error(codes.InvalidArgument, "request has invalid fields"),
		},
		"when_providing_an_empty_organization_serial_number": {
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(t, context.Background())
			},
			request: &external.RequestClaimRequest{
				OrderReference: "order-reference",
				ServiceName:    "service-name",
			},
			wantErr: status.Error(codes.InvalidArgument, "request has invalid fields"),
		},
		"when_providing_an_empty_service_name": {
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(t, context.Background())
			},
			request: &external.RequestClaimRequest{
				OrderReference: "order-reference",
				ServiceName:    "",
			},
			wantErr: status.Error(codes.InvalidArgument, "request has invalid fields"),
		},
		"when_requester_is_not_delegatee": {
			setup: func(t *testing.T, certBundle *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:                 "00000000000000000005",
						PublicKeyPEM:              "public-key",
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{},
						ValidUntil:                time.Now().Add(time.Hour),
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.New(codes.PermissionDenied, external.ErrorReason_ERROR_REASON_ORDER_NOT_FOUND_FOR_ORG, "order does not exist for your organization", nil),
		},
		"when_public_key_is_invalid": {
			setup: func(t *testing.T, certBundle *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    certBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: "arbitrary-invalid-public-key-in-pem-format",
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
						ValidUntil: time.Now().Add(time.Hour),
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.NewInternal("invalid public key format", nil),
		},
		"when_order_is revoked": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), orgCerts)

				publicKeyPEM, err := orgCerts.PublicKeyPEM()
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    orgCerts.Certificate().Subject.SerialNumber,
						PublicKeyPEM: publicKeyPEM,
						RevokedAt: sql.NullTime{
							Time:  now.Add(-1 * time.Hour),
							Valid: true,
						},
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_REVOKED, "order is revoked", nil),
		},
		"when_order_is_no_longer_valid": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    orgCerts.Certificate().Subject.SerialNumber,
						PublicKeyPEM: publicKeyPEM,
						ValidUntil:   now.Add(-1 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_EXPIRED, "order is expired", nil),
		},
		"when_service_not_found_in_access_proofs": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterPublicKeyPEM, err := requesterCertBundle.PublicKeyPEM()
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:                 requesterCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM:              requesterPublicKeyPEM,
						ValidUntil:                now.Add(4 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "non-existing-service",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.New(codes.NotFound, external.ErrorReason_ERROR_REASON_ORDER_DOES_NOT_CONTAIN_SERVICE, "service not found in order", nil),
		},
		"when_outway_not_found": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterPublicKeyPEM, err := requesterCertBundle.PublicKeyPEM()
				require.NoError(t, err)

				// nolint:dupl // this is a test
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    requesterCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: requesterPublicKeyPEM,
						ValidUntil:   now.Add(4 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetOutwaysByPublicKeyFingerprint(ctx, "public-key-fingerprint").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.NewInternal("could not find outway", nil),
		},
		"when_outway_sign_call_fails": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterPublicKeyPEM, err := requesterCertBundle.PublicKeyPEM()
				require.NoError(t, err)

				// nolint:dupl // this is a test
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    requesterCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: requesterPublicKeyPEM,
						ValidUntil:   now.Add(4 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetOutwaysByPublicKeyFingerprint(ctx, "public-key-fingerprint").
					Return([]*database.Outway{
						{
							SelfAddressAPI: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:                     "00000000000000000002",
						DelegateePublicKeyFingerprint: requesterCertBundle.PublicKeyFingerprint(),
						OrderReference:                "order-reference",
						AccessProof: &outwayapi.AccessProof{
							ServiceName:              "service-name",
							OrganizationSerialNumber: "00000000000000000001",
							PublicKeyFingerprint:     "public-key-fingerprint",
						},
						ExpiresAt: timestamppb.New(now.Add(4 * time.Hour)),
					}).
					Return(nil, errors.New("arbitrary error"))

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: grpcerrors.NewInternal("could not sign order claim via outway", nil),
		},
		// nolint:dupl // this is a test
		"happy_flow_with_short_valid_until": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterPublicKeyPEM, err := requesterCertBundle.PublicKeyPEM()
				require.NoError(t, err)

				// nolint:dupl // this is a test
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    requesterCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: requesterPublicKeyPEM,
						ValidUntil:   now.Add(2 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetOutwaysByPublicKeyFingerprint(ctx, "public-key-fingerprint").
					Return([]*database.Outway{
						{
							SelfAddressAPI: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:                     "00000000000000000002",
						DelegateePublicKeyFingerprint: requesterCertBundle.PublicKeyFingerprint(),
						OrderReference:                "order-reference",
						AccessProof: &outwayapi.AccessProof{
							ServiceName:              "service-name",
							OrganizationSerialNumber: "00000000000000000001",
							PublicKeyFingerprint:     "public-key-fingerprint",
						},
						ExpiresAt: timestamppb.New(now.Add(2 * time.Hour)),
					}).
					Return(&outwayapi.SignOrderClaimResponse{
						SignedOrderClaim: "signed-string",
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			want: &external.RequestClaimResponse{
				Claim: "signed-string",
			},
		},
		// nolint:dupl // this is a test
		"happy_flow": {
			setup: func(t *testing.T, orgCerts *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterPublicKeyPEM, err := requesterCertBundle.PublicKeyPEM()
				require.NoError(t, err)

				// nolint:dupl // this is a test
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    requesterCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: requesterPublicKeyPEM,
						ValidUntil:   now.Add(4 * time.Hour),
						OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000001",
										},
										ServiceName:          "service-name",
										PublicKeyFingerprint: "public-key-fingerprint",
									},
								},
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetOutwaysByPublicKeyFingerprint(ctx, "public-key-fingerprint").
					Return([]*database.Outway{
						{
							SelfAddressAPI: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:                     "00000000000000000002",
						DelegateePublicKeyFingerprint: requesterCertBundle.PublicKeyFingerprint(),
						OrderReference:                "order-reference",
						AccessProof: &outwayapi.AccessProof{
							ServiceName:              "service-name",
							OrganizationSerialNumber: "00000000000000000001",
							PublicKeyFingerprint:     "public-key-fingerprint",
						},
						ExpiresAt: timestamppb.New(now.Add(4 * time.Hour)),
					}).
					Return(&outwayapi.SignOrderClaimResponse{
						SignedOrderClaim: "signed-string",
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:                  "order-reference",
				ServiceName:                     "service-name",
				ServiceOrganizationSerialNumber: "00000000000000000001",
			},
			want: &external.RequestClaimResponse{
				Claim: "signed-string",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t, nil)

			ctx := tt.setup(t, bundle, mocks)

			actual, err := service.RequestClaim(ctx, tt.request)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}

			assert.Equal(t, tt.want, actual)
		})
	}
}

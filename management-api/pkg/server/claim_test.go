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
	outwayapi "go.nlx.io/nlx/outway/api"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

// nolint:funlen // this is a test
func TestRequestClaim(t *testing.T) {
	now := time.Now()

	arbitraryPublicKeyPEM := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArN5xGkM73tJsCpKny59e
5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63
pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5mBZO+Ku21V2QFr44tvMh5IZDX3RbMB/4K
ad6sapmSF00HWrqTVMkrEsZ98DTb5nwGLh3kISnct4tLyVSpsl9s1rtkSgGUcs1T
IvWxS2D2mOsSL1HRdUNcFQmzchbfG87kXPvicoOISAZDJKDqWp3iuH0gJpQ+XMBf
mcD90I7Z/cRQjWP3P93B3V06cJkd00cEIRcIQqF8N+lE01H88Fi+wePhZRy92NP5
4wIDAQAB
-----END PUBLIC KEY-----
`

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
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_providing_an_empty_order_reference": {
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(t, context.Background())
			},
			request: &external.RequestClaimRequest{
				OrderReference: "",
				ServiceName:    "service-name",
			},
			wantErr: status.Error(codes.InvalidArgument, "an order reference must be provided"),
		},
		"when_providing_an_empty_service_name": {
			setup: func(*testing.T, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(t, context.Background())
			},
			request: &external.RequestClaimRequest{
				OrderReference: "order-reference",
				ServiceName:    "",
			},
			wantErr: status.Error(codes.InvalidArgument, "an service name must be provided"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Internal, "invalid public key format"),
		},
		"when_public_key_fingerprint_does_not_equal_metadata_fingerprint": {
			setup: func(t *testing.T, certBundle *common_tls.CertificateBundle, mocks serviceMocks) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				// nolint:dupl // this is a test
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(ctx, "order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    certBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM: arbitraryPublicKeyPEM,
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Unauthenticated, "invalid public key for order"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Unauthenticated, "order is revoked"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Unauthenticated, "order is no longer valid"),
		},
		"when_service_not_found_in_access_proofs": {
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "non-existing-service",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.NotFound, "order with reference order-reference and organization serialnumber 00000000000000000002 and service name non-existing-service not found"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
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
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Internal, "could not find outway"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
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
							SelfAddress: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:      "00000000000000000002",
						OrderReference: "order-reference",
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
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			wantErr: status.Error(codes.Internal, "could not sign order claim via outway"),
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
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
							SelfAddress: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:      "00000000000000000002",
						OrderReference: "order-reference",
						AccessProof: &outwayapi.AccessProof{
							ServiceName:              "service-name",
							OrganizationSerialNumber: "00000000000000000001",
							PublicKeyFingerprint:     "public-key-fingerprint",
						},
						ExpiresAt: timestamppb.New(now.Add(2 * time.Hour)),
					}).
					Return(&outwayapi.SignOrderClaimResponse{
						SignedOrderclaim: "signed-string",
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
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
							{
								AccessProof: &database.AccessProof{
									OutgoingAccessRequest: &database.OutgoingAccessRequest{
										Organization: database.Organization{
											SerialNumber: "00000000000000000002",
										},
										ServiceName:          "service-name-2",
										PublicKeyFingerprint: "public-key-fingerprint-2",
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
							SelfAddress: "self-address",
						},
					}, nil)

				mocks.oc.
					EXPECT().
					SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
						Delegatee:      "00000000000000000002",
						OrderReference: "order-reference",
						AccessProof: &outwayapi.AccessProof{
							ServiceName:              "service-name",
							OrganizationSerialNumber: "00000000000000000001",
							PublicKeyFingerprint:     "public-key-fingerprint",
						},
						ExpiresAt: timestamppb.New(now.Add(4 * time.Hour)),
					}).
					Return(&outwayapi.SignOrderClaimResponse{
						SignedOrderclaim: "signed-string",
					}, nil)

				return ctx
			},
			request: &external.RequestClaimRequest{
				OrderReference:           "order-reference",
				ServiceName:              "service-name",
				OrganizationSerialNumber: "00000000000000000001",
			},
			want: &external.RequestClaimResponse{
				Claim: "signed-string",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t)

			ctx := tt.setup(t, bundle, mocks)

			actual, err := service.RequestClaim(ctx, tt.request)

			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)

				assert.Equal(t, tt.want, actual)
			}
		})
	}
}

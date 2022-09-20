// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/common/tls"
	directory "go.nlx.io/nlx/directory-api/api"
	mock_directory "go.nlx.io/nlx/directory-api/api/mock"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
)

var testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
	"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

//nolint:funlen // this is a test
func TestDelegationPlugin(t *testing.T) {
	tests := map[string]struct {
		wantErr                    *httperrors.NLXNetworkError
		wantHTTPStatusCode         int
		setup                      func(*mock_directory.MockDirectoryClient, *mock_management.MockClient, *DelegationPlugin)
		createManagementClientFunc func(client management.Client) createManagementClientFunc
		setHeaders                 func(*http.Request)
	}{
		"missing_order_reference_returns_an_errors": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.UnableToParseDelegationMetadataErr,
				Message:  "failed to parse delegation metadata",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
			},
		},

		"missing_delegator_returns_an_errors": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.UnableToParseDelegationMetadataErr,
				Message:  "failed to parse delegation metadata",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
		},

		"creating_management_client_errors": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.UnableToSetupManagementClientErr,
				Message:  "unable to setup the external management client",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return nil, fmt.Errorf("arbitrary error")
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)
			},
		},

		"getting_organization_management_proxy_errors": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.UnableToSetupManagementClientErr,
				Message:  "unable to setup the external management client",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(nil, fmt.Errorf("arbitrary error"))
			},
		},

		"error_while_retrieving_claim_returns_an_error": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.UnableToRequestClaimErr,
				Message:  "unable to request claim from 00000000000000000001",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference:                  "test-ref-123",
						ServiceName:                     "service-name",
						ServiceOrganizationSerialNumber: "00000000000000000002",
					}).
					Return(nil, errors.New("something went wrong"))
			},
		},

		"error_when_retrieving_invalid_jwt": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.ReceivedInvalidClaimErr,
				Message:  "received an invalid claim from 00000000000000000001",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference:                  "test-ref-123",
						ServiceName:                     "service-name",
						ServiceOrganizationSerialNumber: "00000000000000000002",
					}).
					Return(&external.RequestClaimResponse{
						Claim: "invalid_jwt",
					}, nil)
			},
		},

		//nolint:dupl // this is a test
		"order_not_found": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.OrderNotFoundErr,
				Message:  "order not found",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().RequestClaim(gomock.Any(), &external.RequestClaimRequest{
					OrderReference:                  "test-ref-123",
					ServiceOrganizationSerialNumber: "00000000000000000002",
					ServiceName:                     "service-name",
				}).Return(nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_NOT_FOUND, "order not found", nil))
			},
		},
		//nolint:dupl // this is a test
		"order_not_found_for_org": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.OrderDoesNotExistForYourOrganizationErr,
				Message:  "order does not exist for your organization",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().RequestClaim(gomock.Any(), &external.RequestClaimRequest{
					OrderReference:                  "test-ref-123",
					ServiceOrganizationSerialNumber: "00000000000000000002",
					ServiceName:                     "service-name",
				}).Return(nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_NOT_FOUND_FOR_ORG, "order not found for organization", nil))
			},
		},
		//nolint:dupl // this is a test
		"order_has_expired": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.OrderExpiredErr,
				Message:  "the order has expired",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().RequestClaim(gomock.Any(), &external.RequestClaimRequest{
					OrderReference:                  "test-ref-123",
					ServiceOrganizationSerialNumber: "00000000000000000002",
					ServiceName:                     "service-name",
				}).Return(nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_EXPIRED, "order has expired", nil))
			},
		},
		//nolint:dupl // this is a test
		"order_does_not_contain_service": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.OrderDoesNotContainServiceErr,
				Message:  "order does not contain the service 'service-name'",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().RequestClaim(gomock.Any(), &external.RequestClaimRequest{
					OrderReference:                  "test-ref-123",
					ServiceOrganizationSerialNumber: "00000000000000000002",
					ServiceName:                     "service-name",
				}).Return(nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_DOES_NOT_CONTAIN_SERVICE, "order does not contain service", nil))
			},
		},

		"order_has_been_revoked": {
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.OrderRevokedErr,
				Message:  "order is revoked",
			},
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)

				managementClient.EXPECT().RequestClaim(gomock.Any(), &external.RequestClaimRequest{
					OrderReference:                  "test-ref-123",
					ServiceOrganizationSerialNumber: "00000000000000000002",
					ServiceName:                     "service-name",
				}).Return(nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ERROR_REASON_ORDER_REVOKED, "order is revoked", nil))
			},
		},

		"invalid_claim_in_memory_should_trigger_requesting_a_new_claim": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)
				managementClient.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference:                  "test-ref-123",
						ServiceName:                     "service-name",
						ServiceOrganizationSerialNumber: "00000000000000000002",
					}).
					Return(&external.RequestClaimResponse{
						Claim: testToken,
					}, nil)

				plugin.claims.Store("00000000000000000001/test-ref-123/service-name", &claimData{
					Raw: testToken,
					JWTClaims: delegation.JWTClaims{
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now()),
						},
						Delegatee:      "00000000000000000001",
						OrderReference: "test-ref-123",
					},
				})
			},
		},

		"happy_flow_with_claim_in_memory": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				plugin.claims.Store("00000000000000000001/test-ref-123/service-name", &claimData{
					Raw: "claim",
					JWTClaims: delegation.JWTClaims{
						RegisteredClaims: jwt.RegisteredClaims{},
						Delegatee:        "00000000000000000001",
						OrderReference:   "test-ref-123",
					},
				})
			},
		},

		"happy_flow": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add(delegation.HTTPHeaderDelegator, "00000000000000000001")
				r.Header.Add(delegation.HTTPHeaderOrderReference, "test-ref-123")
			},
			createManagementClientFunc: func(client management.Client) createManagementClientFunc {
				return func(context.Context, string, *tls.CertificateBundle) (management.Client, error) {
					return client, nil
				}
			},
			setup: func(directoryClient *mock_directory.MockDirectoryClient, managementClient *mock_management.MockClient, plugin *DelegationPlugin) {
				directoryClient.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directory.GetOrganizationManagementAPIProxyAddressRequest{OrganizationSerialNumber: "00000000000000000001"}).Return(&directory.GetOrganizationManagementAPIProxyAddressResponse{Address: "management-proxy-address:8443"}, nil)
				managementClient.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference:                  "test-ref-123",
						ServiceName:                     "service-name",
						ServiceOrganizationSerialNumber: "00000000000000000002",
					}).
					Return(&external.RequestClaimResponse{
						Claim: testToken,
					}, nil)
			},
		},

		"happy_flow_without_delegation_headers": {
			wantHTTPStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			requestContext := fakeContext(&Destination{
				Service:                  "service-name",
				OrganizationSerialNumber: "00000000000000000002",
			})

			if tt.setHeaders != nil {
				tt.setHeaders(requestContext.Request)
			}

			directoryClient := mock_directory.NewMockDirectoryClient(ctrl)

			externalManagementClient := mock_management.NewMockClient(ctrl)
			pluginArgs := &NewDelegationPluginArgs{
				Logger:          zap.NewNop(),
				OrgCertificate:  nil,
				DirectoryClient: directoryClient,
			}

			if tt.createManagementClientFunc != nil {
				pluginArgs.CreateManagementClientFunc = tt.createManagementClientFunc(externalManagementClient)
			}

			plugin := NewDelegationPlugin(pluginArgs)

			if tt.setup != nil {
				tt.setup(directoryClient, externalManagementClient, plugin)
			}

			err := plugin.Serve(nopServeFunc)(requestContext)
			assert.NoError(t, err)

			response := requestContext.Response.(*httptest.ResponseRecorder).Result()

			defer response.Body.Close()

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			} else if tt.setHeaders != nil {
				assert.Equal(t, "00000000000000000001", requestContext.LogData["delegator"])
				assert.Equal(t, "test-ref-123", requestContext.LogData["orderReference"])
			}

			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)
		})
	}
}

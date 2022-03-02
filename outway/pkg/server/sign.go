package server

import (
	"context"
	"crypto"

	"github.com/golang-jwt/jwt/v4"
	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/outway/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OutwayService) SignOrderClaim(ctx context.Context, req *api.SignOrderClaimRequest) (*api.SignOrderClaimResponse, error) {
	logger := s.logger
	logger.Info("rpc request GetTermsOfServiceStatus")

	if err := req.ExpiresAt.CheckValid(); err != nil {
		s.logger.Error("invalid expiry time provided", zap.Error(err))
		return nil, status.Error(codes.Internal, "invalid expiry time provided")
	}

	signedClaim, err := s.signFunction(s.orgCert.PrivateKey(), delegation.JWTClaims{
		Delegatee:      req.Delegatee,
		OrderReference: req.OrderReference,
		AccessProof: &delegation.AccessProof{
			ServiceName:              req.AccessProof.ServiceName,
			OrganizationSerialNumber: req.AccessProof.OrganizationSerialNumber,
			PublicKeyFingerprint:     req.AccessProof.PublicKeyFingerprint,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(req.ExpiresAt.AsTime()),
			Issuer:    s.orgCert.Certificate().Subject.SerialNumber,
		},
	})

	if err != nil {
		s.logger.Error("unable to create signed string from private key", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to sign claim")
	}

	return &api.SignOrderClaimResponse{
		SignedOrderclaim: signedClaim,
	}, nil
}

func SignAsRS512(privateKey crypto.PrivateKey, claims delegation.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

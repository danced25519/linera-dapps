//nolint:nolintlint,dupl
package token

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	tokenproto "github.com/linera-hacker/linera-dapps/service/kline/proto/kline/zeus/v1/token"
	token "github.com/linera-hacker/linera-dapps/service/kline/zeus/pkg/mw/v1/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetToken(ctx context.Context, in *tokenproto.GetTokenRequest) (*tokenproto.GetTokenResponse, error) {
	handler, err := token.NewHandler(
		ctx,
		token.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetToken",
			"In", in,
			"Error", err,
		)
		return &tokenproto.GetTokenResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.GetToken(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetToken",
			"In", in,
			"Error", err,
		)
		return &tokenproto.GetTokenResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &tokenproto.GetTokenResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTokens(ctx context.Context, in *tokenproto.GetTokensRequest) (*tokenproto.GetTokensResponse, error) {
	handler, err := token.NewHandler(
		ctx,
		token.WithConds(in.Conds),
		token.WithOffset(in.GetOffset()),
		token.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTokens",
			"In", in,
			"Error", err,
		)
		return &tokenproto.GetTokensResponse{}, status.Error(codes.Internal, "internal server err")
	}

	infos, err := handler.GetTokens(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTokens",
			"In", in,
			"Error", err,
		)
		return &tokenproto.GetTokensResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &tokenproto.GetTokensResponse{
		Infos: infos,
	}, nil
}

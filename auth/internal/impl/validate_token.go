package impl

import (
	"context"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/pkg/client"
)

func ValidateToken(clt *client.Client, fields operations.PostAuthValidateTokenBody) (*bool, error) {
	ctx := context.Background()
	introspectedToken, err := clt.Client.RetrospectToken(ctx, fields.Token, clt.Config.Client, clt.Config.ClientSecret,
		clt.Config.Realm)
	if err != nil {
		return nil, err
	}
	return introspectedToken.Active, nil
}

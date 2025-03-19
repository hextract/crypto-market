package impl

import (
	"context"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/pkg/client"
)

func LoginUser(clt *client.Client, fields operations.PostAuthLoginBody) (*string, error) {
	ctx := context.Background()
	userToken, err := clt.Client.Login(ctx, clt.Config.Client, clt.Config.ClientSecret,
		clt.Config.Realm, fields.Login, fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}

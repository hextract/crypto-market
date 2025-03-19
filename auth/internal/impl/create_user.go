package impl

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/pkg/client"
)

func CreateUser(clt *client.Client, fields operations.PostAuthRegisterBody) (*string, error) {
	ctx := context.Background()
	user := gocloak.User{
		Email:    &fields.Email,
		Enabled:  gocloak.BoolP(true),
		Username: &fields.Login,
	}
	token, err := clt.GetAdminToken()
	if err != nil {
		return nil, err
	}
	userId, err := clt.Client.CreateUser(ctx, token.AccessToken, clt.Config.Realm, user)
	if err != nil {
		return nil, err
	}
	err = clt.Client.SetPassword(ctx, token.AccessToken, userId, clt.Config.Realm, fields.Password, false)
	if err != nil {
		return nil, err
	}
	userToken, err := clt.Client.Login(ctx, clt.Config.Client, clt.Config.ClientSecret, clt.Config.Realm, fields.Login, fields.Password)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}

package impl

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
	"github.com/h4x4d/crypto-market/pkg/client"
)

func ChangePasswordUser(clt *client.Client, fields operations.PostAuthChangePasswordBody) (*string, error) {
	ctx := context.Background()

	// Check Old Password
	_, err := clt.Client.Login(ctx, clt.Config.Client, clt.Config.ClientSecret,
		clt.Config.Realm, fields.Login, fields.OldPassword)
	if err != nil {
		return nil, err
	}

	token, err := clt.GetAdminToken()
	if err != nil {
		return nil, err
	}
	params := gocloak.GetUsersParams{
		Username: &fields.Login,
	}
	user, _ := clt.Client.GetUsers(ctx, token.AccessToken, clt.Config.Realm, params)
	err = clt.Client.SetPassword(ctx, token.AccessToken, *user[0].ID,
		clt.Config.Realm, fields.NewPassword, false)
	if err != nil {
		return nil, err
	}

	userToken, err := clt.Client.Login(ctx, clt.Config.Client, clt.Config.ClientSecret,
		clt.Config.Realm, fields.Login, fields.NewPassword)
	if err != nil {
		return nil, err
	}
	return &userToken.AccessToken, nil
}

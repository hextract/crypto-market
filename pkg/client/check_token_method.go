package client

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/h4x4d/crypto-market/pkg/models"
)

func (c Client) CheckToken(token string) (user *models.User, err error) {
	ctx := context.Background()
	usrInfo, err := c.Client.GetUserInfo(ctx, token, c.Config.Realm)
	if err != nil {
		return nil, err
	}
	exact := true
	params := gocloak.GetUsersParams{
		Email: usrInfo.Email,
		Exact: &exact,
	}
	adminToken, err := c.GetAdminToken()
	if err != nil {
		return nil, err
	}
	users, err := c.Client.GetUsers(ctx, adminToken.AccessToken, c.Config.Realm, params)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	userId := *users[0].ID
	return &models.User{
		UserID: userId,
	}, nil
}

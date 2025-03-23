package client

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
)

func (c Client) GetAdminToken() (*gocloak.JWT, error) {
	ctx := context.Background()
	token, err := c.Client.LoginAdmin(ctx, c.Config.Admin, c.Config.AdminPassword, c.Config.MasterRealm)

	if err != nil {
		return nil, err
	}
	return token, nil
}

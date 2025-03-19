package client

import (
	"context"
	"strconv"
)

func (c Client) GetTelegramId(userId string) (tgId int, err error) {
	ctx := context.Background()
	adminToken, err := c.GetAdminToken()
	if err != nil {
		return 0, err
	}
	user, err := c.Client.GetUserByID(ctx, adminToken.AccessToken, c.Config.Realm, userId)
	if err != nil {
		return 0, err
	}
	tgId, err = strconv.Atoi((*user.Attributes)["telegram_id"][0])
	if err != nil {
		return 0, err
	}
	return tgId, nil
}

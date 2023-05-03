package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ApplicationCommandPermissionsUpdate struct {
	Data *discord.GuildApplicationCommandPermissions `json:"d"`
}

func NewApplicationCommandPermissionsUpdate(rest *rest.Client, data []byte) (*ApplicationCommandPermissionsUpdate, error) {
	pk := new(ApplicationCommandPermissionsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}

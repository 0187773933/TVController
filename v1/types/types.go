package types

import (
	ir_controller_types "github.com/0187773933/IRController/v1/types"
)

type ConfigFile struct {
	Brand string `yaml:"brand"`
	IP string `yaml:"ip"`
	MAC string `yaml:"mac"`
	TimeoutSeconds int `yaml:"timeout_seconds"`
	DefaultInput int `yaml:"default_input"`
	DefaultVolume int `yaml:"default_volume"`
	VolumeResetLimit int `yaml:"volume_reset_limit"`
	WakeOnLan bool `yaml:"wake_on_lan"`
	LGWebSocketPort string `yaml:"lg_web_socket_port"`
	LGClientKey string `yaml:"lg_client_key"`
	VizioAuthToken string `yaml:"vizio_auth_token"`
	IRConfig ir_controller_types.ConfigFile `yaml:"ir"`
}
package connector

import (
	"maunium.net/go/mautrix/bridgev2/bridgeconfig"
	up "go.mau.fi/util/configupgrade"
)

type Config struct {
	// Configuración específica de Reddit si es necesaria
	// Por ejemplo, configuración del cliente OAuth
}

func (rc *RedditConnector) GetConfig() any {
	return rc.Config
}

func upgradeConfig(helper up.Helper) {
	helper.Copy(up.Str, "client_id")
	helper.Copy(up.Str, "client_secret")
	helper.Copy(up.Str, "user_agent")
}

var configUpgrader = up.SimpleUpgrader(upgradeConfig)

func init() {
	bridgeconfig.ExampleUpdaters[configUpgrader.YAMLPath] = configUpgrader
}

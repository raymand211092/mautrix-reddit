package connector

import (
	"context"
	"fmt"
	"net/http"
	
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/id"
)

// Verificar que RedditConnector implementa NetworkConnector
var _ bridgev2.NetworkConnector = (*RedditConnector)(nil)

type RedditConnector struct {
	Bridge *bridgev2.Bridge
	Config *Config
}

func (rc *RedditConnector) GetName() bridgev2.BridgeName {
	return bridgev2.BridgeName{
		DisplayName:      "Reddit",
		NetworkURL:       "https://reddit.com",
		NetworkIcon:      "mxc://maunium.net/reddit-icon",
		NetworkID:        "reddit",
		BeeperBridgeType: "reddit",
		DefaultPort:      29320,
	}
}

func (rc *RedditConnector) Init(bridge *bridgev2.Bridge) {
	rc.Bridge = bridge
}

func (rc *RedditConnector) Start(ctx context.Context) error {
	// Inicialización del conector si es necesaria
	return nil
}

func (rc *RedditConnector) Stop() {
	// Limpieza al detener el bridge
}

func (rc *RedditConnector) GetCapabilities() *bridgev2.NetworkGeneralCapabilities {
	return &bridgev2.NetworkGeneralCapabilities{
		DisappearingMessages: false,
		AggressiveUpdateInfo: false,
	}
}

func (rc *RedditConnector) GetLoginFlows() []bridgev2.LoginFlow {
	return []bridgev2.LoginFlow{{
		Name:        "Username and Password",
		Description: "Log in with your Reddit username and password",
		ID:          "password",
	}}
}

func (rc *RedditConnector) CreateLogin(ctx context.Context, user *bridgev2.User, flowID string) (bridgev2.LoginProcess, error) {
	if flowID != "password" {
		return nil, fmt.Errorf("unknown login flow ID: %s", flowID)
	}
	
	return &RedditLogin{
		User:      user,
		Connector: rc,
	}, nil
}

func (rc *RedditConnector) LoadUserLogin(ctx context.Context, login *bridgev2.UserLogin) error {
	client := &RedditClient{
		UserLogin: login,
		Connector: rc,
	}
	
	login.Client = client
	
	return nil
}

// NetworkAPI proporciona una interfaz HTTP para webhooks de Reddit si es necesario
func (rc *RedditConnector) GetDBMetaTypes() bridgev2.DatabaseMetaTypes {
	return bridgev2.DatabaseMetaTypes{}
}

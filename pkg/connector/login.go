package connector

import (
	"context"
	"fmt"
	
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/oauth2"
	
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"
)

// Verificar que RedditLogin implementa las interfaces necesarias
var _ bridgev2.LoginProcess = (*RedditLogin)(nil)
var _ bridgev2.LoginProcessUserInput = (*RedditLogin)(nil)

type RedditLogin struct {
	User      *bridgev2.User
	Connector *RedditConnector
	
	// Datos temporales del proceso de login
	username string
	password string
	clientID string
	clientSecret string
}

func (rl *RedditLogin) Start(ctx context.Context) (*bridgev2.LoginStep, error) {
	return &bridgev2.LoginStep{
		Type:         bridgev2.LoginStepTypeUserInput,
		StepID:       "credentials",
		Instructions: "Enter your Reddit credentials",
		UserInputParams: &bridgev2.LoginUserInputParams{
			Fields: []bridgev2.LoginInputDataField{
				{
					Type:        bridgev2.LoginInputFieldTypeUsername,
					ID:          "username",
					Name:        "Username",
					Description: "Your Reddit username",
				},
				{
					Type:        bridgev2.LoginInputFieldTypePassword,
					ID:          "password",
					Name:        "Password",
					Description: "Your Reddit password",
				},
				{
					Type:        bridgev2.LoginInputFieldTypePassword,
					ID:          "client_id",
					Name:        "Client ID",
					Description: "Reddit OAuth client ID (create app at https://www.reddit.com/prefs/apps)",
				},
				{
					Type:        bridgev2.LoginInputFieldTypePassword,
					ID:          "client_secret",
					Name:        "Client Secret",
					Description: "Reddit OAuth client secret",
				},
			},
		},
	}, nil
}

func (rl *RedditLogin) Cancel() {
	// Limpiar recursos si es necesario
}

func (rl *RedditLogin) SubmitUserInput(ctx context.Context, input map[string]string) (*bridgev2.LoginStep, error) {
	rl.username = input["username"]
	rl.password = input["password"]
	rl.clientID = input["client_id"]
	rl.clientSecret = input["client_secret"]
	
	// Crear cliente de Reddit para validar credenciales
	credentials := &reddit.Credentials{
		ID:       rl.clientID,
		Secret:   rl.clientSecret,
		Username: rl.username,
		Password: rl.password,
	}
	
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create Reddit client: %w", err)
	}
	
	// Verificar que las credenciales son válidas obteniendo info del usuario
	user, _, err := client.User.Get(ctx, rl.username)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}
	
	// Crear el UserLogin
	userLoginID := networkid.UserLoginID(fmt.Sprintf("reddit_%s", user.ID))
	
	userLogin := &bridgev2.UserLogin{
		ID: userLoginID,
		Metadata: &UserLoginMetadata{
			Username:     rl.username,
			ClientID:     rl.clientID,
			ClientSecret: rl.clientSecret,
		},
	}
	
	// Crear el cliente para este login
	redditClient := &RedditClient{
		UserLogin: userLogin,
		Connector: rl.Connector,
		client:    client,
		username:  rl.username,
	}
	
	userLogin.Client = redditClient
	
	// Completar el login
	return &bridgev2.LoginStep{
		Type:         bridgev2.LoginStepTypeComplete,
		StepID:       "complete",
		Instructions: fmt.Sprintf("Successfully logged in as %s", user.Name),
		CompleteParams: &bridgev2.LoginCompleteParams{
			UserLogin:    userLogin,
			UserLoginID:  userLoginID,
			DisplayName:  user.Name,
		},
	}, nil
}

// Metadata del UserLogin
type UserLoginMetadata struct {
	Username     string `json:"username"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

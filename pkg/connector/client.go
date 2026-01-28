package connector

import (
	"context"
	"fmt"
	"time"
	
	"github.com/vartanbeno/go-reddit/v2/reddit"
	
	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/bridgev2/networkid"
	"maunium.net/go/mautrix/event"
)

// Verificar que RedditClient implementa NetworkAPI
var _ bridgev2.NetworkAPI = (*RedditClient)(nil)

type RedditClient struct {
	UserLogin *bridgev2.UserLogin
	Connector *RedditConnector
	
	client   *reddit.Client
	username string
	
	// Canal para detener el polling
	stopPolling chan struct{}
}

func (rc *RedditClient) Connect(ctx context.Context) error {
	// Recrear el cliente de Reddit desde los metadatos
	metadata := rc.UserLogin.Metadata.(*UserLoginMetadata)
	
	credentials := &reddit.Credentials{
		ID:       metadata.ClientID,
		Secret:   metadata.ClientSecret,
		Username: metadata.Username,
		Password: "", // No almacenamos la contraseña
	}
	
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return fmt.Errorf("failed to create Reddit client: %w", err)
	}
	
	rc.client = client
	rc.username = metadata.Username
	
	// Iniciar polling de mensajes
	rc.stopPolling = make(chan struct{})
	go rc.pollMessages(ctx)
	
	return nil
}

func (rc *RedditClient) Disconnect() {
	if rc.stopPolling != nil {
		close(rc.stopPolling)
	}
}

func (rc *RedditClient) IsLoggedIn() bool {
	return rc.client != nil
}

func (rc *RedditClient) LogoutRemote(ctx context.Context) {
	// Reddit no requiere logout explícito
	rc.Disconnect()
}

func (rc *RedditClient) IsThisUser(ctx context.Context, userID networkid.UserID) bool {
	return string(userID) == rc.username
}

func (rc *RedditClient) GetChatInfo(ctx context.Context, portal *bridgev2.Portal) (*bridgev2.ChatInfo, error) {
	// Para DMs en Reddit, el chat info es bastante simple
	chatID := portal.ID
	
	return &bridgev2.ChatInfo{
		Name: &bridgev2.ChatName{
			Name: string(chatID),
		},
		Type: &bridgev2.ChatType{
			Type: "dm",
		},
	}, nil
}

func (rc *RedditClient) GetUserInfo(ctx context.Context, ghost *bridgev2.Ghost) (*bridgev2.UserInfo, error) {
	username := string(ghost.ID)
	
	user, _, err := rc.client.User.Get(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	
	return &bridgev2.UserInfo{
		Name: &user.Name,
		// Avatar podría extraerse si Reddit proporciona URLs de avatar
	}, nil
}

// HandleMatrixMessage envía un mensaje desde Matrix a Reddit
func (rc *RedditClient) HandleMatrixMessage(ctx context.Context, msg *bridgev2.MatrixMessage) (*bridgev2.MatrixMessageResponse, error) {
	// Obtener el destinatario del portal
	recipient := string(msg.Portal.ID)
	
	// Convertir el contenido del mensaje de Matrix a texto
	content := msg.Event.Content.AsMessage()
	text := content.Body
	
	// Enviar el mensaje a través de la API de Reddit
	_, _, err := rc.client.Message.Send(ctx, &reddit.SendMessageRequest{
		To:      recipient,
		Subject: "Message from Matrix",
		Text:    text,
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	
	return &bridgev2.MatrixMessageResponse{
		DB: &bridgev2.Message{
			ID:        networkid.MessageID(fmt.Sprintf("matrix_%d", time.Now().Unix())),
			SenderID:  networkid.UserID(rc.username),
			Timestamp: time.Now(),
		},
	}, nil
}

// pollMessages hace polling periódico de nuevos mensajes de Reddit
func (rc *RedditClient) pollMessages(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-rc.stopPolling:
			return
		case <-ticker.C:
			rc.fetchNewMessages(ctx)
		}
	}
}

func (rc *RedditClient) fetchNewMessages(ctx context.Context) {
	// Obtener mensajes no leídos
	messages, _, err := rc.client.Message.GetUnread(ctx, &reddit.ListMessageOptions{
		ListOptions: reddit.ListOptions{
			Limit: 25,
		},
	})
	
	if err != nil {
		// Log error pero no detener el polling
		return
	}
	
	for _, msg := range messages.Messages {
		rc.handleIncomingMessage(ctx, msg)
	}
}

func (rc *RedditClient) handleIncomingMessage(ctx context.Context, msg *reddit.Message) {
	// Crear o encontrar el portal para esta conversación
	portalKey := networkid.PortalKey{
		ID:       networkid.PortalID(msg.Author),
		Receiver: rc.UserLogin.ID,
	}
	
	portal, err := rc.UserLogin.Bridge.GetPortalByKey(ctx, portalKey)
	if err != nil {
		return
	}
	
	// Convertir el mensaje de Reddit a formato Matrix
	intent := rc.UserLogin.Bridge.GetGhostByID(ctx, networkid.UserID(msg.Author))
	
	matrixMsg := &bridgev2.MatrixMessage{
		Event: &event.Event{
			Content: event.Content{
				Parsed: &event.MessageEventContent{
					MsgType: event.MsgText,
					Body:    msg.Body,
				},
			},
		},
	}
	
	// Enviar el mensaje al portal
	portal.SendMessage(ctx, intent, matrixMsg)
	
	// Marcar como leído
	rc.client.Message.Read(ctx, []string{msg.FullID})
}

func (rc *RedditClient) GetCapabilities(ctx context.Context, portal *bridgev2.Portal) *bridgev2.NetworkRoomCapabilities {
	return &bridgev2.NetworkRoomCapabilities{
		FormattedText: false,
		Replies:       false,
		Edits:         false,
		Deletes:       false,
		Reactions:     false,
		ReadReceipts:  false,
		Typing:        false,
	}
}

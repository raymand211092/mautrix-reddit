# Ejemplos y Configuración Avanzada

## Ejemplos de Configuración

### Configuración Básica (Desarrollo)

```yaml
homeserver:
  address: http://localhost:8008
  domain: localhost

appservice:
  address: http://localhost:29320
  hostname: 0.0.0.0
  port: 29320
  
  database:
    type: sqlite3-fk-wal
    uri: file:mautrix-reddit.db?_txlock=immediate
  
  id: reddit
  bot:
    username: redditbot
    displayname: Reddit Bridge Bot
  
  ephemeral_events: true
  as_token: "generated-token"
  hs_token: "generated-token"

bridge:
  username_template: "reddit_{userid}"
  displayname_template: "{displayname} (Reddit)"
  allow_all_users: true
  
  permissions:
    "*": user

logging:
  min_level: debug
  writers:
    - type: stdout
      format: pretty-colored
```

### Configuración Producción (Docker + PostgreSQL)

```yaml
homeserver:
  address: https://matrix.example.com
  domain: example.com

appservice:
  address: https://reddit-bridge.example.com
  hostname: 0.0.0.0
  port: 29320
  
  database:
    type: postgres
    uri: postgres://mautrix_reddit:password@postgres:5432/mautrix_reddit?sslmode=require
  
  id: reddit
  bot:
    username: redditbot
    displayname: Reddit Bridge
    avatar: mxc://example.com/reddit-icon
  
  ephemeral_events: true

bridge:
  username_template: "reddit_{userid}"
  displayname_template: "{displayname}"
  allow_all_users: false
  
  permissions:
    "example.com": user
    "@admin:example.com": admin

logging:
  min_level: info
  writers:
    - type: stdout
      format: json
    - type: file
      format: json
      filename: /data/logs/mautrix-reddit.log
      max_size: 100
      max_backups: 10
      compress: true
```

### Configuración con Proxy

```yaml
homeserver:
  address: https://matrix.example.com
  domain: example.com
  http_proxy: http://proxy.example.com:8080
  https_proxy: http://proxy.example.com:8080

# ... resto de configuración
```

## Ejemplos de Uso

### Comandos del Bot

```
# En conversación con @redditbot:example.com

# Ver ayuda
help

# Login
login
> Username: tu_usuario_reddit
> Password: ********
> Client ID: abc123def456
> Client Secret: xyz789...

# Ver sesiones activas
sessions

# Iniciar conversación con usuario
pm spez
> Hola desde Matrix!

# Logout
logout
```

### Uso Programático

Si quieres integrar con scripts:

```python
# ejemplo_matrix.py
from matrix_client.client import MatrixClient

client = MatrixClient("https://matrix.example.com")
client.login("@usuario:example.com", "password")

# Enviar comando al bot
room = client.join_room("!roomid:example.com")
room.send_text("login")

# Esperar respuesta...
```

## Casos de Uso

### 1. Notificaciones de Modmail

```yaml
# En config.yaml, añadir configuración personalizada
reddit:
  enable_modmail: true
  modmail_subreddits:
    - mysubreddit
    - anothersubreddit
```

### 2. Múltiples Cuentas

```bash
# Login con cuenta 1
login
# Ingresar credenciales de cuenta1

# Login con cuenta 2 (en otro mensaje)
login
# Ingresar credenciales de cuenta2

# Listar sesiones
sessions
# Verás ambas cuentas
```

### 3. Bot de Respuestas Automáticas

```go
// custom_bot.go - Extensión personalizada
package main

import (
    "context"
    "strings"
    
    "maunium.net/go/mautrix/bridgev2"
)

func customMessageHandler(ctx context.Context, msg *bridgev2.MatrixMessage) {
    // Si el mensaje contiene ciertas palabras clave
    if strings.Contains(strings.ToLower(msg.Event.Content.AsMessage().Body), "precio") {
        // Responder automáticamente
        msg.Portal.SendMessage(ctx, intent, &bridgev2.MatrixMessage{
            Event: &event.Event{
                Content: event.Content{
                    Parsed: &event.MessageEventContent{
                        MsgType: event.MsgText,
                        Body:    "Para precios, visita: https://ejemplo.com/precios",
                    },
                },
            },
        })
    }
}
```

## Integraciones

### Con Home Assistant

```yaml
# configuration.yaml
notify:
  - platform: matrix
    name: reddit_bridge
    homeserver: https://matrix.example.com
    target_room: "!roomid:example.com"
    username: "@bot:example.com"
    password: "password"
```

Automation:
```yaml
automation:
  - alias: "Notificar nuevo mensaje Reddit"
    trigger:
      platform: state
      entity_id: sensor.reddit_messages
    action:
      service: notify.reddit_bridge
      data:
        message: "Nuevo mensaje de Reddit: {{ trigger.to_state.state }}"
```

### Con Webhooks

```python
# webhook_server.py
from flask import Flask, request
import requests

app = Flask(__name__)

@app.route('/reddit-webhook', methods=['POST'])
def reddit_webhook():
    data = request.json
    
    # Enviar a Matrix via bot
    requests.post(
        'https://matrix.example.com/_matrix/client/r0/rooms/!room:example.com/send/m.room.message',
        headers={'Authorization': 'Bearer YOUR_TOKEN'},
        json={
            'msgtype': 'm.text',
            'body': f"Nuevo post en r/{data['subreddit']}: {data['title']}"
        }
    )
    
    return 'OK'
```

## Troubleshooting Avanzado

### Debug de Mensajes

```bash
# Habilitar logging verbose
export DEBUG=*
./mautrix-reddit -c config.yaml -l trace
```

### Analizar Base de Datos

```bash
# SQLite
sqlite3 mautrix-reddit.db

# Ver tablas
.tables

# Ver mensajes
SELECT * FROM message ORDER BY timestamp DESC LIMIT 10;

# Ver sesiones
SELECT * FROM user_login;
```

### Limpiar Cache

```bash
# Detener bridge
docker-compose down

# Limpiar base de datos (¡CUIDADO! Esto borra todos los datos)
rm data/mautrix-reddit.db*

# Regenerar
docker-compose up -d
```

### Performance Tuning

```yaml
# config.yaml
appservice:
  database:
    # Para PostgreSQL
    max_open_conns: 20
    max_idle_conns: 2
    conn_max_lifetime: 0

# Ajustar polling interval en client.go
# ticker := time.NewTicker(5 * time.Second) // Más frecuente
```

## Migración

### Desde otro bridge

```bash
# 1. Exportar datos del bridge anterior
# (depende del bridge)

# 2. Importar a mautrix-reddit
# Contactar con los mantenedores para scripts de migración

# 3. Actualizar registration.yaml
# 4. Reiniciar servidor Matrix
```

### Backup y Restore

```bash
# Backup
tar -czf backup-$(date +%Y%m%d).tar.gz \
    data/mautrix-reddit.db* \
    data/config.yaml \
    data/registration.yaml

# Restore
tar -xzf backup-20260127.tar.gz -C /
docker-compose restart
```

## Monitoreo

### Prometheus Metrics

```yaml
# Añadir a config.yaml
metrics:
  enabled: true
  listen: 0.0.0.0:8001
```

Prometheus config:
```yaml
scrape_configs:
  - job_name: 'mautrix-reddit'
    static_configs:
      - targets: ['localhost:8001']
```

### Healthcheck

```bash
# HTTP endpoint
curl http://localhost:29320/health

# Docker healthcheck
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:29320/health || exit 1
```

## Recursos Adicionales

- [Reddit API Rate Limits](https://github.com/reddit-archive/reddit/wiki/API)
- [Matrix Spec](https://spec.matrix.org/)
- [mautrix-go Examples](https://github.com/mautrix/go/tree/main/example)

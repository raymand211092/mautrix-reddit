# mautrix-reddit

Un puente (bridge) de Matrix para Reddit que permite enviar y recibir mensajes directos de Reddit desde Matrix usando la nueva arquitectura bridgev2 de mautrix-go.

## Características

- ✅ Enviar y recibir mensajes directos de Reddit
- ✅ Autenticación OAuth con Reddit
- ✅ Sincronización automática de mensajes
- ✅ Soporte para múltiples cuentas de Reddit
- 🚧 Notificaciones de posts (próximamente)
- 🚧 Integración con subreddits (próximamente)

## Requisitos

- Go 1.21 o superior
- Un servidor Matrix (Synapse, Dendrite, Conduit, etc.)
- Una cuenta de Reddit
- Aplicación de Reddit OAuth (ver sección de configuración)

## Instalación

### 1. Compilar el bridge

```bash
git clone https://github.com/yourusername/mautrix-reddit.git
cd mautrix-reddit
go build -o mautrix-reddit ./cmd/mautrix-reddit
```

### 2. Generar configuración

```bash
./mautrix-reddit -e > config.yaml
```

Edita `config.yaml` y configura:
- `homeserver.address` - URL de tu servidor Matrix
- `homeserver.domain` - Dominio de tu servidor Matrix
- `appservice.database` - Ruta a la base de datos SQLite
- `bridge.permissions` - Usuarios permitidos

### 3. Generar registro de appservice

```bash
./mautrix-reddit -g -c config.yaml -r registration.yaml
```

### 4. Registrar el appservice en tu servidor Matrix

Copia el archivo `registration.yaml` a tu servidor Matrix y añádelo a la configuración:

**Synapse** (`homeserver.yaml`):
```yaml
app_service_config_files:
  - /path/to/mautrix-reddit/registration.yaml
```

Reinicia tu servidor Matrix.

### 5. Ejecutar el bridge

```bash
./mautrix-reddit -c config.yaml
```

## Configuración de Reddit OAuth

Para usar este bridge necesitas crear una aplicación en Reddit:

1. Ve a https://www.reddit.com/prefs/apps
2. Haz clic en "create another app..." o "are you a developer? create an app..."
3. Selecciona "script" como tipo de aplicación
4. Rellena los campos:
   - **name**: Cualquier nombre (ej: "Matrix Bridge")
   - **description**: Opcional
   - **about url**: Opcional
   - **redirect uri**: `http://localhost:8080` (no se usa pero es requerido)
5. Haz clic en "create app"
6. Anota el **client ID** (debajo del nombre de la app) y el **client secret**

## Uso

### Login

1. Inicia una conversación con el bot del bridge en Matrix: `@redditbot:tu-servidor.com`
2. Envía el comando: `login`
3. Sigue las instrucciones para ingresar:
   - Username de Reddit
   - Password de Reddit
   - Client ID de tu app OAuth
   - Client Secret de tu app OAuth

### Enviar mensajes

Una vez autenticado, puedes:
- Iniciar una conversación con un usuario de Reddit usando: `pm <username>`
- Los mensajes que recibas en Reddit aparecerán automáticamente como salas en Matrix
- Envía mensajes normalmente en Matrix y se enviarán a Reddit

### Comandos disponibles

- `login` - Autenticarse con Reddit
- `logout` - Cerrar sesión de Reddit
- `pm <username>` - Iniciar conversación con un usuario de Reddit
- `help` - Ver ayuda y comandos disponibles

## Arquitectura

Este bridge está construido usando:

- **mautrix-go bridgev2**: Framework moderno para bridges de Matrix
- **go-reddit**: Cliente de Go para la API de Reddit
- **Reddit OAuth API**: Para autenticación y envío de mensajes

### Estructura del proyecto

```
mautrix-reddit/
├── cmd/
│   └── mautrix-reddit/
│       └── main.go           # Punto de entrada
├── pkg/
│   └── connector/
│       ├── connector.go      # NetworkConnector principal
│       ├── client.go         # Cliente de Reddit (NetworkAPI)
│       ├── login.go          # Proceso de autenticación
│       └── config.go         # Configuración
├── config.yaml               # Configuración del bridge
├── registration.yaml         # Registro del appservice
├── go.mod
└── README.md
```

## Limitaciones conocidas

- Reddit está migrando de mensajes privados (PMs) a Reddit Chat. Este bridge actualmente usa la API de PMs que será deprecated.
- Rate limiting de Reddit puede afectar el envío masivo de mensajes
- No soporta edición o eliminación de mensajes (limitación de Reddit API)
- No soporta reacciones (no disponible en Reddit)

## Migración a Reddit Chat API

Reddit está en proceso de migrar a su nuevo sistema de Chat. Cuando la API de Chat esté completamente disponible y documentada, este bridge se actualizará para soportarla.

## Desarrollo

### Compilar

```bash
go build -o mautrix-reddit ./cmd/mautrix-reddit
```

### Ejecutar en modo debug

```bash
./mautrix-reddit -c config.yaml -l debug
```

### Contribuir

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/amazing-feature`)
3. Commit tus cambios (`git commit -m 'Add amazing feature'`)
4. Push a la rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

## Licencia

MIT License - ver [LICENSE](LICENSE) para más detalles

## Agradecimientos

- [mautrix-go](https://github.com/mautrix/go) - Framework de bridges de Matrix
- [mautrix-twilio](https://github.com/mautrix/twilio) - Ejemplo de bridge con bridgev2
- [go-reddit](https://github.com/vartanbeno/go-reddit) - Cliente de Reddit para Go
- [TextsHQ platform-reddit](https://github.com/textshq/platform-reddit) - Inspiración para la integración con Reddit

## Soporte

- Matrix room: `#mautrix-reddit:maunium.net`
- Issues: https://github.com/yourusername/mautrix-reddit/issues

## Roadmap

- [x] Mensajes directos básicos
- [x] Autenticación OAuth
- [ ] Soporte para Reddit Chat API
- [ ] Notificaciones de posts
- [ ] Integración con subreddits
- [ ] Soporte para modmail
- [ ] Backfill de mensajes históricos
- [ ] Media/imágenes en mensajes
- [ ] Markdown mejorado

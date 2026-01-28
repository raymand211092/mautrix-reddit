# mautrix-reddit

<p align="center">
  <img src="https://img.shields.io/github/workflow/status/USUARIO/mautrix-reddit/Build%20and%20Release?style=for-the-badge" alt="Build Status">
  <img src="https://img.shields.io/github/v/release/USUARIO/mautrix-reddit?style=for-the-badge" alt="Release">
  <img src="https://img.shields.io/github/license/USUARIO/mautrix-reddit?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/github/go-mod/go-version/USUARIO/mautrix-reddit?style=for-the-badge" alt="Go Version">
</p>

<p align="center">
  <b>Un puente Matrix-Reddit para mensajería unificada</b>
</p>

<p align="center">
  <a href="#características">Características</a> •
  <a href="#instalación">Instalación</a> •
  <a href="#uso">Uso</a> •
  <a href="#documentación">Documentación</a> •
  <a href="#contribuir">Contribuir</a>
</p>

---

## 🎯 Características

- ✅ **Mensajería Bidireccional** - Envía y recibe mensajes entre Matrix y Reddit
- ✅ **Autenticación OAuth** - Login seguro con Reddit
- ✅ **Multi-cuenta** - Soporta múltiples cuentas de Reddit simultáneamente
- ✅ **Sincronización Automática** - Polling continuo de nuevos mensajes
- ✅ **Docker Ready** - Despliegue fácil con Docker y docker-compose
- ✅ **Base de Datos Flexible** - SQLite para desarrollo, PostgreSQL para producción
- ✅ **Documentación Completa** - Guías detalladas para usuarios y desarrolladores

## 📦 Instalación

### Opción 1: Binarios Pre-compilados (Recomendado)

Descarga el binario para tu plataforma desde [Releases](https://github.com/USUARIO/mautrix-reddit/releases):

```bash
# Linux (amd64)
wget https://github.com/USUARIO/mautrix-reddit/releases/latest/download/mautrix-reddit-linux-amd64.tar.gz
tar -xzf mautrix-reddit-linux-amd64.tar.gz
./mautrix-reddit-linux-amd64 --help

# macOS (Apple Silicon)
curl -L https://github.com/USUARIO/mautrix-reddit/releases/latest/download/mautrix-reddit-darwin-arm64.tar.gz -o mautrix-reddit.tar.gz
tar -xzf mautrix-reddit.tar.gz
./mautrix-reddit-darwin-arm64 --help
```

### Opción 2: Docker

```bash
docker pull ghcr.io/USUARIO/mautrix-reddit:latest

# O con docker-compose
wget https://raw.githubusercontent.com/USUARIO/mautrix-reddit/main/docker-compose.yml
mkdir -p data
docker-compose run --rm mautrix-reddit -e > data/config.yaml
# Editar data/config.yaml
docker-compose up -d
```

### Opción 3: Compilar desde Código Fuente

```bash
# Requisitos: Go 1.21+
git clone https://github.com/USUARIO/mautrix-reddit.git
cd mautrix-reddit
make build
# o: go build ./cmd/mautrix-reddit
```

## 🚀 Inicio Rápido

### 1. Crear App en Reddit

1. Ve a https://www.reddit.com/prefs/apps
2. Click "create another app..."
3. Tipo: **script**
4. Redirect URI: `http://localhost:8080`
5. Anota el **Client ID** y **Client Secret**

### 2. Configurar el Bridge

```bash
./mautrix-reddit -e > config.yaml
nano config.yaml  # Editar con tu configuración de Matrix
```

### 3. Generar Registro

```bash
./mautrix-reddit -g -c config.yaml -r registration.yaml
```

### 4. Registrar en Matrix

Copia `registration.yaml` a tu servidor Matrix y añádelo a la configuración:

```yaml
# homeserver.yaml (Synapse)
app_service_config_files:
  - /path/to/registration.yaml
```

Reinicia tu servidor Matrix.

### 5. Ejecutar

```bash
./mautrix-reddit -c config.yaml
```

## 💬 Uso

1. Inicia conversación con el bot: `@redditbot:tu-servidor.com`
2. Comando: `login`
3. Ingresa tus credenciales de Reddit
4. Comando: `pm nombreusuario` para chatear

### Comandos Disponibles

| Comando | Descripción |
|---------|-------------|
| `help` | Muestra ayuda y comandos disponibles |
| `login` | Autenticarse con Reddit |
| `logout` | Cerrar sesión |
| `pm <usuario>` | Iniciar chat con usuario de Reddit |
| `sessions` | Ver sesiones activas |
| `ping` | Verificar conexión |

## 📚 Documentación

- 📖 [README Completo](README.md) - Documentación principal
- ⚡ [Guía Rápida](QUICKSTART.md) - Comenzar en 5 minutos
- 🔧 [Ejemplos Avanzados](EXAMPLES.md) - Configuración y casos de uso
- 🤝 [Guía de Contribución](CONTRIBUTING.md) - Cómo contribuir
- 🚀 [GitHub Actions](GITHUB_ACTIONS_GUIDE.md) - Releases automáticos

## 🏗️ Arquitectura

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Matrix    │ ◄─────► │  mautrix-    │ ◄─────► │   Reddit    │
│   Server    │         │   reddit     │         │     API     │
└─────────────┘         └──────────────┘         └─────────────┘
      ↑                        ↑                        ↑
      │                        │                        │
  Usuarios                  Bridge                 Polling
  Matrix               (bridgev2)              Automático
```

**Tecnologías:**
- [mautrix-go](https://github.com/mautrix/go) - Framework de bridges
- [go-reddit](https://github.com/vartanbeno/go-reddit) - Cliente Reddit API
- Go 1.21+ - Lenguaje de programación

## 🐳 Docker

### Imágenes Disponibles

```bash
# GitHub Container Registry (Recomendado)
ghcr.io/USUARIO/mautrix-reddit:latest
ghcr.io/USUARIO/mautrix-reddit:v0.1.0

# Soporta múltiples arquitecturas
# - linux/amd64
# - linux/arm64
```

### docker-compose.yml

```yaml
version: '3.8'
services:
  mautrix-reddit:
    image: ghcr.io/USUARIO/mautrix-reddit:latest
    restart: unless-stopped
    volumes:
      - ./data:/data
    ports:
      - "29320:29320"
```

## 🛠️ Desarrollo

```bash
# Clonar
git clone https://github.com/USUARIO/mautrix-reddit.git
cd mautrix-reddit

# Instalar dependencias
make deps

# Compilar
make build

# Tests
make test

# Linter
make lint

# Ejecutar
make run
```

## 📋 Requisitos

- **Runtime:**
  - Servidor Matrix (Synapse, Dendrite, Conduit)
  - Cuenta de Reddit
  - Aplicación OAuth de Reddit

- **Compilación:**
  - Go 1.21 o superior
  - Git

## ⚠️ Limitaciones

- Reddit está migrando a Chat API (este bridge usa PM API legacy)
- Rate limiting de Reddit puede afectar mensajes masivos
- No soporta edición/eliminación de mensajes
- No soporta reacciones
- Solo mensajes de texto (multimedia en desarrollo)

## 🗺️ Roadmap

- [ ] Migración a Reddit Chat API
- [ ] Soporte para multimedia (imágenes, videos)
- [ ] Backfill de mensajes históricos
- [ ] Integración con subreddits
- [ ] Notificaciones de posts
- [ ] Modmail para moderadores
- [ ] Markdown mejorado

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Por favor lee [CONTRIBUTING.md](CONTRIBUTING.md) para detalles.

### Contributors

<a href="https://github.com/USUARIO/mautrix-reddit/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=USUARIO/mautrix-reddit" />
</a>

## 📄 Licencia

[MIT License](LICENSE) - Este proyecto es software libre y de código abierto.

## 🙏 Agradecimientos

- [mautrix-go](https://github.com/mautrix/go) por Tulir Asokan
- [go-reddit](https://github.com/vartanbeno/go-reddit) por Vartan Benohanian
- Inspirado en [mautrix-twilio](https://github.com/mautrix/twilio)
- Conceptos de [TextsHQ platform-reddit](https://github.com/textshq/platform-reddit)

## 💬 Soporte

- 🐛 [Reportar Bug](https://github.com/USUARIO/mautrix-reddit/issues/new?template=bug_report.md)
- 💡 [Solicitar Feature](https://github.com/USUARIO/mautrix-reddit/issues/new?template=feature_request.md)
- 💬 Matrix: `#mautrix-reddit:maunium.net`

## 📊 Estado

![GitHub issues](https://img.shields.io/github/issues/USUARIO/mautrix-reddit?style=flat-square)
![GitHub pull requests](https://img.shields.io/github/issues-pr/USUARIO/mautrix-reddit?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/USUARIO/mautrix-reddit?style=flat-square)
![GitHub stars](https://img.shields.io/github/stars/USUARIO/mautrix-reddit?style=social)

---

<p align="center">
  Hecho con ❤️ para la comunidad de Matrix
</p>

<p align="center">
  <sub>Si este proyecto te es útil, considera darle una ⭐</sub>
</p>

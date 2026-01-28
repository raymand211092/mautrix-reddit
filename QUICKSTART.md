# Guía de Inicio Rápido - mautrix-reddit

## Instalación Rápida (Docker)

```bash
# 1. Clonar el repositorio
git clone https://github.com/yourusername/mautrix-reddit.git
cd mautrix-reddit

# 2. Crear directorio de datos
mkdir -p data

# 3. Generar configuración
docker-compose run --rm mautrix-reddit -e > data/config.yaml

# 4. Editar configuración
nano data/config.yaml
```

Configura:
- `homeserver.address`: URL de tu servidor Matrix (ej: `https://matrix.example.com`)
- `homeserver.domain`: Dominio de tu servidor (ej: `example.com`)
- `bridge.permissions`: Añade tu usuario de Matrix

```yaml
bridge:
  permissions:
    "example.com": user
    "@tuusuario:example.com": admin
```

```bash
# 5. Generar archivo de registro
docker-compose run --rm mautrix-reddit -g

# 6. Copiar registration.yaml a tu servidor Matrix
# Para Synapse, añade en homeserver.yaml:
# app_service_config_files:
#   - /path/to/mautrix-reddit/data/registration.yaml

# 7. Reiniciar tu servidor Matrix

# 8. Iniciar el bridge
docker-compose up -d

# 9. Ver logs
docker-compose logs -f
```

## Instalación Manual

```bash
# 1. Compilar
go build -o mautrix-reddit ./cmd/mautrix-reddit

# 2. Generar configuración
./mautrix-reddit -e > config.yaml

# 3. Editar config.yaml (ver arriba)

# 4. Generar registro
./mautrix-reddit -g

# 5. Copiar registration.yaml a tu servidor Matrix y reiniciarlo

# 6. Ejecutar el bridge
./mautrix-reddit
```

## Primer Uso

### 1. Crear Aplicación en Reddit

1. Ve a https://www.reddit.com/prefs/apps
2. Clic en "create another app..."
3. Configuración:
   - **name**: "Matrix Bridge" (o cualquier nombre)
   - **type**: Selecciona "script"
   - **redirect uri**: `http://localhost:8080`
4. Anota el **Client ID** y **Client Secret**

### 2. Iniciar Sesión en Matrix

1. Abre tu cliente de Matrix (Element, FluffyChat, etc.)
2. Inicia conversación con: `@redditbot:tu-servidor.com`
3. Envía: `login`
4. Proporciona:
   - Username de Reddit
   - Password de Reddit
   - Client ID
   - Client Secret

### 3. Enviar Mensajes

```
# Iniciar chat con un usuario de Reddit
pm nombreusuario

# Envía mensajes normales en la sala que se crea
Hola desde Matrix!
```

## Comandos Útiles

- `help` - Ver todos los comandos
- `login` - Autenticarse con Reddit
- `logout` - Cerrar sesión
- `pm <usuario>` - Iniciar chat con usuario de Reddit
- `ping` - Verificar que el bridge funciona

## Solución de Problemas

### El bridge no responde

```bash
# Ver logs
docker-compose logs -f mautrix-reddit

# o si lo ejecutas manualmente
./mautrix-reddit -l debug
```

### Error de autenticación

- Verifica que el Client ID y Secret sean correctos
- Asegúrate de que la aplicación de Reddit sea tipo "script"
- Verifica que tu usuario y contraseña de Reddit sean correctos

### No recibo mensajes

- Reddit tiene rate limiting. Espera unos segundos entre mensajes
- Verifica que el polling esté funcionando en los logs
- Asegúrate de que tienes mensajes sin leer en Reddit

### Error al conectar con Matrix

- Verifica que `homeserver.address` en config.yaml sea correcto
- Asegúrate de que el archivo `registration.yaml` esté en la configuración de tu servidor Matrix
- Reinicia tu servidor Matrix después de añadir el appservice

## Configuración Avanzada

### Múltiples Cuentas de Reddit

Puedes autenticarte con varias cuentas de Reddit. Simplemente ejecuta `login` múltiples veces con diferentes credenciales.

### Permisos Granulares

En `config.yaml`:

```yaml
bridge:
  permissions:
    "*": relay              # Todos: solo relay
    "example.com": user     # Tu servidor: usuarios normales
    "@admin:example.com": admin  # Tú: administrador
```

### Base de Datos PostgreSQL

En lugar de SQLite:

```yaml
appservice:
  database:
    type: postgres
    uri: postgres://usuario:contraseña@localhost/mautrix_reddit?sslmode=disable
```

## Próximos Pasos

- Únete a la sala de soporte: `#mautrix-reddit:maunium.net`
- Reporta bugs: https://github.com/yourusername/mautrix-reddit/issues
- Contribuye: Lee CONTRIBUTING.md

## Recursos

- [Documentación completa](README.md)
- [Configuración de Reddit OAuth](https://github.com/reddit-archive/reddit/wiki/OAuth2)
- [Documentación de mautrix-go](https://github.com/mautrix/go)
- [Documentación de Matrix](https://matrix.org/docs/)

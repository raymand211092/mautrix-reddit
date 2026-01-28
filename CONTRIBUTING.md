# Guía de Contribución

¡Gracias por tu interés en contribuir a mautrix-reddit! Este documento proporciona pautas para contribuir al proyecto.

## Código de Conducta

Este proyecto sigue el [Código de Conducta de Contributor Covenant](https://www.contributor-covenant.org/). Al participar, se espera que mantengas este código.

## Cómo Contribuir

### Reportar Bugs

1. Verifica que el bug no haya sido reportado antes en [Issues](https://github.com/yourusername/mautrix-reddit/issues)
2. Si no existe, crea un nuevo issue con:
   - Título descriptivo
   - Pasos para reproducir
   - Comportamiento esperado vs actual
   - Versión de mautrix-reddit
   - Logs relevantes (sanitiza información sensible)

### Sugerir Funcionalidades

1. Verifica que la funcionalidad no haya sido sugerida antes
2. Crea un issue con:
   - Título descriptivo
   - Caso de uso detallado
   - Beneficios esperados
   - Posible implementación (opcional)

### Pull Requests

1. **Fork el repositorio**
   ```bash
   git clone https://github.com/yourusername/mautrix-reddit.git
   cd mautrix-reddit
   ```

2. **Crea una rama para tu feature**
   ```bash
   git checkout -b feature/nombre-descriptivo
   ```

3. **Realiza tus cambios**
   - Sigue las convenciones de código
   - Añade tests si es aplicable
   - Actualiza documentación si es necesario

4. **Verifica que todo funcione**
   ```bash
   make fmt      # Formatear código
   make lint     # Ejecutar linter
   make test     # Ejecutar tests
   make build    # Compilar
   ```

5. **Commit con mensajes descriptivos**
   ```bash
   git commit -m "feat: añadir soporte para mensajes multimedia"
   ```
   
   Formato de commits:
   - `feat:` - Nueva funcionalidad
   - `fix:` - Corrección de bug
   - `docs:` - Cambios en documentación
   - `style:` - Formato, espacios, etc.
   - `refactor:` - Refactorización de código
   - `test:` - Añadir o modificar tests
   - `chore:` - Tareas de mantenimiento

6. **Push a tu fork**
   ```bash
   git push origin feature/nombre-descriptivo
   ```

7. **Crea un Pull Request**
   - Describe los cambios claramente
   - Referencia issues relacionados
   - Incluye screenshots si es relevante

## Estándares de Código

### Go

- Usa `gofmt` para formatear código
- Sigue las [Effective Go guidelines](https://golang.org/doc/effective_go)
- Usa nombres descriptivos para variables y funciones
- Comenta código complejo
- Mantén funciones pequeñas y enfocadas

### Estructura de Paquetes

```
pkg/connector/
├── connector.go    # NetworkConnector principal
├── client.go       # NetworkAPI/cliente
├── login.go        # Proceso de autenticación
├── config.go       # Configuración
└── utils.go        # Utilidades compartidas
```

### Tests

- Escribe tests para nueva funcionalidad
- Mantén >80% de cobertura para código crítico
- Usa nombres descriptivos: `TestFunctionName_Scenario_ExpectedResult`

```go
func TestRedditClient_SendMessage_Success(t *testing.T) {
    // Test implementation
}
```

### Documentación

- Comenta funciones exportadas
- Actualiza README.md para cambios de funcionalidad
- Actualiza QUICKSTART.md para cambios de configuración
- Añade ejemplos cuando sea útil

## Proceso de Revisión

1. Un mantenedor revisará tu PR
2. Pueden solicitar cambios
3. Una vez aprobado, será merged

### Criterios de Aceptación

- [ ] El código compila sin errores
- [ ] Los tests pasan
- [ ] El linter no reporta errores
- [ ] La documentación está actualizada
- [ ] Los commits son descriptivos
- [ ] No hay conflictos con main

## Configuración de Desarrollo

### Requisitos

- Go 1.21+
- Git
- Docker (opcional, para testing)
- golangci-lint (para linting)

### Setup

```bash
# Clonar
git clone https://github.com/yourusername/mautrix-reddit.git
cd mautrix-reddit

# Instalar dependencias
make deps

# Compilar
make build

# Ejecutar tests
make test

# Ejecutar linter
make lint
```

### Testing Local

1. Copia `example-config.yaml` a `config.yaml`
2. Configura con tu servidor Matrix de desarrollo
3. Ejecuta `make run`

### Debugging

Usa nivel de log debug:
```bash
./mautrix-reddit -c config.yaml -l debug
```

## Arquitectura

### Componentes Principales

1. **Connector** (`connector.go`)
   - Implementa `bridgev2.NetworkConnector`
   - Gestión de alto nivel del bridge

2. **Client** (`client.go`)
   - Implementa `bridgev2.NetworkAPI`
   - Comunicación con Reddit API
   - Polling de mensajes

3. **Login** (`login.go`)
   - Implementa `bridgev2.LoginProcess`
   - Flujo de autenticación OAuth

### Flujo de Datos

```
Matrix → Bridge → Client → Reddit API
  ↑                               ↓
  └──────── Polling ←─────────────┘
```

## Áreas que Necesitan Ayuda

- [ ] Soporte para Reddit Chat API (nueva)
- [ ] Backfill de mensajes históricos
- [ ] Soporte para multimedia
- [ ] Mejoras en markdown
- [ ] Tests de integración
- [ ] Documentación adicional

## Preguntas

Si tienes preguntas:

1. Revisa la documentación
2. Busca en issues existentes
3. Pregunta en la sala de Matrix: `#mautrix-reddit:maunium.net`
4. Crea un issue con la etiqueta "question"

## Recursos

- [mautrix-go docs](https://pkg.go.dev/maunium.net/go/mautrix)
- [bridgev2 docs](https://pkg.go.dev/maunium.net/go/mautrix/bridgev2)
- [Reddit API docs](https://www.reddit.com/dev/api)
- [Matrix spec](https://spec.matrix.org/)

## Licencia

Al contribuir, aceptas que tus contribuciones se licencien bajo la licencia MIT del proyecto.

# 🚀 Guía de GitHub Actions y Releases Automáticos

## 📋 Resumen

Este proyecto está configurado con **GitHub Actions** para compilar automáticamente binarios multiplataforma cuando creas un release o tag.

## 🎯 ¿Qué se compila automáticamente?

Cuando creas un tag de versión (ej: `v0.1.0`), GitHub Actions compilará:

### Binarios para múltiples plataformas:
- ✅ **Linux** (amd64, arm64)
- ✅ **macOS** (amd64, arm64) - Intel y Apple Silicon
- ✅ **Windows** (amd64)

### Imágenes Docker:
- ✅ **Multi-arquitectura** (amd64, arm64)
- ✅ Publicadas en GitHub Container Registry
- ✅ Opcionalmente en Docker Hub

## 📦 Formato de Releases

Cada release incluirá:
```
mautrix-reddit-linux-amd64.tar.gz
mautrix-reddit-linux-arm64.tar.gz
mautrix-reddit-darwin-amd64.tar.gz
mautrix-reddit-darwin-arm64.tar.gz
mautrix-reddit-windows-amd64.zip
```

Cada archivo contendrá:
- El binario compilado
- `example-config.yaml`
- `README.md`
- `LICENSE`

## 🔧 Configuración Inicial

### 1. Subir a GitHub

Opción A - Usar el script automático:
```bash
chmod +x setup-github.sh
./setup-github.sh tu-usuario-github
```

Opción B - Manual:
```bash
# Inicializar git
git init
git add .
git commit -m "Initial commit"

# Crear repositorio en GitHub (en la web)
# Luego:
git remote add origin https://github.com/TU_USUARIO/mautrix-reddit.git
git branch -M main
git push -u origin main
```

### 2. Configurar Secrets (Opcional - solo para Docker)

Si quieres publicar en Docker Hub:

1. Ve a tu repo → `Settings` → `Secrets and variables` → `Actions`
2. Añade estos secrets:
   - `DOCKERHUB_USERNAME`: Tu usuario de Docker Hub
   - `DOCKERHUB_TOKEN`: Token de acceso de Docker Hub

**Nota**: Las imágenes Docker se publican automáticamente en GitHub Container Registry sin configuración adicional.

## 📝 Crear un Release

### Método 1: Crear Tag desde la línea de comandos

```bash
# Crear y publicar un tag
git tag -a v0.1.0 -m "Release v0.1.0

Features:
- OAuth authentication
- Bidirectional messaging
- Multi-account support"

git push origin v0.1.0
```

### Método 2: Crear Release desde GitHub

1. Ve a tu repositorio en GitHub
2. Click en `Releases` → `Create a new release`
3. En "Choose a tag", escribe `v0.1.0` y selecciona "Create new tag"
4. Título: `v0.1.0`
5. Descripción: Añade notas de la versión
6. Click en `Publish release`

### Método 3: Usar GitHub CLI

```bash
gh release create v0.1.0 \
  --title "Release v0.1.0" \
  --notes "Initial release with OAuth and messaging support"
```

## ⚡ Proceso Automático

Cuando creas un tag/release:

1. **GitHub Actions se activa automáticamente**
2. **Compila** los binarios para todas las plataformas
3. **Crea** archivos comprimidos (.tar.gz y .zip)
4. **Sube** los binarios al release en GitHub
5. **Construye** imágenes Docker multiplataforma
6. **Publica** las imágenes en registros

## 📊 Monitorear el Progreso

1. Ve a tu repositorio en GitHub
2. Click en la pestaña `Actions`
3. Verás el workflow "Build and Release" ejecutándose
4. Click en él para ver los logs en tiempo real

## 🐳 Usar las Imágenes Docker

Después del release, las imágenes estarán disponibles:

```bash
# Desde GitHub Container Registry (siempre disponible)
docker pull ghcr.io/TU_USUARIO/mautrix-reddit:v0.1.0
docker pull ghcr.io/TU_USUARIO/mautrix-reddit:latest

# Desde Docker Hub (si configuraste los secrets)
docker pull TU_USUARIO/mautrix-reddit:v0.1.0
```

## 💾 Descargar Binarios

Los usuarios pueden descargar los binarios desde:
```
https://github.com/TU_USUARIO/mautrix-reddit/releases
```

Ejemplo de uso:
```bash
# Linux
wget https://github.com/TU_USUARIO/mautrix-reddit/releases/download/v0.1.0/mautrix-reddit-linux-amd64.tar.gz
tar -xzf mautrix-reddit-linux-amd64.tar.gz
./mautrix-reddit-linux-amd64 --help

# macOS
curl -L https://github.com/TU_USUARIO/mautrix-reddit/releases/download/v0.1.0/mautrix-reddit-darwin-amd64.tar.gz -o mautrix-reddit.tar.gz
tar -xzf mautrix-reddit.tar.gz
./mautrix-reddit-darwin-amd64 --help

# Windows
# Descargar mautrix-reddit-windows-amd64.zip desde el navegador
# Extraer y ejecutar mautrix-reddit-windows-amd64.exe
```

## 🔄 Workflow de Desarrollo

### Commits regulares
```bash
git add .
git commit -m "feat: añadir nueva funcionalidad"
git push
```
→ GitHub Actions **compila y prueba** pero **NO crea release**

### Crear nueva versión
```bash
git tag -a v0.2.0 -m "Version 0.2.0"
git push origin v0.2.0
```
→ GitHub Actions **compila, prueba Y crea release** con binarios

## 📈 Versionado Semántico

Usa [Semantic Versioning](https://semver.org/):

- **v1.0.0** - Release mayor (breaking changes)
- **v0.2.0** - Release menor (nuevas features)
- **v0.1.1** - Patch (bug fixes)

Ejemplos:
```bash
# Bug fix
git tag -a v0.1.1 -m "Fix: corregir error en autenticación"

# Nueva feature
git tag -a v0.2.0 -m "Feature: soporte para multimedia"

# Breaking change
git tag -a v1.0.0 -m "Major: migración a Reddit Chat API"
```

## 🎨 Personalizar el Workflow

Edita `.github/workflows/build.yml` para:

### Añadir más plataformas:
```yaml
- goos: freebsd
  goarch: amd64
```

### Cambiar flags de compilación:
```yaml
go build -ldflags="-s -w -X 'main.CustomVar=value'"
```

### Añadir pasos adicionales:
```yaml
- name: Run integration tests
  run: go test -tags=integration ./...
```

## 🔍 Troubleshooting

### El workflow falla en "Build and push Docker"
- **Causa**: No hay secrets de Docker Hub configurados
- **Solución**: Es normal, las imágenes se publican en GitHub Container Registry. Si quieres Docker Hub, añade los secrets.

### No se crean los releases automáticamente
- **Causa**: El tag no tiene formato `v*` (ej: `v0.1.0`)
- **Solución**: Asegúrate de usar el prefijo `v` en los tags

### Error: "permission denied"
- **Causa**: GitHub Actions no tiene permisos para crear releases
- **Solución**: Ve a Settings → Actions → General → Workflow permissions → Selecciona "Read and write permissions"

### Los binarios no se suben al release
- **Causa**: El workflow se ejecutó pero hubo errores en el paso de release
- **Solución**: Revisa los logs en la pestaña Actions

## 📚 Recursos Adicionales

- [GitHub Actions Docs](https://docs.github.com/en/actions)
- [Creating Releases](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)
- [Semantic Versioning](https://semver.org/)
- [Docker Multi-platform](https://docs.docker.com/build/building/multi-platform/)

## ✅ Checklist de Primera Release

- [ ] Código subido a GitHub
- [ ] Script `setup-github.sh` ejecutado (o pasos manuales completados)
- [ ] Permisos de workflow configurados (read/write)
- [ ] Secrets de Docker configurados (opcional)
- [ ] Tag `v0.1.0` creado y pusheado
- [ ] Workflow ejecutándose en Actions
- [ ] Release creado con binarios
- [ ] Binarios descargados y probados

## 🎉 ¡Listo!

Una vez configurado, cada vez que hagas:
```bash
git tag -a v0.X.Y -m "Release notes"
git push origin v0.X.Y
```

Tendrás automáticamente:
- ✅ Binarios para 6 plataformas diferentes
- ✅ Imágenes Docker multi-arquitectura
- ✅ Release notes automáticos
- ✅ Todo listo para que los usuarios descarguen

**No necesitas compilar manualmente nunca más!** 🚀

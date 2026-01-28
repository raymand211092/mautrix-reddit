# 📤 Guía Paso a Paso: Subir a GitHub

## 🎯 Objetivo

Subir el proyecto mautrix-reddit a GitHub y configurar compilación automática de binarios.

## 📋 Pre-requisitos

- [ ] Cuenta de GitHub
- [ ] Git instalado en tu computadora
- [ ] (Opcional) GitHub CLI (`gh`) instalado

## 🚀 Método 1: Automático con Script (RECOMENDADO)

### Paso 1: Descomprimir el Proyecto

```bash
# Descomprimir
tar -xzf mautrix-reddit-github-ready.tar.gz
cd mautrix-reddit
```

### Paso 2: Ejecutar el Script

```bash
# Dar permisos
chmod +x setup-github.sh

# Ejecutar (reemplaza TU_USUARIO con tu usuario de GitHub)
./setup-github.sh TU_USUARIO mautrix-reddit
```

El script te guiará por todo el proceso automáticamente:
- ✅ Inicializa Git
- ✅ Crea el commit inicial
- ✅ Crea el repositorio en GitHub (si tienes `gh` CLI)
- ✅ Sube el código
- ✅ Crea el tag v0.1.0
- ✅ GitHub Actions compila automáticamente

### Paso 3: Esperar la Compilación

1. Ve a `https://github.com/TU_USUARIO/mautrix-reddit/actions`
2. Verás el workflow "Build and Release" ejecutándose
3. Espera ~5-10 minutos

### Paso 4: Descargar Binarios

1. Ve a `https://github.com/TU_USUARIO/mautrix-reddit/releases`
2. Verás el release v0.1.0 con los binarios compilados
3. ¡Listo! 🎉

---

## 🔧 Método 2: Manual (Sin GitHub CLI)

### Paso 1: Descomprimir el Proyecto

```bash
tar -xzf mautrix-reddit-github-ready.tar.gz
cd mautrix-reddit
```

### Paso 2: Inicializar Git

```bash
# Inicializar repositorio
git init

# Añadir todos los archivos
git add .

# Crear primer commit
git commit -m "Initial commit: mautrix-reddit bridge

- Complete Reddit-Matrix bridge implementation
- OAuth authentication support
- Bidirectional messaging
- Multi-account support
- Docker deployment ready
- Comprehensive documentation"

# Renombrar rama a main
git branch -M main
```

### Paso 3: Crear Repositorio en GitHub

1. **Ve a** https://github.com/new

2. **Configura el repositorio:**
   - Repository name: `mautrix-reddit`
   - Description: `A Matrix-Reddit bridge for unified messaging`
   - Público o Privado: Elige según prefieras
   - **NO marques:** "Initialize with README", ".gitignore", o "license"

3. **Haz clic en** "Create repository"

### Paso 4: Conectar y Subir

```bash
# Añadir remote (reemplaza TU_USUARIO)
git remote add origin https://github.com/TU_USUARIO/mautrix-reddit.git

# Subir código
git push -u origin main
```

**Nota:** Si te pide autenticación, usa un Personal Access Token:
1. Ve a https://github.com/settings/tokens
2. "Generate new token (classic)"
3. Selecciona scopes: `repo`, `workflow`
4. Copia el token y úsalo como contraseña

### Paso 5: Configurar Permisos de Workflow

1. Ve a tu repositorio en GitHub
2. Settings → Actions → General
3. En "Workflow permissions", selecciona:
   - ✅ "Read and write permissions"
4. Click "Save"

### Paso 6: Crear el Primer Release

```bash
# Crear tag
git tag -a v0.1.0 -m "Initial release v0.1.0

Features:
- Reddit OAuth authentication
- Bidirectional messaging (Matrix ↔ Reddit)
- Multi-account support
- Docker deployment
- SQLite/PostgreSQL support
- Comprehensive documentation"

# Subir tag
git push origin v0.1.0
```

### Paso 7: Verificar Compilación

1. **Ve a** https://github.com/TU_USUARIO/mautrix-reddit/actions
2. **Verás** el workflow "Build and Release" ejecutándose
3. **Click** en él para ver el progreso en tiempo real
4. **Espera** ~5-10 minutos a que termine

### Paso 8: Verificar Release

1. **Ve a** https://github.com/TU_USUARIO/mautrix-reddit/releases
2. **Verás** el release v0.1.0
3. **Encontrarás** 5 archivos compilados:
   - `mautrix-reddit-linux-amd64.tar.gz`
   - `mautrix-reddit-linux-arm64.tar.gz`
   - `mautrix-reddit-darwin-amd64.tar.gz`
   - `mautrix-reddit-darwin-arm64.tar.gz`
   - `mautrix-reddit-windows-amd64.zip`

### Paso 9: Probar un Binario

```bash
# Descargar (reemplaza TU_USUARIO)
wget https://github.com/TU_USUARIO/mautrix-reddit/releases/download/v0.1.0/mautrix-reddit-linux-amd64.tar.gz

# Extraer
tar -xzf mautrix-reddit-linux-amd64.tar.gz

# Probar
./mautrix-reddit-linux-amd64 --help
```

¡Debería funcionar! 🎉

---

## 🐳 Configurar Docker (Opcional)

Si quieres publicar imágenes Docker en Docker Hub además de GitHub:

### Paso 1: Crear Token de Docker Hub

1. Ve a https://hub.docker.com/settings/security
2. Click "New Access Token"
3. Nombre: "GitHub Actions"
4. Copia el token

### Paso 2: Añadir Secrets en GitHub

1. Ve a tu repositorio → Settings → Secrets and variables → Actions
2. Click "New repository secret"
3. Añade:
   - Name: `DOCKERHUB_USERNAME`
   - Value: Tu usuario de Docker Hub
4. Click "Add secret"
5. Repite para:
   - Name: `DOCKERHUB_TOKEN`
   - Value: El token que copiaste

### Paso 3: Crear Nuevo Release

```bash
git tag -a v0.1.1 -m "v0.1.1 with Docker Hub support"
git push origin v0.1.1
```

Ahora las imágenes se publicarán también en Docker Hub.

---

## 🔄 Crear Releases Futuros

Cuando quieras hacer un nuevo release:

```bash
# Hacer cambios
git add .
git commit -m "feat: nueva funcionalidad"
git push

# Crear nuevo tag
git tag -a v0.2.0 -m "Version 0.2.0

- Nueva funcionalidad X
- Fix bug Y
- Mejora Z"

# Subir tag
git push origin v0.2.0
```

GitHub Actions compilará automáticamente y creará el release.

---

## ❓ Troubleshooting

### Problema: "failed to push some refs"

**Solución:**
```bash
git pull origin main --rebase
git push -u origin main
```

### Problema: "Permission denied (publickey)"

**Solución:** Usa HTTPS en lugar de SSH:
```bash
git remote set-url origin https://github.com/TU_USUARIO/mautrix-reddit.git
```

### Problema: Workflow falla en "Upload Release Asset"

**Solución:** Verifica los permisos del workflow:
- Settings → Actions → General
- "Read and write permissions" debe estar marcado

### Problema: No se crean los binarios

**Solución:** Verifica que el tag tenga formato `v*`:
```bash
# Correcto
git tag -a v0.1.0 -m "Release"

# Incorrecto (sin la 'v')
git tag -a 0.1.0 -m "Release"
```

### Problema: "refusing to allow an OAuth App to create or update workflow"

**Solución:** Crea un nuevo Personal Access Token con scope `workflow`:
1. https://github.com/settings/tokens
2. Generate new token (classic)
3. Selecciona: `repo` y `workflow`
4. Usa este token como contraseña

---

## ✅ Checklist Final

Después de seguir esta guía, deberías tener:

- [x] Proyecto en GitHub
- [x] Código subido
- [x] GitHub Actions configurado
- [x] Tag v0.1.0 creado
- [x] Release con binarios compilados
- [x] Binarios descargables para:
  - Linux (amd64, arm64)
  - macOS (amd64, arm64)
  - Windows (amd64)
- [x] Imágenes Docker en GitHub Container Registry

---

## 📚 Recursos Adicionales

- [Crear repositorio en GitHub](https://docs.github.com/en/get-started/quickstart/create-a-repo)
- [Personal Access Tokens](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Releases](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)

---

## 🎉 ¡Felicidades!

Tu proyecto está ahora en GitHub con compilación automática de binarios. Cada vez que hagas un nuevo tag, GitHub compilará automáticamente para todas las plataformas.

**Próximos pasos:**
1. Personaliza el README.github.md (reemplaza "USUARIO" con tu usuario)
2. Añade una descripción al repositorio
3. Crea topics en GitHub: `matrix`, `reddit`, `bridge`, `golang`
4. Comparte tu proyecto con la comunidad

---

**¿Necesitas ayuda?** 
- Revisa [GITHUB_ACTIONS_GUIDE.md](GITHUB_ACTIONS_GUIDE.md)
- Abre un issue en el repositorio
- Contacta en Matrix: `#mautrix-reddit:maunium.net`

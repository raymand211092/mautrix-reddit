# 🚀 Instrucciones para Subir a GitHub

## ⚡ Opción 1: Script Automático (RECOMENDADO)

```bash
# En tu máquina local, después de descomprimir el proyecto:
cd mautrix-reddit
chmod +x UPLOAD_NOW.sh
./UPLOAD_NOW.sh
```

**¡Eso es todo!** El script hace todo automáticamente:
- ✅ Inicializa Git
- ✅ Crea el commit
- ✅ Crea el repositorio en GitHub
- ✅ Sube el código
- ✅ Crea el tag v0.1.0
- ✅ GitHub Actions compila los binarios

---

## 🔧 Opción 2: Comandos Manuales

Si prefieres ejecutar los comandos uno por uno:

### 1. Inicializar Git

```bash
cd mautrix-reddit
git init
git branch -M main
```

### 2. Configurar usuario

```bash
git config user.name "raymand211092"
git config user.email "raymand211092@users.noreply.github.com"
```

### 3. Crear commit

```bash
git add .
git commit -m "Initial commit: mautrix-reddit bridge"
```

### 4. Crear repositorio en GitHub

```bash
curl -X POST \
  -H "Authorization: token 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" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/user/repos \
  -d '{"name":"mautrix-reddit","description":"A Matrix-Reddit bridge for unified messaging","private":false}'
```

O crea el repositorio manualmente en: https://github.com/new
- Nombre: `mautrix-reddit`
- Descripción: `A Matrix-Reddit bridge for unified messaging`
- Público
- NO inicialices con README

### 5. Añadir remote y hacer push

```bash
git remote add origin https://raymand211092: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@github.com/raymand211092/mautrix-reddit.git

git push -u origin main
```

### 6. Crear tag para el primer release

```bash
git tag -a v0.1.0 -m "Initial release v0.1.0"
git push origin v0.1.0
```

---

## ✅ Verificar que Funcionó

1. **Código subido:**
   https://github.com/raymand211092/mautrix-reddit

2. **GitHub Actions ejecutándose:**
   https://github.com/raymand211092/mautrix-reddit/actions

3. **Releases (después de 5-10 min):**
   https://github.com/raymand211092/mautrix-reddit/releases

---

## 🎯 Qué Esperar

Después de ejecutar el script o los comandos:

### Inmediatamente:
- ✅ Código visible en GitHub
- ✅ GitHub Actions empieza a compilar

### Después de 5-10 minutos:
- ✅ Release v0.1.0 creado
- ✅ 5 binarios compilados disponibles:
  - mautrix-reddit-linux-amd64.tar.gz
  - mautrix-reddit-linux-arm64.tar.gz
  - mautrix-reddit-darwin-amd64.tar.gz (macOS Intel)
  - mautrix-reddit-darwin-arm64.tar.gz (macOS M1/M2)
  - mautrix-reddit-windows-amd64.zip

### Imágenes Docker:
- ✅ ghcr.io/raymand211092/mautrix-reddit:latest
- ✅ ghcr.io/raymand211092/mautrix-reddit:v0.1.0

---

## 🔍 Monitorear el Progreso

```bash
# Ver el estado del workflow
open https://github.com/raymand211092/mautrix-reddit/actions

# O con GitHub CLI (si está instalado)
gh run list --repo raymand211092/mautrix-reddit
gh run watch
```

---

## 🆘 Si Algo Sale Mal

### Error: "repository already exists"
Está bien, continúa con el paso 5 (git push)

### Error: "failed to push"
```bash
# Verificar el remote
git remote -v

# Si no está configurado correctamente:
git remote remove origin
git remote add origin https://raymand211092:TU_TOKEN@github.com/raymand211092/mautrix-reddit.git
git push -u origin main
```

### Error: Workflow no se ejecuta
1. Ve a Settings → Actions → General
2. Asegúrate de que "Allow all actions" esté seleccionado
3. Cambia "Workflow permissions" a "Read and write permissions"
4. Guarda cambios

---

## 🎉 ¡Listo!

Una vez completado, tendrás:
- ✅ Proyecto en GitHub
- ✅ Binarios compilados para 5 plataformas
- ✅ Imágenes Docker publicadas
- ✅ Todo automático para futuros releases

**Para crear nuevos releases en el futuro:**
```bash
git tag -a v0.2.0 -m "Version 0.2.0"
git push origin v0.2.0
# ¡GitHub compila automáticamente!
```

---

## 📝 Notas de Seguridad

**IMPORTANTE:** El token incluido en los scripts tiene permisos de `repo` y `workflow`. Considera:

1. **Revocar el token** después de usarlo si no lo necesitas más
2. **Crear un nuevo token** con permisos mínimos para uso continuo
3. **No compartir** estos scripts con el token incluido

Para revocar el token:
https://github.com/settings/tokens

---

**¿Necesitas ayuda?** Revisa los logs en:
https://github.com/raymand211092/mautrix-reddit/actions

# ⚡ Быстрая публикация на GitHub

Простая инструкция для публикации модуля за 5 минут.

## 📋 Предварительные требования

- [ ] Аккаунт на GitHub
- [ ] Git установлен
- [ ] Код готов к публикации

---

## 🚀 Шаги публикации

### 1️⃣ Создайте репозиторий на GitHub

1. Откройте https://github.com/new
2. Назовите: `fingerprint-injector`
3. Выберите **Public** или **Private**
4. **НЕ** добавляйте README, .gitignore или лицензию
5. Нажмите **Create repository**

### 2️⃣ Обновите go.mod

Откройте `go.mod` и измените первую строку:

```go
// Замените YOUR_USERNAME на ваш GitHub username
module github.com/YOUR_USERNAME/fingerprint-injector
```

Пример: `module github.com/elaine/fingerprint-injector`

### 3️⃣ Выполните команды

```bash
# 1. Инициализация Git
git init
git add .
git commit -m "Initial commit: fingerprint-injector v1.0.0"

# 2. Подключите GitHub (замените YOUR_USERNAME)
git remote add origin https://github.com/YOUR_USERNAME/fingerprint-injector.git
git branch -M main

# 3. Загрузите код
git push -u origin main

# 4. Создайте тег версии
git tag v1.0.0
git push origin v1.0.0
```

### ✅ Готово!

Теперь модуль опубликован и его можно устанавливать:

```bash
go get github.com/YOUR_USERNAME/fingerprint-injector@v1.0.0
```

---

## 🔄 Обновление модуля

Когда вы внесли изменения и хотите выпустить новую версию:

### Вариант A: Используйте скрипт (рекомендуется)

**Windows:**

```powershell
.\publish.ps1
```

**Linux/Mac:**

```bash
chmod +x publish.sh
./publish.sh
```

### Вариант B: Вручную

```bash
# 1. Коммит изменений
git add .
git commit -m "Описание изменений"

# 2. Тест (опционально)
go test ./...

# 3. Загрузите
git push

# 4. Создайте новый тег
git tag v1.1.0
git push origin v1.1.0
```

---

## 📦 Использование в других проектах

### Создайте новый проект:

```bash
mkdir my-project
cd my-project
go mod init my-project
```

### Установите модуль:

```bash
go get github.com/YOUR_USERNAME/fingerprint-injector@latest
```

### Используйте в коде:

```go
package main

import (
    "context"
    "github.com/chromedp/chromedp"
    fp "github.com/YOUR_USERNAME/fingerprint-injector"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Используем ваш модуль!
    fingerprint := fp.GenerateRandomFingerprint()
    injector := fp.NewInjector(fingerprint)

    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://example.com"),
    )
}
```

### Запустите:

```bash
go run main.go
```

---

## 🎯 Проверка публикации

Проверьте что модуль доступен:

```bash
go list -m github.com/YOUR_USERNAME/fingerprint-injector@latest
```

Посмотрите все версии:

```bash
go list -m -versions github.com/YOUR_USERNAME/fingerprint-injector
```

---

## ❓ Troubleshooting

### Ошибка: "cannot find module"

```bash
# Убедитесь что тег создан
git tag

# Если тега нет, создайте
git tag v1.0.0
git push origin v1.0.0

# Подождите 1-2 минуты для индексации Go
```

### Ошибка: "permission denied"

```bash
# Проверьте что вы залогинены в GitHub
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Для HTTPS используйте Personal Access Token
# GitHub → Settings → Developer settings → Personal access tokens
```

### Ошибка: "tag already exists"

```bash
# Удалите старый тег локально
git tag -d v1.0.0

# Удалите на GitHub
git push origin :refs/tags/v1.0.0

# Создайте заново
git tag v1.0.0
git push origin v1.0.0
```

---

## 💡 Tips

### 1. Версионирование

Следуйте Semantic Versioning:

- `v1.0.0` → `v1.0.1` - bug fixes (патч)
- `v1.0.0` → `v1.1.0` - новые функции (минор)
- `v1.0.0` → `v2.0.0` - breaking changes (мажор)

### 2. Обновите CHANGELOG.md

После каждого релиза обновляйте:

```markdown
## [1.1.0] - 2024-10-12

### Added

- Smart fingerprint generator
- iOS support

### Fixed

- Viewport bug on mobile
```

### 3. Создайте GitHub Release

На странице репозитория:

1. Releases → Create a new release
2. Tag: v1.0.0
3. Title: "v1.0.0 - Initial Release"
4. Description: скопируйте из CHANGELOG
5. Publish release

---

## 🎉 Готово!

Теперь ваш модуль:

- ✅ Опубликован на GitHub
- ✅ Доступен для установки через `go get`
- ✅ Можно использовать в любом проекте

---

## 📚 Дополнительно

Подробное руководство: [PUBLISHING_GUIDE.md](PUBLISHING_GUIDE.md)

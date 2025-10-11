# 📦 Руководство по публикации модуля

Это руководство покажет, как опубликовать fingerprint-injector для использования в других проектах.

## 🎯 Способы публикации

### 1. GitHub (Рекомендуется) ⭐

### 2. GitLab / Bitbucket

### 3. Private модуль

### 4. Локальный путь

---

## 📤 Способ 1: Публикация на GitHub (Публичный)

Это **самый простой** и **популярный** способ.

### Шаг 1: Создайте репозиторий на GitHub

1. Откройте https://github.com/new
2. Заполните:
   - **Repository name**: `fingerprint-injector` (или любое имя)
   - **Description**: "Browser fingerprint injection library for chromedp (Golang)"
   - **Visibility**: Public ✅
3. **НЕ** инициализируйте с README (у нас уже есть)
4. Нажмите "Create repository"

### Шаг 2: Измените go.mod

Замените первую строку в `go.mod`:

```go
// Было:
module github.com/vitaliitsarov/fingerprint-injector-go

// Должно быть:
module github.com/ВАШ_USERNAME/fingerprint-injector
```

Например, если ваш GitHub username `vitaliitsarov`:

```go
module github.com/vitaliitsarov/fingerprint-injector
```

### Шаг 3: Инициализируйте Git и загрузите

```bash
# Инициализация
git init
git add .
git commit -m "Initial commit: fingerprint-injector v1.0.0"

# Добавьте remote (замените на свой URL)
git remote add origin https://github.com/ВАШ_USERNAME/fingerprint-injector.git

# Загрузите
git branch -M main
git push -u origin main
```

### Шаг 4: Создайте тег версии (важно!)

```bash
# Создайте тег версии
git tag v1.0.0
git push origin v1.0.0
```

### ✅ Готово! Теперь можно использовать:

```bash
# В другом проекте
go get github.com/ВАШ_USERNAME/fingerprint-injector@v1.0.0
```

---

## 🔒 Способ 2: Приватный репозиторий на GitHub

Если модуль приватный:

### Шаг 1: Создайте приватный репозиторий

1. GitHub → New repository
2. **Visibility**: Private ✅
3. Создайте

### Шаг 2: Настройте Git credentials

```bash
# Настройте Git для приватных репозиториев
git config --global url."https://YOUR_GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"
```

Где `YOUR_GITHUB_TOKEN` - это Personal Access Token:

1. GitHub → Settings → Developer settings → Personal access tokens
2. Generate new token (classic)
3. Дайте доступ к `repo`

### Шаг 3: Загрузите код

```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/ВАШ_USERNAME/fingerprint-injector.git
git branch -M main
git push -u origin main
git tag v1.0.0
git push origin v1.0.0
```

### Шаг 4: Используйте в другом проекте

```bash
# Укажите токен в переменной окружения
export GOPRIVATE=github.com/ВАШ_USERNAME/fingerprint-injector

# Установите
go get github.com/ВАШ_USERNAME/fingerprint-injector@v1.0.0
```

---

## 💻 Способ 3: Локальный путь (для разработки)

Если вы хотите использовать модуль локально без публикации:

### Вариант A: Replace в go.mod

В **вашем проекте** (не в fingerprint-injector):

```bash
# Создайте новый проект
mkdir my-project
cd my-project
go mod init my-project

# Добавьте replace директиву
go mod edit -replace github.com/vitaliitsarov/fingerprint-injector-go=C:/Users/Elaine/Desktop/test
```

Теперь в `go.mod` вашего проекта:

```go
module my-project

go 1.21

require github.com/vitaliitsarov/fingerprint-injector-go v0.0.0

replace github.com/vitaliitsarov/fingerprint-injector-go => C:/Users/Elaine/Desktop/test
```

И используйте:

```go
package main

import (
    "context"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    fingerprint := fp.GenerateRandomFingerprint()
    // ...
}
```

### Вариант B: Используйте file://

```bash
go get github.com/vitaliitsarov/fingerprint-injector-go@v0.0.0
go mod edit -replace github.com/vitaliitsarov/fingerprint-injector-go=file://C:/Users/Elaine/Desktop/test
```

---

## 🚀 Полный пример: От создания до использования

### 1️⃣ Подготовка модуля

```bash
cd C:/Users/Elaine/Desktop/test

# Измените go.mod
# module github.com/YOUR_USERNAME/fingerprint-injector

# Коммит
git init
git add .
git commit -m "Initial commit: fingerprint-injector v1.0.0"
```

### 2️⃣ Публикация на GitHub

```bash
# Создайте репозиторий на GitHub (через веб-интерфейс)
# Затем:

git remote add origin https://github.com/YOUR_USERNAME/fingerprint-injector.git
git branch -M main
git push -u origin main

# Создайте тег
git tag v1.0.0
git push origin v1.0.0
```

### 3️⃣ Использование в новом проекте

```bash
# Создайте новый проект
mkdir my-scraper
cd my-scraper
go mod init my-scraper

# Установите модуль
go get github.com/YOUR_USERNAME/fingerprint-injector@v1.0.0
```

### 4️⃣ Код в новом проекте

```go
// main.go
package main

import (
    "context"
    "log"

    "github.com/chromedp/chromedp"
    fp "github.com/YOUR_USERNAME/fingerprint-injector"
)

func main() {
    // Используем ваш модуль!
    fingerprint := fp.GenerateRandomFingerprint()

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    injector := fp.NewInjector(fingerprint)

    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://example.com"),
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Success! Используем fingerprint-injector как модуль")
}
```

```bash
# Запуск
go run main.go
```

---

## 🔄 Обновление модуля

### После изменений в модуле:

```bash
cd C:/Users/Elaine/Desktop/test

# Коммит изменений
git add .
git commit -m "Add new features"
git push

# Создайте новый тег
git tag v1.1.0
git push origin v1.1.0
```

### В проекте, который использует модуль:

```bash
# Обновите до новой версии
go get github.com/YOUR_USERNAME/fingerprint-injector@v1.1.0

# Или обновите до последней
go get -u github.com/YOUR_USERNAME/fingerprint-injector
```

---

## 📋 Чеклист публикации

### Перед публикацией:

- [ ] Обновите `go.mod` с правильным путём модуля
- [ ] Проверьте что все файлы добавлены в git
- [ ] README.md содержит инструкции по установке
- [ ] Код протестирован (`go test ./...`)
- [ ] Нет чувствительных данных (токенов, паролей)

### Публикация:

- [ ] Создан репозиторий на GitHub/GitLab
- [ ] Код загружен (`git push`)
- [ ] Создан тег версии (`git tag v1.0.0`)
- [ ] Тег загружен (`git push origin v1.0.0`)

### После публикации:

- [ ] Протестируйте установку в чистом проекте
- [ ] Обновите документацию с правильным путём установки
- [ ] Добавьте badge в README (опционально)

---

## 🎨 Обновите README после публикации

После публикации обновите инструкцию по установке в README.md:

````markdown
## 📦 Установка

```bash
go get github.com/YOUR_USERNAME/fingerprint-injector@latest
```
````

## 🚀 Быстрый старт

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

    fingerprint := fp.NewChrome119Windows11()
    injector := fp.NewInjector(fingerprint)

    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://example.com"),
    )
}
```

````

---

## 🔍 Проверка публикации

После публикации проверьте:

```bash
# Проверьте что модуль доступен
go list -m github.com/YOUR_USERNAME/fingerprint-injector@latest

# Посмотрите доступные версии
go list -m -versions github.com/YOUR_USERNAME/fingerprint-injector
````

---

## 💡 Best Practices

### 1. Версионирование (Semantic Versioning)

- `v1.0.0` - Initial release
- `v1.0.1` - Bug fixes
- `v1.1.0` - New features (backwards compatible)
- `v2.0.0` - Breaking changes

### 2. Changelog

Ведите CHANGELOG.md:

```markdown
## [1.1.0] - 2024-10-12

### Added

- New device database with 30+ GPUs
- Smart fingerprint generator
- iOS support

### Fixed

- Viewport issues on mobile devices
```

### 3. GitHub Releases

Создайте Release на GitHub:

1. GitHub → Releases → Create new release
2. Tag: v1.0.0
3. Title: "v1.0.0 - Initial Release"
4. Описание: скопируйте из CHANGELOG.md

---

## 🌟 Дополнительно: Badges для README

Добавьте красивые badges:

```markdown
![Go Version](https://img.shields.io/github/go-mod/go-version/YOUR_USERNAME/fingerprint-injector)
![Release](https://img.shields.io/github/v/release/YOUR_USERNAME/fingerprint-injector)
![License](https://img.shields.io/github/license/YOUR_USERNAME/fingerprint-injector)
![Stars](https://img.shields.io/github/stars/YOUR_USERNAME/fingerprint-injector)
```

---

## ❓ FAQ

### Q: Нужно ли каждый раз создавать тег?

A: Да, для каждой версии нужен свой тег. Go использует теги для версионирования.

### Q: Можно ли использовать без GitHub?

A: Да, можете использовать GitLab, Bitbucket или даже свой сервер с git.

### Q: Как сделать приватный модуль для команды?

A: Используйте приватный репозиторий + настройте GOPRIVATE или используйте корпоративный proxy (Athens, JFrog).

### Q: Можно ли не публиковать, а использовать локально?

A: Да, используйте `replace` директиву в go.mod (см. Способ 3).

---

## 📚 Дополнительные ресурсы

- [Go Modules Reference](https://go.dev/ref/mod)
- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules)
- [Semantic Versioning](https://semver.org/)

---

**Готово!** Теперь ваш модуль можно использовать в любом Go проекте! 🎉

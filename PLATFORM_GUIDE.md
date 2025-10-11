# 🎯 Руководство по выбору платформ

Это руководство покажет, как выбирать и использовать разные платформы в fingerprint-injector.

## 📱 Доступные платформы

| Платформа      | Функция                   | Браузер    | Описание                    |
| -------------- | ------------------------- | ---------- | --------------------------- |
| **Windows 11** | `NewChrome119Windows11()` | Chrome 119 | Стандартная Windows система |
| **Linux**      | `NewChrome119Linux()`     | Chrome 119 | Linux десктоп               |
| **MacOS**      | `NewChrome119MacOS()`     | Chrome 119 | Apple MacOS                 |
| **Android**    | `NewChrome119Android()`   | Chrome 119 | Android мобильный           |
| **iOS Safari** | `NewSafari17iOS()`        | Safari 17  | iPhone с Safari             |
| **iOS Chrome** | `NewChrome119iOS()`       | Chrome 119 | iPhone с Chrome             |

---

## 🚀 Способы выбора платформы

### Способ 1: Прямой выбор в коде

```go
package main

import (
    "context"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Выберите нужную платформу, раскомментировав строку:

    fingerprint := fp.NewChrome119Windows11()  // Windows
    // fingerprint := fp.NewChrome119Linux()      // Linux
    // fingerprint := fp.NewChrome119MacOS()      // MacOS
    // fingerprint := fp.NewChrome119Android()    // Android
    // fingerprint := fp.NewSafari17iOS()         // iOS Safari
    // fingerprint := fp.NewChrome119iOS()        // iOS Chrome

    injector := fp.NewInjector(fingerprint)

    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://whoer.net"),
    )

    select {}
}
```

### Способ 2: Через параметр командной строки

```bash
# Используйте готовый пример platform-selector
cd examples/platform-selector

# Windows
go run main.go -platform=windows

# Linux
go run main.go -platform=linux

# MacOS
go run main.go -platform=macos

# Android
go run main.go -platform=android

# iOS Safari
go run main.go -platform=ios

# iOS Chrome
go run main.go -platform=ios-chrome
```

### Способ 3: Через Makefile

```bash
# Windows (по умолчанию)
make run-platform

# Linux
make run-platform PLATFORM=linux

# MacOS
make run-platform PLATFORM=macos

# Android
make run-platform PLATFORM=android

# iOS Safari
make run-platform PLATFORM=ios

# iOS Chrome
make run-platform PLATFORM=ios-chrome
```

### Способ 4: Динамический выбор через переменную

```go
package main

import (
    "context"
    "os"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func getPlatformFingerprint(platform string) *fp.Fingerprint {
    switch platform {
    case "windows":
        return fp.NewChrome119Windows11()
    case "linux":
        return fp.NewChrome119Linux()
    case "macos":
        return fp.NewChrome119MacOS()
    case "android":
        return fp.NewChrome119Android()
    case "ios":
        return fp.NewSafari17iOS()
    case "ios-chrome":
        return fp.NewChrome119iOS()
    default:
        return fp.NewChrome119Windows11()
    }
}

func main() {
    // Получаем платформу из переменной окружения
    platform := os.Getenv("PLATFORM")
    if platform == "" {
        platform = "windows"
    }

    fingerprint := getPlatformFingerprint(platform)

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    injector := fp.NewInjector(fingerprint)

    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://whoer.net"),
    )

    select {}
}
```

Использование:

```bash
# Linux
PLATFORM=linux go run main.go

# MacOS
PLATFORM=macos go run main.go

# iOS
PLATFORM=ios go run main.go
```

---

## 📊 Сравнение характеристик платформ

### Desktop платформы

| Параметр   | Windows       | Linux              | MacOS        |
| ---------- | ------------- | ------------------ | ------------ |
| Platform   | Win32         | Linux x86_64       | MacIntel     |
| Разрешение | 1920×1080     | 1920×1080          | 2560×1440    |
| DPI        | 1.0           | 1.0                | 2.0          |
| CPU Cores  | 8             | 12                 | 10           |
| RAM        | 8 GB          | 16 GB              | 8 GB         |
| GPU        | Intel UHD 630 | NVIDIA GTX 1080 Ti | Apple M1 Pro |

### Mobile платформы

| Параметр   | Android      | iOS Safari | iOS Chrome |
| ---------- | ------------ | ---------- | ---------- |
| Platform   | Linux armv8l | iPhone     | iPhone     |
| Разрешение | 412×915      | 390×844    | 390×844    |
| DPI        | 2.625        | 3.0        | 3.0        |
| CPU Cores  | 8            | 6          | 6          |
| RAM        | 8 GB         | 6 GB       | 6 GB       |
| GPU        | Adreno 730   | Apple A16  | Apple A16  |

---

## 🎨 Примеры использования

### Пример 1: Web Scraping с ротацией платформ

```go
package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func scrapeWithPlatform(url string, platform string) error {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    var fingerprint *fp.Fingerprint
    switch platform {
    case "windows":
        fingerprint = fp.NewChrome119Windows11()
    case "linux":
        fingerprint = fp.NewChrome119Linux()
    case "macos":
        fingerprint = fp.NewChrome119MacOS()
    }

    injector := fp.NewInjector(fingerprint)

    var content string
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate(url),
        chromedp.OuterHTML("body", &content),
    )

    fmt.Printf("Scraped %d bytes from %s using %s\n",
        len(content), url, platform)

    return err
}

func main() {
    urls := []string{
        "https://example.com/1",
        "https://example.com/2",
        "https://example.com/3",
    }

    platforms := []string{"windows", "linux", "macos"}

    for i, url := range urls {
        platform := platforms[i%len(platforms)]
        scrapeWithPlatform(url, platform)
    }
}
```

### Пример 2: A/B тестирование на разных платформах

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func testOnPlatform(url string, platformName string, wg *sync.WaitGroup) {
    defer wg.Done()

    platforms := map[string]*fp.Fingerprint{
        "windows": fp.NewChrome119Windows11(),
        "linux":   fp.NewChrome119Linux(),
        "macos":   fp.NewChrome119MacOS(),
        "android": fp.NewChrome119Android(),
        "ios":     fp.NewSafari17iOS(),
    }

    fingerprint := platforms[platformName]

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    injector := fp.NewInjector(fingerprint)

    var title string
    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate(url),
        chromedp.Title(&title),
    )

    fmt.Printf("[%s] Title: %s\n", platformName, title)
}

func main() {
    url := "https://example.com"
    platforms := []string{"windows", "linux", "macos", "android", "ios"}

    var wg sync.WaitGroup
    for _, platform := range platforms {
        wg.Add(1)
        go testOnPlatform(url, platform, &wg)
    }
    wg.Wait()
}
```

### Пример 3: Mobile vs Desktop проверка

```go
package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func checkMobileDesktop(url string) {
    platforms := map[string]*fp.Fingerprint{
        "Desktop (Windows)": fp.NewChrome119Windows11(),
        "Desktop (MacOS)":   fp.NewChrome119MacOS(),
        "Mobile (Android)":  fp.NewChrome119Android(),
        "Mobile (iOS)":      fp.NewSafari17iOS(),
    }

    for name, fingerprint := range platforms {
        ctx, cancel := chromedp.NewContext(context.Background())

        injector := fp.NewInjector(fingerprint)

        var isMobile bool
        chromedp.Run(ctx,
            injector.ApplyAll(ctx),
            chromedp.Navigate(url),
            chromedp.Evaluate(`/Mobi|Android/i.test(navigator.userAgent)`, &isMobile),
        )

        fmt.Printf("%s: Mobile=%v\n", name, isMobile)

        cancel()
    }
}

func main() {
    checkMobileDesktop("https://example.com")
}
```

---

## 🔧 Кастомизация платформ

Вы можете взять любой preset и модифицировать его:

```go
// Начинаем с Windows preset
fingerprint := fp.NewChrome119Windows11()

// Изменяем на русский язык
fingerprint.Language = "ru-RU"
fingerprint.Languages = []string{"ru-RU", "ru", "en"}

// Меняем timezone на Москву
fingerprint.Timezone = &fp.Timezone{
    ID:     "Europe/Moscow",
    Offset: -180,
}

// Увеличиваем разрешение
fingerprint.Screen.Width = 2560
fingerprint.Screen.Height = 1440

// Отключаем WebRTC
fingerprint.WebRTC.Disable = true

// Используем модифицированный fingerprint
injector := fp.NewInjector(fingerprint)
```

---

## 🧪 Тестирование платформ

Лучшие сайты для проверки fingerprint:

### Universal тесты

- https://whoer.net - Общая проверка
- https://browserleaks.com - Детальная проверка
- https://amiunique.org - Уникальность fingerprint

### Desktop тесты

- https://www.whatismybrowser.com
- https://browserleaks.com/javascript

### Mobile тесты

- https://www.whatismybrowser.com/detect/is-this-a-mobile-device
- https://mobiletest.me

### Bot detection

- https://bot.sannysoft.com
- https://arh.antoinevastel.com/bots/areyouheadless

---

## 💡 Tips & Tricks

### 1. Соответствие платформы и IP

Если используете прокси, убедитесь, что fingerprint соответствует локации:

```go
// Прокси в США -> Windows US
fingerprint := fp.NewChrome119Windows11()
fingerprint.Timezone.ID = "America/New_York"
fingerprint.Language = "en-US"

// Прокси в Германии -> Windows EU
fingerprint := fp.NewChrome119Windows11()
fingerprint.Timezone.ID = "Europe/Berlin"
fingerprint.Language = "de-DE"
fingerprint.Languages = []string{"de-DE", "de", "en"}
```

### 2. Mobile для социальных сетей

Для работы с мобильными версиями соцсетей:

```go
// Instagram/TikTok лучше с iOS
fingerprint := fp.NewSafari17iOS()

// Twitter/Facebook работают со всеми
fingerprint := fp.NewChrome119Android()
```

### 3. Desktop для сложных сайтов

Для банков, магазинов используйте desktop:

```go
fingerprint := fp.NewChrome119Windows11()
// или
fingerprint := fp.NewChrome119MacOS()
```

---

## 📚 Дополнительные ресурсы

- [README.md](README.md) - Основная документация
- [QUICKSTART.md](QUICKSTART.md) - Быстрый старт
- [examples/platform-selector/](examples/platform-selector/) - Готовый пример
- [INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md) - Продвинутые сценарии

---

**Версия**: 1.0.0  
**Дата**: 2024-10-11

# Fingerprint Injector для Golang (chromedp)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

Библиотека для инжектирования и модификации browser fingerprint при использовании chromedp в Golang. Аналог популярных Node.js решений типа `fingerprint-injector`.

[English](README.md) | **Русский**

## 🎯 Возможности

- ✅ Полная модификация browser fingerprint
- ✅ Поддержка всех основных параметров (User-Agent, WebGL, Canvas, Screen, Timezone и др.)
- ✅ Защита от Canvas fingerprinting с помощью шума
- ✅ Отключение/модификация WebRTC
- ✅ Скрытие признаков автоматизации (webdriver, chrome.runtime)
- ✅ Готовые пресеты для разных ОС и браузеров
- ✅ Простой API
- ✅ Полная интеграция с chromedp

## 📦 Установка

```bash
go get github.com/vitaliitsarov/fingerprint-injector-go
```

## 🚀 Быстрый старт

### Базовое использование

```go
package main

import (
    "context"
    "log"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Используем готовый preset
    fingerprint := fp.NewChrome119Windows11()

    // Создаем инжектор
    injector := fp.NewInjector(fingerprint)

    // Применяем fingerprint
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://example.com"),
    )

    if err != nil {
        log.Fatal(err)
    }
}
```

### Кастомный fingerprint

```go
fingerprint := &fp.Fingerprint{
    UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)...",
    Platform:  "Win32",
    Vendor:    "Google Inc.",
    Language:  "ru-RU",
    Languages: []string{"ru-RU", "ru", "en"},
    Screen: &fp.Screen{
        Width:            1920,
        Height:           1080,
        ColorDepth:       24,
        DevicePixelRatio: 1.0,
    },
    Timezone: &fp.Timezone{
        ID:     "Europe/Moscow",
        Offset: -180,
    },
    WebGL: &fp.WebGL{
        Vendor:   "Google Inc. (NVIDIA)",
        Renderer: "ANGLE (NVIDIA GeForce RTX 3080)",
    },
    Canvas: &fp.Canvas{
        Noise: 0.02,
    },
    WebRTC: &fp.WebRTC{
        Disable: true,
    },
    HardwareConcurrency: 16,
    DeviceMemory:        32,
}

injector := fp.NewInjector(fingerprint)
```

## 🎨 Готовые пресеты

```go
// Chrome 119 на Windows 11
fp.NewChrome119Windows11()

// Chrome 119 на MacOS
fp.NewChrome119MacOS()

// Chrome 119 на Linux
fp.NewChrome119Linux()

// Chrome 119 на Android
fp.NewChrome119Android()
```

## 📋 Основные параметры

### Основные

- `UserAgent` - User-Agent строка
- `Platform` - Платформа (Win32, MacIntel, Linux x86_64)
- `Vendor` - Производитель браузера
- `Language` - Основной язык
- `Languages` - Список языков
- `HardwareConcurrency` - Количество процессорных ядер
- `DeviceMemory` - Объем памяти устройства (GB)

### Экран

```go
Screen: &fp.Screen{
    Width:            1920,
    Height:           1080,
    AvailWidth:       1920,
    AvailHeight:      1040,
    ColorDepth:       24,
    PixelDepth:       24,
    DevicePixelRatio: 1.0,
}
```

### WebGL

```go
WebGL: &fp.WebGL{
    Vendor:           "Google Inc. (NVIDIA)",
    Renderer:         "ANGLE (NVIDIA GeForce RTX 3080)",
    UnmaskedVendor:   "NVIDIA Corporation",
    UnmaskedRenderer: "NVIDIA GeForce RTX 3080",
}
```

### Canvas защита

```go
Canvas: &fp.Canvas{
    Noise: 0.02, // 0.0 - 1.0
}
```

### WebRTC

```go
WebRTC: &fp.WebRTC{
    Disable:  true,
    PublicIP: "8.8.8.8",
    LocalIP:  "192.168.1.100",
}
```

### Временная зона

```go
Timezone: &fp.Timezone{
    ID:     "Europe/Moscow",
    Offset: -180,
}
```

### Батарея

```go
Battery: &fp.Battery{
    Charging:        false,
    ChargingTime:    0,
    DischargingTime: 18000,
    Level:           0.75,
}
```

## 🛡️ Stealth режим

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("disable-blink-features", "AutomationControlled"),
    chromedp.Flag("exclude-switches", "enable-automation"),
    chromedp.Flag("disable-extensions", false),
    chromedp.Flag("disable-dev-shm-usage", true),
    chromedp.UserDataDir("./chrome-data"),
)

allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
defer cancel()

ctx, cancel := chromedp.NewContext(allocCtx)
defer cancel()

// Fingerprint с усиленной защитой
fingerprint := fp.NewChrome119Windows11()
fingerprint.WebRTC.Disable = true
fingerprint.Canvas.Noise = 0.05

injector := fp.NewInjector(fingerprint)
```

## 📖 Примеры

В папке `examples/` находятся полные примеры:

```bash
# Базовый пример
cd examples/basic && go run main.go

# Кастомный fingerprint
cd examples/custom && go run main.go

# Stealth режим
cd examples/stealth && go run main.go

# С прокси
cd examples/with-proxy && go run main.go

# Множественные сессии
cd examples/multi-session && go run main.go

# Случайный fingerprint
cd examples/random && go run main.go
```

## 🔧 Makefile команды

```bash
make deps          # Установить зависимости
make test          # Запустить тесты
make run-basic     # Запустить базовый пример
make run-custom    # Запустить кастомный пример
make run-stealth   # Запустить stealth режим
make run-proxy     # Запустить пример с прокси
make run-multi     # Запустить множественные сессии
make run-random    # Запустить случайный fingerprint
make clean         # Очистить временные файлы
```

## 🧪 Тестирование

Проверить качество fingerprint можно на следующих сайтах:

- https://browserleaks.com/canvas
- https://browserleaks.com/webgl
- https://bot.sannysoft.com/
- https://arh.antoinevastel.com/bots/areyouheadless
- https://deviceandbrowserinfo.com/are_you_a_bot
- https://amiunique.org/

## 📚 Документация

- [QUICKSTART.md](QUICKSTART.md) - Быстрый старт за 5 минут
- [ARCHITECTURE.md](ARCHITECTURE.md) - Архитектура проекта
- [INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md) - Примеры интеграции
- [PROJECT_OVERVIEW.md](PROJECT_OVERVIEW.md) - Обзор проекта
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - Структура проекта

## 🎲 Генерация случайных fingerprint

```go
// Полностью случайный fingerprint
fingerprint := fp.GenerateRandomFingerprint()

// Или на основе существующего с вариациями
base := fp.NewChrome119Windows11()
fingerprint := fp.RandomizeFingerprint(base)

// Случайное разрешение экрана
screen := fp.RandomResolution()

// Случайная временная зона
timezone := fp.RandomTimezone()
```

## 🔧 API Reference

### Создание инжектора

```go
func NewInjector(fingerprint *Fingerprint) *Injector
```

### Методы Injector

- `ApplyAll(ctx context.Context)` - Применить все настройки fingerprint
- `Inject(ctx context.Context)` - Инжектировать JavaScript код
- `SetUserAgentOverride(ctx context.Context)` - Установить User-Agent через CDP
- `SetTimezoneOverride(ctx context.Context)` - Установить Timezone через CDP
- `GetInjectionScript()` - Получить JavaScript код для инжектирования

### Создание Fingerprint

```go
// Из preset
fp := fp.NewChrome119Windows11()

// Дефолтный
fp := fp.NewDefaultFingerprint()

// Кастомный
fp := &fp.Fingerprint{ /* ... */ }

// Случайный
fp := fp.GenerateRandomFingerprint()
```

## 🤝 Вклад

Пул реквесты приветствуются! Для крупных изменений, пожалуйста, сначала откройте issue для обсуждения.

См. [CONTRIBUTING.md](CONTRIBUTING.md) для подробностей.

## 📝 Лицензия

MIT - см. файл [LICENSE](LICENSE)

## ⚠️ Дисклеймер

Этот инструмент предназначен только для легитимных целей, таких как:

- Тестирование защиты от ботов
- Автоматизация тестирования
- Исследование browser fingerprinting

Не используйте для обхода систем защиты или других незаконных действий.

## 🌟 Благодарности

Вдохновлено проектами:

- puppeteer-extra-plugin-stealth
- fingerprint-injector (Node.js)
- FingerprintJS

## 📊 Статистика проекта

- Языков: Go, JavaScript
- Строк кода: ~4700
- Примеров: 6
- Документации: 7 файлов
- Тестов: 15+
- Presets: 4

---

Made with ❤️ for the Go community

**Версия**: 1.0.0  
**Дата**: 2024-10-11

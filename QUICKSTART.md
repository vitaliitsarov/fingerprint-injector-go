# Быстрый старт

Это руководство поможет вам быстро начать работу с fingerprint-injector.

## Установка

```bash
# Создайте новый Go проект
mkdir my-project
cd my-project
go mod init my-project

# Установите fingerprint-injector
go get github.com/vitaliitsarov/fingerprint-injector-go
go get github.com/chromedp/chromedp
```

## Простейший пример за 5 минут

Создайте файл `main.go`:

```go
package main

import (
    "context"
    "log"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    // Создаём контекст
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Используем готовый preset для Windows 11
    fingerprint := fp.NewChrome119Windows11()
    injector := fp.NewInjector(fingerprint)

    // Применяем fingerprint и открываем сайт
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://www.whatismybrowser.com/"),
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Готово! Проверьте браузер.")
    select {} // Держим программу запущенной
}
```

Запустите:

```bash
go run main.go
```

## Основные сценарии использования

### 1. Использование готовых пресетов

```go
// Windows 11
fp := fp.NewChrome119Windows11()

// MacOS
fp := fp.NewChrome119MacOS()

// Linux
fp := fp.NewChrome119Linux()

// Android
fp := fp.NewChrome119Android()
```

### 2. Создание кастомного fingerprint

```go
fingerprint := &fp.Fingerprint{
    UserAgent: "Mozilla/5.0...",
    Platform:  "Win32",
    Language:  "ru-RU",
    Languages: []string{"ru-RU", "ru"},
    Screen: &fp.Screen{
        Width:  1920,
        Height: 1080,
    },
    Timezone: &fp.Timezone{
        ID:     "Europe/Moscow",
        Offset: -180,
    },
    WebRTC: &fp.WebRTC{
        Disable: true, // Отключить WebRTC
    },
    HardwareConcurrency: 8,
    DeviceMemory:        16,
}
```

### 3. Генерация случайного fingerprint

```go
// Полностью случайный fingerprint
fingerprint := fp.GenerateRandomFingerprint()

// Или на основе существующего с вариациями
base := fp.NewChrome119Windows11()
fingerprint := fp.RandomizeFingerprint(base)
```

### 4. Stealth режим (максимальная защита)

```go
// Настройки chromedp для stealth
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("disable-blink-features", "AutomationControlled"),
    chromedp.Flag("exclude-switches", "enable-automation"),
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

### 5. Использование с прокси

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.ProxyServer("http://proxy:port"),
    chromedp.Flag("disable-blink-features", "AutomationControlled"),
)

allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
defer cancel()

// Настройте fingerprint под локацию прокси
fingerprint := fp.NewChrome119Windows11()
fingerprint.Timezone = &fp.Timezone{
    ID:     "Europe/London", // Если прокси в UK
    Offset: 0,
}
fingerprint.Language = "en-GB"
```

## Параметры fingerprint

### Основные

| Параметр    | Описание          | Пример                    |
| ----------- | ----------------- | ------------------------- |
| `UserAgent` | User-Agent строка | `"Mozilla/5.0..."`        |
| `Platform`  | Платформа         | `"Win32"`, `"MacIntel"`   |
| `Language`  | Основной язык     | `"en-US"`, `"ru-RU"`      |
| `Languages` | Список языков     | `[]string{"en-US", "en"}` |

### Продвинутые

```go
// WebGL
WebGL: &fp.WebGL{
    Vendor:   "Google Inc. (NVIDIA)",
    Renderer: "ANGLE (NVIDIA GeForce RTX 3080)",
}

// Canvas защита от fingerprinting
Canvas: &fp.Canvas{
    Noise: 0.02, // 0.0 - 1.0
}

// WebRTC
WebRTC: &fp.WebRTC{
    Disable: true, // Отключить полностью
}

// Экран
Screen: &fp.Screen{
    Width:            1920,
    Height:           1080,
    DevicePixelRatio: 1.0,
}
```

## Проверка fingerprint

Рекомендуемые сайты для проверки:

1. **https://browserleaks.com/** - Комплексная проверка
2. **https://bot.sannysoft.com/** - Проверка на бота
3. **https://www.whatismybrowser.com/** - Информация о браузере
4. **https://amiunique.org/** - Уникальность fingerprint

## Примеры

В директории `examples/` вы найдёте готовые примеры:

```bash
# Базовый пример
cd examples/basic && go run main.go

# Stealth режим
cd examples/stealth && go run main.go

# Случайный fingerprint
cd examples/random && go run main.go

# Множественные сессии
cd examples/multi-session && go run main.go
```

## Makefile команды

```bash
make deps          # Установить зависимости
make test          # Запустить тесты
make run-basic     # Запустить базовый пример
make run-stealth   # Запустить stealth режим
make run-random    # Случайный fingerprint
make clean         # Очистить временные файлы
```

## Troubleshooting

### Браузер не запускается

Убедитесь, что установлен Chrome или Chromium:

```bash
# Windows: скачайте Chrome с google.com/chrome
# Linux:
sudo apt install chromium-browser

# MacOS:
brew install --cask google-chrome
```

### WebDriver detected

Используйте полный stealth режим:

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("disable-blink-features", "AutomationControlled"),
    chromedp.Flag("exclude-switches", "enable-automation"),
    chromedp.Flag("disable-extensions", false),
)
```

### Fingerprint не применяется

Убедитесь, что `ApplyAll()` вызывается до навигации:

```go
chromedp.Run(ctx,
    injector.ApplyAll(ctx),        // Сначала применяем
    chromedp.Navigate("url"),      // Потом переходим
)
```

## Следующие шаги

- Прочитайте [README.md](README.md) для детальной документации
- Изучите примеры в директории `examples/`
- Посмотрите [CONTRIBUTING.md](CONTRIBUTING.md) если хотите внести вклад

## Поддержка

Если у вас возникли вопросы или проблемы:

1. Проверьте существующие Issues
2. Создайте новый Issue с детальным описанием
3. Приложите код для воспроизведения проблемы

---

**Готово!** Теперь вы можете начать использовать fingerprint-injector в своих проектах. 🚀

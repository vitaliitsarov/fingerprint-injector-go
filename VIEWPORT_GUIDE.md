# 📐 Viewport & Device Metrics - Руководство

Viewport - это видимая область браузера, которая отличается для разных устройств. Правильная настройка viewport критически важна для эмуляции мобильных устройств.

## 🎯 Что такое Viewport?

**Viewport** - это размер окна браузера, в котором отображается веб-страница.

### Desktop vs Mobile

| Параметр           | Desktop              | Mobile                |
| ------------------ | -------------------- | --------------------- |
| Viewport           | Большой (1920×1080+) | Маленький (360-414px) |
| Device Pixel Ratio | 1.0-2.0              | 2.0-3.0               |
| Touch              | Нет                  | Да                    |
| Orientation        | Горизонтальная       | Вертикальная          |

---

## 🔧 Как это работает в fingerprint-injector

### Автоматическая настройка

Библиотека **автоматически** настраивает viewport на основе выбранной платформы:

```go
fingerprint := fp.NewChrome119Android() // Мобильный viewport
// или
fingerprint := fp.NewChrome119Windows11() // Desktop viewport

injector := fp.NewInjector(fingerprint)

chromedp.Run(ctx,
    injector.ApplyAll(ctx), // Автоматически применяет viewport
    chromedp.Navigate("https://example.com"),
)
```

### Что устанавливается

При вызове `injector.ApplyAll(ctx)` библиотека устанавливает:

1. **Device Metrics** (через CDP):

   - Width и Height
   - Device Pixel Ratio
   - Mobile flag

2. **Touch Emulation** (для мобильных):

   - Touch events
   - Max touch points

3. **Screen properties** (через JavaScript):
   - screen.width
   - screen.height
   - screen.availWidth/Height
   - window.devicePixelRatio

---

## 📱 Viewport для разных платформ

### Windows Desktop

```go
fingerprint := fp.NewChrome119Windows11()
// Viewport: 1920×1080
// DPI: 1.0
// Touch: Нет
// Type: Desktop
```

**Характеристики:**

- Большой экран
- Стандартный DPI
- Мышь и клавиатура
- Горизонтальная ориентация

### MacOS Desktop (Retina)

```go
fingerprint := fp.NewChrome119MacOS()
// Viewport: 2560×1440
// DPI: 2.0
// Touch: Нет (но может быть трекпад)
// Type: Desktop
```

**Характеристики:**

- Высокое разрешение
- Retina дисплей (2x DPI)
- Четкие шрифты и изображения

### Android Mobile

```go
fingerprint := fp.NewChrome119Android()
// Viewport: 412×915
// DPI: 2.625
// Touch: Да
// Type: Mobile
```

**Характеристики:**

- Компактный экран
- Высокий DPI
- Touch события
- Вертикальная ориентация

### iOS Mobile

```go
fingerprint := fp.NewSafari17iOS()
// Viewport: 390×844
// DPI: 3.0
// Touch: Да
// Type: Mobile
```

**Характеристики:**

- Super Retina (3x DPI)
- Точные touch события
- iOS специфичные особенности

---

## 🎨 Примеры использования

### Пример 1: Проверка viewport

```go
package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Выбираем устройство
    fingerprint := fp.NewChrome119Android()
    injector := fp.NewInjector(fingerprint)

    var width, height int
    var dpr float64

    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://example.com"),

        // Проверяем viewport
        chromedp.Evaluate(`window.innerWidth`, &width),
        chromedp.Evaluate(`window.innerHeight`, &height),
        chromedp.Evaluate(`window.devicePixelRatio`, &dpr),
    )

    fmt.Printf("Viewport: %d×%d (DPI: %.1f)\n", width, height, dpr)
    fmt.Printf("Physical pixels: %.0f×%.0f\n",
        float64(width)*dpr, float64(height)*dpr)
}
```

### Пример 2: Адаптивный контент

```go
func testResponsive(url string) {
    devices := map[string]*fp.Fingerprint{
        "Desktop": fp.NewChrome119Windows11(),
        "Tablet":  fp.NewChrome119Android(), // Можно настроить
        "Mobile":  fp.NewSafari17iOS(),
    }

    for name, fingerprint := range devices {
        ctx, cancel := chromedp.NewContext(context.Background())
        injector := fp.NewInjector(fingerprint)

        var isMobile bool
        chromedp.Run(ctx,
            injector.ApplyAll(ctx),
            chromedp.Navigate(url),
            chromedp.Evaluate(`window.innerWidth < 768`, &isMobile),
        )

        fmt.Printf("%s: Mobile view = %v\n", name, isMobile)
        cancel()
    }
}
```

### Пример 3: Скриншоты разных размеров

```go
func captureScreenshots(url string) {
    devices := []struct {
        name string
        fp   *fp.Fingerprint
    }{
        {"desktop", fp.NewChrome119Windows11()},
        {"mobile", fp.NewChrome119Android()},
        {"tablet", fp.NewChrome119iOS()},
    }

    for _, device := range devices {
        ctx, cancel := chromedp.NewContext(context.Background())
        injector := fp.NewInjector(device.fp)

        var screenshot []byte
        chromedp.Run(ctx,
            injector.ApplyAll(ctx),
            chromedp.Navigate(url),
            chromedp.FullScreenshot(&screenshot, 100),
        )

        // Сохраняем screenshot
        os.WriteFile(device.name+".png", screenshot, 0644)
        cancel()
    }
}
```

### Пример 4: Кастомный viewport

```go
// Создаем кастомный viewport для планшета
fingerprint := fp.NewChrome119Android()

// Модифицируем под планшет (iPad размер)
fingerprint.Screen.Width = 768
fingerprint.Screen.Height = 1024
fingerprint.Screen.AvailWidth = 768
fingerprint.Screen.AvailHeight = 1024
fingerprint.Screen.DevicePixelRatio = 2.0

injector := fp.NewInjector(fingerprint)
// Viewport будет автоматически установлен как 768×1024
```

---

## 📊 Сравнение Viewport

Запустите демонстрацию чтобы увидеть разницу:

```bash
# Через Go
cd examples/viewport-demo
go run main.go

# Через Makefile
make run-viewport
```

**Пример вывода:**

```
[1/4] Тестируем: 🖥️  Desktop Windows (1920×1080)
  ┌─────────────────────────────────────────────────────────┐
  │ Platform: Win32                                         │
  │ Type: Desktop                                           │
  ├─────────────────────────────────────────────────────────┤
  │ Screen Resolution:   1920 × 1080 px                     │
  │ Viewport Size:       1920 × 1040 px                     │
  │ Device Pixel Ratio:  1.00x                              │
  │ Physical Pixels:     1920 × 1040 px                     │
  ├─────────────────────────────────────────────────────────┤
  │ Touch Support:       false                              │
  └─────────────────────────────────────────────────────────┘

[2/4] Тестируем: 📱 iPhone (390×844, 3x)
  ┌─────────────────────────────────────────────────────────┐
  │ Platform: iPhone                                        │
  │ Type: Mobile                                            │
  ├─────────────────────────────────────────────────────────┤
  │ Screen Resolution:    390 ×  844 px                     │
  │ Viewport Size:        390 ×  844 px                     │
  │ Device Pixel Ratio:  3.00x                              │
  │ Physical Pixels:     1170 × 2532 px                     │
  ├─────────────────────────────────────────────────────────┤
  │ Touch Support:       true                               │
  └─────────────────────────────────────────────────────────┘
```

---

## 🔍 Проверка Viewport

### Лучшие сайты для проверки:

1. **https://whatismyviewport.com/**

   - Показывает текущий viewport
   - Device Pixel Ratio
   - Touch support

2. **https://www.mydevice.io/**

   - Детальная информация об устройстве
   - Screen vs Viewport
   - CSS pixels vs Physical pixels

3. **https://responsively.app/**

   - Тестирование responsive дизайна
   - Сравнение разных viewport

4. **Browser DevTools**
   ```javascript
   console.log("Viewport:", window.innerWidth, "x", window.innerHeight);
   console.log("Screen:", screen.width, "x", screen.height);
   console.log("DPR:", window.devicePixelRatio);
   console.log("Touch:", "ontouchstart" in window);
   ```

---

## 💡 Best Practices

### 1. Соответствие устройства и viewport

```go
// ✅ Правильно - мобильный User-Agent + мобильный viewport
fingerprint := fp.NewChrome119Android()
// Viewport автоматически 412×915

// ❌ Неправильно - мобильный UA с desktop viewport
fingerprint := fp.NewChrome119Android()
fingerprint.Screen.Width = 1920  // Не делайте так!
fingerprint.Screen.Height = 1080
```

### 2. DPI должен соответствовать устройству

```go
// iPhone - всегда высокий DPI
fingerprint := fp.NewSafari17iOS()
// DPI автоматически 3.0

// Desktop - обычно 1.0 или 2.0
fingerprint := fp.NewChrome119Windows11()
// DPI автоматически 1.0
```

### 3. Touch для мобильных устройств

```go
// Для мобильных touch включается автоматически
fingerprint := fp.NewChrome119Android()
injector := fp.NewInjector(fingerprint)
// Touch emulation будет включен автоматически
```

### 4. Учитывайте orientation

```go
// Portrait (вертикальная) - для телефонов
fingerprint := fp.NewSafari17iOS()
// Width < Height (390 < 844)

// Landscape (горизонтальная) - для десктопов
fingerprint := fp.NewChrome119Windows11()
// Width > Height (1920 > 1080)
```

---

## 🎯 Продвинутые техники

### Динамическое изменение viewport

```go
// Базовый fingerprint
fingerprint := fp.NewChrome119Android()

// Изменяем под конкретное устройство
// Samsung Galaxy S21
fingerprint.Screen.Width = 360
fingerprint.Screen.Height = 800
fingerprint.Screen.DevicePixelRatio = 3.0

// OnePlus 9
// fingerprint.Screen.Width = 412
// fingerprint.Screen.Height = 915
// fingerprint.Screen.DevicePixelRatio = 2.625
```

### Viewport для Web Scraping

```go
// Для мобильной версии сайта
func scrapeMobile(url string) {
    fingerprint := fp.NewChrome119Android()
    // Автоматически получит мобильную версию
}

// Для desktop версии
func scrapeDesktop(url string) {
    fingerprint := fp.NewChrome119Windows11()
    // Автоматически получит desktop версию
}
```

### A/B Testing viewport

```go
func testViewportResponsive(url string) {
    viewports := []struct {
        name   string
        width  int
        height int
    }{
        {"Mobile S", 320, 568},
        {"Mobile M", 375, 667},
        {"Mobile L", 414, 896},
        {"Tablet", 768, 1024},
        {"Desktop", 1920, 1080},
    }

    for _, vp := range viewports {
        fingerprint := fp.NewChrome119Windows11()
        fingerprint.Screen.Width = vp.width
        fingerprint.Screen.Height = vp.height

        // Тестируем...
    }
}
```

---

## 📚 Дополнительные ресурсы

- [README.md](README.md) - Основная документация
- [PLATFORM_GUIDE.md](PLATFORM_GUIDE.md) - Выбор платформ
- [examples/viewport-demo/](examples/viewport-demo/) - Готовый пример
- [examples/platform-selector/](examples/platform-selector/) - Интерактивный выбор

---

## 🔧 Troubleshooting

### Проблема: Сайт показывает desktop версию на мобильном fingerprint

**Решение**: Убедитесь, что `ApplyAll()` вызван ДО навигации:

```go
// ✅ Правильно
chromedp.Run(ctx,
    injector.ApplyAll(ctx),  // Сначала viewport
    chromedp.Navigate(url),  // Потом навигация
)
```

### Проблема: Touch события не работают

**Решение**: Проверьте что используется мобильный preset:

```go
// ✅ Touch будет работать
fingerprint := fp.NewChrome119Android()

// ❌ Touch не будет работать
fingerprint := fp.NewChrome119Windows11()
```

### Проблема: Неправильный DPI

**Решение**: Используйте готовые presets, они уже настроены:

```go
// ✅ DPI правильный (3.0 для iPhone)
fingerprint := fp.NewSafari17iOS()

// ❌ Не меняйте DPI вручную без необходимости
fingerprint.Screen.DevicePixelRatio = 1.0 // Плохо!
```

---

**Версия**: 1.0.0  
**Дата**: 2024-10-11  
**Автор**: fingerprint-injector team

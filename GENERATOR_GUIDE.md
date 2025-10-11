## 🎲 Умная генерация Fingerprint - Руководство

Библиотека включает **умный генератор** с базой данных реальных устройств, который создает **логически связанные** и **уникальные** fingerprint'ы.

## 🎯 Что это дает?

### ❌ Проблема со случайной генерацией:

```go
// Плохо - нелогичная комбинация
fingerprint.Platform = "iPhone"
fingerprint.UserAgent = "...Android..." // Неправильно!
fingerprint.Screen.Width = 1920        // iPhone не бывает таким
fingerprint.HardwareConcurrency = 32   // У iPhone не 32 ядра
```

### ✅ Умная генерация:

```go
// Хорошо - все логически связано
fingerprint := fp.GenerateRandomFingerprint()
// Platform: iPhone
// UserAgent: ...iPhone OS 17_0... (правильный iOS UA)
// Screen: 390×844 (реальный размер iPhone)
// CPU: 6 cores (правильно для iPhone)
// GPU: Apple A16 (соответствует iPhone)
```

---

## 📊 База данных устройств

Библиотека содержит базу **реальных** устройств:

### Desktop

- **Windows**: различные конфигурации (4-32 ядра, 8-64GB RAM)
- **MacOS**: MacBook Pro с M1/M2/M3 (8-12 ядер, Retina дисплеи)
- **Linux**: серверные и desktop конфигурации

### Mobile

- **iPhone**: 12, 13, 14, 15 (с правильными экранами и GPU)
- **Android**: Samsung, Google Pixel, OnePlus, Xiaomi
- **Tablets**: iPad Pro, Samsung Galaxy Tab

### GPU

- **Desktop**: NVIDIA (RTX 4090-GTX 1660), AMD (RX 7900-5700), Intel (UHD 770-630)
- **Mobile**: Apple (A17-A14), Qualcomm (Adreno 740-640), ARM (Mali)

### Браузеры

- Chrome 119-122
- Firefox 120-121
- Safari 17.0-17.1

---

## 🚀 Использование

### 1. Полностью случайный (с логикой)

```go
package main

import (
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
    // Генерирует случайное, но ЛОГИЧНОЕ устройство
    fingerprint := fp.GenerateRandomFingerprint()
    // Может быть: Windows desktop, MacBook, iPhone, Android и т.д.
    // Все параметры будут соответствовать друг другу!

    injector := fp.NewInjector(fingerprint)
    // ... использование
}
```

### 2. По типу устройства

```go
// Только Desktop устройства
fingerprint := fp.GenerateFingerprintByType("desktop")
// Результат: Windows/Mac/Linux с desktop характеристиками

// Только Mobile устройства
fingerprint := fp.GenerateFingerprintByType("mobile")
// Результат: iPhone/Android с мобильными характеристиками

// Только Tablet устройства
fingerprint := fp.GenerateFingerprintByType("tablet")
// Результат: iPad/Galaxy Tab
```

### 3. По операционной системе

```go
// Только Windows
fingerprint := fp.GenerateFingerprintByOS("windows")
// Случайная Windows конфигурация с правильным GPU, RAM, etc.

// Только MacOS
fingerprint := fp.GenerateFingerprintByOS("macos")
// MacBook с Apple M1/M2/M3, Retina дисплей

// Только iOS
fingerprint := fp.GenerateFingerprintByOS("ios")
// Случайный iPhone с правильными характеристиками

// Только Android
fingerprint := fp.GenerateFingerprintByOS("android")
// Случайный Android телефон

// Только Linux
fingerprint := fp.GenerateFingerprintByOS("linux")
// Linux desktop с правильными характеристиками
```

### 4. По браузеру

```go
// Только Chrome
fingerprint := fp.GenerateFingerprintByBrowser("chrome")
// Chrome на случайной платформе

// Только Firefox
fingerprint := fp.GenerateFingerprintByBrowser("firefox")
// Firefox на desktop (не поддерживает mobile)

// Только Safari
fingerprint := fp.GenerateFingerprintByBrowser("safari")
// Safari только на Apple устройствах (Mac/iPhone/iPad)
```

### 5. Продвинутая генерация с опциями

```go
generator := fp.NewFingerprintGenerator()

// Мобильный Chrome
fingerprint, _ := generator.Generate(&fp.GenerateOptions{
    DeviceType: "mobile",
    Browser:    "chrome",
})

// Windows Firefox
fingerprint, _ := generator.Generate(&fp.GenerateOptions{
    OS:      "windows",
    Browser: "firefox",
})

// iOS Safari (единственная правильная комбинация для iOS браузера)
fingerprint, _ := generator.Generate(&fp.GenerateOptions{
    OS:      "ios",
    Browser: "safari",
})
```

---

## 🎨 Примеры

### Пример 1: Web Scraping с разных устройств

```go
package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func scrapeWithRandomDevice(url string) {
    // Каждый раз новое устройство
    fingerprint := fp.GenerateRandomFingerprint()

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    injector := fp.NewInjector(fingerprint)

    var content string
    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate(url),
        chromedp.OuterHTML("body", &content),
    )

    fmt.Printf("Scraped %d bytes from %s\n", len(content), url)
    fmt.Printf("Used device: %s\n", fingerprint.Platform)
}

func main() {
    urls := []string{
        "https://example.com/1",
        "https://example.com/2",
        "https://example.com/3",
    }

    for _, url := range urls {
        scrapeWithRandomDevice(url)
    }
}
```

### Пример 2: Тестирование на разных OS

```go
func testOnAllOS(url string) {
    oses := []string{"windows", "macos", "linux", "ios", "android"}

    for _, os := range oses {
        fingerprint := fp.GenerateFingerprintByOS(os)

        ctx, cancel := chromedp.NewContext(context.Background())
        injector := fp.NewInjector(fingerprint)

        var title string
        chromedp.Run(ctx,
            injector.ApplyAll(ctx),
            chromedp.Navigate(url),
            chromedp.Title(&title),
        )

        fmt.Printf("[%s] Title: %s\n", os, title)
        cancel()
    }
}
```

### Пример 3: Генерация уникальных профилей

```go
func generateUniqueProfiles(count int) []*fp.Fingerprint {
    profiles := make([]*fp.Fingerprint, count)

    for i := 0; i < count; i++ {
        profiles[i] = fp.GenerateRandomFingerprint()

        // Каждый профиль будет уникальным и логичным
        fmt.Printf("Profile %d: %s, CPU: %d, GPU: %s\n",
            i+1,
            profiles[i].Platform,
            profiles[i].HardwareConcurrency,
            profiles[i].WebGL.UnmaskedRenderer)
    }

    return profiles
}

func main() {
    // Генерируем 100 уникальных профилей
    profiles := generateUniqueProfiles(100)

    // Используем для параллельного scraping
    // Каждый профиль логически связан и уникален
}
```

### Пример 4: Ротация профилей

```go
type ProfileRotator struct {
    profiles []*fp.Fingerprint
    current  int
}

func NewProfileRotator(count int, deviceType string) *ProfileRotator {
    profiles := make([]*fp.Fingerprint, count)

    for i := 0; i < count; i++ {
        if deviceType == "" {
            profiles[i] = fp.GenerateRandomFingerprint()
        } else {
            profiles[i] = fp.GenerateFingerprintByType(deviceType)
        }
    }

    return &ProfileRotator{
        profiles: profiles,
        current:  0,
    }
}

func (r *ProfileRotator) Next() *fp.Fingerprint {
    profile := r.profiles[r.current]
    r.current = (r.current + 1) % len(r.profiles)
    return profile
}

func main() {
    // Создаем ротатор с 10 мобильными профилями
    rotator := NewProfileRotator(10, "mobile")

    urls := []string{"url1", "url2", "url3"} // ... много URL

    for _, url := range urls {
        fingerprint := rotator.Next()
        // Используем fingerprint для scraping
    }
}
```

---

## 🔍 Логические правила генератора

Генератор **автоматически** следит за логическими связями:

### 1. Platform ↔ User-Agent

```go
// ✅ Правильно
Platform: "iPhone"
UserAgent: "...iPhone OS 17_0..."

// ❌ Никогда не сгенерирует
Platform: "iPhone"
UserAgent: "...Android..." // Неправильно!
```

### 2. Platform ↔ GPU

```go
// ✅ Правильно
Platform: "MacIntel"
GPU: "Apple M1 Pro"

Platform: "Win32"
GPU: "NVIDIA GeForce RTX 3080"

Platform: "Linux armv8l" (Android)
GPU: "Qualcomm Adreno 730"

// ❌ Никогда не сгенерирует
Platform: "iPhone"
GPU: "NVIDIA RTX 4090" // iPhone не имеет NVIDIA
```

### 3. Device Type ↔ Screen Size

```go
// ✅ Правильно
Type: "mobile"
Screen: 390×844 (iPhone размер)

Type: "desktop"
Screen: 1920×1080 (desktop размер)

// ❌ Никогда не сгенерирует
Type: "mobile"
Screen: 3840×2160 // Нет мобильных с 4K
```

### 4. Browser ↔ Platform

```go
// ✅ Правильно
Browser: "Safari"
Platform: "MacIntel" или "iPhone"

Browser: "Firefox"
Platform: "Win32", "MacIntel", "Linux x86_64" (только desktop)

// ❌ Никогда не сгенерирует
Browser: "Safari"
Platform: "Win32" // Safari не на Windows

Browser: "Firefox"
Platform: "iPhone" // Firefox не поддерживает iOS (в контексте fingerprint)
```

### 5. RAM ↔ CPU Cores

```go
// ✅ Правильно (логичные комбинации)
CPU: 4 cores, RAM: 8GB
CPU: 16 cores, RAM: 32GB
CPU: 6 cores (iPhone), RAM: 6GB

// ❌ Избегает нелогичных комбинаций
CPU: 32 cores, RAM: 4GB // Не бывает
```

---

## 📊 Статистика базы данных

| Категория            | Количество          |
| -------------------- | ------------------- |
| Типов устройств      | 14                  |
| Desktop конфигураций | 3 (Win/Mac/Linux)   |
| Mobile устройств     | 8 (iPhone/Android)  |
| Tablet устройств     | 2 (iPad/Galaxy Tab) |
| GPU моделей          | 30+                 |
| Браузеров            | 8 версий            |
| OS версий            | 20+                 |

### Возможных комбинаций: **100,000+**

Генератор может создать более 100 тысяч уникальных и логичных комбинаций!

---

## 🎯 Когда использовать?

### Используйте умный генератор:

✅ Когда нужны уникальные fingerprint'ы  
✅ Для массового scraping с ротацией  
✅ Для A/B тестирования на разных устройствах  
✅ Когда важна реалистичность fingerprint  
✅ Для обхода anti-bot систем

### Используйте готовые presets:

✅ Когда нужно конкретное устройство  
✅ Для отладки и тестирования  
✅ Когда важна предсказуемость

---

## 🔧 Расширение базы данных

Вы можете добавить свои устройства в `devices_db.go`:

```go
// Добавьте новое устройство
{
    Name:          "Custom Device",
    Type:          "mobile",
    Platform:      "Linux armv8l",
    CPUCores:      []int{8},
    RAM:           []int{12},
    ScreenWidths:  []int{400},
    ScreenHeights: []int{900},
    DPRs:          []float64{3.0},
}

// Добавьте новый GPU
{
    Vendor:   "Custom Vendor",
    Renderer: "Custom GPU",
    Type:     "mobile",
}
```

---

## 🚀 Демонстрация

Запустите пример чтобы увидеть генератор в действии:

```bash
# Через Go
cd examples/smart-generator
go run main.go

# Через Makefile
make run-generator
```

**Пример вывода:**

```
═══ 1️⃣  Полностью случайный (любое устройство) ═══
  Type: Mobile | Platform: iPhone
  User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 17_0...)
  Screen: 393×852 (DPI: 3.0) | CPU: 6 | RAM: 6GB
  GPU: Apple A16 GPU

═══ 2️⃣  Только Desktop устройства ═══
  Type: Desktop | Platform: Win32
  User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)...
  Screen: 2560×1440 (DPI: 1.5) | CPU: 12 | RAM: 32GB
  GPU: NVIDIA GeForce RTX 3080

═══ 3️⃣  Только Mobile устройства ═══
  Type: Mobile | Platform: Linux armv8l
  User-Agent: Mozilla/5.0 (Linux; Android 13; Google Pixel 8)...
  Screen: 412×915 (DPI: 2.6) | CPU: 8 | RAM: 8GB
  GPU: Qualcomm Adreno 730
```

---

## 💡 Best Practices

### 1. Кэшируйте профили для повторного использования

```go
// ✅ Хорошо
profiles := make([]*fp.Fingerprint, 100)
for i := 0; i < 100; i++ {
    profiles[i] = fp.GenerateRandomFingerprint()
}
// Переиспользуйте profiles

// ❌ Плохо (генерация каждый раз)
for i := 0; i < 1000; i++ {
    fp := fp.GenerateRandomFingerprint() // Медленно
    // ...
}
```

### 2. Выбирайте правильный тип для задачи

```go
// Для мобильных сайтов
fingerprint := fp.GenerateFingerprintByType("mobile")

// Для desktop сайтов
fingerprint := fp.GenerateFingerprintByType("desktop")
```

### 3. Сохраняйте и загружайте профили

```go
import "encoding/json"

// Сохранение
fingerprint := fp.GenerateRandomFingerprint()
data, _ := json.Marshal(fingerprint)
os.WriteFile("profile.json", data, 0644)

// Загрузка
data, _ := os.ReadFile("profile.json")
var fingerprint fp.Fingerprint
json.Unmarshal(data, &fingerprint)
```

---

## 📚 Дополнительные ресурсы

- [README.md](README.md) - Основная документация
- [PLATFORM_GUIDE.md](PLATFORM_GUIDE.md) - Выбор платформ
- [examples/smart-generator/](examples/smart-generator/) - Готовый пример

---

**Версия**: 1.0.0  
**Дата**: 2024-10-11  
**База данных**: 14 типов устройств, 30+ GPU, 100,000+ комбинаций

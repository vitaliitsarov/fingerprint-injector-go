# Архитектура проекта

Документация архитектуры fingerprint-injector для Golang.

## Структура проекта

```
fingerprint-injector/
├── fingerprint.go          # Основные структуры данных
├── injector.go             # Логика инжектирования
├── presets.go              # Готовые пресеты
├── utils.go                # Утилиты и генераторы
├── fingerprint_test.go     # Тесты для fingerprint
├── injector_test.go        # Тесты для injector
├── go.mod                  # Go модуль
├── examples/               # Примеры использования
│   ├── basic/             # Базовый пример
│   ├── custom/            # Кастомный fingerprint
│   ├── stealth/           # Stealth режим
│   ├── with-proxy/        # С прокси
│   ├── multi-session/     # Множественные сессии
│   └── random/            # Случайный fingerprint
└── docs/                  # Документация
```

## Основные компоненты

### 1. Fingerprint (fingerprint.go)

Центральная структура данных, содержащая все параметры отпечатка браузера.

```go
type Fingerprint struct {
    UserAgent           string
    Platform            string
    Vendor              string
    Language            string
    Languages           []string
    Screen              *Screen
    Timezone            *Timezone
    WebGL               *WebGL
    Canvas              *Canvas
    WebRTC              *WebRTC
    Fonts               []string
    Plugins             []Plugin
    HardwareConcurrency int
    DeviceMemory        int
    Audio               *Audio
    Battery             *Battery
}
```

**Назначение**: Хранит все параметры, которые нужно изменить в браузере.

**Методы**:
- `ToJSON()` - Конвертация в JSON
- `NewDefaultFingerprint()` - Создание с дефолтными значениями

### 2. Injector (injector.go)

Отвечает за инжектирование JavaScript кода в браузер через chromedp.

```go
type Injector struct {
    fingerprint *Fingerprint
}
```

**Основные методы**:

- `NewInjector(fp *Fingerprint)` - Создание инжектора
- `GetInjectionScript()` - Генерация JS кода
- `Inject(ctx)` - Инжектирование скрипта
- `SetUserAgentOverride(ctx)` - Установка UA через CDP
- `SetTimezoneOverride(ctx)` - Установка timezone через CDP
- `ApplyAll(ctx)` - Применение всех настроек

**Процесс инжектирования**:

1. Генерация JavaScript кода на основе Fingerprint
2. Установка параметров через Chrome DevTools Protocol (CDP)
3. Инжектирование скрипта через `page.AddScriptToEvaluateOnNewDocument`
4. Выполнение скрипта на текущей странице

### 3. Presets (presets.go)

Готовые конфигурации fingerprint для разных платформ.

**Доступные пресеты**:
- `NewChrome119Windows11()` - Windows 11
- `NewChrome119MacOS()` - MacOS
- `NewChrome119Linux()` - Linux
- `NewChrome119Android()` - Android

### 4. Utils (utils.go)

Вспомогательные функции для работы с fingerprint.

**Основные функции**:
- `RandomUserAgent(platform)` - Генерация случайного UA
- `RandomizeFingerprint(base)` - Добавление вариаций
- `RandomResolution()` - Случайное разрешение экрана
- `RandomTimezone()` - Случайная временная зона
- `GenerateRandomFingerprint()` - Полностью случайный fingerprint

## Поток данных

```
┌─────────────────┐
│   User Code     │
│  (main.go)      │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Fingerprint    │ ◄─── Presets / Custom Config
│   (struct)      │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Injector     │
│  NewInjector()  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  ApplyAll(ctx)  │
└────────┬────────┘
         │
         ├──► SetUserAgentOverride() ──► CDP
         │
         ├──► SetTimezoneOverride() ───► CDP
         │
         └──► Inject() ──► JavaScript ──► Browser
```

## JavaScript инжектирование

### Методы инжектирования

1. **Chrome DevTools Protocol (CDP)**
   - Используется для User-Agent, Timezone
   - Более надежный, нельзя переопределить из JS
   - Применяется до загрузки страницы

2. **JavaScript injection**
   - Используется для navigator, screen, WebGL и т.д.
   - Гибкий, можно модифицировать любые объекты
   - Выполняется через `page.AddScriptToEvaluateOnNewDocument`

### Структура JS кода

```javascript
(function() {
    'use strict';
    
    // 1. Переопределение navigator объектов
    Object.defineProperty(navigator, 'userAgent', {
        get: function() { return 'custom UA'; }
    });
    
    // 2. Переопределение screen параметров
    Object.defineProperty(screen, 'width', {
        get: function() { return 1920; }
    });
    
    // 3. Модификация WebGL
    const getParameter = WebGLRenderingContext.prototype.getParameter;
    WebGLRenderingContext.prototype.getParameter = function(param) {
        if (param === 37445) return 'Custom Vendor';
        return getParameter.call(this, param);
    };
    
    // 4. Canvas fingerprinting защита
    // Добавление шума к canvas
    
    // 5. WebRTC контроль
    // Отключение или подмена IP
    
    // 6. Скрытие automation
    Object.defineProperty(navigator, 'webdriver', {
        get: function() { return undefined; }
    });
})();
```

## Интеграция с chromedp

### Создание контекста

```go
// Базовые опции
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
)

allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
ctx, cancel := chromedp.NewContext(allocCtx)
```

### Применение fingerprint

```go
injector := fp.NewInjector(fingerprint)

chromedp.Run(ctx,
    injector.ApplyAll(ctx),         // 1. Применяем fingerprint
    chromedp.Navigate("url"),       // 2. Переходим на страницу
    chromedp.WaitReady("body"),     // 3. Ждем загрузки
    // ... остальные действия
)
```

## Техники защиты от детекции

### 1. CDP методы
- User-Agent override
- Timezone override
- Geolocation override (будущая функция)

### 2. JavaScript переопределения
- Navigator properties (platform, vendor, languages)
- Screen properties
- Hardware properties (cores, memory)

### 3. Fingerprinting защита
- Canvas noise injection
- Audio context noise
- WebGL parameter randomization

### 4. Скрытие automation
- Удаление `navigator.webdriver`
- Удаление `window.chrome.runtime`
- Модификация permissions API

### 5. WebRTC защита
- Полное отключение
- IP spoofing (будущая функция)

## Расширение функционала

### Добавление нового параметра

1. Добавьте поле в структуру `Fingerprint`:

```go
type Fingerprint struct {
    // ... существующие поля
    NewParameter string `json:"newParameter"`
}
```

2. Добавьте обработку в `GetInjectionScript()`:

```go
script += fmt.Sprintf(`
    Object.defineProperty(navigator, 'newParameter', {
        get: function() { return '%s'; }
    });
`, fp.NewParameter)
```

3. Обновите пресеты в `presets.go`

4. Добавьте тесты в `fingerprint_test.go`

### Добавление нового preset

```go
func NewCustomPreset() *Fingerprint {
    return &Fingerprint{
        UserAgent: "...",
        Platform:  "...",
        // ... остальные параметры
    }
}
```

## Тестирование

### Юнит-тесты

```bash
go test -v ./...
```

### Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Интеграционное тестирование

Проверка на реальных сайтах:
- browserleaks.com
- bot.sannysoft.com
- amiunique.org

## Производительность

### Оптимизации

1. **Кэширование JS кода**
   - Генерация скрипта один раз при создании Injector
   - Переиспользование для множественных страниц

2. **Минимальное количество CDP вызовов**
   - Группировка операций в `ApplyAll()`

3. **Lazy initialization**
   - Создание объектов только при необходимости

### Метрики

- Время инжектирования: ~50-100ms
- Размер JS кода: ~3-5KB
- Память: ~1-2MB на Injector instance

## Безопасность

### Лучшие практики

1. **Не храните чувствительные данные в Fingerprint**
2. **Используйте случайные вариации**
3. **Регулярно обновляйте пресеты**
4. **Тестируйте на детекцию**

### Ограничения

- Не защищает от продвинутых систем детекции
- WebRTC leak (частично решено)
- Font fingerprinting (требует дополнительной работы)
- Canvas/Audio fingerprinting (использует шум, но не идеально)

## Roadmap

### Планируемые функции

- [ ] Geolocation spoofing
- [ ] Media devices spoofing
- [ ] Advanced WebRTC IP handling
- [ ] Более точная Canvas protection
- [ ] Font injection
- [ ] Plugin injection
- [ ] Speech synthesis spoofing
- [ ] Больше мобильных пресетов

## FAQ

**Q: Почему не headless режим?**  
A: Headless режим легко детектится. Рекомендуем использовать обычный режим.

**Q: Можно ли использовать с Selenium?**  
A: Нет, библиотека работает только с chromedp.

**Q: Насколько эффективна защита?**  
A: Зависит от системы детекции. Базовые системы обходятся легко, продвинутые - сложнее.

**Q: Влияет ли на производительность?**  
A: Минимально. Основное время занимает запуск браузера.

## Вклад

См. [CONTRIBUTING.md](CONTRIBUTING.md) для информации о том, как внести вклад в проект.

---

Последнее обновление: 2024-10-11


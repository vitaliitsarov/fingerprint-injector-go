# Fingerprint Injector - Обзор проекта

```
  _____ _                        ____       _       _     ___        _           _
 |  ___(_)_ __   __ _  ___ _ __ |  _ \ _ __(_)_ __ | |_  |_ _|_ __  (_) ___  ___| |_ ___  _ __
 | |_  | | '_ \ / _` |/ _ \ '__|| |_) | '__| | '_ \| __|  | || '_ \ | |/ _ \/ __| __/ _ \| '__|
 |  _| | | | | | (_| |  __/ |   |  __/| |  | | | | | |_   | || | | || |  __/ (__| || (_) | |
 |_|   |_|_| |_|\__, |\___|_|   |_|   |_|  |_|_| |_|\__| |___|_| |_|/ |\___|\___|\__\___/|_|
                |___/                                              |__/
```

## 🎯 Что это?

**fingerprint-injector** - это библиотека для Golang, позволяющая изменять и контролировать browser fingerprint при работе с chromedp. Аналог популярных Node.js решений для защиты от детекции автоматизации.

## 📦 Основные компоненты

### 1. Core Library (4 файла)

- `fingerprint.go` - Структуры данных (300+ строк)
- `injector.go` - Логика инжектирования (500+ строк)
- `presets.go` - Готовые конфигурации (300+ строк)
- `utils.go` - Утилиты и генераторы (200+ строк)

### 2. Tests (2 файла)

- `fingerprint_test.go` - Тесты структур данных
- `injector_test.go` - Тесты инжектирования

### 3. Examples (6 примеров)

- `basic/` - Простейшее использование
- `custom/` - Кастомные настройки
- `stealth/` - Максимальная защита
- `with-proxy/` - С прокси сервером
- `multi-session/` - Множественные сессии
- `random/` - Случайный fingerprint

### 4. Documentation (7 файлов)

- `README.md` - Главная документация
- `QUICKSTART.md` - Быстрый старт
- `ARCHITECTURE.md` - Архитектура
- `INTEGRATION_GUIDE.md` - Интеграция
- `CONTRIBUTING.md` - Для контрибьюторов
- `CHANGELOG.md` - История изменений
- `PROJECT_STRUCTURE.md` - Структура проекта

## 🚀 Возможности

### Базовые

✅ User-Agent подмена  
✅ Platform изменение  
✅ Language настройка  
✅ Screen параметры  
✅ Timezone управление  
✅ Hardware характеристики

### Продвинутые

✅ WebGL fingerprinting  
✅ Canvas fingerprinting защита  
✅ WebRTC управление  
✅ Audio context защита  
✅ Battery API подмена  
✅ Font enumeration

### Защита от детекции

✅ Скрытие navigator.webdriver  
✅ Удаление chrome.runtime  
✅ Permissions API модификация  
✅ Plugin list управление

## 📊 Статистика проекта

| Метрика                 | Значение                |
| ----------------------- | ----------------------- |
| Языков программирования | 2 (Go, JavaScript)      |
| Go файлов               | 6 основных + 2 тестовых |
| Примеров                | 6 полных примеров       |
| Документации            | 7 MD файлов             |
| Строк кода              | ~4700                   |
| Тестов                  | 15+ unit tests          |
| Presets                 | 4 готовых               |

## 🎨 Архитектура

```
┌─────────────────────────────────────────────────────────┐
│                     User Application                     │
└─────────────────┬───────────────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────────────┐
│              Fingerprint Structure                       │
│  (UserAgent, Platform, Screen, WebGL, Canvas, etc.)     │
└─────────────────┬───────────────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────────────┐
│                  Injector                                │
│  • GetInjectionScript()  - Generate JS code              │
│  • Inject()              - Inject to browser             │
│  • SetUserAgentOverride()- CDP override                  │
│  • ApplyAll()            - Apply everything              │
└─────────────────┬───────────────────────────────────────┘
                  │
                  ├──► CDP (Chrome DevTools Protocol)
                  │     • User-Agent
                  │     • Timezone
                  │
                  └──► JavaScript Injection
                        • Navigator properties
                        • Screen properties
                        • WebGL parameters
                        • Canvas/Audio protection
```

## 🔧 Технический стек

- **Go**: 1.21+
- **chromedp**: ^0.9.3
- **cdproto**: Latest
- **Chrome/Chromium**: Любая современная версия

## 📈 Workflow использования

```
1. Create Fingerprint
   ├── Use Preset (NewChrome119Windows11)
   ├── Create Custom (Manual struct)
   └── Generate Random (GenerateRandomFingerprint)

2. Create Injector
   └── injector := fp.NewInjector(fingerprint)

3. Setup chromedp
   ├── ExecAllocator (with stealth flags)
   └── NewContext

4. Apply Fingerprint
   └── injector.ApplyAll(ctx)

5. Navigate & Work
   ├── chromedp.Navigate(url)
   ├── chromedp.Click(selector)
   └── ... your automation
```

## 🎯 Основные use cases

### 1. Web Scraping

Обход anti-bot систем при сборе данных

### 2. Automated Testing

Тестирование приложений с разных устройств/ОС

### 3. Bot Development

Создание ботов с защитой от детекции

### 4. Research

Исследование browser fingerprinting техник

### 5. Privacy Tools

Инструменты для повышения приватности

## 🛡️ Уровни защиты

### Level 1: Basic (Preset)

```go
fingerprint := fp.NewChrome119Windows11()
```

Базовая подмена основных параметров

### Level 2: Custom

```go
fingerprint := &fp.Fingerprint{ /* custom */ }
```

Полный контроль над всеми параметрами

### Level 3: Random

```go
fingerprint := fp.GenerateRandomFingerprint()
```

Уникальный fingerprint для каждой сессии

### Level 4: Stealth

```go
+ Stealth chromedp flags
+ WebRTC disabled
+ Increased noise
+ Proxy rotation
```

Максимальная защита от детекции

## 📚 Документация

### Для начинающих

1. Начните с [QUICKSTART.md](QUICKSTART.md)
2. Изучите [examples/basic/](examples/basic/)
3. Прочитайте [README.md](README.md)

### Для продвинутых

1. [ARCHITECTURE.md](ARCHITECTURE.md) - Как все работает
2. [INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md) - Реальные сценарии
3. [examples/stealth/](examples/stealth/) - Продвинутые техники

### Для контрибьюторов

1. [CONTRIBUTING.md](CONTRIBUTING.md) - Как внести вклад
2. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - Структура кода
3. Tests - Посмотрите существующие тесты

## 🚦 Быстрый старт (30 секунд)

```bash
# 1. Установка
go get github.com/vitaliitsarov/fingerprint-injector-go

# 2. Создайте main.go
cat > main.go << 'EOF'
package main
import (
    "context"
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)
func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()
    injector := fp.NewInjector(fp.NewChrome119Windows11())
    chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://browserleaks.com"),
    )
    select {}
}
EOF

# 3. Запуск
go run main.go
```

## 🧪 Тестирование

### Запуск тестов

```bash
make test
```

### Coverage

```bash
make test
# Откройте coverage.html в браузере
```

### Проверка на сайтах

- https://browserleaks.com/
- https://bot.sannysoft.com/
- https://amiunique.org/

## 🎁 Что включено

### Code

- ✅ 4 основных модуля
- ✅ 2 тестовых файла
- ✅ 6 полных примеров
- ✅ Makefile с командами

### Documentation

- ✅ README с примерами
- ✅ Quick start guide
- ✅ Архитектурная документация
- ✅ Integration guide
- ✅ Contributing guide
- ✅ Changelog

### Development

- ✅ VS Code настройки
- ✅ Launch configurations
- ✅ .gitignore
- ✅ MIT License

## 🎓 Обучающие материалы

### Примеры от простого к сложному

1. **basic/** - Hello World (5 минут)
2. **custom/** - Свои настройки (10 минут)
3. **random/** - Генерация (15 минут)
4. **stealth/** - Защита (20 минут)
5. **with-proxy/** - С прокси (25 минут)
6. **multi-session/** - Масштабирование (30 минут)

## 🔗 Полезные ссылки

### Внутренние

- [README](README.md) - Главная страница
- [QUICKSTART](QUICKSTART.md) - Начало работы
- [ARCHITECTURE](ARCHITECTURE.md) - Как работает
- [INTEGRATION](INTEGRATION_GUIDE.md) - Интеграция

### Внешние

- [chromedp](https://github.com/chromedp/chromedp)
- [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
- [Browser Fingerprinting](https://browserleaks.com/)

## 📝 Чеклист для начала работы

- [ ] Установил Go 1.21+
- [ ] Установил Chrome/Chromium
- [ ] Склонировал/установил проект
- [ ] Запустил `make deps`
- [ ] Прочитал QUICKSTART.md
- [ ] Запустил базовый пример
- [ ] Протестировал на browserleaks.com
- [ ] Прочитал INTEGRATION_GUIDE.md
- [ ] Готов к интеграции в проект!

## 🎉 Поздравляем!

Теперь у вас есть полноценный инструмент для работы с browser fingerprinting в Golang!

---

**Версия**: 1.0.0  
**Лицензия**: MIT  
**Автор**: fingerprint-injector team  
**Дата**: 2024-10-11

Made with ❤️ for the Go community

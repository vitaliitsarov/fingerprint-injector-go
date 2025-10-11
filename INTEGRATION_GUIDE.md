# Руководство по интеграции

Это руководство поможет вам интегрировать fingerprint-injector в ваш проект.

## Установка

### 1. Добавьте в существующий проект

```bash
cd your-project
go get github.com/vitaliitsarov/fingerprint-injector-go
go get github.com/chromedp/chromedp
```

### 2. Импортируйте в код

```go
import (
    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)
```

## Сценарии использования

### Web Scraping

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type Scraper struct {
    ctx      context.Context
    cancel   context.CancelFunc
    injector *fp.Injector
}

func NewScraper() *Scraper {
    // Настройки для scraping
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
        chromedp.Flag("disable-blink-features", "AutomationControlled"),
        chromedp.Flag("user-data-dir", "./scraper-data"),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    ctx, _ := chromedp.NewContext(allocCtx)

    // Используем случайный fingerprint
    fingerprint := fp.GenerateRandomFingerprint()
    injector := fp.NewInjector(fingerprint)

    return &Scraper{
        ctx:      ctx,
        cancel:   cancel,
        injector: injector,
    }
}

func (s *Scraper) ScrapeURL(url string) (string, error) {
    var content string

    err := chromedp.Run(s.ctx,
        s.injector.ApplyAll(s.ctx),
        chromedp.Navigate(url),
        chromedp.WaitReady("body"),
        chromedp.OuterHTML("html", &content),
    )

    return content, err
}

func (s *Scraper) Close() {
    s.cancel()
}

func main() {
    scraper := NewScraper()
    defer scraper.Close()

    content, err := scraper.ScrapeURL("https://example.com")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Scraped %d bytes\n", len(content))
}
```

### Automated Testing

```go
package tests

import (
    "context"
    "testing"
    "time"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type TestBrowser struct {
    ctx      context.Context
    cancel   context.CancelFunc
    injector *fp.Injector
}

func setupBrowser(t *testing.T, preset string) *TestBrowser {
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
        chromedp.Flag("disable-blink-features", "AutomationControlled"),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    ctx, _ := chromedp.NewContext(allocCtx)

    var fingerprint *fp.Fingerprint
    switch preset {
    case "windows":
        fingerprint = fp.NewChrome119Windows11()
    case "mac":
        fingerprint = fp.NewChrome119MacOS()
    default:
        fingerprint = fp.NewDefaultFingerprint()
    }

    injector := fp.NewInjector(fingerprint)

    return &TestBrowser{
        ctx:      ctx,
        cancel:   cancel,
        injector: injector,
    }
}

func (tb *TestBrowser) Close() {
    tb.cancel()
}

func TestLoginFlow(t *testing.T) {
    browser := setupBrowser(t, "windows")
    defer browser.Close()

    var title string
    err := chromedp.Run(browser.ctx,
        browser.injector.ApplyAll(browser.ctx),
        chromedp.Navigate("https://example.com/login"),
        chromedp.WaitVisible(`input[name="username"]`),
        chromedp.SendKeys(`input[name="username"]`, "testuser"),
        chromedp.SendKeys(`input[name="password"]`, "testpass"),
        chromedp.Click(`button[type="submit"]`),
        chromedp.WaitVisible(`#dashboard`),
        chromedp.Title(&title),
    )

    if err != nil {
        t.Fatalf("Login failed: %v", err)
    }

    if title != "Dashboard" {
        t.Errorf("Expected Dashboard, got %s", title)
    }
}
```

### API Testing с разными User Agents

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type UserAgentTest struct {
    Platform string
    Expected string
}

func testUserAgent(test UserAgentTest, wg *sync.WaitGroup) {
    defer wg.Done()

    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    var fingerprint *fp.Fingerprint
    switch test.Platform {
    case "windows":
        fingerprint = fp.NewChrome119Windows11()
    case "mac":
        fingerprint = fp.NewChrome119MacOS()
    case "linux":
        fingerprint = fp.NewChrome119Linux()
    }

    injector := fp.NewInjector(fingerprint)

    var actualUA string
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://httpbin.org/user-agent"),
        chromedp.Text("body", &actualUA),
    )

    if err != nil {
        log.Printf("Error testing %s: %v", test.Platform, err)
        return
    }

    fmt.Printf("✓ %s UA validated\n", test.Platform)
}

func main() {
    tests := []UserAgentTest{
        {"windows", "Chrome 119 Windows"},
        {"mac", "Chrome 119 Mac"},
        {"linux", "Chrome 119 Linux"},
    }

    var wg sync.WaitGroup
    for _, test := range tests {
        wg.Add(1)
        go testUserAgent(test, &wg)
    }
    wg.Wait()

    fmt.Println("All tests completed!")
}
```

### Bot с ротацией fingerprint

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type Bot struct {
    currentFingerprint *fp.Fingerprint
    ctx                context.Context
    cancel             context.CancelFunc
}

func NewBot() *Bot {
    return &Bot{
        currentFingerprint: fp.GenerateRandomFingerprint(),
    }
}

func (b *Bot) initBrowser() {
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", false),
        chromedp.Flag("disable-blink-features", "AutomationControlled"),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    b.ctx, _ = chromedp.NewContext(allocCtx)
    b.cancel = cancel
}

func (b *Bot) RotateFingerprint() {
    // Закрываем старый браузер
    if b.cancel != nil {
        b.cancel()
    }

    // Генерируем новый fingerprint
    b.currentFingerprint = fp.GenerateRandomFingerprint()

    // Создаем новый браузер
    b.initBrowser()

    log.Println("🔄 Fingerprint rotated")
}

func (b *Bot) Visit(url string) error {
    injector := fp.NewInjector(b.currentFingerprint)

    return chromedp.Run(b.ctx,
        injector.ApplyAll(b.ctx),
        chromedp.Navigate(url),
        chromedp.Sleep(2*time.Second),
    )
}

func (b *Bot) Run(urls []string, rotateInterval int) {
    b.initBrowser()
    defer b.cancel()

    for i, url := range urls {
        // Ротация fingerprint каждые N запросов
        if i > 0 && i%rotateInterval == 0 {
            b.RotateFingerprint()
        }

        fmt.Printf("📍 Visiting: %s\n", url)
        if err := b.Visit(url); err != nil {
            log.Printf("Error visiting %s: %v", url, err)
        }

        time.Sleep(1 * time.Second)
    }
}

func main() {
    urls := []string{
        "https://example.com/page1",
        "https://example.com/page2",
        "https://example.com/page3",
        "https://example.com/page4",
        "https://example.com/page5",
    }

    bot := NewBot()
    bot.Run(urls, 2) // Ротация каждые 2 запроса
}
```

### Интеграция с прокси пулом

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/chromedp/chromedp"
    fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type ProxyPool struct {
    proxies []string
    current int
}

func NewProxyPool(proxies []string) *ProxyPool {
    return &ProxyPool{
        proxies: proxies,
        current: 0,
    }
}

func (pp *ProxyPool) Next() string {
    proxy := pp.proxies[pp.current]
    pp.current = (pp.current + 1) % len(pp.proxies)
    return proxy
}

type ProxyBrowser struct {
    proxyPool *ProxyPool
}

func NewProxyBrowser(proxies []string) *ProxyBrowser {
    return &ProxyBrowser{
        proxyPool: NewProxyPool(proxies),
    }
}

func (pb *ProxyBrowser) Request(url string) error {
    proxy := pb.proxyPool.Next()

    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.ProxyServer(proxy),
        chromedp.Flag("disable-blink-features", "AutomationControlled"),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    // Генерируем уникальный fingerprint для каждого запроса
    fingerprint := fp.GenerateRandomFingerprint()
    injector := fp.NewInjector(fingerprint)

    var ip string
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate("https://api.ipify.org"),
        chromedp.Text("body", &ip),
    )

    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }

    fmt.Printf("✓ Request via %s, IP: %s\n", proxy, ip)

    // Теперь делаем реальный запрос
    return chromedp.Run(ctx,
        chromedp.Navigate(url),
    )
}

func main() {
    proxies := []string{
        "http://proxy1:port",
        "http://proxy2:port",
        "http://proxy3:port",
    }

    browser := NewProxyBrowser(proxies)

    urls := []string{
        "https://example.com/1",
        "https://example.com/2",
        "https://example.com/3",
    }

    for _, url := range urls {
        if err := browser.Request(url); err != nil {
            log.Printf("Error: %v", err)
        }
    }
}
```

## Продвинутые техники

### Кастомная генерация fingerprint на основе геолокации

```go
func GenerateFingerprintForCountry(country string) *fp.Fingerprint {
    var base *fp.Fingerprint
    var timezone *fp.Timezone
    var language string
    var languages []string

    switch country {
    case "US":
        base = fp.NewChrome119Windows11()
        timezone = &fp.Timezone{ID: "America/New_York", Offset: -240}
        language = "en-US"
        languages = []string{"en-US", "en"}
    case "UK":
        base = fp.NewChrome119Windows11()
        timezone = &fp.Timezone{ID: "Europe/London", Offset: 0}
        language = "en-GB"
        languages = []string{"en-GB", "en"}
    case "RU":
        base = fp.NewChrome119Windows11()
        timezone = &fp.Timezone{ID: "Europe/Moscow", Offset: -180}
        language = "ru-RU"
        languages = []string{"ru-RU", "ru", "en"}
    default:
        return fp.NewDefaultFingerprint()
    }

    base.Timezone = timezone
    base.Language = language
    base.Languages = languages

    return base
}
```

### Сохранение и загрузка fingerprint

```go
import (
    "encoding/json"
    "os"
)

func SaveFingerprint(fp *fp.Fingerprint, filename string) error {
    data, err := json.MarshalIndent(fp, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func LoadFingerprint(filename string) (*fp.Fingerprint, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var fingerprint fp.Fingerprint
    if err := json.Unmarshal(data, &fingerprint); err != nil {
        return nil, err
    }

    return &fingerprint, nil
}

// Использование
func main() {
    // Сохранение
    fingerprint := fp.GenerateRandomFingerprint()
    SaveFingerprint(fingerprint, "fingerprint.json")

    // Загрузка
    loaded, _ := LoadFingerprint("fingerprint.json")
    injector := fp.NewInjector(loaded)
    // ...
}
```

## Best Practices

### 1. Переиспользование контекста

```go
// ❌ Плохо - создание нового контекста для каждого запроса
func badPattern(urls []string) {
    for _, url := range urls {
        ctx, cancel := chromedp.NewContext(context.Background())
        // ... работа
        cancel()
    }
}

// ✅ Хорошо - переиспользование контекста
func goodPattern(urls []string) {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    injector := fp.NewInjector(fp.GenerateRandomFingerprint())
    chromedp.Run(ctx, injector.ApplyAll(ctx))

    for _, url := range urls {
        chromedp.Run(ctx, chromedp.Navigate(url))
    }
}
```

### 2. Обработка ошибок

```go
func robustScraping(url string) error {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Таймаут
    ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    injector := fp.NewInjector(fp.GenerateRandomFingerprint())

    var content string
    err := chromedp.Run(ctx,
        injector.ApplyAll(ctx),
        chromedp.Navigate(url),
        chromedp.WaitReady("body", chromedp.ByQuery),
        chromedp.OuterHTML("body", &content),
    )

    if err != nil {
        if err == context.DeadlineExceeded {
            return fmt.Errorf("timeout: %w", err)
        }
        return fmt.Errorf("scraping failed: %w", err)
    }

    return nil
}
```

### 3. Логирование

```go
import "log"

func loggingScraper(url string) {
    log.Printf("🚀 Starting scraper for %s", url)

    fingerprint := fp.GenerateRandomFingerprint()
    log.Printf("📱 Generated fingerprint: Platform=%s, Cores=%d",
        fingerprint.Platform, fingerprint.HardwareConcurrency)

    injector := fp.NewInjector(fingerprint)

    // ... работа

    log.Printf("✅ Completed scraping %s", url)
}
```

## Troubleshooting

### Проблема: Fingerprint не применяется

**Решение**: Убедитесь, что `ApplyAll()` вызывается ДО навигации:

```go
// ✅ Правильно
chromedp.Run(ctx,
    injector.ApplyAll(ctx),
    chromedp.Navigate(url),
)

// ❌ Неправильно
chromedp.Run(ctx,
    chromedp.Navigate(url),
    injector.ApplyAll(ctx),
)
```

### Проблема: WebDriver все еще детектируется

**Решение**: Используйте полный набор stealth флагов:

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("disable-blink-features", "AutomationControlled"),
    chromedp.Flag("exclude-switches", "enable-automation"),
    chromedp.Flag("disable-extensions", false),
    chromedp.UserDataDir("./user-data"),
)
```

### Проблема: Память не освобождается

**Решение**: Всегда вызывайте `cancel()`:

```go
ctx, cancel := chromedp.NewContext(context.Background())
defer cancel() // Важно!
```

## Дополнительные ресурсы

- [chromedp Documentation](https://github.com/chromedp/chromedp)
- [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
- [Browser Fingerprinting Techniques](https://browserleaks.com/)

---

Успешной интеграции! 🚀

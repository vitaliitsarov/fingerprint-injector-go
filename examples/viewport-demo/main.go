package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

type DeviceInfo struct {
	Name        string
	Fingerprint *fp.Fingerprint
}

func main() {
	// Коллекция устройств для демонстрации
	devices := []DeviceInfo{
		{
			Name:        "🖥️  Desktop Windows (1920×1080)",
			Fingerprint: fp.NewChrome119Windows11(),
		},
		{
			Name:        "💻 MacBook Pro (2560×1440, Retina)",
			Fingerprint: fp.NewChrome119MacOS(),
		},
		{
			Name:        "📱 iPhone (390×844, 3x)",
			Fingerprint: fp.NewSafari17iOS(),
		},
		{
			Name:        "🤖 Android Pixel (412×915, 2.6x)",
			Fingerprint: fp.NewChrome119Android(),
		},
	}

	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║        Viewport & Device Metrics Demonstration            ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	for i, device := range devices {
		fmt.Printf("[%d/%d] Тестируем: %s\n", i+1, len(devices), device.Name)
		testDevice(device)

		if i < len(devices)-1 {
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Println("\n✅ Все тесты завершены!")
	fmt.Println("Сравните результаты viewport для разных устройств выше.")
}

func testDevice(device DeviceInfo) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // Headless для быстрого тестирования
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	injector := fp.NewInjector(device.Fingerprint)

	// Переменные для сбора информации
	var screenWidth, screenHeight int
	var innerWidth, innerHeight int
	var devicePixelRatio float64
	var isTouchEnabled bool
	var isMobile bool
	var platform string

	err := chromedp.Run(ctx,
		// Применяем fingerprint с viewport
		injector.ApplyAll(ctx),

		// Создаем простую HTML страницу для тестирования
		chromedp.Navigate("data:text/html,"+generateTestHTML()),

		// Собираем информацию
		chromedp.Evaluate(`screen.width`, &screenWidth),
		chromedp.Evaluate(`screen.height`, &screenHeight),
		chromedp.Evaluate(`window.innerWidth`, &innerWidth),
		chromedp.Evaluate(`window.innerHeight`, &innerHeight),
		chromedp.Evaluate(`window.devicePixelRatio`, &devicePixelRatio),
		chromedp.Evaluate(`'ontouchstart' in window`, &isTouchEnabled),
		chromedp.Evaluate(`/Mobi|Android/i.test(navigator.userAgent)`, &isMobile),
		chromedp.Evaluate(`navigator.platform`, &platform),
	)

	if err != nil {
		log.Printf("  ❌ Ошибка: %v\n", err)
		return
	}

	// Определяем тип устройства
	deviceType := "Desktop"
	if isMobile {
		deviceType = "Mobile"
	}

	// Рассчитываем физические пиксели
	physicalWidth := float64(innerWidth) * devicePixelRatio
	physicalHeight := float64(innerHeight) * devicePixelRatio

	// Выводим результаты в красивом формате
	fmt.Println("  ┌─────────────────────────────────────────────────────────┐")
	fmt.Printf("  │ Platform: %-45s │\n", platform)
	fmt.Printf("  │ Type: %-49s │\n", deviceType)
	fmt.Println("  ├─────────────────────────────────────────────────────────┤")
	fmt.Printf("  │ Screen Resolution:   %4d × %4d px                    │\n", screenWidth, screenHeight)
	fmt.Printf("  │ Viewport Size:       %4d × %4d px                    │\n", innerWidth, innerHeight)
	fmt.Printf("  │ Device Pixel Ratio:  %.2fx                             │\n", devicePixelRatio)
	fmt.Printf("  │ Physical Pixels:     %.0f × %.0f px                │\n", physicalWidth, physicalHeight)
	fmt.Println("  ├─────────────────────────────────────────────────────────┤")
	fmt.Printf("  │ Touch Support:       %-33v │\n", isTouchEnabled)
	fmt.Println("  └─────────────────────────────────────────────────────────┘")
	fmt.Println()
}

func generateTestHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Viewport Test</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .info {
            background: rgba(255,255,255,0.1);
            padding: 30px;
            border-radius: 10px;
            backdrop-filter: blur(10px);
        }
        h1 { margin: 0 0 20px 0; }
        .metric { 
            margin: 10px 0;
            font-size: 18px;
        }
    </style>
</head>
<body>
    <div class="info">
        <h1>📐 Viewport Info</h1>
        <div class="metric">Screen: <span id="screen"></span></div>
        <div class="metric">Viewport: <span id="viewport"></span></div>
        <div class="metric">DPR: <span id="dpr"></span></div>
        <div class="metric">Touch: <span id="touch"></span></div>
    </div>
    <script>
        document.getElementById('screen').textContent = screen.width + '×' + screen.height;
        document.getElementById('viewport').textContent = window.innerWidth + '×' + window.innerHeight;
        document.getElementById('dpr').textContent = window.devicePixelRatio;
        document.getElementById('touch').textContent = 'ontouchstart' in window ? 'Yes' : 'No';
    </script>
</body>
</html>`
}

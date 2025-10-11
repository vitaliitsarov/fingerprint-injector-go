package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Флаг для выбора платформы
	platform := flag.String("platform", "ios", "Выбор платформы: windows, linux, macos, android, ios, ios-chrome")
	flag.Parse()

	// Нормализуем имя платформы
	platformName := strings.ToLower(*platform)

	// Выбираем fingerprint на основе платформы
	var fingerprint *fp.Fingerprint

	switch platformName {
	case "windows", "win":
		log.Println("🪟 Выбрана платформа: Windows 11 (Chrome)")
		fingerprint = fp.NewChrome119Windows11()

	case "linux":
		log.Println("🐧 Выбрана платформа: Linux (Chrome)")
		fingerprint = fp.NewChrome119Linux()

	case "macos", "mac":
		log.Println("🍎 Выбрана платформа: MacOS (Chrome)")
		fingerprint = fp.NewChrome119MacOS()

	case "android":
		log.Println("🤖 Выбрана платформа: Android (Chrome)")
		fingerprint = fp.NewChrome119Android()

	case "ios", "iphone":
		log.Println("📱 Выбрана платформа: iOS (Safari)")
		fingerprint = fp.NewSafari17iOS()

	case "ios-chrome":
		log.Println("📱 Выбрана платформа: iOS (Chrome)")
		fingerprint = fp.NewChrome119iOS()

	default:
		log.Fatalf("❌ Неизвестная платформа: %s\n\nДоступные платформы:\n  • windows (win)\n  • linux\n  • macos (mac)\n  • android\n  • ios (iphone) - Safari\n  • ios-chrome - Chrome на iOS", platformName)
	}

	// Определяем тип устройства
	deviceType := "Desktop"
	if fingerprint.Platform == "Linux armv8l" || fingerprint.Platform == "iPhone" || fingerprint.Platform == "iPad" {
		deviceType = "Mobile"
	}

	// Выводим информацию о fingerprint
	fmt.Println("\n📋 Параметры fingerprint:")
	fmt.Printf("  Device Type: %s\n", deviceType)
	fmt.Printf("  Platform: %s\n", fingerprint.Platform)
	fmt.Printf("  User Agent: %s\n", fingerprint.UserAgent)
	fmt.Printf("  Viewport: %dx%d (DPI: %.1f)\n",
		fingerprint.Screen.Width,
		fingerprint.Screen.Height,
		fingerprint.Screen.DevicePixelRatio)
	fmt.Printf("  Timezone: %s\n", fingerprint.Timezone.ID)
	fmt.Printf("  CPU Cores: %d\n", fingerprint.HardwareConcurrency)
	fmt.Printf("  Memory: %d GB\n", fingerprint.DeviceMemory)
	fmt.Println()

	// Настройки chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	log.Println("🚀 Запуск браузера...")

	// Переменные для проверки
	var userAgent, platform2, language string
	var screenWidth, screenHeight, innerWidth, innerHeight int
	var devicePixelRatio float64
	var isTouchDevice bool

	err := chromedp.Run(ctx,
		// Применяем fingerprint
		injector.ApplyAll(ctx),

		// Переходим на страницу проверки
		chromedp.Navigate("https://whoer.net"),

		// Ждем загрузки
		chromedp.Sleep(3*time.Second),

		// Получаем данные для проверки
		chromedp.Evaluate(`navigator.userAgent`, &userAgent),
		chromedp.Evaluate(`navigator.platform`, &platform2),
		chromedp.Evaluate(`navigator.language`, &language),
		chromedp.Evaluate(`screen.width`, &screenWidth),
		chromedp.Evaluate(`screen.height`, &screenHeight),
		chromedp.Evaluate(`window.innerWidth`, &innerWidth),
		chromedp.Evaluate(`window.innerHeight`, &innerHeight),
		chromedp.Evaluate(`window.devicePixelRatio`, &devicePixelRatio),
		chromedp.Evaluate(`'ontouchstart' in window`, &isTouchDevice),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты проверки
	fmt.Println("\n✅ Результаты проверки:")
	fmt.Printf("  User Agent: %s\n", userAgent)
	fmt.Printf("  Platform: %s\n", platform2)
	fmt.Printf("  Language: %s\n", language)
	fmt.Printf("  Screen: %dx%d\n", screenWidth, screenHeight)
	fmt.Printf("  Viewport: %dx%d\n", innerWidth, innerHeight)
	fmt.Printf("  Device Pixel Ratio: %.1f\n", devicePixelRatio)
	fmt.Printf("  Touch Enabled: %v\n", isTouchDevice)
	fmt.Println()

	log.Println("========================================")
	log.Println("Браузер запущен. Проверьте fingerprint на странице.")
	log.Println("Нажмите Ctrl+C для выхода...")
	log.Println("========================================")

	// Держим браузер открытым
	select {}
}

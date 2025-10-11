package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║     Умная генерация Fingerprint с базой устройств        ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Демонстрация разных способов генерации
	examples := []struct {
		name string
		fn   func() *fp.Fingerprint
	}{
		{
			name: "1️⃣  Полностью случайный (любое устройство)",
			fn:   func() *fp.Fingerprint { return fp.GenerateRandomFingerprint() },
		},
		{
			name: "2️⃣  Только Desktop устройства",
			fn:   func() *fp.Fingerprint { return fp.GenerateFingerprintByType("desktop") },
		},
		{
			name: "3️⃣  Только Mobile устройства",
			fn:   func() *fp.Fingerprint { return fp.GenerateFingerprintByType("mobile") },
		},
		{
			name: "4️⃣  Только Windows",
			fn:   func() *fp.Fingerprint { return fp.GenerateFingerprintByOS("windows") },
		},
		{
			name: "5️⃣  Только iOS",
			fn:   func() *fp.Fingerprint { return fp.GenerateFingerprintByOS("ios") },
		},
		{
			name: "6️⃣  Только Firefox",
			fn:   func() *fp.Fingerprint { return fp.GenerateFingerprintByBrowser("firefox") },
		},
	}

	for _, example := range examples {
		fmt.Printf("═══ %s ═══\n", example.name)
		fingerprint := example.fn()
		printFingerprint(fingerprint)
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}

	// Интерактивная демонстрация - генерируем и тестируем
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          Запуск браузера для проверки                    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Генерируем случайный fingerprint
	fingerprint := fp.GenerateRandomFingerprint()

	fmt.Println("Сгенерирован fingerprint:")
	printFullFingerprint(fingerprint)

	// Настройки chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	injector := fp.NewInjector(fingerprint)

	log.Println("🚀 Запуск браузера с сгенерированным fingerprint...")

	var ua, platform string
	var screenW, screenH int

	err := chromedp.Run(ctx,
		injector.ApplyAll(ctx),
		chromedp.Navigate("https://whoer.net"),
		chromedp.Sleep(3*time.Second),

		chromedp.Evaluate(`navigator.userAgent`, &ua),
		chromedp.Evaluate(`navigator.platform`, &platform),
		chromedp.Evaluate(`screen.width`, &screenW),
		chromedp.Evaluate(`screen.height`, &screenH),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n✅ Проверка в браузере:")
	fmt.Printf("  User-Agent: %s\n", ua)
	fmt.Printf("  Platform: %s\n", platform)
	fmt.Printf("  Screen: %dx%d\n", screenW, screenH)

	fmt.Println("\n========================================")
	fmt.Println("Браузер запущен. Проверьте fingerprint на странице.")
	fmt.Println("Нажмите Ctrl+C для выхода...")
	fmt.Println("========================================")

	select {}
}

func printFingerprint(f *fp.Fingerprint) {
	deviceType := "Desktop"
	if f.Platform == "iPhone" || f.Platform == "iPad" || f.Platform == "Linux armv8l" {
		deviceType = "Mobile"
	}

	fmt.Printf("  Type: %s | Platform: %s\n", deviceType, f.Platform)
	fmt.Printf("  User-Agent: %s\n", truncate(f.UserAgent, 80))
	fmt.Printf("  Screen: %dx%d (DPI: %.1f) | CPU: %d | RAM: %dGB\n",
		f.Screen.Width, f.Screen.Height, f.Screen.DevicePixelRatio,
		f.HardwareConcurrency, f.DeviceMemory)
	fmt.Printf("  GPU: %s\n", truncate(f.WebGL.UnmaskedRenderer, 60))
}

func printFullFingerprint(f *fp.Fingerprint) {
	fmt.Println("┌────────────────────────────────────────────────────────┐")
	fmt.Printf("│ Platform: %-45s│\n", f.Platform)
	fmt.Printf("│ User-Agent: %-40s│\n", truncate(f.UserAgent, 40)+"...")
	fmt.Println("├────────────────────────────────────────────────────────┤")
	fmt.Printf("│ Screen: %4d×%4d px | DPI: %.2f                   │\n",
		f.Screen.Width, f.Screen.Height, f.Screen.DevicePixelRatio)
	fmt.Printf("│ CPU Cores: %2d | RAM: %2d GB                          │\n",
		f.HardwareConcurrency, f.DeviceMemory)
	fmt.Printf("│ Language: %-42s│\n", f.Language)
	fmt.Printf("│ Timezone: %-42s│\n", f.Timezone.ID)
	fmt.Println("├────────────────────────────────────────────────────────┤")
	fmt.Printf("│ GPU Vendor: %-40s│\n", truncate(f.WebGL.UnmaskedVendor, 40))
	fmt.Printf("│ GPU Renderer: %-38s│\n", truncate(f.WebGL.UnmaskedRenderer, 38))
	fmt.Println("└────────────────────────────────────────────────────────┘")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

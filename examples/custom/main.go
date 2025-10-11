package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Создаем контекст для chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Устанавливаем таймаут
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Создаем кастомный fingerprint
	fingerprint := &fp.Fingerprint{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		Platform:  "Win32",
		Vendor:    "Google Inc.",
		Language:  "ru-RU",
		Languages: []string{"ru-RU", "ru", "en-US", "en"},
		Screen: &fp.Screen{
			Width:            2560,
			Height:           1440,
			AvailWidth:       2560,
			AvailHeight:      1400,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 1.5,
		},
		Timezone: &fp.Timezone{
			ID:     "Europe/Moscow",
			Offset: -180,
		},
		WebGL: &fp.WebGL{
			Vendor:           "Google Inc. (NVIDIA)",
			Renderer:         "ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			UnmaskedVendor:   "NVIDIA Corporation",
			UnmaskedRenderer: "NVIDIA GeForce RTX 3080",
		},
		Canvas: &fp.Canvas{
			Noise: 0.02,
		},
		WebRTC: &fp.WebRTC{
			Disable: true, // Отключаем WebRTC для большей приватности
		},
		Fonts: []string{
			"Arial", "Courier New", "Georgia", "Times New Roman",
			"Verdana", "Segoe UI", "Tahoma", "Calibri",
		},
		Plugins:             []fp.Plugin{},
		HardwareConcurrency: 16,
		DeviceMemory:        32,
		Audio: &fp.Audio{
			Noise: 0.02,
		},
		Battery: &fp.Battery{
			Charging:        true,
			ChargingTime:    0,
			DischargingTime: 0,
			Level:           0.88,
		},
	}

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	// Применяем fingerprint и тестируем
	var userAgent, platform, language string
	var hardwareConcurrency int
	var webdriver interface{}

	err := chromedp.Run(ctx,
		// Применяем все настройки fingerprint
		injector.ApplyAll(ctx),

		// Переходим на about:blank для тестирования
		chromedp.Navigate("about:blank"),

		// Получаем данные для проверки
		chromedp.Evaluate(`navigator.userAgent`, &userAgent),
		chromedp.Evaluate(`navigator.platform`, &platform),
		chromedp.Evaluate(`navigator.language`, &language),
		chromedp.Evaluate(`navigator.hardwareConcurrency`, &hardwareConcurrency),
		chromedp.Evaluate(`navigator.webdriver`, &webdriver),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты
	log.Println("========================================")
	log.Println("Fingerprint проверка:")
	log.Println("========================================")
	log.Printf("User Agent: %s", userAgent)
	log.Printf("Platform: %s", platform)
	log.Printf("Language: %s", language)
	log.Printf("Hardware Concurrency: %d", hardwareConcurrency)
	log.Printf("WebDriver: %v", webdriver)
	log.Println("========================================")

	// Теперь идем на реальный сайт для проверки
	log.Println("Переходим на browserleaks.com для детальной проверки...")

	err = chromedp.Run(ctx,
		chromedp.Navigate("https://browserleaks.com/javascript"),
		chromedp.Sleep(10*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Проверьте fingerprint на странице.")
	log.Println("Нажмите Ctrl+C для выхода...")

	// Держим браузер открытым
	select {}
}

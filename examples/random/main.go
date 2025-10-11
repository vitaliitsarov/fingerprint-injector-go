package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Генерируем полностью случайный fingerprint
	fingerprint := fp.GenerateRandomFingerprint()

	log.Println("🎲 Сгенерирован случайный fingerprint:")
	log.Printf("  Platform: %s", fingerprint.Platform)
	log.Printf("  User Agent: %s", fingerprint.UserAgent)
	log.Printf("  Screen: %dx%d", fingerprint.Screen.Width, fingerprint.Screen.Height)
	log.Printf("  Timezone: %s", fingerprint.Timezone.ID)
	log.Printf("  CPU Cores: %d", fingerprint.HardwareConcurrency)
	log.Printf("  Memory: %d GB", fingerprint.DeviceMemory)
	log.Println()

	// Настройки chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	log.Println("🚀 Запуск браузера со случайным fingerprint...")

	// Проверяем на нескольких сайтах
	testSites := []string{
		"https://www.whatismybrowser.com/",
		"https://browserleaks.com/javascript",
	}

	for i, site := range testSites {
		log.Printf("📍 [%d/%d] Проверка: %s", i+1, len(testSites), site)

		err := chromedp.Run(ctx,
			injector.ApplyAll(ctx),
			chromedp.Navigate(site),
			chromedp.Sleep(8*time.Second),
		)

		if err != nil {
			log.Printf("❌ Ошибка: %v", err)
			continue
		}

		log.Printf("✓ Проверка завершена")

		if i < len(testSites)-1 {
			time.Sleep(2 * time.Second)
		}
	}

	log.Println("========================================")
	log.Println("✓ Все проверки завершены!")
	log.Println("Браузер остается открытым.")
	log.Println("Нажмите Ctrl+C для выхода...")
	log.Println("========================================")

	// Держим браузер открытым
	select {}
}

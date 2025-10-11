package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Максимально стелс настройки для chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),

		// Отключаем признаки автоматизации
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("disable-extensions", false),

		// Дополнительные флаги для стелса
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),

		// User data dir (для сохранения cookies и кэша)
		chromedp.UserDataDir("./chrome-data"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Используем MacOS fingerprint для разнообразия
	fingerprint := fp.NewChrome119MacOS()

	// Усиливаем стелс - отключаем WebRTC
	fingerprint.WebRTC.Disable = true

	// Увеличиваем шум для Canvas
	fingerprint.Canvas.Noise = 0.05

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	log.Println("🚀 Запуск браузера в stealth режиме...")

	// Список сайтов для тестирования защиты от детекции
	testSites := []string{
		"https://bot.sannysoft.com/",
		"https://arh.antoinevastel.com/bots/areyouheadless",
		"https://deviceandbrowserinfo.com/are_you_a_bot",
	}

	for _, site := range testSites {
		log.Printf("Проверяем: %s", site)

		err := chromedp.Run(ctx,
			// Применяем fingerprint
			injector.ApplyAll(ctx),

			// Переходим на сайт
			chromedp.Navigate(site),

			// Ждем загрузки
			chromedp.Sleep(8*time.Second),
		)

		if err != nil {
			log.Printf("Ошибка при проверке %s: %v", site, err)
			continue
		}

		log.Printf("✓ Проверка завершена: %s", site)
		time.Sleep(2 * time.Second)
	}

	log.Println("========================================")
	log.Println("Все проверки завершены!")
	log.Println("Браузер остается открытым для ручной проверки.")
	log.Println("Нажмите Ctrl+C для выхода...")
	log.Println("========================================")

	// Держим браузер открытым
	select {}
}

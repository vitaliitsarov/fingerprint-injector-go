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

	// Создаем fingerprint (используем preset для Windows 11)
	fingerprint := fp.NewChrome119Windows11()

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	// Применяем fingerprint и открываем сайт
	var result string
	err := chromedp.Run(ctx,
		// Применяем все настройки fingerprint
		injector.ApplyAll(ctx),

		// Переходим на тестовый сайт для проверки fingerprint
		chromedp.Navigate("https://whoer.net/ru#google_vignette"),

		// Ждем загрузки страницы
		chromedp.Sleep(5*time.Second),

		// Получаем user agent со страницы
		chromedp.Evaluate(`navigator.userAgent`, &result),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("User Agent: %s", result)
	log.Println("Браузер запущен. Проверьте fingerprint на странице.")
	log.Println("Нажмите Ctrl+C для выхода...")

	// Держим браузер открытым
	select {}
}

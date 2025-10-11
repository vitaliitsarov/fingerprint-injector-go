package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Настройки прокси (замените на свои)
	proxyServer := "http://proxy-server:port"

	// Опции chromedp с прокси
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),

		// Прокси настройки
		chromedp.ProxyServer(proxyServer),

		// Стелс настройки
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("disable-extensions", false),

		// Дополнительные флаги
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Таймаут
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Создаем fingerprint
	fingerprint := fp.NewChrome119Windows11()

	// Настраиваем timezone под прокси (например, если прокси в UK)
	fingerprint.Timezone = &fp.Timezone{
		ID:     "Europe/London",
		Offset: 0,
	}

	// Настраиваем язык под прокси
	fingerprint.Language = "en-GB"
	fingerprint.Languages = []string{"en-GB", "en"}

	// Отключаем WebRTC для защиты реального IP
	fingerprint.WebRTC.Disable = true

	// Создаем инжектор
	injector := fp.NewInjector(fingerprint)

	log.Println("🌐 Запуск браузера с прокси и fingerprint injection...")

	var ip string
	err := chromedp.Run(ctx,
		// Применяем fingerprint
		injector.ApplyAll(ctx),

		// Проверяем IP адрес
		chromedp.Navigate("https://api.ipify.org"),
		chromedp.Sleep(3*time.Second),
		chromedp.OuterHTML("body", &ip),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("📍 Текущий IP: %s", ip)

	// Проверяем fingerprint на специализированном сайте
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://browserleaks.com/ip"),
		chromedp.Sleep(10*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("========================================")
	log.Println("✓ Прокси и fingerprint применены")
	log.Println("Проверьте данные на странице")
	log.Println("Нажмите Ctrl+C для выхода...")
	log.Println("========================================")

	// Держим браузер открытым
	select {}
}

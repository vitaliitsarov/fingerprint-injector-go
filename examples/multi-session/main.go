package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	fp "github.com/vitaliitsarov/fingerprint-injector-go"
)

func main() {
	// Создаем несколько сессий с разными fingerprints
	sessions := []struct {
		name        string
		fingerprint *fp.Fingerprint
		url         string
	}{
		{
			name:        "Session-Windows",
			fingerprint: fp.NewChrome119Windows11(),
			url:         "https://www.whatismybrowser.com/",
		},
		{
			name:        "Session-MacOS",
			fingerprint: fp.NewChrome119MacOS(),
			url:         "https://www.whatismybrowser.com/",
		},
		{
			name:        "Session-Linux",
			fingerprint: fp.NewChrome119Linux(),
			url:         "https://www.whatismybrowser.com/",
		},
	}

	var wg sync.WaitGroup

	log.Println("🚀 Запуск множественных сессий с разными fingerprints...")

	for i, session := range sessions {
		wg.Add(1)

		go func(index int, sess struct {
			name        string
			fingerprint *fp.Fingerprint
			url         string
		}) {
			defer wg.Done()

			// Создаем отдельные директории для каждой сессии
			userDataDir := fmt.Sprintf("./chrome-data-session-%d", index)

			// Настройки chromedp для каждой сессии
			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false),
				chromedp.Flag("disable-blink-features", "AutomationControlled"),
				chromedp.Flag("exclude-switches", "enable-automation"),
				chromedp.UserDataDir(userDataDir),

				// Разные порты для отладки
				chromedp.Flag("remote-debugging-port", fmt.Sprintf("%d", 9222+index)),
			)

			allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()

			ctx, cancel := chromedp.NewContext(allocCtx)
			defer cancel()

			ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
			defer cancel()

			// Создаем инжектор для этой сессии
			injector := fp.NewInjector(sess.fingerprint)

			log.Printf("📱 Запуск %s...", sess.name)

			var userAgent, platform string
			err := chromedp.Run(ctx,
				// Применяем fingerprint
				injector.ApplyAll(ctx),

				// Переходим на сайт
				chromedp.Navigate(sess.url),
				chromedp.Sleep(5*time.Second),

				// Получаем данные для логирования
				chromedp.Evaluate(`navigator.userAgent`, &userAgent),
				chromedp.Evaluate(`navigator.platform`, &platform),
			)

			if err != nil {
				log.Printf("❌ Ошибка в %s: %v", sess.name, err)
				return
			}

			log.Printf("✓ %s запущен", sess.name)
			log.Printf("  User Agent: %s", userAgent)
			log.Printf("  Platform: %s", platform)

			// Держим сессию открытой
			time.Sleep(30 * time.Second)

		}(i, session)

		// Небольшая задержка между запусками
		time.Sleep(2 * time.Second)
	}

	log.Println("⏳ Ожидание завершения всех сессий...")
	wg.Wait()
	log.Println("✓ Все сессии завершены")
}

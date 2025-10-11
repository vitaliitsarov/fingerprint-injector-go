package fingerprint

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Injector отвечает за инжектирование fingerprint в браузер
type Injector struct {
	fingerprint *Fingerprint
}

// NewInjector создает новый инжектор с заданным fingerprint
func NewInjector(fingerprint *Fingerprint) *Injector {
	return &Injector{
		fingerprint: fingerprint,
	}
}

// GetInjectionScript возвращает JavaScript код для инжектирования fingerprint
func (inj *Injector) GetInjectionScript() string {
	fp := inj.fingerprint

	script := fmt.Sprintf(`
(function() {
	'use strict';

	// Переопределяем navigator.userAgent
	Object.defineProperty(navigator, 'userAgent', {
		get: function() { return '%s'; }
	});

	// Переопределяем navigator.platform
	Object.defineProperty(navigator, 'platform', {
		get: function() { return '%s'; }
	});

	// Переопределяем navigator.vendor
	Object.defineProperty(navigator, 'vendor', {
		get: function() { return '%s'; }
	});

	// Переопределяем navigator.language
	Object.defineProperty(navigator, 'language', {
		get: function() { return '%s'; }
	});

	// Переопределяем navigator.languages
	Object.defineProperty(navigator, 'languages', {
		get: function() { return %s; }
	});

	// Переопределяем navigator.hardwareConcurrency
	Object.defineProperty(navigator, 'hardwareConcurrency', {
		get: function() { return %d; }
	});

	// Переопределяем navigator.deviceMemory
	Object.defineProperty(navigator, 'deviceMemory', {
		get: function() { return %d; }
	});
`,
		fp.UserAgent,
		fp.Platform,
		fp.Vendor,
		fp.Language,
		toJSArray(fp.Languages),
		fp.HardwareConcurrency,
		fp.DeviceMemory,
	)

	// Экран
	if fp.Screen != nil {
		script += fmt.Sprintf(`
	// Переопределяем screen параметры
	Object.defineProperty(screen, 'width', {
		get: function() { return %d; }
	});
	Object.defineProperty(screen, 'height', {
		get: function() { return %d; }
	});
	Object.defineProperty(screen, 'availWidth', {
		get: function() { return %d; }
	});
	Object.defineProperty(screen, 'availHeight', {
		get: function() { return %d; }
	});
	Object.defineProperty(screen, 'colorDepth', {
		get: function() { return %d; }
	});
	Object.defineProperty(screen, 'pixelDepth', {
		get: function() { return %d; }
	});
	Object.defineProperty(window, 'devicePixelRatio', {
		get: function() { return %f; }
	});
`,
			fp.Screen.Width,
			fp.Screen.Height,
			fp.Screen.AvailWidth,
			fp.Screen.AvailHeight,
			fp.Screen.ColorDepth,
			fp.Screen.PixelDepth,
			fp.Screen.DevicePixelRatio,
		)
	}

	// WebGL
	if fp.WebGL != nil {
		script += fmt.Sprintf(`
	// Переопределяем WebGL параметры
	const getParameter = WebGLRenderingContext.prototype.getParameter;
	WebGLRenderingContext.prototype.getParameter = function(parameter) {
		if (parameter === 37445) {
			return '%s';
		}
		if (parameter === 37446) {
			return '%s';
		}
		return getParameter.call(this, parameter);
	};

	const getParameter2 = WebGL2RenderingContext.prototype.getParameter;
	WebGL2RenderingContext.prototype.getParameter = function(parameter) {
		if (parameter === 37445) {
			return '%s';
		}
		if (parameter === 37446) {
			return '%s';
		}
		return getParameter2.call(this, parameter);
	};
`,
			fp.WebGL.Vendor,
			fp.WebGL.Renderer,
			fp.WebGL.Vendor,
			fp.WebGL.Renderer,
		)
	}

	// Canvas fingerprinting защита
	if fp.Canvas != nil && fp.Canvas.Noise > 0 {
		script += fmt.Sprintf(`
	// Добавляем шум к Canvas для защиты от fingerprinting
	const originalToDataURL = HTMLCanvasElement.prototype.toDataURL;
	HTMLCanvasElement.prototype.toDataURL = function() {
		const context = this.getContext('2d');
		if (context) {
			const imageData = context.getImageData(0, 0, this.width, this.height);
			const noise = %f;
			for (let i = 0; i < imageData.data.length; i += 4) {
				imageData.data[i] = imageData.data[i] + Math.random() * noise;
				imageData.data[i + 1] = imageData.data[i + 1] + Math.random() * noise;
				imageData.data[i + 2] = imageData.data[i + 2] + Math.random() * noise;
			}
			context.putImageData(imageData, 0, 0);
		}
		return originalToDataURL.apply(this, arguments);
	};
`,
			fp.Canvas.Noise,
		)
	}

	// WebRTC
	if fp.WebRTC != nil && fp.WebRTC.Disable {
		script += `
	// Отключаем WebRTC
	navigator.getUserMedia = undefined;
	navigator.mediaDevices.getUserMedia = undefined;
	navigator.mediaDevices.enumerateDevices = function() { return Promise.resolve([]); };
	window.RTCPeerConnection = undefined;
	window.RTCSessionDescription = undefined;
	window.RTCIceCandidate = undefined;
`
	}

	// Battery API
	if fp.Battery != nil {
		script += fmt.Sprintf(`
	// Переопределяем Battery API
	const originalGetBattery = navigator.getBattery;
	navigator.getBattery = function() {
		return Promise.resolve({
			charging: %t,
			chargingTime: %f,
			dischargingTime: %f,
			level: %f,
			addEventListener: function() {},
			removeEventListener: function() {}
		});
	};
`,
			fp.Battery.Charging,
			fp.Battery.ChargingTime,
			fp.Battery.DischargingTime,
			fp.Battery.Level,
		)
	}

	// Timezone
	if fp.Timezone != nil {
		script += fmt.Sprintf(`
	// Переопределяем Timezone
	Date.prototype.getTimezoneOffset = function() {
		return %d;
	};
	Intl.DateTimeFormat.prototype.resolvedOptions = function() {
		return {
			locale: '%s',
			calendar: 'gregory',
			numberingSystem: 'latn',
			timeZone: '%s',
			year: 'numeric',
			month: 'numeric',
			day: 'numeric'
		};
	};
`,
			fp.Timezone.Offset,
			fp.Language,
			fp.Timezone.ID,
		)
	}

	// Скрываем следы автоматизации
	script += `
	// Скрываем webdriver
	Object.defineProperty(navigator, 'webdriver', {
		get: function() { return undefined; }
	});

	// Удаляем chrome.runtime
	if (window.chrome && window.chrome.runtime) {
		delete window.chrome.runtime;
	}

	// Переопределяем permissions
	const originalQuery = window.navigator.permissions.query;
	window.navigator.permissions.query = function(parameters) {
		if (parameters.name === 'notifications') {
			return Promise.resolve({ state: 'denied' });
		}
		return originalQuery.apply(this, arguments);
	};

	// Добавляем плагины
	Object.defineProperty(navigator, 'plugins', {
		get: function() {
			return [];
		}
	});

	console.log('🔒 Fingerprint injected successfully');
})();
`

	return script
}

// Inject инжектирует fingerprint в текущую страницу
func (inj *Injector) Inject(ctx context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		script := inj.GetInjectionScript()

		// Инжектируем скрипт на всех страницах
		_, err := page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
		if err != nil {
			return fmt.Errorf("failed to add script: %w", err)
		}

		// Также выполняем скрипт на текущей странице
		var res interface{}
		if err := chromedp.Evaluate(script, &res).Do(ctx); err != nil {
			// Игнорируем ошибку, так как страница может быть еще не загружена
		}

		return nil
	})
}

// SetUserAgentOverride устанавливает User-Agent через CDP
func (inj *Injector) SetUserAgentOverride(ctx context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		return emulation.SetUserAgentOverride(inj.fingerprint.UserAgent).
			WithAcceptLanguage(inj.fingerprint.Language).
			WithPlatform(inj.fingerprint.Platform).
			Do(ctx)
	})
}

// SetTimezoneOverride устанавливает временную зону через CDP
func (inj *Injector) SetTimezoneOverride(ctx context.Context) chromedp.Action {
	if inj.fingerprint.Timezone == nil {
		return chromedp.ActionFunc(func(ctx context.Context) error { return nil })
	}

	return chromedp.ActionFunc(func(ctx context.Context) error {
		return emulation.SetTimezoneOverride(inj.fingerprint.Timezone.ID).Do(ctx)
	})
}

// SetDeviceMetrics устанавливает viewport и метрики устройства через CDP
func (inj *Injector) SetDeviceMetrics(ctx context.Context) chromedp.Action {
	if inj.fingerprint.Screen == nil {
		return chromedp.ActionFunc(func(ctx context.Context) error { return nil })
	}

	return chromedp.ActionFunc(func(ctx context.Context) error {
		screen := inj.fingerprint.Screen

		// Определяем, мобильное ли это устройство
		isMobile := inj.isMobileDevice()

		// Устанавливаем метрики устройства
		return emulation.SetDeviceMetricsOverride(
			int64(screen.Width),
			int64(screen.Height),
			screen.DevicePixelRatio,
			isMobile,
		).WithScreenWidth(int64(screen.Width)).
			WithScreenHeight(int64(screen.Height)).
			Do(ctx)
	})
}

// SetTouchEmulation включает эмуляцию touch событий для мобильных устройств
func (inj *Injector) SetTouchEmulation(ctx context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		if inj.isMobileDevice() {
			// Включаем touch эмуляцию для мобильных
			return emulation.SetTouchEmulationEnabled(true).
				WithMaxTouchPoints(5).
				Do(ctx)
		}
		return nil
	})
}

// isMobileDevice определяет, является ли устройство мобильным
func (inj *Injector) isMobileDevice() bool {
	platform := inj.fingerprint.Platform
	// Проверяем по platform string
	return platform == "Linux armv8l" || // Android
		platform == "iPhone" || // iOS
		platform == "iPad" || // iPad
		inj.fingerprint.Screen.Width <= 768 // Или по размеру экрана
}

// ApplyAll применяет все настройки fingerprint
func (inj *Injector) ApplyAll(ctx context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// Применяем User-Agent
		if err := inj.SetUserAgentOverride(ctx).Do(ctx); err != nil {
			return fmt.Errorf("failed to set user agent: %w", err)
		}

		// Применяем Timezone
		if err := inj.SetTimezoneOverride(ctx).Do(ctx); err != nil {
			return fmt.Errorf("failed to set timezone: %w", err)
		}

		// Применяем Device Metrics (viewport и screen)
		if err := inj.SetDeviceMetrics(ctx).Do(ctx); err != nil {
			return fmt.Errorf("failed to set device metrics: %w", err)
		}

		// Применяем Touch Emulation для мобильных
		if err := inj.SetTouchEmulation(ctx).Do(ctx); err != nil {
			return fmt.Errorf("failed to set touch emulation: %w", err)
		}

		// Инжектируем скрипт
		if err := inj.Inject(ctx).Do(ctx); err != nil {
			return fmt.Errorf("failed to inject script: %w", err)
		}

		return nil
	})
}

// Вспомогательная функция для конвертации слайса в JS массив
func toJSArray(arr []string) string {
	result := "["
	for i, v := range arr {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("'%s'", v)
	}
	result += "]"
	return result
}

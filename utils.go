package fingerprint

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandomUserAgent генерирует случайный User-Agent на основе preset
func RandomUserAgent(platform string) string {
	chromeVersions := []string{"119.0.0.0", "118.0.0.0", "120.0.0.0"}
	version := chromeVersions[randomInt(len(chromeVersions))]

	switch platform {
	case "Win32":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", version)
	case "MacIntel":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", version)
	case "Linux x86_64":
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", version)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", version)
	}
}

// RandomizeFingerprint добавляет случайные вариации к fingerprint
func RandomizeFingerprint(base *Fingerprint) *Fingerprint {
	fp := *base // Копируем базовый fingerprint

	// Варьируем User-Agent
	fp.UserAgent = RandomUserAgent(base.Platform)

	// Варьируем Hardware Concurrency (4, 6, 8, 12, 16)
	cores := []int{4, 6, 8, 12, 16}
	fp.HardwareConcurrency = cores[randomInt(len(cores))]

	// Варьируем Device Memory (4, 8, 16, 32)
	memory := []int{4, 8, 16, 32}
	fp.DeviceMemory = memory[randomInt(len(memory))]

	// Варьируем Canvas noise
	if fp.Canvas != nil {
		fp.Canvas.Noise = 0.01 + float64(randomInt(50))/1000.0
	}

	// Варьируем Audio noise
	if fp.Audio != nil {
		fp.Audio.Noise = 0.01 + float64(randomInt(50))/1000.0
	}

	// Варьируем Battery level
	if fp.Battery != nil {
		fp.Battery.Level = 0.5 + float64(randomInt(50))/100.0
	}

	return &fp
}

// RandomResolution возвращает случайное разрешение экрана
func RandomResolution() *Screen {
	resolutions := []struct {
		width  int
		height int
	}{
		{1920, 1080},
		{1366, 768},
		{1536, 864},
		{1440, 900},
		{2560, 1440},
		{1600, 900},
	}

	res := resolutions[randomInt(len(resolutions))]

	return &Screen{
		Width:            res.width,
		Height:           res.height,
		AvailWidth:       res.width,
		AvailHeight:      res.height - 40, // Вычитаем высоту панели задач
		ColorDepth:       24,
		PixelDepth:       24,
		DevicePixelRatio: 1.0,
	}
}

// RandomTimezone возвращает случайную временную зону
func RandomTimezone() *Timezone {
	timezones := []struct {
		id     string
		offset int
	}{
		{"America/New_York", -240},
		{"America/Los_Angeles", -420},
		{"America/Chicago", -300},
		{"Europe/London", 0},
		{"Europe/Paris", -60},
		{"Europe/Moscow", -180},
		{"Asia/Tokyo", -540},
		{"Asia/Shanghai", -480},
		{"Australia/Sydney", -600},
	}

	tz := timezones[randomInt(len(timezones))]

	return &Timezone{
		ID:     tz.id,
		Offset: tz.offset,
	}
}

// randomInt возвращает случайное число от 0 до max (не включая max)
func randomInt(max int) int {
	if max <= 0 {
		return 0
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}
	return int(n.Int64())
}

// GenerateRandomFingerprint создает полностью случайный fingerprint
// Использует умный генератор с базой данных реальных устройств
func GenerateRandomFingerprint() *Fingerprint {
	generator := NewFingerprintGenerator()

	fingerprint, err := generator.Generate(&GenerateOptions{})
	if err != nil {
		// Fallback на старый метод
		platforms := []string{"Win32", "MacIntel", "Linux x86_64"}
		platform := platforms[randomInt(len(platforms))]

		var baseFingerprint *Fingerprint
		switch platform {
		case "Win32":
			baseFingerprint = NewChrome119Windows11()
		case "MacIntel":
			baseFingerprint = NewChrome119MacOS()
		default:
			baseFingerprint = NewChrome119Linux()
		}

		return RandomizeFingerprint(baseFingerprint)
	}

	return fingerprint
}

// GenerateFingerprintByType создает fingerprint для конкретного типа устройства
func GenerateFingerprintByType(deviceType string) *Fingerprint {
	generator := NewFingerprintGenerator()

	fingerprint, err := generator.Generate(&GenerateOptions{
		DeviceType: deviceType,
	})
	if err != nil {
		return NewDefaultFingerprint()
	}

	return fingerprint
}

// GenerateFingerprintByOS создает fingerprint для конкретной ОС
func GenerateFingerprintByOS(osName string) *Fingerprint {
	generator := NewFingerprintGenerator()

	fingerprint, err := generator.Generate(&GenerateOptions{
		OS: osName,
	})
	if err != nil {
		return NewDefaultFingerprint()
	}

	return fingerprint
}

// GenerateFingerprintByBrowser создает fingerprint для конкретного браузера
func GenerateFingerprintByBrowser(browserName string) *Fingerprint {
	generator := NewFingerprintGenerator()

	fingerprint, err := generator.Generate(&GenerateOptions{
		Browser: browserName,
	})
	if err != nil {
		return NewDefaultFingerprint()
	}

	return fingerprint
}

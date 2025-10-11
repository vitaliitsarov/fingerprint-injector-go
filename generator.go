package fingerprint

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// FingerprintGenerator генератор уникальных fingerprint'ов
type FingerprintGenerator struct {
	db *DeviceDatabase
}

// NewFingerprintGenerator создает новый генератор
func NewFingerprintGenerator() *FingerprintGenerator {
	return &FingerprintGenerator{
		db: GetDeviceDatabase(),
	}
}

// GenerateOptions опции для генерации fingerprint
type GenerateOptions struct {
	DeviceType string // "desktop", "mobile", "tablet", "" (любой)
	OS         string // "windows", "macos", "linux", "ios", "android", "" (любой)
	Browser    string // "chrome", "firefox", "safari", "" (любой)
}

// Generate генерирует логически связанный fingerprint
func (g *FingerprintGenerator) Generate(opts *GenerateOptions) (*Fingerprint, error) {
	if opts == nil {
		opts = &GenerateOptions{}
	}

	// Выбираем устройство
	device, err := g.selectDevice(opts)
	if err != nil {
		return nil, err
	}

	// Выбираем OS
	os := g.selectOS(device, opts)

	// Выбираем браузер
	browser := g.selectBrowser(device, opts)

	// Выбираем GPU
	gpu := g.selectGPU(device)

	// Генерируем User-Agent
	userAgent := g.generateUserAgent(browser, os, device)

	// Генерируем Screen
	screen := g.generateScreen(device)

	// Генерируем остальные параметры
	fingerprint := &Fingerprint{
		UserAgent: userAgent,
		Platform:  device.Platform,
		Vendor:    g.getVendor(browser.Name),
		Language:  g.generateLanguage(),
		Languages: g.generateLanguages(),
		Screen:    screen,
		Timezone:  RandomTimezone(),
		WebGL:     g.generateWebGL(gpu, device.Platform),
		Canvas: &Canvas{
			Noise: 0.01 + float64(randomInt(30))/1000.0,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts:               g.generateFonts(device.Platform),
		Plugins:             []Plugin{},
		HardwareConcurrency: device.CPUCores[randomInt(len(device.CPUCores))],
		DeviceMemory:        device.RAM[randomInt(len(device.RAM))],
		Audio: &Audio{
			Noise: 0.01 + float64(randomInt(30))/1000.0,
		},
		Battery: g.generateBattery(device.Type),
	}

	return fingerprint, nil
}

// selectDevice выбирает устройство на основе опций
func (g *FingerprintGenerator) selectDevice(opts *GenerateOptions) (*DeviceSpec, error) {
	var candidates []DeviceSpec

	for _, device := range g.db.Devices {
		// Фильтруем по типу устройства
		if opts.DeviceType != "" && device.Type != opts.DeviceType {
			continue
		}

		// Фильтруем по OS
		if opts.OS != "" {
			osMatch := false
			switch opts.OS {
			case "windows":
				osMatch = device.Platform == "Win32"
			case "macos":
				osMatch = device.Platform == "MacIntel"
			case "linux":
				osMatch = device.Platform == "Linux x86_64"
			case "ios":
				osMatch = device.Platform == "iPhone" || device.Platform == "iPad"
			case "android":
				osMatch = device.Platform == "Linux armv8l"
			}
			if !osMatch {
				continue
			}
		}

		candidates = append(candidates, device)
	}

	if len(candidates) == 0 {
		return nil, fmt.Errorf("no devices match the criteria")
	}

	// Выбираем случайное устройство
	return &candidates[randomInt(len(candidates))], nil
}

// selectOS выбирает OS для устройства
func (g *FingerprintGenerator) selectOS(device *DeviceSpec, opts *GenerateOptions) *OSVersion {
	var candidates []OSVersion

	for _, os := range g.db.OSes {
		if os.Platform == device.Platform {
			candidates = append(candidates, os)
		}
	}

	if len(candidates) == 0 {
		// Fallback
		return &OSVersion{
			Name:     "Unknown",
			Platform: device.Platform,
			Versions: []string{"1.0"},
		}
	}

	return &candidates[randomInt(len(candidates))]
}

// selectBrowser выбирает браузер
func (g *FingerprintGenerator) selectBrowser(device *DeviceSpec, opts *GenerateOptions) *BrowserVersion {
	var candidates []BrowserVersion

	for _, browser := range g.db.Browsers {
		// Фильтр по опциям
		if opts.Browser != "" && strings.ToLower(browser.Name) != strings.ToLower(opts.Browser) {
			continue
		}

		// Safari только на Apple устройствах
		if browser.Name == "Safari" && device.Platform != "iPhone" && device.Platform != "iPad" && device.Platform != "MacIntel" {
			continue
		}

		candidates = append(candidates, browser)
	}

	if len(candidates) == 0 {
		// Default Chrome
		return &BrowserVersion{Name: "Chrome", Version: "119.0.0.0", Major: 119}
	}

	return &candidates[randomInt(len(candidates))]
}

// selectGPU выбирает GPU для устройства
func (g *FingerprintGenerator) selectGPU(device *DeviceSpec) *GPUSpec {
	var candidates []GPUSpec

	gpuType := "desktop"
	if device.Type == "mobile" || device.Type == "tablet" {
		gpuType = "mobile"
	}

	// Специальная логика для платформ
	for _, gpu := range g.db.GPUs {
		if gpu.Type != gpuType {
			continue
		}

		// Apple устройства - только Apple GPU
		if (device.Platform == "MacIntel" || device.Platform == "iPhone" || device.Platform == "iPad") &&
			gpu.Vendor == "Apple Inc." {
			candidates = append(candidates, gpu)
			continue
		}

		// Android - Qualcomm или ARM
		if device.Platform == "Linux armv8l" &&
			(gpu.Vendor == "Qualcomm" || gpu.Vendor == "ARM") {
			candidates = append(candidates, gpu)
			continue
		}

		// Windows/Linux - NVIDIA, AMD, Intel
		if (device.Platform == "Win32" || device.Platform == "Linux x86_64") &&
			(gpu.Vendor == "NVIDIA Corporation" || gpu.Vendor == "AMD" || gpu.Vendor == "Intel Inc.") {
			candidates = append(candidates, gpu)
			continue
		}
	}

	if len(candidates) == 0 {
		// Fallback GPU
		return &GPUSpec{
			Vendor:   "Google Inc.",
			Renderer: "ANGLE (Generic)",
			Type:     gpuType,
		}
	}

	return &candidates[randomInt(len(candidates))]
}

// generateUserAgent генерирует User-Agent
func (g *FingerprintGenerator) generateUserAgent(browser *BrowserVersion, os *OSVersion, device *DeviceSpec) string {
	switch browser.Name {
	case "Chrome":
		return g.generateChromeUserAgent(browser, os, device)
	case "Firefox":
		return g.generateFirefoxUserAgent(browser, os, device)
	case "Safari":
		return g.generateSafariUserAgent(browser, os, device)
	default:
		return g.generateChromeUserAgent(browser, os, device)
	}
}

// generateChromeUserAgent генерирует Chrome User-Agent
func (g *FingerprintGenerator) generateChromeUserAgent(browser *BrowserVersion, os *OSVersion, device *DeviceSpec) string {
	switch device.Platform {
	case "Win32":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", browser.Version)
	case "MacIntel":
		osVersion := os.Versions[randomInt(len(os.Versions))]
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36",
			strings.Replace(osVersion, ".", "_", -1), browser.Version)
	case "Linux x86_64":
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", browser.Version)
	case "iPhone":
		osVersion := os.Versions[randomInt(len(os.Versions))]
		return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/%s Mobile/15E148 Safari/604.1",
			strings.Replace(osVersion, ".", "_", -1), browser.Major)
	case "Linux armv8l":
		androidVersion := os.Versions[randomInt(len(os.Versions))]
		return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Mobile Safari/537.36",
			androidVersion, device.Name, browser.Version)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", browser.Version)
	}
}

// generateFirefoxUserAgent генерирует Firefox User-Agent
func (g *FingerprintGenerator) generateFirefoxUserAgent(browser *BrowserVersion, os *OSVersion, device *DeviceSpec) string {
	switch device.Platform {
	case "Win32":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:%d.0) Gecko/20100101 Firefox/%s", browser.Major, browser.Version)
	case "MacIntel":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:%d.0) Gecko/20100101 Firefox/%s", browser.Major, browser.Version)
	case "Linux x86_64":
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64; rv:%d.0) Gecko/20100101 Firefox/%s", browser.Major, browser.Version)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:%d.0) Gecko/20100101 Firefox/%s", browser.Major, browser.Version)
	}
}

// generateSafariUserAgent генерирует Safari User-Agent
func (g *FingerprintGenerator) generateSafariUserAgent(browser *BrowserVersion, os *OSVersion, device *DeviceSpec) string {
	if device.Platform == "iPhone" {
		osVersion := os.Versions[randomInt(len(os.Versions))]
		return fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Mobile/15E148 Safari/604.1",
			strings.Replace(osVersion, ".", "_", -1), browser.Version)
	}

	// MacOS
	return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Safari/605.1.15", browser.Version)
}

// generateScreen генерирует параметры экрана
func (g *FingerprintGenerator) generateScreen(device *DeviceSpec) *Screen {
	width := device.ScreenWidths[randomInt(len(device.ScreenWidths))]
	height := device.ScreenHeights[randomInt(len(device.ScreenHeights))]
	dpr := device.DPRs[randomInt(len(device.DPRs))]

	availHeight := height
	if device.Type == "desktop" {
		availHeight = height - 40 // Панель задач
	}

	return &Screen{
		Width:            width,
		Height:           height,
		AvailWidth:       width,
		AvailHeight:      availHeight,
		ColorDepth:       24,
		PixelDepth:       24,
		DevicePixelRatio: dpr,
	}
}

// generateWebGL генерирует WebGL параметры
func (g *FingerprintGenerator) generateWebGL(gpu *GPUSpec, platform string) *WebGL {
	vendor := fmt.Sprintf("Google Inc. (%s)", gpu.Vendor)
	renderer := fmt.Sprintf("ANGLE (%s)", gpu.Renderer)

	if strings.Contains(platform, "Mac") || strings.Contains(platform, "iPhone") || strings.Contains(platform, "iPad") {
		vendor = gpu.Vendor
		renderer = gpu.Renderer
	}

	return &WebGL{
		Vendor:                 vendor,
		Renderer:               renderer,
		UnmaskedVendor:         gpu.Vendor,
		UnmaskedRenderer:       gpu.Renderer,
		ShadingLanguageVersion: "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
	}
}

// generateFonts генерирует список шрифтов
func (g *FingerprintGenerator) generateFonts(platform string) []string {
	commonFonts := []string{"Arial", "Courier New", "Georgia", "Times New Roman", "Verdana"}

	switch platform {
	case "Win32":
		return append(commonFonts, "Segoe UI", "Tahoma", "Trebuchet MS", "Calibri", "Consolas")
	case "MacIntel", "iPhone", "iPad":
		return append(commonFonts, "SF Pro Text", "SF Pro Display", "Helvetica Neue", "Menlo", "Monaco")
	case "Linux x86_64":
		return append(commonFonts, "Ubuntu", "Liberation Sans", "DejaVu Sans", "Noto Sans")
	case "Linux armv8l":
		return []string{"Roboto", "Noto Sans", "Droid Sans"}
	default:
		return commonFonts
	}
}

// generateLanguage генерирует основной язык
func (g *FingerprintGenerator) generateLanguage() string {
	languages := []string{"en-US", "en-GB", "de-DE", "fr-FR", "es-ES", "ru-RU", "zh-CN", "ja-JP"}
	return languages[randomInt(len(languages))]
}

// generateLanguages генерирует список языков
func (g *FingerprintGenerator) generateLanguages() []string {
	primary := g.generateLanguage()
	base := strings.Split(primary, "-")[0]

	result := []string{primary}
	if primary != base {
		result = append(result, base)
	}
	if base != "en" {
		result = append(result, "en-US", "en")
	}

	return result
}

// generateBattery генерирует параметры батареи
func (g *FingerprintGenerator) generateBattery(deviceType string) *Battery {
	if deviceType == "desktop" {
		return &Battery{
			Charging:        true,
			ChargingTime:    0,
			DischargingTime: 0,
			Level:           1.0,
		}
	}

	// Mobile/Tablet
	level := 0.5 + float64(randomInt(50))/100.0
	charging := randomInt(2) == 0

	return &Battery{
		Charging:        charging,
		ChargingTime:    0,
		DischargingTime: float64(10000 + randomInt(10000)),
		Level:           level,
	}
}

// getVendor возвращает vendor для браузера
func (g *FingerprintGenerator) getVendor(browserName string) string {
	switch browserName {
	case "Chrome":
		return "Google Inc."
	case "Firefox":
		return ""
	case "Safari":
		return "Apple Computer, Inc."
	default:
		return "Google Inc."
	}
}

// secureRandomInt безопасный генератор случайных чисел
func secureRandomInt(max int) int {
	if max <= 0 {
		return 0
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}
	return int(n.Int64())
}

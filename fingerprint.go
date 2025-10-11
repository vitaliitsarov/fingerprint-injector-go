package fingerprint

import (
	"encoding/json"
)

// Fingerprint содержит все параметры для изменения отпечатка браузера
type Fingerprint struct {
	UserAgent           string    `json:"userAgent"`
	Platform            string    `json:"platform"`
	Vendor              string    `json:"vendor"`
	Language            string    `json:"language"`
	Languages           []string  `json:"languages"`
	Screen              *Screen   `json:"screen"`
	Timezone            *Timezone `json:"timezone"`
	WebGL               *WebGL    `json:"webgl"`
	Canvas              *Canvas   `json:"canvas"`
	WebRTC              *WebRTC   `json:"webrtc"`
	Fonts               []string  `json:"fonts"`
	Plugins             []Plugin  `json:"plugins"`
	HardwareConcurrency int       `json:"hardwareConcurrency"`
	DeviceMemory        int       `json:"deviceMemory"`
	Audio               *Audio    `json:"audio"`
	Battery             *Battery  `json:"battery"`
}

// Screen параметры экрана
type Screen struct {
	Width            int     `json:"width"`
	Height           int     `json:"height"`
	AvailWidth       int     `json:"availWidth"`
	AvailHeight      int     `json:"availHeight"`
	ColorDepth       int     `json:"colorDepth"`
	PixelDepth       int     `json:"pixelDepth"`
	DevicePixelRatio float64 `json:"devicePixelRatio"`
}

// Timezone параметры временной зоны
type Timezone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"` // в минутах
}

// WebGL параметры WebGL
type WebGL struct {
	Vendor                 string `json:"vendor"`
	Renderer               string `json:"renderer"`
	UnmaskedVendor         string `json:"unmaskedVendor"`
	UnmaskedRenderer       string `json:"unmaskedRenderer"`
	ShadingLanguageVersion string `json:"shadingLanguageVersion"`
}

// Canvas параметры Canvas
type Canvas struct {
	Noise float64 `json:"noise"` // Уровень шума для canvas fingerprinting (0.0 - 1.0)
}

// WebRTC параметры WebRTC
type WebRTC struct {
	PublicIP string `json:"publicIp"`
	LocalIP  string `json:"localIp"`
	Disable  bool   `json:"disable"`
}

// Plugin информация о плагине
type Plugin struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Filename    string     `json:"filename"`
	MimeTypes   []MimeType `json:"mimeTypes"`
}

// MimeType MIME тип
type MimeType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Suffixes    string `json:"suffixes"`
}

// Audio параметры Audio Context
type Audio struct {
	Noise float64 `json:"noise"` // Уровень шума для audio fingerprinting (0.0 - 1.0)
}

// Battery параметры батареи
type Battery struct {
	Charging        bool    `json:"charging"`
	ChargingTime    float64 `json:"chargingTime"`
	DischargingTime float64 `json:"dischargingTime"`
	Level           float64 `json:"level"`
}

// ToJSON конвертирует Fingerprint в JSON строку
func (f *Fingerprint) ToJSON() (string, error) {
	data, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewDefaultFingerprint создает fingerprint с дефолтными значениями
func NewDefaultFingerprint() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		Platform:  "Win32",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            1920,
			Height:           1080,
			AvailWidth:       1920,
			AvailHeight:      1040,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 1.0,
		},
		Timezone: &Timezone{
			ID:     "America/New_York",
			Offset: -240,
		},
		WebGL: &WebGL{
			Vendor:                 "Google Inc. (Intel)",
			Renderer:               "ANGLE (Intel, Intel(R) UHD Graphics 630 Direct3D11 vs_5_0 ps_5_0)",
			UnmaskedVendor:         "Intel Inc.",
			UnmaskedRenderer:       "Intel(R) UHD Graphics 630",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts: []string{
			"Arial", "Courier New", "Georgia", "Times New Roman",
			"Verdana", "Trebuchet MS", "Comic Sans MS",
		},
		Plugins:             []Plugin{},
		HardwareConcurrency: 8,
		DeviceMemory:        8,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        true,
			ChargingTime:    0,
			DischargingTime: 0,
			Level:           1.0,
		},
	}
}

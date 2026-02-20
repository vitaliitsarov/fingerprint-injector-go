package fingerprint

// Chrome Windows 11 presets
func NewChrome119Windows11() *Fingerprint {
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
			Renderer:               "ANGLE (Intel, Intel(R) UHD Graphics 630 Direct3D11 vs_5_0 ps_5_0, D3D11)",
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
			"Verdana", "Trebuchet MS", "Comic Sans MS", "Segoe UI",
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

// Chrome MacOS preset
func NewChrome119MacOS() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		Platform:  "MacIntel",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            2560,
			Height:           1440,
			AvailWidth:       2560,
			AvailHeight:      1417,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 2.0,
		},
		Timezone: &Timezone{
			ID:     "America/Los_Angeles",
			Offset: -420,
		},
		WebGL: &WebGL{
			Vendor:                 "Google Inc. (Apple)",
			Renderer:               "ANGLE (Apple, Apple M1 Pro, OpenGL 4.1 Metal - 83.1)",
			UnmaskedVendor:         "Apple Inc.",
			UnmaskedRenderer:       "Apple M1 Pro",
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
			"Verdana", "Trebuchet MS", "Helvetica Neue", "Menlo",
		},
		Plugins:             []Plugin{},
		HardwareConcurrency: 10,
		DeviceMemory:        8,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        true,
			ChargingTime:    0,
			DischargingTime: 0,
			Level:           0.95,
		},
	}
}

// Chrome Linux preset
func NewChrome119Linux() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		Platform:  "Linux x86_64",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            1920,
			Height:           1080,
			AvailWidth:       1920,
			AvailHeight:      1053,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 1.0,
		},
		Timezone: &Timezone{
			ID:     "Europe/London",
			Offset: 0,
		},
		WebGL: &WebGL{
			Vendor:                 "Google Inc. (NVIDIA)",
			Renderer:               "ANGLE (NVIDIA, NVIDIA GeForce GTX 1080 Ti (0x00001B06) Direct3D11 vs_5_0 ps_5_0, D3D11)",
			UnmaskedVendor:         "NVIDIA Corporation",
			UnmaskedRenderer:       "NVIDIA GeForce GTX 1080 Ti/PCIe/SSE2",
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
			"Verdana", "Trebuchet MS", "Liberation Sans", "Ubuntu",
		},
		Plugins:             []Plugin{},
		HardwareConcurrency: 12,
		DeviceMemory:        16,
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

// Mobile Android preset
func NewChrome119Android() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
		Platform:  "Linux armv8l",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            412,
			Height:           915,
			AvailWidth:       412,
			AvailHeight:      915,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 2.625,
		},
		Timezone: &Timezone{
			ID:     "America/New_York",
			Offset: -240,
		},
		WebGL: &WebGL{
			Vendor:                 "Google Inc. (Qualcomm)",
			Renderer:               "ANGLE (Qualcomm, Adreno (TM) 730, OpenGL ES 3.2)",
			UnmaskedVendor:         "Qualcomm",
			UnmaskedRenderer:       "Adreno (TM) 730",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts:               []string{"Roboto", "Noto Sans"},
		Plugins:             []Plugin{},
		HardwareConcurrency: 8,
		DeviceMemory:        8,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        false,
			ChargingTime:    0,
			DischargingTime: 18000,
			Level:           0.75,
		},
	}
}

// Mobile iOS preset (Safari on iPhone)
func NewSafari17iOS() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
		Platform:  "iPhone",
		Vendor:    "Apple Computer, Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            390,
			Height:           844,
			AvailWidth:       390,
			AvailHeight:      844,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 3.0,
		},
		Timezone: &Timezone{
			ID:     "America/New_York",
			Offset: -240,
		},
		WebGL: &WebGL{
			Vendor:                 "Apple Inc.",
			Renderer:               "Apple GPU",
			UnmaskedVendor:         "Apple Inc.",
			UnmaskedRenderer:       "Apple A16 GPU",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (1.0)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts:               []string{"SF Pro Text", "SF Pro Display", "Helvetica Neue"},
		Plugins:             []Plugin{},
		HardwareConcurrency: 6,
		DeviceMemory:        6,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        false,
			ChargingTime:    0,
			DischargingTime: 14400,
			Level:           0.80,
		},
	}
}

// Mobile iOS preset (Chrome on iPhone)
func NewChrome119iOS() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/119.0.6045.109 Mobile/15E148 Safari/604.1",
		Platform:  "iPhone",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            390,
			Height:           844,
			AvailWidth:       390,
			AvailHeight:      844,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 3.0,
		},
		Timezone: &Timezone{
			ID:     "America/New_York",
			Offset: -240,
		},
		WebGL: &WebGL{
			Vendor:                 "Apple Inc.",
			Renderer:               "Apple GPU",
			UnmaskedVendor:         "Apple Inc.",
			UnmaskedRenderer:       "Apple A16 GPU",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (1.0)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts:               []string{"SF Pro Text", "SF Pro Display", "Helvetica Neue"},
		Plugins:             []Plugin{},
		HardwareConcurrency: 6,
		DeviceMemory:        6,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        false,
			ChargingTime:    0,
			DischargingTime: 14400,
			Level:           0.80,
		},
	}
}

// NewChrome134Android returns a Chrome 134 fingerprint for Android 14 (Pixel 8a)
func NewChrome134Android() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Linux; Android 14; Pixel 8a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Mobile Safari/537.36",
		Platform:  "Linux armv8l",
		Vendor:    "Google Inc.",
		Language:  "en-US",
		Languages: []string{"en-US", "en"},
		Screen: &Screen{
			Width:            1080,
			Height:           2400,
			AvailWidth:       1080,
			AvailHeight:      2340,
			ColorDepth:       24,
			PixelDepth:       24,
			DevicePixelRatio: 2.75,
		},
		Timezone: &Timezone{
			ID:     "Europe/Berlin",
			Offset: -60,
		},
		WebGL: &WebGL{
			Vendor:                 "Google Inc. (Qualcomm)",
			Renderer:               "ANGLE (Qualcomm, Adreno (TM) 740, OpenGL ES 3.2)",
			UnmaskedVendor:         "Qualcomm",
			UnmaskedRenderer:       "Adreno (TM) 740",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts: []string{
			"Roboto", "Noto Sans", "Droid Sans", "sans-serif", "serif", "monospace",
		},
		Plugins:             []Plugin{},
		HardwareConcurrency: 8,
		DeviceMemory:        8,
		Audio: &Audio{
			Noise: 0.01,
		},
		Battery: &Battery{
			Charging:        false,
			ChargingTime:    0,
			DischargingTime: 7200,
			Level:           0.87,
		},
	}
}

// NewChrome134Windows11 returns a Chrome 134 fingerprint for Windows 11
func NewChrome134Windows11() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
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
			Vendor:                 "Google Inc. (NVIDIA)",
			Renderer:               "ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Direct3D11 vs_5_0 ps_5_0, D3D11)",
			UnmaskedVendor:         "NVIDIA Corporation",
			UnmaskedRenderer:       "NVIDIA GeForce RTX 3060/PCIe/SSE2",
			ShadingLanguageVersion: "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)",
		},
		Canvas: &Canvas{
			Noise: 0.01,
		},
		WebRTC: &WebRTC{
			Disable: false,
		},
		Fonts: []string{
			"Arial", "Calibri", "Cambria", "Consolas", "Courier New",
			"Georgia", "Segoe UI", "Times New Roman", "Verdana", "Trebuchet MS",
		},
		Plugins:             []Plugin{},
		HardwareConcurrency: 12,
		DeviceMemory:        16,
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

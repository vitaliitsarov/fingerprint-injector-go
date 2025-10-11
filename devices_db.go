package fingerprint

// DeviceDatabase содержит базу реальных устройств
type DeviceDatabase struct {
	Browsers []BrowserVersion
	Devices  []DeviceSpec
	GPUs     []GPUSpec
	OSes     []OSVersion
}

// BrowserVersion информация о версии браузера
type BrowserVersion struct {
	Name    string
	Version string
	Major   int
}

// DeviceSpec спецификация устройства
type DeviceSpec struct {
	Name          string
	Type          string // "desktop", "mobile", "tablet"
	Platform      string
	CPUCores      []int // Возможные варианты ядер
	RAM           []int // Возможные варианты RAM (GB)
	ScreenWidths  []int
	ScreenHeights []int
	DPRs          []float64
}

// GPUSpec спецификация видеокарты
type GPUSpec struct {
	Vendor   string
	Renderer string
	Type     string // "desktop", "mobile"
}

// OSVersion версия операционной системы
type OSVersion struct {
	Name     string
	Platform string
	Versions []string
}

// GetDeviceDatabase возвращает базу данных устройств
func GetDeviceDatabase() *DeviceDatabase {
	return &DeviceDatabase{
		Browsers: []BrowserVersion{
			{Name: "Chrome", Version: "119.0.0.0", Major: 119},
			{Name: "Chrome", Version: "120.0.0.0", Major: 120},
			{Name: "Chrome", Version: "121.0.0.0", Major: 121},
			{Name: "Chrome", Version: "122.0.0.0", Major: 122},
			{Name: "Firefox", Version: "120.0", Major: 120},
			{Name: "Firefox", Version: "121.0", Major: 121},
			{Name: "Safari", Version: "17.0", Major: 17},
			{Name: "Safari", Version: "17.1", Major: 17},
		},
		Devices: []DeviceSpec{
			// Desktop Windows
			{
				Name:          "Windows Desktop",
				Type:          "desktop",
				Platform:      "Win32",
				CPUCores:      []int{4, 6, 8, 12, 16},
				RAM:           []int{8, 16, 32, 64},
				ScreenWidths:  []int{1920, 2560, 3840, 1680, 1600},
				ScreenHeights: []int{1080, 1440, 2160, 1050, 900},
				DPRs:          []float64{1.0, 1.25, 1.5, 2.0},
			},
			// Desktop MacOS
			{
				Name:          "MacBook Pro",
				Type:          "desktop",
				Platform:      "MacIntel",
				CPUCores:      []int{8, 10, 12},
				RAM:           []int{8, 16, 32, 64},
				ScreenWidths:  []int{1440, 1680, 1920, 2560, 2880},
				ScreenHeights: []int{900, 1050, 1200, 1440, 1800},
				DPRs:          []float64{2.0},
			},
			// Desktop Linux
			{
				Name:          "Linux Desktop",
				Type:          "desktop",
				Platform:      "Linux x86_64",
				CPUCores:      []int{4, 6, 8, 12, 16, 24, 32},
				RAM:           []int{8, 16, 32, 64, 128},
				ScreenWidths:  []int{1920, 2560, 3840, 1680},
				ScreenHeights: []int{1080, 1440, 2160, 1050},
				DPRs:          []float64{1.0, 1.5, 2.0},
			},
			// Mobile - iPhone
			{
				Name:          "iPhone 14 Pro",
				Type:          "mobile",
				Platform:      "iPhone",
				CPUCores:      []int{6},
				RAM:           []int{6},
				ScreenWidths:  []int{393},
				ScreenHeights: []int{852},
				DPRs:          []float64{3.0},
			},
			{
				Name:          "iPhone 15",
				Type:          "mobile",
				Platform:      "iPhone",
				CPUCores:      []int{6},
				RAM:           []int{6},
				ScreenWidths:  []int{393},
				ScreenHeights: []int{852},
				DPRs:          []float64{3.0},
			},
			{
				Name:          "iPhone 13",
				Type:          "mobile",
				Platform:      "iPhone",
				CPUCores:      []int{6},
				RAM:           []int{4},
				ScreenWidths:  []int{390},
				ScreenHeights: []int{844},
				DPRs:          []float64{3.0},
			},
			{
				Name:          "iPhone 12",
				Type:          "mobile",
				Platform:      "iPhone",
				CPUCores:      []int{6},
				RAM:           []int{4},
				ScreenWidths:  []int{390},
				ScreenHeights: []int{844},
				DPRs:          []float64{3.0},
			},
			// Mobile - Android
			{
				Name:          "Samsung Galaxy S23",
				Type:          "mobile",
				Platform:      "Linux armv8l",
				CPUCores:      []int{8},
				RAM:           []int{8, 12},
				ScreenWidths:  []int{360},
				ScreenHeights: []int{780},
				DPRs:          []float64{3.0},
			},
			{
				Name:          "Google Pixel 8",
				Type:          "mobile",
				Platform:      "Linux armv8l",
				CPUCores:      []int{8},
				RAM:           []int{8},
				ScreenWidths:  []int{412},
				ScreenHeights: []int{915},
				DPRs:          []float64{2.625},
			},
			{
				Name:          "OnePlus 11",
				Type:          "mobile",
				Platform:      "Linux armv8l",
				CPUCores:      []int{8},
				RAM:           []int{8, 12, 16},
				ScreenWidths:  []int{412},
				ScreenHeights: []int{919},
				DPRs:          []float64{3.0},
			},
			{
				Name:          "Xiaomi 13",
				Type:          "mobile",
				Platform:      "Linux armv8l",
				CPUCores:      []int{8},
				RAM:           []int{8, 12},
				ScreenWidths:  []int{393},
				ScreenHeights: []int{873},
				DPRs:          []float64{2.75},
			},
			// Tablets
			{
				Name:          "iPad Pro",
				Type:          "tablet",
				Platform:      "iPad",
				CPUCores:      []int{8},
				RAM:           []int{8, 16},
				ScreenWidths:  []int{1024},
				ScreenHeights: []int{1366},
				DPRs:          []float64{2.0},
			},
			{
				Name:          "Samsung Galaxy Tab",
				Type:          "tablet",
				Platform:      "Linux armv8l",
				CPUCores:      []int{8},
				RAM:           []int{6, 8},
				ScreenWidths:  []int{800, 1024},
				ScreenHeights: []int{1280, 1600},
				DPRs:          []float64{2.0, 2.5},
			},
		},
		GPUs: []GPUSpec{
			// Desktop - NVIDIA
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 4090", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 4080", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 4070", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 3090", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 3080", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce RTX 3070", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce GTX 1660 Ti", Type: "desktop"},
			{Vendor: "NVIDIA Corporation", Renderer: "NVIDIA GeForce GTX 1080 Ti", Type: "desktop"},

			// Desktop - AMD
			{Vendor: "AMD", Renderer: "AMD Radeon RX 7900 XTX", Type: "desktop"},
			{Vendor: "AMD", Renderer: "AMD Radeon RX 7800 XT", Type: "desktop"},
			{Vendor: "AMD", Renderer: "AMD Radeon RX 6900 XT", Type: "desktop"},
			{Vendor: "AMD", Renderer: "AMD Radeon RX 6800 XT", Type: "desktop"},
			{Vendor: "AMD", Renderer: "AMD Radeon RX 5700 XT", Type: "desktop"},

			// Desktop - Intel
			{Vendor: "Intel Inc.", Renderer: "Intel(R) UHD Graphics 770", Type: "desktop"},
			{Vendor: "Intel Inc.", Renderer: "Intel(R) UHD Graphics 730", Type: "desktop"},
			{Vendor: "Intel Inc.", Renderer: "Intel(R) UHD Graphics 630", Type: "desktop"},
			{Vendor: "Intel Inc.", Renderer: "Intel(R) Iris Xe Graphics", Type: "desktop"},

			// Desktop - Apple
			{Vendor: "Apple Inc.", Renderer: "Apple M3 Pro", Type: "desktop"},
			{Vendor: "Apple Inc.", Renderer: "Apple M2 Pro", Type: "desktop"},
			{Vendor: "Apple Inc.", Renderer: "Apple M1 Pro", Type: "desktop"},
			{Vendor: "Apple Inc.", Renderer: "Apple M1", Type: "desktop"},

			// Mobile - Qualcomm
			{Vendor: "Qualcomm", Renderer: "Adreno (TM) 740", Type: "mobile"},
			{Vendor: "Qualcomm", Renderer: "Adreno (TM) 730", Type: "mobile"},
			{Vendor: "Qualcomm", Renderer: "Adreno (TM) 650", Type: "mobile"},
			{Vendor: "Qualcomm", Renderer: "Adreno (TM) 640", Type: "mobile"},

			// Mobile - Apple
			{Vendor: "Apple Inc.", Renderer: "Apple A17 Pro GPU", Type: "mobile"},
			{Vendor: "Apple Inc.", Renderer: "Apple A16 GPU", Type: "mobile"},
			{Vendor: "Apple Inc.", Renderer: "Apple A15 GPU", Type: "mobile"},
			{Vendor: "Apple Inc.", Renderer: "Apple A14 GPU", Type: "mobile"},

			// Mobile - ARM Mali
			{Vendor: "ARM", Renderer: "Mali-G710", Type: "mobile"},
			{Vendor: "ARM", Renderer: "Mali-G78", Type: "mobile"},
			{Vendor: "ARM", Renderer: "Mali-G77", Type: "mobile"},
		},
		OSes: []OSVersion{
			{
				Name:     "Windows",
				Platform: "Win32",
				Versions: []string{"10.0", "11.0"},
			},
			{
				Name:     "macOS",
				Platform: "MacIntel",
				Versions: []string{"14.0", "13.6", "13.5", "12.7"},
			},
			{
				Name:     "Linux",
				Platform: "Linux x86_64",
				Versions: []string{"Ubuntu 22.04", "Ubuntu 20.04", "Fedora 39", "Debian 12"},
			},
			{
				Name:     "iOS",
				Platform: "iPhone",
				Versions: []string{"17.0", "17.1", "16.7", "16.6"},
			},
			{
				Name:     "Android",
				Platform: "Linux armv8l",
				Versions: []string{"14", "13", "12", "11"},
			},
		},
	}
}

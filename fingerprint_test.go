package fingerprint

import (
	"encoding/json"
	"testing"
)

func TestNewDefaultFingerprint(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.UserAgent == "" {
		t.Error("UserAgent should not be empty")
	}

	if fp.Platform == "" {
		t.Error("Platform should not be empty")
	}

	if fp.Screen == nil {
		t.Error("Screen should not be nil")
	}

	if fp.Timezone == nil {
		t.Error("Timezone should not be nil")
	}

	if fp.WebGL == nil {
		t.Error("WebGL should not be nil")
	}
}

func TestFingerprintToJSON(t *testing.T) {
	fp := NewDefaultFingerprint()
	
	jsonStr, err := fp.ToJSON()
	if err != nil {
		t.Errorf("ToJSON failed: %v", err)
	}

	if jsonStr == "" {
		t.Error("JSON string should not be empty")
	}

	// Проверяем, что это валидный JSON
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		t.Errorf("Invalid JSON: %v", err)
	}
}

func TestPresets(t *testing.T) {
	presets := []struct {
		name string
		fp   *Fingerprint
	}{
		{"Windows11", NewChrome119Windows11()},
		{"MacOS", NewChrome119MacOS()},
		{"Linux", NewChrome119Linux()},
		{"Android", NewChrome119Android()},
	}

	for _, preset := range presets {
		t.Run(preset.name, func(t *testing.T) {
			if preset.fp.UserAgent == "" {
				t.Errorf("%s: UserAgent should not be empty", preset.name)
			}

			if preset.fp.Platform == "" {
				t.Errorf("%s: Platform should not be empty", preset.name)
			}

			if preset.fp.Screen == nil {
				t.Errorf("%s: Screen should not be nil", preset.name)
			}

			if preset.fp.Screen.Width <= 0 {
				t.Errorf("%s: Screen width should be positive", preset.name)
			}

			if preset.fp.HardwareConcurrency <= 0 {
				t.Errorf("%s: HardwareConcurrency should be positive", preset.name)
			}
		})
	}
}

func TestScreenValidation(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.Screen.Width <= 0 || fp.Screen.Height <= 0 {
		t.Error("Screen dimensions should be positive")
	}

	if fp.Screen.ColorDepth <= 0 {
		t.Error("ColorDepth should be positive")
	}

	if fp.Screen.DevicePixelRatio <= 0 {
		t.Error("DevicePixelRatio should be positive")
	}
}

func TestTimezone(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.Timezone.ID == "" {
		t.Error("Timezone ID should not be empty")
	}
}

func TestWebGL(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.WebGL.Vendor == "" {
		t.Error("WebGL Vendor should not be empty")
	}

	if fp.WebGL.Renderer == "" {
		t.Error("WebGL Renderer should not be empty")
	}
}

func TestCanvasNoise(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.Canvas.Noise < 0 || fp.Canvas.Noise > 1 {
		t.Error("Canvas noise should be between 0 and 1")
	}
}

func TestBattery(t *testing.T) {
	fp := NewDefaultFingerprint()

	if fp.Battery.Level < 0 || fp.Battery.Level > 1 {
		t.Error("Battery level should be between 0 and 1")
	}
}


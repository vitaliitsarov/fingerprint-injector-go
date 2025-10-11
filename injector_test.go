package fingerprint

import (
	"strings"
	"testing"
)

func TestNewInjector(t *testing.T) {
	fp := NewDefaultFingerprint()
	injector := NewInjector(fp)

	if injector == nil {
		t.Error("Injector should not be nil")
	}

	if injector.fingerprint == nil {
		t.Error("Injector fingerprint should not be nil")
	}
}

func TestGetInjectionScript(t *testing.T) {
	fp := NewDefaultFingerprint()
	injector := NewInjector(fp)

	script := injector.GetInjectionScript()

	if script == "" {
		t.Error("Injection script should not be empty")
	}

	// Проверяем наличие ключевых частей скрипта
	requiredParts := []string{
		"navigator.userAgent",
		"navigator.platform",
		"navigator.vendor",
		"navigator.language",
		"screen.width",
		"screen.height",
		"navigator.webdriver",
	}

	for _, part := range requiredParts {
		if !strings.Contains(script, part) {
			t.Errorf("Script should contain '%s'", part)
		}
	}
}

func TestGetInjectionScriptWithWebRTCDisabled(t *testing.T) {
	fp := NewDefaultFingerprint()
	fp.WebRTC.Disable = true

	injector := NewInjector(fp)
	script := injector.GetInjectionScript()

	if !strings.Contains(script, "RTCPeerConnection") {
		t.Error("Script should contain WebRTC disable code")
	}
}

func TestGetInjectionScriptWithCanvasNoise(t *testing.T) {
	fp := NewDefaultFingerprint()
	fp.Canvas.Noise = 0.5

	injector := NewInjector(fp)
	script := injector.GetInjectionScript()

	if !strings.Contains(script, "toDataURL") {
		t.Error("Script should contain Canvas noise code")
	}
}

func TestToJSArray(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"en-US", "en"}, "['en-US', 'en']"},
		{[]string{"ru"}, "['ru']"},
		{[]string{}, "[]"},
	}

	for _, test := range tests {
		result := toJSArray(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

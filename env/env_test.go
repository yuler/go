package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Exist", func(t *testing.T) {
		os.Setenv("TEST_VAR", "test value")
		expected := Get("TEST_VAR")
		if expected != "test value" {
			t.Errorf("Expected 'test value', got '%s'", expected)
		}
	})

	t.Run("Panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		os.Unsetenv("NON_EXISTENT_VAR")
		Get("NON_EXISTENT_VAR")
	})
}

func TestGetDefault(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		os.Unsetenv("NON_EXISTENT_VAR")
		if GetDefault("NON_EXISTENT_VAR", "default") != "default" {
			t.Error("GetDefault() failed")
		}
	})
}

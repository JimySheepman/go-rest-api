package env

import "testing"

func TestLoadEnvironmentConfigure(t *testing.T) {
	t.Run("read env file", func(t *testing.T) {
		loaded, _ := LoadEnvironmentConfigure("../../.env")
		expected := true

		if loaded != expected {
			t.Errorf("expected %v but got %v", expected, loaded)
		}
	})

	t.Run("read env file error", func(t *testing.T) {
		loaded, _ := LoadEnvironmentConfigure("../.env")
		expected := false

		if loaded != expected {
			t.Errorf("expected %v but got %v", expected, loaded)
		}
	})
}

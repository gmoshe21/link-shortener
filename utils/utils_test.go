package utils

import (
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("generate", func(t *testing.T) {
			result := GenerateShortUrl();
			if len(result) != 10 {
				t.Error("fail generate")
			}
		})
	}
}
package test

import (
	"permission-cat/internal/web-save"
	"testing"
)

func TestFetcher(t *testing.T) {
	err := web_save.Run("https://www.baidu.com", false)
	if err != nil {
		return
	}
}

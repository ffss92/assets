package assets

import (
	"fmt"
	"path/filepath"
	"strings"
)

func buildScriptTag(path, file string) string {
	return fmt.Sprintf(`<script type="module" src="%s/%s"></script>`, path, file)
}

func buildCssTag(path, file string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s/%s">`, path, file)
}

func isCss(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".css", ".sass", ".scss", ".pcss", ".less", ".postcss", ".styl", ".stylus":
		return true
	default:
		return false
	}
}

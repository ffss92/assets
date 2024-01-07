package assets

import (
	"encoding/json"
	"errors"
	"html/template"
	"os"
	"strings"
)

var (
	ErrUnknownMode = errors.New("unknown assets mode")
	ErrNoManifest  = errors.New("manifest was not loaded")
)

type Mode int

const (
	// Indicates that the application is running in dev mode.
	// Assets will be resolved using DevURL.
	ModeDevelopment Mode = iota
	// Indicates that the application is running in production mode.
	// Assets will be resolved using StaticURL.
	ModeProduction
)

type Vite struct {
	ManifestPath string
	StaticURL    string
	DevURL       string
	Mode         Mode

	manifest Manifest
}

// Resolve files and builds HTML tags for the corresponding asset type. Currently
// supports css and js assets.
func (v Vite) Resolve(files ...string) (template.HTML, error) {
	var builder strings.Builder

	switch v.Mode {
	// Dev mode, use vite url as base.
	case ModeDevelopment:
		path := strings.TrimSuffix(v.DevURL, "/")
		builder.WriteString(buildScriptTag(path, "@vite/client"))

		for _, file := range files {
			switch {
			case isCss(file):
				builder.WriteString(buildCssTag(path, file))
			default:
				builder.WriteString(buildScriptTag(path, file))
			}
		}

	// Prod mode, use static url as base, check for manifest entries.
	case ModeProduction:
		if v.manifest == nil {
			err := v.LoadManifest()
			if err != nil {
				return template.HTML(""), err
			}

			path := strings.TrimSuffix(v.StaticURL, "/")
			for _, file := range files {
				entry, ok := v.manifest[file]
				if !ok {
					continue
				}

				// Build tags for asset css
				for _, css := range entry.Css {
					builder.WriteString(buildCssTag(path, css))
				}

				// Build tag for entry
				switch {
				case isCss(entry.File):
					builder.WriteString(buildCssTag(path, entry.File))
				default:
					builder.WriteString(buildScriptTag(path, entry.File))
				}
			}
		}
	default:
		return template.HTML(""), ErrUnknownMode
	}

	tags := builder.String()
	return template.HTML(tags), nil
}

// Attempts to load the manifest file into memory. This is only required for
// production mode. If the current Mode is ModeDevelopment, this function will
// return nil.
func (v *Vite) LoadManifest() error {
	if v.Mode == ModeDevelopment {
		return nil
	}

	f, err := os.Open(v.ManifestPath)
	if err != nil {
		return err
	}
	defer f.Close()

	manifest := Manifest{}
	err = json.NewDecoder(f).Decode(&manifest)
	if err != nil {
		return err
	}

	v.manifest = manifest
	return nil
}

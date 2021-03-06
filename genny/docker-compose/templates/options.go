package docker_compose

import (
	"strings"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/gobuffalo/buffalo/runtime"
	"github.com/pkg/errors"
)

// Options for generating a new docker file
type Options struct {
	App     meta.App `json:"app"`
	Version string   `json:"version"`
	Style   string   `json:"style"`
}

// Validate options
func (opts *Options) Validate() error {
	if strings.ToLower(opts.Style) == "none" {
		return errors.New("style was none - generator should not be used")
	}
	if (opts.App == meta.App{}) {
		opts.App = meta.New(".")
	}
	if opts.Version == "" {
		opts.Version = runtime.Version
	}
	if opts.Style == "" {
		opts.Style = "deps"
	}
	if _, ok := boxes[opts.Style]; !ok {
		return errors.Errorf("unknown Docker Compose style: %s", opts.Style)
	}

	return nil
}

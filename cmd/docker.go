package cmd

import (
	"strings"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/gobuffalo/buffalo/runtime"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Options for generating a new docker file
type Options struct {
	App     meta.App `json:"app"`
	Version string   `json:"version"`
	Style   string   `json:"style"`
	AsWeb   bool     `json:"as_web"`
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
		opts.Style = "multi"
	}
	opts.AsWeb = opts.App.WithWebpack
	if _, ok := boxes[opts.Style]; !ok {
		return errors.Errorf("unknown Docker style: %s", opts.Style)
	}

	return nil
}

var boxes = map[string]packr.Box{
	"standard": packr.NewBox("./standard/templates"),
	"multi":    packr.NewBox("./multi/templates"),
}

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "generates Dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}

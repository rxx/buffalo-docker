package docker_compose

import (
	"text/template"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

var boxes = map[string]packr.Box{
	"deps": packr.NewBox("./templates/deps"),
	"full": packr.NewBox("./templates/full"),
}

// New generator for docker-compose
func New(opts *Options) (*genny.Generator, error) {
	if err := opts.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	g := genny.New()
	box, ok := boxes[opts.Style]
	if !ok {
		return g, errors.Errorf("unknown Docker Compose style: %s", opts.Style)
	}
	g.Box(box)

	data := map[string]interface{}{
		"opts": opts,
	}
	t := gotools.TemplateTransformer(data, template.FuncMap{})
	g.Transformer(t)
	return g, nil
}

package docker_compose

import (
	"context"
	"testing"

	"github.com/gobuffalo/genny"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{})
	r.NoError(err)
	r.NotNil(g)

	run := genny.DryRunner(context.Background())
	run.With(g)
	r.NoError(run.Run())

	res := run.Results()
	r.Len(res.Commands, 0)
	r.Len(res.Files, 1)

	f := res.Files[0]
	r.Equal("docker-compose.yml", f.Name())
}

func Test_New_deps(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{
		Style: "deps",
	})
	r.NoError(err)
	r.NotNil(g)

	run := genny.DryRunner(context.Background())
	run.With(g)
	r.NoError(run.Run())

	res := run.Results()
	r.Len(res.Commands, 0)
	r.Len(res.Files, 1)

	f := res.Files[0]
	r.Equal("docker-compose.yml", f.Name())
}

func Test_New_none(t *testing.T) {
	r := require.New(t)

	_, err := New(&Options{
		Style: "none",
	})
	r.Error(err)
}

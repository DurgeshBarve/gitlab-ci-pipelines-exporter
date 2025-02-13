package controller

import (
	"testing"

	"github.com/DurgeshBarve/gitlab-ci-pipelines-exporter/pkg/config"
	"github.com/DurgeshBarve/gitlab-ci-pipelines-exporter/pkg/schemas"
)

func TestTriggerRefMetricsPull(_ *testing.T) {
	ctx, c, _, srv := newTestController(config.Config{})
	srv.Close()

	ref1 := schemas.Ref{
		Project: schemas.NewProject("group/foo"),
		Name:    "main",
	}

	p2 := schemas.NewProject("group/bar")
	ref2 := schemas.Ref{
		Project: p2,
		Name:    "main",
	}

	c.Store.SetRef(ctx, ref1)
	c.Store.SetProject(ctx, p2)

	// TODO: Assert results somehow
	c.triggerRefMetricsPull(ctx, ref1)
	c.triggerRefMetricsPull(ctx, ref2)
}

func TestTriggerEnvironmentMetricsPull(_ *testing.T) {
	ctx, c, _, srv := newTestController(config.Config{})
	srv.Close()

	p1 := schemas.NewProject("foo/bar")
	env1 := schemas.Environment{
		ProjectName: p1.Name,
		Name:        "dev",
	}

	env2 := schemas.Environment{
		ProjectName: "foo/baz",
		Name:        "prod",
	}

	c.Store.SetProject(ctx, p1)
	c.Store.SetEnvironment(ctx, env1)
	c.Store.SetEnvironment(ctx, env2)

	// TODO: Assert results somehow
	c.triggerEnvironmentMetricsPull(ctx, env1)
	c.triggerEnvironmentMetricsPull(ctx, env2)
}

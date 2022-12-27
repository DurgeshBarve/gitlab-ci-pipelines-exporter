package main

import (
	"os"

	"github.com/DurgeshBarve/gitlab-ci-pipelines-exporter/internal/cli"
)

var version = "Durgesh_01"

func main() {
	cli.Run(version, os.Args)
}

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/k1LoW/octocov/internal"
)

func (c *Config) Build() {
	// Repository
	if c.Repository == "" {
		c.Repository = os.Getenv("GITHUB_REPOSITORY")
	}

	// Coverage
	if c.Coverage == nil {
		c.Coverage = &Coverage{}
	}
	if c.Coverage.Paths == nil {
		c.Coverage.Paths = []string{}
	}
	if c.Coverage.Path != "" {
		_, _ = fmt.Fprintln(os.Stderr, "Deprecated: coverage.path: has been deprecated. please use coverage.paths: instead.")
		c.Coverage.Paths = append(c.Coverage.Paths, c.Coverage.Path)
	}
	if len(c.Coverage.Paths) == 0 {
		c.Coverage.Paths = append(c.Coverage.Paths, filepath.Dir(c.path))
	} else {
		paths := []string{}
		for _, p := range c.Coverage.Paths {
			paths = append(paths, filepath.Join(filepath.Dir(c.path), p))
		}
		c.Coverage.Paths = paths
	}

	// CodeToTestRatio
	if c.CodeToTestRatio != nil {
		if c.CodeToTestRatio.Code == nil {
			c.CodeToTestRatio.Code = []string{}
		}
		if c.CodeToTestRatio.Test == nil {
			c.CodeToTestRatio.Test = []string{}
		}
	}

	// TestExecutionTime
	if c.TestExecutionTime == nil {
		c.TestExecutionTime = &TestExecutionTime{}
	}

	// Report

	// Central
	if c.Central != nil {
		if c.Central.Root == "" {
			c.Central.Root = "."
		}
		if !strings.HasPrefix(c.Central.Root, "/") {
			c.Central.Root = filepath.Clean(filepath.Join(c.Root(), c.Central.Root))
		}
		if len(c.Central.Reports.Datastores) == 0 {
			c.Central.Reports.Datastores = append(c.Central.Reports.Datastores, defaultReportsDatastore)
		}
		if len(c.Central.Badges.Datastores) == 0 {
			c.Central.Badges.Datastores = append(c.Central.Badges.Datastores, defaultBadgesDatastore)
		}
	}

	// Push

	// Comment

	// Diff

	// GitRoot
	gitRoot, _ := internal.GetRootPath(c.Root())
	c.GitRoot = gitRoot
}

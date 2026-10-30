// Package spark is the Sparki Go sample: it passes the Go starter
// (go build ./... → go test ./...) on golang:1.22.
package spark

import (
	"errors"
	"regexp"
	"strings"
)

// ErrNoSteps is returned for a build with no steps.
var ErrNoSteps = errors.New("no steps")

// Ignite reports whether a build lights: it does only if every step passes.
func Ignite(steps []string) (bool, error) {
	if len(steps) == 0 {
		return false, ErrNoSteps
	}
	for _, s := range steps {
		if s != "pass" {
			return false, nil
		}
	}
	return true, nil
}

var nonSlug = regexp.MustCompile(`[^a-z0-9]+`)

// Slug makes a url-safe name.
func Slug(name string) string {
	return strings.Trim(nonSlug.ReplaceAllString(strings.ToLower(strings.TrimSpace(name)), "-"), "-")
}

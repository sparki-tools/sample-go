package spark

import (
	"errors"
	"testing"
)

func TestIgnite(t *testing.T) {
	cases := []struct {
		name  string
		steps []string
		want  bool
	}{
		{"every step passes", []string{"pass", "pass", "pass"}, true},
		{"one step fails", []string{"pass", "fail", "pass"}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Ignite(c.steps)
			if err != nil || got != c.want {
				t.Fatalf("Ignite(%v) = %v, %v; want %v, nil", c.steps, got, err, c.want)
			}
		})
	}
}

func TestIgniteNoSteps(t *testing.T) {
	if _, err := Ignite(nil); !errors.Is(err, ErrNoSteps) {
		t.Fatalf("want ErrNoSteps, got %v", err)
	}
}

func TestSlug(t *testing.T) {
	if got := Slug("  Hello, Sparki! "); got != "hello-sparki" {
		t.Fatalf("Slug = %q", got)
	}
}

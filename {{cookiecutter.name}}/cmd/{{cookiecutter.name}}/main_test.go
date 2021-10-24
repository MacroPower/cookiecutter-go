package main_test

import (
	"strings"
	"testing"

	main "{{cookiecutter.vcs_root}}/{{cookiecutter.name}}/cmd/{{cookiecutter.name}}"
)

func TestMain(t *testing.T) {
	t.Parallel()

	sb := strings.Builder{}
	main.Hello(&sb)

	if want := "Hello World!"; sb.String() != want {
		t.Fatalf("expected %s, got %s", want, sb.String())
	}
}

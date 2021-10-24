package main_test

import (
	"strings"
	"testing"

	main "{{cookiecutter.vcs_root}}/{{cookiecutter.name}}/cmd/{{cookiecutter.name}}"
)

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sb := strings.Builder{}
		main.Hello(&sb)
		sb.Reset()
	}
}

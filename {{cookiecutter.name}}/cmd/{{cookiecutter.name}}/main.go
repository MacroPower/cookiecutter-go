package main

import (
	"fmt"
	"io"
	"strings"{% if cookiecutter.use_kong|lower == "y" %}

	{% endif %}{% if cookiecutter.use_kong|lower == "y" %}"github.com/alecthomas/kong"{% endif %}
)
{% if cookiecutter.use_kong|lower == "y" %}
var cli struct {
	Log struct {
		Level  string `help:"Log level." default:"info"`
		Format string `help:"Log format. One of: [logfmt, json]" default:"logfmt"`
	} `prefix:"log." embed:""`
}
{% endif %}
func main() {
	{% if cookiecutter.use_kong|lower == "y" %}cliCtx := kong.Parse(&cli, kong.Name("{{cookiecutter.name}}"))

	{% endif %}sb := strings.Builder{}
	Hello(&sb)
	fmt.Println(sb.String())
}

func Hello(r io.Writer) {
	_, err := r.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}

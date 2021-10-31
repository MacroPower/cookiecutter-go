package main

import (
	{%- if cookiecutter.use_kong|lower != "y" %}
	"fmt"
	{%- endif %}
	"io"
	"strings"
	{%- if cookiecutter.use_gokit_log|lower == "y" %}

	"{{cookiecutter.vcs_root}}/{{cookiecutter.org}}/{{cookiecutter.name}}/internal/log"
	{%- endif %}
	{%- if cookiecutter.use_kong|lower == "y" %}

	"github.com/alecthomas/kong"
	{%- endif %}
)
{% if cookiecutter.use_kong|lower == "y" %}
var cli struct {
	{%- if cookiecutter.use_gokit_log|lower == "y" %}
	Log struct {
		Level  string `help:"Log level." default:"info"`
		Format string `help:"Log format. One of: [logfmt, json]" default:"logfmt"`
	} `prefix:"log." embed:""`
	{%- endif %}
}
{% endif %}
func main() {
	{% if cookiecutter.use_kong|lower == "y" -%}
	cliCtx := kong.Parse(&cli, kong.Name("{{cookiecutter.name}}"))

	{% if cookiecutter.use_gokit_log|lower == "y" -%}
	logLevel := &log.AllowedLevel{}
	if err := logLevel.Set(cli.Log.Level); err != nil {
		cliCtx.FatalIfErrorf(err)
	}

	logFormat := &log.AllowedFormat{}
	if err := logFormat.Set(cli.Log.Format); err != nil {
		cliCtx.FatalIfErrorf(err)
	}

	logger := log.New(&log.Config{
		Level:  logLevel,
		Format: logFormat,
	})
	{%- endif %}
	{%- else -%}
	{% if cookiecutter.use_gokit_log|lower == "y" -%}
	logger := log.New(&log.Config{})
	{%- endif %}
	{%- endif -%}

	{% if cookiecutter.use_gokit_log|lower == "y" %}

	err := log.Info(logger).Log("msg", "Starting {{cookiecutter.name}}")
	{% if cookiecutter.use_kong|lower == "y" %}cliCtx.FatalIfErrorf(err){% else %}panic(err){% endif %}

	{% endif -%}

	sb := strings.Builder{}
	Hello(&sb)
	{% if cookiecutter.use_kong|lower == "y" %}cliCtx.Printf("%s", sb.String()){% else %}fmt.Println(sb.String()){% endif %}
}

func Hello(r io.Writer) {
	_, err := r.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}

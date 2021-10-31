module {{cookiecutter.vcs_root}}/{{cookiecutter.org}}/{{cookiecutter.name}}

go {{cookiecutter.go_version}}
{%- if cookiecutter.use_kong|lower == "y" or cookiecutter.use_gokit_log|lower == "y" %}

require (
    {%- if cookiecutter.use_kong|lower == "y" %}
    github.com/alecthomas/kong v0.2.17
    {%- endif %}
    {%- if cookiecutter.use_gokit_log|lower == "y" %}
    github.com/go-kit/log v0.2.0
    {%- endif %}
)
{%- endif %}

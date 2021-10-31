module {{cookiecutter.vcs_root}}/{{cookiecutter.org}}/{{cookiecutter.name}}

go {{cookiecutter.go_version}}{% if cookiecutter.kong|lower == "y" %}

require (
    {% if cookiecutter.kong|lower == "y" %}github.com/alecthomas/kong v0.2.17{% endif %}
){% endif %}

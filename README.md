# cookiecutter-go

A [Cookiecutter](https://github.com/cookiecutter/cookiecutter) for my Go projects.

Features:

- Configured linter (using [golangci](https://golangci-lint.run)).
- Development dependency handling (using [bingo](https://github.com/bwplotka/bingo)).
- Build and release automation (using [goreleaser](https://goreleaser.com)).
- Benchmarks (using [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)).
- Makefile with help, format, lint, test, bench, and build targets.

Optional features:

- **gh_actions**: Enable to create GitHub Actions.
- **vscode**: Enable to create VS Code launch and container config.

## Thanks

Some inspiration & snippets were taken from the following projects:

- [prometheus](https://github.com/prometheus)
- [thanos](https://github.com/thanos-io/thanos)
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang)

# cookiecutter-go

A [Cookiecutter](https://github.com/cookiecutter/cookiecutter) for my Go projects.

Features:

- Configured linter (using [golangci](https://golangci-lint.run)).
- Development dependency handling (using [bingo](https://github.com/bwplotka/bingo)).
- Build and release automation (using [goreleaser](https://goreleaser.com)).
- Benchmarks (using [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)).
- Makefile with help, format, lint, test, bench, and build targets.

Optional features:

- **use_gh_actions**: Enable to create GitHub Actions.
- **use_vscode**: Enable to create VS Code launch and container config.
- **use_kong**: Enable to add [Kong](https://github.com/alecthomas/kong) command-line parser.
- **use_gokit_log**: Enable to add [go-kit/log](https://github.com/go-kit/log) logger.

## Coming soon

- More sensible linter configuration.
- Version handling (waiting on Go 1.18).
- Optional Prometheus collector.
- Store and compare benchmarks via [benchdiff](https://github.com/WillAbides/benchdiff).

## Contributing

Please feel free to submit issues and/or PRs with fixes or updates. However,
note that this repo represents my personal standards and opinions. I will not
accept anything that I don't agree with or that I think I will not personally
use, even if it's an optional feature.

This is purely to make this Cookiecutter no more complex than I personally need
it to be.

If you would like to contribute something that you're not sure about or think
might be divisive, please just ask about it in an issue and we can talk about
it. Or, if you want it just for yourself, fork the repo and do whatever you'd
like!

## Thanks

Some inspiration & snippets were taken from the following projects:

- [prometheus](https://github.com/prometheus)
- [thanos](https://github.com/thanos-io/thanos)
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang)

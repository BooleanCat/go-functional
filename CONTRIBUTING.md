# Contributing to go-functional

Contributions to this project are very welcome! This guide should help with instructions on how to
submit changes. Contributions can be made in the form of GitHub
[issues](https://github.com/BooleanCat/go-functional/issues) or
[pull requests](https://github.com/BooleanCat/go-functional/pulls).

When submitting an issue, please choose the relevant template or choose a blank issue if your query
doesn't naturally fit into an existing template. The pull request template contains a contribution
checklist.

## Zero-dependency

This project is a zero-dependency project - which means that consumers using this project's packages
must only incur one dependency: go-functional.

Development dependencies are OK as they will not be included as dependencies to end-users (such as
`golangci-lint`).

## Development dependencies

1. [golangci-lint](https://github.com/golangci/golangci-lint) is used to lint the project when
   running `make check`.
2. [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) is used to generate new fakes
   (using `go generate ./...`). Fakes are declared in `internal/fakes/fakes.go`.

## Commit hygiene

- Commits should contain only a single change
- Commit messages must use imperative language (e.g. `Add iter.Fold collection function`)
- Commit messages must explain what is changed, not how it is changed
- The first line of a commit message should be a terse description of the change containing 72
  characters or fewer

## Running tests

Run tests with `make test` from the project root directory.

Tests are written using Go's `testing` package and helpers are available in `internal/assert`.

Code is linted using `golangci-lint`. The linter may be run using `make lint`.

## Different types of changes

### Bug fixes

Bug reports are appreciated ahead of bug fixes as early reporting allows the community to be aware
of any issues ahead of a fix being submitted. If you intend to fix a bug after reporting, that is
greatly appreciated - just make sure to mention you intend to work on it on the issue report so the
maintainers are aware and leave you the chance to make a contribution.

When submitting a bug fix PR, a test must be added (or an existing test modified) that exposes the
bug and your change must make that test pass.

### New features

Issues should be opened ahead of submitting a PR to added a new feature. This is to prevent you
wasting your time should a feature not be desirable and allows others to have input into the
conversation.

All new functionality must be fully tested and new public functions must include an
[`Example` test](https://go.dev/blog/examples) that will be used by the reference docs to
demonstrate its use.

Mark pull requests as "Draft" if you intend to use the pull request as a workspace but are not yet
ready to receive unsolicited feedback on specifics like commit messages or failing tests.

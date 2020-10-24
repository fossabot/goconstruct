<div align="center">
  <img src="./docs/imgs/logo.png"><br>
  <p align="center">Go project construction from opinionated templates.</p>
</div>

[![Build Status][build-badge]][build-url]
[![GoDev][godev-badge]][godev-url]
[![License][license-badge]][license-url]
[![codecov][codecov-badge]][codecov-url]
[![Release][release-badge]][release-url]

[build-badge]: https://github.com/mccurdyc/goconstruct/workflows/build-test/badge.svg
[build-url]: https://github.com/mccurdyc/goconstruct/actions
[godev-badge]: https://pkg.go.dev/badge/github.com/mccurdyc/goconstruct
[godev-url]: https://pkg.go.dev/github.com/mccurdyc/goconstruct?tab=overview
[license-badge]: https://img.shields.io/github/license/mccurdyc/goconstruct
[license-url]: LICENSE
[codecov-badge]: https://codecov.io/gh/mccurdyc/goconstruct/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/mccurdyc/goconstruct
[release-badge]: https://img.shields.io/github/release/mccurdyc/goconstruct.svg
[release-url]: https://github.com/mccurdyc/goconstruct/releases/latest

## Overview

goconstruct provides a way to construct Go projects from opinionated, templated,
components.

## Installing

```sh
(
  d=$(mktemp -d) cd $d
  go get -u -v github.com/mccurdyc/goconstruct
)
```

## Usage

```sh
% goconstruct project -h
USAGE
  project <subcommand>

SUBCOMMANDS
  generate  Generates a new project.

FLAGS
  -config config.toml       A config file defining values for the required variables for all templates used.
  -destination .            The destination directory where the project should be created.
  -templates glue           A comma-separate list of template names.
  -templates-path ./_tmpls  Path to templates.
```

## Contributing

TODO

## License

+ [GNU General Public License Version 3](./LICENSE)

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmccurdyc%2Fgoconstruct.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmccurdyc%2Fgoconstruct?ref=badge_large)

## TODOs

- Handle `goconstruct` with missing subcommands.
- GitHub Action for CI (build and test).
- Backport changes to goconstruct to the template files.
- Render templated filenames.
- Configurable template delimeters.
- Move templates to a separate GitHub repository.
- Generate README usage section via Make target or something.

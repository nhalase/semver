# semver

## Usage

```shell
Usage:
  semver [command]

Available Commands:
  get         Extract the given part of <version>. Part may be release.
  help        Help about any command
```

### Commands

#### get

Extract the given part of <version>, where part is one of major, minor, patch,
pre, build, or release.

The extracted part will be written to stdout.
If <version> is not a valid semantic version, an error will be written to
stderr.

```shell
Usage:
  semver get [command]

Examples:
get [major|minor|patch|pre|build|release] <version>

Available Commands:
  build       Extract the build metadata part of <version>.
  major       Extract the MAJOR part of <version>.
  minor       Extract the MINOR part of <version>.
  patch       Extract the PATCH part of <version>.
  pre         Extract the pre-release part of <version>.
  release     Extract only MAJOR.MINOR.PATCH.
```

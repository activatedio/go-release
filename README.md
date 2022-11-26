# go-release

## Overview

Build wrapper to automate version tagging and incrementing. Supports the
following standard build flow:

* Run verification build (optional)
* Tag git with current version
* Run release build (if applicable)
* Increment the version and push to git remote

## Install

Install the package as follows:

```
go get -u github.com/activatedio/go-release
```

Create a file `.version` in your go project with the current version.

```
v1.0.0
```

Create a g`go-release.yaml` file in your project with build defintions:

``` yaml
---
go-release:
  verify: go test ./...
```

## Configuration

Support yaml values are:

| Name | Description |
| ---- | ----------- |
| `go-release.verify` | Command to run before tagging the project |
| `go-release.perform` | Command to run to perform release |

The `go-release.verify` command is recommeded to check the project via
compliation and any automated tests.

The `go-release.perform` command is used for projects which require a build step
to release. Projects such as go modules generally to not require this.

## Usage

To check the release before performing

```
go-release check
```

Perform the release
```
go-release perform
```

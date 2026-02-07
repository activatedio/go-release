> ## go-release
>
> Build wrapper to automate version tagging and incrementing. Supports the
> following standard build flow:
>

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/activatedio/go-release/ci.yaml?branch=main&style=flat-square)](https://github.com/activatedio/go-release/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/activatedio/go-release?style=flat-square)](https://goreportcard.com/report/github.com/activatedio/go-release)


## Overeview

* Run verification build (optional)
* Tag git with current version
* Run release build (if applicable)
* Increment the version and commit
* Push to origin (optional)

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
verify: go test ./...
```

## Configuration

Support yaml values are:

| Name | Description |
| ---- | ----------- |
| `verify` | Command to run before tagging the project |
| `perform` | Command to run to perform release |
| `skip-push` | Skip push to origin |

The `verify` command is recommeded to check the project via compliation and any
automated tests.

The `perform` command is used for projects which require a build step to
release. Projects such as go modules generally to not require this.

## Usage

To check the release before performing

```
go-release check
```

Perform the release
```
go-release perform
```

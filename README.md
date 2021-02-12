# MrAndreID / Go Log

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/golog.svg)](https://pkg.go.dev/github.com/MrAndreID/golog) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoLog` package is a collection of functions in the go language for Log.

---

## Table of Contents

* [Install](#install)
* [Usage](#usage)
* [Full Example](#full-example)
* [Versioning](#versioning)
* [Authors](#authors)
* [License](#license)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [More](#more)

---

## Install

To use The `MrAndreID/GoLog` package, you must follow the steps below:

```sh
go get -u github.com/MrAndreID/golog
```

## Usage

### Error

```go
golog.Error("primary", "Andrea Adam")
golog.Error("secondary", "Andrea Adam")
```

Output:

```sh
2021/02/12 19:09:02 [ Error ] Andrea Adam.
2021/02/12 19:09:02 [ Error ] -> Andrea Adam.
```

### Success

```go
golog.Success("primary", "Andrea Adam")
golog.Success("secondary", "Andrea Adam")
```

Output:

```sh
2021/02/12 19:09:02 [ Success ] Andrea Adam.
2021/02/12 19:09:02 [ Success ] -> Andrea Adam.
```

### Warning

```go
golog.Warning("primary", "Andrea Adam")
golog.Warning("secondary", "Andrea Adam")
```

Output:

```sh
2021/02/12 19:09:02 [ Warning ] Andrea Adam.
2021/02/12 19:09:02 [ Warning ] -> Andrea Adam.
```

### Info

```go
golog.Info("primary", "Andrea Adam")
golog.Info("secondary", "Andrea Adam")
```

Output:

```sh
2021/02/12 19:09:02 [ Info ] Andrea Adam.
2021/02/12 19:09:02 [ Info ] -> Andrea Adam.
```

## Full Example

Full Example can be found on the [Go Playground website](https://play.golang.com/p/luIJzFDShiy).

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/golog).

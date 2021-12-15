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

### Init

To use Init Func., The `MrAndreID/GoLog` package will add/modify .gitignore file and add `logs/` and the last one will create a `logs` folder. Inside the `logs` folder will contain the log files based on the log file creation date. The contents of the first parameter is used to `limit` the number of log files created other than the day's log files. The second parameter (`logLevel`) are `all`, `error`, `success`, `warning`, & `info`. The second parameter (`logLevel`) is used for the log level to be stored in the log file. The third parameter (`timezone`) is used to set the time zone used. The fourth parameter (`style`) is used to set the style. Style is to add color.

```go
golog.Init(30, "all", "Asia/Jakarta", false)
```

### Error

```go
golog.Error("Andrea Adam")
```

Output:

```sh
2021-02-15 18:45:18 [ ERROR ] Andrea Adam
```

### Success

```go
golog.Success("Andrea Adam")
```

Output:

```sh
2021-02-15 18:45:18 [ SUCCESS ] Andrea Adam
```

### Warning

```go
golog.Warning("Andrea Adam")
```

Output:

```sh
2021-02-15 18:45:18 [ WARNING ] Andrea Adam
```

### Info

```go
golog.Info("Andrea Adam")
```

Output:

```sh
2021-02-15 18:45:18 [ INFO ] Andrea Adam
```

## Full Example

Full Example can be found on the [Go Playground website](https://play.golang.com/p/zk2GlYUvClU).

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

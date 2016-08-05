# govvv

The simple Go binary versioning tool that wraps the `go build` command. 

![](https://cl.ly/0U2m441v392Q/intro-1.gif)

Stop worrying about `-ldflags` and **`go get github.com/ahmetalpbalkan/govvv`** now.

## Build Variables

| Variable | Description | Example |
|----------|-------------|---------|
| **`main.GitCommit`** | short commit hash of source tree | `0b5ed7a` |
| **`main.GitBranch`** | current branch name the code is built off | `master` |
| **`main.GitState`** | whether there are uncommitted changes | `clean` or `dirty` | 
| **`main.BuildDate`** | RFC3339 formatted UTC date | `2016-08-04T18:07:54Z` |
| **`main.Version`** | contents of `./VERSION` file, if exists | `2.0.0` |

## Using govvv is easy

Just add the build variables you want to the `main` package and run:

| old          | :sparkles: new :sparkles: |
| -------------|-----------------|
| `go build`   | `govvv build`   |
| `go install` | `govvv install` | 

## Version your app with govvv

Create a `VERSION` file in your build root directory and add a `Version`
variable to your `main` package.

![](https://cl.ly/3Q1K1R2D3b2K/intro-2.gif)

Do you have your own way of specifying `Version`? No problem:

## govvv lets you specify custom `-ldflags`

Your existing `-ldflags` argument will still be preserved:

    govvv build -ldflags "-X main.BuildNumber=$buildnum" myapp

and the `-ldflags` constructed by govvv will be appended to your flag.

## Try govvv today

    $ go get github.com/ahmetalpbalkan/govvv

------

govvv is distributed under [Apache 2.0 License](LICENSE).

Copyright 2016 Ahmet Alp Balkan 

------

[![Build Status](https://travis-ci.org/ahmetalpbalkan/govvv.svg?branch=master)](https://travis-ci.org/ahmetalpbalkan/govvv)
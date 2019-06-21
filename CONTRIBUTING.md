# Contributing

## Requirements

* [npm](https://www.npmjs.com/): to validate a commit message and generate the Change Log
* [Golang](https://golang.org/)
* [golangci-lint](https://github.com/golangci/golangci-lint)

```console
$ npm i
```

## Test

```console
$ npm t
```

## Commit Message Format

The commit message format of this project conforms to the [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).
We validate the commit message with git's `commit-msg` hook using [commitlint](http://marionebl.github.io/commitlint/#/) and [husky](https://www.npmjs.com/package/husky).

## Coding Guide

* https://github.com/golang/go/wiki/CodeReviewComments

go vet

```console
$ npm run vet
```

golangci-lint

```console
$ npm run lint
```

gofmt

```console
$ npm run fmt
```

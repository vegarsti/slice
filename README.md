# slice

<p>
    <a href="https://github.com/vegarsti/slice/releases"><img src="https://img.shields.io/github/release/vegarsti/slice.svg" alt="Latest Release"></a>
    <a href="https://github.com/vegarsti/slice/actions"><img src="https://github.com/vegarsti/slice/workflows/test/badge.svg" alt="Build Status"></a>
    <a href="http://goreportcard.com/report/github.com/vegarsti/slice"><img src="http://goreportcard.com/badge/vegarsti/slice" alt="Go ReportCard"></a>
</p>

Like [`cut`](https://en.wikipedia.org/wiki/Cut_(Unix)), but with the negative slicing we know and love.

## Installation

```sh
$ go get "github.com/vegarsti/slice"
```


## Usage

```sh
$ echo "hello world" | slice 1:-1
ello worl

$ echo "hello\nyou" | slice 1:-1
ell
o

$ echo "hello\nyou" | slice 1:
ello
ou

$ echo "hello\nyou" | slice :2
he
yo
```

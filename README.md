# slice

Like [`cut`](https://en.wikipedia.org/wiki/Cut_(Unix)), but with the negative slicing we know and love.

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

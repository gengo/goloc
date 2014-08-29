Goloc
=====

Goloc is a tool for counting the number of statements in a Go file.

### Installation & Usage

To install Goloc, simply run:

```
go get github.com/gengo/goloc/goloc
```

This installs goloc locally, and if `$GOPATH/bin` is in your `$PATH`, you will now be able to run the `goloc` command:

```
$ goloc --help

Goloc is a tool for counting the number of statements in go files.

Usage:

	cat path/to/file | goloc
```

Goloc reads from stdin, so the easiest is to `cat` a file and feed it to goloc:

```
$ cat myfile.go | goloc
20
```

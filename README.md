# Playground and cheatsheet for learning Golang

> This is a collection of Golang that are split by [topics](#table-of-contents) and contain code examples with explanations, different use cases and links to further readings.


It is a **playground** because you may change or add the code to see how it works.
Altogether it might make your learning process to be more interactive and it might help you to keep
code quality pretty high from very beginning.

It is a **cheatsheet** because you may get back to these code examples once you want to recap the
syntax of [standard Golang statements and constructions](#table-of-contents).

## How to Use This Repository

This repository is source code that belong to a [Video tutorial](https://www.bilibili.com/video/av24365381?from=search&seid=7202638767983264130) about Golang. The video is maked by a Google engineer. If you can understand Chinese,you can watch it.If you don't know Chinese, don't worry, just write the code.

So normally you might want to do the following:

- [Find the topic](#table-of-contents) you want to learn or recap.
- Look at code examples and assertions to see usage examples and expected output.
- Change code or add new assertions to see how things work.
- [Run tests](#testing-the-code) and [lint the code](#linting-the-code) to see if it work and is
written correctly.

## Table of Contents

1. **Getting Started**
    - [What is Golang](https://golang.org/)
    - [Golang Syntax](https://blog.golang.org/gos-declaration-syntax)
2. **Basic**
    - [atomic](./basic/atomic)
    - [basic](./basic/basic) (`var-declare`,`enum`,`const`,`zeroValue`,`initialValue`,`typeDeduction`)
    - [branch](./basic/branch) (`if-else`, `switch-case`)
    - [loop](./basic/loop) (`for`)
    - [regex](./basic/regex)
3. **Data Types**
    - [Channel](./channel) (including select)
    - [array](./container/arrays)
    - [map](./container/maps)
    - [slice](./container/slices)
    - [string](./container/strings)
4. **Functional programming**
    - [functional](./functional)
5. **Err handiling**
    - [defer](./errhandiling/defer)
    - [recover](./errhandiling/recover)
    - [file listing server](./errhandiling/filelistingserver) (Practice project)
6. **Goroutine**
    - [Goroutine](./goroutine)
7. **Http**
    - [http](./http) (http request)
8. **Reflect**
    - [reflect](./reflect)
9. **Tcp**
    - [tcp](./tcp) (basic using)
    - [chatrom](./chatroom)(a practice about tcp-programming)
10. **RPC**
    - [rpc](./rpc)(basic using)
11. **Interface**
    - [interface](./retriever) (`duck-type`,`interface`)
12. **Algorithms**
    - [data-structure](./algorithms) (`tree`,`queue`,`linklist`,`hashtable`.eg)
    - [algorithms](./algorithms) (`sort`,`josephu`.eg)
13. **Comprehensive practical**
    - [crawler](./crawler)(`single-task`,`Concurrent-crawler`)
    - [crawler_distributed](./crawler_distributed)(`distributed`)

## Prerequisites

**Installing Golang**

Make sure that you have [Golang installed](https://golang.org/doc/install) on your machine.

You may check your Golang version by running:

```bash
go version
```

Make sure that you have [Golang environment variable](https://golang.org/doc/code.html) on your machine.

**Installing dependencies**

Install all dependencies that are required for the project by running:

```bash
go get github.com/golang/example/hello
```

## Testing the Code

Tests are made using **go-test comand**.

You may add new tests for yourself by adding files and functions with `test_` prefix
(i.e. `test_topic.go` with `func TestSubTopic(t *testing.T)` function inside).

To run all the tests please execute the following command from the project root folder:

```bash
go test
```

To run specific tests please execute:

```bash
go test ./test_topic.go
```

## Linting the Code

Linting is done using [golint](https://github.com/golang/lint).

### Golint

To check if the code is written with respect
to Golang style guide please run:

```bash
golint ./src/
```

In case if linter will detect error (i.e. `missing-docstring`) you may want to read more about
specific error by running:

```bash
main.go:5:6: exported type Hero should have comment or be unexported
```

[More about Golint](https://github.com/golang/lint)

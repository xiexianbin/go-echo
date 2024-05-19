# go-echo-demo

[![build](https://github.com/xiexianbin/go-echo-demo/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/go-echo-demo/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/go-echo-demo?status.svg)](https://pkg.go.dev/github.com/xiexianbin/go-echo-demo)
[![Go Report Card](https://goreportcard.com/badge/github.com/xiexianbin/go-echo-demo)](https://goreportcard.com/report/github.com/xiexianbin/go-echo-demo)

## Feature

- [swaggo/echo-swagger](https://github.com/swaggo/echo-swagger) echo middleware to automatically generate RESTful API documentation with Swagger 2.0.
- Prometheus metrics([ref](https://www.xiexianbin.cn/monitor/prometheus/clientlibs/))
- [go-playground/validator](https://github.com/go-playground/validator) [detail](https://echo.labstack.com/docs/request#validate-data)

## usage

- release

```
git tag -s v0.1.0
git push origin v0.1.0
# git push origin --tags
```

- download

```
curl -Lfs -o main https://github.com/xiexianbin/go-echo-demo/releases/latest/download/main-{linux|darwin|windows}-{amd64|arm64}
chmod +x main
./main
```

## reference

- https://echo.labstack.com/
- [xiexianbin/go-actions-demo](https://github.com/xiexianbin/go-actions-demo) for github actions
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md) reference only, not officially recognized by [proposal: doc: make an "official" go-standards/project-layout documentation within Go project](https://github.com/golang/go/issues/45861)
- [mattn/go-generics-example](https://github.com/mattn/go-generics-example) Example code for Go generics

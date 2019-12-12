# Package puzh [![Build Status](https://travis-ci.com/puzh/puzh-go.svg?branch=master)](https://travis-ci.com/puzh/puzh-go) [![GoDoc](https://godoc.org/github.com/puzh/puzh-go?status.svg)](https://godoc.org/github.com/puzh/puzh-go) [![Go Report Card](https://goreportcard.com/badge/github.com/puzh/puzh-go)](https://goreportcard.com/report/github.com/puzh/puzh-go)

```go
import "github.com/puzh/puzh-go"
```


## Index
- [func Init(token string)](#func-init)
- [func It(a ...interface{}) error](#func-it)
- [func Itf(format string, a ...interface{}) error](#func-itf)
- [type Puzh](#type-puzh)
    - [func New(token string, silent bool) *Puzh](#func-new)
    - [func (p *Puzh) It(a ...interface{}) error](#func-puzh-it)
    - [func (p *Puzh) Itf(format string, a ...interface{}) error](#func-puzh-itf)


## func [Init](puzh.go#L25)
```go
func Init(token string)
```
Init initializes the package with the given authentication token.


## func [It](puzh.go#L31)
```go
func It(a ...interface{}) error
```
It formats its parameters analogous to fmt.Sprint and sends the resulting string to the  Telegram [@puzhbot](https://t.me/puzhbot).


## func [Itf](puzh.go#L40)
```go
func Itf(format string, a ...interface{}) error
```
Itf formats its parameters analogous to fmt.Sprintf and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).


## type [Puzh](puzh.go#L48)
```go
type Puzh struct {
	// contains filtered or unexported fields
}
```
Puzh implements the [puzh.it](https://puzh.it) API.


### func [New](puzh.go#L56)
```go
func New(token string, silent bool) *Puzh
```
New returns a new instance of Puzh given an authentication token.


### func (*Puzh) [It](puzh.go#L73)
```go
func (p *Puzh) It(a ...interface{}) error
```
It formats its parameters analogous to fmt.Sprint and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).


### func (*Puzh) [Itf](puzh.go#L79)
```go
func (p *Puzh) Itf(format string, a ...interface{}) error
```
Itf formats its parameters analogous to fmt.Sprintf and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).

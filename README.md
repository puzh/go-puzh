# Package puzh [![Build Status](https://travis-ci.com/puzh/go-puzh.svg?branch=master)](https://travis-ci.com/puzh/go-puzh) [![GoDoc](https://godoc.org/github.com/puzh/go-puzh?status.svg)](https://godoc.org/github.com/puzh/go-puzh) [![Go Report Card](https://goreportcard.com/badge/github.com/puzh/go-puzh)](https://goreportcard.com/report/github.com/puzh/go-puzh)

```go
import "github.com/puzh/go-puzh"
```


## Index
- [Variables](#variables)
- [func Init(token string)](#func-init)
- [func It(a ...interface{}) error](#func-it)
- [func Itf(format string, a ...interface{}) error](#func-itf)
- [type Puzh](#type-puzh)
    - [func NewPuzh(token string) *Puzh](#func-newpuzh)
    - [func (p *Puzh) It(a ...interface{}) error](#func-puzh-it)
    - [func (p *Puzh) Itf(format string, a ...interface{}) error](#func-puzh-itf)


## Variables
```go
var (
	// ErrNotInitialized indicates that the package has not been initialized.
	ErrNotInitialized = errors.New("puzh: package not initialized")
)
```


## func [Init](puzh.go#L26)
```go
func Init(token string)
```
Init initializes the package.


## func [It](puzh.go#L32)
```go
func It(a ...interface{}) error
```
It formats its parameters analogous to fmt.Sprint and sends the resulting string to the  Telegram [@puzhbot](https://t.me/puzhbot).


## func [Itf](puzh.go#L41)
```go
func Itf(format string, a ...interface{}) error
```
Itf formats its parameters analogous to fmt.Sprintf and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).


## type [Puzh](puzh.go#L49)
```go
type Puzh struct {
	// contains filtered or unexported fields
}
```
Puzh implements the [puzh.it](https://puzh.it) API.


## func [NewPuzh](puzh.go#L56)
```go
func NewPuzh(token string) *Puzh
```
NewPuzh returns a new instance of Puzh given an authentication token.


## func (*Puzh) [It](puzh.go#L69)
```go
func (p *Puzh) It(a ...interface{}) error
```
It formats its parameters analogous to fmt.Sprint and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).


## func (*Puzh) [Itf](puzh.go#L75)
```go
func (p *Puzh) Itf(format string, a ...interface{}) error
```
Itf formats its parameters analogous to fmt.Sprintf and sends the resulting string to the Telegram [@puzhbot](https://t.me/puzhbot).

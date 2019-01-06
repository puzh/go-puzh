package puzh

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	// ErrNotInitialized indicates that the package has not been initialized.
	ErrNotInitialized = errors.New("puzh: package not initialized")

	p *Puzh
)

func init() {
	if token := os.Getenv("PUZH_TOKEN"); token != "" {
		p = NewPuzh(token)
	}
}

// Init initializes the package.
func Init(token string) {
	p = NewPuzh(token)
}

// It formats its parameters analogous to fmt.Sprint and sends the resulting
// string to the Telegram @puzhbot.
func It(a ...interface{}) error {
	if p == nil {
		return ErrNotInitialized
	}
	return p.It(a...)
}

// Itf formats its parameters analogous to fmt.Sprintf and sends the resulting
// string to the Telegram @puzhbot.
func Itf(format string, a ...interface{}) error {
	if p == nil {
		return ErrNotInitialized
	}
	return p.Itf(format, a...)
}

// Puzh implements the puzh.it API.
type Puzh struct {
	client *http.Client

	token string
}

// NewPuzh returns a new instance of Puzh given an authentication token.
func NewPuzh(token string) *Puzh {
	tr := &http.Transport{
		MaxIdleConns:    1,
		IdleConnTimeout: 30 * time.Second,
	}
	return &Puzh{
		client: &http.Client{Transport: tr},
		token:  token,
	}
}

// It formats its parameters analogous to fmt.Sprint and sends the resulting
// string to the Telegram @puzhbot.
func (p *Puzh) It(a ...interface{}) error {
	return p.it(fmt.Sprint(a...))
}

// Itf formats its parameters analogous to fmt.Sprintf and sends the resulting
// string to the Telegram @puzhbot.
func (p *Puzh) Itf(format string, a ...interface{}) error {
	return p.it(fmt.Sprintf(format, a...))
}

func (p *Puzh) it(message string) error {
	data := url.Values{
		"token": []string{p.token},
		"text":  []string{message},
	}
	_, err := p.client.PostForm("https://api.puzh.it", data)
	return err
}

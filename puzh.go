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
	errNotInitialized = errors.New("puzh: package not initialized")

	p *Puzh
)

func init() {
	if token := os.Getenv("PUZH_TOKEN"); token != "" {
		p = New(token, false)
	}
}

// Init initializes the package with the given authentication token.
func Init(token string) {
	p = New(token, false)
}

// It formats its parameters analogous to fmt.Sprint and sends the resulting
// string to the Telegram @puzhbot.
func It(a ...interface{}) error {
	if p == nil {
		return errNotInitialized
	}
	return p.It(a...)
}

// Itf formats its parameters analogous to fmt.Sprintf and sends the resulting
// string to the Telegram @puzhbot.
func Itf(format string, a ...interface{}) error {
	if p == nil {
		return errNotInitialized
	}
	return p.Itf(format, a...)
}

// Puzh implements the puzh.it API.
type Puzh struct {
	client *http.Client

	token  string // api token
	silent bool   // send messages silently, without sound
}

// New returns a new instance of Puzh given an authentication token.
func New(token string, silent bool) *Puzh {
	return &Puzh{
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:    1,
				IdleConnTimeout: 30 * time.Second,
			},
		},

		token:  token,
		silent: silent,
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
	if p.silent {
		data["silent"] = []string{"true"}
	}

	_, err := p.client.PostForm("https://api.puzh.it", data)
	return err
}

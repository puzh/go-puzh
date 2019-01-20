package puzh

import (
	"os"

	"github.com/puzh/puzh"
)

func ExampleItf() {
	puzh.Init("my-token")

	puzh.Itf("Hi, *%s*", os.Getenv("USER"))
}

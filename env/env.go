package env

import (
	"flag"
)

var BindAddress = flag.String("bindAddress", ":8001", "Bind address for the server")

func Parse() {
	flag.Parse()
}

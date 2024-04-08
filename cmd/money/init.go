package main

import (
	"flag"
)

var (
	a, r, w, t *string
)

func init() {
	r = flag.String("r", "", "RUN_ADDRESS")
	w = flag.String("w", "", "WORKPLACE_DB")
	a = flag.String("a", "", "AUTH_DB")
	t = flag.String("t", "", "AUTH_TYPE")
}

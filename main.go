package main

import (
	"github.com/miyachi000/go-openssl-cmd-wrapper/v1/openssl"
	"github.com/miyachi000/go-openssl-cmd-wrapper/v1/openssl/genpkey"
)

func main() {
	out, err := openssl.Genpkey(genpkey.Algorithm(genpkey.ED25519), genpkey.Out("outfile.pkey")))

	if err != nil {
		print(err)
	}
	print(out)
}
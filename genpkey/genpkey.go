package genpkey

import (
	"fmt"

	"github.com/miyachi000/go-openssl-cmd-wrapper/v1/types"
)

type algorithm string

const (
	RSA     Algorithm = "RSA"
	RSAPSS            = "RSA-PSS"
	EC                = "EC"
	X25519            = "X25519"
	X448              = "X448"
	ED25519           = "ED25519"
	ED448             = "ED448"
)

type outform string

const (
	DER outform = "DER"
	PEM         = "PEM"
)

// Verbose output "status dots" while generating keys.
// -verbose
func Verbose(g *types.Cmd) {
	g.AddOptions("-verbose")
}

// Out output the key to the specified file. If this argument is not specified then standard output is used.
// -out <filename>
func Out(filename string) func(*type.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("-out %s", filename))
	}
}

// Outform The output format, except when -genparam is given; the default is PEM.
// -outform DER|PEM
func Outform(form outform) func(*type.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("-outform %s", form))
	}
}

// Algorithm
// -algorithm <alg>
func Algorithm(alg algorithm) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("-algorithm %s", alg))
	}
}

// Pkeyopt
// -pkeyopt opt:value
func Pkeyopt(opt, value string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(fmt.Sprintf("-pkeyopt %s:%s", opt, value))
	}
}
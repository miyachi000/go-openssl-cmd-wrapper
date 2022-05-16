package openssl

import (
	"context"
	"io"
	"log"

	"github.com/miyachi000/go-openssl-cmd-wrapper/v1/types"
)

// Genpkey
func Genpkey(options ...types.Option) (string, error) {
	return command(context.Background(), "genpkey", options...)
}

func command(ctx context.Context, name string, options ...types.Option) (string, error) {
	g := types.NewCmd(name)
	g.ApplyOptions(options...)

	return g.Exec(ctx, g.Base, g.Debug, g.Options...)
}
//go:build unix && !darwin && (amd64 || arm64)

package ffi

import (
	_ "embed"
)

const libName = "libffi.so"

//go:embed lib/libffi.8.so
var libData []byte

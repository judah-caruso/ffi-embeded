//go:build darwin && arm64

package ffi

import (
	_ "embed"
)

const libName = "libffi.dylib"

//go:embed lib/libffi.8.dylib
var libData []byte

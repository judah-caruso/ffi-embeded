//go:build windows && (amd64 || arm64)

package ffi

import (
	_ "embed"
)

const libName = "libffi.dll"

//go:embed lib/libffi-8.dll
var libData []byte

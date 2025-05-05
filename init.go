//go:generate go run lib/generate.go
//go:build (freebsd || linux || windows || darwin) && (amd64 || arm64)

package ffi

import (
	"os"
)

func init() {
	file, err := os.CreateTemp("", "*-"+libName)
	if err != nil {
		panic(err)
	}

	path := file.Name()
	if _, err := file.Write(libData); err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}

	libffi, err := Load(path)
	if err != nil {
		panic(err)
	}

	os.Remove(path)

	prepCif, err = libffi.Get("ffi_prep_cif")
	if err != nil {
		panic(err)
	}

	prepCifVar, err = libffi.Get("ffi_prep_cif_var")
	if err != nil {
		panic(err)
	}

	call, err = libffi.Get("ffi_call")
	if err != nil {
		panic(err)
	}

	closureAlloc, err = libffi.Get("ffi_closure_alloc")
	if err != nil {
		panic(err)
	}

	closureFree, err = libffi.Get("ffi_closure_free")
	if err != nil {
		panic(err)
	}

	prepClosureLoc, err = libffi.Get("ffi_prep_closure_loc")
	if err != nil {
		panic(err)
	}
}

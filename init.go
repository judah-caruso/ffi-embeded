//go:build (freebsd || linux || windows || darwin) && (amd64 || arm64)

package ffi

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func init() {
	name := fmt.Sprintf("%d-%s", time.Now().UnixNano(), libName)
	path := filepath.Join(os.TempDir(), name)

	if err := os.WriteFile(path, libData, 0644); err != nil {
		panic(err)
	}

	defer os.Remove(path)

	libffi, err := Load(path)
	if err != nil {
		panic(err)
	}

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

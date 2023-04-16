package gptj

// #cgo LDFLAGS: -L./build/ -L./build/gpt4all-chat -lm -lstdc++
// #include "binding.h"
import "C"

type GPTJParams struct {
	params *C.struct_gptj_hparams
}

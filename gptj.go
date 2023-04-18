package gptj

// #cgo CXXFLAGS: -I./gpt4all-chat/ggml/include
// #cgo LDFLAGS: -L./build -lbinding -L./build/gpt4all-chat/ggml/src -lggml -L./build/gpt4all-chat/ggml/examples -lggml_utils -lm -lstdc++
// #include "binding.h"
import "C"

type GptJParams struct {
	params *C.struct_gptj_hparams
}

type GptJContext struct {
	context *C.struct_PromptContext
}

type GptJ struct {
	handle C.GPTJHandle
}

func Load(path string) GptJ {
	cstr := C.CString(path)
	handle := &C.go_gptj_load(cstr)
	return GptJ{handle}
}

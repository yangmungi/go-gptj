package gptj

/*
#cgo CXXFLAGS: -I./gpt4all-chat/ggml/include
#cgo LDFLAGS: -L./build -lbinding -L./build/gpt4all-chat/ggml/src -lggml -L./build/gpt4all-chat/ggml/examples -lggml_utils -lm -lstdc++

#include "binding.h"

extern int cgo_concatString(const char*);
*/
import "C"
import "fmt"

type GptJ struct {
	handle *C.struct_GPTJContext
}

func Load(path string) *GptJ {
	cstr := C.CString(path)
	handle := C.go_gptj_load(cstr)
	return &GptJ{handle}
}

var globalString string

//export concatString
func concatString(s *C.char) C.int {
	chunk := C.GoString(s)
	globalString += chunk
	fmt.Printf("%s", chunk)
	return 0
}

func (g *GptJ) Prompt(prompt string) string {
	globalString = ""
	C.go_gptj_prompt(*g.handle, C.CString(prompt), C.Callback(C.cgo_concatString), 200, 40, 0.9, 0.5, 9)
	return globalString
}

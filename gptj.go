package gptj

/*
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
	if handle == nil {
		return nil
	}

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

type PromptOptions struct {
	Tokens int
	TopK   int
	TopP   float64
	Temp   float64
	Batch  int
}

func (g *GptJ) Prompt(prompt string, po PromptOptions) string {
	globalString = ""
	if po.Tokens == 0 {
		po.Tokens = 200
	}

	if po.TopK == 0 {
		po.TopK = 100
	}

	if po.TopP == 0 {
		po.TopP = 0.9
	}

	if po.Temp == 0 {
		po.Temp = 1.0
	}

	if po.Batch == 0 {
		po.Batch = 9
	}

	C.go_gptj_prompt(*g.handle, C.CString(prompt), C.Callback(C.cgo_concatString),
		C.int32_t(po.Tokens), C.int32_t(po.TopK),
		C.float(po.TopP), C.float(po.Temp),
		(C.int32_t)(po.Batch))

	return globalString
}

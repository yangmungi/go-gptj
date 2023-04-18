package main

import (
	"flag"
	"fmt"

	"github.com/yangmungi/go-gptj"
)

func main() {
	var modelPath string
	var promptor string
	flag.StringVar(&modelPath, "model-path", "./", "GPT-J model path")
	flag.StringVar(&promptor, "prompt", "Hello!", "Prompt")
	flag.Parse()

	gptjHandle := gptj.Load(modelPath)
	fmt.Printf("[%+v] %s \n", gptjHandle, promptor)

	response := gptjHandle.Prompt(promptor)
	fmt.Printf("\n%s\n", response)
}

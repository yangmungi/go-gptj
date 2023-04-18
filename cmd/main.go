package main

import (
	"flag"
	"fmt"

	"github.com/yangmungi/go-gptj"
)

func main() {
	var modelPath string
	flag.StringVar(&modelPath, "model-path", "./", "GPT-J model path")
	flag.Parse()

	gptjHandle := gptj.Load(modelPath)

	fmt.Printf("%+v\n", gptjHandle)

}

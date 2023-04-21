package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yangmungi/go-gptj"
)

type RespBody struct {
	Response string
	Elapsed  time.Duration
}

type ReqBody struct {
	Prompt  string
	Options gptj.PromptOptions
}

func main() {
	var modelPath string
	var addr string
	flag.StringVar(&modelPath, "model-path", "./", "GPT-J model path")
	flag.StringVar(&addr, "addr", ":8089", "Server address")

	flag.Parse()

	gptjHandle := gptj.Load(modelPath)
	if gptjHandle == nil {
		fmt.Println()
		return
	}

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading", http.StatusBadRequest)
			return
		}

		var rBody ReqBody
		err = json.Unmarshal(body, &rBody)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		log.Printf("%+v", rBody)
		start := time.Now()
		responsed := gptjHandle.Prompt(rBody.Prompt, rBody.Options)
		elapsed := time.Now().Sub(start)

		enc := json.NewEncoder(w)

		err = enc.Encode(RespBody{responsed, elapsed})
		if err != nil {
			log.Println(err)
		}

	})

	log.Printf("Start server %s", addr)
	err := http.ListenAndServe(addr, m)
	if err != nil {
		log.Print(err)
	}
}

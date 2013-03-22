package main

import (
	"log"
	"net/http"
	"webpipes"
)

func EchoPipe(inputs map[string]string) (map[string]string, error) {
	return inputs, nil
}

func main() {
	block := webpipes.Block{
		"echo",
		"/echo",
		"Echo Service",
		[]webpipes.InputParameter{
			webpipes.InputParameter{
				"input",
				"string",
				"input stream to echo",
				false,
				"",
			},
		},
		[]webpipes.OutputParameter{
			webpipes.OutputParameter{
				"output",
				"string",
				"echoed string",
			},
		},
	}

	pipe := webpipes.GoWebPipe{block, EchoPipe}
	http.Handle("/echo", pipe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

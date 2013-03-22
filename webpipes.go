package webpipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InputParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Optional    bool   `json:"optional?"`
	Default     string `json:"default"`
}

type OutputParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Block struct {
	Name        string            `json:"name"`
	Url         string            `json:"url"`
	Description string            `json:"description"`
	Inputs      []InputParameter  `json:"inputs"`
	Outputs     []OutputParameter `json:"outputs"`
}

type InputEnvelope struct {
	Inputs map[string]string `json:"inputs"`
}

type OutputEnvelope struct {
	Inputs map[string]string `json:"outputs"`
}

type BlockDefinition struct {
	Name string `json:"name"`
	url  string `json:"url"`
}

type BlockPath struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type Pipe struct {
	SourceBlock  int    `json:"source_block"`
	SourceOutput string `json:"source_output"`
	SourceValue  string `json:"source_value"`
	TargetBlock  int    `json:"target_block"`
	TargetInput  string `json:"target_block"`
}

type Pipeline struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Blocks      []BlockPath       `json:"blocks"`
	Pipes       []Pipe            `json:"pipes"`
	Inputs      []InputParameter  `json:"inputs"`
	Outputs     []OutputParameter `json:"outputs"`
}

type WebpipesError struct {
}

func (error *WebpipesError) Error() string {
	return "Webpipes Error"
}

type GoWebPipe struct {
	Block   Block
	Handler func(map[string]string) (map[string]string, error)
}

func (goWebPipe GoWebPipe) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		blockBytes, err := json.Marshal(&goWebPipe.Block)
		if err == nil {
			fmt.Fprintf(w, "%s", string(blockBytes))
		}
	} else if req.Method == "POST" {
		body, readErr := ioutil.ReadAll(req.Body)
		if readErr == nil {
			inputEnvelope := InputEnvelope{}
			unmarshalErr := json.Unmarshal(body, &inputEnvelope)
			if unmarshalErr == nil {
				// call the handler passed in
				outputs, webpipesError := goWebPipe.Handler(inputEnvelope.Inputs)
				if webpipesError == nil {
					outputEnvelope := OutputEnvelope{outputs}
					outputJson, marshalErr := json.Marshal(outputEnvelope)
					if marshalErr == nil {
						w.Write(outputJson)
					} else {
						w.WriteHeader(500)
						fmt.Fprintf(w, "%s", marshalErr)
					}
				} else {
					w.WriteHeader(500)
					fmt.Fprintf(w, "%s", webpipesError)
				}
			} else {
				w.WriteHeader(500)
				fmt.Fprintf(w, "%s", unmarshalErr)
			}
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", readErr)
		}
	}
}

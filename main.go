package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type presentation struct {
	Echo             bool   `json:"echo"`
	Reveal           string `json:"reveal"`
	Focus            bool   `json:"focus"`
	Panel            string `json:"panel"`
	ShowReuseMessage bool   `json:"showReuseMessage"`
	Clear            bool   `json:"clear"`
}

type task struct {
	Label        string       `json:"label"`
	Type         string       `json:"type"`
	Command      string       `json:"command"`
	Presentation presentation `json:"presentation"`
}

var defaultTask = task{
	Label:   "run",
	Type:    "shell",
	Command: "go build -o out && ./out",
	Presentation: presentation{
		Echo:             true,
		Reveal:           "always",
		Focus:            false,
		Panel:            "shared",
		ShowReuseMessage: true,
		Clear:            true,
	},
}

type document struct {
	Version string `json:"version"`
	Tasks   []task `json:"tasks"`
}

var defaultDocument = document{
	Version: "2.0.0",
	Tasks:   []task{defaultTask},
}

func main() {
	err := os.Mkdir(".vscode", 0700)
	if err != nil {
		log.Fatal("Could not create directory .vscode: ", err)
	}

	vs, err := os.Create(filepath.Join(".vscode", "tasks.json"))
	if err != nil {
		log.Fatal("Could not create tasks.json: ", err)
	}
	defer vs.Close()

	enc := json.NewEncoder(vs)
	enc.SetIndent("", "    ")
	enc.Encode(defaultDocument)

}

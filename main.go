package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
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

func makeDotVsCode(dir string) error {
	err := os.Mkdir(filepath.Join(dir, ".vscode"), 0700)
	if err != nil {
		return fmt.Errorf("Could not create directory .vscode: %w", err)
	}

	vs, err := os.Create(filepath.Join(dir, ".vscode", "tasks.json"))
	if err != nil {
		return fmt.Errorf("Could not create tasks.json: %w", err)
	}

	document := defaultDocument
	document.Tasks[0].Command = fmt.Sprintf("go build -o %v && ./%v", dir, dir)

	enc := json.NewEncoder(vs)
	enc.SetIndent("", "    ")
	enc.Encode(defaultDocument)
	return vs.Close()
}

func goInit(dir, module string) error {
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = dir
	return cmd.Run()
}

func gitInit(dir, module string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	dat, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(dat))
		return err
	}

	if strings.HasPrefix(module, "github.com") || strings.HasPrefix(module, "gitlab.com") {
		cmd = exec.Command("git", "remote", "add", "origin", module)
		cmd.Dir = dir
		dat, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(string(dat))
			return err
		}
	}
	return nil
}

func main() {

	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n %s [options] <module>\n", name, name)
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	module := flag.Arg(0)
	dir := path.Base(module)

	if err := os.Mkdir(dir, 0700); err != nil {
		log.Fatal(err)
	}

	if err := makeDotVsCode(dir); err != nil {
		log.Fatal(err)
	}

	if err := goInit(dir, module); err != nil {
		log.Fatal(err)
	}

	if err := gitInit(dir, module); err != nil {
		log.Fatal(err)
	}
}

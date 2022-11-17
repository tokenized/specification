package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tokenized/specification/internal/golang"
	"github.com/tokenized/specification/internal/json"
	"github.com/tokenized/specification/internal/markdown"
	"github.com/tokenized/specification/internal/platform/parser"
	"github.com/tokenized/specification/internal/protobuf"
	"github.com/tokenized/specification/internal/python"
	"github.com/tokenized/specification/internal/swagger"
	"github.com/tokenized/specification/internal/typescript"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	srcPath := filepath.FromSlash(dir + "/src")
	distPath := filepath.FromSlash(dir + "/dist")

	// --------------------------------------------------------------------
	// Prepare Schemas

	actions, err := parser.NewSchema(filepath.FromSlash(srcPath + "/actions/develop"))
	if err != nil {
		log.Fatal(err)
	}

	instruments, err := parser.NewSchema(filepath.FromSlash(srcPath + "/instruments/develop"))
	if err != nil {
		log.Fatal(err)
	}

	messages, err := parser.NewSchema(filepath.FromSlash(srcPath + "/messages/develop"))
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------------------------------------------------
	// Compile Languages

	protobuf.Compile(srcPath, distPath, *actions, *instruments, *messages)

	golang.Compile(srcPath, distPath, *actions, *instruments, *messages)

	json.Compile(srcPath, distPath, *actions, *instruments, *messages)

	markdown.Compile(srcPath, distPath, *actions, *instruments, *messages)

	python.Compile(srcPath, distPath, *actions, *instruments, *messages)

	typescript.Compile(srcPath, distPath, *actions, *instruments, *messages)

	swagger.Compile(srcPath, distPath, *actions, *instruments, *messages)
}

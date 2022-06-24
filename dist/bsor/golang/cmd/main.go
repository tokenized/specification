package main

import (
	"encoding/json"
	"fmt"
	"os"

	bsor "github.com/tokenized/specification/dist/bsor/golang"
)

func main() {
	path := os.Args[1]

	actionDefs, err := bsor.Generateactions()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate actions : %s", err))
	}

	instrumentDefs, err := bsor.Generateinstruments()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate instruments : %s", err))
	}

	messageDefs, err := bsor.Generatemessages()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate messages : %s", err))
	}

	actionsPath := path + "/actions.bsor"
	actionsFile, err := os.Create(actionsPath)
	if err != nil {
		panic(fmt.Sprintf("Create actions bsor file : %s", err))
	}
	actionsFile.Write([]byte(actionDefs.String() + "\n"))
	actionsFile.Close()

	instrumentsPath := path + "/instruments.bsor"
	instrumentsFile, err := os.Create(instrumentsPath)
	if err != nil {
		panic(fmt.Sprintf("Create instruments bsor file : %s", err))
	}
	instrumentsFile.Write([]byte(instrumentDefs.String() + "\n"))
	instrumentsFile.Close()

	messagesPath := path + "/messages.bsor"
	messagesFile, err := os.Create(messagesPath)
	if err != nil {
		panic(fmt.Sprintf("Create messages bsor file : %s", err))
	}
	messagesFile.Write([]byte(messageDefs.String() + "\n"))
	messagesFile.Close()

	actionsPath = path + "/actions.json"
	actionsFile, err = os.Create(actionsPath)
	if err != nil {
		panic(fmt.Sprintf("Create actions json file : %s", err))
	}
	actionsEncoder := json.NewEncoder(actionsFile)
	actionsEncoder.SetIndent("", "  ")
	if err := actionsEncoder.Encode(actionDefs); err != nil {
		panic(fmt.Sprintf("Encode actions json file : %s", err))
	}
	actionsFile.Close()

	instrumentsPath = path + "/instruments.json"
	instrumentsFile, err = os.Create(instrumentsPath)
	if err != nil {
		panic(fmt.Sprintf("Create instruments json file : %s", err))
	}
	instrumentsEncoder := json.NewEncoder(instrumentsFile)
	instrumentsEncoder.SetIndent("", "  ")
	if err := instrumentsEncoder.Encode(instrumentDefs); err != nil {
		panic(fmt.Sprintf("Encode instruments json file : %s", err))
	}
	instrumentsFile.Close()

	messagesPath = path + "/messages.json"
	messagesFile, err = os.Create(messagesPath)
	if err != nil {
		panic(fmt.Sprintf("Create messages json file : %s", err))
	}
	messagesEncoder := json.NewEncoder(messagesFile)
	messagesEncoder.SetIndent("", "  ")
	if err := messagesEncoder.Encode(messageDefs); err != nil {
		panic(fmt.Sprintf("Encode messages json file : %s", err))
	}
	messagesFile.Close()
}

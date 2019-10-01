# Go Tokenized Protocol

Usage

    package main

    import (
        "fmt"

        "github.com/tokenized/specification/dist/golang/actions"
        "github.com/tokenized/specification/dist/golang/protocol"
    )

    func main() {
        permissions := protocol.Permissions{
            protocol.Permission{
                Permitted:              true,
                AdministrationProposal: true,
                HolderProposal:         false,
                AdministrativeMatter:   false,
                VotingSystemsAllowed:   []bool{true, true},
                // Fields: nil, // Leave Fields blank to use as default for all fields.
            },
        }

        // Serialize permissions
        permBytes, err := permissions.Bytes()
        if err != nil {
            fmt.Printf("Failed to serialize permissions\n")
            return
        }

        // Generate a new contract offer action
        action := actions.ContractOffer{
            ContractName:        "Test",
            BodyOfAgreementType: 2,
            BodyOfAgreement:     []byte("<contract agreement>"),
            ContractType:        "Test Type",
            ContractPermissions: permBytes,
            // Specify any other fields necessary
            // ...
        }

        // Serialize action
        isTest := true // use "test.tokenized" OP_RETURN signature
        actionData, err := protocol.Serialize(&action, isTest)
        if err != nil {
            fmt.Printf("Failed to serialize payload\n")
            return
        }

        // Convert to hex
        fmt.Printf("Contract Offer Hex :\n%x\n\n", actionData)

        // Output as human readable string
        fmt.Printf("Contract Offer text : \n%s\n", action.String())
    }

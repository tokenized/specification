# Go Tokenized Protocol

Usage

    package main

    import (
        "fmt"

        "github.com/tokenized/specification/dist/golang/actions"
        "github.com/tokenized/specification/dist/golang/protocol"
    )

    func main() {
        voteSystemsAllowed := make([]bool, 2)
        voteSystemsAllowed[0] = true
        voteSystemsAllowed[1] = true

        permission := protocol.Permission{
            Permitted:              true,
            AdministrationProposal: true,
            HolderProposal:         false,
            VotingSystemsAllowed:   voteSystemsAllowed,
        }

        // Note: Permissions can be different for each field.
        permissions := make([]protocol.Permission, 0, 20)
        for i := 0; i < 20; i++ { // 20 fields in contract
            permissions = append(permissions, permission)
        }

        // Serialize auth flags
        authFlags, err := protocol.WriteAuthFlags(permissions)
        if err != nil {
            fmt.Printf("Failed to serialize auth flags\n")
            return
        }

        // Generate a new contract offer action
        action := actions.ContractOffer{
            ContractName:        "Test",
            BodyOfAgreementType: 2,
            BodyOfAgreement:     []byte("<contract agreement>"),
            ContractType:        "Test Type",
            ContractAuthFlags:   authFlags,
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

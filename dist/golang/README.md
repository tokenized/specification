# Go Tokenized Protocol

Usage

    import "github.com/tokenized/specification/dist/protocol"

    // Generate a new action
    action := protocol.ContractOffer{
        ContractName:     "Test",
        ContractFileType: 2,
        ContractFile:     []byte("sha25666666666666666666666666"),
        // ...
    }

    // Serialize action
    actionBytes, err := protocol.Serialize(&action)
    if err != nil {
        return errors.Wrap(err, "serializing payload")
    }

    // Convert to hex
    actionHex := fmt.Sprintf("%x", actionBytes)

    // Output as human readible string
    fmt.Println(action.String())

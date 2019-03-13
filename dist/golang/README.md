# Go Tokenized Protocol

Usage

    import "github.com/tokenized/specification/dist/protocol"

    // Create a new protocol action (parameters: Contract Name, Issuer Name)
    action := protocol.NewContractOffer([]byte("My Contract"), []byte("Acme Corporation"))

    // Serialize the action as byte array
    output, err := action.Serialize()
    if err != nil {
        return err
    }

    // Output bytes
    fmt.Println(output)

    // Output as human readible string
    fmt.Println(action.String())

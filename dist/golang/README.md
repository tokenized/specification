# Go Tokenized Protocol

Usage

    import "github.com/tokenized/specification/dist/protocol"

    // Create a new protocol action
    action := protocol.NewContractOffer()
    action.ContractName = Nvarchar32("My Contract")
    action.IssuerName = Nvarchar32("Acme Corporation")

    // Serialize the action as byte array
    output, err := action.Serialize()
    if err != nil {
        return err
    }

    // Output bytes
    fmt.Println(output)

    // Output as human readible string
    fmt.Println(action.String())

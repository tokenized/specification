# Tokenized Typescript

***Warning*** - This library is new and changing rapidly.

# Install

```bash
cd $GOPATH/src/github.com/tokenized/specification/dist/typescript
npm install
```

# Build

```bash
cd $GOPATH/src/github.com/tokenized/specification/dist/typescript
node node_modules/gulp/bin/gulp.js
```

# Generate

```bash
cd $GOPATH/src/github.com/tokenized/specification/
go run cmd/tokenized/main.go generate
```

# Unit Tests

```bash
cd $GOPATH/src/github.com/tokenized/specification/dist/typescript
npm run test
```

# Install

```bash
cd $GOPATH/src/github.com/tokenized/specification/dist/typescript
npm run test
```

# Usage

```typescript
import {
	ContractOffer, OpReturnMessage, Permission, Timestamp, Entity, PublicKeyHash
} from '@tokenized/tokenized';

let voteSystemsAllowed = [ true, true ];

let permission = new Permission();
permission.permitted = true;
permission.administration_proposal = true;
permission.holder_proposal = false;
permission.voting_systems_allowed = voteSystemsAllowed;

	// Note: Permissions can be different for each field.
	let permissions = [...Array(20)].map(_ => permission);

// Serialize auth flags
let authFlags = Permission.WriteAuthFlags(permissions);

// Generate a new contract offer action
let action = new ContractOffer();
action.contract_name = 'Test';
action.body_of_agreement_type = 2;
action.body_of_agreement = [...Buffer.from('<contract agreement>', 'ascii')];
action.contract_type = 'Test Type';
action.contract_auth_flags = authFlags;
action.supporting_docs = [];
action.contract_expiration = new Timestamp();
action.issuer = new Entity();
action.voting_systems = [];
action.oracles = [];
action.master_pkh = new PublicKeyHash();
// Specify any other fields necessary
// ...


// Serialize action
let isTest = true; // use "test.tokenized" OP_RETURN signature
let actionData = OpReturnMessage.Serialize(action, isTest);

// Convert to hex
console.log("Contract Offer Hex :\n", actionData);

// Output as human readable string
console.log("Contract Offer text :\n", JSON.stringify(action));

```

# License

Open BSV

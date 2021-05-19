# Proposal: Add identity management APIs

## Background

In order to make it easier for users to use the service without having to manually enter authentication 
information (e.g. ak, sk, hmac, etc.) and service endpoint information each time, 
we have added identity management-related functionality to the prototype, which requires support such as:

- Add identity
- Delete identity
- List identities

and so on. The identity information needs to be stored persistently and can be reused when adding tasks at the `server`.

## Propose

So, I propose to design the API to be called by the front-end, including:

- List Identities
- Create Identity
- Delete Identity
- Get Identity

Where `Identity` is defined as follows:

```go
type Identity struct {
    Name string // my_id_1
    Type string // qingstor, fs
    Credential struct {
        Protocol string // hmac, file
        Value string // ak:sk, /path/to/token
    }
    Endpoint struct {
        Protocol string // http, https
        Value string // qingstor.com:443
    }
}
```

Persistent storage can generate `identity_key` in the form: `id:{id_type}:{id_name}`, 
which ensures that there is only one `key` of the same name for each different type, 
and makes it easier to filter by `type` later.

The relevant API definition mentioned above is as follows:

```graphql
type Query {
    identities(type: IdentityType): [Identity!]
    identity(type: IdentityType!, name: String!): Identity!
}

type Mutation {
    createIdentity(input: CreateIdentity): Identity!
    deleteIdentity(input: DeleteIdentity): Identity!
}

type Identity {
    name: String!
    type: IdentityType!
    credential: Credential!
    endpoint: Endpoint!
}

input CredentialInput {
    protocol: String!
    value: String!
}

type Credential {
    protocol: String!
    value: String!
}

input EndpointInput {
    protocol: String!
    value: String!
}

type Endpoint {
    protocol: String!
    value: String!
}

enum IdentityType {
    Qingstor # Only Qingstor type identity is supported for now
}

input CreateIdentity {
    name: String!
    type: IdentityType!
    credential: CredentialInput!
    endpoint: EndpointInput!
}

input DeleteIdentity {
    name: String!
    type: IdentityType!
}
```

Later, when the `createTask` interface is called, the `credential` information in the `option` of the `task` can be reused by resolving the `identity`.

## Rationale

Prototype reference design <https://www.figma.com/file/tZBW1fMDLlcdFpaHJYih9B/Data-Migration-Prototype?node-id=1191%3A5>

## Compatibility

None

## Implementation

Most of the work would be done by the author of this proposal.

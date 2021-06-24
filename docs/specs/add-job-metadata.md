---
author: Prnyself <https://github.com/Prnyself>
status: draft
updated_at: 2021-06-24
---

# Proposal: Add job metadata

## Background

Job is the unit we do some action, such as `CopyMultipart`, `CopySingleFile`. When a job is done, we may need get result
from the action. For now, we do not have a mechanism to get the results of a job. The function of handler is 
`func(ctx context.Context, msg protobuf.Message) error`, the only return value is an `error`.

## Propose

So, I propose to save the result as job's metadata in DB. 

The key of job metadata is like: `jmeta:<job_id>`, and the content is like:

```go
type JobMetadata struct {
    JobID string // redundant jobID for convenience
    Data  map[string]interface{}
}
```

In this way, when job's action is done, conduct the `JobMetadata` and rpc call to save it into task's DB. So that it can
be used as needed.

## Rationale

### Add Set Job Metadata Api

The rpc api for service `Worker` as follows:

```protobuf
rpc SetJobMetadata (SetJobMetadataRequest) returns (SetJobMetadataReply) {
}
rpc GetJobMetadata (GetJobMetadataRequest) returns (GetJobMetadataReply) {
}
```

```protobuf
message SetJobMetadataRequest {
    string job_id = 1;
    bytes metadata = 2;
}

message SetJobMetadataReply {
}

message GetJobMetadataRequest {
    string job_id = 1;
}

message GetJobMetadataReply {
    bytes metadata = 1;
}
```

### Set Job Metadata in Handler

When job is run, rpc call `SetJobMetadata` to save metadata into DB, and call `GetJobMetadata` as needed.

For more, user **should** ensure job metadata available by `Await` when job is handled asynchronously.

## Compatibility

None

## Implementation

Most of the work would be done by the author of this proposal.
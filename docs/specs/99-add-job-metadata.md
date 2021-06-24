---
author: Prnyself <https://github.com/Prnyself>
status: draft
updated_at: 2021-06-24
---

# DMP-99: Add job metadata

## Background

Job is the unit we do some action, such as `CopyMultipart`, `CopySingleFile`. When a job is done, we may need get result
from the action. For now, we do not have a mechanism to get the results of a job. The function of handler is 
`func(ctx context.Context, msg protobuf.Message) error`, the only return value is an `error`.

## Propose

So, I propose to save the result as job's metadata in DB.

The key of job metadata is like: `jmeta:<job_id>`, and 
we can define strongly typed `JobMetadata` to ensure the content and struct of metadata. 

Take the `WriteMultipartJobMetadata` as example:

```protobuf
message WriteMultipartJobMetadata {
    string etag = 1;
}
```

In this way, when job's action is done, construct the `JobMetadata` and rpc call to save it into task's DB. So that it can
be used as needed. For example:

```go
result, _ := protobuf.Marshal(&models.WriteMultipartJobMetadata{
	Etag: part.ETag,
})
_, err = rn.grpcClient.SetJobMetadata(ctx, &models.SetJobMetadataRequest{
	JobId: rn.j.Id,
	Metadata: result,
})
```

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

When job is run, conduct rpc call `SetJobMetadata` to save metadata into DB, and call `GetJobMetadata` as needed.

For more, user **should** ensure job metadata available by `Await` when job is handled asynchronously.

### Remove Metadata

When job is finished, the metadata can be GC by directly delete key with given prefix `jmeta:<job_id>` in DB. 

## Compatibility

None

## Implementation

Most of the work would be done by the author of this proposal.
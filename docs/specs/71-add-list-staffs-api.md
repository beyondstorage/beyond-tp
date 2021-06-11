---
author: Prnyself <https://github.com/Prnyself>
status: draft
updated_at: 2021-05-19
---

# DMP-71: Add list staffs API

## Background

When creating a task, we should choose on which staff the task run. For now, we have already inserted staffID into DB,
so we should support an `listStaffs` API for frontend to list all available staffs.

## Propose

So, I propose to design the list staff API to be called by the front-end, whose definition is as follows:

```graphql
type Query {
    staffs: [Staff!]!
}

type Staff {
    id: String!
}
```

For now, we only record staff ID, and more information will be added in the future.

In this way, when `createTask` is called, the `staffID` could be filled by the result of `listStaffs`.

## Rationale

The `createTask` input is defined as follows:

```graphql
input CreateTask {
    name: String!
    type: TaskType!
    storages: [StorageInput!]!
    options: [PairInput!]!
    staffs: [StaffInput!]!
}

input StaffInput {
    id: String!
}
```

## Compatibility

None

## Implementation

Most of the work would be done by the author of this proposal.
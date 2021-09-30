- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-09-29
- RFC PR: [beyondstorage/beyond-tp#176](https://github.com/beyondstorage/beyond-tp/issues/176)
- Tracking Issue: [beyondstorage/beyond-tp#0](https://github.com/beyondstorage/beyond-tp/issues/0)

# BTP-176: Streamlined Architecture

## Background

BeyondTP is designed to be a neutral data migration service, providing private deployment, SaaS services, and other delivery forms.

So we have a complex architecture:

```
Manager ----> Agent --+--> Leader
                      |
                      |
                      +--> Worker --+--> Runner
                                    |
                                    +--> Runner
```

Manager is the manager node of BeyondTP, it will take the following jobs:

- Agent management, including register, monitor, and so on.
- Task management, including task create/start/delete/distribute and so on.
- GraphQL API and GUI.

Agent is the client node of BeyondTP, it has two roles: Leader and Worker. Every time Agent got a task, it will try to elect as a leader. Only one of them will succeed and others will before a worker.

Every worker will try to contact the leader to poll/create/finish jobs. Inside a worker, it will start multiple runners so that we can run jobs concurrently.

However, it proved to be over-designed. Current architecture made us hard to implement features like:

- Speed limit
- Retry failed job
- Speed monitor
- Task report
- ...

## Proposal

So I propose to use streamlined architecture: merge manager and leader.

In this design, our architecture will be quite simple:

```
Server --> Client --+->  Worker
                    |
                    +->  Worker
                    |
                    +->  Worker
```

- Server will responsible for all coordinated operations
- client will focus on executing tasks
- Every client will spawn multiple workers to executing jobs.

## Rationale

N/A

## Compatibility

Brand new design, we will refactor the whole architecture.

## Implementation

As described in the tracking issue.

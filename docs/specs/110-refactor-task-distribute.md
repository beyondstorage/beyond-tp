---
author: Prnyself <https://github.com/Prnyself>
status: draft
updated_at: 2021-07-05
---

# DMP-110: Refactor Task Distribute

## Background

Task is the basic unit submitted by user. We store the metadata of a task into it. The struct is:

```go
type Task struct {
    Id        string                
    Name      string                
    Type      TaskType              
    Status    TaskStatus            
    CreatedAt *timestamppb.Timestamp
    UpdatedAt *timestamppb.Timestamp
    Storages  []*Storage            
    Options   []*Pair               
    StaffIds  []string              
}
```

The staff will try to poll task when started. For now, we use for-loop to monitor task submitting, and distribute
the task into staffs which are selected by user.
When execution, the staffs will elect to be the `leader`, then split the task into jobs and distribute them into 
different runner to run.

While waiting the task, manager returns `PollStatus_Empty` by rpc and then sleep for another round. 
There are two problems to handle waiting in this way:

- interval of sleep cannot be set properly
  - too frequent if interval is small
  - delay the response if interval is too big  
- waste of server's resource 

We need an elegant way to monitor task submitting.

## Propose

So, I propose to delay the distribution of task from creating to running.

For one hand, we do not need to insert `staff-task relation` when create task. It is enough to know which staff will run
this task, before task's distribution.

For another hand, when a staff started, it can call `Poll` to subscribe task change with prefix `s_t:{staff-id}:` 
(replace by its own staff ID), and do not insert `staff-task relation` when create task, but insert relation with 
key `s_t:{staff-id}:{task-id}` when run task.

To ensure the data consistency, we can use `transaction` between modify task status and insert `staff-task relation`.

When a staff start to poll task, it will monitor the `staff-task relation` change, and the process will be hang up
until a real `staff-task relation` has been set up, which indicates the task is going to run by this staff.

## Rationale

### Integrate Subscribe into Register

Otherwise, we can integrate subscription into `Register`, so that only one rpc call is needed instead of rpc call separately.

## Compatibility

None

## Implementation

Most of the work would be done by the author of this proposal.
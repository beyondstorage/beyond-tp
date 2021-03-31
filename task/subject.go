package task

import "fmt"

// Tasks subject will be holding on Manager.
// Manager use this subject to publish new tasks.
func SubjectTasks() string {
	return "tasks"
}

// TaskReply subject will be holding on Manager.
func SubjectTaskReply(taskId string) string {
	return fmt.Sprintf("task.reply.%s", taskId)
}

// Task subject will be holding on Staff.
// All staff (both leader and worker) will use this subject to publish new jobs.
func SubjectTask(taskId string) string {
	return fmt.Sprintf("task.%s", taskId)
}

// subject is the ElectReply.Subject
func SubjectClockin(subject string) string {
	return fmt.Sprintf("clockin.%s", subject)
}

// subject is the ElectReply.Subject
func SubjectClockout(subject string) string {
	return fmt.Sprintf("clockout.%s", subject)
}

// subject is the ElectReply.Subject
func SubjectClockoutNotify(subject string) string {
	return fmt.Sprintf("checkout.notify.%s", subject)
}

// JobReply subject will be holding on Staff.
// Staff will use this subject to wait for replies on specific job.
func SubjectJobReply(jobId string) string {
	return fmt.Sprintf("job.reply.%s", jobId)
}

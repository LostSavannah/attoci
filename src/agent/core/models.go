package core

type JobStep struct {
	Name     string
	Commands []string
}

type Job struct {
	JobId string
	Steps []JobStep
}

type JobStepEvent struct {
	JobStepName  string
	EventName    string
	EventContent string
	EventMeta    map[string]any
}

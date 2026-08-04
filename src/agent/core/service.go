package core

import (
	"encoding/json"
)

func Example(number int) int {
	return number * 2
}

type AttociAgentIdentifier struct {
	AgentId   string
	AgentType string
}

type AttociService struct {
	HttpClient
	AttociAgentIdentifier
}

func (service AttociService) GetCurrentJob() (job Job, err error) {
	data, err := service.Send("POST", "/api/requests/job", service.AttociAgentIdentifier)
	if err == nil {
		err = json.Unmarshal(data, &job)
	}
	return
}

func (service AttociService) SendJobEvent(event_id string, event JobStepEvent) (err error) {
	url := "/api/jobs/" + event_id + "/events"
	_, err = service.Send("POST", url, event)
	return
}

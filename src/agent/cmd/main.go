package main

import (
	"fmt"
	"os"
	"time"

	"ybcdmi.org/attoci-agent/core"
)

func main() {
	const JobStartedEventName = "job.started"
	const JobFailedEventName = "job.failed"
	const JobFinishedEventName = "job.finished"
	const StepStartedEventName = "step.started"
	const StepStdoutEventName = "step.stdout"
	const StepStderrEventName = "step.stderr"
	const StepFailedEventName = "step.failed"
	const StepFinishedEventName = "step.finished"

	baseUrl := os.Getenv("ATTOCI_URL")
	apiKey := os.Getenv("ATTOCI_API_KEY")
	agentId := os.Getenv("ATTOCI_AGENT_ID")
	agentType := os.Getenv("ATTOCI_AGENT_TYPE")
	srv := core.AttociService{
		AttociAgentIdentifier: core.AttociAgentIdentifier{
			AgentId:   agentId,
			AgentType: agentType,
		},
		HttpClient: core.HttpClient{
			BaseUrl:        baseUrl,
			ApiKey:         apiKey,
			TimeoutSeconds: 10,
		},
	}

	fmt.Println("Agent :", agentId, ": started...")
	for {
		job, err := srv.GetCurrentJob()
		if err == nil {
			fmt.Println("Running job: ", job.JobId)
			sendEvent := func(event core.JobStepEvent) {
				srv.SendJobEvent(job.JobId, core.JobStepEvent{
					JobStepName:  event.JobStepName,
					EventName:    event.EventName,
					EventContent: event.EventContent,
					EventMeta: map[string]any{
						"Agent-Epoch": time.Now().Unix(),
					},
				})
			}
			sendEvent(core.JobStepEvent{
				EventName: JobStartedEventName,
			})
			failure := false
			for _, step := range job.Steps {
				fmt.Println("\tStep :", step.Name, ": running (...)")
				sendEvent(core.JobStepEvent{
					JobStepName: step.Name,
					EventName:   StepStartedEventName,
				})
				step.Execute(func(result core.CommandOutput) {
					eventName := StepStdoutEventName
					if result.IsError {
						eventName = StepStderrEventName
					} else if result.IsFailure {
						eventName = StepFailedEventName
					}
					sendEvent(core.JobStepEvent{
						JobStepName:  step.Name,
						EventName:    eventName,
						EventContent: result.Text,
					})
					failure = result.IsFailure
				})
				if failure {
					fmt.Println("\tStep :", step.Name, ": failed (!)")
					break
				}
				fmt.Println("\tStep :", step.Name, ": succeded (.)")
				sendEvent(core.JobStepEvent{
					JobStepName: step.Name,
					EventName:   StepFinishedEventName,
				})
			}
			if failure {
				sendEvent(core.JobStepEvent{
					EventName: JobFailedEventName,
				})
			} else {
				sendEvent(core.JobStepEvent{
					EventName: JobFinishedEventName,
				})
			}
		}
		time.Sleep(time.Second * time.Duration(4))
	}
}

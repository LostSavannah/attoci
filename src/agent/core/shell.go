package core

import (
	"bufio"
	"os/exec"
	"runtime"
)

type CommandOutput struct {
	IsError   bool
	IsFailure bool
	Text      string
}

func Error(text string) CommandOutput {
	return CommandOutput{
		IsError:   true,
		IsFailure: false,
		Text:      text,
	}
}

func Failure(text string) CommandOutput {
	return CommandOutput{
		IsError:   true,
		IsFailure: true,
		Text:      text,
	}
}

func Output(text string) CommandOutput {
	return CommandOutput{
		IsError:   false,
		IsFailure: false,
		Text:      text,
	}
}

func RunCommand(command string, callback func(CommandOutput)) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	case "linux":
		cmd = exec.Command("sh", "-c", command)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		callback(Error(err.Error()))
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		callback(Error(err.Error()))
		return
	}

	if err = cmd.Start(); err != nil {
		callback(Error(err.Error()))
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			callback(Output(scanner.Text()))
		}
		if err = scanner.Err(); err != nil {
			callback(Output(err.Error()))
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			callback(Error(scanner.Text()))
		}
		if err = scanner.Err(); err != nil {
			callback(Output(err.Error()))
		}
	}()

	if err := cmd.Wait(); err != nil {
		callback(Failure(err.Error()))
	}
}

func (step JobStep) Execute(callback func(CommandOutput)) {
	for _, command := range step.Commands {
		commandFailure := false
		RunCommand(command, func(output CommandOutput) {
			callback(output)
			commandFailure = output.IsFailure
		})
		if commandFailure {
			break
		}
	}
}

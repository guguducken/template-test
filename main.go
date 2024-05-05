package main

import (
	"fmt"

	"github.com/guguducken/template-test/pkg/github"
)

func main() {
	key := github.GetInput("access-key")
	github.SetOutput("test-key", key)

	github.Debug("this is debug message for action-go-template")
	github.Warning(nil, "this is a nil property warning")
	github.Warning(github.NewWarningProperty("warning-test", "main.go", 1, 2, 1, 2), "this is a property warning")

	github.Error(nil, "this is an nil property error message for action-go-template")
	github.Error(github.NewErrorProperty("error-test", "main.go", 1, 2, 1, 2), "this is a property error message for action-go-template")

	github.Notice(nil, "this is a notice message for action-go-template")
	github.Notice(github.NewNoticeProperty("notice-test", "main.go", 1, 2, 1, 2), "this is a notice message for action-go-template")

	github.SetEnv(github.Properties{github.NewEnvProperty("ENV_TEST", "env-test")})

	github.AddMask("shanghai")
	fmt.Printf("china shanghai guguducken\n")

	github.AddPath("/test/bin")

	github.StartGroup("group-test")
	for i := 0; i < 100; i++ {
		fmt.Printf("group log %d\n", i)
	}
	github.StopGroup()

	github.Stop("stop-token-token")
	github.Warning(nil, "this is a warning property warning after stop command")
	github.Start("stop-token-token")
	github.Warning(nil, "this is a warning property warning after start command")

	github.Summary("### This is last summary :rocket:")
}

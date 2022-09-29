package main

import (
	"github.com/timsutton/learn-go/cmd"
	"github.com/timsutton/learn-go/util"
)

func main() {
	util.SetupEnvironment()

	cmd.Execute()
}

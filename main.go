package main

import (
	"github.com/timsutton/speedwagon/cmd"
	"github.com/timsutton/speedwagon/util"
)

func main() {
	util.SetupEnvironment()

	cmd.Execute()
}

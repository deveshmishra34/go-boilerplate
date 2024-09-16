package main

import "github.com/deveshmishra34/groot/cmd"

var VERSION string = "2.2.1-default"

func main() {
	cmd.Version = VERSION
	cmd.Execute()
}

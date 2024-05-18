package main

import (
	start "go-reloaded/pkg/run"
)

func main() {
	start.Run("./sample.txt", "result.txt")
	start.Test()
}

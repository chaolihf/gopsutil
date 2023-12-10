package gopsutil

import (
	"fmt"
	process "github.com/chaolihf/gopsutil/process"
	"testing"
)

func Test_Process(t *testing.T) {
	p, _ := process.Processes()
	fmt.Println(len(p))
}

package gopsutil

import (
	"fmt"
	"testing"

	process "github.com/chaolihf/gopsutil/process"
)

func Test_Process(t *testing.T) {
	allProcess, err := process.Processes()
	if err != nil {
		t.Error(err.Error())
	} else {
		for _, item := range allProcess {
			nsPid, _ := item.GetContainerPid()
			fmt.Println(nsPid)

			username, _ := item.Username()
			fmt.Println(username)
			name, _ := item.Name()
			fmt.Println(name)
			command, _ := item.Cmdline()
			fmt.Println(command)
			memory, _ := item.MemoryInfo()
			fmt.Println(memory)
			numThread, _ := item.NumThreads()
			fmt.Println(numThread)
			numOpenFiles, _ := item.NumFDs()
			fmt.Println(numOpenFiles)
			createTime, _ := item.CreateTime()
			fmt.Println(createTime)
			parentId, _ := item.Ppid()
			fmt.Println(parentId)
			cpu, _ := item.CPUPercent()
			fmt.Println(cpu)
			exec, _ := item.Exe()
			fmt.Println(exec)
			ioCounters, _ := item.IOCounters()
			fmt.Println(ioCounters)
			if memory != nil {
				fmt.Println(int64(memory.RSS))
				fmt.Println(int64(memory.VMS))
			}
			fmt.Println(item.Pid)
			if ioCounters != nil {
				fmt.Println(int64(ioCounters.ReadBytes))
				fmt.Println(int64(ioCounters.WriteBytes))
				fmt.Println(int64(ioCounters.ReadCount))
				fmt.Println(int64(ioCounters.WriteCount))
			}
		}
	}
}

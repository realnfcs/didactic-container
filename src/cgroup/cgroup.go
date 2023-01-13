package cgroup

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/realnfcs/didactic-container/src/utils"
)

func CreateCgroup() {
	cgroups := "/sys/fs/cgroup/"

	err := os.Mkdir(filepath.Join(cgroups, "pids"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	pids := filepath.Join(cgroups, "pids")
	err = os.Mkdir(filepath.Join(pids, "didactic-container"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	// Limit of 20 Process
	utils.Must(os.WriteFile(filepath.Join(pids, "pids.max"), []byte("20"), 0700))

	// Remove the new cgroup in place after the container exists
	// must(os.WriteFile(filepath.Join(pids, "notify_on_release"), []byte("1"), 0700))

	// Get the current PID e write in a file named cgroup.procs, that's mean now the process is a member on of this control group with limit
	utils.Must(os.WriteFile(filepath.Join(pids, "didactic-container/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

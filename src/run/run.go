package run

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/realnfcs/didactic-container/src/cgroup"
	"github.com/realnfcs/didactic-container/src/utils"
)

func Run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create a namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Unix Timesharing System 		 // Process ID 			// Namespaces
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
		Credential:   &syscall.Credential{Uid: 0, Gid: 0},
		/*
			UidMappings: []syscall.SysProcIDMap{
				{ContainerID: 0, HostID: os.Getuid(), Size: 1},
			},
			GidMappings: []syscall.SysProcIDMap{
				{ContainerID: 0, HostID: os.Getgid(), Size: 1},
			},
		*/
	}

	utils.Must(cmd.Run())
}

func Child() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cgroup.CreateCgroup()

	utils.Must(syscall.Sethostname([]byte("container")))
	utils.Must(syscall.Chroot("./fs"))
	utils.Must(syscall.Chdir("/"))
	utils.Must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	syscall.Unmount("/proc", 0)
}

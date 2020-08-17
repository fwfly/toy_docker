package main

import (
    "log"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    cmd := exec.Command("bash")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
        syscall.CLONE_NEWNS |
        syscall.CLONE_NEWUSER,

// Default gid and uid are : uid=65534(nobody) gid=65534(nogroup) groups=65534(nogroup)
// SysProcAttr.Credential isn't supported in Ubuntu 19.10.
// Use Mappings instead

/*        UidMappings: []syscall.SysProcIDMap{
            {
                ContainerID: 1234,
                HostID:      0,
                Size:        1,
            },
        },
        GidMappings: []syscall.SysProcIDMap{
            {
                ContainerID: 1234,
                HostID:      0,
                Size:        1,
            },
        },
*/
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
    os.Exit(-1)
}

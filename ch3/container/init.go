package container

import (
    log "github.com/Sirupsen/logrus"
    "syscall"
    "os"
)

func RunContainerInitProcess(command string, args []string) error {
    log.Info("container Init Command %s", command)
    defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
    syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
    argv := []string{command}
    if err := syscall.Exec(command, argv, os.Environ()); err != nil {
        log.Errorf(err.Error())
    }
    return nil
}

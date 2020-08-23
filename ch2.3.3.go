/*
Ch 2.2.3

Verification:
1. You should see below results
$ make 2.2.3
sudo go run ch2.3.3.go
Cmd docker
3483Current pid 1
stress: info: [7] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd

2. Check if testmemorylimit is created under /sys/fs/cgroup/memory
$ ls /sys/fs/cgroup/memory | grep testmemorylimit
testmemorylimit

3. Use top to check if the memory has been limited.
*/

package main

import (
    "path"
    "os"
    "fmt"
    "io/ioutil"
    "os/exec"
    "syscall"
    "strconv"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
    if os.Args[0] == "/proc/self/exe"{
        fmt.Printf("Current pid %d",syscall.Getpid())
        fmt.Println()
        cmd := exec.Command("sh", "-c", `stress --vm-bytes 100m --vm-keep -m 1`)
        cmd.SysProcAttr = &syscall.SysProcAttr{
        }
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
            fmt.Println(err)
            os.Exit(-1)
        }
    }

    fmt.Println("Cmd docker")
    cmd := exec.Command("/proc/self/exe")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    } else {
        fmt.Printf("%v", cmd.Process.Pid)
        os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("50m"), 0644)
    }
    cmd.Process.Wait()

}

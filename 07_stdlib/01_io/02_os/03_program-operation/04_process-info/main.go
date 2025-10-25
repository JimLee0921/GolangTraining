package main

import (
	"fmt"
	"os"
)

func main() {
	// 平台不同，结果不同
	fmt.Println("Process ID (PID):", os.Getpid())
	fmt.Println("Parent Process ID (PPID):", os.Getppid())
	fmt.Println("Real User ID (UID):", os.Getuid())
	fmt.Println("Effective User ID (EUID):", os.Geteuid())
	fmt.Println("Real Group ID (GID):", os.Getgid())
	fmt.Println("Effective Group ID (EGID):", os.Getegid())

	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println("Error getting groups:", err)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	fmt.Println("Supplementary Group IDs:", groups)
	fmt.Println("Hostname:", hostname)
}

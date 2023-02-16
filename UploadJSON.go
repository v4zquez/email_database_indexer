package main

import (
    "fmt"
    "os/exec"
)

func execute(cmd string, url string, arg1 string, arg2 string, userPassword string, arg3 string, file string) {

    out, err := exec.Command(cmd, url, arg1, arg2, userPassword, arg3, file).Output()

    if err != nil {
        fmt.Printf("%s", err)
    }

    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)
}

func main() {
	execute("curl", "http://localhost:4080/api/_bulk", "-i", "-u", "admin:Complexpass#123", "--data-binary", "@olympics.ndjson")
}

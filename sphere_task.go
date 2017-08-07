package main

import (
    s "strings"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
 args := os.Args[1:]
 total_res := 0
 for i := 0; i < len (args); i++ {
    dat, err := ioutil.ReadFile (args[i])
    check (err)
    res := s.Count (string(dat), "\n")
    if len (args) > 1 {
        fmt.Print (fmt.Sprintf ("%7[1]d %[2]s\n", res, args[i]))
    } else {
        fmt.Print (fmt.Sprintf ("%[1]d %[2]s\n", res, args[i]))
    }
    total_res += res
 }

 if len (args) > 1 {
    fmt.Print (fmt.Sprintf ("%7[1]d total\n", total_res))
 }
}

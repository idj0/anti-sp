package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CheckCredentials(user User, path string, fpath string) bool {
    // if file doesn't exist then create one
    if _, err := os.Stat(fpath); os.IsNotExist(err) {
        fmt.Println("credential.json doesn't exist! creating file...")

        os.MkdirAll(path, os.ModePerm)

        fmt.Printf("Username: ")
        fmt.Scan(&user.Username)
        fmt.Printf("Password: ")
        fmt.Scan(&user.Password)
    // if file exist then check credential
    } else {
        file, _ := ioutil.ReadFile(fpath)
        if err := json.Unmarshal([]byte(file), &user); err != nil {
            log.Fatal(err)
        }

        if user.Username == "" ||  user.Password == "" {
            fmt.Printf("Username: ")
            fmt.Scan(&user.Username)
            fmt.Printf("Password: ")
            fmt.Scan(&user.Password)
        }
    }

    file, _ := json.MarshalIndent(user, "", " ")
    ioutil.WriteFile(fpath, file, 0644)

    return true
}


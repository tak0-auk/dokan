package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if lengh := len(os.Args); lengh < 2 {
		fmt.Println("doken is clone git repository and build Dockerfile. dokan need a  repository url")
		return
	} else if lengh > 2 {
		fmt.Println("only 1 argument")
		return
	}

	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", os.ModeDir)
		fmt.Println("make tmp dir")
		os.Chmod("tmp", 0777)
	}

	err := os.Chdir("tmp")
	if err != nil {
		fmt.Println(err)
	}

	url := os.Args[1]
	cmd := exec.Command("git", "clone", url)
	err = cmd.Run()
	os.Chdir("..")
	err = os.RemoveAll("tmp")
	fmt.Println(err)
}

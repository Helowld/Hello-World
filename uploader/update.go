package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	filePath := "./README.md"
	re, err := regexp.Compile(`List\[[0-9]{1,2}\]`)
	read, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
	args := os.Args[1:]
	fmt.Println(args)
	if n := len(args); n == 0 || args[0] == "" {
		panic("empty name")
	}

	content := string(read)
	loc := re.FindStringIndex(content)
	// fmt.Println(loc)
	listOfLanguages, err := strconv.Atoi(content[loc[0]+5 : loc[1]-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(listOfLanguages)
	newContents := strings.Replace(content, strconv.Itoa(listOfLanguages), strconv.Itoa(listOfLanguages+1), 1)
	fmt.Println(newContents[:700])

	// err = ioutil.WriteFile(filePath, []byte(newContents), 0)
	// if err != nil {
	// 	panic(err)
	// }

	// printContents()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func printContents() {
	data, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println("CONTENTS:", string(data))
}
func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("git", "status").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
		panic("not found the command")
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

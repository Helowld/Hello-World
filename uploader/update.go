package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := "./README.md"
	read, err := ioutil.ReadFile(filePath)
	isError(err)
	content := string(read)

	args := os.Args[1:]
	fmt.Println(args)
	if n := len(args); n == 0 || args[0] == "" {
		panic("Please specify the language name")
	} else if n > 1 {
		panic("expected only one argument")
	}

	do(&content, args[0])
	fmt.Println(content)
	fmt.Println("Press y to update the file")
	var write string
	fmt.Scanln(&write)
	writeToFile(write, &content)

	// printContents()
}

func writeToFile(write string, newContents *string) {
	filePath := "./README.md"
	yes := "yY"
	if strings.Contains(write, yes) {
		err := ioutil.WriteFile(filePath, []byte(*newContents), 0)
		isError(err)
	}

}

func do(content *string, pName string) {
	s := *content
	if fileName := execute(); fileName != "" && strings.HasPrefix(fileName, "hello_world.") {
		path := "- [" + pName + "](https://github.com/rustiever/Hello-World/blob/main/" + strings.Trim(fileName, "\n") + ")"
		i := strings.Index(s, "<details>")
		j := strings.Index(s, "</details>")
		n := strings.Index(s, "List[")
		num := s[n+5 : n+7]
		k := s[i+10 : j-1]
		l := strings.Split(k, "\n")
		l = append(l, path)
		sort.Strings(l)
		var pos int
		flag := false
		for m, n := range l {
			if n != "" {
				l[m] += "\n"
				if flag == false {
					l[m] = "\n" + l[m]
					pos = m
				}
				flag = true
			}
		}
		l = l[pos:]
		m := strings.Join(l, "\n")
		ni, err := strconv.Atoi(num)
		isError(err)
		ni += 1
		*content = strings.Replace(s, num, strconv.Itoa(ni), 1)
		*content = strings.Replace(*content, k, m, 1)
		// println(strings.Replace(s, k, m, 1))
	} else {
		println(fileName)
	}
}

func execute() string {
	out, err := exec.Command("git", "ls-files", "--others", "--exclude-standard").Output()

	if err != nil {
		fmt.Printf("%s", err)
		panic("not found the command")
	}
	output := string(out)
	return output
}

func printContents() {
	data, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println("CONTENTS:", string(data))
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return (err != nil)
}

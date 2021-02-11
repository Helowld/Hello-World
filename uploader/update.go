package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/thatisuday/commando"
)

const (
	filePath string = "../README.md"
)

func main() {
	read, err := ioutil.ReadFile(filePath)
	isError(err)
	content := string(read)

	commando.SetExecutableName("updater").SetVersion("0.0.1").SetDescription("Use this tool to update the Readme file after adding new language")
	commando.Register(nil).AddArgument("name", "language name", "").
		SetAction(func(action map[string]commando.ArgValue, flag map[string]commando.FlagValue) {
			println(action["name"].Value)
			do(&content, action["name"].Value)
			fmt.Println(content)
			if input("Update the file?(y/N)") {
				writeToFile(&content)
				isError(err)
				fmt.Println("File Updated Successfully")
			} else {
				fmt.Println("File Update is cancelled")
			}
		})
	commando.Parse(nil)
}

func input(stmt string) bool {
	fmt.Println(stmt)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	isError(err)
	return strings.Contains("yY", string(char))
}
func writeToFile(newContents *string) {
	err := ioutil.WriteFile(filePath, []byte(*newContents), 0)
	isError(err)
}

func do(content *string, pName string) {
	s := *content
	fileName, err := executeGit()
	isError(err)
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
	counter := 0
	for m, n := range l {
		if n != "" {
			counter++
			l[m] += "\n"
			if !flag {
				l[m] = "\n" + l[m]
				pos = m
			}
			flag = true
		}
	}
	l = l[pos:]
	m := strings.Join(l, "\n")
	*content = strings.Replace(s, num, strconv.Itoa(counter), 1)
	*content = strings.Replace(*content, k, m, 1)
}

func executeGit() (string, error) {
	out, err := exec.Command("git", "ls-files", "--others", "--exclude-standard").Output()
	isError(err)
	output := strings.Split(string(out), "\n")
	if len(output) == 2 {
		if strings.HasPrefix(output[0], "hello_world.") {
			return output[0], nil
		} else {

			return "", errors.New(`fileName must have prefix "hello_world."`)

		}
	}
	return "", errors.New("multiple new files has been found ")
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
		log.Fatalln(err.Error())
	}
	return (err != nil)
}

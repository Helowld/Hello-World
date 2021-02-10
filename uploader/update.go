/* package main */

/* import ( */
/* "bufio" */
/* "fmt" */
/* "io/ioutil" */
/* "os" */
/* "os/exec" */
/* "sort" */
/* "strconv" */
/* "strings" */

/* "github.com/thatisuday/commando" */
/* ) */

/* const ( */
/* filePath string = "./README.md" */
/* ) */

/* func main() { */

/* commando.SetExecutableName("updater").SetVersion("0.0.1").SetDescription("Use this tool to update the Readme file after adding new language") */
/* commando.Register(nil).AddArgument("name", "language name", "").SetAction(func(action map[string]commando.ArgValue, flag map[string]commando.FlagValue) { */
/* read, err := ioutil.ReadFile(filePath) */
/* isError(err) */
/* content := string(read) */

/* fileName := execute() */

/* fmt.Printf("proceed with this file => %s ?(y/n)", fileName) */
/* r := input() */
/* fmt.Printf("%v", r) */

/* do(&content, action["name"].Value, fileName) */
/* // fmt.Print("Press y to update the file") */
/* // var write string */
/* // fmt.Scanln(&write) */
/* writeToFile(string(r), &content) */
/* }) */
/* // printContents() */
/* commando.Parse(nil) */
/* } */

/* func input() rune { */
/* reader := bufio.NewReader(os.Stdin) */
/* char, _, err := reader.ReadRune() */
/* isError(err) */
/* return char */
/* } */

/* func do(content *string, pName string, fileName string) { */
/* s := *content */
/* if strings.HasPrefix(fileName, "hello_world.") { */
/* path := "- [" + pName + "](https://github.com/rustiever/Hello-World/blob/main/" + strings.Trim(fileName, "\n") + ")" */
/* i := strings.Index(s, "<details>") */
/* j := strings.Index(s, "</details>") */
/* n := strings.Index(s, "List[") */
/* num := s[n+5 : n+7] */
/* k := s[i+10 : j-1] */
/* l := strings.Split(k, "\n") */
/* l = append(l, path) */
/* sort.Strings(l) */
/* var pos int */
/* flag := false */
/* for m, n := range l { */
/* if n != "" { */
/* l[m] += "\n" */
/* if !flag { */
/* l[m] = "\n" + l[m] */
/* pos = m */
/* } */
/* flag = true */
/* } */
/* } */
/* l = l[pos:] */
/* m := strings.Join(l, "\n") */
/* ni, err := strconv.Atoi(num) */
/* isError(err) */
/* ni += 1 */
/* *content = strings.Replace(s, num, strconv.Itoa(ni), 1) */
/* *content = strings.Replace(*content, k, m, 1) */
/* // println(strings.Replace(s, k, m, 1)) */
/* } else { */
/* println(fileName) */
/* println(pName) */
/* } */
/* } */

/* func execute() string { */
/* out, err := exec.Command("git", "ls-files", "--others", "--exclude-standard").Output() */

/* if err != nil { */
/* fmt.Printf("%s", err) */
/* panic("not found the command") */
/* } */
/* output := string(out) */
/* strings.Split(output, "\n") */
/* println(output) */
/* return output */
/* } */

/* func printContents() { */
/* data, err := ioutil.ReadFile("./README.md") */
/* if err != nil { */
/* panic(err) */
/* } */
/* fmt.Println("CONTENTS:", string(data)) */
/* } */

/* func writeToFile(write string, newContents *string) { */
/* yes := "yY" */
/* if strings.Contains(write, yes) { */
/* err := ioutil.WriteFile(filePath, []byte(*newContents), 0) */
/* isError(err) */
/* } */

/* } */

/* func isError(err error) bool { */
/* if err != nil { */
/* fmt.Println(err.Error()) */
/* panic(err) */
/* } */
/* return (err != nil) */
/* } */

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
	println(strings.Contains("yY", write))
	writeToFile(write, &content, filePath)
	isError(err)

	// printContents()
}

func writeToFile(write string, newContents *string, path string) {
	yes := "yY"
	if strings.Contains(yes, write) {
		err := ioutil.WriteFile(path, []byte(*newContents), 0)
		isError(err)
	} else {
		println("why man")
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

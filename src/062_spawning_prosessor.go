package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello grep")

	grepIn ,_ := grepCmd.StdinPipe()
	grepOut,_ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\n good bye grep"))
	grepIn.Close()
	//grepBytes , _ := ioutil.ReadAll(grepOut)
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c","ls -a -l -h")
	lsOut, _ := lsCmd.Output()
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

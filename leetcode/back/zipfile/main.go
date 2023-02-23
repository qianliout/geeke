package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// if err := os.Chdir("/Users/liuqianli/Download"); err != nil {
	// 	fmt.Println("chdir err is ", err.Error())
	// }

	// exec.Command("cd", "/Users/liuqianli/Download")

	osCmd := exec.Command("zip", "-u", "/Users/liuqianli/Download/111.zip", "/Users/liuqianli/Download/hello.go")
	err := osCmd.Run()
	if err != nil {
		fmt.Println("err is ", err.Error(), osCmd.Stderr)
	}
}

// 调用系统命令去做增量压缩
func Zip(pre, new string) error {
	// -u 是增量压缩，-j是指不解析文件的绝对路径
	command := exec.Command("zip", "-uj", pre, new)
	err := command.Run()
	return err
}

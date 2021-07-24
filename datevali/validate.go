package main

import (
	"fmt"
	"os"
	"os/exec"
)

type DateValier struct {
	OperatorPath string
	Result       []byte
}

func BeginValidate() ([]byte, error) {
	ver := NewDefaultValier()
	ver.vali()
	return ver.Result, nil
}

func (ver *DateValier) vali() {
	pe, _ := PathExists(ver.OperatorPath)
	if pe {
		exec.Command("rm", "-rf", ver.OperatorPath).CombinedOutput()
		osExecClone(ver.OperatorPath, "git@git.hortorgames.com:lihao/cultivation_makedata.git")
	} else {
		osExecClone(ver.OperatorPath, "git@git.hortorgames.com:lihao/cultivation_makedata.git")
	}
	cmd := exec.Command("hortor-cli", "config", "./config-go.js")
	cmd.Dir = ver.OperatorPath
	var err error
	if ver.Result, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func NewDefaultValier() *DateValier {
	return &DateValier{OperatorPath: "/Users/hcm-b0192/validate/"}
}

//建筑 建筑  建筑

func main0() {

	var totalStr string
	var result []byte
	var err error

	//建筑筑
	exec.Command("rm", "-rf", "/data2/validata/dataSpace").CombinedOutput()
	osExecClone("/data2/validata/dataSpace", "git@git.hortorgames.com:lihao/cultivation_makedata.git")
	cmd := exec.Command("hortor-cli", "config", "/data2/validata/dataSpace/config-go.js")
	cmd.Dir = "/data2/validata/dataSpace/"
	if result, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	totalStr += string(result)

	fmt.Println(totalStr)
}

func osExecClone(workspace, url string) error {
	cmd := exec.Command("git", "clone", url, workspace)
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	return nil
}

func osExecPull(workspace string) error {
	cmd := exec.Command("git", "pull", workspace)
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
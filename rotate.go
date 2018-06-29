package log

import (
	"fmt"
	"os"
	"syscall"
)

/*
func RotateFile(name string, max int) {
	var pre string = fmt.Sprintf("%s.%d", name, max-1)
	var cur string
	for i := max - 2; i >= 0; i-- {
		if i == 0 {
			cur = name
		} else {
			cur = fmt.Sprintf("%s.%d", name, i)
		}
		_, err := os.Stat(cur)
		if err == os.ErrNotExist {
			continue
		}
		syscall.Rename(cur, pre)
		pre = cur
	}
}
*/

func RotateFiles(name string, max int) {

	if max <= 1 {
		panic(fmt.Sprintf("RoateFile's parameter 'max' must greater than 1: %d", max))
	}

	rename_list := []string{}
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		return
	}
	rename_list = append(rename_list, name)

	for i := 1; i < max; i++ {
		cur := fmt.Sprintf("%s.%d", name, i)
		rename_list = append(rename_list, cur)
		_, err := os.Stat(cur)
		if err != nil && os.IsNotExist(err) {
			break
		}
	}
	//fmt.Printf("%+v\n", rename_list)

	for i := len(rename_list) - 2; i >= 0; i-- {
		syscall.Rename(rename_list[i], rename_list[i+1])
	}
}

package common

import "fmt"

func Recover() {
	if err := recover(); err != nil {
		fmt.Println("Recover: ", err)
	}
}

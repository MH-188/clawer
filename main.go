package main

import (
	"clawer/xhs"
	"fmt"
)

func main() {
	xhsClient := xhs.NewXhsClient()
	_, err := xhsClient.GetNotePageInfo("")
	//err := xhsClient.GetPersonalPageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
}

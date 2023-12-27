package main

import (
	"fmt"
	"github.com/MH-188/clawer/xhs"
)

func main() {
	xhsClient := xhs.NewXhsClient()
	_, err := xhsClient.GetNotePageInfo("http://xhslink.com/80a39x")
	//err := xhsClient.GetPersonalPageInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
}

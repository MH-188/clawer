package xhs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMarshalTest(t *testing.T) {
	list1 := `{"fileId":"486a63bf-3c46-0785-9853-1ffc08e3836e","height":1920,"width":1272,"url":"https:\u002F\u002Fsns-img-hw.xhscdn.net\u002F486a63bf-3c46-0785-9853-1ffc08e3836e","traceId":"1040g00830ofkelo42e3g5oaf902gkmqaq63bli8"}`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(list1), &m)
	fmt.Println(err, m)
}

func TestGetNotePageInfo(t *testing.T) {
	client := NewXhsClient()
	_, err := client.GetNotePageInfo("http://xhslink.com/7H5xLx")
	if err != nil {
		fmt.Println(err)
		return
	}
}

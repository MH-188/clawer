package xhs

import (
	"clawer/tools"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type NoteDetail struct {
	Note NoteInner `json:"note"`
}

type NoteOuter struct {
	FirstNoteId   string                `json:"firstNoteId"`
	NoteDetailMap map[string]NoteDetail `json:"noteDetailMap"`
}

type NoteInner struct {
	ImageList []Image `json:"imageList"`
}

type BackgroundPictureResponse struct {
	Note NoteOuter `json:"note"`
}

type Image struct {
	FileId         string      `json:"fileId"`
	Height         int         `json:"height"`
	Width          int         `json:"width"`
	Url            string      `json:"url"`
	TraceId        string      `json:"traceId"`
	InfoList       []InfoLists `json:"infoList"`
	NoWatermarkUrl string
}

type InfoLists struct {
	Url string `json:"url"`
}

// 图片URL示例：    https://sns-img-hw.xhscdn.net/851b040f-d291-e19b-ebde-7525e707111a?imageView2/2/h/1920/format/webp|imageMogr2/strip
// 响应回来的URL：  https://sns-img-hw.xhscdn.net/037c63c1-ee5e-06c4-a38c-6ea2df0e071e

// GetBackgroundPicture 获取文章背景照片
func GetBackgroundPicture(respStr string) ([]string, error) {
	result := make([]string, 0, 10)
	index := strings.Index(respStr, "__INITIAL_STATE__")
	if index < 0 {
		err := fmt.Errorf("没找到__INITIAL_STATE__结构")
		return result, err
	}

	indexOffset := strings.Index(respStr[index:], "</script>")
	if indexOffset < 0 {
		err := fmt.Errorf("没找到__INITIAL_STATE__结构")
		return result, err
	}

	strNew := tools.RmUndefinedOfResponse(respStr[index+18 : index+indexOffset])
	var backGroundResp BackgroundPictureResponse
	err := json.Unmarshal([]byte(strNew), &backGroundResp)
	if err != nil {
		return result, err
	}

	for _, value := range backGroundResp.Note.NoteDetailMap {
		for _, image := range value.Note.ImageList {
			url := image.InfoList[0].Url
			traceId := ""
			strs := strings.Split(url, "/")
			if len(strs) < 1 {
				err = errors.New("图片链接异常")
				return result, err
			}
			strArray := strings.Split(strs[len(strs)-1], "!")
			if len(strArray) < 1 {
				err = errors.New("找不到traceId")
				return result, err
			}
			traceId = strArray[0]
			result = append(result, fmt.Sprintf("https://sns-img-qc.xhscdn.com/%s?imageView2/format/png", traceId))
		}
	}
	return result, nil
}

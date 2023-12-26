package xhs

import (
	"clawer/network"
	"fmt"
	"net/http"
)

type XClient struct {
	HttpClient *http.Client
	XhsHeader  map[string]string
}

func NewXhsClient() *XClient {
	return &XClient{
		HttpClient: http.DefaultClient,
		XhsHeader: map[string]string{
			"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		},
	}
}

// GetNotePageInfo 笔记主页
func (xc *XClient) GetNotePageInfo(url string) ([]string, error) {
	result := make([]string, 0)
	respBytes, err := network.HttpDoRequest(xc.HttpClient, "GET", url, xc.XhsHeader, nil)
	if err != nil {
		return result, err
	}

	// 响应的html页面转换成字符串
	htmlStr := string(respBytes)

	result, err = GetBackgroundPicture(htmlStr)
	if err != nil {
		return result, err
	}

	fmt.Println(result)
	return result, nil
}

func (xc *XClient) GetPersonalPageInfo() error {
	personalPageUrl := "https://www.xiaohongshu.com/user/profile/62f4761c000000001f004c86?xhsshare=CopyLink&appuid=5d477c6d0000000012025c94&apptime=1693878714"
	respBytes, err := network.HttpDoRequest(xc.HttpClient, "GET", personalPageUrl, xc.XhsHeader, nil)
	if err != nil {
		return err
	}

	// 响应的html页面转换成字符串
	htmlStr := string(respBytes)
	err = GetPersonalPageInfo(htmlStr)
	if err != nil {
		return err
	}

	return nil
}

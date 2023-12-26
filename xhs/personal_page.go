package xhs

import "os"

// https://www.xiaohongshu.com/user/profile/62f4761c000000001f004c86?xhsshare=CopyLink&appuid=5d477c6d0000000012025c94&apptime=1693878714
// 个人主页

// GetPersonalPageInfo  获取个人主页信息
func GetPersonalPageInfo(respStr string) error {
	file, err := os.OpenFile("./data/personal_page.html", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(respStr)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func fech(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Cookie", "Hm_lvt_866c9be12d4a814454792b1fd0fed295=1672926691; SL_G_WPT_TO=zh-CN; _ga=GA1.2.393270645.1672926691; _gid=GA1.2.801201576.1672926691; _gat_gtag_UA_476124_1=1; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8GXQNXLgcs5PrnWvMs4xAGNPLC_K5LAP40-VvnXqzWlb4-Hbys4-I2Tgjadw7MqjsoHuLgCjbOq5jBWLhHS4biDYYFaXugkuE62wOA3QT7fZ9nw_5ZUiLcmqzOR3lQkobI8IpDJRypbkKgpCB5P4xfo; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1672926715")
	resp, err := client.Do(req)
	if err != err {
		fmt.Println("Http get err:", err)
		return ""
	}

	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}
func main() {
	url := "https://zzk.cnblogs.com/s?w=golang"
	s := fech(url)
	fmt.Printf("s,:%v\n", s)
}

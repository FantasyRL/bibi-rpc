package main

import (
	"bibi/cmd/api/biz/model/api"
	"bibi/pkg/errno"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func register() string {
	t := time.Now()
	username := strconv.FormatInt(t.UTC().Unix(), 10)
	url := strings.Join([]string{"http://127.0.0.1:10001/bibi/user/register?username=", username, "&email=", username, "@qq.com&password=", "114514"}, "")
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "127.0.0.1:10001")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	return username
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
func login(username string) string {
	url := strings.Join([]string{"http://127.0.0.1:10001/bibi/user/login?username=", username, "&password=", "114514"}, "")
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "127.0.0.1:10001")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	resp := new(api.LoginResponse)
	sonic.Unmarshal(body, resp)
	return *(resp.AccessToken)
}

func putAvatar(token string, path string) {
	url := "http://127.0.0.1:10001/bibi/user/avatar/upload"
	method := "PUT"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(path)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("avatar_file", filepath.Base(path))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("access-token", token)
	req.Header.Add("refresh-token", "")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "127.0.0.1:10001")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------135599731246320599346196")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := new(api.AvatarResponse)
	sonic.Unmarshal(body, resp)
	if resp.Base.Code != errno.SuccessCode {
		fmt.Println(string(body))
		time.Sleep(time.Second * 3)
	}

}
func main() {
	filepath.Walk("ai/auto_insert_client/images", func(path string, info fs.FileInfo, err error) error {
		//fmt.Printf("visited file or dir: %q\n", path)
		putAvatar(login(register()), path)
		return nil
	})

}

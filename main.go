package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "ForumApp",
		Usage: "A simple CLI forum application",
		Commands: []*cli.Command{
			{
				Name:    "SignIn",
				Aliases: []string{"s"},
				Usage:   "签到",
				Action: func(c *cli.Context) error {
					userCookie := c.Args().First()
					if userCookie == "" {
						return fmt.Errorf("请输入用户的Cookie")
					}

					Qiandao(userCookie)
					fmt.Println("签到成功")
					return nil
				},
			},
			{
				Name:    "apply task",
				Aliases: []string{"a"},
				Usage:   "申领任务",
				Action: func(c *cli.Context) error {
					userCookie := c.Args().First()
					taskId := c.Args().Get(1)
					if userCookie == "" || taskId == "" {
						return fmt.Errorf("请输入完整的信息")
					}

					applyTask(userCookie, taskId)
					fmt.Println("申领任务成功")
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
func Qiandao(userCookie string) {
	// 定义请求 URL
	requestURL := "http://www.1050qm.com/plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=0&inajax=0"

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置 Cookie
	cookie := &http.Cookie{
		Name:  "b-user-id",
		Value: userCookie,
	}
	req.AddCookie(cookie)

	// 添加 Referer 和 User-Agent
	req.Header.Set("Referer", "http://www.1000qm.vip/home.php?mod=task")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 通过 charset.NewReader 解决中文乱码问题
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Error creating charset reader:", err)
		return
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
func applyTask(userCookie string, taskId string) {
	// 定义请求 URL
	requestURL := "http://www.1000qm.vip/home.php?mod=task&do=apply&id=" + taskId

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置 Cookie
	cookie := &http.Cookie{
		Name:  "b-user-id",
		Value: userCookie,
	}
	req.AddCookie(cookie)

	// 添加 Referer 和 User-Agent
	req.Header.Set("Referer", "http://www.1000qm.vip/home.php?mod=task")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 通过 charset.NewReader 解决中文乱码问题
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Error creating charset reader:", err)
		return
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}

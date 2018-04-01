package notes

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/urfave/cli"
)

func init() {
	addCommand(cli.Command{
		Name:    "baiduwangpan",
		Aliases: []string{"bdwp"},
		Usage:   "百度网盘的一些操作",
		Subcommands: []cli.Command{
			{
				Name:    "quota",
				Aliases: []string{"q"},
				Usage:   "查询配额",
				Action:  quota,
			},
		},
	})
}

var app_id = "260149"

var pcsURL = &url.URL{
	Scheme: "http",
	Host:   "pcs.baidu.com",
}

var panURL = &url.URL{
	Scheme: "http",
	Host:   "pan.baidu.com",
}

func quota(c *cli.Context) (err error) {
	bduss := c.Args().First()

	cookie := &http.Cookie{
		Name:  "BDUSS",
		Value: bduss,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	jar.SetCookies(pcsURL, []*http.Cookie{
		cookie,
	})

	jar.SetCookies(panURL, []*http.Cookie{
		cookie,
	})

	client := &http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", "https://pcs.baidu.com/rest/2.0/pcs/quota", nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	v := req.URL.Query()
	v.Add("method", "info")
	v.Add("app_id", app_id)
	req.URL.RawQuery = v.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Println(string(body))

	return
}

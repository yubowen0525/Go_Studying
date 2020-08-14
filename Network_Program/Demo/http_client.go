package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{
		//CheckRedirect: redirePolicyFunc,
	}

	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.Header.Add("User-Agent", "Our Custom User-Agent")
	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)

	resp, err = client.Do(req)

	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	io.Copy(os.Stdout, resp.Body)

}

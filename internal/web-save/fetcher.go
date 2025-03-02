package web_save

import (
	"compress/gzip"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"os"
	configPkg "permission-cat/config"
	"time"
)

var config = configPkg.Conf

var HEADERS = map[string]string{
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Accept-encoding":           "gzip, deflate, br, zstd",
	"Accept-language":           "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"Cache-Control":             "max-age=0",
	"Priority":                  "u=0, i",
	"Referer":                   "https://www.google.com/",
	"sec-ch-ua":                 `"Not(A:Brand";v="99", "Microsoft Edge";v="133", "Chromium";v="133"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"Windows"`,
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0`,
}

func fetchContent(targetURL string, useProxy bool) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range HEADERS {
		req.Header.Set(k, v)
	}

	var client *http.Client

	if useProxy {
		proxy, _ := url.Parse(config.WebSave.Proxy)

		transport := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 5,
		}
	} else {
		client = http.DefaultClient
	}

	resp, _ := client.Do(req)
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	defer gzipReader.Close()
	bodyText, _ := io.ReadAll(gzipReader)

	return bodyText, nil
}

func Run(url string, useProxy bool) error {
	html, err := fetchContent(url, useProxy)
	if err != nil {
		return err
	}

	fd, err := os.Create("../test.html")
	if err != nil {
		return err
	}
	fd.Write(html)
	defer fd.Close()

	return nil
	//convertHtmlToMarkdown(html)
}

package main

import "github.com/valyala/fasthttp"

func DoRequest(url string) (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", nil
	}

	bodyBytes := resp.Body()

	return string(bodyBytes), nil
}

func DoPostRequestExample(url string) (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethodBytes([]byte("POST"))
	req.PostArgs().Add("value", "0.13")

	client := fasthttp.Client{}
	client.MaxConnsPerHost = 1000

	if err := client.Do(req, resp); err != nil {
		return "", nil
	}

	bodyBytes := resp.Body()

	return string(bodyBytes), nil
}

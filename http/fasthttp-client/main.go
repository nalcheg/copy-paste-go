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

package middleware

import "github.com/valyala/fasthttp"

type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler

func Get(router fasthttp.RequestHandler) fasthttp.RequestHandler {
	return MultipleMiddleware(router, firstMiddleware, secondMiddleware)
}

func MultipleMiddleware(h fasthttp.RequestHandler, m ...Middleware) fasthttp.RequestHandler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func firstMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Passed-First-Middleware", "true")
		next(ctx)
	}
}

// example for decision
func secondMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if string(ctx.FormValue("pass")) == "true" {
			ctx.Response.Header.Set("Passed-Second-Middleware", "true")
			next(ctx)
		} else {
			ctx.Response.Header.Set("Passed-Second-Middleware", "false")
			ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
		}
	}
}

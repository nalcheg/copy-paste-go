package consts

const (
	RabbitmqDSN         = "amqp://guest:guest@127.0.0.1:5672"
	ExchangeName        = "apigateway"
	RequestsQueue       = "requests"
	RequestsRoutingKey  = "request"
	ResponsesQueue      = "responses"
	ResponsesRoutingKey = "response"
	RequestIDHeader     = "reqID"
)

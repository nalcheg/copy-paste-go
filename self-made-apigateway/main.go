package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/valyala/fasthttp"

	"github.com/nalcheg/copy-paste-go/self-made-apigateway/consts"
	"github.com/nalcheg/copy-paste-go/self-made-apigateway/rabbitmq"
)

type Requests struct {
	sync.RWMutex

	r map[uuid.UUID]*fasthttp.RequestCtx
}

type Responses struct {
	sync.RWMutex

	r map[uuid.UUID][]byte
}

type ApiGateway struct {
	requests            Requests
	responses           Responses
	publisher           *rabbitmq.Publisher
	handlerSleepTimeout time.Duration
}

func main() {
	rabbitConn, rabbitCh, err := rabbitmq.Connect(consts.RabbitmqDSN)
	if err != nil {
		log.Panic(err)
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	if err := rabbitmq.Prepare(rabbitCh); err != nil {
		log.Panic(err)
	}

	handlerTimeout := 50 * time.Microsecond
	if len(os.Args) >= 2 {
		handlerTimeout, err = time.ParseDuration(os.Args[1])
		if err != nil {
			log.Panic(err)
		}
	}

	apiGateway := &ApiGateway{
		requests: Requests{
			r: make(map[uuid.UUID]*fasthttp.RequestCtx),
		},
		responses: Responses{
			r: make(map[uuid.UUID][]byte),
		},
		publisher:           rabbitmq.NewPublisher(rabbitCh),
		handlerSleepTimeout: handlerTimeout,
	}
	go apiGateway.respListener(rabbitCh)

	log.Panic(fasthttp.ListenAndServe(":8080", apiGateway.handler))
}

func (ag *ApiGateway) respListener(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		consts.ResponsesQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	for m := range msgs {
		ag.responses.Lock()
		ag.responses.r[uuid.MustParse(fmt.Sprintf("%s", m.Headers[consts.RequestIDHeader]))] = m.Body
		ag.responses.Unlock()

		if err := m.Ack(true); err != nil {
			log.Fatal(err)
		}
	}
}

func (ag *ApiGateway) handler(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	reqID := uuid.New()
	ctx.Response.Header.Add(consts.RequestIDHeader, reqID.String())

	ag.requests.Lock()
	ag.requests.r[reqID] = ctx
	ag.requests.Unlock()

	if err := ag.publisher.Publish(consts.ExchangeName, consts.RequestsRoutingKey, reqID, body); err != nil {
		log.Panic(err)
	}

	for {
		ag.responses.RLock()
		for k, v := range ag.responses.r {
			if k == reqID {
				ag.requests.Lock()
				delete(ag.requests.r, k)
				ag.requests.Unlock()

				ag.responses.RUnlock()
				ag.responses.Lock()
				delete(ag.responses.r, k)
				ag.responses.Unlock()

				ctx.SetBody(v)

				return
			}
		}
		ag.responses.RUnlock()
		time.Sleep(ag.handlerSleepTimeout)
	}
}

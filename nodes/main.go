package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
)

type Node struct {
	IsPrimary *atomic.Bool
	ID        int64
	Addr      string
}

type nodeService struct {
	sync.Mutex
	ID   int64
	List []*Node
}

func NewNodeService(ID int64, list []*Node) *nodeService {
	return &nodeService{ID: ID, List: list}
}

const ElectionDuration = 200 * time.Millisecond

func main() {
	idUint, err := strconv.ParseInt(os.Getenv("ID"), 10, 8)
	if err != nil {
		log.Fatal(err)
	}

	ns := NewNodeService(idUint,
		[]*Node{
			{
				IsPrimary: atomic.NewBool(false),
				ID:        0,
				Addr:      "http://127.0.0.1:58080/state",
			}, {
				IsPrimary: atomic.NewBool(false),
				ID:        1,
				Addr:      "http://127.0.0.1:58081/state",
			}, {
				IsPrimary: atomic.NewBool(false),
				ID:        2,
				Addr:      "http://127.0.0.1:58082/state",
			},
		},
	)

	r := router.New()
	r.GET("/state", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetStatusCode(204)
	})

	port := "58080"
	if os.Getenv("PORT") != "" {
	}
	port = os.Getenv("PORT")

	go func() {
		log.Fatal(fasthttp.ListenAndServe(":"+port, r.Handler))
	}()

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		for {
			ch <- struct{}{}
			time.Sleep(ElectionDuration)
		}
	}(ch)

	if err := ns.Run(ch); err != nil {
		return
	}
}

func (ns *nodeService) Run(ch chan struct{}) error {
	for {
		select {
		case <-ch:
			ns.checkPrimary()
		}
	}
}

func (ns *nodeService) checkPrimary() {
	ns.Lock()
	defer ns.Unlock()

	gr := errgroup.Group{}
	for _, node := range ns.List {
		node := node
		gr.Go(func() error {
			req := fasthttp.AcquireRequest()
			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseRequest(req)
			defer fasthttp.ReleaseResponse(resp)

			req.SetRequestURI(node.Addr)

			if err := fasthttp.Do(req, resp); err != nil {
				node.IsPrimary.Store(false)

				return nil
			}

			if resp.StatusCode() == 204 {
				node.IsPrimary.Store(true)
			}

			return nil
		})
	}
	if err := gr.Wait(); err != nil {
		log.Fatal(err)
	}

	f := false
	for _, node := range ns.List {
		if f {
			node.IsPrimary.Store(false)
		} else if node.IsPrimary.Load() == true {
			f = true
		}
	}

	log.Printf(
		"%t - %t - %t",
		ns.List[0].IsPrimary.Load(),
		ns.List[1].IsPrimary.Load(),
		ns.List[2].IsPrimary.Load(),
	)
}

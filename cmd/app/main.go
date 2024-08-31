package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mirhijinam/wb-l0/internal/api"
	"github.com/mirhijinam/wb-l0/internal/config"
	"github.com/mirhijinam/wb-l0/internal/models"
	"github.com/mirhijinam/wb-l0/internal/pkg/db"
	"github.com/mirhijinam/wb-l0/internal/repository"
	"github.com/mirhijinam/wb-l0/internal/service"
	"github.com/nats-io/stan.go"
)

var (
	clusterID    = "test-cluster"
	subscriberID = "subid"
	publisherID  = "pubid"
	subject      = "foo"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := db.MustOpenDB(ctx, cfg)
	if err != nil {
		panic(err)
	}

	r := repository.New(pool)

	sc, err := stan.Connect(clusterID, subscriberID)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
		return
	}
	defer func() {
		if err := sc.Close(); err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}()

	s := service.New(sc, &r)

	err = s.PutIntoCache(ctx)
	if err != nil {
		log.Fatalf("failed to fill cache: %v", err)
	}

	sub, err := sc.Subscribe(subject, func(m *stan.Msg) {
		order := models.OrderData{}

		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Printf("failed to unmarshal msg: %v", err.Error())
			return
		}

		if err := r.Save(ctx, order); err != nil {
			log.Println("failed to save order:", err)
		}

		s.Cache[order.OrderUid] = order
	})
	if err != nil {
		return
	}
	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			log.Println("failed to unsubscribe:", err)
		}
	}()

	api := api.New(&s)

	server := http.NewServeMux()
	server.HandleFunc("GET /order/{id}", api.GetOrderById)
	log.Fatal(http.ListenAndServe(":8080", server))
}

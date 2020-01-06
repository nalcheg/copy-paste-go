package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DatabaseName   = "temp"
	CollectionName = "temp"
)

type Mongo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

type Config struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type Ticker struct {
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Avg     float64 `json:"avg"`
	Vol     float64 `json:"vol"`
	VolCur  float64 `json:"vol_cur"`
	Last    float64 `json:"last"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int     `json:"updated"`
}

type TickerDocument struct {
	ID        string    `bson:"id"`
	Ticker    Ticker    `bson:"ticker"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (m *Mongo) Connect(dsn string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	m.client = client
	m.database = client.Database(DatabaseName)
	m.collection = client.Database(DatabaseName).Collection(CollectionName)

	return nil
}

func (m *Mongo) LoadDump(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var dump []interface{}
	if err := json.Unmarshal(data, &dump); err != nil {
		return err
	}

	if _, err := m.collection.InsertMany(context.Background(), dump); err != nil {
		return err
	}

	return nil
}

func (m *Mongo) UnloadDump(s string) error {
	cursor, err := m.collection.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return err
	}

	var result []interface{}
	if err := cursor.All(context.Background(), &result); err != nil {
		return err
	}

	f, err := os.Create("tempdata.json")
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

func (m *Mongo) InsertIfNotExists(value Config) error {
	filter := bson.M{"id": value.ID}

	var result Config
	err := m.collection.FindOne(nil, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		_, err := m.collection.InsertOne(nil, value)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	singleResult := m.collection.FindOneAndReplace(nil, filter, value)
	if singleResult.Err() != nil {
		return singleResult.Err()
	}

	return nil
}

func (m *Mongo) SelectExampleOne(currency string) ([]string, error) {
	filter := bson.M{
		"ID": bson.M{
			"$regex": currency,
		},
	}

	var limit int64 = 10
	opts := options.FindOptions{Limit: &limit, Sort: bson.M{"Ticker.vol": -1}}

	c, err := m.collection.Find(context.Background(), filter, &opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := c.Close(context.Background()); err != nil {
			panic(err)
		}
	}()

	var pairs []string
	for c.Next(context.Background()) == true {
		var tickerDocument TickerDocument
		if err = c.Decode(&tickerDocument); err != nil {
			return nil, err
		}
		pairs = append(pairs, tickerDocument.ID)
	}

	return pairs, nil
}

func (m *Mongo) SelectExampleTwoStraight(pairs []string) ([]string, error) {
	var linkedPairs []string

	for _, value := range pairs {
		currencies := strings.Split(value, "_")
		pairs, err := m.SelectExampleOne(currencies[0])
		if err != nil {
			panic(err)
		}

		for _, pair := range pairs {
			linkedPairs = append(linkedPairs, pair)
		}
	}

	return linkedPairs, nil
}

func (m *Mongo) SelectExampleThreeConcurent(pairs []string) ([]string, error) {
	resultChannel := make(chan []string, 10)
	wg := sync.WaitGroup{}

	for _, value := range pairs {
		wg.Add(1)
		currencies := strings.Split(value, "_")
		go func() {
			pairs, err := m.SelectExampleOne(currencies[0])
			if err != nil {
				panic(err)
			}
			resultChannel <- pairs
		}()
	}

	var linkedPairs []string
	go func() {
		for {
			select {
			case pairs := <-resultChannel:
				for _, subValue := range pairs {
					linkedPairs = append(linkedPairs, subValue)
				}
				wg.Done()
			}
		}
	}()

	wg.Wait()

	return linkedPairs, nil
}

func (m *Mongo) FlushCollection() error {
	if err := m.collection.Drop(context.Background()); err != nil {
		return err
	}

	return nil
}

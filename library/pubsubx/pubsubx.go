package pubsubx

import (
	"context"
	"encoding/json"
	"go-learn/entities"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func PushMessage(payload entities.LogsPayload) {
	log.Println("PushMessage Running........")
	ctx := context.Background()

	key, err := ioutil.ReadFile(os.Getenv("ACCOUNT_PATH"))
	if err != nil {
		log.Println("readFile:", err)
		return
	}

	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECTID"), option.WithCredentialsJSON(key))
	if err != nil {
		log.Println("pubsub.NewClient:", err)
		return
	}

	defer client.Close()
	t := client.Topic(os.Getenv("TOPICID"))

	pb, _ := json.Marshal(payload)

	result := t.Publish(ctx, &pubsub.Message{
		Data: pb,
	})
	var wg sync.WaitGroup
	var errors uint64

	wg.Add(1)
	go func(res *pubsub.PublishResult) {
		defer wg.Done()
		id, err := res.Get(ctx)
		if err != nil {
			log.Println("Failed Publish Message", err)
			atomic.AddUint64(&errors, 1)
			return
		}
		log.Println("Success Publish Message")
		log.Println("message id", id)
	}(result)
	wg.Wait()

	if errors > 0 {
		log.Println("messages did not publish successfully")
		return
	}

	return
}

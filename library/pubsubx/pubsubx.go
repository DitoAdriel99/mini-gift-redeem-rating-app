package pubsubx

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
)

func PushMessage(payload entities.LogsPayload) error {
	log.Println("PushMessage Running........")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
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
		return fmt.Errorf("messages did not publish successfully")
	}

	return nil
}

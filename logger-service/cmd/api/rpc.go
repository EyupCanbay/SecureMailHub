package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// RPCServer is the type for our rpc server. methods that take this as a receiver are available
// over RPC, as long as they are exported.
type RPCServer struct {
}

//RPCServer is the type for data we receive form rpc
type RPCPayload struct {
	Name string
	Data string
}

//logInfo is the type for data we receive form rpc
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	//resp is the message sent back to the rpc caller
	*resp = "Processed payload via RPC" + payload.Name
	return nil
}

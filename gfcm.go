package go_fcm

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

func SetFirebaseAdmin(serviceKeyAccount string) (app *firebase.App, err error) {
	opt := option.WithCredentialsFile(serviceKeyAccount)
	app, err = firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		err = fmt.Errorf("error initializing app: %v", err)
	}

	return app, err
}

func SentSingleClient(app *firebase.App, receiver string, data map[string]string) (
	response string, err error) {

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := receiver

	// See firebase messaging documentation on defining a message payload.
	message := &messaging.Message{
		Data: data,
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err = client.Send(ctx, message)

	// Response is a message ID string.
	return response, err
}

func SentMultiClient(app *firebase.App, receivers []string, data map[string]string) (
	batchResponse *messaging.BatchResponse, err error){

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// Create a list containing up to 100 registration tokens.
	// This registration tokens come from the client FCM SDKs.
	registrationTokens := receivers

	// See firebase messaging documentation on defining a message payload.
	message := &messaging.MulticastMessage{
		Data: data,
		Tokens: registrationTokens,
	}

	// See the BatchResponse reference documentation
	// for the contents of response.
	batchResponse, err = client.SendMulticast(context.Background(), message)

	return batchResponse, err
}

func SentToTopic(app *firebase.App, topic string, data map[string]string) (
	response string, err error){

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// The topic name can be optionally prefixed with "/topics/".
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: data,
		Topic: topic,
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err = client.Send(ctx, message)

	return response, err
}

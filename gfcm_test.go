package go_fcm

import (
	firebase "firebase.google.com/go"
	"os"
	"testing"
)

var App *firebase.App
var Data = map[string]string{
	"title": "tes",
	"content": "content",
}
var Topic = "tesTopic"

func TestSetFirebaseAdmin(t *testing.T) {
	fApp, err := SetFirebaseAdmin(os.Getenv("service_account_key"))

	if err != nil {
		t.Fatalf("error set firebase admin %v", err)
	}

	App = fApp
}

func TestSentSingleClient(t *testing.T) {
	_, err := SentSingleClient(App, os.Getenv("client1"), Data)

	if err != nil {
		t.Fatalf("error sent single client; %v", err)
	}
}

func TestSentMultiClient(t *testing.T) {
	_, err := SentMultiClient(App, []string{
		os.Getenv("client1"),
		os.Getenv("client2"),
	}, Data)

	if err != nil {
		t.Fatalf("error sent multi client: %v", err)
	}
}

func TestSentToTopic(t *testing.T) {
	_, err := SentToTopic(App, Topic, Data)

	if err != nil {
		t.Fatalf("error sent to topic: %v", err)
	}
}
### Example

- installation
``` go get github.com/michaelwp/go-fcm  ```

- Set service account key, To create the file you need to go to this link :
https://console.firebase.google.com/project/**YOUR_PROJECT**/settings/serviceaccounts/adminsdk

- import
```go
import (
	gfcm "github.com/michaelwp/go-fcm"
)
```

- Set Firebase Admin
```go
app, err := gfcm.SetFirebaseAdmin("serviceAccountKey.json")
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

- Single Receiver
```go
singleReceiver := "eOznMhsdSxmp74ZMNLe0V2:APA91bGQE-VjFMkKSFZ5E9sdRaecJr2wDn-rRDYsH6u52ewkI6-FrHgIfbghoe7zRjoPYeM5kQOg-hJCWZEfXtZaHjzYlMnJVzxi_-CVcEoCiz7cr7SXQFgoNKi7uxZCHnFthh0JAwpP"

data := map[string] string {
    "score": "850",
    "time":  "2:45",
}

res, err := SentSingleClient(app, singleReceiver, data)
if err != nil {
    fmt.Println("Error sent single client: ", err)
}
fmt.Println("Successfully sent single client: ", res)
```

- Multi Receivers
```go
multiReceiver := []string{
        "fuv9sAo_uJ4:APA91bEAgx7yL0cRNSsiSQlPnKtgJvAtWAQlnu6yT4jnxrLR4YmXAnev5bDPVRRRcIoWp58F6wCi07oS_-g8Pza82dYbxrLmCMTgGv7LsScmsviVuwxgtiMvUm6rKppPY11_lM7lqQBD",
        "eOznMhsdSxmp74ZMNLe0V2:APA91bGQE-VjFMkKSFZ5E9sdRaecJr2wDn-rRDYsH6u52ewkI6-FrHgIfbghoe7zRjoPYeM5kQOg-hJCWZEfXtZaHjzYlMnJVzxi_-CVcEoCiz7cr7SXQFgoNKi7uxZCHnFthh0JAwpP",
    }

data := map[string] string {
    "score": "850",
    "time":  "2:45",
}

bRes, err := SentMultiClient(app, multiReceiver, data)
if err != nil {
    fmt.Println("Error sent multi client: ", err)
}
fmt.Printf("%d messages were sent successfully\n", bRes.SuccessCount)
```

- Topic
```go
topic := "highScores"

data := map[string] string {
    "score": "850",
    "time":  "2:45",
}


res, err = SentToTopic(app, topic, data)
if err != nil {
    fmt.Println("Error sent to topic: ", err)
}
fmt.Println("Successfully sent message:", res)
```
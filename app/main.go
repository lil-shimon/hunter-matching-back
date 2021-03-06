package main

import (
"context"
	"log"

firebase "firebase.google.com/go"
"google.golang.org/api/option"
)

func main() {
	//初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//データを追加する
	_, err = client.Collection("users").Doc("user2").Set(ctx, map[string]interface{} {
		"first": "example.first",
		"middle": "example.middle",
		"last": "example.last",
		"born": "example.born",
	})
	if err != nil{
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	//切断
	defer client.Close()
}

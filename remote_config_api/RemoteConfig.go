package main

import (
	"fmt"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		print(err.Error)
	}
	print("success" + fmt.Errorf("success: %v", app))
}

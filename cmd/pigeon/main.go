package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/EdoardoLaGreca/pigeon"
	"github.com/EdoardoLaGreca/pigeon/credentials"
)

func main() {
	// Parse arguments to run this function.
	detects := DetectionsParse(os.Args[1:])

	if args := detects.Args(); len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", os.Args[0])
		detects.Usage()
		os.Exit(1)
	}

	// Provide credentials
	creds, err := credentials.NewApplicationCredentials("")
	if err != nil {
		log.Fatalf("unable to fetch credentials: %v\n", err)
		os.Exit(1)
	}

	client := pigeon.NewClient(context.Background(), creds.Provide())
	client.SetLanguageHints([]string{detects.language}, false)
	res, err := client.RequestImageAnnotation(detects.Args(), detects.Feature())
	if err != nil {
		log.Fatalf("unable to request image annotation: %v", err)
	}

	// Marshal annotations from responses
	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		log.Fatalf("unable to marshal the response: %v\n", err)
	}
	fmt.Println(string(body))
}

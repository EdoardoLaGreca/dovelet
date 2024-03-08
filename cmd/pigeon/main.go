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
	args := ParseArgs(os.Args[1:])

	if len(args.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", os.Args[0])
		args.Usage()
		os.Exit(1)
	}

	// Provide credentials
	creds, err := credentials.NewApplicationCredentials("")
	if err != nil {
		log.Fatalf("unable to fetch credentials: %v\n", err)
	}

	client := pigeon.NewClient(context.Background(), creds.Provide())
	if len(args.Language()) > 0 {
		client.SetLanguageHints(args.Language(), false)
	}
	res, err := client.RequestImageAnnotation(args.Args(), args.Feature())
	if err != nil {
		log.Fatalf("unable to request image annotation: %v\n", err)
	}

	// Marshal annotations from responses
	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		log.Fatalf("unable to marshal the response: %v\n", err)
	}
	fmt.Println(string(body))
}

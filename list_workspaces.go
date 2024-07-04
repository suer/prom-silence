package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/amp"
)

func ListWorkspaces() ([]byte, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := amp.NewFromConfig(cfg)
	response, error := client.ListWorkspaces(ctx, &amp.ListWorkspacesInput{})
	if error != nil {
		return nil, error
	}

	return json.Marshal(response.Workspaces)
}

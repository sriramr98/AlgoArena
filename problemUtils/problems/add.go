package problems

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"os"
)

type ProblemMeta struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Topic       string   `json:"topic"`
}

var BatchSize = 25 // BatchSize defines the number of problems to be processed in a single batch.

func AddProblems(problemListMetaPath string, problemsPath string) error {
	client := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
	)

	problemsList, err := readProblemListMeta(problemListMetaPath)
	if err != nil {
		return fmt.Errorf("failed to read problem list metadata: %w", err)
	}

	if len(problemsList) == 0 {
		return fmt.Errorf("no problems found in the provided metadata file")
	}

	var requests []anthropic.MessageBatchNewParamsRequest
	for idx, problem := range problemsList {
		requests = append(requests, anthropic.MessageBatchNewParamsRequest{
			CustomID: fmt.Sprintf("problem-%s-%d", problem.Title, idx),
		})
	}

	var batchRequests [][]anthropic.MessageBatchNewParamsRequest
	for i := 0; i < len(requests); i += BatchSize {
		end := i + BatchSize
		if end > len(requests) {
			end = len(requests)
		}
		batchRequests = append(batchRequests, requests[i:end])
	}

	for _, batch := range batchRequests {
		processBatch(client, batch)
	}

	return nil
}

func readProblemListMeta(problemListMetaPath string) ([]ProblemMeta, error) {
	file, err := os.Open(problemListMetaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open problem list metadata file: %w", err)
	}
	defer file.Close()
	var problemsList []ProblemMeta
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&problemsList)
	if err != nil {
		return nil, fmt.Errorf("failed to decode problem list metadata: %w", err)
	}
	return problemsList, nil
}

func processBatch(client anthropic.Client, requests []anthropic.MessageBatchNewParamsRequest) {
	batchResp, err := client.Messages.Batches.New(context.TODO(), anthropic.MessageBatchNewParams{
		Requests: requests,
	})
	if err != nil {
		fmt.Printf("Error processing batch: %v\n", err)
		return
	}

	fmt.Printf("Batch initiated successfully: %+v requests\n", batchResp)

	if err := monitorBatchStatus(client, batchResp.ID); err != nil {
		fmt.Printf("Error monitoring batch status: %v\n", err)
		return
	}

	fmt.Println("Batch processed successfully.")
}

func monitorBatchStatus(client anthropic.Client, batchID string) error {
	var batchEnded bool = false

	for !batchEnded {
		batchStatus, err := client.Messages.Batches.Get(context.TODO(), batchID)
		if err != nil {
			return fmt.Errorf("failed to get batch status: %w", err)
		}

		if batchStatus.ProcessingStatus == anthropic.MessageBatchProcessingStatusEnded {
			batchEnded = true
		}

		fmt.Printf("Batch %s processed successfully with status: %+v\n", batchID, batchStatus)
	}

	return nil
}

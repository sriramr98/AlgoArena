package problems

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"os"
	"path"
)

type ProblemMeta struct {
	Title      string   `json:"title"`
	Difficulty string   `json:"difficulty"`
	Tags       []string `json:"tags"`
	Topic      string   `json:"topic"`
}

type BatchResponse struct {
}

var AISystemPromptRole string = `
You are an assistant at AlgoArena which is a DSA platform which is used by millions of software engineers to prepare Data Structures and Algorithms interviews. 
Your job at AlgoArena is to create new problems that users can solve to prepare for their interviews.
You will be provided a Problem Title, some tags regarding the problem and a difficulty. 
You will use this information to create a new problem in the JSON format specified. 
Here are some rules to follow while creating the problem:
`

var AISystemPromptRules string = `
Rules:
1. The problem description should not be a generic description of the problem, but rather modelled after a real-world problem that can be solved using the given data structures and algorithms.
2. The problem description should be formatted in HTML, with appropriate headings, bullet points, and code blocks to enhance readability.
3. The test cases should be comprehensive and cover edge cases, ensuring that the problem is well-defined and can be solved correctly.

Schema:
{schema}
`
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

	prepareAIRules(problemsPath)

	var requests []anthropic.MessageBatchNewParamsRequest
	for idx, problem := range problemsList {
		requests = append(requests, anthropic.MessageBatchNewParamsRequest{
			CustomID: fmt.Sprintf("problem-%s-%d", problem.Title, idx),
			Params: anthropic.MessageBatchNewParamsRequestParams{
				Model: anthropic.ModelClaude3_7SonnetLatest,
				System: []anthropic.TextBlockParam{
					{Text: AISystemPromptRole, Type: "text", CacheControl: anthropic.NewCacheControlEphemeralParam()},
					{Text: AISystemPromptRules, Type: "text", CacheControl: anthropic.NewCacheControlEphemeralParam()},
				},
				Messages: []anthropic.MessageParam{
					{
						Content: []anthropic.ContentBlockParamUnion{
							{
								OfText: &anthropic.TextBlockParam{
									Text: fmt.Sprintf(
										"Create a new problem with the title: %s, tags: %v, topic: %s and difficulty: %s",
										problem.Title, problem.Tags, problem.Topic, problem.Difficulty,
									),
								},
							},
						},
						Role: anthropic.MessageParamRoleUser,
					},
				},
			},
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
		batchId, err := processBatch(client, batch)
		if err != nil {
			fmt.Printf("Error processing batch: %v\n", err)
			return err
		}
		results, err := fetchBatchResults(client, batchId)
		if err != nil {
			fmt.Printf("Error fetching batch results: %v\n", err)
			return err
		}

		//TODO: Create problem solutions, validate against test cases and create problem json files
		fmt.Println(results)
	}

	return nil
}

func prepareAIRules(problemsPath string) {
	schemaPath := path.Join(problemsPath, "schema.json")
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		fmt.Printf("Schema file not found at path: %s\n", schemaPath)
		fmt.Printf("Please ensure the schema file is present in the specified path %s.\n", problemsPath)
		return
	}
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		fmt.Printf("Failed to open schema file: %v\n", err)
		return
	}
	defer schemaFile.Close()
	var schema map[string]interface{}
	decoder := json.NewDecoder(schemaFile)
	err = decoder.Decode(&schema)
	if err != nil {
		fmt.Printf("Failed to decode schema file: %v\n", err)
		return
	}

	// Removing solutions since it'll be generated separately and appended later
	delete(schema, "solutions")

	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		fmt.Printf("Failed to marshal schema to JSON: %v\n", err)
		return
	}
	AISystemPromptRules = fmt.Sprintf(AISystemPromptRules, string(schemaJSON))
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

func processBatch(client anthropic.Client, requests []anthropic.MessageBatchNewParamsRequest) (string, error) {
	batchResp, err := client.Messages.Batches.New(context.TODO(), anthropic.MessageBatchNewParams{
		Requests: requests,
	})
	if err != nil {
		fmt.Printf("Error processing batch: %v\n", err)
		return "", err
	}

	fmt.Printf("Batch initiated successfully: %+v requests\n", batchResp)

	if err := monitorBatchStatus(client, batchResp.ID); err != nil {
		fmt.Printf("Error monitoring batch status: %v\n", err)
		return "", err
	}

	fmt.Println("Batch processed successfully.")
	return batchResp.ID, nil
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

func fetchBatchResults(client anthropic.Client, batchId string) ([]BatchResponse, error) {
	return nil, nil
}

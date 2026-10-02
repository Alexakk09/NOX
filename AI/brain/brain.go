package brain

import (
	"encoding/json"

	"Jarvis/clients"
	"Jarvis/tools/registry"
	"Jarvis/voice"
)

type Brain struct {
	api            *clients.APIClient
	tools          *registry.Registry
	conversationID int64
	tts            *voice.TTS
	stt            *voice.STT
}

func toolDefinitions(registry *registry.Registry) []registry.Definition {
	return registry.Definitions()
}

func NewBrain(
	api *clients.APIClient,
	tools *registry.Registry,
) *Brain {
	return &Brain{
		api:   api,
		tools: tools,
		tts:   voice.NewTTS(),
		stt: voice.NewSTT(
			`C:\Users\veerj\whisper.cpp\build\bin\Release\whisper-stream.exe`,
			`C:\Users\veerj\whisper.cpp\ggml-tiny.en.bin`,
		),
	}
}

func (b *Brain) StartConversation() error {
	conversationID, err := b.api.CreateConversation()
	if err != nil {
		return err
	}

	b.conversationID = conversationID

	return nil
}

func (b *Brain) Run(input string) (string, error) {
	definitions := toolDefinitions(b.tools)

	clientDefinitions := make([]clients.ToolDefinition, 0, len(definitions))

	for _, definition := range definitions {
		parameters := make([]clients.ToolParameter, 0, len(definition.Parameters))

		for _, parameter := range definition.Parameters {
			parameters = append(parameters, clients.ToolParameter{
				Name:        parameter.Name,
				Type:        parameter.Type,
				Description: parameter.Description,
				Required:    parameter.Required,
			})
		}

		clientDefinitions = append(clientDefinitions, clients.ToolDefinition{
			Name:        definition.Name,
			Description: definition.Description,
			Parameters:  parameters,
		})
	}

	response, err := b.api.SendChat(clients.ChatRequest{
		ConversationID:  b.conversationID,
		Message:         input,
		ToolDefinitions: clientDefinitions,
	})
	if err != nil {
		return "", err
	}

	for len(response.ToolCalls) > 0 {
		for _, call := range response.ToolCalls {
			result, err := b.tools.Execute(
				call.Name,
				call.Arguments,
			)
			if err != nil {
				return "", err
			}

			resultJSON, err := json.Marshal(result)
			if err != nil {
				return "", err
			}

			response, err = b.api.SendChat(clients.ChatRequest{
				ConversationID: b.conversationID,
				ToolResult: &clients.ToolResult{
					Call: clients.ToolCall{
						ID:        call.ID,
						Name:      call.Name,
						Arguments: call.Arguments,
					},
					Result: string(resultJSON),
				},
				ToolDefinitions: clientDefinitions,
			})
			if err != nil {
				return "", err
			}
		}
	}

	return response.Response, nil
}

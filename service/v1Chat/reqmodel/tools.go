package reqmodel

import (
	"free-gpt3.5-2api/service/v1"
	"github.com/google/uuid"
)

func copyParameter(parameter ApiFunctionParameter) ChatFunctionParameter {
	var cpyParam ChatFunctionParameter
	switch parameter.Type {
	case "string":
		cpyParam = ChatFunctionParameter{
			Type:        parameter.Type,
			Enum:        parameter.Enum,
			Description: parameter.Description,
		}
	case "object":
		cpyProperties := make(map[string]ChatFunctionParameter)
		for key, value := range parameter.Properties {
			cpyProperties[key] = copyParameter(value)
		}
		cpyParam = ChatFunctionParameter{
			Type:       parameter.Type,
			Properties: cpyProperties,
			Required:   parameter.Required,
		}
	}
	return cpyParam
}

func ApiReq2ChatReq35(apiReq *ApiReq) (chatReq *ChatReq35) {
	messages := make([]ChatMessages, 0)
	for _, apiMessage := range apiReq.Messages {
		chatMessage := ChatMessages{
			Author: ChatAuthor{
				Role: apiMessage.Role,
			},
			Content: ChatContent{
				ContentType: "text",
				Parts:       []string{apiMessage.Content},
			},
		}
		messages = append(messages, chatMessage)
	}

	tools := make([]ChatTool, 0)
	for _, apiTool := range apiReq.Tools {
		apiFunctionParameters := apiTool.Function.Parameters
		chatTool := ChatTool{
			Type: apiTool.Type,
			Function: ChatFunction{
				Name:        apiTool.Function.Name,
				Description: apiTool.Function.Description,
				Parameters:  copyParameter(apiFunctionParameters),
			},
		}
		tools = append(tools, chatTool)
	}

	chatReq = &ChatReq35{
		Action:                     "next",
		Messages:                   messages,
		ParentMessageId:            uuid.New().String(),
		Model:                      v1.MappingModel(apiReq.Model),
		TimeZoneOffsetMin:          -180,
		Suggestions:                make([]string, 0),
		HistoryAndTrainingDisabled: true,
		ConversationMode: ChatConversationMode{
			Kind: "primary_assistant",
		},
		WebsocketRequestId: uuid.New().String(),
		Tools:              tools,
	}
	return chatReq
}

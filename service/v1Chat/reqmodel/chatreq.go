package reqmodel

type ChatAuthor struct {
	Role string `json:"role"`
}

type ChatContent struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type ChatMessages struct {
	Author  ChatAuthor  `json:"author"`
	Content ChatContent `json:"content"`
}

type ChatConversationMode struct {
	Kind string `json:"kind"`
}

type ChatTool struct {
	Type     string       `json:"type"`
	Function ChatFunction `json:"function,omitempty"`
}

type ChatFunction struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Parameters  ChatFunctionParameter `json:"parameters"`
}

type ChatFunctionParameter struct {
	Type        string                           `json:"type"`
	Description string                           `json:"description,omitempty"`
	Properties  map[string]ChatFunctionParameter `json:"properties,omitempty"`
	Required    []string                         `json:"required,omitempty"`
	Enum        []string                         `json:"enum,omitempty"`
}

type ChatReq35 struct {
	Action                     string               `json:"action"`
	Messages                   []ChatMessages       `json:"messages"`
	ParentMessageId            string               `json:"parent_message_id"`
	Model                      string               `json:"model"`
	TimeZoneOffsetMin          int                  `json:"timezone_offset_min"`
	Suggestions                []string             `json:"suggestions"`
	HistoryAndTrainingDisabled bool                 `json:"history_and_training_disabled"`
	ConversationMode           ChatConversationMode `json:"conversation_mode"`
	WebsocketRequestId         string               `json:"websocket_request_id"`
	Tools                      []ChatTool           `json:"tools,omitempty"`
}

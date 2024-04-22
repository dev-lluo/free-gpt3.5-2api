package reqmodel

type ApiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ApiReq struct {
	Messages    []ApiMessage `json:"messages"`
	Model       string       `json:"model"`
	Stream      bool         `json:"stream"`
	PluginIds   []string     `json:"plugin_ids"`
	NewMessages string
	Tools       []ApiTool `json:"tools,omitempty"`
}

type ApiTool struct {
	Type     string      `json:"type"`
	Function ApiFunction `json:"function,omitempty"`
}

type ApiFunction struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Parameters  ApiFunctionParameter `json:"parameters"`
}

type ApiFunctionParameter struct {
	Type        string                          `json:"type"`
	Description string                          `json:"description,omitempty"`
	Properties  map[string]ApiFunctionParameter `json:"properties,omitempty"`
	Required    []string                        `json:"required,omitempty"`
	Enum        []string                        `json:"enum,omitempty"`
}

package model

import "context"

// A Model sends prompts to a specific GenAI model using an Endpoint location, where the model is enabled and billed
type Model struct {
	prompt  func(ctx context.Context, modelName string, cfg Config, args []string) error
	mFamily string
	mType   string
	mName   string
}

var Models map[string]Model = map[string]Model{
	"gemini-1.0-pro-001": {
		prompt:  UseGeminiModel,
		mFamily: "Gemini",
		mType:   "text",
		mName:   "gemini-1.0-pro-001",
	},
	"gemini-1.0-ultra-001": {
		prompt:  UseGeminiModel,
		mFamily: "Gemini",
		mType:   "text",
		mName:   "gemini-1.0-ultra-001",
	},
	"gemini-1.0-pro-vision-001": {
		prompt:  UseGeminiModel,
		mFamily: "Gemini",
		mType:   "text",
		mName:   "gemini-1.0-pro-vision-001",
	},
	"gemini-1.0-ultra-vision-001": {
		prompt:  UseGeminiModel,
		mFamily: "Gemini",
		mType:   "text",
		mName:   "gemini-1.0-ultra-vision-001",
	},
	"gemini-1.5-pro-preview-0215": {
		prompt:  UseGeminiModel,
		mFamily: "Gemini",
		mType:   "text",
		mName:   "gemini-1.5-pro-preview-0215",
	},
	"text-bison": {
		prompt:  UsePaLMModel,
		mFamily: "text",
		mType:   "text",
		mName:   "text-bison",
	},
	"text-bison@001": {
		prompt:  UsePaLMModel,
		mFamily: "text",
		mType:   "text",
		mName:   "text-bison@001",
	},
	"text-bison@002": {
		prompt:  UsePaLMModel,
		mFamily: "text",
		mType:   "text",
		mName:   "text-bison@002",
	},
	"text-unicorn@001": {
		prompt:  UsePaLMModel,
		mFamily: "text",
		mType:   "text",
		mName:   "text-unicorn@001",
	},
	"medlm-medium": {
		prompt:  UsePaLMModel,
		mFamily: "MultiModal",
		mType:   "MultiModal",
		mName:   "medlm-medium",
	},
	"medlm-large": {
		prompt:  UsePaLMModel,
		mFamily: "MultiModal",
		mType:   "MultiModal",
		mName:   "medlm-large",
	},
	"medpalm2@preview": {
		prompt:  UsePaLMModel,
		mFamily: "MultiModal",
		mType:   "MultiModal",
		mName:   "medpalm2@preview",
	},
	"code-bison": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-bison",
	},
	"code-bison@001": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-bison@001",
	},
	"code-bison@002": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-bison@002",
	},
	"code-bison-32k": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-bison-32k",
	},
	"code-bison-32k@002": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-bison-32k@002",
	},
	"code-gecko": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-gecko",
	},
	"code-gecko@001": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-gecko@001",
	},
	"code-gecko@002": {
		mFamily: "Embeddings",
		mType:   "Embeddings",
		mName:   "code-gecko@002",
	},
	"claude-3-haiku@20240307": {
		prompt:  UseClaudeModel,
		mFamily: "MultiModal",
		mType:   "MultiModal",
		mName:   "claude-3-haiku@20240307",
	},
}

// TODO - Ideally would like to avoid this level of indirection, but suing it for the
//
//	time being to get course grained refactoring working
func (m Model) Use(ctx context.Context, cfg Config, args []string) error {
	return m.prompt(ctx, m.mName, cfg, args)
}
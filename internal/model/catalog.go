package model

import (
	"context"
	"embed"
	"encoding/csv"
	"fmt"
	"strings"
)

//go:embed models*
var modelfiles embed.FS

// A Model sends prompts to a specific GenAI model using an Endpoint location, where the model is enabled and billed
type Model struct {
	prompt func(ctx context.Context, modelName string, cfg Config, args []string) error `json:"-"`
	Family string                                                                       `json:"family"`
	Mode   string                                                                       `json:"mode"`
	Name   string                                                                       `json:"name"`
}

// listToModels returns a slice of Models from the embedded CSV file of models
func listToModels() ([]Model, error) {
	var records [][]string
	for _, modelfile := range []string{"models", "models.internal"} {
		var modelrecords [][]string
		modellist, err := modelfiles.ReadFile(modelfile)
		if err != nil {
			continue
		}
		r := csv.NewReader(strings.NewReader(string(modellist)))
		modelrecords, _ = r.ReadAll()
		records = append(records, modelrecords...)
	}

	models := make([]Model, 0, len(records))
	for _, record := range records {
		if strings.HasPrefix(record[0], "#") {
			continue
		}
		model := Model{
			Family: record[0],
			Mode:   record[1],
			Name:   record[2],
		}
		if strings.HasPrefix(model.Family, "gemini") {
			model.prompt = UseGeminiModel
		}
		if strings.HasPrefix(model.Family, "palm") {
			model.prompt = UsePaLMModel
		}
		if strings.HasPrefix(model.Family, "anthropic") {
			model.prompt = UseClaudeModel
		}
		models = append(models, model)
	}
	return models, nil
}

func List() ([]Model, error) {
	return listToModels()
}

func Get(name string) (Model, error) {
	models, err := listToModels()
	if err != nil {
		return Model{}, err
	}
	for _, model := range models {
		if model.Name == name {
			return model, nil
		}
	}
	return Model{}, fmt.Errorf("Model: `%s` not found", name)
}

// TODO - Ideally would like to avoid this level of indirection, but using it for the
//
//	time being to get course grained refactoring working
func (m Model) Use(ctx context.Context, cfg Config, args []string) error {
	if m.prompt != nil {
		return m.prompt(ctx, m.Name, cfg, args)
	}
	return fmt.Errorf("Model: `%s` does not currently implement the `Use` method", m.Name)
}

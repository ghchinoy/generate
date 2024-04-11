package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ghchinoy/gen/internal/model"
	"github.com/spf13/cobra"
	// "google.golang.org/protobuf/encoding/protojson"
	// "google.golang.org/protobuf/types/known/structpb"
	// TODO - Use this import after moving model files to internal/model directory and 'model' package
	// "github.com/ghchinoy/gen/internal/model"
)

var (
	// TODO - Look at ways to remove the need to export these two variable outside the package
	modelName       string
	ModelConfigFile string
	//modelConfig     map[string]interface{}
)

func init() {
	rootCmd.AddCommand(promptCmd)

	//promptCmd.AddCommand(generateContentCmd)

	promptCmd.PersistentFlags().StringVarP(&modelName, "model", "m", "gemini-1.0-pro", "model name")
	promptCmd.PersistentFlags().StringVarP(&ModelConfigFile, "config", "c", "", "model parameters")

	//flag.StringVar(&modelName, "model", "gemini-1.0-pro", "generative model to use")
	//flag.StringVar(&region, "region", "us-central1", "region to use")
	//flag.Parse()
}

var promptCmd = &cobra.Command{
	Use:     "prompt",
	Aliases: []string{"p"},
	Short:   "Prompt a model",
	Long:    `Provide prompt parts to a model to generate content`,
	// Run:     generateContentForModel,
	Run: generateContent,
}

// // generateContentForModel prompts a model to generate content based on the provided prompt.
// func generateContentForModel(cmd *cobra.Command, args []string) {

// 	// Resolve unused argument
// 	_ = cmd

// 	if len(args) == 0 {
// 		fmt.Println("please provide prompt")
// 		os.Exit(1)
// 	}
// 	if logtype != "none" {
// 		log.Printf("model: %s", modelName)
// 		log.Printf("prompt: %s", args)
// 	}

// 	// TODO - Not sure why, but I don't think this is printing to stdout
// 	fmt.Printf("/n --- Model name: %s ---", modelName)

// 	// TODO better as a switch guarded by model list
// 	var err error
// 	if strings.HasPrefix(modelName, "gemini") {
// 		err = useGeminiModel(projectID, region, modelName, args)
// 	} else if strings.Contains(modelName, "text-bison") || strings.Contains(modelName, "text-unicorn") {
// 		err = usePaLMModel(projectID, region, modelName, args)
// 	} else if strings.HasPrefix(modelName, "medlm-") || strings.HasPrefix(modelName, "medpalm") {
// 		err = usePaLMModel(projectID, region, modelName, args)
// 	} else if strings.HasPrefix(modelName, "claude") {
// 		err = useClaudeModel(projectID, region, modelName, args)
// 	} else {
// 		err = fmt.Errorf("model '%s' is not supported", modelName)
// 	}
// 	if err != nil {
// 		log.Printf("error generating content: %v", err)
// 		os.Exit(1)
// 	}
// }

// generateContent prompts a model to generate content based on the provided prompt.
func generateContent(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		// TODO replace with log fatal
		fmt.Println("please provide prompt")
		os.Exit(1)
	}

	cfgB := model.ConfigBuilder{}

	// Set the model configuration
	cfgB.ProjectID(projectID).RegionID(region).ConfigFile(cfgFile).OutputType(Outputtype).LogType(Logtype)
	cfg, err := cfgB.Build()
	if err != nil {
		log.Fatalf("error building config: %v", err)
	}

	if Logtype != "none" {
		log.Printf("model: %s", modelName)
		log.Printf("prompt: %s", args)
	}

	fmt.Printf("/n Model name: %s", modelName)

	// Lookup the model based on the name
	m, ok := model.Models[modelName]
	if !ok {
		log.Printf("model '%s' is not supported", modelName)
		// TODO replace with log.fatal
		os.Exit(1)
	}

	ctx := context.Background()

	err = m.Use(ctx, cfg, args)
	if err != nil {
		log.Printf("error generating content: %v", err)
		os.Exit(1)
	}

}

// // readModelConfigFile reads the model configuration file (JSON text file)
// func readModelConfigFile(configFile string) (map[string]interface{}, error) {

// 	// Resolve unused argument
// 	_ = configFile
// 	var config map[string]interface{}
// 	data, err := os.ReadFile(ModelConfigFile)
// 	if err != nil {
// 		return config, fmt.Errorf("error reading model config: %v", err)

// 	}

// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		return config, fmt.Errorf("error unmarshalling model config: %v", err)
// 	}
// 	return config, nil
// }

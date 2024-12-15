// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SourceDB string `yaml:"source_db"`
	TargetDB string `yaml:"target_db"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file",
	Long:  `This command will guide you through setting up the configuration file for the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI configuration setup!")

		// Questions for the survey
		var questions = []*survey.Question{
			{
				Name:     "sourceDB",
				Prompt:   &survey.Input{Message: "输入源数据库名称:"},
				Validate: survey.Required,
			},
			{
				Name:     "targetDB",
				Prompt:   &survey.Input{Message: "输入目标数据库名称:"},
				Validate: survey.Required,
			},
		}

		// Answers will hold the input data
		answers := struct {
			SourceDB string
			TargetDB string
		}{}

		// Run the survey
		err := survey.Ask(questions, &answers)
		if err != nil {
			fmt.Printf("Error collecting input: %v\n", err)
			os.Exit(1)
		}

		// Create configuration object
		config := Config{
			SourceDB: answers.SourceDB,
			TargetDB: answers.TargetDB,
		}

		// Write to config.yaml
		writeConfig(config)

		fmt.Println("配置文件已经保存至 config.yaml")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// Write configuration to a file
func writeConfig(config Config) {
	file, err := os.Create("config.yaml")
	if err != nil {
		fmt.Printf("Error creating config file: %v\n", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Error closing config file")
		}
	}(file)

	encoder := yaml.NewEncoder(file)
	defer func(encoder *yaml.Encoder) {
		err := encoder.Close()
		if err != nil {
			println("Error closing config file")
		}
	}(encoder)

	if err := encoder.Encode(&config); err != nil {
		fmt.Printf("Error writing to config file: %v\n", err)
		os.Exit(1)
	}
}

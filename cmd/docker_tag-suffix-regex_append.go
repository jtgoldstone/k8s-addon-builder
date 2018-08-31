package cmd

import (
	abd "github.com/GoogleCloudPlatform/addon-builder/pkg"

	"github.com/spf13/cobra"
)

var DockerTagSuffixRegexAppend = &cobra.Command{
	Use:  "append <REGEX> <TAG_SUFFIX>",
	Args: cobra.ExactArgs(2),
	RunE: appendTagSuffixWrapper,
}

func init() {
	DockerTagSuffixRegexCmd.AddCommand(DockerTagSuffixRegexAppend)
}

func appendTagSuffixWrapper(cmd *cobra.Command, args []string) error {
	return abd.EditTagSuffixWrapper(cmd, args, true)
}

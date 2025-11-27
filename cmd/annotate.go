package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const StdIn = "-"

func NewAnnotateCommand() {
	var command = &cobra.Command{
		Use:   "annotate <path>",
		Short: "Check for and annotate namespaces",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Error("<path> argument required to annotate manifests")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var manifests []unstructured.Unstructured
			var err error

			path := args[0]
			if path == StdIn {
				manifests, err = readManifestsData(cmd.InOrStdin())
				if err != nil {
					return err
				}
			} else {
				files, err := listFiles(path)
				if len(files) < 1 {
					return fmt.Errorf("no YAML or JSON files were found in %s", path)
				}
				if err != nil {
					return err
				}

				var errs []error
				manifests, errs = readFilesAsManifests(files)
				if len(errs) != 0 {
					errMessages := make([]string, len(errs))
					for idx, err := range errs {
						errMessages[idx] = err.Error()
					}
					return fmt.Errorf("could not read YAML/JSON files:\n%s", strings.Join(errMessages, "\n"))
				}
			}

			return nil
		},
	}
}

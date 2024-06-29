package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
)

type RootOptions struct {
	IsAdd     bool
	IsDelete  bool
	IsList    bool
	Query     string
	EndPoint  string
	SilenceId string
}

func rootCmd() *cobra.Command {
	opts := &RootOptions{}
	cmd := &cobra.Command{
		Use:           "amp-silence",
		Short:         "",
		Args:          cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
		SilenceUsage:  true,
		SilenceErrors: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IsAdd && !opts.IsDelete && !opts.IsList {
				cmd.Help()
			}
			return run(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.EndPoint, "endpoint", "e", "", "API endpoint URL (ex: https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/)")
	cmd.MarkFlagRequired("endpoint")
	cmd.Flags().BoolVarP(&opts.IsAdd, "add", "a", false, "add silence")
	cmd.Flags().BoolVarP(&opts.IsDelete, "delete", "d", false, "delete silence")
	cmd.Flags().BoolVarP(&opts.IsList, "list", "l", false, "list silences")
	cmd.Flags().StringVarP(&opts.Query, "query", "q", "", "JMESPath query (ex: 'silenceID')")
	cmd.Flags().StringVarP(&opts.SilenceId, "silenceid", "s", "", "silence id (ex: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)")

	return cmd
}

func run(opts *RootOptions) error {
	result := []byte{}
	if opts.IsAdd {
		r, err := AddSilence(opts.EndPoint, os.Stdin)
		if err != nil {
			return err
		}
		result = r
	} else if opts.IsDelete {
		r, err := DeleteSilence(opts.EndPoint, opts.SilenceId)
		if err != nil {
			return err
		}
		result = r
	} else if opts.IsList {
		r, err := ListSilences(opts.EndPoint)
		if err != nil {
			return err
		}
		result = r
	}

	if len(opts.Query) > 0 {
		r, err := applyJMESPath(opts.Query, result)
		if err != nil {
			return err
		}
		result = r
	}

	fmt.Println(string(result))

	return nil
}

func applyJMESPath(query string, data []byte) ([]byte, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	result, err := jmespath.Search(query, jsonData)
	if err != nil {
		return nil, err
	}
	json := fmt.Sprintf("%v", result)
	return []byte(json), nil
}

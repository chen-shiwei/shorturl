package cmd

import (
	"net/http"

	"github.com/chen-shiwei/shorturl/internal/shorturl/router"
	"github.com/spf13/cobra"
)

var (
	addr string

	apiCmd = &cobra.Command{
		Use: "api",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := router.New()

			err := http.ListenAndServe(addr, r)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	apiCmd.Flags().StringVar(&addr, "addr", "", "运行")
}

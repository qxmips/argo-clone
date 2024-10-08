package commands

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {

	var (
		otlpAddress string
	)

	command := &cobra.Command{
		Use:               cliName,
		Short:             "Run the Argo CD API server",
		Long:              "Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes.",
		DisableAutoGenTag: true, // disable cobra auto generation tag
		Run: func(c *cobra.Command, args []string) {
			ctx := c.Context()


			for {
				var closer func()
				ctx, cancel := context.WithCancel(ctx)

				if otlpAddress != "" {
					closer, err = traceutil.InitTracer(ctx, "argocd-server", otlpAddress, otlpInsecure, otlpHeaders, otlpAttrs)
					if err != nil {
						log.Fatalf("failed to initialize tracing: %v", err)
					}
				}

	}

}

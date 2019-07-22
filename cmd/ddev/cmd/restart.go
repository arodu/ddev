package cmd

import (
	"strings"

	"github.com/drud/ddev/pkg/dockerutil"
	"github.com/drud/ddev/pkg/output"
	"github.com/drud/ddev/pkg/util"
	"github.com/spf13/cobra"
)

var restartAll bool

// RestartCmd rebuilds an apps settings
var RestartCmd = &cobra.Command{
	Use:   "restart [projects]",
	Short: "Restart a project or several projects.",
	Long:  `Stops named projects and then starts them back up again.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		dockerutil.EnsureDdevNetwork()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := getRequestedProjects(args, restartAll)
		if err != nil {
			util.Failed("Failed to get project(s): %v", err)
		}

		for _, app := range projects {

			output.UserOut.Printf("Restarting project %s...", app.GetName())
			err = app.Stop(false, false)
			if err != nil {
				util.Failed("Failed to restart %s: %v", app.GetName(), err)
			}

			err = app.Start()
			if err != nil {
				util.Failed("Failed to restart %s: %v", app.GetName(), err)
			}

			util.Success("Restarted %s", app.GetName())
			util.Success("Your project can be reached at %s", strings.Join(app.GetAllURLs(), ", "))
		}
	},
}

func init() {
	RestartCmd.Flags().BoolVarP(&restartAll, "all", "a", false, "restart all projects")
	RootCmd.AddCommand(RestartCmd)
}

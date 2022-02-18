package tests

import (
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common/validator"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewWatchExecutionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "watch <executionID>",
		Aliases: []string{"w"},
		Short:   "Watch logs output from executor pod",
		Long:    `Gets test execution details, until it's in success/error state, blocks until gets complete state`,
		Args:    validator.ExecutionID,
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := common.GetClient(cmd)

			executionID := args[0]
			execution, err := client.GetExecution(executionID)
			ui.ExitOnError("get execution failed", err)

			if execution.ExecutionResult.IsCompleted() {
				ui.Completed("execution is already finished")
			} else {
				watchLogs(executionID, client)
			}

		},
	}
}

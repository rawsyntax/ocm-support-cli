package create

import (
	"fmt"
	"ocm-support-cli/pkg/access_token"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"
)

// Cmd represents the version command
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Creates registry credentials for the current account.",
	Long:  "Created registry credentials for the current account.",
	RunE:  run,
}

func run(cmd *cobra.Command, argv []string) error {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	err = access_token.CreateAccessToken(connection)
	if err != nil {
		return fmt.Errorf("failed to create access token: %v", err)
	}
	fmt.Println("Token generated successfully")
	return nil
}

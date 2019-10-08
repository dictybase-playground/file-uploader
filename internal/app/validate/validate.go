package validate

import (
	"fmt"

	cli "github.com/urfave/cli"
)

// ValidateMinioArgs checks that the necessary flags are not missing
func ValidateMinioArgs(c *cli.Context) error {
	for _, p := range []string{
		"minio-endpoint",
		"minio-access-key",
		"minio-secret-key",
		"minio-bucket",
		"folder",
	} {
		if len(c.String(p)) == 0 {
			return cli.NewExitError(
				fmt.Sprintf("argument %s is missing", p),
				2,
			)
		}
	}
	return nil
}

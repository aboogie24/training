package main

import (
	"training.alfredbrowniii.io/internals/infrastructure"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		// Deploy all infrastructure resources

		infra, err := infrastructure.CreateInfrastructure(ctx)
		if err != nil {
			return err
		}

		ctx.Export("vpcId: ", infra.Vpc.Id)

		return nil
	})
}

// func checkIfErr(err error) {
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 		return
// 	}
// }

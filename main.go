package main

import (
	"fmt"

	"training.alfredbrowniii.io/internals/infrastructure"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := infrastructure.CreateInfrastructure(ctx)
		if err != nil {
			return err
		}
		fmt.Print("you are here")

		return nil
	})
}

// func checkIfErr(err error) {
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 		return
// 	}
// }

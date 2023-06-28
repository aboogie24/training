package infrastructure

import (
	"reflect"

	"training.alfredbrowniii.io/internals/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type infrastructure struct {
	vpc *vpc.Vpc
}

// ElementType implements pulumi.Input.
func (*infrastructure) ElementType() reflect.Type {
	panic("unimplemented")
}

func CreateInfrastructure(ctx *pulumi.Context) (*infrastructure, error) {
	_, err := vpc.CreateVpc(ctx, "training-vpc", "10.0.0.0/16")
	if err != nil {
		return nil, err
	}

	return &infrastructure{
		vpc: &vpc.Vpc{},
	}, nil
}

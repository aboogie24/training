package infrastructure

import (
	"reflect"

	"training.alfredbrowniii.io/internals/subnets"
	"training.alfredbrowniii.io/internals/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type infrastructure struct {
	Vpc           *vpc.Vpc
	PrivateSubnet *subnets.Subnet
	PublicSubnet  *subnets.Subnet
}

// ElementType implements pulumi.Input.
func (*infrastructure) ElementType() reflect.Type {
	panic("unimplemented")
}

func CreateInfrastructure(ctx *pulumi.Context) (*infrastructure, error) {

	vpcs, err := vpc.CreateVpc(ctx, "training-vpc", "10.0.0.0/16")
	if err != nil {
		return nil, err
	}

	privateSubnet, err := subnets.CreateSubnets(ctx, "10.0.0.0/20", vpcs.Id.ToStringOutput())
	if err != nil {
		return nil, err
	}

	return &infrastructure{
		Vpc:           vpcs,
		PrivateSubnet: privateSubnet,
		PublicSubnet:  &subnets.Subnet{},
	}, nil
}

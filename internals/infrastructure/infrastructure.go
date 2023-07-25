package infrastructure

import (
	"fmt"
	"reflect"

	"training.alfredbrowniii.io/internals/config"
	"training.alfredbrowniii.io/internals/subnets"
	"training.alfredbrowniii.io/internals/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Infrastructure struct {
	Config        *config.Config
	Vpc           *vpc.Vpc
	PrivateSubnet *subnets.Subnet
	PublicSubnet  *subnets.Subnet
	Igw           *subnets.IGW
}

// ElementType implements pulumi.Input.
func (*Infrastructure) ElementType() reflect.Type {
	panic("unimplemented")
}

func CreateInfrastructure(ctx *pulumi.Context) (*Infrastructure, error) {

	config, err := config.GetConfig(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", config.Vpc.CidrBlock)

	vpcs, err := vpc.CreateVpc(ctx, "training-vpc", "10.0.0.0/16")
	if err != nil {
		return nil, err
	}

	privateSubnet, err := subnets.CreatePrivateSubnets(ctx, "10.0.0.0/20", vpcs.Id.ToStringOutput())
	if err != nil {
		return nil, err
	}

	publicSubnet, err := subnets.CreatePublicSubnets(ctx, "10.0.16.0/20", vpcs.Id.ToStringOutput())
	if err != nil {
		return nil, err
	}

	igw, err := subnets.CreateIGW(ctx, vpcs.Id.ToStringOutput())
	if err != nil {
		return nil, err
	}

	return &Infrastructure{
		Vpc:           vpcs,
		PrivateSubnet: privateSubnet,
		PublicSubnet:  publicSubnet,
		Config:        &config,
		Igw:           &igw,
	}, nil
}

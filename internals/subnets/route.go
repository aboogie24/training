package subnets

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"training.alfredbrowniii.io/internals/infrastructure"
)

func RouteAssociation(ctx *pulumi.Context, infra &infrastructure.Infrastructure) {

	route_id = GetRouteTable(ctx, infra.PublicSubnet.ID.ToStringOutput)

	args := []string{
		SubnetId:     infra.PublicSubnet.ID,
		RouteTableId: route_id,
	}

	_, err := ec2.NewRouteTableAssociation(ctx, "route-1", args)
	if err != nil {
		return
	}
}

func GetRouteTable(ctx, subnet_id pulumi.StringOutput) (route_id string) {

	selected, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
		SubnetId: pulumi.StringPtr(subnet_id),
	}, nil)
	if err != nil {
		return nil
	}

	return selected.Id

}

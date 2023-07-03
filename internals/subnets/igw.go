package subnets

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IGW struct {
	Id  pulumi.StringOutput
	Arn pulumi.StringOutput
}

func CreateIGW(ctx *pulumi.Context, vpc_id pulumi.StringOutput) (gateway IGW, err error) {

	ig, err := ec2.NewInternetGateway(ctx, "gw", &ec2.InternetGatewayArgs{
		VpcId: vpc_id,
		Tags: pulumi.StringMap{
			"Name": pulumi.String("main"),
		},
	})
	if err != nil {
		return IGW{}, err
	}

	_, err = ec2.NewInternetGatewayAttachment(ctx, "gw-attachedment", &ec2.InternetGatewayAttachmentArgs{
		InternetGatewayId: ig.ID(),
		VpcId:             vpc_id,
	})
	if err != nil {
		return IGW{}, err
	}

	ig_id := ig.ID().ApplyT(func(id string) string {
		return id
	}).(pulumi.StringOutput)

	igw := IGW{
		Id:  ig_id,
		Arn: ig.Arn,
	}

	return igw, nil

}

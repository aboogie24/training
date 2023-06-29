package vpc

import (
	"log"

	nativeEc2 "github.com/pulumi/pulumi-aws-native/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	cidrBlock string
	tags      []map[string]string
	Id        pulumi.StringOutput
}

func CreateVpc(ctx *pulumi.Context, vpcName string, cidrBlock string) (*Vpc, error) {

	if vpcName == "" {
		log.Panic("No vpc name provided")
	}

	vpc, err := nativeEc2.NewVPC(ctx, vpcName, &nativeEc2.VPCArgs{
		CidrBlock:          pulumi.StringPtr(cidrBlock),
		EnableDnsHostnames: pulumi.BoolPtr(true),
		EnableDnsSupport:   pulumi.BoolPtr(true),
		InstanceTenancy:    nil,
		Ipv4IpamPoolId:     nil,
		Ipv4NetmaskLength:  nil,
		Tags:               nil,
	})
	if err != err {
		log.Default()
		return nil, err
	}
	vpc_id := vpc.ID().ApplyT(func(id string) string {
		return id
	}).(pulumi.StringOutput)

	v := &Vpc{
		cidrBlock: cidrBlock,
		tags:      []map[string]string{},
		Id:        vpc_id,
	}

	return v, nil

}

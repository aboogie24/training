package vpc

import (
	"log"

	nativeEc2 "github.com/pulumi/pulumi-aws-native/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	cidrBlock string
	tags      []map[string]string
}

func CreateVpc(ctx *pulumi.Context, vpcName string, cidrBlock string) (*Vpc, error) {

	if vpcName == "" {
		log.Panic("No vpc name provided")
	}

	v := &Vpc{
		cidrBlock: cidrBlock,
		tags:      []map[string]string{},
	}

	_, err := nativeEc2.NewVPC(ctx, vpcName, &nativeEc2.VPCArgs{
		CidrBlock:          pulumi.StringPtr(v.cidrBlock),
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

	return v, nil
}

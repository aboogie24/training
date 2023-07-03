package subnets

import (
	"log"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subnet struct {
	Cidr string
}

func CreatePrivateSubnets(ctx *pulumi.Context, cidr string, id pulumi.StringOutput) (*Subnet, error) {

	// if cidr is  empty return  err
	if cidr == "" {
		log.Panic("Cidr not provided  for subnet")
	}

	s := &Subnet{
		Cidr: cidr,
	}

	_, err := ec2.NewSubnet(ctx, "Private-subnet", &ec2.SubnetArgs{
		AssignIpv6AddressOnCreation:   nil,
		AvailabilityZone:              nil,
		AvailabilityZoneId:            nil,
		CidrBlock:                     pulumi.StringPtr(s.Cidr),
		EnableDns64:                   nil,
		Ipv6CidrBlock:                 nil,
		Ipv6Native:                    nil,
		MapPublicIpOnLaunch:           nil,
		OutpostArn:                    nil,
		PrivateDnsNameOptionsOnLaunch: nil,
		Tags:                          nil,
		VpcId:                         id,
	}, nil)

	if err != nil {
		return nil, err
	}

	return s, err
}

func CreatePublicSubnets(ctx *pulumi.Context, cidr string, id pulumi.StringOutput) (*Subnet, error) {

	// if cidr is  empty return  err
	if cidr == "" {
		log.Panic("Cidr not provided  for subnet")
	}

	s := &Subnet{
		Cidr: cidr,
	}

	_, err := ec2.NewSubnet(ctx, "public-subnet", &ec2.SubnetArgs{
		AssignIpv6AddressOnCreation:   nil,
		AvailabilityZone:              nil,
		AvailabilityZoneId:            nil,
		CidrBlock:                     pulumi.StringPtr(s.Cidr),
		EnableDns64:                   nil,
		Ipv6CidrBlock:                 nil,
		Ipv6Native:                    nil,
		MapPublicIpOnLaunch:           nil,
		OutpostArn:                    nil,
		PrivateDnsNameOptionsOnLaunch: nil,
		Tags:                          nil,
		VpcId:                         id,
	}, nil)

	if err != nil {
		return nil, err
	}

	return s, err
}

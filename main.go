package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Infrastructure struct {
	tags          Tags
	vpc           VPC
	privateSubnet Subnet
	publicSubnet  Subnet
}

type Tags struct {
	Version  string
	Author   string
	Team     string
	Features string
	Region   string
}

type VPC struct {
	ID        pulumi.String
	cidrBlock string
	name      string
}

type Subnet struct {
	CidrBlock string
}

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		tags := &Tags{
			Version:  getEnv(ctx, "tags:version", "unknown"),
			Author:   getEnv(ctx, "tags:author", "unknown"),
			Team:     getEnv(ctx, "tags:team", "unknown"),
			Features: getEnv(ctx, "tags:feature", "unknown"),
			Region:   getEnv(ctx, "aws-native:region", "unknown"),
		}
		// Unmarshal struct strings to map
		t, _ := json.Marshal(&tags)
		var m map[string]string
		_ = json.Unmarshal(t, &m)

		vpc := &VPC{
			cidrBlock: getEnv(ctx, "vpc:cidrBlock", "unknown"),
			name:      getEnv(ctx, "vpc:name", "vpc-01"),
		}

		infra := &Infrastructure{
			tags: *tags,
			vpc:  *vpc,
		}

		vpcArgs := &ec2.VpcArgs{
			CidrBlock: pulumi.String(infra.vpc.cidrBlock),
			Tags:      pulumi.ToStringMap(m),
		}

		vpcs, err := ec2.NewVpc(ctx, infra.vpc.name, vpcArgs)
		if err != nil {
			panic("Vpc not created")
		}

		pub_sub := &Subnet{
			CidrBlock: getEnv(ctx, "vpc:public-subnet-ips", "10.0.16.0/20"),
		}

		public_subnetArgs := &ec2.SubnetArgs{
			AssignIpv6AddressOnCreation:             nil,
			AvailabilityZone:                        pulumi.String("us-east-1a"),
			AvailabilityZoneId:                      nil,
			CidrBlock:                               pulumi.StringPtr(pub_sub.CidrBlock),
			CustomerOwnedIpv4Pool:                   nil,
			EnableDns64:                             nil,
			EnableResourceNameDnsARecordOnLaunch:    nil,
			EnableResourceNameDnsAaaaRecordOnLaunch: nil,
			Ipv6CidrBlock:                           nil,
			Ipv6Native:                              nil,
			MapCustomerOwnedIpOnLaunch:              nil,
			MapPublicIpOnLaunch:                     nil,
			OutpostArn:                              nil,
			PrivateDnsHostnameTypeOnLaunch:          nil,
			Tags:                                    pulumi.ToStringMap(m),
			VpcId:                                   vpcs.ID(),
		}

		publicSubnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("%s-pub-subnet", vpc.name), public_subnetArgs)
		if err != nil {
			return err
		}

		priv_sub := &Subnet{
			CidrBlock: getEnv(ctx, "vpc:private-subnet-ips", "10.0.0.0/20"),
		}

		priv_subnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("%s-priv-subnet", vpc.name), &ec2.SubnetArgs{
			AssignIpv6AddressOnCreation:             nil,
			AvailabilityZone:                        pulumi.String("us-east-1b"),
			AvailabilityZoneId:                      nil,
			CidrBlock:                               pulumi.StringPtr(priv_sub.CidrBlock),
			CustomerOwnedIpv4Pool:                   nil,
			EnableDns64:                             nil,
			EnableResourceNameDnsARecordOnLaunch:    nil,
			EnableResourceNameDnsAaaaRecordOnLaunch: nil,
			Ipv6CidrBlock:                           nil,
			Ipv6Native:                              nil,
			MapCustomerOwnedIpOnLaunch:              nil,
			MapPublicIpOnLaunch:                     nil,
			OutpostArn:                              nil,
			PrivateDnsHostnameTypeOnLaunch:          nil,
			Tags:                                    pulumi.ToStringMap(m),
			VpcId:                                   vpcs.ID(),
		})
		if err != nil {
			return err
		}

		igw, err := ec2.NewInternetGateway(ctx, "gw", &ec2.InternetGatewayArgs{
			VpcId: vpcs.ID(),
			Tags:  pulumi.ToStringMap(m),
		})
		if err != nil {
			return err
		}

		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"eks.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		})

		checkErr(err)

		eksRole, err := iam.NewRole(ctx, "eks-training", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(assumeRole.Json),
		})
		checkErr(err)

		_, err = iam.NewRolePolicyAttachment(ctx, "EKSClusterPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
			Role:      eksRole.Name,
		})
		checkErr(err)
		_, err = iam.NewRolePolicyAttachment(ctx, "AmazonEKSVPCResourceController", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"),
			Role:      eksRole.Name,
		})
		checkErr(err)

		ClusterEKS, err := eks.NewCluster(ctx, "Training-eks", &eks.ClusterArgs{
			DefaultAddonsToRemoves:  nil,
			EnabledClusterLogTypes:  nil,
			EncryptionConfig:        nil,
			KubernetesNetworkConfig: nil,
			Name:                    pulumi.StringPtr("Training-eks"),
			OutpostConfig:           nil,
			RoleArn:                 eksRole.Arn,
			Tags:                    nil,
			Version:                 nil,
			VpcConfig: &eks.ClusterVpcConfigArgs{
				ClusterSecurityGroupId: nil,
				EndpointPrivateAccess:  nil,
				EndpointPublicAccess:   nil,
				PublicAccessCidrs:      nil,
				SecurityGroupIds:       nil,
				SubnetIds:              pulumi.StringArray{publicSubnet.ID(), priv_subnet.ID()},
				VpcId:                  nil,
			},
		})
		checkErr(err)

		ctx.Export("vpc_ID:", vpcs.ID())
		ctx.Export("public-subnet:", publicSubnet.ID())
		ctx.Export("private-subnet:", priv_subnet.ID())
		ctx.Export("IGW-ID:", igw.ID())
		ctx.Export("ClusterID:", ClusterEKS.ID())
		ctx.Export("K8s Endpoint:", ClusterEKS.Endpoint)

		return nil
	})
}

func getEnv(ctx *pulumi.Context, key string, fallback string) string {
	if value, ok := ctx.GetConfig(key); ok {
		return value
	}
	return fallback
}

func checkErr(err error) {
	if err != nil {
		fmt.Print("you fucked up")
		return
	}
	return
}

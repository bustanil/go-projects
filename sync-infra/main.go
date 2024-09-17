package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		amazonLinux, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"amazon",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "architecture",
					Values: []string{
						"arm64",
					},
				},
				{
					Name: "name",
					Values: []string{
						"al2023-ami-2023*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}

		vpc, err := ec2.NewVpc(ctx, "sync-vpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			EnableDnsSupport:   pulumi.Bool(true),
		})

		igw, err := ec2.NewInternetGateway(ctx, "sync-igw", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}

		// Create a route table
		rt, err := ec2.NewRouteTable(ctx, "igw-rt", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: igw.ID(),
				},
			},
		})
		if err != nil {
			return err
		}

		// Create two public subnets
		subnet1, err := ec2.NewSubnet(ctx, "my-subnet-1", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.1.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		subnet2, err := ec2.NewSubnet(ctx, "my-subnet-2", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.2.0/24"),
			MapPublicIpOnLaunch: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Associate the route table with the subnets
		_, err = ec2.NewRouteTableAssociation(ctx, "subnet1-association", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet1.ID(),
			RouteTableId: rt.ID(),
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewRouteTableAssociation(ctx, "subnet2-association", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet2.ID(),
			RouteTableId: rt.ID(),
		})
		if err != nil {
			return err
		}

		sg, err := ec2.NewSecurityGroup(ctx, "http-ssh-sg", &ec2.SecurityGroupArgs{
			VpcId:       vpc.ID(),
			Description: pulumi.String("Allow HTTP and SSH traffic"),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(80),
					ToPort:     pulumi.Int(80),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
				&ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(22),
					ToPort:     pulumi.Int(22),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
		})
		if err != nil {
			return err
		}

		bucket, err := s3.NewBucket(ctx, "sync-bucket", &s3.BucketArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("sync-bucket"),
			},
		})
		if err != nil {
			return err
		}

		// Create an IAM role
		role, err := iam.NewRole(ctx, "ec2Role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Action": "sts:AssumeRole",
					"Principal": {
						"Service": "ec2.amazonaws.com"
					},
					"Effect": "Allow",
					"Sid": ""
				}]
			}`),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("ec2Role"),
			},
		})
		if err != nil {
			return err
		}

		// Attach the S3 read/write policy to the IAM role
		policy := pulumi.Sprintf(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Action": [
						"s3:ListBucket",
						"s3:PutObject",
						"s3:GetObject",
						"s3:DeleteObject"
					],
					"Resource": [
						"%s"
					]
				}]
			}`, bucket.Arn)
		_, err = iam.NewRolePolicy(ctx, "S3-access-role", &iam.RolePolicyArgs{
			Role:   role.Name,
			Policy: policy,
		})
		if err != nil {
			return err
		}

		// Create an IAM instance profile for the role
		profile, err := iam.NewInstanceProfile(ctx, "sync-S3-access", &iam.InstanceProfileArgs{
			Role: role.Name,
		})
		if err != nil {
			return err
		}

		// file API EC2 instance
		_, err = ec2.NewInstance(ctx, "file-api", &ec2.InstanceArgs{
			Ami: pulumi.String(amazonLinux.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("file-api"),
			},
			InstanceType:             ec2.InstanceType_T4g_Micro,
			VpcSecurityGroupIds:      pulumi.StringArray{sg.ID()},
			IamInstanceProfile:       profile,
			AssociatePublicIpAddress: pulumi.Bool(true),
			SubnetId:                 subnet1.ID(),
		})
		if err != nil {
			return err
		}

		// TODO outputs
		return nil
	})
}

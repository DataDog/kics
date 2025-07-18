---
title: "Public and Private EC2 Share Role"
group-id: "rules/terraform/aws"
meta:
  name: "aws/public_and_private_ec2_share_role"
  id: "c53c7a89-f9d7-4c7b-8b66-8a555be99593"
  display_name: "Public and Private EC2 Share Role"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `c53c7a89-f9d7-4c7b-8b66-8a555be99593`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#iam_instance_profile)

### Description

 Public and private EC2 instances should not share the same IAM role. Assigning the same IAM role to both public and private instances greatly increases the blast radius in the event of a compromise. If an attacker gains control over a public EC2 instance, they could leverage the shared role’s permissions to access sensitive AWS resources with the same privileges as a private, internal instance, possibly leading to data breaches and lateral movement within your AWS environment. Isolating IAM roles for public and private instances helps minimize risk by ensuring that externally exposed instances have only the minimal permissions required.


## Compliant Code Examples
```terraform
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name            = "Ec2RoleShareRule1"
  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  cidr                          = "10.0.0.0/16"
  manage_default_security_group = true
  default_security_group_ingress = [
    {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "ssh"
      cidr_blocks = "0.0.0.0/0"
  }]
  default_security_group_egress = []
  version                       = "3.7.0"
}

module "ec2_public_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = module.vpc.public_subnets[0]
  iam_instance_profile   = aws_iam_instance_profile.test_profile5.name

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

module "ec2_private_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = module.vpc.private_subnets[0]
  iam_instance_profile   = aws_iam_instance_profile.test_profile4.name

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

resource "aws_iam_instance_profile" "test_profile4" {
  name = "test_profile"
  role = "aws_iam_role.test_role4.name"
}

resource "aws_iam_instance_profile" "test_profile5" {
  name = "test_profile"
  role = "aws_iam_role.test_role5.name"
}

```

```terraform
provider "aws" {
  region = "us-east-1"
}

locals {
  test_description = "two EC2s, one public, one private"
  test_name        = "Ec2RoleShareRule test - use case 1"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"

  name = "Ec2RoleShareRule1"
  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  cidr = "10.0.0.0/16"
  manage_default_security_group= true
  default_security_group_ingress = [
          {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "ssh"
      cidr_blocks = "0.0.0.0/0"
    }]
  default_security_group_egress =[]
  version = "3.7.0"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_iam_role" "test_role2" {
  name = "test_role2"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

resource "aws_iam_instance_profile" "test_profile2" {
  name = "test_profile"
  role = "aws_iam_role.test_role2.name"
}


resource "aws_iam_role_policy" "test_policy2" {
  name = "test_policy"
  role = "aws_iam_role.test_role2.id"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}


resource "aws_instance" "pub_ins2" {
  ami           = "data.aws_ami.ubuntu.id"
  instance_type = "t2.micro"
  subnet_id = module.vpc.public_subnets[0]
  iam_instance_profile = aws_iam_instance_profile.test_profile2.name

}

resource "aws_instance" "priv_ins2" {
  ami           = "data.aws_ami.ubuntu.id"
  instance_type = "t2.micro"
  subnet_id = module.vpc.private_subnets[0]
  iam_instance_profile = aws_iam_instance_profile.test_profile3.name
}

```
## Non-Compliant Code Examples
```terraform
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name            = "Ec2RoleShareRule1"
  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  cidr                          = "10.0.0.0/16"
  manage_default_security_group = true
  default_security_group_ingress = [
    {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "ssh"
      cidr_blocks = "0.0.0.0/0"
  }]
  default_security_group_egress = []
  version                       = "3.7.0"
}

module "ec2_public_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = module.vpc.public_subnets[0]
  iam_instance_profile   = aws_iam_instance_profile.test_profile1.name

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

module "ec2_private_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = module.vpc.private_subnets[0]
  iam_instance_profile   = aws_iam_instance_profile.test_profile1.name

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

resource "aws_iam_instance_profile" "test_profile1" {
  name = "test_profile"
  role = "aws_iam_role.test_role1.name"
}

```

```terraform
provider "aws" {
  region = "us-east-1"
}

locals {
  test_description = "two EC2s, one public, one private"
  test_name        = "Ec2RoleShareRule test - use case 1"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"

  name = "Ec2RoleShareRule1"
  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  cidr = "10.0.0.0/16"
  manage_default_security_group= true
  default_security_group_ingress = [
          {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "ssh"
      cidr_blocks = "0.0.0.0/0"
    }]
  default_security_group_egress =[]
  version = "3.7.0"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_iam_role" "test_role" {
  name = "test_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

resource "aws_iam_instance_profile" "test_profile" {
  name = "test_profile"
  role = "aws_iam_role.test_role.name"
}


resource "aws_iam_role_policy" "test_policy" {
  name = "test_policy"
  role = "aws_iam_role.test_role.id"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}


resource "aws_instance" "pub_ins" {
  ami           = "data.aws_ami.ubuntu.id"
  instance_type = "t2.micro"
  subnet_id = module.vpc.public_subnets[0]
  iam_instance_profile = aws_iam_instance_profile.test_profile.name

}

resource "aws_instance" "priv_ins" {
  ami           = "data.aws_ami.ubuntu.id"
  instance_type = "t2.micro"
  subnet_id = module.vpc.private_subnets[0]
  iam_instance_profile = aws_iam_instance_profile.test_profile.name
}

```
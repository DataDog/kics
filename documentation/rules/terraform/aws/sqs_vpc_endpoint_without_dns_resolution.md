---
title: "SQS VPC Endpoint Without DNS Resolution"
group-id: "rules/terraform/aws"
meta:
  name: "aws/sqs_vpc_endpoint_without_dns_resolution"
  id: "e9b7acf9-9ba0-4837-a744-31e7df1e434d"
  display_name: "SQS VPC Endpoint Without DNS Resolution"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `e9b7acf9-9ba0-4837-a744-31e7df1e434d`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Low

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc#enable_dns_support)

### Description

 When configuring an Amazon SQS VPC Endpoint using Terraform, it is essential to have DNS resolution enabled for the associated VPC by setting `enable_dns_support = true`. If this is not set, such as in the configuration `enable_dns_support = false`, private DNS hostnames are not resolved within the VPC, causing communication failures for instances relying on DNS-based access to AWS services via the endpoint. This can result in application connectivity issues and may force the use of less secure public network routes, increasing the attack surface.


## Compliant Code Examples
```terraform
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"
  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true
  enable_dns_support = true

  tags = {
    Terraform = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_vpc" "main2" {
  cidr_block = local.cidr_block
  enable_dns_support = true
  enable_dns_hostnames = false
}

resource "aws_vpc_endpoint" "sqs-vpc-endpoint2" {
  vpc_id            = aws_vpc.main2.id
  service_name      = "com.amazonaws.${local.region}.sqs"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true
  subnet_ids = [aws_subnet.public-subnet.id]
  security_group_ids = [aws_security_group.public-internet-sg.id]
}

```
## Non-Compliant Code Examples
```terraform
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"
  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true
  enable_dns_support = false

  tags = {
    Terraform = "true"
    Environment = "dev"
  }
}

```

```terraform
locals {
  region = "us-east-1"
  cidr_block = "172.16.0.0/16"
  public_subnet_cidr_block = "172.16.100.0/24"
  quad_zero_cidr_block = "0.0.0.0/0"
}

provider "aws" {
  region = local.region
}

resource "aws_vpc" "main" {
  cidr_block = local.cidr_block
  enable_dns_support = false
  enable_dns_hostnames = false
}

resource "aws_subnet" "public-subnet" {
  vpc_id     = aws_vpc.main.id
  cidr_block = local.public_subnet_cidr_block

  tags = {
    Name = "public-subnet"
  }
}

resource "aws_route_table" "public-rtb" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = local.cidr_block
    vpc_endpoint_id = aws_vpc_endpoint.sqs-vpc-endpoint.id
  }

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "public-rtb"
  }
}

resource "aws_route_table_association" "public-rtb-assoc" {
  subnet_id      = aws_subnet.public-subnet.id
  route_table_id = aws_route_table.public-rtb.id
}

resource "aws_security_group" "public-internet-sg" {
  name        = "public-internet-sg"
  description = "Allow all local traffic with internet access"
  vpc_id      = aws_vpc.main.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [local.quad_zero_cidr_block]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [local.cidr_block]
  }

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

resource "aws_instance" "test-ec2-instance" {
  ami = data.aws_ami.ubuntu.id
  instance_type = "t2.micro"
  subnet_id = aws_subnet.public-subnet.id
  vpc_security_group_ids = [aws_security_group.public-internet-sg.id]
}

resource "aws_vpc_endpoint" "sqs-vpc-endpoint" {
  vpc_id            = aws_vpc.main.id
  service_name      = "com.amazonaws.${local.region}.sqs"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true
  subnet_ids = [aws_subnet.public-subnet.id]
  security_group_ids = [aws_security_group.public-internet-sg.id]
}

resource "aws_sqs_queue" "test-queue" {
  name = "test-queue"
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id
}

```
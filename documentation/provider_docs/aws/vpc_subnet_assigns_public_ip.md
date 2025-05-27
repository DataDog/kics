---
title: "VPC Subnet Assigns Public IP"
meta:
  name: "aws/vpc_subnet_assigns_public_ip"
  id: "52f04a44-6bfa-4c41-b1d3-4ae99a2de05c"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/vpc_subnet_assigns_public_ip`

**Id:** `52f04a44-6bfa-4c41-b1d3-4ae99a2de05c`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
VPC Subnet should not assign public IP

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet#map_public_ip_on_launch)

## Non-Compliant Code Examples
```terraform
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]

  map_public_ip_on_launch = true
  enable_nat_gateway      = true
  enable_vpn_gateway      = true

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "positive" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Positive"
  }

  map_public_ip_on_launch = true
}

```

## Compliant Code Examples
```terraform
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]

  map_public_ip_on_launch = false
  enable_nat_gateway      = true
  enable_vpn_gateway      = true

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_vpc" "main3" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "negative2" {
  vpc_id     = aws_vpc.main3.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Negative2"
  }

  map_public_ip_on_launch = false
}

```

```terraform
resource "aws_vpc" "main2" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "negative1" {
  vpc_id     = aws_vpc.main2.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Negative1"
  }
}

```
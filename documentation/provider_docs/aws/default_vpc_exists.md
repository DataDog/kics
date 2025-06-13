---
title: "Default VPC Exists"
meta:
  name: "aws/default_vpc_exists"
  id: "96ed3526-0179-4c73-b1b2-372fde2e0d13"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---

## Metadata
**Name:** `aws/default_vpc_exists`

**Id:** `96ed3526-0179-4c73-b1b2-372fde2e0d13`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Defaults

## Description
It isn't recommended to use resources in default VPC

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_vpc)

## Non-Compliant Code Examples
```terraform
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.7.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true
  default_vpc_name   = "my-default-vpc"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_default_vpc" "positive1" {
  tags = {
    Name = "Default VPC"
  }
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
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = true

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_vpc" "negative1" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"

  tags = {
    Name = "main"
  }
}



```
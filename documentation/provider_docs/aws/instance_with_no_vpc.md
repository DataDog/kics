---
title: "Instance With No VPC"
meta:
  name: "aws/instance_with_no_vpc"
  id: "a31a5a29-718a-4ff4-8001-a69e5e4d029e"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/instance_with_no_vpc`

**Id:** `a31a5a29-718a-4ff4-8001-a69e5e4d029e`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Insecure Configurations

## Description
EC2 Instances should be configured under a VPC network. AWS VPCs provide the controls to facilitate a formal process for approving and testing all network connections and changes to the firewall and router configurations.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance)

## Non-Compliant Code Examples
```terraform
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  subnet_id              = "subnet-eddcdzz4"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_instance" "positive1" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

}
```

## Compliant Code Examples
```terraform
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = true
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = "subnet-eddcdzz4"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
resource "aws_instance" "negative1" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  vpc_security_group_ids = ["aws_security_group.instance.id"]

}

```
---
title: "EC2 Instance Using Default VPC"
meta:
  name: "aws/ec2_instance_using_default_vpc"
  id: "7e4a6e76-568d-43ef-8c4e-36dea481bff1"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/ec2_instance_using_default_vpc`

**Id:** `7e4a6e76-568d-43ef-8c4e-36dea481bff1`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
EC2 Instances should not be configured under a default VPC network

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#subnet_id)

## Non-Compliant Code Examples
```terraform
resource "aws_instance" "positive1" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  subnet_id   = aws_subnet.my_subnet.id

}

resource "aws_subnet" "my_subnet" {
  vpc_id     = aws_vpc.default.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Main"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_instance" "negative1" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  subnet_id   = aws_subnet.my_subnet2.id

}

resource "aws_subnet" "my_subnet2" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "Main"
  }
}

```
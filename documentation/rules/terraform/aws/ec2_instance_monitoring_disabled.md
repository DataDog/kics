---
title: "EC2 Instance Monitoring Disabled"
meta:
  name: "aws/ec2_instance_monitoring_disabled"
  id: "23b70e32-032e-4fa6-ba5c-82f56b9980e6"
  display_name: "EC2 Instance Monitoring Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Name:** `aws/ec2_instance_monitoring_disabled`

**Query Name** `EC2 Instance Monitoring Disabled`

**Id:** `23b70e32-032e-4fa6-ba5c-82f56b9980e6`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Observability

## Description
Enabling detailed monitoring on EC2 instances ensures that Amazon CloudWatch captures metrics every minute, rather than the default five-minute interval. This is achieved in Terraform by setting the `monitoring` attribute to `true` within the `aws_instance` resource. Without detailed monitoring (`monitoring = false` or omitted), operational visibility is significantly reduced, which can delay the detection of performance issues, outages, or security incidents. If this vulnerability is left unaddressed, administrators may miss critical events or be unable to react promptly to changes in resource behavior, potentially leading to service disruptions or undiagnosed anomalies.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#monitoring)


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
  associate_public_ip_address = false

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "monitoring_negative1" {
  ami           = data.aws_ami.ubuntu.id
  monitoring    = true
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }
}

```
## Non-Compliant Code Examples
```terraform
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "monitoring_positive2" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"
  monitoring    = false

  tags = {
    Name = "HelloWorld"
  }
}

```

```terraform
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = "subnet-eddcdzz4"
  associate_public_ip_address = false

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```

```terraform
module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "single-instance"

  ami                    = "ami-ebd02392"
  instance_type          = "t2.micro"
  key_name               = "user1"
  monitoring             = false
  vpc_security_group_ids = ["sg-12345678"]
  subnet_id              = "subnet-eddcdzz4"
  associate_public_ip_address = false

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}

```
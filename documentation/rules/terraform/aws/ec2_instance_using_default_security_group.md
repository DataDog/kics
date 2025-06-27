---
title: "EC2 Instance Using Default Security Group"
meta:
  name: "aws/ec2_instance_using_default_security_group"
  id: "f1adc521-f79a-4d71-b55b-a68294687432"
  display_name: "EC2 Instance Using Default Security Group"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Name:** `aws/ec2_instance_using_default_security_group`

**Query Name** `EC2 Instance Using Default Security Group`

**Id:** `f1adc521-f79a-4d71-b55b-a68294687432`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Access Control

## Description
Amazon EC2 instances should not use the default security group, as this group is open by default and permits unrestricted inbound and outbound traffic (`security_groups = [aws_security_group.default.id]`). Using the default security group increases the risk of unauthorized access, lateral movement, and exposure of sensitive workloads. To reduce this attack surface, instances should be assigned to tightly controlled, custom security groups with explicit ingress and egress rules.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#security_groups)


## Compliant Code Examples
```terraform
resource "aws_instance" "negative2" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.sg.id]
}

```

```terraform
resource "aws_instance" "negative1" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }

  security_groups = [aws_security_group.sg.id]
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_instance" "positive2" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.default.id]
}

```

```terraform
resource "aws_instance" "positive1" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }

  security_groups = [aws_security_group.default.id]
}

```
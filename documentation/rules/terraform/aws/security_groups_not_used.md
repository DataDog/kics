---
title: "Security group not used"
group_id: "rules/terraform/aws"
meta:
  name: "aws/security_groups_not_used"
  id: "4849211b-ac39-479e-ae78-5694d506cb24"
  display_name: "Security group not used"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "INFO"
  category: "Access Control"
---
## Metadata

**Id:** `4849211b-ac39-479e-ae78-5694d506cb24`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Info

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

### Description

 This check ensures that AWS load balancers are associated with appropriate security groups, which control network traffic to and from the resource. When the `security_groups` attribute is omitted from an `aws_lb` resource, as shown below, the load balancer may become exposed to unrestricted network access, increasing the risk of unauthorized access or attacks:

```
resource "aws_lb" "test" {
  name = "test"
  load_balancer_type = "application"
  subnets = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]
  internal = true
}
```

By explicitly defining `security_groups`, you can restrict inbound and outbound traffic to only trusted sources:

```
resource "aws_lb" "test" {
  name = "test"
  load_balancer_type = "application"
  subnets = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]
  internal = true
  security_groups = [aws_security_group.allow_tls.id]
}
```


## Compliant Code Examples
```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  required_version = ">= 1.1.0"
}

variable "iam_role" {
  type        = string
  default     = "AmazonSSMRoleForInstancesQuickSetup"
  description = "Set AWS IAM role."
}

variable "ami_owner" {
  type        = string
  default     = "self"
  description = "Set AWS image owner."
}

variable "region" {
  type        = string
  default     = "eu-west-3"
  description = "Set AWS region."
}

variable "secgroups" {
  type        = list(string)
  default     = ["CowrieSSH"]
  description = "Set AWS security groups."
}

data "aws_ami" "cowrie" {
  most_recent = true
  owners      = ["var.ami_owner"]

  filter {
    name   = "name"
    values = ["cowrie-packer-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

provider "aws" {
  profile = "default"
  region  = var.region
}

resource "aws_security_group" "cowrie" {
  name        = "CowrieSSH"
  description = "CowrieSSH Terraform security group"

  ingress {
    description = "Allow anyone to connect to the honeypot."
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description      = "Allow all outgoing traffic."
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name    = "cowrie_ssh_sg"
    purpose = "honeypot"
  }
}

resource "aws_instance" "cowrie_server" {
  ami                  = data.aws_ami.cowrie.id
  instance_type        = "t3.nano"
  security_groups      = var.secgroups
  iam_instance_profile = var.iam_role

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }

  tags = {
    Name    = "cowrie",
    author  = "foo"
    vcs-url = "https://github.com/foo/bar"
    purpose = "honeypot"
  }
}

```

```terraform
resource "aws_security_group" "allow_tls" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "TLS from VPC"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_tls"
  }
}

resource "aws_lb" "test" {
  name = "test"
  load_balancer_type = "application"
  subnets = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]
  internal = true
  security_groups = [aws_security_group.allow_tls.id]
}

```

```terraform
# given:
#  - used security group
#  - aws_instance
# when:
#  - used security group attached to aws_instance
# then:
#  - do not detect any unused security group

resource "aws_security_group" "used_sg" {
  name        = "used-sg"
  description = "Used security group"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "Some port"
    from_port        = 43
    to_port          = 43
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

}

resource "aws_instance" "negative3" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"

  vpc_security_group_ids = [ "aws_security_group.used_sg.id" ]

}


```
## Non-Compliant Code Examples
```terraform
# given:
#  - unused security group
#  - aws_instance
# when:
#  - no security group attached to aws_instance
# then:
#  - detect unused security group as unused

resource "aws_instance" "positive1" {
  ami = "ami-003634241a8fcdec0"

  instance_type = "t2.micro"
}

resource "aws_security_group" "unused-sg" {
  name        = "unused-sg"
  description = "Unused security group"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "Some port"
    from_port        = 42
    to_port          = 42
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

}

```

```terraform
resource "aws_security_group" "example" {
  name        = "example"
  description = "Allow Redis traffic"
  vpc_id      = data.aws_vpc.selected.id

  ingress {
    
    from_port   = 6379
    to_port     = 6379
    protocol    = "tcp"
    cidr_blocks = [data.aws_vpc.selected.cidr_block]
  }
}

resource "aws_elasticache_replication_group" "redis" {
  replication_group_id       = "Example"
  parameter_group_name       = "default.redis6.x"
  engine                     = "redis"
  engine_version             = "6.x"
  automatic_failover_enabled = false
  
}
```

```terraform
resource "aws_lb" "test" {
  name = "test"
  load_balancer_type = "application"
  subnets = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]
  internal = true
}

resource "aws_security_group" "allow_tls" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "TLS from VPC"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_tls"
  }
}

```
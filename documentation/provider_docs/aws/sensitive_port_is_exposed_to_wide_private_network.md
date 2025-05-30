---
title: "Sensitive Port Is Exposed To Wide Private Network"
meta:
  name: "aws/sensitive_port_is_exposed_to_wide_private_network"
  id: "92fe237e-074c-4262-81a4-2077acb928c1"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/sensitive_port_is_exposed_to_wide_private_network`

**Id:** `92fe237e-074c-4262-81a4-2077acb928c1`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
A sensitive port, such as port 23 or port 110, is open for a wide private network in either TCP or UDP protocol

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```terraform
resource "aws_security_group" "positive6" {
  name        = "allow_tls6"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 135
    to_port     = 170
    protocol    = "udp"
    cidr_blocks = ["10.68.0.0", "172.16.0.0/12"]
  }
}

```

```terraform
resource "aws_security_group" "positive2" {
  name        = "allow_tls2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 20
    to_port     = 60
    protocol    = "tcp"
    cidr_blocks = ["192.168.0.0/16"]
  }
}

```

```terraform
resource "aws_security_group" "positive3" {
  name        = "allow_tls3"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 5000
    to_port     = 6000
    protocol    = "-1"
    cidr_blocks = ["172.16.0.0/12"]
  }
}

```

```terraform
resource "aws_security_group" "positive7" {
  name        = "allow_tls7"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "udp"
    cidr_blocks = ["192.168.0.0/16", "10.0.0.0/8"]
  }
}

```

```terraform
resource "aws_security_group" "positive8" {
  name        = "allow_tls8"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["172.16.0.0/12"]
  }
}

```

```terraform
module "vote_service_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.3.0"

  name        = "user-service"
  description = "Security group for user-service with custom ports open within VPC, and PostgreSQL publicly open"
  vpc_id      = "vpc-12345678"

  ingress_with_cidr_blocks = [
    {
      description = "TLS from VPC"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_blocks = ["172.16.0.0/12"]
    }
  ]
}

```

```terraform
resource "aws_security_group" "positive4" {
  name        = "allow_tls4"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 20
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/8"]
  }
}

```

```terraform
resource "aws_security_group" "positive5" {
  name        = "allow_tls5"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 445
    to_port     = 500
    protocol    = "udp"
    cidr_blocks = ["192.168.0.0/16", "0.0.0.0/0", "2.2.3.4/12"]
  }
}

```

```terraform
resource "aws_security_group" "positive1" {
  name        = "allow_tls1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2200
    to_port     = 2500
    protocol    = "-1"
    cidr_blocks = ["10.0.0.0/8"]
  }
}

```

## Compliant Code Examples
```terraform
module "vote_service_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.3.0"

  name        = "user-service"
  description = "Security group for user-service with custom ports open within VPC, and PostgreSQL publicly open"
  vpc_id      = "vpc-12345678"

  ingress_with_cidr_blocks = [
    {
      description = "TLS from VPC"
      from_port   = 0
      to_port     = 0
      protocol    = "-1"
      cidr_blocks = ["1.2.3.4", "0.0.0.0/0"]
    }
  ]
}

```

```terraform
resource "aws_security_group" "negative3" {
  name        = "allow_tls3"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "tcp"
    cidr_blocks = ["1.2.3.4/0"]
  }
}

```

```terraform
resource "aws_security_group" "negative2" {
  name        = "allow_tls2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2384
    to_port     = 2386
    protocol    = "tcp"
    cidr_blocks = ["/0"]
  }
}

```

```terraform
resource "aws_security_group" "negative6" {
  name        = "allow_tls6"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["1.2.3.4", "0.0.0.0/0"]
  }
}

```

```terraform
resource "aws_security_group" "negative1" {
  name        = "allow_tls1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.main.cidr_block]
  }
}

```

```terraform
resource "aws_security_group" "negative5" {
  name        = "allow_tls5"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "udp"
    cidr_blocks = ["1.2.3.4/5", "0.0.0.0/12"]
  }
}

```

```terraform
resource "aws_security_group" "negative4" {
  name        = "allow_tls4"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 25
    to_port     = 2500
    protocol    = "tcp"
    cidr_blocks = ["1.2.3.4/5"]
  }
}

```
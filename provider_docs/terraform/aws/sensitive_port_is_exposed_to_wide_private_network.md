---
title: "Sensitive Port Is Exposed To Wide Private Network"
meta:
  name: "terraform/sensitive_port_is_exposed_to_wide_private_network"
  id: "92fe237e-074c-4262-81a4-2077acb928c1"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/sensitive_port_is_exposed_to_wide_private_network`
**Id:** `92fe237e-074c-4262-81a4-2077acb928c1`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Networking and Firewall
## Description
Leaving sensitive ports such as port 23 (Telnet) or port 110 (POP3) open to a wide private network via insecure security group rules can expose resources to unnecessary risk, as these ports are frequently targeted by attackers seeking to exploit legacy or weakly protected protocols. In Terraform, a misconfiguration like:

```
ingress {
  from_port   = 23
  to_port     = 23
  protocol    = "tcp"
  cidr_blocks = ["10.0.0.0/8"]
}
```

makes internal resources within the VPC accessible to all hosts in the private address range, greatly increasing the attack surface if any host in that range is compromised. Restricting access to only necessary subnets and ports—for example, using

```
ingress {
  from_port   = 2383
  to_port     = 2383
  protocol    = "tcp"
  cidr_blocks = [aws_vpc.main.cidr_block]
}
```

—significantly reduces the risk of lateral movement and unauthorized access within your network.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)

## Non-Compliant Code Examples
```aws
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

```aws
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

```aws
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

## Compliant Code Examples
```aws
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

```aws
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

```aws
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
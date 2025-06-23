---
title: "SQL Analysis Services Port 2383 (TCP) Is Publicly Accessible"
meta:
  name: "aws/sql_analysis_services_port_2383_is_publicly_accessible"
  id: "54c417bf-c762-48b9-9d31-b3d87047e3f0"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `aws/sql_analysis_services_port_2383_is_publicly_accessible`
**Id:** `54c417bf-c762-48b9-9d31-b3d87047e3f0`
**Cloud Provider:** aws
**Severity:** Medium
**Category:** Networking and Firewall
## Description
This check verifies whether TCP port 2383 is publicly accessible by inspecting the `cidr_blocks` attribute in security group ingress rules for overly broad IP ranges, such as `["0.0.0.0/0"]`. Allowing public access to this port exposes the associated resource to the internet, increasing the risk of unauthorized access or attacks. Restricting access to trusted networks is recommended to minimize the attack surface and protect sensitive services listening on that port.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)


## Compliant Code Examples
```terraform
resource "aws_security_group" "negative1" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2383
    to_port     = 2383
    protocol    = "tcp"
    cidr_blocks = ["0.1.0.0/0"]
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_security_group" "positive1" {
  name        = "allow_tls_1"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2300
    to_port     = 2400
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
resource "aws_security_group" "positive2" {
  name        = "allow_tls_2"
  description = "Allow TLS inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "TLS from VPC"
    from_port   = 2380
    to_port     = 2390
    protocol    = "tcp"
    cidr_blocks = ["0.1.0.0/0"]
  }

   ingress {
    description = "TLS from VPC"
    from_port   = 2350
    to_port     = 2384
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

```
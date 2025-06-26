---
title: "Redshift Using Default Port"
meta:
  name: "terraform/redshift_using_default_port"
  id: "41abc6cc-dde1-4217-83d3-fb5f0cc09d8f"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/redshift_using_default_port`
**Id:** `41abc6cc-dde1-4217-83d3-fb5f0cc09d8f`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Networking and Firewall
## Description
Amazon Redshift clusters listen on a default port (`5439`) unless otherwise specified using the `port` attribute in Terraform. Using the default port increases the risk of brute-force and automated attacks, as malicious actors frequently scan for commonly used service ports. Setting a custom port adds an obscurity layer by making the service less predictable to attackers, thereby reducing its exposure to opportunistic network attacks.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#port)

## Non-Compliant Code Examples
```aws
resource "aws_redshift_cluster" "positive1" {
  cluster_identifier    = "tf-redshift-cluster"
  database_name         = "mydb"
  master_username       = "foo"
  master_password       = "Mustbe8characters"
  node_type             = "dc1.large"
  cluster_type          = "single-node"
  publicly_accessible   = false
}

```

```aws
resource "aws_redshift_cluster" "positive2" {
  cluster_identifier    = "tf-redshift-cluster"
  database_name         = "mydb"
  master_username       = "foo"
  master_password       = "Mustbe8characters"
  node_type             = "dc1.large"
  cluster_type          = "single-node"
  publicly_accessible   = false
  port                  = 5439
}

```

## Compliant Code Examples
```aws
resource "aws_redshift_cluster" "negative1" {
  cluster_identifier    = "tf-redshift-cluster"
  database_name         = "mydb"
  master_username       = "foo"
  master_password       = "Mustbe8characters"
  node_type             = "dc1.large"
  cluster_type          = "single-node"
  publicly_accessible   = false
  port                  = 1150
}

```
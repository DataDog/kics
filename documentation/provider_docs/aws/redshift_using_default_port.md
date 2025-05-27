---
title: "Redshift Using Default Port"
meta:
  name: "aws/redshift_using_default_port"
  id: "41abc6cc-dde1-4217-83d3-fb5f0cc09d8f"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/redshift_using_default_port`

**Id:** `41abc6cc-dde1-4217-83d3-fb5f0cc09d8f`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
Redshift should not use the default port (5439) because an attacker can easily guess the port

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#port)

## Non-Compliant Code Examples
```terraform
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

```terraform
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

## Compliant Code Examples
```terraform
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
---
title: "Redshift Publicly Accessible"
meta:
  name: "aws/redshift_publicly_accessible"
  id: "af173fde-95ea-4584-b904-bb3923ac4bda"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/redshift_publicly_accessible`

**Id:** `af173fde-95ea-4584-b904-bb3923ac4bda`

**Cloud Provider:** aws

**Severity:** High

**Category:** Insecure Configurations

## Description
Amazon Redshift clusters with public accessibility enabled can be accessed from the internet, creating a significant security risk by potentially exposing sensitive data to unauthorized users. By default, the 'publicly_accessible' parameter is set to true if not explicitly defined, meaning your cluster is publicly accessible unless you specifically configure otherwise. To secure your Redshift cluster, always set 'publicly_accessible = false' as shown below:

```hcl
resource "aws_redshift_cluster" "secure_example" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  publicly_accessible = false
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster)

## Non-Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "positive1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}

resource "aws_redshift_cluster" "positive2" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  publicly_accessible = true
}
```

## Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "negative1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  publicly_accessible = false
}
```
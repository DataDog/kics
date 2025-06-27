---
title: "DAX Cluster Not Encrypted"
meta:
  name: "aws/dax_cluster_not_encrypted"
  id: "f11aec39-858f-4b6f-b946-0a1bf46c0c87"
  display_name: "DAX Cluster Not Encrypted"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Name:** `aws/dax_cluster_not_encrypted`

**Query Name** `DAX Cluster Not Encrypted`

**Id:** `f11aec39-858f-4b6f-b946-0a1bf46c0c87`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** High

**Category:** Encryption

## Description
This check verifies that AWS DAX (DynamoDB Accelerator) clusters have server-side encryption enabled to protect data at rest. Without encryption, sensitive data stored in DAX clusters could be exposed if unauthorized access to the storage media occurs, potentially leading to data breaches and compliance violations.

To secure a DAX cluster, you must include a 'server_side_encryption' block with 'enabled = true' as shown below:
```
resource "aws_dax_cluster" "secure_example" {
  cluster_name       = "cluster-example"
  // other configuration...
  
  server_side_encryption {
    enabled = true
  }
}
```
Insecure configurations either omit the server_side_encryption block entirely, include an empty block, or explicitly set 'enabled = false'.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/dax_cluster#enabled)


## Compliant Code Examples
```terraform
resource "aws_dax_cluster" "bar" {
  cluster_name       = "cluster-example"
  iam_role_arn       = data.aws_iam_role.example.arn
  node_type          = "dax.r4.large"
  replication_factor = 1

  server_side_encryption {
    enabled = true
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_dax_cluster" "bar_1" {
  cluster_name       = "cluster-example"
  iam_role_arn       = data.aws_iam_role.example.arn
  node_type          = "dax.r4.large"
  replication_factor = 1
}

resource "aws_dax_cluster" "bar_2" {
  cluster_name       = "cluster-example"
  iam_role_arn       = data.aws_iam_role.example.arn
  node_type          = "dax.r4.large"
  replication_factor = 1

  server_side_encryption {
  }
}

resource "aws_dax_cluster" "bar_3" {
  cluster_name       = "cluster-example"
  iam_role_arn       = data.aws_iam_role.example.arn
  node_type          = "dax.r4.large"
  replication_factor = 1

  server_side_encryption {
    enabled = false
  }
}

```
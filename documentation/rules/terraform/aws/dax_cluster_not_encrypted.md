---
title: "DAX cluster not encrypted"
group_id: "rules/terraform/aws"
meta:
  name: "aws/dax_cluster_not_encrypted"
  id: "f11aec39-858f-4b6f-b946-0a1bf46c0c87"
  display_name: "DAX cluster not encrypted"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `f11aec39-858f-4b6f-b946-0a1bf46c0c87`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/dax_cluster#enabled)

### Description

 This check verifies that AWS DynamoDB Accelerator (DAX) clusters have server-side encryption enabled to protect data at rest. Without encryption, sensitive data stored in DAX clusters could be exposed if unauthorized access to the storage media occurs, potentially leading to data breaches and compliance violations.

To secure a DAX cluster, you must include a `server_side_encryption` block with `enabled = true`, as shown below:
```
resource "aws_dax_cluster" "secure_example" {
  cluster_name       = "cluster-example"
  // other configuration...
  
  server_side_encryption {
    enabled = true
  }
}
```
Insecure configurations either omit the `server_side_encryption` block entirely, include an empty block, or explicitly set `enabled = false`.


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

```terraform
module "dax" {
  source = "terraform-aws-modules/dax/aws"
  version = "2.0.0"

  name                   = "cluster"
  server_side_encryption = { enabled = true }
  iam_role_arn           = "arn:aws:iam::aws_account_id:role/aws_service_role"
  description            = "DAX cluster"
  node_type              = "dax.r4.large"
  replication_factor     = 2

  parameter_group = {
    name        = "default.dax1.0"
    parameters  = [
      {
        name  = "query-ttl-millis"
        value = "100000"
      },
      {
        name  = "record-ttl-millis"
        value = "100000"
      }
    ]
  }

  tags = {
    Environment = "test"
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

```terraform
module "dax" {
  source = "terraform-aws-modules/dax/aws"
  version = "2.0.0"

  name                   = "cluster"
  server_side_encryption = {}
  iam_role_arn           = "arn:aws:iam::aws_account_id:role/aws_service_role"
  description            = "DAX cluster"
  node_type              = "dax.r4.large"
  replication_factor     = 2

  parameter_group = {
    name        = "default.dax1.0"
    parameters  = [
      {
        name  = "query-ttl-millis"
        value = "100000"
      },
      {
        name  = "record-ttl-millis"
        value = "100000"
      }
    ]
  }

  tags = {
    Environment = "test"
  }
}
```
---
title: "Redshift Cluster Logging Disabled"
meta:
  name: "aws/redshift_cluster_logging_disabled"
  id: "15ffbacc-fa42-4f6f-a57d-2feac7365caa"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `aws/redshift_cluster_logging_disabled`
**Id:** `15ffbacc-fa42-4f6f-a57d-2feac7365caa`
**Cloud Provider:** aws
**Severity:** Medium
**Category:** Observability
## Description
Amazon Redshift clusters should have logging enabled to ensure that audit and diagnostic data, such as query execution and connection attempts, are recorded and stored for security and compliance purposes. If the `logging` block is omitted, or the attribute `enable` is set to `false` in the Redshift cluster Terraform resource, as shown below,

```
resource "aws_redshift_cluster" "example" {
  // ... other configuration ...
  logging {
      enable = false
  }
}
```

critical events may go unrecorded, making it difficult to detect unauthorized access or investigate security incidents. Enabling logging with

```
logging {
    enable = true
    bucket_name = "nameOfAnExistingS3Bucket"
}
```

helps ensure that activity is continuously monitored and available for review in the case of an incident.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#enable)


## Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "negative1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  logging {
      enable = true
      bucket_name = "nameOfAnExistingS3Bucket"
  }
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "positive1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  logging {
      enable = false
  }
}

resource "aws_redshift_cluster" "positive2" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}
```
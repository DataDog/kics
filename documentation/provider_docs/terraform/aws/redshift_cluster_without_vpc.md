---
title: "Redshift Cluster Without VPC"
meta:
  name: "aws/redshift_cluster_without_vpc"
  id: "0a494a6a-ebe2-48a0-9d77-cf9d5125e1b3"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `aws/redshift_cluster_without_vpc`
**Id:** `0a494a6a-ebe2-48a0-9d77-cf9d5125e1b3`
**Cloud Provider:** aws
**Severity:** Low
**Category:** Insecure Configurations
## Description
Amazon Redshift clusters should be deployed within an Amazon VPC to ensure network isolation and control over access to the cluster. If the attributes `vpc_security_group_ids` and `cluster_subnet_group_name` are not specified, the cluster is created outside a VPC and could be exposed to the public internet, increasing the risk of unauthorized access and data breaches. A secure configuration includes these attributes as shown:

```
resource "aws_redshift_cluster" "secure_example" {
  // other attributes
  vpc_security_group_ids     = [aws_security_group.redshift.id]
  cluster_subnet_group_name  = aws_redshift_subnet_group.redshift_subnet_group.id
}
```


#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#vpc_security_group_ids)


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
  vpc_security_group_ids = [
    aws_security_group.redshift.id
  ]
  cluster_subnet_group_name = aws_redshift_subnet_group.redshift_subnet_group.id
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
      enable = true
      bucket_name = "nameOfAnExistingS3Bucket"
  }
}

```
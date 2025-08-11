---
title: "Redshift publicly accessible"
group_id: "rules/terraform/aws"
meta:
  name: "aws/redshift_publicly_accessible"
  id: "af173fde-95ea-4584-b904-bb3923ac4bda"
  display_name: "Redshift publicly accessible"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `af173fde-95ea-4584-b904-bb3923ac4bda`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster)

### Description

 Amazon Redshift clusters with public accessibility enabled can be accessed from the internet, creating a significant security risk by potentially exposing sensitive data to unauthorized users. By default, the `publicly_accessible` parameter is set to `true` if not explicitly defined, meaning your cluster is publicly accessible unless specifically configured otherwise. To secure your Redshift cluster, always set `publicly_accessible = false`, as shown below:

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

```terraform
module "redshift" {
  source  = "terraform-aws-modules/redshift/aws"
  version = "~> 3.0"

  cluster_identifier = "redshift-cluster-demo"
  database_name      = "demodb"
  master_username    = "adminuser"
  master_password    = "UserPass123"
  node_type          = "dc2.large"
  publicly_accessible = false

  cluster_type = "single-node"

  # enhanced_vpc_routing = true

  # This is to avoid "trying to delete resource repeatedly" issue with this resource
  skip_final_snapshot = true
  iam_roles           = []

  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  depends_on = [
    aws_iam_role_policy_attachment.redshift-access,
    aws_iam_role_policy_attachment.redshift-access2,
  ]
}
```
## Non-Compliant Code Examples
```terraform
module "redshift" {
  source  = "terraform-aws-modules/redshift/aws"
  version = "~> 3.0"

  cluster_identifier = "redshift-cluster-demo"
  database_name      = "demodb"
  master_username    = "adminuser"
  master_password    = "UserPass123"
  node_type          = "dc2.large"
  publicly_accessible = true

  cluster_type = "single-node"

  # enhanced_vpc_routing = true

  # This is to avoid "trying to delete resource repeatedly" issue with this resource
  skip_final_snapshot = true
  iam_roles           = []

  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  depends_on = [
    aws_iam_role_policy_attachment.redshift-access,
    aws_iam_role_policy_attachment.redshift-access2,
  ]
}
```

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
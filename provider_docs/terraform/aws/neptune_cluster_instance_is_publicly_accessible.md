---
title: "Neptune Cluster Instance is Publicly Accessible"
meta:
  name: "terraform/neptune_cluster_instance_is_publicly_accessible"
  id: "9ba198e0-fef4-464a-8a4d-75ea55300de7"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/neptune_cluster_instance_is_publicly_accessible`
**Id:** `9ba198e0-fef4-464a-8a4d-75ea55300de7`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
Amazon Neptune cluster instances should not be publicly accessible to minimize the risk of unauthorized access to sensitive graph data. When a Neptune instance is publicly accessible, it can be accessed directly from the internet, potentially exposing it to attacks and unauthorized access attempts. To properly secure Neptune instances, set the 'publicly_accessible' attribute to 'false' as shown in the following example: 
```
resource "aws_neptune_cluster_instance" "example" {
  // ... other configurations
  publicly_accessible = false
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/neptune_cluster_instance#publicly_accessible)

## Non-Compliant Code Examples
```aws
resource "aws_neptune_cluster_instance" "example" {
  count              = 2
  cluster_identifier = aws_neptune_cluster.default.id
  engine             = "neptune"
  instance_class     = "db.r4.large"
  apply_immediately  = true
  publicly_accessible = true
}

```

## Compliant Code Examples
```aws
resource "aws_neptune_cluster_instance" "negative" {
  count              = 2
  cluster_identifier = aws_neptune_cluster.default.id
  engine             = "neptune"
  instance_class     = "db.r4.large"
  apply_immediately  = true
  publicly_accessible = false
}

```
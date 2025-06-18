---
title: "ECR Image Tag Not Immutable"
meta:
  name: "terraform/ecr_image_tag_not_immutable"
  id: "d1846b12-20c5-4d45-8798-fc35b79268eb"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/ecr_image_tag_not_immutable`
**Id:** `d1846b12-20c5-4d45-8798-fc35b79268eb`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
To ensure the integrity of container images, Amazon Elastic Container Registry (ECR) repositories should have image tag immutability enabled by setting `image_tag_mutability` to `IMMUTABLE`. When image tags are set as mutable, existing image tags can be overwritten with new images, which may enable attackers or unauthorized users to replace trusted container images with malicious ones without changing the tag reference. This vulnerability can compromise the supply chain, leading to the deployment of untrusted or harmful workloads in production environments. Enforcing image tag immutability helps maintain a consistent and auditable history of deployed images, preventing accidental or intentional tampering of container tags.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository)

## Non-Compliant Code Examples
```aws
resource "aws_ecr_repository" "foo2" {
  name                 = "bar"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "foo3" {
  name                 = "bar"

  image_scanning_configuration {
    scan_on_push = true
  }
}

```

## Compliant Code Examples
```aws
resource "aws_ecr_repository" "foo" {
  name                 = "bar"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

```
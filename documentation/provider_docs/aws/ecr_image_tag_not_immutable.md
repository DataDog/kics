---
title: "ECR Image Tag Not Immutable"
meta:
  name: "aws/ecr_image_tag_not_immutable"
  id: "d1846b12-20c5-4d45-8798-fc35b79268eb"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/ecr_image_tag_not_immutable`

**Id:** `d1846b12-20c5-4d45-8798-fc35b79268eb`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
ECR should have an image tag be immutable. This prevents image tags from being overwritten.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository)

## Non-Compliant Code Examples
```terraform
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
```terraform
resource "aws_ecr_repository" "foo" {
  name                 = "bar"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

```
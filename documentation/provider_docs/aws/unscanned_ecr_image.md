---
title: "Unscanned ECR Image"
meta:
  name: "aws/unscanned_ecr_image"
  id: "9630336b-3fed-4096-8173-b9afdfe346a7"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/unscanned_ecr_image`

**Id:** `9630336b-3fed-4096-8173-b9afdfe346a7`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
Checks if the ECR Image has been scanned

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository#scan_on_push)

## Non-Compliant Code Examples
```terraform
resource "aws_ecr_repository" "positive1" {
  name                 = "img_p_2"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_repository" "positive2" {
  name                 = "img_p_1"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}
```

## Compliant Code Examples
```terraform
resource "aws_ecr_repository" "negative1" {
  name                 = "bar"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
```
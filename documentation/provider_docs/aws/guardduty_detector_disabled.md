---
title: "GuardDuty Detector Disabled"
meta:
  name: "aws/guardduty_detector_disabled"
  id: "704dadd3-54fc-48ac-b6a0-02f170011473"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/guardduty_detector_disabled`

**Id:** `704dadd3-54fc-48ac-b6a0-02f170011473`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Make sure that Amazon GuardDuty is Enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/guardduty_detector#example-usage)

## Non-Compliant Code Examples
```terraform
resource "aws_guardduty_detector" "positive1" {
  enable = false
}


```

## Compliant Code Examples
```terraform
resource "aws_guardduty_detector" "negative1" {
  enable = true
}

```
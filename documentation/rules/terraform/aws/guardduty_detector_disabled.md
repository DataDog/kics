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
This check ensures that Amazon GuardDuty is enabled in your AWS environment by verifying that the `enable` attribute in the `aws_guardduty_detector` Terraform resource is set to `true`. GuardDuty is a threat detection service that continuously monitors for malicious or unauthorized behavior, helping identify and prioritize potential security risks. If GuardDuty is disabled, suspicious activities such as anomalous API calls, potentially unauthorized deployments, or account compromise may go undetected, leaving cloud resources vulnerable to attack. Enabling GuardDuty is a crucial security best practice to maintain visibility into potential threats and respond to incidents promptly.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/guardduty_detector#example-usage)


## Compliant Code Examples
```terraform
resource "aws_guardduty_detector" "negative1" {
  enable = true
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_guardduty_detector" "positive1" {
  enable = false
}


```
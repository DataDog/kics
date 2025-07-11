---
title: "CloudWatch AWS Organizations Changes Missing Alarm"
group-id: "rules/terraform/aws"
meta:
  name: "aws/cloudwatch_aws_organizations_changes_missing_alarm"
  id: "38b85c45-e772-4de8-a247-69619ca137b3"
  display_name: "CloudWatch AWS Organizations Changes Missing Alarm"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "INFO"
  category: "Observability"
---
## Metadata

**Id:** `38b85c45-e772-4de8-a247-69619ca137b3`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Info

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_metric_filter#pattern)

### Description

 This check ensures that a log metric filter and alarm are configured to monitor changes to AWS Organizations within CloudWatch. Without a properly configured metric filter and alarm, critical events such as the creation, deletion, or modification of organizational accounts and policies may go undetected. This lack of visibility increases the risk of malicious or unauthorized changes going unnoticed, which could compromise the security and governance of the entire AWS environment. Left unaddressed, this misconfiguration may allow threat actors to make unauthorized changes to the organization structure, potentially leading to privilege escalation or data exfiltration.


## Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-2"
}

resource "aws_cloudwatch_log_group" "CIS_CloudWatch_LogsGroup" {
  name = "CIS_CloudWatch_LogsGroup"
}

resource "aws_sns_topic" "cis_alerts_sns_topic" {
  name = "cis-alerts-sns-topic"
}

resource "aws_cloudwatch_metric_alarm" "cis_aws_organizations" {
  alarm_name                = "CIS-4.15-AWS-Organizations"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_aws_organizations.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_aws_organizations" {
  name           = "CIS-4.15-AWS-Organizations"
  pattern        = "{ ($.eventSource = \"organizations.amazonaws.com\") && (($.eventName = AcceptHandshake) || ($.eventName = AttachPolicy) || ($.eventName = CreateAccount) || ($.eventName = PutBucketLifecycle) || ($.eventName = CreateOrganizationalUnit) || ($.eventName = CreatePolicy) || ($.eventName = DeclineHandshake) || ($.eventName = DeleteOrganization) || ($.eventName = DeleteOrganizationalUnit) || ($.eventName = DeletePolicy) || ($.eventName = DetachPolicy) || ($.eventName = DisablePolicyType) || ($.eventName = EnablePolicyType) || ($.eventName = InviteAccountToOrganization) || ($.eventName = LeaveOrganization) || ($.eventName = MoveAccount) || ($.eventName = RemoveAccountFromOrganization) || ($.eventName = UpdatePolicy) || ($.eventName = UpdateOrganizationalUni)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.15-AWS-Organizations"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

```
## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-2"
}

resource "aws_cloudwatch_log_group" "CIS_CloudWatch_LogsGroup" {
  name = "CIS_CloudWatch_LogsGroup"
}

resource "aws_sns_topic" "cis_alerts_sns_topic" {
  name = "cis-alerts-sns-topic"
}

resource "aws_cloudwatch_metric_alarm" "cis_aws_organizations" {
  alarm_name                = "CIS-4.15-AWS-Organizations"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_aws_organizations.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_aws_organizations" {
  name           = "CIS-4.15-AWS-Organizations"
  pattern        = "{ ($.eventSource = \"organizations.amazonaws.com\") && (($.eventName = AttachPolicy) || ($.eventName = CreateAccount) || ($.eventName = PutBucketLifecycle) || ($.eventName = CreateOrganizationalUnit) || ($.eventName = CreatePolicy) || ($.eventName = DeclineHandshake) || ($.eventName = DeleteOrganization) || ($.eventName = DeleteOrganizationalUnit) || ($.eventName = DeletePolicy) || ($.eventName = DetachPolicy) || ($.eventName = DisablePolicyType) || ($.eventName = EnablePolicyType) || ($.eventName = InviteAccountToOrganization) || ($.eventName = LeaveOrganization) || ($.eventName = MoveAccount) || ($.eventName = RemoveAccountFromOrganization) || ($.eventName = UpdatePolicy) || ($.eventName = UpdateOrganizationalUni)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.15-AWS-Organizations"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

```

```terraform
provider "aws" {
  region = "us-east-2"
}

resource "aws_cloudwatch_log_group" "CIS_CloudWatch_LogsGroup" {
  name = "CIS_CloudWatch_LogsGroup"
}

resource "aws_sns_topic" "cis_alerts_sns_topic" {
  name = "cis-alerts-sns-topic"
}

resource "aws_cloudwatch_metric_alarm" "cis_aws_organizations" {
  alarm_name                = "CIS-4.15-AWS-Organizations"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = "OTHER FILTER"
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_aws_organizations" {
  name           = "CIS-4.15-AWS-Organizations"
  pattern        = "{ ($.eventSource = \"organizations.amazonaws.com\") && (($.eventName = \"AcceptHandshake\") || ($.eventName = 'AttachPolicy') || ($.eventName = CreateAccount) || ($.eventName = PutBucketLifecycle) || ($.eventName = CreateOrganizationalUnit) || ($.eventName = CreatePolicy) || ($.eventName = DeclineHandshake) || ($.eventName = DeleteOrganization) || ($.eventName = DeleteOrganizationalUnit) || ($.eventName = DeletePolicy) || ($.eventName = DetachPolicy) || ($.eventName = DisablePolicyType) || ($.eventName = EnablePolicyType) || ($.eventName = InviteAccountToOrganization) || ($.eventName = LeaveOrganization) || ($.eventName = MoveAccount) || ($.eventName = RemoveAccountFromOrganization) || ($.eventName = UpdatePolicy) || ($.eventName = UpdateOrganizationalUni)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.15-AWS-Organizations"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

```
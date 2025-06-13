---
title: "CloudWatch Disabling Or Scheduled Deletion Of Customer Created CMK Alarm Missing"
meta:
  name: "aws/cloudwatch_disabling_or_scheduled_deletion_of_customer_created_cmk_alarm_missing"
  id: "56a585f5-555c-48b2-8395-e64e4740a9cf"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudwatch_disabling_or_scheduled_deletion_of_customer_created_cmk_alarm_missing`

**Id:** `56a585f5-555c-48b2-8395-e64e4740a9cf`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Ensure a log metric filter and alarm exist for disabling or scheduled deletion of customer created CMK

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_metric_filter#pattern)

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

resource "aws_cloudwatch_metric_alarm" "cis_disable_delete_cmk" {
  alarm_name                = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_disable_delete_cmk.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_disable_delete_cmk" {
  name           = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  pattern        = "{ ($.eventSource = \"kms.amazonaws.com\") || (($.eventName = DisableKey) || ($.eventName = ScheduleKeyDeletion)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.7-Disable-Scheduled-Delete-CMK"
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

resource "aws_cloudwatch_metric_alarm" "cis_disable_delete_cmk" {
  alarm_name                = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_disable_delete_cmk.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_disable_delete_cmk" {
  name           = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  pattern        = "{ ($.eventSource = \"kms.amazonaws.com\") && (($.eventName = ScheduleKeyDeletion)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.7-Disable-Scheduled-Delete-CMK"
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

resource "aws_cloudwatch_metric_alarm" "cis_disable_delete_cmk" {
  alarm_name                = "CIS-4.7-Disable-Scheduled-Delete-CMK"
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

resource "aws_cloudwatch_log_metric_filter" "cis_disable_delete_cmk" {
  name           = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  pattern        = "{ ($.eventSource = \"kms.amazonaws.com\") && (($.eventName = DisableKey) || ($.eventName = ScheduleKeyDeletion)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.7-Disable-Scheduled-Delete-CMK"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

```

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

resource "aws_cloudwatch_metric_alarm" "cis_disable_delete_cmk" {
  alarm_name                = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_disable_delete_cmk.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_disable_delete_cmk" {
  name           = "CIS-4.7-Disable-Scheduled-Delete-CMK"
  pattern        = "{ ($.eventSource = \"kms.amazonaws.com\") && (($.eventName = DisableKey) || ($.eventName = ScheduleKeyDeletion)) }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-4.7-Disable-Scheduled-Delete-CMK"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

```
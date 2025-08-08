---
title: "CloudWatch console sign-in without MFA alarm missing"
group_id: "rules/terraform/aws"
meta:
  name: "aws/cloudwatch_management_console_sign_in_without_mfa_alarm_missing"
  id: "44ceb4fa-0897-4fd2-b676-30e7a58f2933"
  display_name: "CloudWatch console sign-in without MFA alarm missing"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `44ceb4fa-0897-4fd2-b676-30e7a58f2933`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Low

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_metric_filter#pattern)

### Description

 This check ensures that a CloudWatch log metric filter and alarm are properly configured to detect AWS Management Console sign-ins that occur without multi-factor authentication (MFA), using log patterns such as `{ ($.eventName = "ConsoleLogin") && ($.additionalEventData.MFAUsed != "Yes") }`. Without this monitoring, administrators may be unaware of insecure single-factor logins to the management console, increasing the risk of unauthorized or compromised account access. If left unaddressed, attackers or malicious insiders could exploit accounts lacking MFA to gain elevated access, potentially resulting in data breaches or unauthorized changes to AWS resources.


## Compliant Code Examples
```terraform
resource "aws_cloudwatch_log_metric_filter" "cis_unauthorized_api_calls_metric_filter" {
  name           = "CIS-UnauthorizedAPICalls"
  pattern        = "{ ($.errorCode = \"*UnauthorizedOperation\") || ($.errorCode = \"AccessDenied*\") }"
  log_group_name = aws_cloudwatch_log_group.cis_cloudwatch_logsgroup.name

  metric_transformation {
    name      = "CIS-UnauthorizedAPICalls"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_unauthorized_api_calls_cw_alarm" {
  alarm_name                = "CIS-3.1-UnauthorizedAPICalls"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_unauthorized_api_calls_metric_filter.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring unauthorized API calls will help reveal application errors and may reduce time to detect malicious activity."
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

resource "aws_cloudwatch_log_metric_filter" "cis_no_mfa_console_signin_metric_filter" {
  name           = "CIS-ConsoleSigninWithoutMFA"
  pattern        = "{ ($.eventName = \"ConsoleLogin\") && ($.additionalEventData.MFAUsed != \"Yes\") }"
  log_group_name = aws_cloudwatch_log_group.cis_cloudwatch_logsgroup.name

  metric_transformation {
    name      = "CIS-ConsoleSigninWithoutMFA"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_no_mfa_console_signin_cw_alarm" {
  alarm_name                = "CIS-3.2-ConsoleSigninWithoutMFA"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_no_mfa_console_signin_metric_filter.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring for single-factor console logins will increase visibility into accounts that are not protected by MFA."
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_cloudwatch_log_metric_filter" "cis_no_mfa_console_signin_metric_filter" {
  name           = "CIS-ConsoleSigninWithoutMFA"
  pattern        = "{ $.additionalEventData.MFAUsed != \"Yes\" }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-ConsoleSigninWithoutMFA"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_no_mfa_console_signin_cw_alarm" {
  alarm_name                = "CIS-3.2-ConsoleSigninWithoutMFA"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_no_mfa_console_signin_metric_filter.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring for single-factor console logins will increase visibility into accounts that are not protected by MFA."
  alarm_actions             = [aws_sns_topic.CIS_Alerts_SNS_Topic.arn]
  insufficient_data_actions = []
}

```

```terraform
resource "aws_cloudwatch_log_metric_filter" "cis_unauthorized_api_calls_metric_filter" {
  name           = "CIS-UnauthorizedAPICalls"
  pattern        = "{ ($.errorCode = \"*UnauthorizedOperation\") || ($.errorCode = \"AccessDenied*\") }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-UnauthorizedAPICalls"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_unauthorized_api_calls_cw_alarm" {
  alarm_name                = "CIS-3.1-UnauthorizedAPICalls"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_unauthorized_api_calls_metric_filter.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring unauthorized API calls will help reveal application errors and may reduce time to detect malicious activity."
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}
resource "aws_cloudwatch_log_metric_filter" "cis_no_mfa_console_signin_metric_filter" {
  name           = "CIS-ConsoleSigninWithoutMFA"
  pattern        = "{ ($.eventName = \"ConsoleLogin\") && ($.additionalEventData.MFAUsed != \"Yes\") }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-ConsoleSigninWithoutMFA"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_no_mfa_console_signin_cw_alarm" {
  alarm_name                = "CIS-3.2-ConsoleSigninWithoutMFA"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = "XXXX NOT YOUR FILTER XXXX"
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring for single-factor console logins will increase visibility into accounts that are not protected by MFA."
  alarm_actions             = [aws_sns_topic.cis_alerts_sns_topic.arn]
  insufficient_data_actions = []
}

```

```terraform
resource "aws_cloudwatch_log_metric_filter" "cis_no_mfa_console_signin_metric_filter" {
  name           = "CIS-ConsoleSigninWithoutMFA"
  pattern        = "{ ($.eventName = \"ConsoleLogin\") || ($.additionalEventData.MFAUsed != \"Yes\") }"
  log_group_name = aws_cloudwatch_log_group.CIS_CloudWatch_LogsGroup.name

  metric_transformation {
    name      = "CIS-ConsoleSigninWithoutMFA"
    namespace = "CIS_Metric_Alarm_Namespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_metric_alarm" "cis_no_mfa_console_signin_cw_alarm" {
  alarm_name                = "CIS-3.2-ConsoleSigninWithoutMFA"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "1"
  metric_name               = aws_cloudwatch_log_metric_filter.cis_no_mfa_console_signin_metric_filter.id
  namespace                 = "CIS_Metric_Alarm_Namespace"
  period                    = "300"
  statistic                 = "Sum"
  threshold                 = "1"
  alarm_description         = "Monitoring for single-factor console logins will increase visibility into accounts that are not protected by MFA."
  alarm_actions             = [aws_sns_topic.CIS_Alerts_SNS_Topic.arn]
  insufficient_data_actions = []
}

```
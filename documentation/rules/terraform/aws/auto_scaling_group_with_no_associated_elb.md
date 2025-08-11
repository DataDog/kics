---
title: "Auto scaling group with no associated ELB"
group_id: "rules/terraform/aws"
meta:
  name: "aws/auto_scaling_group_with_no_associated_elb"
  id: "8e94dced-9bcc-4203-8eb7-7e41202b2505"
  display_name: "Auto scaling group with no associated ELB"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata

**Id:** `8e94dced-9bcc-4203-8eb7-7e41202b2505`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Availability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group#load_balancers)

### Description

 AWS Auto Scaling Groups (ASGs) should have associated Elastic Load Balancers (ELBs) to ensure high availability and efficient distribution of incoming traffic. Without specifying the `load_balancers` attribute, as shown in the example below, instances launched by the ASG may become inaccessible or experience uneven traffic distribution, leading to reduced reliability and performance:

```
resource "aws_autoscaling_group" "bar" {
  ...
  // missing `load_balancers`
}
```

To remediate this, the `load_balancers` attribute should be defined, as in the secure example below:

```
resource "aws_autoscaling_group" "bar3" {
  ...
  load_balancers = ["elb_1"]
}
```


## Compliant Code Examples
```terraform
resource "aws_autoscaling_group" "foo" {
  name_prefix          = "bar-"
  vpc_zone_identifier  = ["subnet-abcd1234", "subnet-1a2b3c4d"]
  launch_configuration = aws_launch_configuration.foobar.name
  target_group_arns    = ["bar", "baz", "qux"]
  min_size             = 1
  max_size             = 3
  desired_capacity     = 2
  instance_refresh {
    strategy = "Rolling"
    preferences {
      min_healthy_percentage = 50
    }
  }
}

```

```terraform
resource "aws_autoscaling_group" "bar3" {
  availability_zones = ["us-east-1a"]
  desired_capacity   = 1
  max_size           = 1
  min_size           = 1

  launch_template {
    id      = aws_launch_template.foobar.id
    version = "$Latest"
  }

  load_balancers = ["elb_1"]
}

```

```terraform
resource "aws_autoscaling_group" "my_asg" {
  name_prefix          = format("%s-", var.name)
  vpc_zone_identifier  = var.private_zone_identifiers
  launch_configuration = aws_launch_configuration.config.name
  target_group_arns    = [var.target_group_arns]
  min_size             = 1
  max_size             = 2
  desired_capacity     = 1
  instance_refresh {
    strategy = "Rolling"
    preferences {
      min_healthy_percentage = 50
    }
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_autoscaling_group" "foo" {
  name_prefix          = "bar-"
  vpc_zone_identifier  = ["subnet-abcd1234", "subnet-1a2b3c4d"]
  launch_configuration = aws_launch_configuration.foobar.name
  min_size             = 1
  max_size             = 3
  desired_capacity     = 2
  instance_refresh {
    strategy = "Rolling"
    preferences {
      min_healthy_percentage = 50
    }
  }
}

```

```terraform
resource "aws_autoscaling_group" "positive2" {
  availability_zones = ["us-east-1a"]
  desired_capacity   = 1
  max_size           = 1
  min_size           = 1

  launch_template {
    id      = aws_launch_template.foobar.id
    version = "$Latest"
  }

  load_balancers = []
}

```

```terraform
resource "aws_autoscaling_group" "foo" {
  name_prefix          = "bar-"
  vpc_zone_identifier  = ["subnet-abcd1234", "subnet-1a2b3c4d"]
  launch_configuration = aws_launch_configuration.foobar.name
  target_group_arns    = []
  min_size             = 1
  max_size             = 3
  desired_capacity     = 2
  instance_refresh {
    strategy = "Rolling"
    preferences {
      min_healthy_percentage = 50
    }
  }
}

```
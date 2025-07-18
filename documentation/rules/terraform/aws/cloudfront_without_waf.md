---
title: "CloudFront Without WAF"
group-id: "rules/terraform/aws"
meta:
  name: "aws/cloudfront_without_waf"
  id: "1419b4c6-6d5c-4534-9cf6-6a5266085333"
  display_name: "CloudFront Without WAF"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `1419b4c6-6d5c-4534-9cf6-6a5266085333`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution)

### Description

 AWS CloudFront distributions should be integrated with AWS Web Application Firewall (WAF) by specifying the `web_acl_id` attribute. If a CloudFront distribution is created without a WAF association (`web_acl_id` missing), the application is left vulnerable to common web attacks such as SQL injection and cross-site scripting. Failing to link a WAF to your CloudFront distribution increases the risk of security breaches, potential data loss, and service interruption.


## Compliant Code Examples
```terraform


module "acm" {
  source      = "terraform-aws-modules/acm/aws"
  version     = "~> v2.0"
  domain_name = var.site_domain
  zone_id     = data.aws_route53_zone.this.zone_id
  tags        = var.tags

  providers = {
    aws = aws.us_east_1 # cloudfront needs acm certificate to be from "us-east-1" region
  }
}

resource "aws_cloudfront_distribution" "negative1" {
  origin {
    domain_name = var.public_alb_domain
    origin_id   = "alb"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
  }

  enabled         = true
  is_ipv6_enabled = true
  comment         = var.site_domain
  web_acl_id      = "test"

  aliases = [var.site_domain]

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["*"]

      cookies {
        forward = "all"
      }

    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
    compress               = true
  }

  # Cache behavior with precedence 0
  ordered_cache_behavior {
    path_pattern     = "wp-content/*"
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["Host"]

      cookies {
        forward = "all"
      }
    }

    min_ttl                = 900
    default_ttl            = 900
    max_ttl                = 900
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  # Cache behavior with precedence 1
  ordered_cache_behavior {
    path_pattern     = "wp-includes/*"
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["Host"]

      cookies {
        forward = "all"
      }
    }

    min_ttl                = 3600
    default_ttl            = 3600
    max_ttl                = 3600
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }
  price_class = var.cf_price_class
  tags        = var.tags
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }


  viewer_certificate {
    acm_certificate_arn      = module.acm.this_acm_certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.1_2016"
  }

  # By default, cloudfront caches error for five minutes. There can be situation when a developer has accidentally broken the website and you would not want to wait for five minutes for the error response to be cached.
  # https://docs.aws.amazon.com/AmazonS3/latest/dev/CustomErrorDocSupport.html
  custom_error_response {
    error_code            = 400
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 403
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 404
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 405
    error_caching_min_ttl = var.error_ttl
  }

  depends_on = [
    aws_ecs_service.this
  ]
}




```
## Non-Compliant Code Examples
```terraform


module "acm" {
  source      = "terraform-aws-modules/acm/aws"
  version     = "~> v2.0"
  domain_name = var.site_domain
  zone_id     = data.aws_route53_zone.this.zone_id
  tags        = var.tags

  providers = {
    aws = aws.us_east_1 # cloudfront needs acm certificate to be from "us-east-1" region
  }
}

resource "aws_cloudfront_distribution" "positive1" {
  origin {
    domain_name = var.public_alb_domain
    origin_id   = "alb"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
  }

  enabled         = true
  is_ipv6_enabled = true
  comment         = var.site_domain

  aliases = [var.site_domain]

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["*"]

      cookies {
        forward = "all"
      }

    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
    compress               = true
  }

  # Cache behavior with precedence 0
  ordered_cache_behavior {
    path_pattern     = "wp-content/*"
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["Host"]

      cookies {
        forward = "all"
      }
    }

    min_ttl                = 900
    default_ttl            = 900
    max_ttl                = 900
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  # Cache behavior with precedence 1
  ordered_cache_behavior {
    path_pattern     = "wp-includes/*"
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "alb"

    forwarded_values {
      query_string = true
      headers      = ["Host"]

      cookies {
        forward = "all"
      }
    }

    min_ttl                = 3600
    default_ttl            = 3600
    max_ttl                = 3600
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }
  price_class = var.cf_price_class
  tags        = var.tags
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }


  viewer_certificate {
    acm_certificate_arn      = module.acm.this_acm_certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.1_2016"
  }

  # By default, cloudfront caches error for five minutes. There can be situation when a developer has accidentally broken the website and you would not want to wait for five minutes for the error response to be cached.
  # https://docs.aws.amazon.com/AmazonS3/latest/dev/CustomErrorDocSupport.html
  custom_error_response {
    error_code            = 400
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 403
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 404
    error_caching_min_ttl = var.error_ttl
  }

  custom_error_response {
    error_code            = 405
    error_caching_min_ttl = var.error_ttl
  }

  depends_on = [
    aws_ecs_service.this
  ]
}




```
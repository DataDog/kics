---
title: "CloudFront without minimum protocol TLS 1.2"
group_id: "rules/terraform/aws"
meta:
  name: "aws/cloudfront_without_minimum_protocol_tls_1.2"
  id: "00e5e55e-c2ff-46b3-a757-a7a1cd802456"
  display_name: "CloudFront without minimum protocol TLS 1.2"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `00e5e55e-c2ff-46b3-a757-a7a1cd802456`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution)

### Description

 Amazon CloudFront distributions should enforce a minimum TLS protocol version of at least TLS 1.2 to ensure secure encryption between clients and CloudFront. Allowing earlier versions such as TLS 1.0 or TLS 1.1 exposes distributions to known vulnerabilities and weak ciphers, increasing the risk of data interception and man-in-the-middle attacks. This can be securely enforced in Terraform using the `viewer_certificate` block with `minimum_protocol_version = "TLSv1.2_2018"` or higher, as shown below:

```
viewer_certificate {
  cloudfront_default_certificate = false
  minimum_protocol_version = "TLSv1.2_2018"
}
```


## Compliant Code Examples
```terraform
resource "aws_cloudfront_distribution" "negative1" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  comment             = "Some comment"
  default_root_object = "index.html"

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA", "GB", "DE"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = false
    minimum_protocol_version = "TLSv1.2_2018"
  }
}

resource "aws_cloudfront_distribution" "negative2" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  comment             = "Some comment"
  default_root_object = "index.html"

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA", "GB", "DE"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = false
    minimum_protocol_version = "TLSv1.2_2019"
  }
}

```

```terraform
module "my_distribution" {
  source = "terraform-aws-modules/cloudfront/aws"
  version = "~> 3.0"

  aliases = ["mysite.example.com"]

  comment             = "Some comment"
  enabled             = true
  default_root_object = "index.html"
  is_ipv6_enabled     = true
  price_class         = "PriceClass_100"
  retain_on_delete    = false
  wait_for_deployment = true
  tags = {
    my-products = "cloudfront"
  }

  create_origin_access_identity = true
  origin_access_identities = {
    s3_bucket_one = "My awesome CloudFront can access"
    s3_bucket_two = "CloudFront can also access this one"
  }

  origin = [
    {
      s3_origin_config = {
        origin_access_identity = "s3_bucket_one"
      }
      domain_name = "MY_S3_DNS_NAME"
      origin_id   = "S3-my-group"
    },
    {
      custom_origin_config = {
        http_port              = 80
        https_port             = 443
        origin_keepalive_timeout = 5
        origin_protocol_policy = "match-viewer"
        origin_read_timeout    = 30
        origin_ssl_protocols   = ["SSLv3", "TLSv1", "TLSv1.1", "TLSv1.2"]
      }
      domain_name = "MY_DOMAIN"
      origin_id   = "stable"
      origin_path = "/some-path"
      custom_header = {
        name  = "x-authentication-token"
        value = "123456aaabbbcc"
      }
    }
  ]

  default_cache_behavior = {
    allowed_methods            = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods             = ["GET", "HEAD"]
    target_origin_id           = "S3-my-group"
    viewer_protocol_policy     = "allow-all"
    compress                   = true
    query_string               = false
    query_string_cache_keys    = ["querystring1"]
    cookie_forward             = "all"
    lambda_function_association = {
      event_type = "viewer-request"
      lambda_arn = "arn:aws:lambda:us-east-1:123456789012:function:static-web-content"
    }
  }

  ordered_cache_behavior = [
    {
      path_pattern     = "/content/immutable/*"
      allowed_methods  = ["GET", "HEAD"]
      cached_methods   = ["GET", "HEAD"]
      target_origin_id = "S3-my-group"

      forwarded_values = {
        query_string = false
        headers      = ["Origin"]
        cookies = {
          forward = "none"
        }
      }

      min_ttl                = 0
      default_ttl            = 86400
      max_ttl                = 31536000
      compress               = true
      viewer_protocol_policy = "redirect-to-https"
      lambda_function_association = {
        event_type   = "viewer-response"
        lambda_arn   = "arn:aws:lambda:us-east-1:123456789012:function:static-web-content"
        include_body = "false"
      }
    },
    {
      path_pattern           = "/content/*"
      allowed_methods        = ["GET", "HEAD", "OPTIONS"]
      cached_methods         = ["GET", "HEAD"]
      target_origin_id       = "S3-my-group"
      compress               = true
      viewer_protocol_policy = "redirect-to-https"

      forwarded_values = {
        query_string = false
        headers      = ["Origin"]
        cookies = {
          forward = "none"
        }
      }
    }
  ]

  viewer_certificate {
    cloudfront_default_certificate = false
    minimum_protocol_version = "TLSv1.2_2021"
  }

  logging_config = {
    include_cookies = false
    bucket          = "my_logs.s3.amazonaws.com"
    prefix          = "myprefix"
  }

  restriction = {
    locations = ["US", "CA", "GB", "DE"]
  }
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_cloudfront_distribution" "positive3" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  comment             = "Some comment"
  default_root_object = "index.html"

  default_cache_behavior {
    #settings
  }

  restrictions {
    #restrictions
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

```

```terraform
resource "aws_cloudfront_distribution" "positive1" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  comment             = "Some comment"
  default_root_object = "index.html"

  default_cache_behavior {
    #settings
  }

  restrictions {
    #restrictions
  }
}

```

```terraform
resource "aws_cloudfront_distribution" "positive4" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  comment             = "Some comment"
  default_root_object = "index.html"

  default_cache_behavior {
    #settings
  }

  restrictions {
    #restrictions
  }

  viewer_certificate {
    cloudfront_default_certificate = false
  }
}

```
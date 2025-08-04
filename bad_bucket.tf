terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

resource "aws_s3_bucket" "example4" {
  bucket = "my-tf-example-bucket4"
  versioning {
    enabled = false
  }
}

resource "aws_s3_bucket" "example23" {
  bucket = "my-tf-example-bucket4"
  versioning {
    enabled = false
  }
}

provider "aws" {
  # Configuration options
}

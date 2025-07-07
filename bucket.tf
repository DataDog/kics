terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

provider "aws" {
  # Configuration options
}

resource "aws_s3_bucket" "example6" {
  bucket = "my-tf-example-bucket6"
  versioning {
    enabled = false
  }
}

resource "aws_s3_bucket" "example3" {
  bucket = "my-tf-example-bucket3"
  versioning {
    enabled = false
  }
}

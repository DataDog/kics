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

resource "aws_s3_bucket" "example4" {
  bucket = "my-tf-example-bucket4"
}


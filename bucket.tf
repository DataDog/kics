terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

resource "aws_s3_bucket" "example2" {
  bucket = "my-tf-example-bucket2"
}

provider "aws" {
  # Configuration options
}

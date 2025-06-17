terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

resource "aws_s3_bucket" "example" {
  bucket = "my-tf-example-bucket"
}

provider "aws" {
  # Configuration options
}

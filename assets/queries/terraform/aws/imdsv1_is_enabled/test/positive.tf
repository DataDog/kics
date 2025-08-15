# Test case 2: aws_instance with incorrect http_tokens = "optional"
resource "aws_instance" "bad_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "optional" # ❌ Should be "required"
  }
}

# Test case 3: aws_launch_template with incorrect http_tokens = "optional"
resource "aws_launch_template" "bad_example" {
  name_prefix   = "example"
  image_id      = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "optional" # ❌ Should be "required"
  }
}

# Test case 1: Missing metadata_options entirely (should trigger finding) - K9VULN-7671 scenario
resource "aws_launch_template" "missing_metadata_options" {
  name_prefix   = "missing-metadata"
  image_id      = "ami-123456"
  instance_type = "t2.micro"
}

# This is not a test case, but validates case https://datadoghq.atlassian.net/browse/K9VULN-7671
resource "aws_launch_template" "good_example" {
  name_prefix   = "secure"
  image_id      = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "required" # ✅ Correct value
  }
}
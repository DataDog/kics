resource "aws_elasticsearch_domain" "good_example" {
  domain_name = "example"

  vpc_options {

  
  encrypt_at_rest {
  	enabled = true
  	}
  }
}

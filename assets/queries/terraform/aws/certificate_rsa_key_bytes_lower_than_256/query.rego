package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {


	services := {"aws_api_gateway_domain_name", "aws_iam_server_certificate", "aws_acm_certificate"}



	resource[name].certificate_body.rsa_key_bytes < 256

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

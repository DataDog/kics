package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {

	services := {"aws_api_gateway_domain_name", "aws_iam_server_certificate", "aws_acm_certificate"}


	expiration_date := resource[name].certificate_body.expiration_date

	common_lib.expired(expiration_date)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

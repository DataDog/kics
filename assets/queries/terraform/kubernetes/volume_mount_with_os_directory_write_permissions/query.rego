package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	is_object(containers[y].volume_mount) == true
	is_os_dir(containers[y].volume_mount)

	containers[y].volume_mount.read_only == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	volumeMounts := containers[y].volume_mount
	is_array(volumeMounts) == true
	is_os_dir(volumeMounts[j])
	volumeMounts[j].read_only == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	volumeMounts := containers.volume_mount
	is_object(volumeMounts) == true
	is_os_dir(volumeMounts)
	volumeMounts.read_only == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": json.marshal({
			"before": "false",
			"after": "true",
		}),
		"remediationType": "replacement",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	volumeMounts := containers.volume_mount
	is_array(volumeMounts) == true
	is_os_dir(volumeMounts[j])
	volumeMounts[j].read_only == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": json.marshal({
			"before": "false",
			"after": "true",
		}),
		"remediationType": "replacement",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	volumeMounts := containers[y].volume_mount
	is_object(volumeMounts) == true
	is_os_dir(volumeMounts)

	not common_lib.valid_key(volumeMounts, "read_only")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",



		"remediation": "read_only = true",
		"remediationType": "addition",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	volumeMounts := containers[y].volume_mount
	is_array(volumeMounts) == true
	is_os_dir(volumeMounts[j])
	volumeMountTypes := volumeMounts[_]

	not common_lib.valid_key(volumeMountTypes, "read_only")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",



	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	volumeMounts := containers.volume_mount
	is_object(volumeMounts) == true
	is_os_dir(volumeMounts)

	not common_lib.valid_key(volumeMounts, "read_only")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",



		"remediation": "read_only = true",
		"remediationType": "addition",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	volumeMounts := containers.volume_mount
	is_array(volumeMounts) == true
	is_os_dir(volumeMounts[j])
	volumeMountTypes := volumeMounts[j]

	not common_lib.valid_key(volumeMountTypes, "read_only")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",



		"remediation": "read_only = true",
		"remediationType": "addition",
	}
}

is_os_dir(volumeMounts) {
	hostSensitiveDir = {"/bin", "/sbin", "/boot", "/cdrom", "/dev", "/etc", "/home", "/lib", "/media", "/proc", "/root", "/run", "/seLinux", "/srv", "/usr", "/var"}
	startswith(volumeMounts.mount_path, hostSensitiveDir[_])
} else {
	volumeMounts.mount_path == "/"
}

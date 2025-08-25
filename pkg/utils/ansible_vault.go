/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package utils

import "context"

// DecryptAnsibleVault verifies if the fileContent is encrypted by ansible-vault. If yes, the function decrypts it
func DecryptAnsibleVault(ctx context.Context, fileContent []byte, secret string) []byte {
	// logger := log.Ctx(ctx)
	// match, err := regexp.MatchString(`^\s*\$ANSIBLE_VAULT.*`, string(fileContent))
	// if err != nil {
	// 	return fileContent
	// }
	// if secret != "" && match {
	// 	content, err := vault.Decrypt(string(fileContent), secret)

	// 	if err == nil {
	// 		logger.Info().Msg("Decrypting Ansible Vault file")
	// 		fileContent = []byte(content)
	// 	}
	// }
	return fileContent
}

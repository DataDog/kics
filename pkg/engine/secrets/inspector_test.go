/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package secrets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntropyInterval(t *testing.T) {
	inputs := []struct {
		name    string
		entropy Entropy
		token   string
		want    bool
	}{
		{
			name: "empty_within_bounds",
			entropy: Entropy{
				Group: 0, // not relevant for this test
				Min:   0,
				Max:   0,
			},
			token: "",
			want:  true,
		},
		{
			name: "empty_outside_bounds",
			entropy: Entropy{
				Group: 0, // not relevant for this test
				Min:   1,
				Max:   2,
			},
			token: "",
			want:  false,
		},
		{
			name: "token_larger_than_max_entropy",
			entropy: Entropy{
				Group: 0,
				Min:   1,
				Max:   2, // 3.655152
			},
			token: "d68d2897add9bc2203a5ed0632a5cdd8ff8cefb0",
			want:  false,
		},
		{
			name: "token_within_bounds",
			entropy: Entropy{
				Group: 0,
				Min:   1,
				Max:   3.7,
			},
			token: "d68d2897add9bc2203a5ed0632a5cdd8ff8cefb0",
			want:  true,
		},
		{
			name: "password_less_than_minimum",
			entropy: Entropy{
				Group: 0,
				Min:   3, // 2.321928
				Max:   4,
			},
			token: "passx",
			want:  false,
		},
		{
			name: "password_within_bound",
			entropy: Entropy{
				Group: 0,
				Min:   2, // 2.321928
				Max:   3,
			},
			token: "passx",
			want:  true,
		},
	}
	for _, in := range inputs {
		highEntropy, entropyValue := CheckEntropyInterval(in.entropy, in.token)
		require.Equal(t, in.want, highEntropy, "test[%s] CheckEntropyInterval(%+v, %s) = %v, want %v :: entropyValue %f",
			in.name,
			in.entropy,
			in.token,
			highEntropy,
			in.want,
			entropyValue,
		)
	}
}

package topologymanager

import (
	"testing"
)

func TestCanAdmitPodResult(t *testing.T) {
	tcases := []struct {
		name     string
		affinity bool
		expected bool
	}{
		{
			name:     "affinity is set to false in topology hints",
			affinity: false,
			expected: false,
		},
		{
			name:     "affinity is set to true in topology hints",
			affinity: true,
			expected: true,
		},
	}

	for _, tc := range tcases {
		policy := NewStrictPolicy()
		hints := TopologyHints{
			Affinity: tc.affinity,
		}
		result := policy.CanAdmitPodResult(hints)

		if result.Admit != tc.expected {
			t.Errorf("expected Admit field in result to be %t, got %t", tc.expected, result.Admit)
		}

		if tc.expected == false {
			if len(result.Reason) == 0 {
				t.Errorf("expected Reason field to be not empty")
			}
			if len(result.Message) == 0 {
				t.Errorf("expected Message field to be not empty")
			}
		}
	}
}

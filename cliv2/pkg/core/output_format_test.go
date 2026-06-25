package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/snyk/go-application-framework/pkg/local_workflows/output_workflow"
)

func TestResolveAgentDefaultFormat(t *testing.T) {
	tests := []struct {
		name          string
		snykOutputEnv string
		agentDetected bool
		expected      string
	}{
		{
			name:     "no agent, no env → empty (GAF uses text default)",
			expected: "",
		},
		{
			name:          "agent detected, no env → TOON",
			agentDetected: true,
			expected:      output_workflow.TOON_MIME_TYPE,
		},
		{
			name:          "agent + SNYK_OUTPUT=json → JSON (env beats detection)",
			agentDetected: true,
			snykOutputEnv: "json",
			expected:      output_workflow.JSON_MIME_TYPE,
		},
		{
			name:          "agent + SNYK_OUTPUT=human → empty (human = text default)",
			agentDetected: true,
			snykOutputEnv: "human",
			expected:      "",
		},
		{
			name:          "agent + SNYK_OUTPUT=toon → TOON",
			agentDetected: true,
			snykOutputEnv: "toon",
			expected:      output_workflow.TOON_MIME_TYPE,
		},
		{
			name:          "no agent + SNYK_OUTPUT=toon → TOON",
			agentDetected: false,
			snykOutputEnv: "toon",
			expected:      output_workflow.TOON_MIME_TYPE,
		},
		{
			name:          "SNYK_OUTPUT unknown value treated as empty (ignored)",
			agentDetected: true,
			snykOutputEnv: "garbage",
			expected:      output_workflow.TOON_MIME_TYPE,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := resolveAgentDefaultFormat(tc.snykOutputEnv, tc.agentDetected)
			assert.Equal(t, tc.expected, got)
		})
	}
}

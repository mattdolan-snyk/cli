package core

import (
	"os"

	"github.com/snyk/cli/cliv2/internal/persona/agent"
	"github.com/snyk/go-application-framework/pkg/configuration"
	"github.com/snyk/go-application-framework/pkg/local_workflows/output_workflow"
)

// applyAgentOutputDefault sets OUTPUT_CONFIG_KEY_DEFAULT_FORMAT on config when the
// environment signals an AI agent and no explicit format flag is present.
// Precedence: SNYK_OUTPUT env > agent detection > nothing (GAF falls back to text).
func applyAgentOutputDefault(config configuration.Configuration) {
	_, detected := agent.DetectAgent()
	snykOutput := os.Getenv("SNYK_OUTPUT")
	if v := resolveAgentDefaultFormat(snykOutput, detected); v != "" {
		config.Set(output_workflow.OUTPUT_CONFIG_KEY_DEFAULT_FORMAT, v)
	}
}

// resolveAgentDefaultFormat returns the MIME type that should be set as the default
// output format, or "" if no override is needed (caller leaves GAF's text default).
// snykOutputEnv is the value of the SNYK_OUTPUT environment variable (empty if unset).
// agentDetected is true when a known AI agent was detected in the environment.
func resolveAgentDefaultFormat(snykOutputEnv string, agentDetected bool) string {
	switch snykOutputEnv {
	case "toon":
		return output_workflow.TOON_MIME_TYPE
	case "json":
		return output_workflow.JSON_MIME_TYPE
	case "sarif":
		return output_workflow.SARIF_MIME_TYPE
	case "human", "text":
		return ""
	case "":
		if agentDetected {
			return output_workflow.TOON_MIME_TYPE
		}
		return ""
	default:
		// Unknown value: fall through to agent detection.
		if agentDetected {
			return output_workflow.TOON_MIME_TYPE
		}
		return ""
	}
}

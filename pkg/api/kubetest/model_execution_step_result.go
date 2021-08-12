/*
 * Kubetest
 *
 * Efficient testing of k8s applications mandates a k8s native approach to test mgmt/definition/execution - kubetest provides a “quality control plane” that natively integrates testing activities into k8s development and operational workflows
 *
 * API version: 0.0.5
 * Contact: api@kubetest.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kubetest
import (
	"time"
)

// execution result data
type ExecutionStepResult struct {
	// step name
	Name string `json:"name,omitempty"`
	// script execution start time
	StartTime time.Time `json:"start-time,omitempty"`
	// script execution start time
	EndTime time.Time `json:"end-time,omitempty"`
	// execution step status
	Status string `json:"status,omitempty"`
	AssertionResults *AssertionResult `json:"assertionResults,omitempty"`
}

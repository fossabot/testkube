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

// object which should be returned by REST based executors
type Execution struct {
	// execution id
	Id string `json:"id,omitempty"`
	// script metadata content
	ScriptContent string `json:"script-content,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
	// execution status
	Status string `json:"status,omitempty"`
	Result *ExecutionResult `json:"result,omitempty"`
	// test start time
	StartTime time.Time `json:"start-time,omitempty"`
	// test end time
	EndTime time.Time `json:"end-time,omitempty"`
}

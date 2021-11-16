/*
 * TestKube API
 *
 * TestKube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

import (
	"time"
)

// API server script execution
type Execution struct {
	// execution id
	Id string `json:"id,omitempty"`
	// unique script name (CRD Script name)
	ScriptName string `json:"scriptName,omitempty"`
	// script type e.g. postman/collection
	ScriptType string `json:"scriptType,omitempty"`
	// execution name
	Name string `json:"name,omitempty"`
	// execution envs passed to executor
	Envs map[string]string `json:"envs,omitempty"`
	// additional flags passed to executor binary
	Flags map[string]string `json:"flags,omitempty"`
	// script metadata content
	ScriptContent string      `json:"scriptContent,omitempty"`
	Repository    *Repository `json:"repository,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
	// test start time
	StartTime time.Time `json:"startTime,omitempty"`
	// test end time
	EndTime         time.Time        `json:"endTime,omitempty"`
	ExecutionResult *ExecutionResult `json:"executionResult,omitempty"`
}

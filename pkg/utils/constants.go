/*
Copyright © 2021 The LitmusChaos Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

const (
	DefaultFileName = ".litmusconfig"

	// Default username
	DefaultUsername = "admin"

	// Default installation mode
	DefaultMode = "cluster"

	// Platform list
	PlatformList = "1. AWS\n2. GKE\n3. Openshift\n4. Rancher\n5. Others"

	// AWS identifier
	AWSIdentifier = "aws://"

	// GKE identifier
	GKEIdentifier = "gce://"

	// Openshift identifier
	OpenshiftIdentifier = "node.openshift.io/os_id"

	// Default platform name
	DefaultPlatform = "Others"

	// Label of subscriber agent being deployed
	ChaosAgentLabel = "app=subscriber"

	// Agent type is "external" for agents connected via litmusctl
	AgentType = "external"

	// Default namespace for agent installation
	DefaultNs = "litmus"

	// Default service account used for agent installation
	DefaultSA = "litmus"

	// Chaos agent connection yaml path
	ChaosYamlPath = "api/file"

	ChaosAgentPath = "targets"

	// Graphql server api path
	GQLAPIPath = "/api/query"

	// Auth server api path
	AuthAPIPath = "/auth"
)

var (
	// CompatibilityMatrix to store the compatible versions of litmusctl and ChaosCenter
	CompatibilityMatrix = map[string][]string{
		"0.6.0":  {"2.2.0", "2.3.0"},
		"0.7.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.8.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.9.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.10.0": {"2.10.0", "2.9.0", "2.11.0"},
		"0.11.0": {"2.10.0", "2.9.0", "2.11.0"},
	}
)

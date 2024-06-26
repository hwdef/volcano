/*
Copyright 2022 The Volcano Authors.

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

package jobflow

const (
	// Volcano string of volcano apiVersion
	Volcano = "volcano"
	// JobFlow kind of jobFlow
	JobFlow = "JobFlow"
	// CreatedByJobTemplate the vcjob annotation and label of created by jobTemplate
	CreatedByJobTemplate = "volcano.sh/createdByJobTemplate"
	// CreatedByJobFlow the vcjob annotation and label of created by jobFlow
	CreatedByJobFlow = "volcano.sh/createdByJobFlow"
)

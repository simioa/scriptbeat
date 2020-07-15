// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}

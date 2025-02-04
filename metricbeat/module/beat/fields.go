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

package beat

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "beat", asset.ModuleFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded zlib format compressed contents of module/beat.
func AssetBeat() string {
	return "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"
}

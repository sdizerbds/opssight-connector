/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package hub

import (
	"encoding/json"
	"fmt"
)

type RiskProfileCategory int

const (
	RiskProfileCategoryActivity      RiskProfileCategory = iota
	RiskProfileCategoryLicense       RiskProfileCategory = iota
	RiskProfileCategoryOperational   RiskProfileCategory = iota
	RiskProfileCategoryVersion       RiskProfileCategory = iota
	RiskProfileCategoryVulnerability RiskProfileCategory = iota
)

func (r RiskProfileCategory) String() string {
	switch r {
	case RiskProfileCategoryActivity:
		return "ACTIVITY"
	case RiskProfileCategoryLicense:
		return "LICENSE"
	case RiskProfileCategoryOperational:
		return "OPERATIONAL"
	case RiskProfileCategoryVersion:
		return "VERSION"
	case RiskProfileCategoryVulnerability:
		return "VULNERABILITY"
	default:
		panic(fmt.Errorf("invalid RiskProfileCategory value: %d", r))
	}
}

// func (r RiskProfileCategory) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`"%s"`, r.String())
// 	return []byte(jsonString), nil
// }

func (r *RiskProfileCategory) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	status, err := parseHubRiskProfileCategory(str)
	if err != nil {
		return err
	}
	*r = status
	return nil
}

func (r RiskProfileCategory) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RiskProfileCategory) UnmarshalText(text []byte) (err error) {
	status, err := parseHubRiskProfileCategory(string(text))
	if err != nil {
		return err
	}
	*r = status
	return nil
}

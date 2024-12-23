// Copyright 2024 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import "fmt"

type MachineClientInterface interface {
	GetMachines() ([]*Machine, error)
	GetMachine(name string) (*Machine, error)
}

func NewMachineClient(providerType string, accessKeyId string, accessKeySecret string, region string) (*MachineClientInterface, error) {
	var client MachineClientInterface
	var err error
	switch providerType {
	case "Aliyun":
		client, err = NewMachineAliyunClient(accessKeyId, accessKeySecret, region)
		return &client, err
	case "VMware":
		client, err = NewMachineVMwareClient(accessKeyId, accessKeySecret, region)
		return &client, err
	}

	return nil, fmt.Errorf("unsupported provider type: %s", providerType)
}

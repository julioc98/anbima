// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package enterpriseknowledgegraphservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/enterpriseknowledgegraphservice/2018-12-03/enterpriseknowledgegraphservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	S1 SkuName = original.S1
)

type BaseClient = original.BaseClient
type EnterpriseKnowledgeGraph = original.EnterpriseKnowledgeGraph
type EnterpriseKnowledgeGraphClient = original.EnterpriseKnowledgeGraphClient
type EnterpriseKnowledgeGraphProperties = original.EnterpriseKnowledgeGraphProperties
type EnterpriseKnowledgeGraphResponseList = original.EnterpriseKnowledgeGraphResponseList
type EnterpriseKnowledgeGraphResponseListIterator = original.EnterpriseKnowledgeGraphResponseListIterator
type EnterpriseKnowledgeGraphResponseListPage = original.EnterpriseKnowledgeGraphResponseListPage
type Error = original.Error
type ErrorBody = original.ErrorBody
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEnterpriseKnowledgeGraphClient(subscriptionID string) EnterpriseKnowledgeGraphClient {
	return original.NewEnterpriseKnowledgeGraphClient(subscriptionID)
}
func NewEnterpriseKnowledgeGraphClientWithBaseURI(baseURI string, subscriptionID string) EnterpriseKnowledgeGraphClient {
	return original.NewEnterpriseKnowledgeGraphClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnterpriseKnowledgeGraphResponseListIterator(page EnterpriseKnowledgeGraphResponseListPage) EnterpriseKnowledgeGraphResponseListIterator {
	return original.NewEnterpriseKnowledgeGraphResponseListIterator(page)
}
func NewEnterpriseKnowledgeGraphResponseListPage(getNextPage func(context.Context, EnterpriseKnowledgeGraphResponseList) (EnterpriseKnowledgeGraphResponseList, error)) EnterpriseKnowledgeGraphResponseListPage {
	return original.NewEnterpriseKnowledgeGraphResponseListPage(getNextPage)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

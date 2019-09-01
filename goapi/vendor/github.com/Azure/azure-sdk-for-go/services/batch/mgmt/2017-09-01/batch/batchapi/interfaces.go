package batchapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2017-09-01/batch"
	"github.com/Azure/go-autorest/autorest"
)

// AccountClientAPI contains the set of methods on the AccountClient type.
type AccountClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters batch.AccountCreateParameters) (result batch.AccountCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result batch.AccountDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result batch.Account, err error)
	GetKeys(ctx context.Context, resourceGroupName string, accountName string) (result batch.AccountKeys, err error)
	List(ctx context.Context) (result batch.AccountListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result batch.AccountListResultPage, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters batch.AccountRegenerateKeyParameters) (result batch.AccountKeys, err error)
	SynchronizeAutoStorageKeys(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters batch.AccountUpdateParameters) (result batch.Account, err error)
}

var _ AccountClientAPI = (*batch.AccountClient)(nil)

// ApplicationPackageClientAPI contains the set of methods on the ApplicationPackageClient type.
type ApplicationPackageClientAPI interface {
	Activate(ctx context.Context, resourceGroupName string, accountName string, applicationID string, version string, parameters batch.ActivateApplicationPackageParameters) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, applicationID string, version string) (result batch.ApplicationPackage, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, applicationID string, version string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, applicationID string, version string) (result batch.ApplicationPackage, err error)
}

var _ ApplicationPackageClientAPI = (*batch.ApplicationPackageClient)(nil)

// ApplicationClientAPI contains the set of methods on the ApplicationClient type.
type ApplicationClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, applicationID string, parameters *batch.ApplicationCreateParameters) (result batch.Application, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, applicationID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, applicationID string) (result batch.Application, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (result batch.ListApplicationsResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, applicationID string, parameters batch.ApplicationUpdateParameters) (result autorest.Response, err error)
}

var _ ApplicationClientAPI = (*batch.ApplicationClient)(nil)

// LocationClientAPI contains the set of methods on the LocationClient type.
type LocationClientAPI interface {
	CheckNameAvailability(ctx context.Context, locationName string, parameters batch.CheckNameAvailabilityParameters) (result batch.CheckNameAvailabilityResult, err error)
	GetQuotas(ctx context.Context, locationName string) (result batch.LocationQuota, err error)
}

var _ LocationClientAPI = (*batch.LocationClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result batch.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*batch.OperationsClient)(nil)

// CertificateClientAPI contains the set of methods on the CertificateClient type.
type CertificateClientAPI interface {
	CancelDeletion(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result batch.Certificate, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters batch.CertificateCreateOrUpdateParameters, ifMatch string, ifNoneMatch string) (result batch.CertificateCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result batch.CertificateDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result batch.Certificate, err error)
	ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result batch.ListCertificatesResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters batch.CertificateCreateOrUpdateParameters, ifMatch string) (result batch.Certificate, err error)
}

var _ CertificateClientAPI = (*batch.CertificateClient)(nil)

// PoolClientAPI contains the set of methods on the PoolClient type.
type PoolClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters batch.Pool, ifMatch string, ifNoneMatch string) (result batch.PoolCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result batch.PoolDeleteFuture, err error)
	DisableAutoScale(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result batch.Pool, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result batch.Pool, err error)
	ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result batch.ListPoolsResultPage, err error)
	StopResize(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result batch.Pool, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters batch.Pool, ifMatch string) (result batch.Pool, err error)
}

var _ PoolClientAPI = (*batch.PoolClient)(nil)

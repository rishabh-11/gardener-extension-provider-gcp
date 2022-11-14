// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package bastion

import (
	"context"
	"errors"
	"fmt"

	"github.com/gardener/gardener-extension-provider-gcp/pkg/apis/gcp"
	"github.com/gardener/gardener-extension-provider-gcp/pkg/apis/gcp/helper"
	gcpclient "github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client"

	"github.com/gardener/gardener/extensions/pkg/controller/bastion"
	"github.com/gardener/gardener/extensions/pkg/controller/common"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/extensions"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type configValidator struct {
	common.ClientContext
	gcpClientFactory gcpclient.Factory
	logger           logr.Logger
}

// NewConfigValidator creates a new ConfigValidator.
func NewConfigValidator(gcpClientFactory gcpclient.Factory, logger logr.Logger) bastion.ConfigValidator {
	return &configValidator{
		gcpClientFactory: gcpClientFactory,
		logger:           logger.WithName("gcp-bastion-config-validator"),
	}
}

// Validate validates the provider config of the given bastion resource with the cloud provider.
func (c *configValidator) Validate(ctx context.Context, bastion *extensionsv1alpha1.Bastion, cluster *extensions.Cluster) field.ErrorList {
	allErrs := field.ErrorList{}

	logger := c.logger.WithValues("bastion", client.ObjectKeyFromObject(bastion))

	infrastructureStatus, subnet, err := getInfrastructureStatus(ctx, c.Client(), cluster)
	if err != nil {
		allErrs = append(allErrs, field.InternalError(nil, err))
		return allErrs
	}

	secretReference := &corev1.SecretReference{
		Namespace: cluster.ObjectMeta.Name,
		Name:      v1beta1constants.SecretNameCloudProvider,
	}

	// Create GCP compute client
	computeClient, err := c.gcpClientFactory.NewComputeClient(ctx, c.Client(), *secretReference)
	if err != nil {
		allErrs = append(allErrs, field.InternalError(nil, err))
		return allErrs
	}

	// Validate bastion config
	logger.Info("Validating bastion configuration")
	allErrs = append(allErrs, c.validateInfrastructureStatus(ctx, computeClient, cluster.Shoot.Spec.Region, infrastructureStatus, subnet)...)

	return allErrs
}

func getInfrastructureStatus(ctx context.Context, c client.Client, cluster *extensions.Cluster) (*gcp.InfrastructureStatus, string, error) {
	var infrastructureStatus *gcp.InfrastructureStatus
	var nodeSubnet string

	worker := &extensionsv1alpha1.Worker{}

	err := c.Get(ctx, client.ObjectKey{Namespace: cluster.ObjectMeta.Name, Name: cluster.Shoot.Name}, worker)
	if err != nil {
		return nil, "", err
	}

	if worker == nil || worker.Spec.InfrastructureProviderStatus == nil {
		return nil, "", errors.New("infrastructure provider status must be not empty for worker")
	}

	if infrastructureStatus, err = helper.InfrastructureStatusFromRaw(worker.Spec.InfrastructureProviderStatus); err != nil {
		return nil, "", err
	}

	if infrastructureStatus.Networks.VPC.Name == "" {
		return nil, "", errors.New("virtual network must be not empty for infrastructure provider status")
	}

	if len(infrastructureStatus.Networks.Subnets) == 0 || infrastructureStatus.Networks.Subnets[0].Name == "" {
		return nil, "", errors.New("subnet must be not empty for infrastructure provider status")
	}

	for _, s := range infrastructureStatus.Networks.Subnets {
		if s.Purpose == gcp.PurposeNodes && s.Name != "" {
			nodeSubnet = s.Name
			break
		}
	}

	if nodeSubnet == "" {
		return nil, "", errors.New("could not get subnet purpose node from infrastructure status")
	}

	return infrastructureStatus, nodeSubnet, nil
}

func (c *configValidator) validateInfrastructureStatus(ctx context.Context, computeClient gcpclient.ComputeClient, region string, infrastructureStatus *gcp.InfrastructureStatus, nodeSubnet string) field.ErrorList {
	allErrs := field.ErrorList{}

	vpc, err := computeClient.GetVPC(ctx, infrastructureStatus.Networks.VPC.Name)
	if err != nil || vpc == nil || vpc.Name == "" {
		allErrs = append(allErrs, field.InternalError(field.NewPath("vpc"), fmt.Errorf("could not get vpc %s from gcp provider: %w", infrastructureStatus.Networks.VPC.Name, err)))
		return allErrs
	}

	subnet, err := computeClient.GetSubnet(ctx, region, nodeSubnet)
	if err != nil || subnet == nil || subnet.Name == "" {
		allErrs = append(allErrs, field.InternalError(field.NewPath("subnet"), fmt.Errorf("could not get subnet %s from gcp provider: %w", nodeSubnet, err)))
		return allErrs
	}

	return allErrs
}

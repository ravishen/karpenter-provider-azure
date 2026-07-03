/*
Portions Copyright (c) Microsoft Corporation.

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

package consts

// Condition reasons surfaced on the NodeClaim Launched condition when instance
// creation fails. Consumers (dashboards, alerts, automation) may branch on these
// machine-readable values.
const (
	NodeClassReadinessUnknownReason    = "NodeClassReadinessUnknown"
	InstanceTypeResolutionFailedReason = "InstanceTypeResolutionFailed"
	// CreateInstanceFailedReason is the generic fallback reason for instance-creation
	// failures that are not classified by the offerings error handlers below.
	CreateInstanceFailedReason = "CreateInstanceFailed"

	// Classified instance-creation failure reasons set by the offerings error handlers.
	SubscriptionQuotaReachedReason              = "SubscriptionQuotaReached"
	AllocationFailureReason                     = "AllocationFailure"
	ZonalAllocationFailureReason                = "ZonalAllocationFailure"
	OverconstrainedZonalAllocationFailureReason = "OverconstrainedZonalAllocationFailure"
	OverconstrainedAllocationFailureReason      = "OverconstrainedAllocationFailure"
	SKUNotAvailableReason                       = "SKUNotAvailable"
)

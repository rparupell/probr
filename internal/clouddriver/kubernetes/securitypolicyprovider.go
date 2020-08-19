package kubernetes

// SecurityPolicyProvider ...
type SecurityPolicyProvider interface {
	HasSecurityPolicies() (*bool, error)
	HasPrivilegedAccessRestriction() (*bool, error)
	HasHostPIDRestriction() (*bool, error)
	HasHostIPCRestriction() (*bool, error)
	HasHostNetworkRestriction() (*bool, error)
	HasAllowPrivilegeEscalationRestriction() (*bool, error)
	HasRootUserRestriction() (*bool, error)
	HasNETRAWRestriction() (*bool, error)
	HasAllowedCapabilitiesRestriction() (*bool, error)
	HasAssignedCapabilitiesRestriction() (*bool, error)
}
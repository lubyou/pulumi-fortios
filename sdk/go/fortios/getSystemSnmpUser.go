// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios systemsnmp user
func LookupSystemSnmpUser(ctx *pulumi.Context, args *LookupSystemSnmpUserArgs, opts ...pulumi.InvokeOption) (*LookupSystemSnmpUserResult, error) {
	var rv LookupSystemSnmpUserResult
	err := ctx.Invoke("fortios:index/getSystemSnmpUser:GetSystemSnmpUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemSnmpUser.
type LookupSystemSnmpUserArgs struct {
	// Specify the name of the desired systemsnmp user.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemSnmpUser.
type LookupSystemSnmpUserResult struct {
	// Authentication protocol.
	AuthProto string `pulumi:"authProto"`
	// Password for authentication protocol.
	AuthPwd string `pulumi:"authPwd"`
	// SNMP notifications (traps) to send.
	Events string `pulumi:"events"`
	// Enable/disable direct management of HA cluster members.
	HaDirect string `pulumi:"haDirect"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SNMP user name.
	Name string `pulumi:"name"`
	// SNMP managers to send notifications (traps) to.
	NotifyHosts string `pulumi:"notifyHosts"`
	// IPv6 SNMP managers to send notifications (traps) to.
	NotifyHosts6 string `pulumi:"notifyHosts6"`
	// Privacy (encryption) protocol.
	PrivProto string `pulumi:"privProto"`
	// Password for privacy (encryption) protocol.
	PrivPwd string `pulumi:"privPwd"`
	// Enable/disable SNMP queries for this user.
	Queries string `pulumi:"queries"`
	// SNMPv3 query port (default = 161).
	QueryPort int `pulumi:"queryPort"`
	// Security level for message authentication and encryption.
	SecurityLevel string `pulumi:"securityLevel"`
	// Source IP for SNMP trap.
	SourceIp string `pulumi:"sourceIp"`
	// Source IPv6 for SNMP trap.
	SourceIpv6 string `pulumi:"sourceIpv6"`
	// Enable/disable this SNMP user.
	Status string `pulumi:"status"`
	// SNMPv3 local trap port (default = 162).
	TrapLport int `pulumi:"trapLport"`
	// SNMPv3 trap remote port (default = 162).
	TrapRport int `pulumi:"trapRport"`
	// Enable/disable traps for this SNMP user.
	TrapStatus string  `pulumi:"trapStatus"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

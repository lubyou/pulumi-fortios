// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system resourcelimits
func LookupSystemResourceLimits(ctx *pulumi.Context, args *LookupSystemResourceLimitsArgs, opts ...pulumi.InvokeOption) (*LookupSystemResourceLimitsResult, error) {
	var rv LookupSystemResourceLimitsResult
	err := ctx.Invoke("fortios:index/getSystemResourceLimits:GetSystemResourceLimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemResourceLimits.
type LookupSystemResourceLimitsArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemResourceLimits.
type LookupSystemResourceLimitsResult struct {
	// Maximum number of firewall custom services.
	CustomService int `pulumi:"customService"`
	// Maximum number of dial-up tunnels.
	DialupTunnel int `pulumi:"dialupTunnel"`
	// Maximum number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress int `pulumi:"firewallAddress"`
	// Maximum number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp int `pulumi:"firewallAddrgrp"`
	// Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy int `pulumi:"firewallPolicy"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of VPN IPsec phase1 tunnels.
	IpsecPhase1 int `pulumi:"ipsecPhase1"`
	// Maximum number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface int `pulumi:"ipsecPhase1Interface"`
	// Maximum number of VPN IPsec phase2 tunnels.
	IpsecPhase2 int `pulumi:"ipsecPhase2"`
	// Maximum number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface int `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in MB.
	LogDiskQuota int `pulumi:"logDiskQuota"`
	// Maximum number of firewall one-time schedules.
	OnetimeSchedule int `pulumi:"onetimeSchedule"`
	// Maximum number of concurrent proxy users.
	Proxy int `pulumi:"proxy"`
	// Maximum number of firewall recurring schedules.
	RecurringSchedule int `pulumi:"recurringSchedule"`
	// Maximum number of firewall service groups.
	ServiceGroup int `pulumi:"serviceGroup"`
	// Maximum number of sessions.
	Session int `pulumi:"session"`
	// Maximum number of SSL-VPN.
	Sslvpn int `pulumi:"sslvpn"`
	// Maximum number of local users.
	User int `pulumi:"user"`
	// Maximum number of user groups.
	UserGroup int     `pulumi:"userGroup"`
	Vdomparam *string `pulumi:"vdomparam"`
}

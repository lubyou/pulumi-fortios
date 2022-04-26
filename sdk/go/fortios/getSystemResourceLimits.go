// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system resourcelimits
func LookupSystemResourceLimits(ctx *pulumi.Context, args *LookupSystemResourceLimitsArgs, opts ...pulumi.InvokeOption) (*LookupSystemResourceLimitsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
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

func LookupSystemResourceLimitsOutput(ctx *pulumi.Context, args LookupSystemResourceLimitsOutputArgs, opts ...pulumi.InvokeOption) LookupSystemResourceLimitsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemResourceLimitsResult, error) {
			args := v.(LookupSystemResourceLimitsArgs)
			r, err := LookupSystemResourceLimits(ctx, &args, opts...)
			var s LookupSystemResourceLimitsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemResourceLimitsResultOutput)
}

// A collection of arguments for invoking GetSystemResourceLimits.
type LookupSystemResourceLimitsOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemResourceLimitsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemResourceLimitsArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemResourceLimits.
type LookupSystemResourceLimitsResultOutput struct{ *pulumi.OutputState }

func (LookupSystemResourceLimitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemResourceLimitsResult)(nil)).Elem()
}

func (o LookupSystemResourceLimitsResultOutput) ToLookupSystemResourceLimitsResultOutput() LookupSystemResourceLimitsResultOutput {
	return o
}

func (o LookupSystemResourceLimitsResultOutput) ToLookupSystemResourceLimitsResultOutputWithContext(ctx context.Context) LookupSystemResourceLimitsResultOutput {
	return o
}

// Maximum number of firewall custom services.
func (o LookupSystemResourceLimitsResultOutput) CustomService() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.CustomService }).(pulumi.IntOutput)
}

// Maximum number of dial-up tunnels.
func (o LookupSystemResourceLimitsResultOutput) DialupTunnel() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.DialupTunnel }).(pulumi.IntOutput)
}

// Maximum number of firewall addresses (IPv4, IPv6, multicast).
func (o LookupSystemResourceLimitsResultOutput) FirewallAddress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.FirewallAddress }).(pulumi.IntOutput)
}

// Maximum number of firewall address groups (IPv4, IPv6).
func (o LookupSystemResourceLimitsResultOutput) FirewallAddrgrp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.FirewallAddrgrp }).(pulumi.IntOutput)
}

// Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
func (o LookupSystemResourceLimitsResultOutput) FirewallPolicy() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.FirewallPolicy }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemResourceLimitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of VPN IPsec phase1 tunnels.
func (o LookupSystemResourceLimitsResultOutput) IpsecPhase1() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.IpsecPhase1 }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase1 interface tunnels.
func (o LookupSystemResourceLimitsResultOutput) IpsecPhase1Interface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.IpsecPhase1Interface }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase2 tunnels.
func (o LookupSystemResourceLimitsResultOutput) IpsecPhase2() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.IpsecPhase2 }).(pulumi.IntOutput)
}

// Maximum number of VPN IPsec phase2 interface tunnels.
func (o LookupSystemResourceLimitsResultOutput) IpsecPhase2Interface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.IpsecPhase2Interface }).(pulumi.IntOutput)
}

// Log disk quota in MB.
func (o LookupSystemResourceLimitsResultOutput) LogDiskQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.LogDiskQuota }).(pulumi.IntOutput)
}

// Maximum number of firewall one-time schedules.
func (o LookupSystemResourceLimitsResultOutput) OnetimeSchedule() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.OnetimeSchedule }).(pulumi.IntOutput)
}

// Maximum number of concurrent proxy users.
func (o LookupSystemResourceLimitsResultOutput) Proxy() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.Proxy }).(pulumi.IntOutput)
}

// Maximum number of firewall recurring schedules.
func (o LookupSystemResourceLimitsResultOutput) RecurringSchedule() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.RecurringSchedule }).(pulumi.IntOutput)
}

// Maximum number of firewall service groups.
func (o LookupSystemResourceLimitsResultOutput) ServiceGroup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.ServiceGroup }).(pulumi.IntOutput)
}

// Maximum number of sessions.
func (o LookupSystemResourceLimitsResultOutput) Session() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.Session }).(pulumi.IntOutput)
}

// Maximum number of SSL-VPN.
func (o LookupSystemResourceLimitsResultOutput) Sslvpn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.Sslvpn }).(pulumi.IntOutput)
}

// Maximum number of local users.
func (o LookupSystemResourceLimitsResultOutput) User() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.User }).(pulumi.IntOutput)
}

// Maximum number of user groups.
func (o LookupSystemResourceLimitsResultOutput) UserGroup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) int { return v.UserGroup }).(pulumi.IntOutput)
}

func (o LookupSystemResourceLimitsResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemResourceLimitsResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemResourceLimitsResultOutput{})
}

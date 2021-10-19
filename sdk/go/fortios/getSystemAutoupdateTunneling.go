// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios systemautoupdate tunneling
func LookupSystemAutoupdateTunneling(ctx *pulumi.Context, args *LookupSystemAutoupdateTunnelingArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutoupdateTunnelingResult, error) {
	var rv LookupSystemAutoupdateTunnelingResult
	err := ctx.Invoke("fortios:index/getSystemAutoupdateTunneling:GetSystemAutoupdateTunneling", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutoupdateTunneling.
type LookupSystemAutoupdateTunnelingArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutoupdateTunneling.
type LookupSystemAutoupdateTunnelingResult struct {
	// Web proxy IP address or FQDN.
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Web proxy password.
	Password string `pulumi:"password"`
	// Web proxy port.
	Port int `pulumi:"port"`
	// Enable/disable web proxy tunnelling.
	Status string `pulumi:"status"`
	// Web proxy username.
	Username  string  `pulumi:"username"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system alias
func LookupSystemAlias(ctx *pulumi.Context, args *LookupSystemAliasArgs, opts ...pulumi.InvokeOption) (*LookupSystemAliasResult, error) {
	var rv LookupSystemAliasResult
	err := ctx.Invoke("fortios:index/getSystemAlias:GetSystemAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAlias.
type LookupSystemAliasArgs struct {
	// Specify the name of the desired system alias.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAlias.
type LookupSystemAliasResult struct {
	// Command list to execute.
	Command string `pulumi:"command"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Alias command name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}
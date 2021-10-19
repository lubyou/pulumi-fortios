// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemSnmpUser`.
func GetSystemSnmpUserList(ctx *pulumi.Context, args *GetSystemSnmpUserListArgs, opts ...pulumi.InvokeOption) (*GetSystemSnmpUserListResult, error) {
	var rv GetSystemSnmpUserListResult
	err := ctx.Invoke("fortios:index/getSystemSnmpUserList:GetSystemSnmpUserList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemSnmpUserList.
type GetSystemSnmpUserListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemSnmpUserList.
type GetSystemSnmpUserListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemSnmpUser`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

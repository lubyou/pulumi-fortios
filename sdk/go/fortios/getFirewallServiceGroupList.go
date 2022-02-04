// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallServiceGroup`.
func GetFirewallServiceGroupList(ctx *pulumi.Context, args *GetFirewallServiceGroupListArgs, opts ...pulumi.InvokeOption) (*GetFirewallServiceGroupListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallServiceGroupListResult
	err := ctx.Invoke("fortios:index/getFirewallServiceGroupList:GetFirewallServiceGroupList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceGroupList.
type GetFirewallServiceGroupListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceGroupList.
type GetFirewallServiceGroupListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallServiceGroup`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallServiceGroupListOutput(ctx *pulumi.Context, args GetFirewallServiceGroupListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallServiceGroupListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallServiceGroupListResult, error) {
			args := v.(GetFirewallServiceGroupListArgs)
			r, err := GetFirewallServiceGroupList(ctx, &args, opts...)
			return *r, err
		}).(GetFirewallServiceGroupListResultOutput)
}

// A collection of arguments for invoking GetFirewallServiceGroupList.
type GetFirewallServiceGroupListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallServiceGroupListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallServiceGroupListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallServiceGroupList.
type GetFirewallServiceGroupListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallServiceGroupListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallServiceGroupListResult)(nil)).Elem()
}

func (o GetFirewallServiceGroupListResultOutput) ToGetFirewallServiceGroupListResultOutput() GetFirewallServiceGroupListResultOutput {
	return o
}

func (o GetFirewallServiceGroupListResultOutput) ToGetFirewallServiceGroupListResultOutputWithContext(ctx context.Context) GetFirewallServiceGroupListResultOutput {
	return o
}

func (o GetFirewallServiceGroupListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallServiceGroupListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallServiceGroupListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallServiceGroupListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallServiceGroup`.
func (o GetFirewallServiceGroupListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallServiceGroupListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallServiceGroupListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallServiceGroupListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallServiceGroupListResultOutput{})
}

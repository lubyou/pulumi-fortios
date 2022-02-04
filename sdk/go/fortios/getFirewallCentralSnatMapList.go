// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallCentralSnatMap`.
func GetFirewallCentralSnatMapList(ctx *pulumi.Context, args *GetFirewallCentralSnatMapListArgs, opts ...pulumi.InvokeOption) (*GetFirewallCentralSnatMapListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallCentralSnatMapListResult
	err := ctx.Invoke("fortios:index/getFirewallCentralSnatMapList:GetFirewallCentralSnatMapList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallCentralSnatMapList.
type GetFirewallCentralSnatMapListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallCentralSnatMapList.
type GetFirewallCentralSnatMapListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallCentralSnatMap`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetFirewallCentralSnatMapListOutput(ctx *pulumi.Context, args GetFirewallCentralSnatMapListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallCentralSnatMapListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallCentralSnatMapListResult, error) {
			args := v.(GetFirewallCentralSnatMapListArgs)
			r, err := GetFirewallCentralSnatMapList(ctx, &args, opts...)
			return *r, err
		}).(GetFirewallCentralSnatMapListResultOutput)
}

// A collection of arguments for invoking GetFirewallCentralSnatMapList.
type GetFirewallCentralSnatMapListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallCentralSnatMapListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallCentralSnatMapListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallCentralSnatMapList.
type GetFirewallCentralSnatMapListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallCentralSnatMapListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallCentralSnatMapListResult)(nil)).Elem()
}

func (o GetFirewallCentralSnatMapListResultOutput) ToGetFirewallCentralSnatMapListResultOutput() GetFirewallCentralSnatMapListResultOutput {
	return o
}

func (o GetFirewallCentralSnatMapListResultOutput) ToGetFirewallCentralSnatMapListResultOutputWithContext(ctx context.Context) GetFirewallCentralSnatMapListResultOutput {
	return o
}

func (o GetFirewallCentralSnatMapListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallCentralSnatMapListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallCentralSnatMapListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallCentralSnatMapListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallCentralSnatMap`.
func (o GetFirewallCentralSnatMapListResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetFirewallCentralSnatMapListResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetFirewallCentralSnatMapListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallCentralSnatMapListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallCentralSnatMapListResultOutput{})
}

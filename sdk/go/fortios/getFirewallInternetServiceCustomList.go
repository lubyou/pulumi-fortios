// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallInternetServiceCustom`.
func GetFirewallInternetServiceCustomList(ctx *pulumi.Context, args *GetFirewallInternetServiceCustomListArgs, opts ...pulumi.InvokeOption) (*GetFirewallInternetServiceCustomListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallInternetServiceCustomListResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceCustomList:GetFirewallInternetServiceCustomList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceCustomList.
type GetFirewallInternetServiceCustomListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceCustomList.
type GetFirewallInternetServiceCustomListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallInternetServiceCustom`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallInternetServiceCustomListOutput(ctx *pulumi.Context, args GetFirewallInternetServiceCustomListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallInternetServiceCustomListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallInternetServiceCustomListResult, error) {
			args := v.(GetFirewallInternetServiceCustomListArgs)
			r, err := GetFirewallInternetServiceCustomList(ctx, &args, opts...)
			return *r, err
		}).(GetFirewallInternetServiceCustomListResultOutput)
}

// A collection of arguments for invoking GetFirewallInternetServiceCustomList.
type GetFirewallInternetServiceCustomListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallInternetServiceCustomListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetServiceCustomListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallInternetServiceCustomList.
type GetFirewallInternetServiceCustomListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallInternetServiceCustomListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetServiceCustomListResult)(nil)).Elem()
}

func (o GetFirewallInternetServiceCustomListResultOutput) ToGetFirewallInternetServiceCustomListResultOutput() GetFirewallInternetServiceCustomListResultOutput {
	return o
}

func (o GetFirewallInternetServiceCustomListResultOutput) ToGetFirewallInternetServiceCustomListResultOutputWithContext(ctx context.Context) GetFirewallInternetServiceCustomListResultOutput {
	return o
}

func (o GetFirewallInternetServiceCustomListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetServiceCustomListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallInternetServiceCustomListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallInternetServiceCustomListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallInternetServiceCustom`.
func (o GetFirewallInternetServiceCustomListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallInternetServiceCustomListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallInternetServiceCustomListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetServiceCustomListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallInternetServiceCustomListResultOutput{})
}

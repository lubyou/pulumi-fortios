// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func GetFirewallShaperPerIpShaperList(ctx *pulumi.Context, args *GetFirewallShaperPerIpShaperListArgs, opts ...pulumi.InvokeOption) (*GetFirewallShaperPerIpShaperListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallShaperPerIpShaperListResult
	err := ctx.Invoke("fortios:index/getFirewallShaperPerIpShaperList:GetFirewallShaperPerIpShaperList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallShaperPerIpShaperList.
type GetFirewallShaperPerIpShaperListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallShaperPerIpShaperList.
type GetFirewallShaperPerIpShaperListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallShaperPerIpShaperListOutput(ctx *pulumi.Context, args GetFirewallShaperPerIpShaperListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallShaperPerIpShaperListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallShaperPerIpShaperListResult, error) {
			args := v.(GetFirewallShaperPerIpShaperListArgs)
			r, err := GetFirewallShaperPerIpShaperList(ctx, &args, opts...)
			var s GetFirewallShaperPerIpShaperListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallShaperPerIpShaperListResultOutput)
}

// A collection of arguments for invoking GetFirewallShaperPerIpShaperList.
type GetFirewallShaperPerIpShaperListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallShaperPerIpShaperListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallShaperPerIpShaperListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallShaperPerIpShaperList.
type GetFirewallShaperPerIpShaperListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallShaperPerIpShaperListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallShaperPerIpShaperListResult)(nil)).Elem()
}

func (o GetFirewallShaperPerIpShaperListResultOutput) ToGetFirewallShaperPerIpShaperListResultOutput() GetFirewallShaperPerIpShaperListResultOutput {
	return o
}

func (o GetFirewallShaperPerIpShaperListResultOutput) ToGetFirewallShaperPerIpShaperListResultOutputWithContext(ctx context.Context) GetFirewallShaperPerIpShaperListResultOutput {
	return o
}

func (o GetFirewallShaperPerIpShaperListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetFirewallShaperPerIpShaperListResult] {
	return pulumix.Output[GetFirewallShaperPerIpShaperListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetFirewallShaperPerIpShaperListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallShaperPerIpShaperListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallShaperPerIpShaperListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallShaperPerIpShaperListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallShaperPerIpShaperListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallShaperPerIpShaperListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallShaperPerIpShaperListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallShaperPerIpShaperListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallShaperPerIpShaperListResultOutput{})
}

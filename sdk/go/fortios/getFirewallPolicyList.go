// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallPolicy`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sample1, err := fortios.GetFirewallPolicyList(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("output1", sample1)
// 		sample2, err := fortios.GetFirewallPolicyList(ctx, &GetFirewallPolicyListArgs{
// 			Filter: pulumi.StringRef("schedule==always&action==accept,action==deny"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("sample2Output", sample2.Policyidlists)
// 		return nil
// 	})
// }
// ```
func GetFirewallPolicyList(ctx *pulumi.Context, args *GetFirewallPolicyListArgs, opts ...pulumi.InvokeOption) (*GetFirewallPolicyListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallPolicyListResult
	err := ctx.Invoke("fortios:index/getFirewallPolicyList:GetFirewallPolicyList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallPolicyList.
type GetFirewallPolicyListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallPolicyList.
type GetFirewallPolicyListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallPolicy`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetFirewallPolicyListOutput(ctx *pulumi.Context, args GetFirewallPolicyListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallPolicyListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallPolicyListResult, error) {
			args := v.(GetFirewallPolicyListArgs)
			r, err := GetFirewallPolicyList(ctx, &args, opts...)
			return *r, err
		}).(GetFirewallPolicyListResultOutput)
}

// A collection of arguments for invoking GetFirewallPolicyList.
type GetFirewallPolicyListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallPolicyListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicyListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallPolicyList.
type GetFirewallPolicyListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallPolicyListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicyListResult)(nil)).Elem()
}

func (o GetFirewallPolicyListResultOutput) ToGetFirewallPolicyListResultOutput() GetFirewallPolicyListResultOutput {
	return o
}

func (o GetFirewallPolicyListResultOutput) ToGetFirewallPolicyListResultOutputWithContext(ctx context.Context) GetFirewallPolicyListResultOutput {
	return o
}

func (o GetFirewallPolicyListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicyListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallPolicyListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallPolicyListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallPolicy`.
func (o GetFirewallPolicyListResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetFirewallPolicyListResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetFirewallPolicyListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicyListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallPolicyListResultOutput{})
}

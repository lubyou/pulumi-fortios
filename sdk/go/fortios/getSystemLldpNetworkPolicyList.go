// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemLldpNetworkPolicyList(ctx *pulumi.Context, args *GetSystemLldpNetworkPolicyListArgs, opts ...pulumi.InvokeOption) (*GetSystemLldpNetworkPolicyListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemLldpNetworkPolicyListResult
	err := ctx.Invoke("fortios:index/getSystemLldpNetworkPolicyList:GetSystemLldpNetworkPolicyList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemLldpNetworkPolicyListOutput(ctx *pulumi.Context, args GetSystemLldpNetworkPolicyListOutputArgs, opts ...pulumi.InvokeOption) GetSystemLldpNetworkPolicyListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemLldpNetworkPolicyListResult, error) {
			args := v.(GetSystemLldpNetworkPolicyListArgs)
			r, err := GetSystemLldpNetworkPolicyList(ctx, &args, opts...)
			var s GetSystemLldpNetworkPolicyListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemLldpNetworkPolicyListResultOutput)
}

// A collection of arguments for invoking GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemLldpNetworkPolicyListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemLldpNetworkPolicyListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListResultOutput struct{ *pulumi.OutputState }

func (GetSystemLldpNetworkPolicyListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemLldpNetworkPolicyListResult)(nil)).Elem()
}

func (o GetSystemLldpNetworkPolicyListResultOutput) ToGetSystemLldpNetworkPolicyListResultOutput() GetSystemLldpNetworkPolicyListResultOutput {
	return o
}

func (o GetSystemLldpNetworkPolicyListResultOutput) ToGetSystemLldpNetworkPolicyListResultOutputWithContext(ctx context.Context) GetSystemLldpNetworkPolicyListResultOutput {
	return o
}

func (o GetSystemLldpNetworkPolicyListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemLldpNetworkPolicyListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemLldpNetworkPolicyListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemLldpNetworkPolicyListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemLldpNetworkPolicyListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemLldpNetworkPolicyListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemLldpNetworkPolicyListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemLldpNetworkPolicyListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemLldpNetworkPolicyListResultOutput{})
}

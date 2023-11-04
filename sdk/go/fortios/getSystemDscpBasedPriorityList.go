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

func GetSystemDscpBasedPriorityList(ctx *pulumi.Context, args *GetSystemDscpBasedPriorityListArgs, opts ...pulumi.InvokeOption) (*GetSystemDscpBasedPriorityListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemDscpBasedPriorityListResult
	err := ctx.Invoke("fortios:index/getSystemDscpBasedPriorityList:GetSystemDscpBasedPriorityList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDscpBasedPriorityList.
type GetSystemDscpBasedPriorityListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDscpBasedPriorityList.
type GetSystemDscpBasedPriorityListResult struct {
	Filter     *string `pulumi:"filter"`
	Fosidlists []int   `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemDscpBasedPriorityListOutput(ctx *pulumi.Context, args GetSystemDscpBasedPriorityListOutputArgs, opts ...pulumi.InvokeOption) GetSystemDscpBasedPriorityListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemDscpBasedPriorityListResult, error) {
			args := v.(GetSystemDscpBasedPriorityListArgs)
			r, err := GetSystemDscpBasedPriorityList(ctx, &args, opts...)
			var s GetSystemDscpBasedPriorityListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemDscpBasedPriorityListResultOutput)
}

// A collection of arguments for invoking GetSystemDscpBasedPriorityList.
type GetSystemDscpBasedPriorityListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemDscpBasedPriorityListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDscpBasedPriorityListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDscpBasedPriorityList.
type GetSystemDscpBasedPriorityListResultOutput struct{ *pulumi.OutputState }

func (GetSystemDscpBasedPriorityListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDscpBasedPriorityListResult)(nil)).Elem()
}

func (o GetSystemDscpBasedPriorityListResultOutput) ToGetSystemDscpBasedPriorityListResultOutput() GetSystemDscpBasedPriorityListResultOutput {
	return o
}

func (o GetSystemDscpBasedPriorityListResultOutput) ToGetSystemDscpBasedPriorityListResultOutputWithContext(ctx context.Context) GetSystemDscpBasedPriorityListResultOutput {
	return o
}

func (o GetSystemDscpBasedPriorityListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSystemDscpBasedPriorityListResult] {
	return pulumix.Output[GetSystemDscpBasedPriorityListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetSystemDscpBasedPriorityListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDscpBasedPriorityListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetSystemDscpBasedPriorityListResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemDscpBasedPriorityListResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemDscpBasedPriorityListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemDscpBasedPriorityListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemDscpBasedPriorityListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDscpBasedPriorityListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemDscpBasedPriorityListResultOutput{})
}

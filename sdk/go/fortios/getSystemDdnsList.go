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

func GetSystemDdnsList(ctx *pulumi.Context, args *GetSystemDdnsListArgs, opts ...pulumi.InvokeOption) (*GetSystemDdnsListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemDdnsListResult
	err := ctx.Invoke("fortios:index/getSystemDdnsList:GetSystemDdnsList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDdnsList.
type GetSystemDdnsListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDdnsList.
type GetSystemDdnsListResult struct {
	Ddnsidlists []int   `pulumi:"ddnsidlists"`
	Filter      *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemDdnsListOutput(ctx *pulumi.Context, args GetSystemDdnsListOutputArgs, opts ...pulumi.InvokeOption) GetSystemDdnsListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemDdnsListResult, error) {
			args := v.(GetSystemDdnsListArgs)
			r, err := GetSystemDdnsList(ctx, &args, opts...)
			var s GetSystemDdnsListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemDdnsListResultOutput)
}

// A collection of arguments for invoking GetSystemDdnsList.
type GetSystemDdnsListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemDdnsListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDdnsListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDdnsList.
type GetSystemDdnsListResultOutput struct{ *pulumi.OutputState }

func (GetSystemDdnsListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDdnsListResult)(nil)).Elem()
}

func (o GetSystemDdnsListResultOutput) ToGetSystemDdnsListResultOutput() GetSystemDdnsListResultOutput {
	return o
}

func (o GetSystemDdnsListResultOutput) ToGetSystemDdnsListResultOutputWithContext(ctx context.Context) GetSystemDdnsListResultOutput {
	return o
}

func (o GetSystemDdnsListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSystemDdnsListResult] {
	return pulumix.Output[GetSystemDdnsListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetSystemDdnsListResultOutput) Ddnsidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemDdnsListResult) []int { return v.Ddnsidlists }).(pulumi.IntArrayOutput)
}

func (o GetSystemDdnsListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDdnsListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemDdnsListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemDdnsListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemDdnsListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDdnsListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemDdnsListResultOutput{})
}

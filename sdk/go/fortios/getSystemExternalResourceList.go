// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemExternalResourceList(ctx *pulumi.Context, args *GetSystemExternalResourceListArgs, opts ...pulumi.InvokeOption) (*GetSystemExternalResourceListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemExternalResourceListResult
	err := ctx.Invoke("fortios:index/getSystemExternalResourceList:GetSystemExternalResourceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemExternalResourceList.
type GetSystemExternalResourceListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemExternalResourceList.
type GetSystemExternalResourceListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemExternalResourceListOutput(ctx *pulumi.Context, args GetSystemExternalResourceListOutputArgs, opts ...pulumi.InvokeOption) GetSystemExternalResourceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemExternalResourceListResult, error) {
			args := v.(GetSystemExternalResourceListArgs)
			r, err := GetSystemExternalResourceList(ctx, &args, opts...)
			var s GetSystemExternalResourceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemExternalResourceListResultOutput)
}

// A collection of arguments for invoking GetSystemExternalResourceList.
type GetSystemExternalResourceListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemExternalResourceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemExternalResourceListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemExternalResourceList.
type GetSystemExternalResourceListResultOutput struct{ *pulumi.OutputState }

func (GetSystemExternalResourceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemExternalResourceListResult)(nil)).Elem()
}

func (o GetSystemExternalResourceListResultOutput) ToGetSystemExternalResourceListResultOutput() GetSystemExternalResourceListResultOutput {
	return o
}

func (o GetSystemExternalResourceListResultOutput) ToGetSystemExternalResourceListResultOutputWithContext(ctx context.Context) GetSystemExternalResourceListResultOutput {
	return o
}

func (o GetSystemExternalResourceListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemExternalResourceListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemExternalResourceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemExternalResourceListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemExternalResourceListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemExternalResourceListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemExternalResourceListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemExternalResourceListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemExternalResourceListResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemWccpList(ctx *pulumi.Context, args *GetSystemWccpListArgs, opts ...pulumi.InvokeOption) (*GetSystemWccpListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemWccpListResult
	err := ctx.Invoke("fortios:index/getSystemWccpList:GetSystemWccpList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemWccpList.
type GetSystemWccpListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemWccpList.
type GetSystemWccpListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	ServiceIdlists []string `pulumi:"serviceIdlists"`
	Vdomparam      *string  `pulumi:"vdomparam"`
}

func GetSystemWccpListOutput(ctx *pulumi.Context, args GetSystemWccpListOutputArgs, opts ...pulumi.InvokeOption) GetSystemWccpListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemWccpListResult, error) {
			args := v.(GetSystemWccpListArgs)
			r, err := GetSystemWccpList(ctx, &args, opts...)
			var s GetSystemWccpListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemWccpListResultOutput)
}

// A collection of arguments for invoking GetSystemWccpList.
type GetSystemWccpListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemWccpListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemWccpListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemWccpList.
type GetSystemWccpListResultOutput struct{ *pulumi.OutputState }

func (GetSystemWccpListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemWccpListResult)(nil)).Elem()
}

func (o GetSystemWccpListResultOutput) ToGetSystemWccpListResultOutput() GetSystemWccpListResultOutput {
	return o
}

func (o GetSystemWccpListResultOutput) ToGetSystemWccpListResultOutputWithContext(ctx context.Context) GetSystemWccpListResultOutput {
	return o
}

func (o GetSystemWccpListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemWccpListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemWccpListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemWccpListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemWccpListResultOutput) ServiceIdlists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemWccpListResult) []string { return v.ServiceIdlists }).(pulumi.StringArrayOutput)
}

func (o GetSystemWccpListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemWccpListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemWccpListResultOutput{})
}

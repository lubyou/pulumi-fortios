// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemProxyArpList(ctx *pulumi.Context, args *GetSystemProxyArpListArgs, opts ...pulumi.InvokeOption) (*GetSystemProxyArpListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemProxyArpListResult
	err := ctx.Invoke("fortios:index/getSystemProxyArpList:GetSystemProxyArpList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemProxyArpList.
type GetSystemProxyArpListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemProxyArpList.
type GetSystemProxyArpListResult struct {
	Filter     *string `pulumi:"filter"`
	Fosidlists []int   `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemProxyArpListOutput(ctx *pulumi.Context, args GetSystemProxyArpListOutputArgs, opts ...pulumi.InvokeOption) GetSystemProxyArpListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemProxyArpListResult, error) {
			args := v.(GetSystemProxyArpListArgs)
			r, err := GetSystemProxyArpList(ctx, &args, opts...)
			var s GetSystemProxyArpListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemProxyArpListResultOutput)
}

// A collection of arguments for invoking GetSystemProxyArpList.
type GetSystemProxyArpListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemProxyArpListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemProxyArpListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemProxyArpList.
type GetSystemProxyArpListResultOutput struct{ *pulumi.OutputState }

func (GetSystemProxyArpListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemProxyArpListResult)(nil)).Elem()
}

func (o GetSystemProxyArpListResultOutput) ToGetSystemProxyArpListResultOutput() GetSystemProxyArpListResultOutput {
	return o
}

func (o GetSystemProxyArpListResultOutput) ToGetSystemProxyArpListResultOutputWithContext(ctx context.Context) GetSystemProxyArpListResultOutput {
	return o
}

func (o GetSystemProxyArpListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemProxyArpListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetSystemProxyArpListResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemProxyArpListResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemProxyArpListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemProxyArpListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemProxyArpListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemProxyArpListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemProxyArpListResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemIpv6NeighborCacheList(ctx *pulumi.Context, args *GetSystemIpv6NeighborCacheListArgs, opts ...pulumi.InvokeOption) (*GetSystemIpv6NeighborCacheListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemIpv6NeighborCacheListResult
	err := ctx.Invoke("fortios:index/getSystemIpv6NeighborCacheList:GetSystemIpv6NeighborCacheList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemIpv6NeighborCacheList.
type GetSystemIpv6NeighborCacheListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemIpv6NeighborCacheList.
type GetSystemIpv6NeighborCacheListResult struct {
	Filter     *string `pulumi:"filter"`
	Fosidlists []int   `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemIpv6NeighborCacheListOutput(ctx *pulumi.Context, args GetSystemIpv6NeighborCacheListOutputArgs, opts ...pulumi.InvokeOption) GetSystemIpv6NeighborCacheListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemIpv6NeighborCacheListResult, error) {
			args := v.(GetSystemIpv6NeighborCacheListArgs)
			r, err := GetSystemIpv6NeighborCacheList(ctx, &args, opts...)
			var s GetSystemIpv6NeighborCacheListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemIpv6NeighborCacheListResultOutput)
}

// A collection of arguments for invoking GetSystemIpv6NeighborCacheList.
type GetSystemIpv6NeighborCacheListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemIpv6NeighborCacheListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemIpv6NeighborCacheListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemIpv6NeighborCacheList.
type GetSystemIpv6NeighborCacheListResultOutput struct{ *pulumi.OutputState }

func (GetSystemIpv6NeighborCacheListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemIpv6NeighborCacheListResult)(nil)).Elem()
}

func (o GetSystemIpv6NeighborCacheListResultOutput) ToGetSystemIpv6NeighborCacheListResultOutput() GetSystemIpv6NeighborCacheListResultOutput {
	return o
}

func (o GetSystemIpv6NeighborCacheListResultOutput) ToGetSystemIpv6NeighborCacheListResultOutputWithContext(ctx context.Context) GetSystemIpv6NeighborCacheListResultOutput {
	return o
}

func (o GetSystemIpv6NeighborCacheListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemIpv6NeighborCacheListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetSystemIpv6NeighborCacheListResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemIpv6NeighborCacheListResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemIpv6NeighborCacheListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemIpv6NeighborCacheListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemIpv6NeighborCacheListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemIpv6NeighborCacheListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemIpv6NeighborCacheListResultOutput{})
}

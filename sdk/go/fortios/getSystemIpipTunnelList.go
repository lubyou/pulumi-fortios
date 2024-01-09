// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemIpipTunnelList(ctx *pulumi.Context, args *GetSystemIpipTunnelListArgs, opts ...pulumi.InvokeOption) (*GetSystemIpipTunnelListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemIpipTunnelListResult
	err := ctx.Invoke("fortios:index/getSystemIpipTunnelList:GetSystemIpipTunnelList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemIpipTunnelList.
type GetSystemIpipTunnelListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemIpipTunnelList.
type GetSystemIpipTunnelListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemIpipTunnelListOutput(ctx *pulumi.Context, args GetSystemIpipTunnelListOutputArgs, opts ...pulumi.InvokeOption) GetSystemIpipTunnelListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemIpipTunnelListResult, error) {
			args := v.(GetSystemIpipTunnelListArgs)
			r, err := GetSystemIpipTunnelList(ctx, &args, opts...)
			var s GetSystemIpipTunnelListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemIpipTunnelListResultOutput)
}

// A collection of arguments for invoking GetSystemIpipTunnelList.
type GetSystemIpipTunnelListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemIpipTunnelListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemIpipTunnelListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemIpipTunnelList.
type GetSystemIpipTunnelListResultOutput struct{ *pulumi.OutputState }

func (GetSystemIpipTunnelListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemIpipTunnelListResult)(nil)).Elem()
}

func (o GetSystemIpipTunnelListResultOutput) ToGetSystemIpipTunnelListResultOutput() GetSystemIpipTunnelListResultOutput {
	return o
}

func (o GetSystemIpipTunnelListResultOutput) ToGetSystemIpipTunnelListResultOutputWithContext(ctx context.Context) GetSystemIpipTunnelListResultOutput {
	return o
}

func (o GetSystemIpipTunnelListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemIpipTunnelListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemIpipTunnelListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemIpipTunnelListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemIpipTunnelListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemIpipTunnelListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemIpipTunnelListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemIpipTunnelListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemIpipTunnelListResultOutput{})
}

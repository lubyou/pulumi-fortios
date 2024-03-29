// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemGreTunnelList(ctx *pulumi.Context, args *GetSystemGreTunnelListArgs, opts ...pulumi.InvokeOption) (*GetSystemGreTunnelListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemGreTunnelListResult
	err := ctx.Invoke("fortios:index/getSystemGreTunnelList:GetSystemGreTunnelList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemGreTunnelList.
type GetSystemGreTunnelListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemGreTunnelList.
type GetSystemGreTunnelListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemGreTunnelListOutput(ctx *pulumi.Context, args GetSystemGreTunnelListOutputArgs, opts ...pulumi.InvokeOption) GetSystemGreTunnelListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemGreTunnelListResult, error) {
			args := v.(GetSystemGreTunnelListArgs)
			r, err := GetSystemGreTunnelList(ctx, &args, opts...)
			var s GetSystemGreTunnelListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemGreTunnelListResultOutput)
}

// A collection of arguments for invoking GetSystemGreTunnelList.
type GetSystemGreTunnelListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemGreTunnelListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemGreTunnelListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemGreTunnelList.
type GetSystemGreTunnelListResultOutput struct{ *pulumi.OutputState }

func (GetSystemGreTunnelListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemGreTunnelListResult)(nil)).Elem()
}

func (o GetSystemGreTunnelListResultOutput) ToGetSystemGreTunnelListResultOutput() GetSystemGreTunnelListResultOutput {
	return o
}

func (o GetSystemGreTunnelListResultOutput) ToGetSystemGreTunnelListResultOutputWithContext(ctx context.Context) GetSystemGreTunnelListResultOutput {
	return o
}

func (o GetSystemGreTunnelListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemGreTunnelListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemGreTunnelListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemGreTunnelListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemGreTunnelListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemGreTunnelListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemGreTunnelListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemGreTunnelListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemGreTunnelListResultOutput{})
}

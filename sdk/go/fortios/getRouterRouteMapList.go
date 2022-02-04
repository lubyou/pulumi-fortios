// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterRouteMap`.
func GetRouterRouteMapList(ctx *pulumi.Context, args *GetRouterRouteMapListArgs, opts ...pulumi.InvokeOption) (*GetRouterRouteMapListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterRouteMapListResult
	err := ctx.Invoke("fortios:index/getRouterRouteMapList:GetRouterRouteMapList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterRouteMapList.
type GetRouterRouteMapListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterRouteMapList.
type GetRouterRouteMapListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterRouteMap`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterRouteMapListOutput(ctx *pulumi.Context, args GetRouterRouteMapListOutputArgs, opts ...pulumi.InvokeOption) GetRouterRouteMapListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterRouteMapListResult, error) {
			args := v.(GetRouterRouteMapListArgs)
			r, err := GetRouterRouteMapList(ctx, &args, opts...)
			return *r, err
		}).(GetRouterRouteMapListResultOutput)
}

// A collection of arguments for invoking GetRouterRouteMapList.
type GetRouterRouteMapListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterRouteMapListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterRouteMapListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterRouteMapList.
type GetRouterRouteMapListResultOutput struct{ *pulumi.OutputState }

func (GetRouterRouteMapListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterRouteMapListResult)(nil)).Elem()
}

func (o GetRouterRouteMapListResultOutput) ToGetRouterRouteMapListResultOutput() GetRouterRouteMapListResultOutput {
	return o
}

func (o GetRouterRouteMapListResultOutput) ToGetRouterRouteMapListResultOutputWithContext(ctx context.Context) GetRouterRouteMapListResultOutput {
	return o
}

func (o GetRouterRouteMapListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterRouteMapListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterRouteMapListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterRouteMapListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `RouterRouteMap`.
func (o GetRouterRouteMapListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterRouteMapListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterRouteMapListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterRouteMapListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterRouteMapListResultOutput{})
}

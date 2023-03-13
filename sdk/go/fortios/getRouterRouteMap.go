// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouterRouteMap(ctx *pulumi.Context, args *LookupRouterRouteMapArgs, opts ...pulumi.InvokeOption) (*LookupRouterRouteMapResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterRouteMapResult
	err := ctx.Invoke("fortios:index/getRouterRouteMap:GetRouterRouteMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterRouteMap.
type LookupRouterRouteMapArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterRouteMap.
type LookupRouterRouteMapResult struct {
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                  `pulumi:"id"`
	Name      string                  `pulumi:"name"`
	Rules     []GetRouterRouteMapRule `pulumi:"rules"`
	Vdomparam *string                 `pulumi:"vdomparam"`
}

func LookupRouterRouteMapOutput(ctx *pulumi.Context, args LookupRouterRouteMapOutputArgs, opts ...pulumi.InvokeOption) LookupRouterRouteMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterRouteMapResult, error) {
			args := v.(LookupRouterRouteMapArgs)
			r, err := LookupRouterRouteMap(ctx, &args, opts...)
			var s LookupRouterRouteMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterRouteMapResultOutput)
}

// A collection of arguments for invoking GetRouterRouteMap.
type LookupRouterRouteMapOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterRouteMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRouteMapArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterRouteMap.
type LookupRouterRouteMapResultOutput struct{ *pulumi.OutputState }

func (LookupRouterRouteMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRouteMapResult)(nil)).Elem()
}

func (o LookupRouterRouteMapResultOutput) ToLookupRouterRouteMapResultOutput() LookupRouterRouteMapResultOutput {
	return o
}

func (o LookupRouterRouteMapResultOutput) ToLookupRouterRouteMapResultOutputWithContext(ctx context.Context) LookupRouterRouteMapResultOutput {
	return o
}

func (o LookupRouterRouteMapResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRouteMapResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterRouteMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRouteMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterRouteMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRouteMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouterRouteMapResultOutput) Rules() GetRouterRouteMapRuleArrayOutput {
	return o.ApplyT(func(v LookupRouterRouteMapResult) []GetRouterRouteMapRule { return v.Rules }).(GetRouterRouteMapRuleArrayOutput)
}

func (o LookupRouterRouteMapResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterRouteMapResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterRouteMapResultOutput{})
}

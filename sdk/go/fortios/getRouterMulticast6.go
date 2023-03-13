// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouterMulticast6(ctx *pulumi.Context, args *LookupRouterMulticast6Args, opts ...pulumi.InvokeOption) (*LookupRouterMulticast6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterMulticast6Result
	err := ctx.Invoke("fortios:index/getRouterMulticast6:GetRouterMulticast6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterMulticast6.
type LookupRouterMulticast6Args struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterMulticast6.
type LookupRouterMulticast6Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                           `pulumi:"id"`
	Interfaces       []GetRouterMulticast6Interface   `pulumi:"interfaces"`
	MulticastPmtu    string                           `pulumi:"multicastPmtu"`
	MulticastRouting string                           `pulumi:"multicastRouting"`
	PimSmGlobals     []GetRouterMulticast6PimSmGlobal `pulumi:"pimSmGlobals"`
	Vdomparam        *string                          `pulumi:"vdomparam"`
}

func LookupRouterMulticast6Output(ctx *pulumi.Context, args LookupRouterMulticast6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterMulticast6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterMulticast6Result, error) {
			args := v.(LookupRouterMulticast6Args)
			r, err := LookupRouterMulticast6(ctx, &args, opts...)
			var s LookupRouterMulticast6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterMulticast6ResultOutput)
}

// A collection of arguments for invoking GetRouterMulticast6.
type LookupRouterMulticast6OutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterMulticast6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterMulticast6Args)(nil)).Elem()
}

// A collection of values returned by GetRouterMulticast6.
type LookupRouterMulticast6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterMulticast6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterMulticast6Result)(nil)).Elem()
}

func (o LookupRouterMulticast6ResultOutput) ToLookupRouterMulticast6ResultOutput() LookupRouterMulticast6ResultOutput {
	return o
}

func (o LookupRouterMulticast6ResultOutput) ToLookupRouterMulticast6ResultOutputWithContext(ctx context.Context) LookupRouterMulticast6ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterMulticast6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterMulticast6ResultOutput) Interfaces() GetRouterMulticast6InterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) []GetRouterMulticast6Interface { return v.Interfaces }).(GetRouterMulticast6InterfaceArrayOutput)
}

func (o LookupRouterMulticast6ResultOutput) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.MulticastPmtu }).(pulumi.StringOutput)
}

func (o LookupRouterMulticast6ResultOutput) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.MulticastRouting }).(pulumi.StringOutput)
}

func (o LookupRouterMulticast6ResultOutput) PimSmGlobals() GetRouterMulticast6PimSmGlobalArrayOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) []GetRouterMulticast6PimSmGlobal { return v.PimSmGlobals }).(GetRouterMulticast6PimSmGlobalArrayOutput)
}

func (o LookupRouterMulticast6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterMulticast6ResultOutput{})
}

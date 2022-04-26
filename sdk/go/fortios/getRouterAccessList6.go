// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router accesslist6
func LookupRouterAccessList6(ctx *pulumi.Context, args *LookupRouterAccessList6Args, opts ...pulumi.InvokeOption) (*LookupRouterAccessList6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterAccessList6Result
	err := ctx.Invoke("fortios:index/getRouterAccessList6:GetRouterAccessList6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAccessList6.
type LookupRouterAccessList6Args struct {
	// Specify the name of the desired router accesslist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAccessList6.
type LookupRouterAccessList6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRouterAccessList6Rule `pulumi:"rules"`
	Vdomparam *string                    `pulumi:"vdomparam"`
}

func LookupRouterAccessList6Output(ctx *pulumi.Context, args LookupRouterAccessList6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterAccessList6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterAccessList6Result, error) {
			args := v.(LookupRouterAccessList6Args)
			r, err := LookupRouterAccessList6(ctx, &args, opts...)
			var s LookupRouterAccessList6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterAccessList6ResultOutput)
}

// A collection of arguments for invoking GetRouterAccessList6.
type LookupRouterAccessList6OutputArgs struct {
	// Specify the name of the desired router accesslist6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterAccessList6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAccessList6Args)(nil)).Elem()
}

// A collection of values returned by GetRouterAccessList6.
type LookupRouterAccessList6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterAccessList6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAccessList6Result)(nil)).Elem()
}

func (o LookupRouterAccessList6ResultOutput) ToLookupRouterAccessList6ResultOutput() LookupRouterAccessList6ResultOutput {
	return o
}

func (o LookupRouterAccessList6ResultOutput) ToLookupRouterAccessList6ResultOutputWithContext(ctx context.Context) LookupRouterAccessList6ResultOutput {
	return o
}

// Comment.
func (o LookupRouterAccessList6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessList6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterAccessList6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessList6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRouterAccessList6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessList6Result) string { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o LookupRouterAccessList6ResultOutput) Rules() GetRouterAccessList6RuleArrayOutput {
	return o.ApplyT(func(v LookupRouterAccessList6Result) []GetRouterAccessList6Rule { return v.Rules }).(GetRouterAccessList6RuleArrayOutput)
}

func (o LookupRouterAccessList6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterAccessList6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterAccessList6ResultOutput{})
}

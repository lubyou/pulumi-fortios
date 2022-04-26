// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router prefixlist6
func LookupRouterPrefixList6(ctx *pulumi.Context, args *LookupRouterPrefixList6Args, opts ...pulumi.InvokeOption) (*LookupRouterPrefixList6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterPrefixList6Result
	err := ctx.Invoke("fortios:index/getRouterPrefixList6:GetRouterPrefixList6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPrefixList6.
type LookupRouterPrefixList6Args struct {
	// Specify the name of the desired router prefixlist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPrefixList6.
type LookupRouterPrefixList6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules     []GetRouterPrefixList6Rule `pulumi:"rules"`
	Vdomparam *string                    `pulumi:"vdomparam"`
}

func LookupRouterPrefixList6Output(ctx *pulumi.Context, args LookupRouterPrefixList6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterPrefixList6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterPrefixList6Result, error) {
			args := v.(LookupRouterPrefixList6Args)
			r, err := LookupRouterPrefixList6(ctx, &args, opts...)
			var s LookupRouterPrefixList6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterPrefixList6ResultOutput)
}

// A collection of arguments for invoking GetRouterPrefixList6.
type LookupRouterPrefixList6OutputArgs struct {
	// Specify the name of the desired router prefixlist6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterPrefixList6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPrefixList6Args)(nil)).Elem()
}

// A collection of values returned by GetRouterPrefixList6.
type LookupRouterPrefixList6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterPrefixList6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPrefixList6Result)(nil)).Elem()
}

func (o LookupRouterPrefixList6ResultOutput) ToLookupRouterPrefixList6ResultOutput() LookupRouterPrefixList6ResultOutput {
	return o
}

func (o LookupRouterPrefixList6ResultOutput) ToLookupRouterPrefixList6ResultOutputWithContext(ctx context.Context) LookupRouterPrefixList6ResultOutput {
	return o
}

// Comment.
func (o LookupRouterPrefixList6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixList6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterPrefixList6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixList6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRouterPrefixList6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixList6Result) string { return v.Name }).(pulumi.StringOutput)
}

// IPv6 prefix list rule. The structure of `rule` block is documented below.
func (o LookupRouterPrefixList6ResultOutput) Rules() GetRouterPrefixList6RuleArrayOutput {
	return o.ApplyT(func(v LookupRouterPrefixList6Result) []GetRouterPrefixList6Rule { return v.Rules }).(GetRouterPrefixList6RuleArrayOutput)
}

func (o LookupRouterPrefixList6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterPrefixList6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterPrefixList6ResultOutput{})
}

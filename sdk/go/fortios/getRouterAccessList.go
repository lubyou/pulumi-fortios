// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router accesslist
func LookupRouterAccessList(ctx *pulumi.Context, args *LookupRouterAccessListArgs, opts ...pulumi.InvokeOption) (*LookupRouterAccessListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterAccessListResult
	err := ctx.Invoke("fortios:index/getRouterAccessList:GetRouterAccessList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAccessList.
type LookupRouterAccessListArgs struct {
	// Specify the name of the desired router accesslist.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAccessList.
type LookupRouterAccessListResult struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRouterAccessListRule `pulumi:"rules"`
	Vdomparam *string                   `pulumi:"vdomparam"`
}

func LookupRouterAccessListOutput(ctx *pulumi.Context, args LookupRouterAccessListOutputArgs, opts ...pulumi.InvokeOption) LookupRouterAccessListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterAccessListResult, error) {
			args := v.(LookupRouterAccessListArgs)
			r, err := LookupRouterAccessList(ctx, &args, opts...)
			return *r, err
		}).(LookupRouterAccessListResultOutput)
}

// A collection of arguments for invoking GetRouterAccessList.
type LookupRouterAccessListOutputArgs struct {
	// Specify the name of the desired router accesslist.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterAccessListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAccessListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterAccessList.
type LookupRouterAccessListResultOutput struct{ *pulumi.OutputState }

func (LookupRouterAccessListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAccessListResult)(nil)).Elem()
}

func (o LookupRouterAccessListResultOutput) ToLookupRouterAccessListResultOutput() LookupRouterAccessListResultOutput {
	return o
}

func (o LookupRouterAccessListResultOutput) ToLookupRouterAccessListResultOutputWithContext(ctx context.Context) LookupRouterAccessListResultOutput {
	return o
}

// Comment.
func (o LookupRouterAccessListResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessListResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterAccessListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessListResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRouterAccessListResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAccessListResult) string { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o LookupRouterAccessListResultOutput) Rules() GetRouterAccessListRuleArrayOutput {
	return o.ApplyT(func(v LookupRouterAccessListResult) []GetRouterAccessListRule { return v.Rules }).(GetRouterAccessListRuleArrayOutput)
}

func (o LookupRouterAccessListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterAccessListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterAccessListResultOutput{})
}

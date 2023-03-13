// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallServiceCategory(ctx *pulumi.Context, args *LookupFirewallServiceCategoryArgs, opts ...pulumi.InvokeOption) (*LookupFirewallServiceCategoryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallServiceCategoryResult
	err := ctx.Invoke("fortios:index/getFirewallServiceCategory:GetFirewallServiceCategory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceCategory.
type LookupFirewallServiceCategoryArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceCategory.
type LookupFirewallServiceCategoryResult struct {
	Comment      string `pulumi:"comment"`
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallServiceCategoryOutput(ctx *pulumi.Context, args LookupFirewallServiceCategoryOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallServiceCategoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallServiceCategoryResult, error) {
			args := v.(LookupFirewallServiceCategoryArgs)
			r, err := LookupFirewallServiceCategory(ctx, &args, opts...)
			var s LookupFirewallServiceCategoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallServiceCategoryResultOutput)
}

// A collection of arguments for invoking GetFirewallServiceCategory.
type LookupFirewallServiceCategoryOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallServiceCategoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceCategoryArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallServiceCategory.
type LookupFirewallServiceCategoryResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallServiceCategoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceCategoryResult)(nil)).Elem()
}

func (o LookupFirewallServiceCategoryResultOutput) ToLookupFirewallServiceCategoryResultOutput() LookupFirewallServiceCategoryResultOutput {
	return o
}

func (o LookupFirewallServiceCategoryResultOutput) ToLookupFirewallServiceCategoryResultOutputWithContext(ctx context.Context) LookupFirewallServiceCategoryResultOutput {
	return o
}

func (o LookupFirewallServiceCategoryResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCategoryResult) string { return v.Comment }).(pulumi.StringOutput)
}

func (o LookupFirewallServiceCategoryResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCategoryResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallServiceCategoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCategoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallServiceCategoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCategoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallServiceCategoryResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallServiceCategoryResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallServiceCategoryResultOutput{})
}

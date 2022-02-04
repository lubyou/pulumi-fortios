// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicecustom
func LookupFirewallInternetServiceCustom(ctx *pulumi.Context, args *LookupFirewallInternetServiceCustomArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceCustomResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetServiceCustomResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceCustom:GetFirewallInternetServiceCustom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomArgs struct {
	// Specify the name of the desired firewall internetservicecustom.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
	Entries []GetFirewallInternetServiceCustomEntry `pulumi:"entries"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Select the destination address or address group object from available options.
	Name string `pulumi:"name"`
	// Reputation level of the custom Internet Service.
	Reputation int     `pulumi:"reputation"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetServiceCustomOutput(ctx *pulumi.Context, args LookupFirewallInternetServiceCustomOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetServiceCustomResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetServiceCustomResult, error) {
			args := v.(LookupFirewallInternetServiceCustomArgs)
			r, err := LookupFirewallInternetServiceCustom(ctx, &args, opts...)
			return *r, err
		}).(LookupFirewallInternetServiceCustomResultOutput)
}

// A collection of arguments for invoking GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomOutputArgs struct {
	// Specify the name of the desired firewall internetservicecustom.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetServiceCustomOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceCustomArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetServiceCustomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceCustomResult)(nil)).Elem()
}

func (o LookupFirewallInternetServiceCustomResultOutput) ToLookupFirewallInternetServiceCustomResultOutput() LookupFirewallInternetServiceCustomResultOutput {
	return o
}

func (o LookupFirewallInternetServiceCustomResultOutput) ToLookupFirewallInternetServiceCustomResultOutputWithContext(ctx context.Context) LookupFirewallInternetServiceCustomResultOutput {
	return o
}

// Comment.
func (o LookupFirewallInternetServiceCustomResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
func (o LookupFirewallInternetServiceCustomResultOutput) Entries() GetFirewallInternetServiceCustomEntryArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) []GetFirewallInternetServiceCustomEntry {
		return v.Entries
	}).(GetFirewallInternetServiceCustomEntryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetServiceCustomResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) string { return v.Id }).(pulumi.StringOutput)
}

// Select the destination address or address group object from available options.
func (o LookupFirewallInternetServiceCustomResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) string { return v.Name }).(pulumi.StringOutput)
}

// Reputation level of the custom Internet Service.
func (o LookupFirewallInternetServiceCustomResultOutput) Reputation() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) int { return v.Reputation }).(pulumi.IntOutput)
}

func (o LookupFirewallInternetServiceCustomResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceCustomResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetServiceCustomResultOutput{})
}

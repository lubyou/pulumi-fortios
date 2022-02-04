// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system arptable
func LookupSystemArpTable(ctx *pulumi.Context, args *LookupSystemArpTableArgs, opts ...pulumi.InvokeOption) (*LookupSystemArpTableResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemArpTableResult
	err := ctx.Invoke("fortios:index/getSystemArpTable:GetSystemArpTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemArpTable.
type LookupSystemArpTableArgs struct {
	// Specify the fosid of the desired system arptable.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemArpTable.
type LookupSystemArpTableResult struct {
	// Unique integer ID of the entry.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IP address.
	Ip string `pulumi:"ip"`
	// MAC address.
	Mac       string  `pulumi:"mac"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemArpTableOutput(ctx *pulumi.Context, args LookupSystemArpTableOutputArgs, opts ...pulumi.InvokeOption) LookupSystemArpTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemArpTableResult, error) {
			args := v.(LookupSystemArpTableArgs)
			r, err := LookupSystemArpTable(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemArpTableResultOutput)
}

// A collection of arguments for invoking GetSystemArpTable.
type LookupSystemArpTableOutputArgs struct {
	// Specify the fosid of the desired system arptable.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemArpTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemArpTableArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemArpTable.
type LookupSystemArpTableResultOutput struct{ *pulumi.OutputState }

func (LookupSystemArpTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemArpTableResult)(nil)).Elem()
}

func (o LookupSystemArpTableResultOutput) ToLookupSystemArpTableResultOutput() LookupSystemArpTableResultOutput {
	return o
}

func (o LookupSystemArpTableResultOutput) ToLookupSystemArpTableResultOutputWithContext(ctx context.Context) LookupSystemArpTableResultOutput {
	return o
}

// Unique integer ID of the entry.
func (o LookupSystemArpTableResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemArpTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface name.
func (o LookupSystemArpTableResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) string { return v.Interface }).(pulumi.StringOutput)
}

// IP address.
func (o LookupSystemArpTableResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) string { return v.Ip }).(pulumi.StringOutput)
}

// MAC address.
func (o LookupSystemArpTableResultOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) string { return v.Mac }).(pulumi.StringOutput)
}

func (o LookupSystemArpTableResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemArpTableResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemArpTableResultOutput{})
}

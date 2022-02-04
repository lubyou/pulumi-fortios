// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallservice group
func LookupFirewallServiceGroup(ctx *pulumi.Context, args *LookupFirewallServiceGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallServiceGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallServiceGroupResult
	err := ctx.Invoke("fortios:index/getFirewallServiceGroup:GetFirewallServiceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceGroup.
type LookupFirewallServiceGroupArgs struct {
	// Specify the name of the desired firewallservice group.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceGroup.
type LookupFirewallServiceGroupResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members []GetFirewallServiceGroupMember `pulumi:"members"`
	// Address name.
	Name string `pulumi:"name"`
	// Enable/disable web proxy service group.
	Proxy     string  `pulumi:"proxy"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallServiceGroupOutput(ctx *pulumi.Context, args LookupFirewallServiceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallServiceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallServiceGroupResult, error) {
			args := v.(LookupFirewallServiceGroupArgs)
			r, err := LookupFirewallServiceGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupFirewallServiceGroupResultOutput)
}

// A collection of arguments for invoking GetFirewallServiceGroup.
type LookupFirewallServiceGroupOutputArgs struct {
	// Specify the name of the desired firewallservice group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallServiceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceGroupArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallServiceGroup.
type LookupFirewallServiceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallServiceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceGroupResult)(nil)).Elem()
}

func (o LookupFirewallServiceGroupResultOutput) ToLookupFirewallServiceGroupResultOutput() LookupFirewallServiceGroupResultOutput {
	return o
}

func (o LookupFirewallServiceGroupResultOutput) ToLookupFirewallServiceGroupResultOutputWithContext(ctx context.Context) LookupFirewallServiceGroupResultOutput {
	return o
}

// Color of icon on the GUI.
func (o LookupFirewallServiceGroupResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupFirewallServiceGroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupFirewallServiceGroupResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallServiceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service objects contained within the group. The structure of `member` block is documented below.
func (o LookupFirewallServiceGroupResultOutput) Members() GetFirewallServiceGroupMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) []GetFirewallServiceGroupMember { return v.Members }).(GetFirewallServiceGroupMemberArrayOutput)
}

// Address name.
func (o LookupFirewallServiceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable web proxy service group.
func (o LookupFirewallServiceGroupResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) string { return v.Proxy }).(pulumi.StringOutput)
}

func (o LookupFirewallServiceGroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallServiceGroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallServiceGroupResultOutput{})
}

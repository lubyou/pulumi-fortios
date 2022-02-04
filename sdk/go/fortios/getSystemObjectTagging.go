// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system objecttagging
func LookupSystemObjectTagging(ctx *pulumi.Context, args *LookupSystemObjectTaggingArgs, opts ...pulumi.InvokeOption) (*LookupSystemObjectTaggingResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemObjectTaggingResult
	err := ctx.Invoke("fortios:index/getSystemObjectTagging:GetSystemObjectTagging", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemObjectTagging.
type LookupSystemObjectTaggingArgs struct {
	// Specify the category of the desired system objecttagging.
	Category string `pulumi:"category"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemObjectTagging.
type LookupSystemObjectTaggingResult struct {
	// Address.
	Address string `pulumi:"address"`
	// Tag Category.
	Category string `pulumi:"category"`
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Device.
	Device string `pulumi:"device"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface.
	Interface string `pulumi:"interface"`
	// Allow multiple tag selection.
	Multiple string `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags      []GetSystemObjectTaggingTag `pulumi:"tags"`
	Vdomparam *string                     `pulumi:"vdomparam"`
}

func LookupSystemObjectTaggingOutput(ctx *pulumi.Context, args LookupSystemObjectTaggingOutputArgs, opts ...pulumi.InvokeOption) LookupSystemObjectTaggingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemObjectTaggingResult, error) {
			args := v.(LookupSystemObjectTaggingArgs)
			r, err := LookupSystemObjectTagging(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemObjectTaggingResultOutput)
}

// A collection of arguments for invoking GetSystemObjectTagging.
type LookupSystemObjectTaggingOutputArgs struct {
	// Specify the category of the desired system objecttagging.
	Category pulumi.StringInput `pulumi:"category"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemObjectTaggingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemObjectTaggingArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemObjectTagging.
type LookupSystemObjectTaggingResultOutput struct{ *pulumi.OutputState }

func (LookupSystemObjectTaggingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemObjectTaggingResult)(nil)).Elem()
}

func (o LookupSystemObjectTaggingResultOutput) ToLookupSystemObjectTaggingResultOutput() LookupSystemObjectTaggingResultOutput {
	return o
}

func (o LookupSystemObjectTaggingResultOutput) ToLookupSystemObjectTaggingResultOutputWithContext(ctx context.Context) LookupSystemObjectTaggingResultOutput {
	return o
}

// Address.
func (o LookupSystemObjectTaggingResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Address }).(pulumi.StringOutput)
}

// Tag Category.
func (o LookupSystemObjectTaggingResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Category }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o LookupSystemObjectTaggingResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) int { return v.Color }).(pulumi.IntOutput)
}

// Device.
func (o LookupSystemObjectTaggingResultOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Device }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemObjectTaggingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface.
func (o LookupSystemObjectTaggingResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Allow multiple tag selection.
func (o LookupSystemObjectTaggingResultOutput) Multiple() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) string { return v.Multiple }).(pulumi.StringOutput)
}

// Tags. The structure of `tags` block is documented below.
func (o LookupSystemObjectTaggingResultOutput) Tags() GetSystemObjectTaggingTagArrayOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) []GetSystemObjectTaggingTag { return v.Tags }).(GetSystemObjectTaggingTagArrayOutput)
}

func (o LookupSystemObjectTaggingResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemObjectTaggingResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemObjectTaggingResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebfilterIpsUrlfilterCacheSetting struct {
	pulumi.CustomResourceState

	DnsRetryInterval pulumi.IntOutput       `pulumi:"dnsRetryInterval"`
	ExtendedTtl      pulumi.IntOutput       `pulumi:"extendedTtl"`
	Vdomparam        pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterIpsUrlfilterCacheSetting registers a new resource with the given unique name, arguments, and options.
func NewWebfilterIpsUrlfilterCacheSetting(ctx *pulumi.Context,
	name string, args *WebfilterIpsUrlfilterCacheSettingArgs, opts ...pulumi.ResourceOption) (*WebfilterIpsUrlfilterCacheSetting, error) {
	if args == nil {
		args = &WebfilterIpsUrlfilterCacheSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterIpsUrlfilterCacheSetting
	err := ctx.RegisterResource("fortios:index/webfilterIpsUrlfilterCacheSetting:WebfilterIpsUrlfilterCacheSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterIpsUrlfilterCacheSetting gets an existing WebfilterIpsUrlfilterCacheSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterIpsUrlfilterCacheSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterIpsUrlfilterCacheSettingState, opts ...pulumi.ResourceOption) (*WebfilterIpsUrlfilterCacheSetting, error) {
	var resource WebfilterIpsUrlfilterCacheSetting
	err := ctx.ReadResource("fortios:index/webfilterIpsUrlfilterCacheSetting:WebfilterIpsUrlfilterCacheSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterIpsUrlfilterCacheSetting resources.
type webfilterIpsUrlfilterCacheSettingState struct {
	DnsRetryInterval *int    `pulumi:"dnsRetryInterval"`
	ExtendedTtl      *int    `pulumi:"extendedTtl"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

type WebfilterIpsUrlfilterCacheSettingState struct {
	DnsRetryInterval pulumi.IntPtrInput
	ExtendedTtl      pulumi.IntPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (WebfilterIpsUrlfilterCacheSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterIpsUrlfilterCacheSettingState)(nil)).Elem()
}

type webfilterIpsUrlfilterCacheSettingArgs struct {
	DnsRetryInterval *int    `pulumi:"dnsRetryInterval"`
	ExtendedTtl      *int    `pulumi:"extendedTtl"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterIpsUrlfilterCacheSetting resource.
type WebfilterIpsUrlfilterCacheSettingArgs struct {
	DnsRetryInterval pulumi.IntPtrInput
	ExtendedTtl      pulumi.IntPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (WebfilterIpsUrlfilterCacheSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterIpsUrlfilterCacheSettingArgs)(nil)).Elem()
}

type WebfilterIpsUrlfilterCacheSettingInput interface {
	pulumi.Input

	ToWebfilterIpsUrlfilterCacheSettingOutput() WebfilterIpsUrlfilterCacheSettingOutput
	ToWebfilterIpsUrlfilterCacheSettingOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingOutput
}

func (*WebfilterIpsUrlfilterCacheSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (i *WebfilterIpsUrlfilterCacheSetting) ToWebfilterIpsUrlfilterCacheSettingOutput() WebfilterIpsUrlfilterCacheSettingOutput {
	return i.ToWebfilterIpsUrlfilterCacheSettingOutputWithContext(context.Background())
}

func (i *WebfilterIpsUrlfilterCacheSetting) ToWebfilterIpsUrlfilterCacheSettingOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsUrlfilterCacheSettingOutput)
}

// WebfilterIpsUrlfilterCacheSettingArrayInput is an input type that accepts WebfilterIpsUrlfilterCacheSettingArray and WebfilterIpsUrlfilterCacheSettingArrayOutput values.
// You can construct a concrete instance of `WebfilterIpsUrlfilterCacheSettingArrayInput` via:
//
//	WebfilterIpsUrlfilterCacheSettingArray{ WebfilterIpsUrlfilterCacheSettingArgs{...} }
type WebfilterIpsUrlfilterCacheSettingArrayInput interface {
	pulumi.Input

	ToWebfilterIpsUrlfilterCacheSettingArrayOutput() WebfilterIpsUrlfilterCacheSettingArrayOutput
	ToWebfilterIpsUrlfilterCacheSettingArrayOutputWithContext(context.Context) WebfilterIpsUrlfilterCacheSettingArrayOutput
}

type WebfilterIpsUrlfilterCacheSettingArray []WebfilterIpsUrlfilterCacheSettingInput

func (WebfilterIpsUrlfilterCacheSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (i WebfilterIpsUrlfilterCacheSettingArray) ToWebfilterIpsUrlfilterCacheSettingArrayOutput() WebfilterIpsUrlfilterCacheSettingArrayOutput {
	return i.ToWebfilterIpsUrlfilterCacheSettingArrayOutputWithContext(context.Background())
}

func (i WebfilterIpsUrlfilterCacheSettingArray) ToWebfilterIpsUrlfilterCacheSettingArrayOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsUrlfilterCacheSettingArrayOutput)
}

// WebfilterIpsUrlfilterCacheSettingMapInput is an input type that accepts WebfilterIpsUrlfilterCacheSettingMap and WebfilterIpsUrlfilterCacheSettingMapOutput values.
// You can construct a concrete instance of `WebfilterIpsUrlfilterCacheSettingMapInput` via:
//
//	WebfilterIpsUrlfilterCacheSettingMap{ "key": WebfilterIpsUrlfilterCacheSettingArgs{...} }
type WebfilterIpsUrlfilterCacheSettingMapInput interface {
	pulumi.Input

	ToWebfilterIpsUrlfilterCacheSettingMapOutput() WebfilterIpsUrlfilterCacheSettingMapOutput
	ToWebfilterIpsUrlfilterCacheSettingMapOutputWithContext(context.Context) WebfilterIpsUrlfilterCacheSettingMapOutput
}

type WebfilterIpsUrlfilterCacheSettingMap map[string]WebfilterIpsUrlfilterCacheSettingInput

func (WebfilterIpsUrlfilterCacheSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (i WebfilterIpsUrlfilterCacheSettingMap) ToWebfilterIpsUrlfilterCacheSettingMapOutput() WebfilterIpsUrlfilterCacheSettingMapOutput {
	return i.ToWebfilterIpsUrlfilterCacheSettingMapOutputWithContext(context.Background())
}

func (i WebfilterIpsUrlfilterCacheSettingMap) ToWebfilterIpsUrlfilterCacheSettingMapOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterIpsUrlfilterCacheSettingMapOutput)
}

type WebfilterIpsUrlfilterCacheSettingOutput struct{ *pulumi.OutputState }

func (WebfilterIpsUrlfilterCacheSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (o WebfilterIpsUrlfilterCacheSettingOutput) ToWebfilterIpsUrlfilterCacheSettingOutput() WebfilterIpsUrlfilterCacheSettingOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingOutput) ToWebfilterIpsUrlfilterCacheSettingOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingOutput) DnsRetryInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WebfilterIpsUrlfilterCacheSetting) pulumi.IntOutput { return v.DnsRetryInterval }).(pulumi.IntOutput)
}

func (o WebfilterIpsUrlfilterCacheSettingOutput) ExtendedTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *WebfilterIpsUrlfilterCacheSetting) pulumi.IntOutput { return v.ExtendedTtl }).(pulumi.IntOutput)
}

func (o WebfilterIpsUrlfilterCacheSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebfilterIpsUrlfilterCacheSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebfilterIpsUrlfilterCacheSettingArrayOutput struct{ *pulumi.OutputState }

func (WebfilterIpsUrlfilterCacheSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (o WebfilterIpsUrlfilterCacheSettingArrayOutput) ToWebfilterIpsUrlfilterCacheSettingArrayOutput() WebfilterIpsUrlfilterCacheSettingArrayOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingArrayOutput) ToWebfilterIpsUrlfilterCacheSettingArrayOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingArrayOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingArrayOutput) Index(i pulumi.IntInput) WebfilterIpsUrlfilterCacheSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterIpsUrlfilterCacheSetting {
		return vs[0].([]*WebfilterIpsUrlfilterCacheSetting)[vs[1].(int)]
	}).(WebfilterIpsUrlfilterCacheSettingOutput)
}

type WebfilterIpsUrlfilterCacheSettingMapOutput struct{ *pulumi.OutputState }

func (WebfilterIpsUrlfilterCacheSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterIpsUrlfilterCacheSetting)(nil)).Elem()
}

func (o WebfilterIpsUrlfilterCacheSettingMapOutput) ToWebfilterIpsUrlfilterCacheSettingMapOutput() WebfilterIpsUrlfilterCacheSettingMapOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingMapOutput) ToWebfilterIpsUrlfilterCacheSettingMapOutputWithContext(ctx context.Context) WebfilterIpsUrlfilterCacheSettingMapOutput {
	return o
}

func (o WebfilterIpsUrlfilterCacheSettingMapOutput) MapIndex(k pulumi.StringInput) WebfilterIpsUrlfilterCacheSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterIpsUrlfilterCacheSetting {
		return vs[0].(map[string]*WebfilterIpsUrlfilterCacheSetting)[vs[1].(string)]
	}).(WebfilterIpsUrlfilterCacheSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsUrlfilterCacheSettingInput)(nil)).Elem(), &WebfilterIpsUrlfilterCacheSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsUrlfilterCacheSettingArrayInput)(nil)).Elem(), WebfilterIpsUrlfilterCacheSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterIpsUrlfilterCacheSettingMapInput)(nil)).Elem(), WebfilterIpsUrlfilterCacheSettingMap{})
	pulumi.RegisterOutputType(WebfilterIpsUrlfilterCacheSettingOutput{})
	pulumi.RegisterOutputType(WebfilterIpsUrlfilterCacheSettingArrayOutput{})
	pulumi.RegisterOutputType(WebfilterIpsUrlfilterCacheSettingMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebProxyUrlMatch struct {
	pulumi.CustomResourceState

	CacheExemption pulumi.StringOutput    `pulumi:"cacheExemption"`
	Comment        pulumi.StringPtrOutput `pulumi:"comment"`
	FastFallback   pulumi.StringOutput    `pulumi:"fastFallback"`
	ForwardServer  pulumi.StringOutput    `pulumi:"forwardServer"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Status         pulumi.StringOutput    `pulumi:"status"`
	UrlPattern     pulumi.StringOutput    `pulumi:"urlPattern"`
	Vdomparam      pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebProxyUrlMatch registers a new resource with the given unique name, arguments, and options.
func NewWebProxyUrlMatch(ctx *pulumi.Context,
	name string, args *WebProxyUrlMatchArgs, opts ...pulumi.ResourceOption) (*WebProxyUrlMatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlPattern == nil {
		return nil, errors.New("invalid value for required argument 'UrlPattern'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebProxyUrlMatch
	err := ctx.RegisterResource("fortios:index/webProxyUrlMatch:WebProxyUrlMatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebProxyUrlMatch gets an existing WebProxyUrlMatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebProxyUrlMatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebProxyUrlMatchState, opts ...pulumi.ResourceOption) (*WebProxyUrlMatch, error) {
	var resource WebProxyUrlMatch
	err := ctx.ReadResource("fortios:index/webProxyUrlMatch:WebProxyUrlMatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebProxyUrlMatch resources.
type webProxyUrlMatchState struct {
	CacheExemption *string `pulumi:"cacheExemption"`
	Comment        *string `pulumi:"comment"`
	FastFallback   *string `pulumi:"fastFallback"`
	ForwardServer  *string `pulumi:"forwardServer"`
	Name           *string `pulumi:"name"`
	Status         *string `pulumi:"status"`
	UrlPattern     *string `pulumi:"urlPattern"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

type WebProxyUrlMatchState struct {
	CacheExemption pulumi.StringPtrInput
	Comment        pulumi.StringPtrInput
	FastFallback   pulumi.StringPtrInput
	ForwardServer  pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Status         pulumi.StringPtrInput
	UrlPattern     pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (WebProxyUrlMatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyUrlMatchState)(nil)).Elem()
}

type webProxyUrlMatchArgs struct {
	CacheExemption *string `pulumi:"cacheExemption"`
	Comment        *string `pulumi:"comment"`
	FastFallback   *string `pulumi:"fastFallback"`
	ForwardServer  *string `pulumi:"forwardServer"`
	Name           *string `pulumi:"name"`
	Status         *string `pulumi:"status"`
	UrlPattern     string  `pulumi:"urlPattern"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyUrlMatch resource.
type WebProxyUrlMatchArgs struct {
	CacheExemption pulumi.StringPtrInput
	Comment        pulumi.StringPtrInput
	FastFallback   pulumi.StringPtrInput
	ForwardServer  pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Status         pulumi.StringPtrInput
	UrlPattern     pulumi.StringInput
	Vdomparam      pulumi.StringPtrInput
}

func (WebProxyUrlMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyUrlMatchArgs)(nil)).Elem()
}

type WebProxyUrlMatchInput interface {
	pulumi.Input

	ToWebProxyUrlMatchOutput() WebProxyUrlMatchOutput
	ToWebProxyUrlMatchOutputWithContext(ctx context.Context) WebProxyUrlMatchOutput
}

func (*WebProxyUrlMatch) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyUrlMatch)(nil)).Elem()
}

func (i *WebProxyUrlMatch) ToWebProxyUrlMatchOutput() WebProxyUrlMatchOutput {
	return i.ToWebProxyUrlMatchOutputWithContext(context.Background())
}

func (i *WebProxyUrlMatch) ToWebProxyUrlMatchOutputWithContext(ctx context.Context) WebProxyUrlMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyUrlMatchOutput)
}

// WebProxyUrlMatchArrayInput is an input type that accepts WebProxyUrlMatchArray and WebProxyUrlMatchArrayOutput values.
// You can construct a concrete instance of `WebProxyUrlMatchArrayInput` via:
//
//	WebProxyUrlMatchArray{ WebProxyUrlMatchArgs{...} }
type WebProxyUrlMatchArrayInput interface {
	pulumi.Input

	ToWebProxyUrlMatchArrayOutput() WebProxyUrlMatchArrayOutput
	ToWebProxyUrlMatchArrayOutputWithContext(context.Context) WebProxyUrlMatchArrayOutput
}

type WebProxyUrlMatchArray []WebProxyUrlMatchInput

func (WebProxyUrlMatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyUrlMatch)(nil)).Elem()
}

func (i WebProxyUrlMatchArray) ToWebProxyUrlMatchArrayOutput() WebProxyUrlMatchArrayOutput {
	return i.ToWebProxyUrlMatchArrayOutputWithContext(context.Background())
}

func (i WebProxyUrlMatchArray) ToWebProxyUrlMatchArrayOutputWithContext(ctx context.Context) WebProxyUrlMatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyUrlMatchArrayOutput)
}

// WebProxyUrlMatchMapInput is an input type that accepts WebProxyUrlMatchMap and WebProxyUrlMatchMapOutput values.
// You can construct a concrete instance of `WebProxyUrlMatchMapInput` via:
//
//	WebProxyUrlMatchMap{ "key": WebProxyUrlMatchArgs{...} }
type WebProxyUrlMatchMapInput interface {
	pulumi.Input

	ToWebProxyUrlMatchMapOutput() WebProxyUrlMatchMapOutput
	ToWebProxyUrlMatchMapOutputWithContext(context.Context) WebProxyUrlMatchMapOutput
}

type WebProxyUrlMatchMap map[string]WebProxyUrlMatchInput

func (WebProxyUrlMatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyUrlMatch)(nil)).Elem()
}

func (i WebProxyUrlMatchMap) ToWebProxyUrlMatchMapOutput() WebProxyUrlMatchMapOutput {
	return i.ToWebProxyUrlMatchMapOutputWithContext(context.Background())
}

func (i WebProxyUrlMatchMap) ToWebProxyUrlMatchMapOutputWithContext(ctx context.Context) WebProxyUrlMatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyUrlMatchMapOutput)
}

type WebProxyUrlMatchOutput struct{ *pulumi.OutputState }

func (WebProxyUrlMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyUrlMatch)(nil)).Elem()
}

func (o WebProxyUrlMatchOutput) ToWebProxyUrlMatchOutput() WebProxyUrlMatchOutput {
	return o
}

func (o WebProxyUrlMatchOutput) ToWebProxyUrlMatchOutputWithContext(ctx context.Context) WebProxyUrlMatchOutput {
	return o
}

func (o WebProxyUrlMatchOutput) CacheExemption() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.CacheExemption }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WebProxyUrlMatchOutput) FastFallback() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.FastFallback }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) ForwardServer() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.ForwardServer }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) UrlPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringOutput { return v.UrlPattern }).(pulumi.StringOutput)
}

func (o WebProxyUrlMatchOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyUrlMatch) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebProxyUrlMatchArrayOutput struct{ *pulumi.OutputState }

func (WebProxyUrlMatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyUrlMatch)(nil)).Elem()
}

func (o WebProxyUrlMatchArrayOutput) ToWebProxyUrlMatchArrayOutput() WebProxyUrlMatchArrayOutput {
	return o
}

func (o WebProxyUrlMatchArrayOutput) ToWebProxyUrlMatchArrayOutputWithContext(ctx context.Context) WebProxyUrlMatchArrayOutput {
	return o
}

func (o WebProxyUrlMatchArrayOutput) Index(i pulumi.IntInput) WebProxyUrlMatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebProxyUrlMatch {
		return vs[0].([]*WebProxyUrlMatch)[vs[1].(int)]
	}).(WebProxyUrlMatchOutput)
}

type WebProxyUrlMatchMapOutput struct{ *pulumi.OutputState }

func (WebProxyUrlMatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyUrlMatch)(nil)).Elem()
}

func (o WebProxyUrlMatchMapOutput) ToWebProxyUrlMatchMapOutput() WebProxyUrlMatchMapOutput {
	return o
}

func (o WebProxyUrlMatchMapOutput) ToWebProxyUrlMatchMapOutputWithContext(ctx context.Context) WebProxyUrlMatchMapOutput {
	return o
}

func (o WebProxyUrlMatchMapOutput) MapIndex(k pulumi.StringInput) WebProxyUrlMatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebProxyUrlMatch {
		return vs[0].(map[string]*WebProxyUrlMatch)[vs[1].(string)]
	}).(WebProxyUrlMatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyUrlMatchInput)(nil)).Elem(), &WebProxyUrlMatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyUrlMatchArrayInput)(nil)).Elem(), WebProxyUrlMatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyUrlMatchMapInput)(nil)).Elem(), WebProxyUrlMatchMap{})
	pulumi.RegisterOutputType(WebProxyUrlMatchOutput{})
	pulumi.RegisterOutputType(WebProxyUrlMatchArrayOutput{})
	pulumi.RegisterOutputType(WebProxyUrlMatchMapOutput{})
}

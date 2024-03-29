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

type FirewallSslSetting struct {
	pulumi.CustomResourceState

	AbbreviateHandshake    pulumi.StringOutput    `pulumi:"abbreviateHandshake"`
	CertCacheCapacity      pulumi.IntOutput       `pulumi:"certCacheCapacity"`
	CertCacheTimeout       pulumi.IntOutput       `pulumi:"certCacheTimeout"`
	KxpQueueThreshold      pulumi.IntOutput       `pulumi:"kxpQueueThreshold"`
	NoMatchingCipherAction pulumi.StringOutput    `pulumi:"noMatchingCipherAction"`
	ProxyConnectTimeout    pulumi.IntOutput       `pulumi:"proxyConnectTimeout"`
	SessionCacheCapacity   pulumi.IntOutput       `pulumi:"sessionCacheCapacity"`
	SessionCacheTimeout    pulumi.IntOutput       `pulumi:"sessionCacheTimeout"`
	SslDhBits              pulumi.StringOutput    `pulumi:"sslDhBits"`
	SslQueueThreshold      pulumi.IntOutput       `pulumi:"sslQueueThreshold"`
	SslSendEmptyFrags      pulumi.StringOutput    `pulumi:"sslSendEmptyFrags"`
	Vdomparam              pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSslSetting registers a new resource with the given unique name, arguments, and options.
func NewFirewallSslSetting(ctx *pulumi.Context,
	name string, args *FirewallSslSettingArgs, opts ...pulumi.ResourceOption) (*FirewallSslSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertCacheCapacity == nil {
		return nil, errors.New("invalid value for required argument 'CertCacheCapacity'")
	}
	if args.CertCacheTimeout == nil {
		return nil, errors.New("invalid value for required argument 'CertCacheTimeout'")
	}
	if args.NoMatchingCipherAction == nil {
		return nil, errors.New("invalid value for required argument 'NoMatchingCipherAction'")
	}
	if args.ProxyConnectTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ProxyConnectTimeout'")
	}
	if args.SessionCacheCapacity == nil {
		return nil, errors.New("invalid value for required argument 'SessionCacheCapacity'")
	}
	if args.SessionCacheTimeout == nil {
		return nil, errors.New("invalid value for required argument 'SessionCacheTimeout'")
	}
	if args.SslDhBits == nil {
		return nil, errors.New("invalid value for required argument 'SslDhBits'")
	}
	if args.SslSendEmptyFrags == nil {
		return nil, errors.New("invalid value for required argument 'SslSendEmptyFrags'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallSslSetting
	err := ctx.RegisterResource("fortios:index/firewallSslSetting:FirewallSslSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSslSetting gets an existing FirewallSslSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSslSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSslSettingState, opts ...pulumi.ResourceOption) (*FirewallSslSetting, error) {
	var resource FirewallSslSetting
	err := ctx.ReadResource("fortios:index/firewallSslSetting:FirewallSslSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSslSetting resources.
type firewallSslSettingState struct {
	AbbreviateHandshake    *string `pulumi:"abbreviateHandshake"`
	CertCacheCapacity      *int    `pulumi:"certCacheCapacity"`
	CertCacheTimeout       *int    `pulumi:"certCacheTimeout"`
	KxpQueueThreshold      *int    `pulumi:"kxpQueueThreshold"`
	NoMatchingCipherAction *string `pulumi:"noMatchingCipherAction"`
	ProxyConnectTimeout    *int    `pulumi:"proxyConnectTimeout"`
	SessionCacheCapacity   *int    `pulumi:"sessionCacheCapacity"`
	SessionCacheTimeout    *int    `pulumi:"sessionCacheTimeout"`
	SslDhBits              *string `pulumi:"sslDhBits"`
	SslQueueThreshold      *int    `pulumi:"sslQueueThreshold"`
	SslSendEmptyFrags      *string `pulumi:"sslSendEmptyFrags"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

type FirewallSslSettingState struct {
	AbbreviateHandshake    pulumi.StringPtrInput
	CertCacheCapacity      pulumi.IntPtrInput
	CertCacheTimeout       pulumi.IntPtrInput
	KxpQueueThreshold      pulumi.IntPtrInput
	NoMatchingCipherAction pulumi.StringPtrInput
	ProxyConnectTimeout    pulumi.IntPtrInput
	SessionCacheCapacity   pulumi.IntPtrInput
	SessionCacheTimeout    pulumi.IntPtrInput
	SslDhBits              pulumi.StringPtrInput
	SslQueueThreshold      pulumi.IntPtrInput
	SslSendEmptyFrags      pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (FirewallSslSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslSettingState)(nil)).Elem()
}

type firewallSslSettingArgs struct {
	AbbreviateHandshake    *string `pulumi:"abbreviateHandshake"`
	CertCacheCapacity      int     `pulumi:"certCacheCapacity"`
	CertCacheTimeout       int     `pulumi:"certCacheTimeout"`
	KxpQueueThreshold      *int    `pulumi:"kxpQueueThreshold"`
	NoMatchingCipherAction string  `pulumi:"noMatchingCipherAction"`
	ProxyConnectTimeout    int     `pulumi:"proxyConnectTimeout"`
	SessionCacheCapacity   int     `pulumi:"sessionCacheCapacity"`
	SessionCacheTimeout    int     `pulumi:"sessionCacheTimeout"`
	SslDhBits              string  `pulumi:"sslDhBits"`
	SslQueueThreshold      *int    `pulumi:"sslQueueThreshold"`
	SslSendEmptyFrags      string  `pulumi:"sslSendEmptyFrags"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSslSetting resource.
type FirewallSslSettingArgs struct {
	AbbreviateHandshake    pulumi.StringPtrInput
	CertCacheCapacity      pulumi.IntInput
	CertCacheTimeout       pulumi.IntInput
	KxpQueueThreshold      pulumi.IntPtrInput
	NoMatchingCipherAction pulumi.StringInput
	ProxyConnectTimeout    pulumi.IntInput
	SessionCacheCapacity   pulumi.IntInput
	SessionCacheTimeout    pulumi.IntInput
	SslDhBits              pulumi.StringInput
	SslQueueThreshold      pulumi.IntPtrInput
	SslSendEmptyFrags      pulumi.StringInput
	Vdomparam              pulumi.StringPtrInput
}

func (FirewallSslSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslSettingArgs)(nil)).Elem()
}

type FirewallSslSettingInput interface {
	pulumi.Input

	ToFirewallSslSettingOutput() FirewallSslSettingOutput
	ToFirewallSslSettingOutputWithContext(ctx context.Context) FirewallSslSettingOutput
}

func (*FirewallSslSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslSetting)(nil)).Elem()
}

func (i *FirewallSslSetting) ToFirewallSslSettingOutput() FirewallSslSettingOutput {
	return i.ToFirewallSslSettingOutputWithContext(context.Background())
}

func (i *FirewallSslSetting) ToFirewallSslSettingOutputWithContext(ctx context.Context) FirewallSslSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSettingOutput)
}

// FirewallSslSettingArrayInput is an input type that accepts FirewallSslSettingArray and FirewallSslSettingArrayOutput values.
// You can construct a concrete instance of `FirewallSslSettingArrayInput` via:
//
//	FirewallSslSettingArray{ FirewallSslSettingArgs{...} }
type FirewallSslSettingArrayInput interface {
	pulumi.Input

	ToFirewallSslSettingArrayOutput() FirewallSslSettingArrayOutput
	ToFirewallSslSettingArrayOutputWithContext(context.Context) FirewallSslSettingArrayOutput
}

type FirewallSslSettingArray []FirewallSslSettingInput

func (FirewallSslSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslSetting)(nil)).Elem()
}

func (i FirewallSslSettingArray) ToFirewallSslSettingArrayOutput() FirewallSslSettingArrayOutput {
	return i.ToFirewallSslSettingArrayOutputWithContext(context.Background())
}

func (i FirewallSslSettingArray) ToFirewallSslSettingArrayOutputWithContext(ctx context.Context) FirewallSslSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSettingArrayOutput)
}

// FirewallSslSettingMapInput is an input type that accepts FirewallSslSettingMap and FirewallSslSettingMapOutput values.
// You can construct a concrete instance of `FirewallSslSettingMapInput` via:
//
//	FirewallSslSettingMap{ "key": FirewallSslSettingArgs{...} }
type FirewallSslSettingMapInput interface {
	pulumi.Input

	ToFirewallSslSettingMapOutput() FirewallSslSettingMapOutput
	ToFirewallSslSettingMapOutputWithContext(context.Context) FirewallSslSettingMapOutput
}

type FirewallSslSettingMap map[string]FirewallSslSettingInput

func (FirewallSslSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslSetting)(nil)).Elem()
}

func (i FirewallSslSettingMap) ToFirewallSslSettingMapOutput() FirewallSslSettingMapOutput {
	return i.ToFirewallSslSettingMapOutputWithContext(context.Background())
}

func (i FirewallSslSettingMap) ToFirewallSslSettingMapOutputWithContext(ctx context.Context) FirewallSslSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSettingMapOutput)
}

type FirewallSslSettingOutput struct{ *pulumi.OutputState }

func (FirewallSslSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslSetting)(nil)).Elem()
}

func (o FirewallSslSettingOutput) ToFirewallSslSettingOutput() FirewallSslSettingOutput {
	return o
}

func (o FirewallSslSettingOutput) ToFirewallSslSettingOutputWithContext(ctx context.Context) FirewallSslSettingOutput {
	return o
}

func (o FirewallSslSettingOutput) AbbreviateHandshake() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.StringOutput { return v.AbbreviateHandshake }).(pulumi.StringOutput)
}

func (o FirewallSslSettingOutput) CertCacheCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.CertCacheCapacity }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) CertCacheTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.CertCacheTimeout }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) KxpQueueThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.KxpQueueThreshold }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) NoMatchingCipherAction() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.StringOutput { return v.NoMatchingCipherAction }).(pulumi.StringOutput)
}

func (o FirewallSslSettingOutput) ProxyConnectTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.ProxyConnectTimeout }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) SessionCacheCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.SessionCacheCapacity }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) SessionCacheTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.SessionCacheTimeout }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

func (o FirewallSslSettingOutput) SslQueueThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.IntOutput { return v.SslQueueThreshold }).(pulumi.IntOutput)
}

func (o FirewallSslSettingOutput) SslSendEmptyFrags() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.StringOutput { return v.SslSendEmptyFrags }).(pulumi.StringOutput)
}

func (o FirewallSslSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSslSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSslSettingArrayOutput struct{ *pulumi.OutputState }

func (FirewallSslSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslSetting)(nil)).Elem()
}

func (o FirewallSslSettingArrayOutput) ToFirewallSslSettingArrayOutput() FirewallSslSettingArrayOutput {
	return o
}

func (o FirewallSslSettingArrayOutput) ToFirewallSslSettingArrayOutputWithContext(ctx context.Context) FirewallSslSettingArrayOutput {
	return o
}

func (o FirewallSslSettingArrayOutput) Index(i pulumi.IntInput) FirewallSslSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSslSetting {
		return vs[0].([]*FirewallSslSetting)[vs[1].(int)]
	}).(FirewallSslSettingOutput)
}

type FirewallSslSettingMapOutput struct{ *pulumi.OutputState }

func (FirewallSslSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslSetting)(nil)).Elem()
}

func (o FirewallSslSettingMapOutput) ToFirewallSslSettingMapOutput() FirewallSslSettingMapOutput {
	return o
}

func (o FirewallSslSettingMapOutput) ToFirewallSslSettingMapOutputWithContext(ctx context.Context) FirewallSslSettingMapOutput {
	return o
}

func (o FirewallSslSettingMapOutput) MapIndex(k pulumi.StringInput) FirewallSslSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSslSetting {
		return vs[0].(map[string]*FirewallSslSetting)[vs[1].(string)]
	}).(FirewallSslSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSettingInput)(nil)).Elem(), &FirewallSslSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSettingArrayInput)(nil)).Elem(), FirewallSslSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSettingMapInput)(nil)).Elem(), FirewallSslSettingMap{})
	pulumi.RegisterOutputType(FirewallSslSettingOutput{})
	pulumi.RegisterOutputType(FirewallSslSettingArrayOutput{})
	pulumi.RegisterOutputType(FirewallSslSettingMapOutput{})
}

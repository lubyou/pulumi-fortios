// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallSshHostKey struct {
	pulumi.CustomResourceState

	Hostname  pulumi.StringOutput    `pulumi:"hostname"`
	Ip        pulumi.StringOutput    `pulumi:"ip"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Nid       pulumi.StringOutput    `pulumi:"nid"`
	Port      pulumi.IntOutput       `pulumi:"port"`
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	Status    pulumi.StringOutput    `pulumi:"status"`
	Type      pulumi.StringOutput    `pulumi:"type"`
	Usage     pulumi.StringOutput    `pulumi:"usage"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSshHostKey registers a new resource with the given unique name, arguments, and options.
func NewFirewallSshHostKey(ctx *pulumi.Context,
	name string, args *FirewallSshHostKeyArgs, opts ...pulumi.ResourceOption) (*FirewallSshHostKey, error) {
	if args == nil {
		args = &FirewallSshHostKeyArgs{}
	}

	if args.PublicKey != nil {
		args.PublicKey = pulumi.ToSecret(args.PublicKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"publicKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallSshHostKey
	err := ctx.RegisterResource("fortios:index/firewallSshHostKey:FirewallSshHostKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSshHostKey gets an existing FirewallSshHostKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSshHostKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSshHostKeyState, opts ...pulumi.ResourceOption) (*FirewallSshHostKey, error) {
	var resource FirewallSshHostKey
	err := ctx.ReadResource("fortios:index/firewallSshHostKey:FirewallSshHostKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSshHostKey resources.
type firewallSshHostKeyState struct {
	Hostname  *string `pulumi:"hostname"`
	Ip        *string `pulumi:"ip"`
	Name      *string `pulumi:"name"`
	Nid       *string `pulumi:"nid"`
	Port      *int    `pulumi:"port"`
	PublicKey *string `pulumi:"publicKey"`
	Status    *string `pulumi:"status"`
	Type      *string `pulumi:"type"`
	Usage     *string `pulumi:"usage"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallSshHostKeyState struct {
	Hostname  pulumi.StringPtrInput
	Ip        pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Nid       pulumi.StringPtrInput
	Port      pulumi.IntPtrInput
	PublicKey pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Usage     pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSshHostKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshHostKeyState)(nil)).Elem()
}

type firewallSshHostKeyArgs struct {
	Hostname  *string `pulumi:"hostname"`
	Ip        *string `pulumi:"ip"`
	Name      *string `pulumi:"name"`
	Nid       *string `pulumi:"nid"`
	Port      *int    `pulumi:"port"`
	PublicKey *string `pulumi:"publicKey"`
	Status    *string `pulumi:"status"`
	Type      *string `pulumi:"type"`
	Usage     *string `pulumi:"usage"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSshHostKey resource.
type FirewallSshHostKeyArgs struct {
	Hostname  pulumi.StringPtrInput
	Ip        pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Nid       pulumi.StringPtrInput
	Port      pulumi.IntPtrInput
	PublicKey pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Usage     pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSshHostKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshHostKeyArgs)(nil)).Elem()
}

type FirewallSshHostKeyInput interface {
	pulumi.Input

	ToFirewallSshHostKeyOutput() FirewallSshHostKeyOutput
	ToFirewallSshHostKeyOutputWithContext(ctx context.Context) FirewallSshHostKeyOutput
}

func (*FirewallSshHostKey) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshHostKey)(nil)).Elem()
}

func (i *FirewallSshHostKey) ToFirewallSshHostKeyOutput() FirewallSshHostKeyOutput {
	return i.ToFirewallSshHostKeyOutputWithContext(context.Background())
}

func (i *FirewallSshHostKey) ToFirewallSshHostKeyOutputWithContext(ctx context.Context) FirewallSshHostKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshHostKeyOutput)
}

// FirewallSshHostKeyArrayInput is an input type that accepts FirewallSshHostKeyArray and FirewallSshHostKeyArrayOutput values.
// You can construct a concrete instance of `FirewallSshHostKeyArrayInput` via:
//
//	FirewallSshHostKeyArray{ FirewallSshHostKeyArgs{...} }
type FirewallSshHostKeyArrayInput interface {
	pulumi.Input

	ToFirewallSshHostKeyArrayOutput() FirewallSshHostKeyArrayOutput
	ToFirewallSshHostKeyArrayOutputWithContext(context.Context) FirewallSshHostKeyArrayOutput
}

type FirewallSshHostKeyArray []FirewallSshHostKeyInput

func (FirewallSshHostKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshHostKey)(nil)).Elem()
}

func (i FirewallSshHostKeyArray) ToFirewallSshHostKeyArrayOutput() FirewallSshHostKeyArrayOutput {
	return i.ToFirewallSshHostKeyArrayOutputWithContext(context.Background())
}

func (i FirewallSshHostKeyArray) ToFirewallSshHostKeyArrayOutputWithContext(ctx context.Context) FirewallSshHostKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshHostKeyArrayOutput)
}

// FirewallSshHostKeyMapInput is an input type that accepts FirewallSshHostKeyMap and FirewallSshHostKeyMapOutput values.
// You can construct a concrete instance of `FirewallSshHostKeyMapInput` via:
//
//	FirewallSshHostKeyMap{ "key": FirewallSshHostKeyArgs{...} }
type FirewallSshHostKeyMapInput interface {
	pulumi.Input

	ToFirewallSshHostKeyMapOutput() FirewallSshHostKeyMapOutput
	ToFirewallSshHostKeyMapOutputWithContext(context.Context) FirewallSshHostKeyMapOutput
}

type FirewallSshHostKeyMap map[string]FirewallSshHostKeyInput

func (FirewallSshHostKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshHostKey)(nil)).Elem()
}

func (i FirewallSshHostKeyMap) ToFirewallSshHostKeyMapOutput() FirewallSshHostKeyMapOutput {
	return i.ToFirewallSshHostKeyMapOutputWithContext(context.Background())
}

func (i FirewallSshHostKeyMap) ToFirewallSshHostKeyMapOutputWithContext(ctx context.Context) FirewallSshHostKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshHostKeyMapOutput)
}

type FirewallSshHostKeyOutput struct{ *pulumi.OutputState }

func (FirewallSshHostKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshHostKey)(nil)).Elem()
}

func (o FirewallSshHostKeyOutput) ToFirewallSshHostKeyOutput() FirewallSshHostKeyOutput {
	return o
}

func (o FirewallSshHostKeyOutput) ToFirewallSshHostKeyOutputWithContext(ctx context.Context) FirewallSshHostKeyOutput {
	return o
}

func (o FirewallSshHostKeyOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Nid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Nid }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o FirewallSshHostKeyOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o FirewallSshHostKeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringOutput { return v.Usage }).(pulumi.StringOutput)
}

func (o FirewallSshHostKeyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshHostKey) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSshHostKeyArrayOutput struct{ *pulumi.OutputState }

func (FirewallSshHostKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshHostKey)(nil)).Elem()
}

func (o FirewallSshHostKeyArrayOutput) ToFirewallSshHostKeyArrayOutput() FirewallSshHostKeyArrayOutput {
	return o
}

func (o FirewallSshHostKeyArrayOutput) ToFirewallSshHostKeyArrayOutputWithContext(ctx context.Context) FirewallSshHostKeyArrayOutput {
	return o
}

func (o FirewallSshHostKeyArrayOutput) Index(i pulumi.IntInput) FirewallSshHostKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSshHostKey {
		return vs[0].([]*FirewallSshHostKey)[vs[1].(int)]
	}).(FirewallSshHostKeyOutput)
}

type FirewallSshHostKeyMapOutput struct{ *pulumi.OutputState }

func (FirewallSshHostKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshHostKey)(nil)).Elem()
}

func (o FirewallSshHostKeyMapOutput) ToFirewallSshHostKeyMapOutput() FirewallSshHostKeyMapOutput {
	return o
}

func (o FirewallSshHostKeyMapOutput) ToFirewallSshHostKeyMapOutputWithContext(ctx context.Context) FirewallSshHostKeyMapOutput {
	return o
}

func (o FirewallSshHostKeyMapOutput) MapIndex(k pulumi.StringInput) FirewallSshHostKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSshHostKey {
		return vs[0].(map[string]*FirewallSshHostKey)[vs[1].(string)]
	}).(FirewallSshHostKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshHostKeyInput)(nil)).Elem(), &FirewallSshHostKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshHostKeyArrayInput)(nil)).Elem(), FirewallSshHostKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshHostKeyMapInput)(nil)).Elem(), FirewallSshHostKeyMap{})
	pulumi.RegisterOutputType(FirewallSshHostKeyOutput{})
	pulumi.RegisterOutputType(FirewallSshHostKeyArrayOutput{})
	pulumi.RegisterOutputType(FirewallSshHostKeyMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type FirewallSshLocalKey struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput    `pulumi:"name"`
	Password   pulumi.StringPtrOutput `pulumi:"password"`
	PrivateKey pulumi.StringOutput    `pulumi:"privateKey"`
	PublicKey  pulumi.StringOutput    `pulumi:"publicKey"`
	Source     pulumi.StringOutput    `pulumi:"source"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSshLocalKey registers a new resource with the given unique name, arguments, and options.
func NewFirewallSshLocalKey(ctx *pulumi.Context,
	name string, args *FirewallSshLocalKeyArgs, opts ...pulumi.ResourceOption) (*FirewallSshLocalKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	if args.PublicKey != nil {
		args.PublicKey = pulumi.ToSecret(args.PublicKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"privateKey",
		"publicKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallSshLocalKey
	err := ctx.RegisterResource("fortios:index/firewallSshLocalKey:FirewallSshLocalKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSshLocalKey gets an existing FirewallSshLocalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSshLocalKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSshLocalKeyState, opts ...pulumi.ResourceOption) (*FirewallSshLocalKey, error) {
	var resource FirewallSshLocalKey
	err := ctx.ReadResource("fortios:index/firewallSshLocalKey:FirewallSshLocalKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSshLocalKey resources.
type firewallSshLocalKeyState struct {
	Name       *string `pulumi:"name"`
	Password   *string `pulumi:"password"`
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
	Source     *string `pulumi:"source"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type FirewallSshLocalKeyState struct {
	Name       pulumi.StringPtrInput
	Password   pulumi.StringPtrInput
	PrivateKey pulumi.StringPtrInput
	PublicKey  pulumi.StringPtrInput
	Source     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallSshLocalKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshLocalKeyState)(nil)).Elem()
}

type firewallSshLocalKeyArgs struct {
	Name       *string `pulumi:"name"`
	Password   *string `pulumi:"password"`
	PrivateKey string  `pulumi:"privateKey"`
	PublicKey  string  `pulumi:"publicKey"`
	Source     *string `pulumi:"source"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSshLocalKey resource.
type FirewallSshLocalKeyArgs struct {
	Name       pulumi.StringPtrInput
	Password   pulumi.StringPtrInput
	PrivateKey pulumi.StringInput
	PublicKey  pulumi.StringInput
	Source     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallSshLocalKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshLocalKeyArgs)(nil)).Elem()
}

type FirewallSshLocalKeyInput interface {
	pulumi.Input

	ToFirewallSshLocalKeyOutput() FirewallSshLocalKeyOutput
	ToFirewallSshLocalKeyOutputWithContext(ctx context.Context) FirewallSshLocalKeyOutput
}

func (*FirewallSshLocalKey) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalKey)(nil)).Elem()
}

func (i *FirewallSshLocalKey) ToFirewallSshLocalKeyOutput() FirewallSshLocalKeyOutput {
	return i.ToFirewallSshLocalKeyOutputWithContext(context.Background())
}

func (i *FirewallSshLocalKey) ToFirewallSshLocalKeyOutputWithContext(ctx context.Context) FirewallSshLocalKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalKeyOutput)
}

func (i *FirewallSshLocalKey) ToOutput(ctx context.Context) pulumix.Output[*FirewallSshLocalKey] {
	return pulumix.Output[*FirewallSshLocalKey]{
		OutputState: i.ToFirewallSshLocalKeyOutputWithContext(ctx).OutputState,
	}
}

// FirewallSshLocalKeyArrayInput is an input type that accepts FirewallSshLocalKeyArray and FirewallSshLocalKeyArrayOutput values.
// You can construct a concrete instance of `FirewallSshLocalKeyArrayInput` via:
//
//	FirewallSshLocalKeyArray{ FirewallSshLocalKeyArgs{...} }
type FirewallSshLocalKeyArrayInput interface {
	pulumi.Input

	ToFirewallSshLocalKeyArrayOutput() FirewallSshLocalKeyArrayOutput
	ToFirewallSshLocalKeyArrayOutputWithContext(context.Context) FirewallSshLocalKeyArrayOutput
}

type FirewallSshLocalKeyArray []FirewallSshLocalKeyInput

func (FirewallSshLocalKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshLocalKey)(nil)).Elem()
}

func (i FirewallSshLocalKeyArray) ToFirewallSshLocalKeyArrayOutput() FirewallSshLocalKeyArrayOutput {
	return i.ToFirewallSshLocalKeyArrayOutputWithContext(context.Background())
}

func (i FirewallSshLocalKeyArray) ToFirewallSshLocalKeyArrayOutputWithContext(ctx context.Context) FirewallSshLocalKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalKeyArrayOutput)
}

func (i FirewallSshLocalKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallSshLocalKey] {
	return pulumix.Output[[]*FirewallSshLocalKey]{
		OutputState: i.ToFirewallSshLocalKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallSshLocalKeyMapInput is an input type that accepts FirewallSshLocalKeyMap and FirewallSshLocalKeyMapOutput values.
// You can construct a concrete instance of `FirewallSshLocalKeyMapInput` via:
//
//	FirewallSshLocalKeyMap{ "key": FirewallSshLocalKeyArgs{...} }
type FirewallSshLocalKeyMapInput interface {
	pulumi.Input

	ToFirewallSshLocalKeyMapOutput() FirewallSshLocalKeyMapOutput
	ToFirewallSshLocalKeyMapOutputWithContext(context.Context) FirewallSshLocalKeyMapOutput
}

type FirewallSshLocalKeyMap map[string]FirewallSshLocalKeyInput

func (FirewallSshLocalKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshLocalKey)(nil)).Elem()
}

func (i FirewallSshLocalKeyMap) ToFirewallSshLocalKeyMapOutput() FirewallSshLocalKeyMapOutput {
	return i.ToFirewallSshLocalKeyMapOutputWithContext(context.Background())
}

func (i FirewallSshLocalKeyMap) ToFirewallSshLocalKeyMapOutputWithContext(ctx context.Context) FirewallSshLocalKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalKeyMapOutput)
}

func (i FirewallSshLocalKeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallSshLocalKey] {
	return pulumix.Output[map[string]*FirewallSshLocalKey]{
		OutputState: i.ToFirewallSshLocalKeyMapOutputWithContext(ctx).OutputState,
	}
}

type FirewallSshLocalKeyOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalKey)(nil)).Elem()
}

func (o FirewallSshLocalKeyOutput) ToFirewallSshLocalKeyOutput() FirewallSshLocalKeyOutput {
	return o
}

func (o FirewallSshLocalKeyOutput) ToFirewallSshLocalKeyOutputWithContext(ctx context.Context) FirewallSshLocalKeyOutput {
	return o
}

func (o FirewallSshLocalKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallSshLocalKey] {
	return pulumix.Output[*FirewallSshLocalKey]{
		OutputState: o.OutputState,
	}
}

func (o FirewallSshLocalKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallSshLocalKeyOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o FirewallSshLocalKeyOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o FirewallSshLocalKeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

func (o FirewallSshLocalKeyOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o FirewallSshLocalKeyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshLocalKey) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSshLocalKeyArrayOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshLocalKey)(nil)).Elem()
}

func (o FirewallSshLocalKeyArrayOutput) ToFirewallSshLocalKeyArrayOutput() FirewallSshLocalKeyArrayOutput {
	return o
}

func (o FirewallSshLocalKeyArrayOutput) ToFirewallSshLocalKeyArrayOutputWithContext(ctx context.Context) FirewallSshLocalKeyArrayOutput {
	return o
}

func (o FirewallSshLocalKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallSshLocalKey] {
	return pulumix.Output[[]*FirewallSshLocalKey]{
		OutputState: o.OutputState,
	}
}

func (o FirewallSshLocalKeyArrayOutput) Index(i pulumi.IntInput) FirewallSshLocalKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSshLocalKey {
		return vs[0].([]*FirewallSshLocalKey)[vs[1].(int)]
	}).(FirewallSshLocalKeyOutput)
}

type FirewallSshLocalKeyMapOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshLocalKey)(nil)).Elem()
}

func (o FirewallSshLocalKeyMapOutput) ToFirewallSshLocalKeyMapOutput() FirewallSshLocalKeyMapOutput {
	return o
}

func (o FirewallSshLocalKeyMapOutput) ToFirewallSshLocalKeyMapOutputWithContext(ctx context.Context) FirewallSshLocalKeyMapOutput {
	return o
}

func (o FirewallSshLocalKeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallSshLocalKey] {
	return pulumix.Output[map[string]*FirewallSshLocalKey]{
		OutputState: o.OutputState,
	}
}

func (o FirewallSshLocalKeyMapOutput) MapIndex(k pulumi.StringInput) FirewallSshLocalKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSshLocalKey {
		return vs[0].(map[string]*FirewallSshLocalKey)[vs[1].(string)]
	}).(FirewallSshLocalKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalKeyInput)(nil)).Elem(), &FirewallSshLocalKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalKeyArrayInput)(nil)).Elem(), FirewallSshLocalKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalKeyMapInput)(nil)).Elem(), FirewallSshLocalKeyMap{})
	pulumi.RegisterOutputType(FirewallSshLocalKeyOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalKeyArrayOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalKeyMapOutput{})
}

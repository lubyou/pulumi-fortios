// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallSshLocalCa struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput    `pulumi:"name"`
	Password   pulumi.StringPtrOutput `pulumi:"password"`
	PrivateKey pulumi.StringOutput    `pulumi:"privateKey"`
	PublicKey  pulumi.StringOutput    `pulumi:"publicKey"`
	Source     pulumi.StringOutput    `pulumi:"source"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSshLocalCa registers a new resource with the given unique name, arguments, and options.
func NewFirewallSshLocalCa(ctx *pulumi.Context,
	name string, args *FirewallSshLocalCaArgs, opts ...pulumi.ResourceOption) (*FirewallSshLocalCa, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSshLocalCa
	err := ctx.RegisterResource("fortios:index/firewallSshLocalCa:FirewallSshLocalCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSshLocalCa gets an existing FirewallSshLocalCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSshLocalCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSshLocalCaState, opts ...pulumi.ResourceOption) (*FirewallSshLocalCa, error) {
	var resource FirewallSshLocalCa
	err := ctx.ReadResource("fortios:index/firewallSshLocalCa:FirewallSshLocalCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSshLocalCa resources.
type firewallSshLocalCaState struct {
	Name       *string `pulumi:"name"`
	Password   *string `pulumi:"password"`
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
	Source     *string `pulumi:"source"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type FirewallSshLocalCaState struct {
	Name       pulumi.StringPtrInput
	Password   pulumi.StringPtrInput
	PrivateKey pulumi.StringPtrInput
	PublicKey  pulumi.StringPtrInput
	Source     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallSshLocalCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshLocalCaState)(nil)).Elem()
}

type firewallSshLocalCaArgs struct {
	Name       *string `pulumi:"name"`
	Password   *string `pulumi:"password"`
	PrivateKey string  `pulumi:"privateKey"`
	PublicKey  string  `pulumi:"publicKey"`
	Source     *string `pulumi:"source"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSshLocalCa resource.
type FirewallSshLocalCaArgs struct {
	Name       pulumi.StringPtrInput
	Password   pulumi.StringPtrInput
	PrivateKey pulumi.StringInput
	PublicKey  pulumi.StringInput
	Source     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallSshLocalCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshLocalCaArgs)(nil)).Elem()
}

type FirewallSshLocalCaInput interface {
	pulumi.Input

	ToFirewallSshLocalCaOutput() FirewallSshLocalCaOutput
	ToFirewallSshLocalCaOutputWithContext(ctx context.Context) FirewallSshLocalCaOutput
}

func (*FirewallSshLocalCa) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalCa)(nil)).Elem()
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaOutput() FirewallSshLocalCaOutput {
	return i.ToFirewallSshLocalCaOutputWithContext(context.Background())
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaOutputWithContext(ctx context.Context) FirewallSshLocalCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaOutput)
}

// FirewallSshLocalCaArrayInput is an input type that accepts FirewallSshLocalCaArray and FirewallSshLocalCaArrayOutput values.
// You can construct a concrete instance of `FirewallSshLocalCaArrayInput` via:
//
//	FirewallSshLocalCaArray{ FirewallSshLocalCaArgs{...} }
type FirewallSshLocalCaArrayInput interface {
	pulumi.Input

	ToFirewallSshLocalCaArrayOutput() FirewallSshLocalCaArrayOutput
	ToFirewallSshLocalCaArrayOutputWithContext(context.Context) FirewallSshLocalCaArrayOutput
}

type FirewallSshLocalCaArray []FirewallSshLocalCaInput

func (FirewallSshLocalCaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshLocalCa)(nil)).Elem()
}

func (i FirewallSshLocalCaArray) ToFirewallSshLocalCaArrayOutput() FirewallSshLocalCaArrayOutput {
	return i.ToFirewallSshLocalCaArrayOutputWithContext(context.Background())
}

func (i FirewallSshLocalCaArray) ToFirewallSshLocalCaArrayOutputWithContext(ctx context.Context) FirewallSshLocalCaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaArrayOutput)
}

// FirewallSshLocalCaMapInput is an input type that accepts FirewallSshLocalCaMap and FirewallSshLocalCaMapOutput values.
// You can construct a concrete instance of `FirewallSshLocalCaMapInput` via:
//
//	FirewallSshLocalCaMap{ "key": FirewallSshLocalCaArgs{...} }
type FirewallSshLocalCaMapInput interface {
	pulumi.Input

	ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput
	ToFirewallSshLocalCaMapOutputWithContext(context.Context) FirewallSshLocalCaMapOutput
}

type FirewallSshLocalCaMap map[string]FirewallSshLocalCaInput

func (FirewallSshLocalCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshLocalCa)(nil)).Elem()
}

func (i FirewallSshLocalCaMap) ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput {
	return i.ToFirewallSshLocalCaMapOutputWithContext(context.Background())
}

func (i FirewallSshLocalCaMap) ToFirewallSshLocalCaMapOutputWithContext(ctx context.Context) FirewallSshLocalCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaMapOutput)
}

type FirewallSshLocalCaOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalCa)(nil)).Elem()
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaOutput() FirewallSshLocalCaOutput {
	return o
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaOutputWithContext(ctx context.Context) FirewallSshLocalCaOutput {
	return o
}

func (o FirewallSshLocalCaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallSshLocalCaOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o FirewallSshLocalCaOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o FirewallSshLocalCaOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

func (o FirewallSshLocalCaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o FirewallSshLocalCaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSshLocalCa) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSshLocalCaArrayOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSshLocalCa)(nil)).Elem()
}

func (o FirewallSshLocalCaArrayOutput) ToFirewallSshLocalCaArrayOutput() FirewallSshLocalCaArrayOutput {
	return o
}

func (o FirewallSshLocalCaArrayOutput) ToFirewallSshLocalCaArrayOutputWithContext(ctx context.Context) FirewallSshLocalCaArrayOutput {
	return o
}

func (o FirewallSshLocalCaArrayOutput) Index(i pulumi.IntInput) FirewallSshLocalCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSshLocalCa {
		return vs[0].([]*FirewallSshLocalCa)[vs[1].(int)]
	}).(FirewallSshLocalCaOutput)
}

type FirewallSshLocalCaMapOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSshLocalCa)(nil)).Elem()
}

func (o FirewallSshLocalCaMapOutput) ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput {
	return o
}

func (o FirewallSshLocalCaMapOutput) ToFirewallSshLocalCaMapOutputWithContext(ctx context.Context) FirewallSshLocalCaMapOutput {
	return o
}

func (o FirewallSshLocalCaMapOutput) MapIndex(k pulumi.StringInput) FirewallSshLocalCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSshLocalCa {
		return vs[0].(map[string]*FirewallSshLocalCa)[vs[1].(string)]
	}).(FirewallSshLocalCaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalCaInput)(nil)).Elem(), &FirewallSshLocalCa{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalCaArrayInput)(nil)).Elem(), FirewallSshLocalCaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSshLocalCaMapInput)(nil)).Elem(), FirewallSshLocalCaMap{})
	pulumi.RegisterOutputType(FirewallSshLocalCaOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalCaArrayOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalCaMapOutput{})
}

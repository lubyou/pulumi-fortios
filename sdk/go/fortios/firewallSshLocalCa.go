// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSH proxy local CA.
//
// ## Import
//
// FirewallSsh LocalCa can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallSshLocalCa:FirewallSshLocalCa labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallSshLocalCa struct {
	pulumi.CustomResourceState

	// SSH proxy local CA name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for SSH private key.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
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
	// SSH proxy local CA name.
	Name *string `pulumi:"name"`
	// Password for SSH private key.
	Password *string `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey *string `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey *string `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallSshLocalCaState struct {
	// SSH proxy local CA name.
	Name pulumi.StringPtrInput
	// Password for SSH private key.
	Password pulumi.StringPtrInput
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringPtrInput
	// SSH proxy public key.
	PublicKey pulumi.StringPtrInput
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSshLocalCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshLocalCaState)(nil)).Elem()
}

type firewallSshLocalCaArgs struct {
	// SSH proxy local CA name.
	Name *string `pulumi:"name"`
	// Password for SSH private key.
	Password *string `pulumi:"password"`
	// SSH proxy private key, encrypted with a password.
	PrivateKey string `pulumi:"privateKey"`
	// SSH proxy public key.
	PublicKey string `pulumi:"publicKey"`
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSshLocalCa resource.
type FirewallSshLocalCaArgs struct {
	// SSH proxy local CA name.
	Name pulumi.StringPtrInput
	// Password for SSH private key.
	Password pulumi.StringPtrInput
	// SSH proxy private key, encrypted with a password.
	PrivateKey pulumi.StringInput
	// SSH proxy public key.
	PublicKey pulumi.StringInput
	// SSH proxy local CA source type. Valid values: `built-in`, `user`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
	return reflect.TypeOf((*FirewallSshLocalCa)(nil))
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaOutput() FirewallSshLocalCaOutput {
	return i.ToFirewallSshLocalCaOutputWithContext(context.Background())
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaOutputWithContext(ctx context.Context) FirewallSshLocalCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaOutput)
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaPtrOutput() FirewallSshLocalCaPtrOutput {
	return i.ToFirewallSshLocalCaPtrOutputWithContext(context.Background())
}

func (i *FirewallSshLocalCa) ToFirewallSshLocalCaPtrOutputWithContext(ctx context.Context) FirewallSshLocalCaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaPtrOutput)
}

type FirewallSshLocalCaPtrInput interface {
	pulumi.Input

	ToFirewallSshLocalCaPtrOutput() FirewallSshLocalCaPtrOutput
	ToFirewallSshLocalCaPtrOutputWithContext(ctx context.Context) FirewallSshLocalCaPtrOutput
}

type firewallSshLocalCaPtrType FirewallSshLocalCaArgs

func (*firewallSshLocalCaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalCa)(nil))
}

func (i *firewallSshLocalCaPtrType) ToFirewallSshLocalCaPtrOutput() FirewallSshLocalCaPtrOutput {
	return i.ToFirewallSshLocalCaPtrOutputWithContext(context.Background())
}

func (i *firewallSshLocalCaPtrType) ToFirewallSshLocalCaPtrOutputWithContext(ctx context.Context) FirewallSshLocalCaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaPtrOutput)
}

// FirewallSshLocalCaArrayInput is an input type that accepts FirewallSshLocalCaArray and FirewallSshLocalCaArrayOutput values.
// You can construct a concrete instance of `FirewallSshLocalCaArrayInput` via:
//
//          FirewallSshLocalCaArray{ FirewallSshLocalCaArgs{...} }
type FirewallSshLocalCaArrayInput interface {
	pulumi.Input

	ToFirewallSshLocalCaArrayOutput() FirewallSshLocalCaArrayOutput
	ToFirewallSshLocalCaArrayOutputWithContext(context.Context) FirewallSshLocalCaArrayOutput
}

type FirewallSshLocalCaArray []FirewallSshLocalCaInput

func (FirewallSshLocalCaArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallSshLocalCa)(nil))
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
//          FirewallSshLocalCaMap{ "key": FirewallSshLocalCaArgs{...} }
type FirewallSshLocalCaMapInput interface {
	pulumi.Input

	ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput
	ToFirewallSshLocalCaMapOutputWithContext(context.Context) FirewallSshLocalCaMapOutput
}

type FirewallSshLocalCaMap map[string]FirewallSshLocalCaInput

func (FirewallSshLocalCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallSshLocalCa)(nil))
}

func (i FirewallSshLocalCaMap) ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput {
	return i.ToFirewallSshLocalCaMapOutputWithContext(context.Background())
}

func (i FirewallSshLocalCaMap) ToFirewallSshLocalCaMapOutputWithContext(ctx context.Context) FirewallSshLocalCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSshLocalCaMapOutput)
}

type FirewallSshLocalCaOutput struct {
	*pulumi.OutputState
}

func (FirewallSshLocalCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallSshLocalCa)(nil))
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaOutput() FirewallSshLocalCaOutput {
	return o
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaOutputWithContext(ctx context.Context) FirewallSshLocalCaOutput {
	return o
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaPtrOutput() FirewallSshLocalCaPtrOutput {
	return o.ToFirewallSshLocalCaPtrOutputWithContext(context.Background())
}

func (o FirewallSshLocalCaOutput) ToFirewallSshLocalCaPtrOutputWithContext(ctx context.Context) FirewallSshLocalCaPtrOutput {
	return o.ApplyT(func(v FirewallSshLocalCa) *FirewallSshLocalCa {
		return &v
	}).(FirewallSshLocalCaPtrOutput)
}

type FirewallSshLocalCaPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallSshLocalCaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSshLocalCa)(nil))
}

func (o FirewallSshLocalCaPtrOutput) ToFirewallSshLocalCaPtrOutput() FirewallSshLocalCaPtrOutput {
	return o
}

func (o FirewallSshLocalCaPtrOutput) ToFirewallSshLocalCaPtrOutputWithContext(ctx context.Context) FirewallSshLocalCaPtrOutput {
	return o
}

type FirewallSshLocalCaArrayOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallSshLocalCa)(nil))
}

func (o FirewallSshLocalCaArrayOutput) ToFirewallSshLocalCaArrayOutput() FirewallSshLocalCaArrayOutput {
	return o
}

func (o FirewallSshLocalCaArrayOutput) ToFirewallSshLocalCaArrayOutputWithContext(ctx context.Context) FirewallSshLocalCaArrayOutput {
	return o
}

func (o FirewallSshLocalCaArrayOutput) Index(i pulumi.IntInput) FirewallSshLocalCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallSshLocalCa {
		return vs[0].([]FirewallSshLocalCa)[vs[1].(int)]
	}).(FirewallSshLocalCaOutput)
}

type FirewallSshLocalCaMapOutput struct{ *pulumi.OutputState }

func (FirewallSshLocalCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallSshLocalCa)(nil))
}

func (o FirewallSshLocalCaMapOutput) ToFirewallSshLocalCaMapOutput() FirewallSshLocalCaMapOutput {
	return o
}

func (o FirewallSshLocalCaMapOutput) ToFirewallSshLocalCaMapOutputWithContext(ctx context.Context) FirewallSshLocalCaMapOutput {
	return o
}

func (o FirewallSshLocalCaMapOutput) MapIndex(k pulumi.StringInput) FirewallSshLocalCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallSshLocalCa {
		return vs[0].(map[string]FirewallSshLocalCa)[vs[1].(string)]
	}).(FirewallSshLocalCaOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallSshLocalCaOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalCaPtrOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalCaArrayOutput{})
	pulumi.RegisterOutputType(FirewallSshLocalCaMapOutput{})
}

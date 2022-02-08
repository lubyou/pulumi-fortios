// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSH proxy host public keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallSshHostKey(ctx, "trname", &fortios.FirewallSshHostKeyArgs{
// 			Hostname: pulumi.String("testmachine"),
// 			Ip:       pulumi.String("1.1.1.1"),
// 			Port:     pulumi.Int(22),
// 			Status:   pulumi.String("trusted"),
// 			Type:     pulumi.String("RSA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// FirewallSsh HostKey can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallSshHostKey:FirewallSshHostKey labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallSshHostKey struct {
	pulumi.CustomResourceState

	// Hostname of the SSH server.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// SSH public key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringOutput `pulumi:"nid"`
	// Port of the SSH server.
	Port pulumi.IntOutput `pulumi:"port"`
	// SSH public key.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringOutput `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSshHostKey registers a new resource with the given unique name, arguments, and options.
func NewFirewallSshHostKey(ctx *pulumi.Context,
	name string, args *FirewallSshHostKeyArgs, opts ...pulumi.ResourceOption) (*FirewallSshHostKey, error) {
	if args == nil {
		args = &FirewallSshHostKeyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
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
	// Hostname of the SSH server.
	Hostname *string `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip *string `pulumi:"ip"`
	// SSH public key name.
	Name *string `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid *string `pulumi:"nid"`
	// Port of the SSH server.
	Port *int `pulumi:"port"`
	// SSH public key.
	PublicKey *string `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status *string `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type *string `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallSshHostKeyState struct {
	// Hostname of the SSH server.
	Hostname pulumi.StringPtrInput
	// IP address of the SSH server.
	Ip pulumi.StringPtrInput
	// SSH public key name.
	Name pulumi.StringPtrInput
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringPtrInput
	// Port of the SSH server.
	Port pulumi.IntPtrInput
	// SSH public key.
	PublicKey pulumi.StringPtrInput
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringPtrInput
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringPtrInput
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSshHostKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSshHostKeyState)(nil)).Elem()
}

type firewallSshHostKeyArgs struct {
	// Hostname of the SSH server.
	Hostname *string `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip *string `pulumi:"ip"`
	// SSH public key name.
	Name *string `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid *string `pulumi:"nid"`
	// Port of the SSH server.
	Port *int `pulumi:"port"`
	// SSH public key.
	PublicKey *string `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status *string `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type *string `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSshHostKey resource.
type FirewallSshHostKeyArgs struct {
	// Hostname of the SSH server.
	Hostname pulumi.StringPtrInput
	// IP address of the SSH server.
	Ip pulumi.StringPtrInput
	// SSH public key name.
	Name pulumi.StringPtrInput
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringPtrInput
	// Port of the SSH server.
	Port pulumi.IntPtrInput
	// SSH public key.
	PublicKey pulumi.StringPtrInput
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringPtrInput
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringPtrInput
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
//          FirewallSshHostKeyArray{ FirewallSshHostKeyArgs{...} }
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
//          FirewallSshHostKeyMap{ "key": FirewallSshHostKeyArgs{...} }
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

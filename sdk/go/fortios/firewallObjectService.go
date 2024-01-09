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

type FirewallObjectService struct {
	pulumi.CustomResourceState

	Category       pulumi.StringOutput    `pulumi:"category"`
	Comment        pulumi.StringPtrOutput `pulumi:"comment"`
	Fqdn           pulumi.StringOutput    `pulumi:"fqdn"`
	Icmpcode       pulumi.StringOutput    `pulumi:"icmpcode"`
	Icmptype       pulumi.StringOutput    `pulumi:"icmptype"`
	Iprange        pulumi.StringOutput    `pulumi:"iprange"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Protocol       pulumi.StringOutput    `pulumi:"protocol"`
	ProtocolNumber pulumi.StringOutput    `pulumi:"protocolNumber"`
	SctpPortrange  pulumi.StringOutput    `pulumi:"sctpPortrange"`
	SessionTtl     pulumi.StringOutput    `pulumi:"sessionTtl"`
	TcpPortrange   pulumi.StringOutput    `pulumi:"tcpPortrange"`
	UdpPortrange   pulumi.StringOutput    `pulumi:"udpPortrange"`
}

// NewFirewallObjectService registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectService(ctx *pulumi.Context,
	name string, args *FirewallObjectServiceArgs, opts ...pulumi.ResourceOption) (*FirewallObjectService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallObjectService
	err := ctx.RegisterResource("fortios:index/firewallObjectService:FirewallObjectService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectService gets an existing FirewallObjectService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectServiceState, opts ...pulumi.ResourceOption) (*FirewallObjectService, error) {
	var resource FirewallObjectService
	err := ctx.ReadResource("fortios:index/firewallObjectService:FirewallObjectService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectService resources.
type firewallObjectServiceState struct {
	Category       *string `pulumi:"category"`
	Comment        *string `pulumi:"comment"`
	Fqdn           *string `pulumi:"fqdn"`
	Icmpcode       *string `pulumi:"icmpcode"`
	Icmptype       *string `pulumi:"icmptype"`
	Iprange        *string `pulumi:"iprange"`
	Name           *string `pulumi:"name"`
	Protocol       *string `pulumi:"protocol"`
	ProtocolNumber *string `pulumi:"protocolNumber"`
	SctpPortrange  *string `pulumi:"sctpPortrange"`
	SessionTtl     *string `pulumi:"sessionTtl"`
	TcpPortrange   *string `pulumi:"tcpPortrange"`
	UdpPortrange   *string `pulumi:"udpPortrange"`
}

type FirewallObjectServiceState struct {
	Category       pulumi.StringPtrInput
	Comment        pulumi.StringPtrInput
	Fqdn           pulumi.StringPtrInput
	Icmpcode       pulumi.StringPtrInput
	Icmptype       pulumi.StringPtrInput
	Iprange        pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Protocol       pulumi.StringPtrInput
	ProtocolNumber pulumi.StringPtrInput
	SctpPortrange  pulumi.StringPtrInput
	SessionTtl     pulumi.StringPtrInput
	TcpPortrange   pulumi.StringPtrInput
	UdpPortrange   pulumi.StringPtrInput
}

func (FirewallObjectServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectServiceState)(nil)).Elem()
}

type firewallObjectServiceArgs struct {
	Category       string  `pulumi:"category"`
	Comment        *string `pulumi:"comment"`
	Fqdn           *string `pulumi:"fqdn"`
	Icmpcode       *string `pulumi:"icmpcode"`
	Icmptype       *string `pulumi:"icmptype"`
	Iprange        *string `pulumi:"iprange"`
	Name           *string `pulumi:"name"`
	Protocol       string  `pulumi:"protocol"`
	ProtocolNumber *string `pulumi:"protocolNumber"`
	SctpPortrange  *string `pulumi:"sctpPortrange"`
	SessionTtl     *string `pulumi:"sessionTtl"`
	TcpPortrange   *string `pulumi:"tcpPortrange"`
	UdpPortrange   *string `pulumi:"udpPortrange"`
}

// The set of arguments for constructing a FirewallObjectService resource.
type FirewallObjectServiceArgs struct {
	Category       pulumi.StringInput
	Comment        pulumi.StringPtrInput
	Fqdn           pulumi.StringPtrInput
	Icmpcode       pulumi.StringPtrInput
	Icmptype       pulumi.StringPtrInput
	Iprange        pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Protocol       pulumi.StringInput
	ProtocolNumber pulumi.StringPtrInput
	SctpPortrange  pulumi.StringPtrInput
	SessionTtl     pulumi.StringPtrInput
	TcpPortrange   pulumi.StringPtrInput
	UdpPortrange   pulumi.StringPtrInput
}

func (FirewallObjectServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectServiceArgs)(nil)).Elem()
}

type FirewallObjectServiceInput interface {
	pulumi.Input

	ToFirewallObjectServiceOutput() FirewallObjectServiceOutput
	ToFirewallObjectServiceOutputWithContext(ctx context.Context) FirewallObjectServiceOutput
}

func (*FirewallObjectService) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectService)(nil)).Elem()
}

func (i *FirewallObjectService) ToFirewallObjectServiceOutput() FirewallObjectServiceOutput {
	return i.ToFirewallObjectServiceOutputWithContext(context.Background())
}

func (i *FirewallObjectService) ToFirewallObjectServiceOutputWithContext(ctx context.Context) FirewallObjectServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectServiceOutput)
}

// FirewallObjectServiceArrayInput is an input type that accepts FirewallObjectServiceArray and FirewallObjectServiceArrayOutput values.
// You can construct a concrete instance of `FirewallObjectServiceArrayInput` via:
//
//	FirewallObjectServiceArray{ FirewallObjectServiceArgs{...} }
type FirewallObjectServiceArrayInput interface {
	pulumi.Input

	ToFirewallObjectServiceArrayOutput() FirewallObjectServiceArrayOutput
	ToFirewallObjectServiceArrayOutputWithContext(context.Context) FirewallObjectServiceArrayOutput
}

type FirewallObjectServiceArray []FirewallObjectServiceInput

func (FirewallObjectServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectService)(nil)).Elem()
}

func (i FirewallObjectServiceArray) ToFirewallObjectServiceArrayOutput() FirewallObjectServiceArrayOutput {
	return i.ToFirewallObjectServiceArrayOutputWithContext(context.Background())
}

func (i FirewallObjectServiceArray) ToFirewallObjectServiceArrayOutputWithContext(ctx context.Context) FirewallObjectServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectServiceArrayOutput)
}

// FirewallObjectServiceMapInput is an input type that accepts FirewallObjectServiceMap and FirewallObjectServiceMapOutput values.
// You can construct a concrete instance of `FirewallObjectServiceMapInput` via:
//
//	FirewallObjectServiceMap{ "key": FirewallObjectServiceArgs{...} }
type FirewallObjectServiceMapInput interface {
	pulumi.Input

	ToFirewallObjectServiceMapOutput() FirewallObjectServiceMapOutput
	ToFirewallObjectServiceMapOutputWithContext(context.Context) FirewallObjectServiceMapOutput
}

type FirewallObjectServiceMap map[string]FirewallObjectServiceInput

func (FirewallObjectServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectService)(nil)).Elem()
}

func (i FirewallObjectServiceMap) ToFirewallObjectServiceMapOutput() FirewallObjectServiceMapOutput {
	return i.ToFirewallObjectServiceMapOutputWithContext(context.Background())
}

func (i FirewallObjectServiceMap) ToFirewallObjectServiceMapOutputWithContext(ctx context.Context) FirewallObjectServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectServiceMapOutput)
}

type FirewallObjectServiceOutput struct{ *pulumi.OutputState }

func (FirewallObjectServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectService)(nil)).Elem()
}

func (o FirewallObjectServiceOutput) ToFirewallObjectServiceOutput() FirewallObjectServiceOutput {
	return o
}

func (o FirewallObjectServiceOutput) ToFirewallObjectServiceOutputWithContext(ctx context.Context) FirewallObjectServiceOutput {
	return o
}

func (o FirewallObjectServiceOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallObjectServiceOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Icmpcode() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Icmpcode }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Icmptype() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Icmptype }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Iprange() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Iprange }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) ProtocolNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.ProtocolNumber }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) SctpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.SctpPortrange }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) SessionTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.SessionTtl }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) TcpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.TcpPortrange }).(pulumi.StringOutput)
}

func (o FirewallObjectServiceOutput) UdpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectService) pulumi.StringOutput { return v.UdpPortrange }).(pulumi.StringOutput)
}

type FirewallObjectServiceArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectService)(nil)).Elem()
}

func (o FirewallObjectServiceArrayOutput) ToFirewallObjectServiceArrayOutput() FirewallObjectServiceArrayOutput {
	return o
}

func (o FirewallObjectServiceArrayOutput) ToFirewallObjectServiceArrayOutputWithContext(ctx context.Context) FirewallObjectServiceArrayOutput {
	return o
}

func (o FirewallObjectServiceArrayOutput) Index(i pulumi.IntInput) FirewallObjectServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectService {
		return vs[0].([]*FirewallObjectService)[vs[1].(int)]
	}).(FirewallObjectServiceOutput)
}

type FirewallObjectServiceMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectService)(nil)).Elem()
}

func (o FirewallObjectServiceMapOutput) ToFirewallObjectServiceMapOutput() FirewallObjectServiceMapOutput {
	return o
}

func (o FirewallObjectServiceMapOutput) ToFirewallObjectServiceMapOutputWithContext(ctx context.Context) FirewallObjectServiceMapOutput {
	return o
}

func (o FirewallObjectServiceMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectService {
		return vs[0].(map[string]*FirewallObjectService)[vs[1].(string)]
	}).(FirewallObjectServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectServiceInput)(nil)).Elem(), &FirewallObjectService{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectServiceArrayInput)(nil)).Elem(), FirewallObjectServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectServiceMapInput)(nil)).Elem(), FirewallObjectServiceMap{})
	pulumi.RegisterOutputType(FirewallObjectServiceOutput{})
	pulumi.RegisterOutputType(FirewallObjectServiceArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectServiceMapOutput{})
}

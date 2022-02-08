// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.
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
// 		_, err := fortios.NewFortimanagerFirewallObjectVip(ctx, "test1", &fortios.FortimanagerFirewallObjectVipArgs{
// 			ArpReply:      pulumi.String("enable"),
// 			Comment:       pulumi.String("test obj vip"),
// 			ConfigDefault: pulumi.String("enable"),
// 			ExtIntf:       pulumi.String("any"),
// 			ExtIp:         pulumi.String("2.2.2.2"),
// 			MappedIp:      pulumi.String("1.1.1.1"),
// 			Type:          pulumi.String("static-nat"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewFortimanagerFirewallObjectVip(ctx, "test2", &fortios.FortimanagerFirewallObjectVipArgs{
// 			ArpReply:      pulumi.String("disable"),
// 			Comment:       pulumi.String("test obj vip"),
// 			ConfigDefault: pulumi.String("enable"),
// 			ExtIp:         pulumi.String("2.2.2.2-2.2.2.100"),
// 			MappedAddr:    pulumi.String("update.microsoft.com"),
// 			Type:          pulumi.String("fqdn"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerFirewallObjectVip struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrOutput `pulumi:"arpReply"`
	// Comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrOutput `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrOutput `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrOutput `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrOutput `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp pulumi.StringPtrOutput `pulumi:"mappedIp"`
	// Virtual IP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFortimanagerFirewallObjectVip registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerFirewallObjectVip(ctx *pulumi.Context,
	name string, args *FortimanagerFirewallObjectVipArgs, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectVip, error) {
	if args == nil {
		args = &FortimanagerFirewallObjectVipArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerFirewallObjectVip
	err := ctx.RegisterResource("fortios:index/fortimanagerFirewallObjectVip:FortimanagerFirewallObjectVip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerFirewallObjectVip gets an existing FortimanagerFirewallObjectVip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerFirewallObjectVip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerFirewallObjectVipState, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectVip, error) {
	var resource FortimanagerFirewallObjectVip
	err := ctx.ReadResource("fortios:index/fortimanagerFirewallObjectVip:FortimanagerFirewallObjectVip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerFirewallObjectVip resources.
type fortimanagerFirewallObjectVipState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply *string `pulumi:"arpReply"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault *string `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf *string `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp *string `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr *string `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp *string `pulumi:"mappedIp"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type *string `pulumi:"type"`
}

type FortimanagerFirewallObjectVipState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrInput
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrInput
	// Mapped Ip address.
	MappedIp pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectVipState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectVipState)(nil)).Elem()
}

type fortimanagerFirewallObjectVipArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply *string `pulumi:"arpReply"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Enable/Disable config default value. enabled by default.
	ConfigDefault *string `pulumi:"configDefault"`
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf *string `pulumi:"extIntf"`
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp *string `pulumi:"extIp"`
	// Mapped FQDN address name.
	MappedAddr *string `pulumi:"mappedAddr"`
	// Mapped Ip address.
	MappedIp *string `pulumi:"mappedIp"`
	// Virtual IP name.
	Name *string `pulumi:"name"`
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FortimanagerFirewallObjectVip resource.
type FortimanagerFirewallObjectVipArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Enable to respond to ARP requests for this virtual IP address. Enabled by default.
	ArpReply pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Enable/Disable config default value. enabled by default.
	ConfigDefault pulumi.StringPtrInput
	// Interface connected to the source network that receives the packets that will be forwarded to the destination network.
	ExtIntf pulumi.StringPtrInput
	// IP address or address range on the external interface that you want to map to an address or address range on the destination network.
	ExtIp pulumi.StringPtrInput
	// Mapped FQDN address name.
	MappedAddr pulumi.StringPtrInput
	// Mapped Ip address.
	MappedIp pulumi.StringPtrInput
	// Virtual IP name.
	Name pulumi.StringPtrInput
	// Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
	Type pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectVipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectVipArgs)(nil)).Elem()
}

type FortimanagerFirewallObjectVipInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectVipOutput() FortimanagerFirewallObjectVipOutput
	ToFortimanagerFirewallObjectVipOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipOutput
}

func (*FortimanagerFirewallObjectVip) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (i *FortimanagerFirewallObjectVip) ToFortimanagerFirewallObjectVipOutput() FortimanagerFirewallObjectVipOutput {
	return i.ToFortimanagerFirewallObjectVipOutputWithContext(context.Background())
}

func (i *FortimanagerFirewallObjectVip) ToFortimanagerFirewallObjectVipOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectVipOutput)
}

// FortimanagerFirewallObjectVipArrayInput is an input type that accepts FortimanagerFirewallObjectVipArray and FortimanagerFirewallObjectVipArrayOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectVipArrayInput` via:
//
//          FortimanagerFirewallObjectVipArray{ FortimanagerFirewallObjectVipArgs{...} }
type FortimanagerFirewallObjectVipArrayInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectVipArrayOutput() FortimanagerFirewallObjectVipArrayOutput
	ToFortimanagerFirewallObjectVipArrayOutputWithContext(context.Context) FortimanagerFirewallObjectVipArrayOutput
}

type FortimanagerFirewallObjectVipArray []FortimanagerFirewallObjectVipInput

func (FortimanagerFirewallObjectVipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (i FortimanagerFirewallObjectVipArray) ToFortimanagerFirewallObjectVipArrayOutput() FortimanagerFirewallObjectVipArrayOutput {
	return i.ToFortimanagerFirewallObjectVipArrayOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectVipArray) ToFortimanagerFirewallObjectVipArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectVipArrayOutput)
}

// FortimanagerFirewallObjectVipMapInput is an input type that accepts FortimanagerFirewallObjectVipMap and FortimanagerFirewallObjectVipMapOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectVipMapInput` via:
//
//          FortimanagerFirewallObjectVipMap{ "key": FortimanagerFirewallObjectVipArgs{...} }
type FortimanagerFirewallObjectVipMapInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectVipMapOutput() FortimanagerFirewallObjectVipMapOutput
	ToFortimanagerFirewallObjectVipMapOutputWithContext(context.Context) FortimanagerFirewallObjectVipMapOutput
}

type FortimanagerFirewallObjectVipMap map[string]FortimanagerFirewallObjectVipInput

func (FortimanagerFirewallObjectVipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (i FortimanagerFirewallObjectVipMap) ToFortimanagerFirewallObjectVipMapOutput() FortimanagerFirewallObjectVipMapOutput {
	return i.ToFortimanagerFirewallObjectVipMapOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectVipMap) ToFortimanagerFirewallObjectVipMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectVipMapOutput)
}

type FortimanagerFirewallObjectVipOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectVipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (o FortimanagerFirewallObjectVipOutput) ToFortimanagerFirewallObjectVipOutput() FortimanagerFirewallObjectVipOutput {
	return o
}

func (o FortimanagerFirewallObjectVipOutput) ToFortimanagerFirewallObjectVipOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipOutput {
	return o
}

type FortimanagerFirewallObjectVipArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectVipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (o FortimanagerFirewallObjectVipArrayOutput) ToFortimanagerFirewallObjectVipArrayOutput() FortimanagerFirewallObjectVipArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectVipArrayOutput) ToFortimanagerFirewallObjectVipArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectVipArrayOutput) Index(i pulumi.IntInput) FortimanagerFirewallObjectVipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerFirewallObjectVip {
		return vs[0].([]*FortimanagerFirewallObjectVip)[vs[1].(int)]
	}).(FortimanagerFirewallObjectVipOutput)
}

type FortimanagerFirewallObjectVipMapOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectVipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallObjectVip)(nil)).Elem()
}

func (o FortimanagerFirewallObjectVipMapOutput) ToFortimanagerFirewallObjectVipMapOutput() FortimanagerFirewallObjectVipMapOutput {
	return o
}

func (o FortimanagerFirewallObjectVipMapOutput) ToFortimanagerFirewallObjectVipMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectVipMapOutput {
	return o
}

func (o FortimanagerFirewallObjectVipMapOutput) MapIndex(k pulumi.StringInput) FortimanagerFirewallObjectVipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerFirewallObjectVip {
		return vs[0].(map[string]*FortimanagerFirewallObjectVip)[vs[1].(string)]
	}).(FortimanagerFirewallObjectVipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectVipInput)(nil)).Elem(), &FortimanagerFirewallObjectVip{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectVipArrayInput)(nil)).Elem(), FortimanagerFirewallObjectVipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectVipMapInput)(nil)).Elem(), FortimanagerFirewallObjectVipMap{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectVipOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectVipArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectVipMapOutput{})
}

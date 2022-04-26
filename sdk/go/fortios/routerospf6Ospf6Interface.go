// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OSPF6 interface configuration.
//
// > The provider supports the definition of Ospf6-Interface in Router Ospf6 `RouterOspf6`, and also allows the definition of separate Ospf6-Interface resources `Routerospf6Ospf6Interface`, but do not use a `RouterOspf6` with in-line Ospf6-Interface in conjunction with any `Routerospf6Ospf6Interface` resources, otherwise conflicts and overwrite will occur.
//
// ## Import
//
// Routerospf6 Ospf6Interface can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type Routerospf6Ospf6Interface struct {
	pulumi.CustomResourceState

	// A.B.C.D, in IPv4 address format.
	AreaId pulumi.StringOutput `pulumi:"areaId"`
	// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringOutput `pulumi:"bfd"`
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntOutput `pulumi:"cost"`
	// Dead interval.
	DeadInterval pulumi.IntOutput `pulumi:"deadInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Hello interval.
	HelloInterval pulumi.IntOutput `pulumi:"helloInterval"`
	// Configuration interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	IpsecAuthAlg pulumi.StringOutput `pulumi:"ipsecAuthAlg"`
	// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
	IpsecEncAlg pulumi.StringOutput `pulumi:"ipsecEncAlg"`
	// IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
	IpsecKeys Routerospf6Ospf6InterfaceIpsecKeyArrayOutput `pulumi:"ipsecKeys"`
	// Key roll-over interval.
	KeyRolloverInterval pulumi.IntOutput `pulumi:"keyRolloverInterval"`
	// MTU for OSPFv3 packets.
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
	MtuIgnore pulumi.StringOutput `pulumi:"mtuIgnore"`
	// Interface entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors Routerospf6Ospf6InterfaceNeighborArrayOutput `pulumi:"neighbors"`
	// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// priority
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Retransmit interval.
	RetransmitInterval pulumi.IntOutput `pulumi:"retransmitInterval"`
	// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Transmit delay.
	TransmitDelay pulumi.IntOutput `pulumi:"transmitDelay"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterospf6Ospf6Interface registers a new resource with the given unique name, arguments, and options.
func NewRouterospf6Ospf6Interface(ctx *pulumi.Context,
	name string, args *Routerospf6Ospf6InterfaceArgs, opts ...pulumi.ResourceOption) (*Routerospf6Ospf6Interface, error) {
	if args == nil {
		args = &Routerospf6Ospf6InterfaceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Routerospf6Ospf6Interface
	err := ctx.RegisterResource("fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterospf6Ospf6Interface gets an existing Routerospf6Ospf6Interface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterospf6Ospf6Interface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Routerospf6Ospf6InterfaceState, opts ...pulumi.ResourceOption) (*Routerospf6Ospf6Interface, error) {
	var resource Routerospf6Ospf6Interface
	err := ctx.ReadResource("fortios:index/routerospf6Ospf6Interface:Routerospf6Ospf6Interface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Routerospf6Ospf6Interface resources.
type routerospf6Ospf6InterfaceState struct {
	// A.B.C.D, in IPv4 address format.
	AreaId *string `pulumi:"areaId"`
	// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
	Authentication *string `pulumi:"authentication"`
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Dead interval.
	DeadInterval *int `pulumi:"deadInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hello interval.
	HelloInterval *int `pulumi:"helloInterval"`
	// Configuration interface name.
	Interface *string `pulumi:"interface"`
	// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	IpsecAuthAlg *string `pulumi:"ipsecAuthAlg"`
	// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
	IpsecEncAlg *string `pulumi:"ipsecEncAlg"`
	// IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
	IpsecKeys []Routerospf6Ospf6InterfaceIpsecKey `pulumi:"ipsecKeys"`
	// Key roll-over interval.
	KeyRolloverInterval *int `pulumi:"keyRolloverInterval"`
	// MTU for OSPFv3 packets.
	Mtu *int `pulumi:"mtu"`
	// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
	MtuIgnore *string `pulumi:"mtuIgnore"`
	// Interface entry name.
	Name *string `pulumi:"name"`
	// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []Routerospf6Ospf6InterfaceNeighbor `pulumi:"neighbors"`
	// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
	NetworkType *string `pulumi:"networkType"`
	// priority
	Priority *int `pulumi:"priority"`
	// Retransmit interval.
	RetransmitInterval *int `pulumi:"retransmitInterval"`
	// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Transmit delay.
	TransmitDelay *int `pulumi:"transmitDelay"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Routerospf6Ospf6InterfaceState struct {
	// A.B.C.D, in IPv4 address format.
	AreaId pulumi.StringPtrInput
	// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
	Authentication pulumi.StringPtrInput
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Dead interval.
	DeadInterval pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Hello interval.
	HelloInterval pulumi.IntPtrInput
	// Configuration interface name.
	Interface pulumi.StringPtrInput
	// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	IpsecAuthAlg pulumi.StringPtrInput
	// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
	IpsecEncAlg pulumi.StringPtrInput
	// IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
	IpsecKeys Routerospf6Ospf6InterfaceIpsecKeyArrayInput
	// Key roll-over interval.
	KeyRolloverInterval pulumi.IntPtrInput
	// MTU for OSPFv3 packets.
	Mtu pulumi.IntPtrInput
	// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
	MtuIgnore pulumi.StringPtrInput
	// Interface entry name.
	Name pulumi.StringPtrInput
	// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors Routerospf6Ospf6InterfaceNeighborArrayInput
	// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
	NetworkType pulumi.StringPtrInput
	// priority
	Priority pulumi.IntPtrInput
	// Retransmit interval.
	RetransmitInterval pulumi.IntPtrInput
	// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Transmit delay.
	TransmitDelay pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Routerospf6Ospf6InterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospf6Ospf6InterfaceState)(nil)).Elem()
}

type routerospf6Ospf6InterfaceArgs struct {
	// A.B.C.D, in IPv4 address format.
	AreaId *string `pulumi:"areaId"`
	// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
	Authentication *string `pulumi:"authentication"`
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Dead interval.
	DeadInterval *int `pulumi:"deadInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hello interval.
	HelloInterval *int `pulumi:"helloInterval"`
	// Configuration interface name.
	Interface *string `pulumi:"interface"`
	// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	IpsecAuthAlg *string `pulumi:"ipsecAuthAlg"`
	// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
	IpsecEncAlg *string `pulumi:"ipsecEncAlg"`
	// IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
	IpsecKeys []Routerospf6Ospf6InterfaceIpsecKey `pulumi:"ipsecKeys"`
	// Key roll-over interval.
	KeyRolloverInterval *int `pulumi:"keyRolloverInterval"`
	// MTU for OSPFv3 packets.
	Mtu *int `pulumi:"mtu"`
	// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
	MtuIgnore *string `pulumi:"mtuIgnore"`
	// Interface entry name.
	Name *string `pulumi:"name"`
	// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []Routerospf6Ospf6InterfaceNeighbor `pulumi:"neighbors"`
	// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
	NetworkType *string `pulumi:"networkType"`
	// priority
	Priority *int `pulumi:"priority"`
	// Retransmit interval.
	RetransmitInterval *int `pulumi:"retransmitInterval"`
	// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Transmit delay.
	TransmitDelay *int `pulumi:"transmitDelay"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Routerospf6Ospf6Interface resource.
type Routerospf6Ospf6InterfaceArgs struct {
	// A.B.C.D, in IPv4 address format.
	AreaId pulumi.StringPtrInput
	// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
	Authentication pulumi.StringPtrInput
	// Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Dead interval.
	DeadInterval pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Hello interval.
	HelloInterval pulumi.IntPtrInput
	// Configuration interface name.
	Interface pulumi.StringPtrInput
	// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
	IpsecAuthAlg pulumi.StringPtrInput
	// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
	IpsecEncAlg pulumi.StringPtrInput
	// IPsec authentication and encryption keys. The structure of `ipsecKeys` block is documented below.
	IpsecKeys Routerospf6Ospf6InterfaceIpsecKeyArrayInput
	// Key roll-over interval.
	KeyRolloverInterval pulumi.IntPtrInput
	// MTU for OSPFv3 packets.
	Mtu pulumi.IntPtrInput
	// Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
	MtuIgnore pulumi.StringPtrInput
	// Interface entry name.
	Name pulumi.StringPtrInput
	// OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors Routerospf6Ospf6InterfaceNeighborArrayInput
	// Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
	NetworkType pulumi.StringPtrInput
	// priority
	Priority pulumi.IntPtrInput
	// Retransmit interval.
	RetransmitInterval pulumi.IntPtrInput
	// Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Transmit delay.
	TransmitDelay pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Routerospf6Ospf6InterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospf6Ospf6InterfaceArgs)(nil)).Elem()
}

type Routerospf6Ospf6InterfaceInput interface {
	pulumi.Input

	ToRouterospf6Ospf6InterfaceOutput() Routerospf6Ospf6InterfaceOutput
	ToRouterospf6Ospf6InterfaceOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceOutput
}

func (*Routerospf6Ospf6Interface) ElementType() reflect.Type {
	return reflect.TypeOf((**Routerospf6Ospf6Interface)(nil)).Elem()
}

func (i *Routerospf6Ospf6Interface) ToRouterospf6Ospf6InterfaceOutput() Routerospf6Ospf6InterfaceOutput {
	return i.ToRouterospf6Ospf6InterfaceOutputWithContext(context.Background())
}

func (i *Routerospf6Ospf6Interface) ToRouterospf6Ospf6InterfaceOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Routerospf6Ospf6InterfaceOutput)
}

// Routerospf6Ospf6InterfaceArrayInput is an input type that accepts Routerospf6Ospf6InterfaceArray and Routerospf6Ospf6InterfaceArrayOutput values.
// You can construct a concrete instance of `Routerospf6Ospf6InterfaceArrayInput` via:
//
//          Routerospf6Ospf6InterfaceArray{ Routerospf6Ospf6InterfaceArgs{...} }
type Routerospf6Ospf6InterfaceArrayInput interface {
	pulumi.Input

	ToRouterospf6Ospf6InterfaceArrayOutput() Routerospf6Ospf6InterfaceArrayOutput
	ToRouterospf6Ospf6InterfaceArrayOutputWithContext(context.Context) Routerospf6Ospf6InterfaceArrayOutput
}

type Routerospf6Ospf6InterfaceArray []Routerospf6Ospf6InterfaceInput

func (Routerospf6Ospf6InterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Routerospf6Ospf6Interface)(nil)).Elem()
}

func (i Routerospf6Ospf6InterfaceArray) ToRouterospf6Ospf6InterfaceArrayOutput() Routerospf6Ospf6InterfaceArrayOutput {
	return i.ToRouterospf6Ospf6InterfaceArrayOutputWithContext(context.Background())
}

func (i Routerospf6Ospf6InterfaceArray) ToRouterospf6Ospf6InterfaceArrayOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Routerospf6Ospf6InterfaceArrayOutput)
}

// Routerospf6Ospf6InterfaceMapInput is an input type that accepts Routerospf6Ospf6InterfaceMap and Routerospf6Ospf6InterfaceMapOutput values.
// You can construct a concrete instance of `Routerospf6Ospf6InterfaceMapInput` via:
//
//          Routerospf6Ospf6InterfaceMap{ "key": Routerospf6Ospf6InterfaceArgs{...} }
type Routerospf6Ospf6InterfaceMapInput interface {
	pulumi.Input

	ToRouterospf6Ospf6InterfaceMapOutput() Routerospf6Ospf6InterfaceMapOutput
	ToRouterospf6Ospf6InterfaceMapOutputWithContext(context.Context) Routerospf6Ospf6InterfaceMapOutput
}

type Routerospf6Ospf6InterfaceMap map[string]Routerospf6Ospf6InterfaceInput

func (Routerospf6Ospf6InterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Routerospf6Ospf6Interface)(nil)).Elem()
}

func (i Routerospf6Ospf6InterfaceMap) ToRouterospf6Ospf6InterfaceMapOutput() Routerospf6Ospf6InterfaceMapOutput {
	return i.ToRouterospf6Ospf6InterfaceMapOutputWithContext(context.Background())
}

func (i Routerospf6Ospf6InterfaceMap) ToRouterospf6Ospf6InterfaceMapOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Routerospf6Ospf6InterfaceMapOutput)
}

type Routerospf6Ospf6InterfaceOutput struct{ *pulumi.OutputState }

func (Routerospf6Ospf6InterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Routerospf6Ospf6Interface)(nil)).Elem()
}

func (o Routerospf6Ospf6InterfaceOutput) ToRouterospf6Ospf6InterfaceOutput() Routerospf6Ospf6InterfaceOutput {
	return o
}

func (o Routerospf6Ospf6InterfaceOutput) ToRouterospf6Ospf6InterfaceOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceOutput {
	return o
}

type Routerospf6Ospf6InterfaceArrayOutput struct{ *pulumi.OutputState }

func (Routerospf6Ospf6InterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Routerospf6Ospf6Interface)(nil)).Elem()
}

func (o Routerospf6Ospf6InterfaceArrayOutput) ToRouterospf6Ospf6InterfaceArrayOutput() Routerospf6Ospf6InterfaceArrayOutput {
	return o
}

func (o Routerospf6Ospf6InterfaceArrayOutput) ToRouterospf6Ospf6InterfaceArrayOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceArrayOutput {
	return o
}

func (o Routerospf6Ospf6InterfaceArrayOutput) Index(i pulumi.IntInput) Routerospf6Ospf6InterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Routerospf6Ospf6Interface {
		return vs[0].([]*Routerospf6Ospf6Interface)[vs[1].(int)]
	}).(Routerospf6Ospf6InterfaceOutput)
}

type Routerospf6Ospf6InterfaceMapOutput struct{ *pulumi.OutputState }

func (Routerospf6Ospf6InterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Routerospf6Ospf6Interface)(nil)).Elem()
}

func (o Routerospf6Ospf6InterfaceMapOutput) ToRouterospf6Ospf6InterfaceMapOutput() Routerospf6Ospf6InterfaceMapOutput {
	return o
}

func (o Routerospf6Ospf6InterfaceMapOutput) ToRouterospf6Ospf6InterfaceMapOutputWithContext(ctx context.Context) Routerospf6Ospf6InterfaceMapOutput {
	return o
}

func (o Routerospf6Ospf6InterfaceMapOutput) MapIndex(k pulumi.StringInput) Routerospf6Ospf6InterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Routerospf6Ospf6Interface {
		return vs[0].(map[string]*Routerospf6Ospf6Interface)[vs[1].(string)]
	}).(Routerospf6Ospf6InterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Routerospf6Ospf6InterfaceInput)(nil)).Elem(), &Routerospf6Ospf6Interface{})
	pulumi.RegisterInputType(reflect.TypeOf((*Routerospf6Ospf6InterfaceArrayInput)(nil)).Elem(), Routerospf6Ospf6InterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Routerospf6Ospf6InterfaceMapInput)(nil)).Elem(), Routerospf6Ospf6InterfaceMap{})
	pulumi.RegisterOutputType(Routerospf6Ospf6InterfaceOutput{})
	pulumi.RegisterOutputType(Routerospf6Ospf6InterfaceArrayOutput{})
	pulumi.RegisterOutputType(Routerospf6Ospf6InterfaceMapOutput{})
}

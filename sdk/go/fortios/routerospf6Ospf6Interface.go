// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Routerospf6Ospf6Interface struct {
	pulumi.CustomResourceState

	AreaId              pulumi.StringOutput                          `pulumi:"areaId"`
	Authentication      pulumi.StringOutput                          `pulumi:"authentication"`
	Bfd                 pulumi.StringOutput                          `pulumi:"bfd"`
	Cost                pulumi.IntOutput                             `pulumi:"cost"`
	DeadInterval        pulumi.IntOutput                             `pulumi:"deadInterval"`
	DynamicSortSubtable pulumi.StringPtrOutput                       `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                       `pulumi:"getAllTables"`
	HelloInterval       pulumi.IntOutput                             `pulumi:"helloInterval"`
	Interface           pulumi.StringOutput                          `pulumi:"interface"`
	IpsecAuthAlg        pulumi.StringOutput                          `pulumi:"ipsecAuthAlg"`
	IpsecEncAlg         pulumi.StringOutput                          `pulumi:"ipsecEncAlg"`
	IpsecKeys           Routerospf6Ospf6InterfaceIpsecKeyArrayOutput `pulumi:"ipsecKeys"`
	KeyRolloverInterval pulumi.IntOutput                             `pulumi:"keyRolloverInterval"`
	Mtu                 pulumi.IntOutput                             `pulumi:"mtu"`
	MtuIgnore           pulumi.StringOutput                          `pulumi:"mtuIgnore"`
	Name                pulumi.StringOutput                          `pulumi:"name"`
	Neighbors           Routerospf6Ospf6InterfaceNeighborArrayOutput `pulumi:"neighbors"`
	NetworkType         pulumi.StringOutput                          `pulumi:"networkType"`
	Priority            pulumi.IntOutput                             `pulumi:"priority"`
	RetransmitInterval  pulumi.IntOutput                             `pulumi:"retransmitInterval"`
	Status              pulumi.StringOutput                          `pulumi:"status"`
	TransmitDelay       pulumi.IntOutput                             `pulumi:"transmitDelay"`
	Vdomparam           pulumi.StringPtrOutput                       `pulumi:"vdomparam"`
}

// NewRouterospf6Ospf6Interface registers a new resource with the given unique name, arguments, and options.
func NewRouterospf6Ospf6Interface(ctx *pulumi.Context,
	name string, args *Routerospf6Ospf6InterfaceArgs, opts ...pulumi.ResourceOption) (*Routerospf6Ospf6Interface, error) {
	if args == nil {
		args = &Routerospf6Ospf6InterfaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	AreaId              *string                             `pulumi:"areaId"`
	Authentication      *string                             `pulumi:"authentication"`
	Bfd                 *string                             `pulumi:"bfd"`
	Cost                *int                                `pulumi:"cost"`
	DeadInterval        *int                                `pulumi:"deadInterval"`
	DynamicSortSubtable *string                             `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                             `pulumi:"getAllTables"`
	HelloInterval       *int                                `pulumi:"helloInterval"`
	Interface           *string                             `pulumi:"interface"`
	IpsecAuthAlg        *string                             `pulumi:"ipsecAuthAlg"`
	IpsecEncAlg         *string                             `pulumi:"ipsecEncAlg"`
	IpsecKeys           []Routerospf6Ospf6InterfaceIpsecKey `pulumi:"ipsecKeys"`
	KeyRolloverInterval *int                                `pulumi:"keyRolloverInterval"`
	Mtu                 *int                                `pulumi:"mtu"`
	MtuIgnore           *string                             `pulumi:"mtuIgnore"`
	Name                *string                             `pulumi:"name"`
	Neighbors           []Routerospf6Ospf6InterfaceNeighbor `pulumi:"neighbors"`
	NetworkType         *string                             `pulumi:"networkType"`
	Priority            *int                                `pulumi:"priority"`
	RetransmitInterval  *int                                `pulumi:"retransmitInterval"`
	Status              *string                             `pulumi:"status"`
	TransmitDelay       *int                                `pulumi:"transmitDelay"`
	Vdomparam           *string                             `pulumi:"vdomparam"`
}

type Routerospf6Ospf6InterfaceState struct {
	AreaId              pulumi.StringPtrInput
	Authentication      pulumi.StringPtrInput
	Bfd                 pulumi.StringPtrInput
	Cost                pulumi.IntPtrInput
	DeadInterval        pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HelloInterval       pulumi.IntPtrInput
	Interface           pulumi.StringPtrInput
	IpsecAuthAlg        pulumi.StringPtrInput
	IpsecEncAlg         pulumi.StringPtrInput
	IpsecKeys           Routerospf6Ospf6InterfaceIpsecKeyArrayInput
	KeyRolloverInterval pulumi.IntPtrInput
	Mtu                 pulumi.IntPtrInput
	MtuIgnore           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Neighbors           Routerospf6Ospf6InterfaceNeighborArrayInput
	NetworkType         pulumi.StringPtrInput
	Priority            pulumi.IntPtrInput
	RetransmitInterval  pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	TransmitDelay       pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (Routerospf6Ospf6InterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospf6Ospf6InterfaceState)(nil)).Elem()
}

type routerospf6Ospf6InterfaceArgs struct {
	AreaId              *string                             `pulumi:"areaId"`
	Authentication      *string                             `pulumi:"authentication"`
	Bfd                 *string                             `pulumi:"bfd"`
	Cost                *int                                `pulumi:"cost"`
	DeadInterval        *int                                `pulumi:"deadInterval"`
	DynamicSortSubtable *string                             `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                             `pulumi:"getAllTables"`
	HelloInterval       *int                                `pulumi:"helloInterval"`
	Interface           *string                             `pulumi:"interface"`
	IpsecAuthAlg        *string                             `pulumi:"ipsecAuthAlg"`
	IpsecEncAlg         *string                             `pulumi:"ipsecEncAlg"`
	IpsecKeys           []Routerospf6Ospf6InterfaceIpsecKey `pulumi:"ipsecKeys"`
	KeyRolloverInterval *int                                `pulumi:"keyRolloverInterval"`
	Mtu                 *int                                `pulumi:"mtu"`
	MtuIgnore           *string                             `pulumi:"mtuIgnore"`
	Name                *string                             `pulumi:"name"`
	Neighbors           []Routerospf6Ospf6InterfaceNeighbor `pulumi:"neighbors"`
	NetworkType         *string                             `pulumi:"networkType"`
	Priority            *int                                `pulumi:"priority"`
	RetransmitInterval  *int                                `pulumi:"retransmitInterval"`
	Status              *string                             `pulumi:"status"`
	TransmitDelay       *int                                `pulumi:"transmitDelay"`
	Vdomparam           *string                             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Routerospf6Ospf6Interface resource.
type Routerospf6Ospf6InterfaceArgs struct {
	AreaId              pulumi.StringPtrInput
	Authentication      pulumi.StringPtrInput
	Bfd                 pulumi.StringPtrInput
	Cost                pulumi.IntPtrInput
	DeadInterval        pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HelloInterval       pulumi.IntPtrInput
	Interface           pulumi.StringPtrInput
	IpsecAuthAlg        pulumi.StringPtrInput
	IpsecEncAlg         pulumi.StringPtrInput
	IpsecKeys           Routerospf6Ospf6InterfaceIpsecKeyArrayInput
	KeyRolloverInterval pulumi.IntPtrInput
	Mtu                 pulumi.IntPtrInput
	MtuIgnore           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Neighbors           Routerospf6Ospf6InterfaceNeighborArrayInput
	NetworkType         pulumi.StringPtrInput
	Priority            pulumi.IntPtrInput
	RetransmitInterval  pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	TransmitDelay       pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
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

func (i *Routerospf6Ospf6Interface) ToOutput(ctx context.Context) pulumix.Output[*Routerospf6Ospf6Interface] {
	return pulumix.Output[*Routerospf6Ospf6Interface]{
		OutputState: i.ToRouterospf6Ospf6InterfaceOutputWithContext(ctx).OutputState,
	}
}

// Routerospf6Ospf6InterfaceArrayInput is an input type that accepts Routerospf6Ospf6InterfaceArray and Routerospf6Ospf6InterfaceArrayOutput values.
// You can construct a concrete instance of `Routerospf6Ospf6InterfaceArrayInput` via:
//
//	Routerospf6Ospf6InterfaceArray{ Routerospf6Ospf6InterfaceArgs{...} }
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

func (i Routerospf6Ospf6InterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Routerospf6Ospf6Interface] {
	return pulumix.Output[[]*Routerospf6Ospf6Interface]{
		OutputState: i.ToRouterospf6Ospf6InterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// Routerospf6Ospf6InterfaceMapInput is an input type that accepts Routerospf6Ospf6InterfaceMap and Routerospf6Ospf6InterfaceMapOutput values.
// You can construct a concrete instance of `Routerospf6Ospf6InterfaceMapInput` via:
//
//	Routerospf6Ospf6InterfaceMap{ "key": Routerospf6Ospf6InterfaceArgs{...} }
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

func (i Routerospf6Ospf6InterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Routerospf6Ospf6Interface] {
	return pulumix.Output[map[string]*Routerospf6Ospf6Interface]{
		OutputState: i.ToRouterospf6Ospf6InterfaceMapOutputWithContext(ctx).OutputState,
	}
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

func (o Routerospf6Ospf6InterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Routerospf6Ospf6Interface] {
	return pulumix.Output[*Routerospf6Ospf6Interface]{
		OutputState: o.OutputState,
	}
}

func (o Routerospf6Ospf6InterfaceOutput) AreaId() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.AreaId }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.Bfd }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) DeadInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.DeadInterval }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) HelloInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.HelloInterval }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) IpsecAuthAlg() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.IpsecAuthAlg }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) IpsecEncAlg() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.IpsecEncAlg }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) IpsecKeys() Routerospf6Ospf6InterfaceIpsecKeyArrayOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) Routerospf6Ospf6InterfaceIpsecKeyArrayOutput { return v.IpsecKeys }).(Routerospf6Ospf6InterfaceIpsecKeyArrayOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) KeyRolloverInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.KeyRolloverInterval }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) MtuIgnore() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.MtuIgnore }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Neighbors() Routerospf6Ospf6InterfaceNeighborArrayOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) Routerospf6Ospf6InterfaceNeighborArrayOutput { return v.Neighbors }).(Routerospf6Ospf6InterfaceNeighborArrayOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) RetransmitInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.RetransmitInterval }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) TransmitDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.IntOutput { return v.TransmitDelay }).(pulumi.IntOutput)
}

func (o Routerospf6Ospf6InterfaceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Routerospf6Ospf6Interface) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
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

func (o Routerospf6Ospf6InterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Routerospf6Ospf6Interface] {
	return pulumix.Output[[]*Routerospf6Ospf6Interface]{
		OutputState: o.OutputState,
	}
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

func (o Routerospf6Ospf6InterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Routerospf6Ospf6Interface] {
	return pulumix.Output[map[string]*Routerospf6Ospf6Interface]{
		OutputState: o.OutputState,
	}
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

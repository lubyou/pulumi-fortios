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

type RouterospfOspfInterface struct {
	pulumi.CustomResourceState

	Authentication      pulumi.StringOutput                      `pulumi:"authentication"`
	AuthenticationKey   pulumi.StringPtrOutput                   `pulumi:"authenticationKey"`
	Bfd                 pulumi.StringOutput                      `pulumi:"bfd"`
	Comments            pulumi.StringPtrOutput                   `pulumi:"comments"`
	Cost                pulumi.IntOutput                         `pulumi:"cost"`
	DatabaseFilterOut   pulumi.StringOutput                      `pulumi:"databaseFilterOut"`
	DeadInterval        pulumi.IntOutput                         `pulumi:"deadInterval"`
	DynamicSortSubtable pulumi.StringPtrOutput                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                   `pulumi:"getAllTables"`
	HelloInterval       pulumi.IntOutput                         `pulumi:"helloInterval"`
	HelloMultiplier     pulumi.IntOutput                         `pulumi:"helloMultiplier"`
	Interface           pulumi.StringOutput                      `pulumi:"interface"`
	Ip                  pulumi.StringOutput                      `pulumi:"ip"`
	Keychain            pulumi.StringOutput                      `pulumi:"keychain"`
	Md5Key              pulumi.StringOutput                      `pulumi:"md5Key"`
	Md5Keychain         pulumi.StringOutput                      `pulumi:"md5Keychain"`
	Md5Keys             RouterospfOspfInterfaceMd5KeyArrayOutput `pulumi:"md5Keys"`
	Mtu                 pulumi.IntOutput                         `pulumi:"mtu"`
	MtuIgnore           pulumi.StringOutput                      `pulumi:"mtuIgnore"`
	Name                pulumi.StringOutput                      `pulumi:"name"`
	NetworkType         pulumi.StringOutput                      `pulumi:"networkType"`
	PrefixLength        pulumi.IntOutput                         `pulumi:"prefixLength"`
	Priority            pulumi.IntOutput                         `pulumi:"priority"`
	ResyncTimeout       pulumi.IntOutput                         `pulumi:"resyncTimeout"`
	RetransmitInterval  pulumi.IntOutput                         `pulumi:"retransmitInterval"`
	Status              pulumi.StringOutput                      `pulumi:"status"`
	TransmitDelay       pulumi.IntOutput                         `pulumi:"transmitDelay"`
	Vdomparam           pulumi.StringPtrOutput                   `pulumi:"vdomparam"`
}

// NewRouterospfOspfInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterospfOspfInterface(ctx *pulumi.Context,
	name string, args *RouterospfOspfInterfaceArgs, opts ...pulumi.ResourceOption) (*RouterospfOspfInterface, error) {
	if args == nil {
		args = &RouterospfOspfInterfaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterospfOspfInterface
	err := ctx.RegisterResource("fortios:index/routerospfOspfInterface:RouterospfOspfInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterospfOspfInterface gets an existing RouterospfOspfInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterospfOspfInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterospfOspfInterfaceState, opts ...pulumi.ResourceOption) (*RouterospfOspfInterface, error) {
	var resource RouterospfOspfInterface
	err := ctx.ReadResource("fortios:index/routerospfOspfInterface:RouterospfOspfInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterospfOspfInterface resources.
type routerospfOspfInterfaceState struct {
	Authentication      *string                         `pulumi:"authentication"`
	AuthenticationKey   *string                         `pulumi:"authenticationKey"`
	Bfd                 *string                         `pulumi:"bfd"`
	Comments            *string                         `pulumi:"comments"`
	Cost                *int                            `pulumi:"cost"`
	DatabaseFilterOut   *string                         `pulumi:"databaseFilterOut"`
	DeadInterval        *int                            `pulumi:"deadInterval"`
	DynamicSortSubtable *string                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                         `pulumi:"getAllTables"`
	HelloInterval       *int                            `pulumi:"helloInterval"`
	HelloMultiplier     *int                            `pulumi:"helloMultiplier"`
	Interface           *string                         `pulumi:"interface"`
	Ip                  *string                         `pulumi:"ip"`
	Keychain            *string                         `pulumi:"keychain"`
	Md5Key              *string                         `pulumi:"md5Key"`
	Md5Keychain         *string                         `pulumi:"md5Keychain"`
	Md5Keys             []RouterospfOspfInterfaceMd5Key `pulumi:"md5Keys"`
	Mtu                 *int                            `pulumi:"mtu"`
	MtuIgnore           *string                         `pulumi:"mtuIgnore"`
	Name                *string                         `pulumi:"name"`
	NetworkType         *string                         `pulumi:"networkType"`
	PrefixLength        *int                            `pulumi:"prefixLength"`
	Priority            *int                            `pulumi:"priority"`
	ResyncTimeout       *int                            `pulumi:"resyncTimeout"`
	RetransmitInterval  *int                            `pulumi:"retransmitInterval"`
	Status              *string                         `pulumi:"status"`
	TransmitDelay       *int                            `pulumi:"transmitDelay"`
	Vdomparam           *string                         `pulumi:"vdomparam"`
}

type RouterospfOspfInterfaceState struct {
	Authentication      pulumi.StringPtrInput
	AuthenticationKey   pulumi.StringPtrInput
	Bfd                 pulumi.StringPtrInput
	Comments            pulumi.StringPtrInput
	Cost                pulumi.IntPtrInput
	DatabaseFilterOut   pulumi.StringPtrInput
	DeadInterval        pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HelloInterval       pulumi.IntPtrInput
	HelloMultiplier     pulumi.IntPtrInput
	Interface           pulumi.StringPtrInput
	Ip                  pulumi.StringPtrInput
	Keychain            pulumi.StringPtrInput
	Md5Key              pulumi.StringPtrInput
	Md5Keychain         pulumi.StringPtrInput
	Md5Keys             RouterospfOspfInterfaceMd5KeyArrayInput
	Mtu                 pulumi.IntPtrInput
	MtuIgnore           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	NetworkType         pulumi.StringPtrInput
	PrefixLength        pulumi.IntPtrInput
	Priority            pulumi.IntPtrInput
	ResyncTimeout       pulumi.IntPtrInput
	RetransmitInterval  pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	TransmitDelay       pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterospfOspfInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospfOspfInterfaceState)(nil)).Elem()
}

type routerospfOspfInterfaceArgs struct {
	Authentication      *string                         `pulumi:"authentication"`
	AuthenticationKey   *string                         `pulumi:"authenticationKey"`
	Bfd                 *string                         `pulumi:"bfd"`
	Comments            *string                         `pulumi:"comments"`
	Cost                *int                            `pulumi:"cost"`
	DatabaseFilterOut   *string                         `pulumi:"databaseFilterOut"`
	DeadInterval        *int                            `pulumi:"deadInterval"`
	DynamicSortSubtable *string                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                         `pulumi:"getAllTables"`
	HelloInterval       *int                            `pulumi:"helloInterval"`
	HelloMultiplier     *int                            `pulumi:"helloMultiplier"`
	Interface           *string                         `pulumi:"interface"`
	Ip                  *string                         `pulumi:"ip"`
	Keychain            *string                         `pulumi:"keychain"`
	Md5Key              *string                         `pulumi:"md5Key"`
	Md5Keychain         *string                         `pulumi:"md5Keychain"`
	Md5Keys             []RouterospfOspfInterfaceMd5Key `pulumi:"md5Keys"`
	Mtu                 *int                            `pulumi:"mtu"`
	MtuIgnore           *string                         `pulumi:"mtuIgnore"`
	Name                *string                         `pulumi:"name"`
	NetworkType         *string                         `pulumi:"networkType"`
	PrefixLength        *int                            `pulumi:"prefixLength"`
	Priority            *int                            `pulumi:"priority"`
	ResyncTimeout       *int                            `pulumi:"resyncTimeout"`
	RetransmitInterval  *int                            `pulumi:"retransmitInterval"`
	Status              *string                         `pulumi:"status"`
	TransmitDelay       *int                            `pulumi:"transmitDelay"`
	Vdomparam           *string                         `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterospfOspfInterface resource.
type RouterospfOspfInterfaceArgs struct {
	Authentication      pulumi.StringPtrInput
	AuthenticationKey   pulumi.StringPtrInput
	Bfd                 pulumi.StringPtrInput
	Comments            pulumi.StringPtrInput
	Cost                pulumi.IntPtrInput
	DatabaseFilterOut   pulumi.StringPtrInput
	DeadInterval        pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HelloInterval       pulumi.IntPtrInput
	HelloMultiplier     pulumi.IntPtrInput
	Interface           pulumi.StringPtrInput
	Ip                  pulumi.StringPtrInput
	Keychain            pulumi.StringPtrInput
	Md5Key              pulumi.StringPtrInput
	Md5Keychain         pulumi.StringPtrInput
	Md5Keys             RouterospfOspfInterfaceMd5KeyArrayInput
	Mtu                 pulumi.IntPtrInput
	MtuIgnore           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	NetworkType         pulumi.StringPtrInput
	PrefixLength        pulumi.IntPtrInput
	Priority            pulumi.IntPtrInput
	ResyncTimeout       pulumi.IntPtrInput
	RetransmitInterval  pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	TransmitDelay       pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterospfOspfInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerospfOspfInterfaceArgs)(nil)).Elem()
}

type RouterospfOspfInterfaceInput interface {
	pulumi.Input

	ToRouterospfOspfInterfaceOutput() RouterospfOspfInterfaceOutput
	ToRouterospfOspfInterfaceOutputWithContext(ctx context.Context) RouterospfOspfInterfaceOutput
}

func (*RouterospfOspfInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterospfOspfInterface)(nil)).Elem()
}

func (i *RouterospfOspfInterface) ToRouterospfOspfInterfaceOutput() RouterospfOspfInterfaceOutput {
	return i.ToRouterospfOspfInterfaceOutputWithContext(context.Background())
}

func (i *RouterospfOspfInterface) ToRouterospfOspfInterfaceOutputWithContext(ctx context.Context) RouterospfOspfInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfOspfInterfaceOutput)
}

func (i *RouterospfOspfInterface) ToOutput(ctx context.Context) pulumix.Output[*RouterospfOspfInterface] {
	return pulumix.Output[*RouterospfOspfInterface]{
		OutputState: i.ToRouterospfOspfInterfaceOutputWithContext(ctx).OutputState,
	}
}

// RouterospfOspfInterfaceArrayInput is an input type that accepts RouterospfOspfInterfaceArray and RouterospfOspfInterfaceArrayOutput values.
// You can construct a concrete instance of `RouterospfOspfInterfaceArrayInput` via:
//
//	RouterospfOspfInterfaceArray{ RouterospfOspfInterfaceArgs{...} }
type RouterospfOspfInterfaceArrayInput interface {
	pulumi.Input

	ToRouterospfOspfInterfaceArrayOutput() RouterospfOspfInterfaceArrayOutput
	ToRouterospfOspfInterfaceArrayOutputWithContext(context.Context) RouterospfOspfInterfaceArrayOutput
}

type RouterospfOspfInterfaceArray []RouterospfOspfInterfaceInput

func (RouterospfOspfInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterospfOspfInterface)(nil)).Elem()
}

func (i RouterospfOspfInterfaceArray) ToRouterospfOspfInterfaceArrayOutput() RouterospfOspfInterfaceArrayOutput {
	return i.ToRouterospfOspfInterfaceArrayOutputWithContext(context.Background())
}

func (i RouterospfOspfInterfaceArray) ToRouterospfOspfInterfaceArrayOutputWithContext(ctx context.Context) RouterospfOspfInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfOspfInterfaceArrayOutput)
}

func (i RouterospfOspfInterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouterospfOspfInterface] {
	return pulumix.Output[[]*RouterospfOspfInterface]{
		OutputState: i.ToRouterospfOspfInterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// RouterospfOspfInterfaceMapInput is an input type that accepts RouterospfOspfInterfaceMap and RouterospfOspfInterfaceMapOutput values.
// You can construct a concrete instance of `RouterospfOspfInterfaceMapInput` via:
//
//	RouterospfOspfInterfaceMap{ "key": RouterospfOspfInterfaceArgs{...} }
type RouterospfOspfInterfaceMapInput interface {
	pulumi.Input

	ToRouterospfOspfInterfaceMapOutput() RouterospfOspfInterfaceMapOutput
	ToRouterospfOspfInterfaceMapOutputWithContext(context.Context) RouterospfOspfInterfaceMapOutput
}

type RouterospfOspfInterfaceMap map[string]RouterospfOspfInterfaceInput

func (RouterospfOspfInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterospfOspfInterface)(nil)).Elem()
}

func (i RouterospfOspfInterfaceMap) ToRouterospfOspfInterfaceMapOutput() RouterospfOspfInterfaceMapOutput {
	return i.ToRouterospfOspfInterfaceMapOutputWithContext(context.Background())
}

func (i RouterospfOspfInterfaceMap) ToRouterospfOspfInterfaceMapOutputWithContext(ctx context.Context) RouterospfOspfInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterospfOspfInterfaceMapOutput)
}

func (i RouterospfOspfInterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterospfOspfInterface] {
	return pulumix.Output[map[string]*RouterospfOspfInterface]{
		OutputState: i.ToRouterospfOspfInterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type RouterospfOspfInterfaceOutput struct{ *pulumi.OutputState }

func (RouterospfOspfInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterospfOspfInterface)(nil)).Elem()
}

func (o RouterospfOspfInterfaceOutput) ToRouterospfOspfInterfaceOutput() RouterospfOspfInterfaceOutput {
	return o
}

func (o RouterospfOspfInterfaceOutput) ToRouterospfOspfInterfaceOutputWithContext(ctx context.Context) RouterospfOspfInterfaceOutput {
	return o
}

func (o RouterospfOspfInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*RouterospfOspfInterface] {
	return pulumix.Output[*RouterospfOspfInterface]{
		OutputState: o.OutputState,
	}
}

func (o RouterospfOspfInterfaceOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) AuthenticationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringPtrOutput { return v.AuthenticationKey }).(pulumi.StringPtrOutput)
}

func (o RouterospfOspfInterfaceOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Bfd }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o RouterospfOspfInterfaceOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) DatabaseFilterOut() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.DatabaseFilterOut }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) DeadInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.DeadInterval }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterospfOspfInterfaceOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterospfOspfInterfaceOutput) HelloInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.HelloInterval }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) HelloMultiplier() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.HelloMultiplier }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Keychain() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Keychain }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Md5Key() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Md5Key }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Md5Keychain() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Md5Keychain }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Md5Keys() RouterospfOspfInterfaceMd5KeyArrayOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) RouterospfOspfInterfaceMd5KeyArrayOutput { return v.Md5Keys }).(RouterospfOspfInterfaceMd5KeyArrayOutput)
}

func (o RouterospfOspfInterfaceOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) MtuIgnore() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.MtuIgnore }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) ResyncTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.ResyncTimeout }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) RetransmitInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.RetransmitInterval }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o RouterospfOspfInterfaceOutput) TransmitDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.IntOutput { return v.TransmitDelay }).(pulumi.IntOutput)
}

func (o RouterospfOspfInterfaceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterospfOspfInterface) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterospfOspfInterfaceArrayOutput struct{ *pulumi.OutputState }

func (RouterospfOspfInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterospfOspfInterface)(nil)).Elem()
}

func (o RouterospfOspfInterfaceArrayOutput) ToRouterospfOspfInterfaceArrayOutput() RouterospfOspfInterfaceArrayOutput {
	return o
}

func (o RouterospfOspfInterfaceArrayOutput) ToRouterospfOspfInterfaceArrayOutputWithContext(ctx context.Context) RouterospfOspfInterfaceArrayOutput {
	return o
}

func (o RouterospfOspfInterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouterospfOspfInterface] {
	return pulumix.Output[[]*RouterospfOspfInterface]{
		OutputState: o.OutputState,
	}
}

func (o RouterospfOspfInterfaceArrayOutput) Index(i pulumi.IntInput) RouterospfOspfInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterospfOspfInterface {
		return vs[0].([]*RouterospfOspfInterface)[vs[1].(int)]
	}).(RouterospfOspfInterfaceOutput)
}

type RouterospfOspfInterfaceMapOutput struct{ *pulumi.OutputState }

func (RouterospfOspfInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterospfOspfInterface)(nil)).Elem()
}

func (o RouterospfOspfInterfaceMapOutput) ToRouterospfOspfInterfaceMapOutput() RouterospfOspfInterfaceMapOutput {
	return o
}

func (o RouterospfOspfInterfaceMapOutput) ToRouterospfOspfInterfaceMapOutputWithContext(ctx context.Context) RouterospfOspfInterfaceMapOutput {
	return o
}

func (o RouterospfOspfInterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterospfOspfInterface] {
	return pulumix.Output[map[string]*RouterospfOspfInterface]{
		OutputState: o.OutputState,
	}
}

func (o RouterospfOspfInterfaceMapOutput) MapIndex(k pulumi.StringInput) RouterospfOspfInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterospfOspfInterface {
		return vs[0].(map[string]*RouterospfOspfInterface)[vs[1].(string)]
	}).(RouterospfOspfInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfOspfInterfaceInput)(nil)).Elem(), &RouterospfOspfInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfOspfInterfaceArrayInput)(nil)).Elem(), RouterospfOspfInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterospfOspfInterfaceMapInput)(nil)).Elem(), RouterospfOspfInterfaceMap{})
	pulumi.RegisterOutputType(RouterospfOspfInterfaceOutput{})
	pulumi.RegisterOutputType(RouterospfOspfInterfaceArrayOutput{})
	pulumi.RegisterOutputType(RouterospfOspfInterfaceMapOutput{})
}

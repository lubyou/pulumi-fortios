// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallAddress struct {
	pulumi.CustomResourceState

	AllowRouting        pulumi.StringOutput                 `pulumi:"allowRouting"`
	AssociatedInterface pulumi.StringPtrOutput              `pulumi:"associatedInterface"`
	CacheTtl            pulumi.IntPtrOutput                 `pulumi:"cacheTtl"`
	ClearpassSpt        pulumi.StringOutput                 `pulumi:"clearpassSpt"`
	Color               pulumi.IntPtrOutput                 `pulumi:"color"`
	Comment             pulumi.StringPtrOutput              `pulumi:"comment"`
	Country             pulumi.StringPtrOutput              `pulumi:"country"`
	DynamicSortSubtable pulumi.StringPtrOutput              `pulumi:"dynamicSortSubtable"`
	EndIp               pulumi.StringOutput                 `pulumi:"endIp"`
	EndMac              pulumi.StringOutput                 `pulumi:"endMac"`
	EpgName             pulumi.StringPtrOutput              `pulumi:"epgName"`
	FabricObject        pulumi.StringOutput                 `pulumi:"fabricObject"`
	Filter              pulumi.StringPtrOutput              `pulumi:"filter"`
	Fqdn                pulumi.StringPtrOutput              `pulumi:"fqdn"`
	FssoGroups          FirewallAddressFssoGroupArrayOutput `pulumi:"fssoGroups"`
	GetAllTables        pulumi.StringPtrOutput              `pulumi:"getAllTables"`
	HwModel             pulumi.StringPtrOutput              `pulumi:"hwModel"`
	HwVendor            pulumi.StringPtrOutput              `pulumi:"hwVendor"`
	Interface           pulumi.StringPtrOutput              `pulumi:"interface"`
	Lists               FirewallAddressListArrayOutput      `pulumi:"lists"`
	Macaddrs            FirewallAddressMacaddrArrayOutput   `pulumi:"macaddrs"`
	Name                pulumi.StringOutput                 `pulumi:"name"`
	NodeIpOnly          pulumi.StringOutput                 `pulumi:"nodeIpOnly"`
	ObjId               pulumi.StringPtrOutput              `pulumi:"objId"`
	ObjTag              pulumi.StringPtrOutput              `pulumi:"objTag"`
	ObjType             pulumi.StringOutput                 `pulumi:"objType"`
	Organization        pulumi.StringPtrOutput              `pulumi:"organization"`
	Os                  pulumi.StringPtrOutput              `pulumi:"os"`
	PolicyGroup         pulumi.StringPtrOutput              `pulumi:"policyGroup"`
	RouteTag            pulumi.IntPtrOutput                 `pulumi:"routeTag"`
	Sdn                 pulumi.StringPtrOutput              `pulumi:"sdn"`
	SdnAddrType         pulumi.StringOutput                 `pulumi:"sdnAddrType"`
	SdnTag              pulumi.StringPtrOutput              `pulumi:"sdnTag"`
	StartIp             pulumi.StringOutput                 `pulumi:"startIp"`
	StartMac            pulumi.StringOutput                 `pulumi:"startMac"`
	SubType             pulumi.StringOutput                 `pulumi:"subType"`
	Subnet              pulumi.StringOutput                 `pulumi:"subnet"`
	SubnetName          pulumi.StringPtrOutput              `pulumi:"subnetName"`
	SwVersion           pulumi.StringPtrOutput              `pulumi:"swVersion"`
	TagDetectionLevel   pulumi.StringPtrOutput              `pulumi:"tagDetectionLevel"`
	TagType             pulumi.StringPtrOutput              `pulumi:"tagType"`
	Taggings            FirewallAddressTaggingArrayOutput   `pulumi:"taggings"`
	Tenant              pulumi.StringPtrOutput              `pulumi:"tenant"`
	Type                pulumi.StringOutput                 `pulumi:"type"`
	Uuid                pulumi.StringOutput                 `pulumi:"uuid"`
	Vdomparam           pulumi.StringPtrOutput              `pulumi:"vdomparam"`
	Visibility          pulumi.StringPtrOutput              `pulumi:"visibility"`
	Wildcard            pulumi.StringOutput                 `pulumi:"wildcard"`
	WildcardFqdn        pulumi.StringPtrOutput              `pulumi:"wildcardFqdn"`
}

// NewFirewallAddress registers a new resource with the given unique name, arguments, and options.
func NewFirewallAddress(ctx *pulumi.Context,
	name string, args *FirewallAddressArgs, opts ...pulumi.ResourceOption) (*FirewallAddress, error) {
	if args == nil {
		args = &FirewallAddressArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallAddress
	err := ctx.RegisterResource("fortios:index/firewallAddress:FirewallAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallAddress gets an existing FirewallAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallAddressState, opts ...pulumi.ResourceOption) (*FirewallAddress, error) {
	var resource FirewallAddress
	err := ctx.ReadResource("fortios:index/firewallAddress:FirewallAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallAddress resources.
type firewallAddressState struct {
	AllowRouting        *string                    `pulumi:"allowRouting"`
	AssociatedInterface *string                    `pulumi:"associatedInterface"`
	CacheTtl            *int                       `pulumi:"cacheTtl"`
	ClearpassSpt        *string                    `pulumi:"clearpassSpt"`
	Color               *int                       `pulumi:"color"`
	Comment             *string                    `pulumi:"comment"`
	Country             *string                    `pulumi:"country"`
	DynamicSortSubtable *string                    `pulumi:"dynamicSortSubtable"`
	EndIp               *string                    `pulumi:"endIp"`
	EndMac              *string                    `pulumi:"endMac"`
	EpgName             *string                    `pulumi:"epgName"`
	FabricObject        *string                    `pulumi:"fabricObject"`
	Filter              *string                    `pulumi:"filter"`
	Fqdn                *string                    `pulumi:"fqdn"`
	FssoGroups          []FirewallAddressFssoGroup `pulumi:"fssoGroups"`
	GetAllTables        *string                    `pulumi:"getAllTables"`
	HwModel             *string                    `pulumi:"hwModel"`
	HwVendor            *string                    `pulumi:"hwVendor"`
	Interface           *string                    `pulumi:"interface"`
	Lists               []FirewallAddressList      `pulumi:"lists"`
	Macaddrs            []FirewallAddressMacaddr   `pulumi:"macaddrs"`
	Name                *string                    `pulumi:"name"`
	NodeIpOnly          *string                    `pulumi:"nodeIpOnly"`
	ObjId               *string                    `pulumi:"objId"`
	ObjTag              *string                    `pulumi:"objTag"`
	ObjType             *string                    `pulumi:"objType"`
	Organization        *string                    `pulumi:"organization"`
	Os                  *string                    `pulumi:"os"`
	PolicyGroup         *string                    `pulumi:"policyGroup"`
	RouteTag            *int                       `pulumi:"routeTag"`
	Sdn                 *string                    `pulumi:"sdn"`
	SdnAddrType         *string                    `pulumi:"sdnAddrType"`
	SdnTag              *string                    `pulumi:"sdnTag"`
	StartIp             *string                    `pulumi:"startIp"`
	StartMac            *string                    `pulumi:"startMac"`
	SubType             *string                    `pulumi:"subType"`
	Subnet              *string                    `pulumi:"subnet"`
	SubnetName          *string                    `pulumi:"subnetName"`
	SwVersion           *string                    `pulumi:"swVersion"`
	TagDetectionLevel   *string                    `pulumi:"tagDetectionLevel"`
	TagType             *string                    `pulumi:"tagType"`
	Taggings            []FirewallAddressTagging   `pulumi:"taggings"`
	Tenant              *string                    `pulumi:"tenant"`
	Type                *string                    `pulumi:"type"`
	Uuid                *string                    `pulumi:"uuid"`
	Vdomparam           *string                    `pulumi:"vdomparam"`
	Visibility          *string                    `pulumi:"visibility"`
	Wildcard            *string                    `pulumi:"wildcard"`
	WildcardFqdn        *string                    `pulumi:"wildcardFqdn"`
}

type FirewallAddressState struct {
	AllowRouting        pulumi.StringPtrInput
	AssociatedInterface pulumi.StringPtrInput
	CacheTtl            pulumi.IntPtrInput
	ClearpassSpt        pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	Country             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	EndIp               pulumi.StringPtrInput
	EndMac              pulumi.StringPtrInput
	EpgName             pulumi.StringPtrInput
	FabricObject        pulumi.StringPtrInput
	Filter              pulumi.StringPtrInput
	Fqdn                pulumi.StringPtrInput
	FssoGroups          FirewallAddressFssoGroupArrayInput
	GetAllTables        pulumi.StringPtrInput
	HwModel             pulumi.StringPtrInput
	HwVendor            pulumi.StringPtrInput
	Interface           pulumi.StringPtrInput
	Lists               FirewallAddressListArrayInput
	Macaddrs            FirewallAddressMacaddrArrayInput
	Name                pulumi.StringPtrInput
	NodeIpOnly          pulumi.StringPtrInput
	ObjId               pulumi.StringPtrInput
	ObjTag              pulumi.StringPtrInput
	ObjType             pulumi.StringPtrInput
	Organization        pulumi.StringPtrInput
	Os                  pulumi.StringPtrInput
	PolicyGroup         pulumi.StringPtrInput
	RouteTag            pulumi.IntPtrInput
	Sdn                 pulumi.StringPtrInput
	SdnAddrType         pulumi.StringPtrInput
	SdnTag              pulumi.StringPtrInput
	StartIp             pulumi.StringPtrInput
	StartMac            pulumi.StringPtrInput
	SubType             pulumi.StringPtrInput
	Subnet              pulumi.StringPtrInput
	SubnetName          pulumi.StringPtrInput
	SwVersion           pulumi.StringPtrInput
	TagDetectionLevel   pulumi.StringPtrInput
	TagType             pulumi.StringPtrInput
	Taggings            FirewallAddressTaggingArrayInput
	Tenant              pulumi.StringPtrInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
	Wildcard            pulumi.StringPtrInput
	WildcardFqdn        pulumi.StringPtrInput
}

func (FirewallAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddressState)(nil)).Elem()
}

type firewallAddressArgs struct {
	AllowRouting        *string                    `pulumi:"allowRouting"`
	AssociatedInterface *string                    `pulumi:"associatedInterface"`
	CacheTtl            *int                       `pulumi:"cacheTtl"`
	ClearpassSpt        *string                    `pulumi:"clearpassSpt"`
	Color               *int                       `pulumi:"color"`
	Comment             *string                    `pulumi:"comment"`
	Country             *string                    `pulumi:"country"`
	DynamicSortSubtable *string                    `pulumi:"dynamicSortSubtable"`
	EndIp               *string                    `pulumi:"endIp"`
	EndMac              *string                    `pulumi:"endMac"`
	EpgName             *string                    `pulumi:"epgName"`
	FabricObject        *string                    `pulumi:"fabricObject"`
	Filter              *string                    `pulumi:"filter"`
	Fqdn                *string                    `pulumi:"fqdn"`
	FssoGroups          []FirewallAddressFssoGroup `pulumi:"fssoGroups"`
	GetAllTables        *string                    `pulumi:"getAllTables"`
	HwModel             *string                    `pulumi:"hwModel"`
	HwVendor            *string                    `pulumi:"hwVendor"`
	Interface           *string                    `pulumi:"interface"`
	Lists               []FirewallAddressList      `pulumi:"lists"`
	Macaddrs            []FirewallAddressMacaddr   `pulumi:"macaddrs"`
	Name                *string                    `pulumi:"name"`
	NodeIpOnly          *string                    `pulumi:"nodeIpOnly"`
	ObjId               *string                    `pulumi:"objId"`
	ObjTag              *string                    `pulumi:"objTag"`
	ObjType             *string                    `pulumi:"objType"`
	Organization        *string                    `pulumi:"organization"`
	Os                  *string                    `pulumi:"os"`
	PolicyGroup         *string                    `pulumi:"policyGroup"`
	RouteTag            *int                       `pulumi:"routeTag"`
	Sdn                 *string                    `pulumi:"sdn"`
	SdnAddrType         *string                    `pulumi:"sdnAddrType"`
	SdnTag              *string                    `pulumi:"sdnTag"`
	StartIp             *string                    `pulumi:"startIp"`
	StartMac            *string                    `pulumi:"startMac"`
	SubType             *string                    `pulumi:"subType"`
	Subnet              *string                    `pulumi:"subnet"`
	SubnetName          *string                    `pulumi:"subnetName"`
	SwVersion           *string                    `pulumi:"swVersion"`
	TagDetectionLevel   *string                    `pulumi:"tagDetectionLevel"`
	TagType             *string                    `pulumi:"tagType"`
	Taggings            []FirewallAddressTagging   `pulumi:"taggings"`
	Tenant              *string                    `pulumi:"tenant"`
	Type                *string                    `pulumi:"type"`
	Uuid                *string                    `pulumi:"uuid"`
	Vdomparam           *string                    `pulumi:"vdomparam"`
	Visibility          *string                    `pulumi:"visibility"`
	Wildcard            *string                    `pulumi:"wildcard"`
	WildcardFqdn        *string                    `pulumi:"wildcardFqdn"`
}

// The set of arguments for constructing a FirewallAddress resource.
type FirewallAddressArgs struct {
	AllowRouting        pulumi.StringPtrInput
	AssociatedInterface pulumi.StringPtrInput
	CacheTtl            pulumi.IntPtrInput
	ClearpassSpt        pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	Country             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	EndIp               pulumi.StringPtrInput
	EndMac              pulumi.StringPtrInput
	EpgName             pulumi.StringPtrInput
	FabricObject        pulumi.StringPtrInput
	Filter              pulumi.StringPtrInput
	Fqdn                pulumi.StringPtrInput
	FssoGroups          FirewallAddressFssoGroupArrayInput
	GetAllTables        pulumi.StringPtrInput
	HwModel             pulumi.StringPtrInput
	HwVendor            pulumi.StringPtrInput
	Interface           pulumi.StringPtrInput
	Lists               FirewallAddressListArrayInput
	Macaddrs            FirewallAddressMacaddrArrayInput
	Name                pulumi.StringPtrInput
	NodeIpOnly          pulumi.StringPtrInput
	ObjId               pulumi.StringPtrInput
	ObjTag              pulumi.StringPtrInput
	ObjType             pulumi.StringPtrInput
	Organization        pulumi.StringPtrInput
	Os                  pulumi.StringPtrInput
	PolicyGroup         pulumi.StringPtrInput
	RouteTag            pulumi.IntPtrInput
	Sdn                 pulumi.StringPtrInput
	SdnAddrType         pulumi.StringPtrInput
	SdnTag              pulumi.StringPtrInput
	StartIp             pulumi.StringPtrInput
	StartMac            pulumi.StringPtrInput
	SubType             pulumi.StringPtrInput
	Subnet              pulumi.StringPtrInput
	SubnetName          pulumi.StringPtrInput
	SwVersion           pulumi.StringPtrInput
	TagDetectionLevel   pulumi.StringPtrInput
	TagType             pulumi.StringPtrInput
	Taggings            FirewallAddressTaggingArrayInput
	Tenant              pulumi.StringPtrInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
	Wildcard            pulumi.StringPtrInput
	WildcardFqdn        pulumi.StringPtrInput
}

func (FirewallAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddressArgs)(nil)).Elem()
}

type FirewallAddressInput interface {
	pulumi.Input

	ToFirewallAddressOutput() FirewallAddressOutput
	ToFirewallAddressOutputWithContext(ctx context.Context) FirewallAddressOutput
}

func (*FirewallAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddress)(nil)).Elem()
}

func (i *FirewallAddress) ToFirewallAddressOutput() FirewallAddressOutput {
	return i.ToFirewallAddressOutputWithContext(context.Background())
}

func (i *FirewallAddress) ToFirewallAddressOutputWithContext(ctx context.Context) FirewallAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddressOutput)
}

// FirewallAddressArrayInput is an input type that accepts FirewallAddressArray and FirewallAddressArrayOutput values.
// You can construct a concrete instance of `FirewallAddressArrayInput` via:
//
//	FirewallAddressArray{ FirewallAddressArgs{...} }
type FirewallAddressArrayInput interface {
	pulumi.Input

	ToFirewallAddressArrayOutput() FirewallAddressArrayOutput
	ToFirewallAddressArrayOutputWithContext(context.Context) FirewallAddressArrayOutput
}

type FirewallAddressArray []FirewallAddressInput

func (FirewallAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddress)(nil)).Elem()
}

func (i FirewallAddressArray) ToFirewallAddressArrayOutput() FirewallAddressArrayOutput {
	return i.ToFirewallAddressArrayOutputWithContext(context.Background())
}

func (i FirewallAddressArray) ToFirewallAddressArrayOutputWithContext(ctx context.Context) FirewallAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddressArrayOutput)
}

// FirewallAddressMapInput is an input type that accepts FirewallAddressMap and FirewallAddressMapOutput values.
// You can construct a concrete instance of `FirewallAddressMapInput` via:
//
//	FirewallAddressMap{ "key": FirewallAddressArgs{...} }
type FirewallAddressMapInput interface {
	pulumi.Input

	ToFirewallAddressMapOutput() FirewallAddressMapOutput
	ToFirewallAddressMapOutputWithContext(context.Context) FirewallAddressMapOutput
}

type FirewallAddressMap map[string]FirewallAddressInput

func (FirewallAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddress)(nil)).Elem()
}

func (i FirewallAddressMap) ToFirewallAddressMapOutput() FirewallAddressMapOutput {
	return i.ToFirewallAddressMapOutputWithContext(context.Background())
}

func (i FirewallAddressMap) ToFirewallAddressMapOutputWithContext(ctx context.Context) FirewallAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddressMapOutput)
}

type FirewallAddressOutput struct{ *pulumi.OutputState }

func (FirewallAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddress)(nil)).Elem()
}

func (o FirewallAddressOutput) ToFirewallAddressOutput() FirewallAddressOutput {
	return o
}

func (o FirewallAddressOutput) ToFirewallAddressOutputWithContext(ctx context.Context) FirewallAddressOutput {
	return o
}

func (o FirewallAddressOutput) AllowRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.AllowRouting }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) AssociatedInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.AssociatedInterface }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) CacheTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.IntPtrOutput { return v.CacheTtl }).(pulumi.IntPtrOutput)
}

func (o FirewallAddressOutput) ClearpassSpt() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.ClearpassSpt }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Color() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.IntPtrOutput { return v.Color }).(pulumi.IntPtrOutput)
}

func (o FirewallAddressOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Country }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.EndIp }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) EndMac() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.EndMac }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) EpgName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.EpgName }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) FssoGroups() FirewallAddressFssoGroupArrayOutput {
	return o.ApplyT(func(v *FirewallAddress) FirewallAddressFssoGroupArrayOutput { return v.FssoGroups }).(FirewallAddressFssoGroupArrayOutput)
}

func (o FirewallAddressOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) HwModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.HwModel }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) HwVendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.HwVendor }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Interface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Interface }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Lists() FirewallAddressListArrayOutput {
	return o.ApplyT(func(v *FirewallAddress) FirewallAddressListArrayOutput { return v.Lists }).(FirewallAddressListArrayOutput)
}

func (o FirewallAddressOutput) Macaddrs() FirewallAddressMacaddrArrayOutput {
	return o.ApplyT(func(v *FirewallAddress) FirewallAddressMacaddrArrayOutput { return v.Macaddrs }).(FirewallAddressMacaddrArrayOutput)
}

func (o FirewallAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) NodeIpOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.NodeIpOnly }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) ObjId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.ObjId }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) ObjTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.ObjTag }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) ObjType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.ObjType }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Os }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) PolicyGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.PolicyGroup }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) RouteTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.IntPtrOutput { return v.RouteTag }).(pulumi.IntPtrOutput)
}

func (o FirewallAddressOutput) Sdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Sdn }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) SdnAddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.SdnAddrType }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) SdnTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.SdnTag }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.StartIp }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) StartMac() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.StartMac }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.SubType }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.SubnetName }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) SwVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.SwVersion }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) TagDetectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.TagDetectionLevel }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) TagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.TagType }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Taggings() FirewallAddressTaggingArrayOutput {
	return o.ApplyT(func(v *FirewallAddress) FirewallAddressTaggingArrayOutput { return v.Taggings }).(FirewallAddressTaggingArrayOutput)
}

func (o FirewallAddressOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Tenant }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.Visibility }).(pulumi.StringPtrOutput)
}

func (o FirewallAddressOutput) Wildcard() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringOutput { return v.Wildcard }).(pulumi.StringOutput)
}

func (o FirewallAddressOutput) WildcardFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddress) pulumi.StringPtrOutput { return v.WildcardFqdn }).(pulumi.StringPtrOutput)
}

type FirewallAddressArrayOutput struct{ *pulumi.OutputState }

func (FirewallAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddress)(nil)).Elem()
}

func (o FirewallAddressArrayOutput) ToFirewallAddressArrayOutput() FirewallAddressArrayOutput {
	return o
}

func (o FirewallAddressArrayOutput) ToFirewallAddressArrayOutputWithContext(ctx context.Context) FirewallAddressArrayOutput {
	return o
}

func (o FirewallAddressArrayOutput) Index(i pulumi.IntInput) FirewallAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallAddress {
		return vs[0].([]*FirewallAddress)[vs[1].(int)]
	}).(FirewallAddressOutput)
}

type FirewallAddressMapOutput struct{ *pulumi.OutputState }

func (FirewallAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddress)(nil)).Elem()
}

func (o FirewallAddressMapOutput) ToFirewallAddressMapOutput() FirewallAddressMapOutput {
	return o
}

func (o FirewallAddressMapOutput) ToFirewallAddressMapOutputWithContext(ctx context.Context) FirewallAddressMapOutput {
	return o
}

func (o FirewallAddressMapOutput) MapIndex(k pulumi.StringInput) FirewallAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallAddress {
		return vs[0].(map[string]*FirewallAddress)[vs[1].(string)]
	}).(FirewallAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddressInput)(nil)).Elem(), &FirewallAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddressArrayInput)(nil)).Elem(), FirewallAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddressMapInput)(nil)).Elem(), FirewallAddressMap{})
	pulumi.RegisterOutputType(FirewallAddressOutput{})
	pulumi.RegisterOutputType(FirewallAddressArrayOutput{})
	pulumi.RegisterOutputType(FirewallAddressMapOutput{})
}

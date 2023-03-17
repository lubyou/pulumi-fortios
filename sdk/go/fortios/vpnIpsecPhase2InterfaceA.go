// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnIpsecPhase2Interface struct {
	pulumi.CustomResourceState

	AddRoute               pulumi.StringOutput    `pulumi:"addRoute"`
	AutoDiscoveryForwarder pulumi.StringOutput    `pulumi:"autoDiscoveryForwarder"`
	AutoDiscoverySender    pulumi.StringOutput    `pulumi:"autoDiscoverySender"`
	AutoNegotiate          pulumi.StringOutput    `pulumi:"autoNegotiate"`
	Comments               pulumi.StringPtrOutput `pulumi:"comments"`
	DhcpIpsec              pulumi.StringOutput    `pulumi:"dhcpIpsec"`
	Dhgrp                  pulumi.StringOutput    `pulumi:"dhgrp"`
	Diffserv               pulumi.StringOutput    `pulumi:"diffserv"`
	Diffservcode           pulumi.StringOutput    `pulumi:"diffservcode"`
	DstAddrType            pulumi.StringOutput    `pulumi:"dstAddrType"`
	DstEndIp               pulumi.StringOutput    `pulumi:"dstEndIp"`
	DstEndIp6              pulumi.StringOutput    `pulumi:"dstEndIp6"`
	DstName                pulumi.StringOutput    `pulumi:"dstName"`
	DstName6               pulumi.StringOutput    `pulumi:"dstName6"`
	DstPort                pulumi.IntOutput       `pulumi:"dstPort"`
	DstStartIp             pulumi.StringOutput    `pulumi:"dstStartIp"`
	DstStartIp6            pulumi.StringOutput    `pulumi:"dstStartIp6"`
	DstSubnet              pulumi.StringOutput    `pulumi:"dstSubnet"`
	DstSubnet6             pulumi.StringOutput    `pulumi:"dstSubnet6"`
	Encapsulation          pulumi.StringOutput    `pulumi:"encapsulation"`
	InboundDscpCopy        pulumi.StringOutput    `pulumi:"inboundDscpCopy"`
	InitiatorTsNarrow      pulumi.StringOutput    `pulumi:"initiatorTsNarrow"`
	Ipv4Df                 pulumi.StringOutput    `pulumi:"ipv4Df"`
	Keepalive              pulumi.StringOutput    `pulumi:"keepalive"`
	KeylifeType            pulumi.StringOutput    `pulumi:"keylifeType"`
	Keylifekbs             pulumi.IntOutput       `pulumi:"keylifekbs"`
	Keylifeseconds         pulumi.IntOutput       `pulumi:"keylifeseconds"`
	L2tp                   pulumi.StringOutput    `pulumi:"l2tp"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	Pfs                    pulumi.StringOutput    `pulumi:"pfs"`
	Phase1name             pulumi.StringOutput    `pulumi:"phase1name"`
	Proposal               pulumi.StringOutput    `pulumi:"proposal"`
	Protocol               pulumi.IntOutput       `pulumi:"protocol"`
	Replay                 pulumi.StringOutput    `pulumi:"replay"`
	RouteOverlap           pulumi.StringOutput    `pulumi:"routeOverlap"`
	SingleSource           pulumi.StringOutput    `pulumi:"singleSource"`
	SrcAddrType            pulumi.StringOutput    `pulumi:"srcAddrType"`
	SrcEndIp               pulumi.StringOutput    `pulumi:"srcEndIp"`
	SrcEndIp6              pulumi.StringOutput    `pulumi:"srcEndIp6"`
	SrcName                pulumi.StringOutput    `pulumi:"srcName"`
	SrcName6               pulumi.StringOutput    `pulumi:"srcName6"`
	SrcPort                pulumi.IntOutput       `pulumi:"srcPort"`
	SrcStartIp             pulumi.StringOutput    `pulumi:"srcStartIp"`
	SrcStartIp6            pulumi.StringOutput    `pulumi:"srcStartIp6"`
	SrcSubnet              pulumi.StringOutput    `pulumi:"srcSubnet"`
	SrcSubnet6             pulumi.StringOutput    `pulumi:"srcSubnet6"`
	Vdomparam              pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnIpsecPhase2Interface registers a new resource with the given unique name, arguments, and options.
func NewVpnIpsecPhase2Interface(ctx *pulumi.Context,
	name string, args *VpnIpsecPhase2InterfaceArgs, opts ...pulumi.ResourceOption) (*VpnIpsecPhase2Interface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Phase1name == nil {
		return nil, errors.New("invalid value for required argument 'Phase1name'")
	}
	if args.Proposal == nil {
		return nil, errors.New("invalid value for required argument 'Proposal'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpnIpsecPhase2Interface
	err := ctx.RegisterResource("fortios:index/vpnIpsecPhase2Interface:VpnIpsecPhase2Interface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnIpsecPhase2Interface gets an existing VpnIpsecPhase2Interface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnIpsecPhase2Interface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnIpsecPhase2InterfaceState, opts ...pulumi.ResourceOption) (*VpnIpsecPhase2Interface, error) {
	var resource VpnIpsecPhase2Interface
	err := ctx.ReadResource("fortios:index/vpnIpsecPhase2Interface:VpnIpsecPhase2Interface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnIpsecPhase2Interface resources.
type vpnIpsecPhase2InterfaceState struct {
	AddRoute               *string `pulumi:"addRoute"`
	AutoDiscoveryForwarder *string `pulumi:"autoDiscoveryForwarder"`
	AutoDiscoverySender    *string `pulumi:"autoDiscoverySender"`
	AutoNegotiate          *string `pulumi:"autoNegotiate"`
	Comments               *string `pulumi:"comments"`
	DhcpIpsec              *string `pulumi:"dhcpIpsec"`
	Dhgrp                  *string `pulumi:"dhgrp"`
	Diffserv               *string `pulumi:"diffserv"`
	Diffservcode           *string `pulumi:"diffservcode"`
	DstAddrType            *string `pulumi:"dstAddrType"`
	DstEndIp               *string `pulumi:"dstEndIp"`
	DstEndIp6              *string `pulumi:"dstEndIp6"`
	DstName                *string `pulumi:"dstName"`
	DstName6               *string `pulumi:"dstName6"`
	DstPort                *int    `pulumi:"dstPort"`
	DstStartIp             *string `pulumi:"dstStartIp"`
	DstStartIp6            *string `pulumi:"dstStartIp6"`
	DstSubnet              *string `pulumi:"dstSubnet"`
	DstSubnet6             *string `pulumi:"dstSubnet6"`
	Encapsulation          *string `pulumi:"encapsulation"`
	InboundDscpCopy        *string `pulumi:"inboundDscpCopy"`
	InitiatorTsNarrow      *string `pulumi:"initiatorTsNarrow"`
	Ipv4Df                 *string `pulumi:"ipv4Df"`
	Keepalive              *string `pulumi:"keepalive"`
	KeylifeType            *string `pulumi:"keylifeType"`
	Keylifekbs             *int    `pulumi:"keylifekbs"`
	Keylifeseconds         *int    `pulumi:"keylifeseconds"`
	L2tp                   *string `pulumi:"l2tp"`
	Name                   *string `pulumi:"name"`
	Pfs                    *string `pulumi:"pfs"`
	Phase1name             *string `pulumi:"phase1name"`
	Proposal               *string `pulumi:"proposal"`
	Protocol               *int    `pulumi:"protocol"`
	Replay                 *string `pulumi:"replay"`
	RouteOverlap           *string `pulumi:"routeOverlap"`
	SingleSource           *string `pulumi:"singleSource"`
	SrcAddrType            *string `pulumi:"srcAddrType"`
	SrcEndIp               *string `pulumi:"srcEndIp"`
	SrcEndIp6              *string `pulumi:"srcEndIp6"`
	SrcName                *string `pulumi:"srcName"`
	SrcName6               *string `pulumi:"srcName6"`
	SrcPort                *int    `pulumi:"srcPort"`
	SrcStartIp             *string `pulumi:"srcStartIp"`
	SrcStartIp6            *string `pulumi:"srcStartIp6"`
	SrcSubnet              *string `pulumi:"srcSubnet"`
	SrcSubnet6             *string `pulumi:"srcSubnet6"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

type VpnIpsecPhase2InterfaceState struct {
	AddRoute               pulumi.StringPtrInput
	AutoDiscoveryForwarder pulumi.StringPtrInput
	AutoDiscoverySender    pulumi.StringPtrInput
	AutoNegotiate          pulumi.StringPtrInput
	Comments               pulumi.StringPtrInput
	DhcpIpsec              pulumi.StringPtrInput
	Dhgrp                  pulumi.StringPtrInput
	Diffserv               pulumi.StringPtrInput
	Diffservcode           pulumi.StringPtrInput
	DstAddrType            pulumi.StringPtrInput
	DstEndIp               pulumi.StringPtrInput
	DstEndIp6              pulumi.StringPtrInput
	DstName                pulumi.StringPtrInput
	DstName6               pulumi.StringPtrInput
	DstPort                pulumi.IntPtrInput
	DstStartIp             pulumi.StringPtrInput
	DstStartIp6            pulumi.StringPtrInput
	DstSubnet              pulumi.StringPtrInput
	DstSubnet6             pulumi.StringPtrInput
	Encapsulation          pulumi.StringPtrInput
	InboundDscpCopy        pulumi.StringPtrInput
	InitiatorTsNarrow      pulumi.StringPtrInput
	Ipv4Df                 pulumi.StringPtrInput
	Keepalive              pulumi.StringPtrInput
	KeylifeType            pulumi.StringPtrInput
	Keylifekbs             pulumi.IntPtrInput
	Keylifeseconds         pulumi.IntPtrInput
	L2tp                   pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Pfs                    pulumi.StringPtrInput
	Phase1name             pulumi.StringPtrInput
	Proposal               pulumi.StringPtrInput
	Protocol               pulumi.IntPtrInput
	Replay                 pulumi.StringPtrInput
	RouteOverlap           pulumi.StringPtrInput
	SingleSource           pulumi.StringPtrInput
	SrcAddrType            pulumi.StringPtrInput
	SrcEndIp               pulumi.StringPtrInput
	SrcEndIp6              pulumi.StringPtrInput
	SrcName                pulumi.StringPtrInput
	SrcName6               pulumi.StringPtrInput
	SrcPort                pulumi.IntPtrInput
	SrcStartIp             pulumi.StringPtrInput
	SrcStartIp6            pulumi.StringPtrInput
	SrcSubnet              pulumi.StringPtrInput
	SrcSubnet6             pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (VpnIpsecPhase2InterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecPhase2InterfaceState)(nil)).Elem()
}

type vpnIpsecPhase2InterfaceArgs struct {
	AddRoute               *string `pulumi:"addRoute"`
	AutoDiscoveryForwarder *string `pulumi:"autoDiscoveryForwarder"`
	AutoDiscoverySender    *string `pulumi:"autoDiscoverySender"`
	AutoNegotiate          *string `pulumi:"autoNegotiate"`
	Comments               *string `pulumi:"comments"`
	DhcpIpsec              *string `pulumi:"dhcpIpsec"`
	Dhgrp                  *string `pulumi:"dhgrp"`
	Diffserv               *string `pulumi:"diffserv"`
	Diffservcode           *string `pulumi:"diffservcode"`
	DstAddrType            *string `pulumi:"dstAddrType"`
	DstEndIp               *string `pulumi:"dstEndIp"`
	DstEndIp6              *string `pulumi:"dstEndIp6"`
	DstName                *string `pulumi:"dstName"`
	DstName6               *string `pulumi:"dstName6"`
	DstPort                *int    `pulumi:"dstPort"`
	DstStartIp             *string `pulumi:"dstStartIp"`
	DstStartIp6            *string `pulumi:"dstStartIp6"`
	DstSubnet              *string `pulumi:"dstSubnet"`
	DstSubnet6             *string `pulumi:"dstSubnet6"`
	Encapsulation          *string `pulumi:"encapsulation"`
	InboundDscpCopy        *string `pulumi:"inboundDscpCopy"`
	InitiatorTsNarrow      *string `pulumi:"initiatorTsNarrow"`
	Ipv4Df                 *string `pulumi:"ipv4Df"`
	Keepalive              *string `pulumi:"keepalive"`
	KeylifeType            *string `pulumi:"keylifeType"`
	Keylifekbs             *int    `pulumi:"keylifekbs"`
	Keylifeseconds         *int    `pulumi:"keylifeseconds"`
	L2tp                   *string `pulumi:"l2tp"`
	Name                   *string `pulumi:"name"`
	Pfs                    *string `pulumi:"pfs"`
	Phase1name             string  `pulumi:"phase1name"`
	Proposal               string  `pulumi:"proposal"`
	Protocol               *int    `pulumi:"protocol"`
	Replay                 *string `pulumi:"replay"`
	RouteOverlap           *string `pulumi:"routeOverlap"`
	SingleSource           *string `pulumi:"singleSource"`
	SrcAddrType            *string `pulumi:"srcAddrType"`
	SrcEndIp               *string `pulumi:"srcEndIp"`
	SrcEndIp6              *string `pulumi:"srcEndIp6"`
	SrcName                *string `pulumi:"srcName"`
	SrcName6               *string `pulumi:"srcName6"`
	SrcPort                *int    `pulumi:"srcPort"`
	SrcStartIp             *string `pulumi:"srcStartIp"`
	SrcStartIp6            *string `pulumi:"srcStartIp6"`
	SrcSubnet              *string `pulumi:"srcSubnet"`
	SrcSubnet6             *string `pulumi:"srcSubnet6"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnIpsecPhase2Interface resource.
type VpnIpsecPhase2InterfaceArgs struct {
	AddRoute               pulumi.StringPtrInput
	AutoDiscoveryForwarder pulumi.StringPtrInput
	AutoDiscoverySender    pulumi.StringPtrInput
	AutoNegotiate          pulumi.StringPtrInput
	Comments               pulumi.StringPtrInput
	DhcpIpsec              pulumi.StringPtrInput
	Dhgrp                  pulumi.StringPtrInput
	Diffserv               pulumi.StringPtrInput
	Diffservcode           pulumi.StringPtrInput
	DstAddrType            pulumi.StringPtrInput
	DstEndIp               pulumi.StringPtrInput
	DstEndIp6              pulumi.StringPtrInput
	DstName                pulumi.StringPtrInput
	DstName6               pulumi.StringPtrInput
	DstPort                pulumi.IntPtrInput
	DstStartIp             pulumi.StringPtrInput
	DstStartIp6            pulumi.StringPtrInput
	DstSubnet              pulumi.StringPtrInput
	DstSubnet6             pulumi.StringPtrInput
	Encapsulation          pulumi.StringPtrInput
	InboundDscpCopy        pulumi.StringPtrInput
	InitiatorTsNarrow      pulumi.StringPtrInput
	Ipv4Df                 pulumi.StringPtrInput
	Keepalive              pulumi.StringPtrInput
	KeylifeType            pulumi.StringPtrInput
	Keylifekbs             pulumi.IntPtrInput
	Keylifeseconds         pulumi.IntPtrInput
	L2tp                   pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Pfs                    pulumi.StringPtrInput
	Phase1name             pulumi.StringInput
	Proposal               pulumi.StringInput
	Protocol               pulumi.IntPtrInput
	Replay                 pulumi.StringPtrInput
	RouteOverlap           pulumi.StringPtrInput
	SingleSource           pulumi.StringPtrInput
	SrcAddrType            pulumi.StringPtrInput
	SrcEndIp               pulumi.StringPtrInput
	SrcEndIp6              pulumi.StringPtrInput
	SrcName                pulumi.StringPtrInput
	SrcName6               pulumi.StringPtrInput
	SrcPort                pulumi.IntPtrInput
	SrcStartIp             pulumi.StringPtrInput
	SrcStartIp6            pulumi.StringPtrInput
	SrcSubnet              pulumi.StringPtrInput
	SrcSubnet6             pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (VpnIpsecPhase2InterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecPhase2InterfaceArgs)(nil)).Elem()
}

type VpnIpsecPhase2InterfaceInput interface {
	pulumi.Input

	ToVpnIpsecPhase2InterfaceOutput() VpnIpsecPhase2InterfaceOutput
	ToVpnIpsecPhase2InterfaceOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceOutput
}

func (*VpnIpsecPhase2Interface) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecPhase2Interface)(nil)).Elem()
}

func (i *VpnIpsecPhase2Interface) ToVpnIpsecPhase2InterfaceOutput() VpnIpsecPhase2InterfaceOutput {
	return i.ToVpnIpsecPhase2InterfaceOutputWithContext(context.Background())
}

func (i *VpnIpsecPhase2Interface) ToVpnIpsecPhase2InterfaceOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecPhase2InterfaceOutput)
}

// VpnIpsecPhase2InterfaceArrayInput is an input type that accepts VpnIpsecPhase2InterfaceArray and VpnIpsecPhase2InterfaceArrayOutput values.
// You can construct a concrete instance of `VpnIpsecPhase2InterfaceArrayInput` via:
//
//	VpnIpsecPhase2InterfaceArray{ VpnIpsecPhase2InterfaceArgs{...} }
type VpnIpsecPhase2InterfaceArrayInput interface {
	pulumi.Input

	ToVpnIpsecPhase2InterfaceArrayOutput() VpnIpsecPhase2InterfaceArrayOutput
	ToVpnIpsecPhase2InterfaceArrayOutputWithContext(context.Context) VpnIpsecPhase2InterfaceArrayOutput
}

type VpnIpsecPhase2InterfaceArray []VpnIpsecPhase2InterfaceInput

func (VpnIpsecPhase2InterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecPhase2Interface)(nil)).Elem()
}

func (i VpnIpsecPhase2InterfaceArray) ToVpnIpsecPhase2InterfaceArrayOutput() VpnIpsecPhase2InterfaceArrayOutput {
	return i.ToVpnIpsecPhase2InterfaceArrayOutputWithContext(context.Background())
}

func (i VpnIpsecPhase2InterfaceArray) ToVpnIpsecPhase2InterfaceArrayOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecPhase2InterfaceArrayOutput)
}

// VpnIpsecPhase2InterfaceMapInput is an input type that accepts VpnIpsecPhase2InterfaceMap and VpnIpsecPhase2InterfaceMapOutput values.
// You can construct a concrete instance of `VpnIpsecPhase2InterfaceMapInput` via:
//
//	VpnIpsecPhase2InterfaceMap{ "key": VpnIpsecPhase2InterfaceArgs{...} }
type VpnIpsecPhase2InterfaceMapInput interface {
	pulumi.Input

	ToVpnIpsecPhase2InterfaceMapOutput() VpnIpsecPhase2InterfaceMapOutput
	ToVpnIpsecPhase2InterfaceMapOutputWithContext(context.Context) VpnIpsecPhase2InterfaceMapOutput
}

type VpnIpsecPhase2InterfaceMap map[string]VpnIpsecPhase2InterfaceInput

func (VpnIpsecPhase2InterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecPhase2Interface)(nil)).Elem()
}

func (i VpnIpsecPhase2InterfaceMap) ToVpnIpsecPhase2InterfaceMapOutput() VpnIpsecPhase2InterfaceMapOutput {
	return i.ToVpnIpsecPhase2InterfaceMapOutputWithContext(context.Background())
}

func (i VpnIpsecPhase2InterfaceMap) ToVpnIpsecPhase2InterfaceMapOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecPhase2InterfaceMapOutput)
}

type VpnIpsecPhase2InterfaceOutput struct{ *pulumi.OutputState }

func (VpnIpsecPhase2InterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecPhase2Interface)(nil)).Elem()
}

func (o VpnIpsecPhase2InterfaceOutput) ToVpnIpsecPhase2InterfaceOutput() VpnIpsecPhase2InterfaceOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceOutput) ToVpnIpsecPhase2InterfaceOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceOutput) AddRoute() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.AddRoute }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) AutoDiscoveryForwarder() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.AutoDiscoveryForwarder }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) AutoDiscoverySender() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.AutoDiscoverySender }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) AutoNegotiate() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.AutoNegotiate }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DhcpIpsec() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DhcpIpsec }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Dhgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Dhgrp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Diffserv() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Diffserv }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Diffservcode }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstAddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstAddrType }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstEndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstEndIp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstEndIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstEndIp6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstName }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstName6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstName6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.IntOutput { return v.DstPort }).(pulumi.IntOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstStartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstStartIp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstStartIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstStartIp6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstSubnet }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) DstSubnet6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.DstSubnet6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Encapsulation() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Encapsulation }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) InboundDscpCopy() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.InboundDscpCopy }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) InitiatorTsNarrow() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.InitiatorTsNarrow }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Ipv4Df() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Ipv4Df }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Keepalive() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Keepalive }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) KeylifeType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.KeylifeType }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Keylifekbs() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.IntOutput { return v.Keylifekbs }).(pulumi.IntOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Keylifeseconds() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.IntOutput { return v.Keylifeseconds }).(pulumi.IntOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) L2tp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.L2tp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Pfs() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Pfs }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Phase1name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Phase1name }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Proposal() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Proposal }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Replay() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.Replay }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) RouteOverlap() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.RouteOverlap }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SingleSource() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SingleSource }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcAddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcAddrType }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcEndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcEndIp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcEndIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcEndIp6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcName }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcName6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcName6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.IntOutput { return v.SrcPort }).(pulumi.IntOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcStartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcStartIp }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcStartIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcStartIp6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcSubnet }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) SrcSubnet6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringOutput { return v.SrcSubnet6 }).(pulumi.StringOutput)
}

func (o VpnIpsecPhase2InterfaceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecPhase2Interface) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnIpsecPhase2InterfaceArrayOutput struct{ *pulumi.OutputState }

func (VpnIpsecPhase2InterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecPhase2Interface)(nil)).Elem()
}

func (o VpnIpsecPhase2InterfaceArrayOutput) ToVpnIpsecPhase2InterfaceArrayOutput() VpnIpsecPhase2InterfaceArrayOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceArrayOutput) ToVpnIpsecPhase2InterfaceArrayOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceArrayOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceArrayOutput) Index(i pulumi.IntInput) VpnIpsecPhase2InterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnIpsecPhase2Interface {
		return vs[0].([]*VpnIpsecPhase2Interface)[vs[1].(int)]
	}).(VpnIpsecPhase2InterfaceOutput)
}

type VpnIpsecPhase2InterfaceMapOutput struct{ *pulumi.OutputState }

func (VpnIpsecPhase2InterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecPhase2Interface)(nil)).Elem()
}

func (o VpnIpsecPhase2InterfaceMapOutput) ToVpnIpsecPhase2InterfaceMapOutput() VpnIpsecPhase2InterfaceMapOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceMapOutput) ToVpnIpsecPhase2InterfaceMapOutputWithContext(ctx context.Context) VpnIpsecPhase2InterfaceMapOutput {
	return o
}

func (o VpnIpsecPhase2InterfaceMapOutput) MapIndex(k pulumi.StringInput) VpnIpsecPhase2InterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnIpsecPhase2Interface {
		return vs[0].(map[string]*VpnIpsecPhase2Interface)[vs[1].(string)]
	}).(VpnIpsecPhase2InterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecPhase2InterfaceInput)(nil)).Elem(), &VpnIpsecPhase2Interface{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecPhase2InterfaceArrayInput)(nil)).Elem(), VpnIpsecPhase2InterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecPhase2InterfaceMapInput)(nil)).Elem(), VpnIpsecPhase2InterfaceMap{})
	pulumi.RegisterOutputType(VpnIpsecPhase2InterfaceOutput{})
	pulumi.RegisterOutputType(VpnIpsecPhase2InterfaceArrayOutput{})
	pulumi.RegisterOutputType(VpnIpsecPhase2InterfaceMapOutput{})
}

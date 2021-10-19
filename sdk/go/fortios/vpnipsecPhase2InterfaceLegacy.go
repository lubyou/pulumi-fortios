// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `VpnIpsecPhase2Interface`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Site To Site/Pre-Shared Key
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
// 		test1, err := fortios.NewVPNIPsecPhase1InterfaceLegacy(ctx, "test1", &fortios.VPNIPsecPhase1InterfaceLegacyArgs{
// 			Authmethod:       pulumi.String("psk"),
// 			AuthmethodRemote: pulumi.String(""),
// 			Comments:         pulumi.String("VPN 001Test P1"),
// 			Interface:        pulumi.String("port2"),
// 			ModeCfg:          pulumi.String("disable"),
// 			Peertype:         pulumi.String("any"),
// 			Proposal:         pulumi.String("aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
// 			Psksecret:        pulumi.String("testscecret112233445566778899"),
// 			RemoteGw:         pulumi.String("1.2.2.2"),
// 			Type:             pulumi.String("static"),
// 			WizardType:       pulumi.String("static-fortigate"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewVPNIPsecPhase2InterfaceLegacy(ctx, "test2", &fortios.VPNIPsecPhase2InterfaceLegacyArgs{
// 			Comments:    pulumi.String("VPN 001Test P2"),
// 			DstAddrType: pulumi.String("name"),
// 			DstName:     pulumi.String("HQ-toBranch_remote"),
// 			Phase1name:  test1.Name,
// 			Proposal:    pulumi.String("aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
// 			SrcAddrType: pulumi.String("name"),
// 			SrcName:     pulumi.String("HQ-toBranch_local"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Site To Site/Signature
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
// 		test1, err := fortios.NewVPNIPsecPhase1InterfaceLegacy(ctx, "test1", &fortios.VPNIPsecPhase1InterfaceLegacyArgs{
// 			Certificates: pulumi.StringArray{
// 				pulumi.String("Fortinet_SSL_ECDSA384"),
// 			},
// 			Comments:   pulumi.String("VPN 001Test P1"),
// 			Interface:  pulumi.String("port2"),
// 			Peer:       pulumi.String("2b_peer"),
// 			Peergrp:    pulumi.String(""),
// 			Peerid:     pulumi.String(""),
// 			Peertype:   pulumi.String("peer"),
// 			Proposal:   pulumi.String("aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
// 			Psksecret:  pulumi.String("testscecret112233445566778899"),
// 			RemoteGw:   pulumi.String("1.2.2.2"),
// 			Type:       pulumi.String("static"),
// 			WizardType: pulumi.String("static-fortigate"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewVPNIPsecPhase2InterfaceLegacy(ctx, "test2", &fortios.VPNIPsecPhase2InterfaceLegacyArgs{
// 			Comments:    pulumi.String("VPN 001Test P2"),
// 			DstAddrType: pulumi.String("subnet"),
// 			DstSubnet:   pulumi.String("2.2.2.2/24"),
// 			Phase1name:  test1.Name,
// 			Proposal:    pulumi.String("aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
// 			SrcAddrType: pulumi.String("range"),
// 			SrcEndIp:    pulumi.String("1.1.1.1"),
// 			SrcStartIp:  pulumi.String("1.1.1.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Remote Access/Pre-Shared Key
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
// 		test1, err := fortios.NewVPNIPsecPhase1InterfaceLegacy(ctx, "test1", &fortios.VPNIPsecPhase1InterfaceLegacyArgs{
// 			Comments:            pulumi.String("VPN 001Test P1"),
// 			Interface:           pulumi.String("port2"),
// 			Ipv4SplitExclude:    pulumi.String(""),
// 			Ipv4SplitInclude:    pulumi.String("d_split"),
// 			Peertype:            pulumi.String("any"),
// 			Proposal:            pulumi.String("aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
// 			Psksecret:           pulumi.String("testscecret112233445566778899"),
// 			RemoteGw:            pulumi.String("0.0.0.0"),
// 			SplitIncludeService: pulumi.String(""),
// 			Type:                pulumi.String("dynamic"),
// 			WizardType:          pulumi.String("dialup-forticlient"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewVPNIPsecPhase2InterfaceLegacy(ctx, "test2", &fortios.VPNIPsecPhase2InterfaceLegacyArgs{
// 			Comments:    pulumi.String("VPN 001Test P2"),
// 			DstAddrType: pulumi.String("subnet"),
// 			DstEndIp:    pulumi.String("0.0.0.0"),
// 			DstStartIp:  pulumi.String("0.0.0.0"),
// 			DstSubnet:   pulumi.String("0.0.0.0 0.0.0.0"),
// 			Phase1name:  test1.Name,
// 			Proposal:    pulumi.String("aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
// 			SrcAddrType: pulumi.String("subnet"),
// 			SrcEndIp:    pulumi.String("0.0.0.0"),
// 			SrcStartIp:  pulumi.String("0.0.0.0"),
// 			SrcSubnet:   pulumi.String("0.0.0.0 0.0.0.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Remote Access/Signature
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
// 		test1, err := fortios.NewVPNIPsecPhase1InterfaceLegacy(ctx, "test1", &fortios.VPNIPsecPhase1InterfaceLegacyArgs{
// 			Certificates: pulumi.StringArray{
// 				pulumi.String("Fortinet_SSL_ECDSA384"),
// 			},
// 			Comments:            pulumi.String("VPN 001Test P1"),
// 			Interface:           pulumi.String("port2"),
// 			Ipv4SplitExclude:    pulumi.String(""),
// 			Ipv4SplitInclude:    pulumi.String("d_split"),
// 			Peer:                pulumi.String("2b_peer"),
// 			Peergrp:             pulumi.String(""),
// 			Peerid:              pulumi.String(""),
// 			Peertype:            pulumi.String("any"),
// 			Proposal:            pulumi.String("aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
// 			Psksecret:           pulumi.String("testscecret112233445566778899"),
// 			RemoteGw:            pulumi.String("1.2.2.2"),
// 			SplitIncludeService: pulumi.String(""),
// 			Type:                pulumi.String("dynamic"),
// 			WizardType:          pulumi.String("dialup-forticlient"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewVPNIPsecPhase2InterfaceLegacy(ctx, "test2", &fortios.VPNIPsecPhase2InterfaceLegacyArgs{
// 			Comments:    pulumi.String("VPN 001Test P2"),
// 			DstAddrType: pulumi.String("subnet"),
// 			DstEndIp:    pulumi.String("0.0.0.0"),
// 			DstStartIp:  pulumi.String("0.0.0.0"),
// 			DstSubnet:   pulumi.String("0.0.0.0 0.0.0.0"),
// 			Phase1name:  test1.Name,
// 			Proposal:    pulumi.String("aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
// 			SrcAddrType: pulumi.String("subnet"),
// 			SrcEndIp:    pulumi.String("0.0.0.0"),
// 			SrcStartIp:  pulumi.String("0.0.0.0"),
// 			SrcSubnet:   pulumi.String("0.0.0.0 0.0.0.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VPNIPsecPhase2InterfaceLegacy struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Local proxy ID type.
	DstAddrType pulumi.StringOutput `pulumi:"dstAddrType"`
	// Remote proxy ID IPv4 end.
	DstEndIp pulumi.StringOutput `pulumi:"dstEndIp"`
	// Remote proxy ID name.
	DstName pulumi.StringOutput `pulumi:"dstName"`
	// Remote proxy ID IPv4 start.
	DstStartIp pulumi.StringOutput `pulumi:"dstStartIp"`
	// Remote proxy ID IPv4 subnet.
	DstSubnet pulumi.StringOutput `pulumi:"dstSubnet"`
	// IPsec tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Phase 1 determines the options required for phase 2.
	Phase1name pulumi.StringOutput `pulumi:"phase1name"`
	// Phase2 proposal.
	Proposal pulumi.StringOutput `pulumi:"proposal"`
	// Local proxy ID type.
	SrcAddrType pulumi.StringOutput `pulumi:"srcAddrType"`
	// Local proxy ID end.
	SrcEndIp pulumi.StringOutput `pulumi:"srcEndIp"`
	// Local proxy ID name.
	SrcName pulumi.StringOutput `pulumi:"srcName"`
	// Local proxy ID start.
	SrcStartIp pulumi.StringOutput `pulumi:"srcStartIp"`
	// Local proxy ID subnet.
	SrcSubnet pulumi.StringOutput `pulumi:"srcSubnet"`
}

// NewVPNIPsecPhase2InterfaceLegacy registers a new resource with the given unique name, arguments, and options.
func NewVPNIPsecPhase2InterfaceLegacy(ctx *pulumi.Context,
	name string, args *VPNIPsecPhase2InterfaceLegacyArgs, opts ...pulumi.ResourceOption) (*VPNIPsecPhase2InterfaceLegacy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Phase1name == nil {
		return nil, errors.New("invalid value for required argument 'Phase1name'")
	}
	var resource VPNIPsecPhase2InterfaceLegacy
	err := ctx.RegisterResource("fortios:index/vPNIPsecPhase2InterfaceLegacy:VPNIPsecPhase2InterfaceLegacy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPNIPsecPhase2InterfaceLegacy gets an existing VPNIPsecPhase2InterfaceLegacy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNIPsecPhase2InterfaceLegacy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPNIPsecPhase2InterfaceLegacyState, opts ...pulumi.ResourceOption) (*VPNIPsecPhase2InterfaceLegacy, error) {
	var resource VPNIPsecPhase2InterfaceLegacy
	err := ctx.ReadResource("fortios:index/vPNIPsecPhase2InterfaceLegacy:VPNIPsecPhase2InterfaceLegacy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPNIPsecPhase2InterfaceLegacy resources.
type vpnipsecPhase2InterfaceLegacyState struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Local proxy ID type.
	DstAddrType *string `pulumi:"dstAddrType"`
	// Remote proxy ID IPv4 end.
	DstEndIp *string `pulumi:"dstEndIp"`
	// Remote proxy ID name.
	DstName *string `pulumi:"dstName"`
	// Remote proxy ID IPv4 start.
	DstStartIp *string `pulumi:"dstStartIp"`
	// Remote proxy ID IPv4 subnet.
	DstSubnet *string `pulumi:"dstSubnet"`
	// IPsec tunnel name.
	Name *string `pulumi:"name"`
	// Phase 1 determines the options required for phase 2.
	Phase1name *string `pulumi:"phase1name"`
	// Phase2 proposal.
	Proposal *string `pulumi:"proposal"`
	// Local proxy ID type.
	SrcAddrType *string `pulumi:"srcAddrType"`
	// Local proxy ID end.
	SrcEndIp *string `pulumi:"srcEndIp"`
	// Local proxy ID name.
	SrcName *string `pulumi:"srcName"`
	// Local proxy ID start.
	SrcStartIp *string `pulumi:"srcStartIp"`
	// Local proxy ID subnet.
	SrcSubnet *string `pulumi:"srcSubnet"`
}

type VPNIPsecPhase2InterfaceLegacyState struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Local proxy ID type.
	DstAddrType pulumi.StringPtrInput
	// Remote proxy ID IPv4 end.
	DstEndIp pulumi.StringPtrInput
	// Remote proxy ID name.
	DstName pulumi.StringPtrInput
	// Remote proxy ID IPv4 start.
	DstStartIp pulumi.StringPtrInput
	// Remote proxy ID IPv4 subnet.
	DstSubnet pulumi.StringPtrInput
	// IPsec tunnel name.
	Name pulumi.StringPtrInput
	// Phase 1 determines the options required for phase 2.
	Phase1name pulumi.StringPtrInput
	// Phase2 proposal.
	Proposal pulumi.StringPtrInput
	// Local proxy ID type.
	SrcAddrType pulumi.StringPtrInput
	// Local proxy ID end.
	SrcEndIp pulumi.StringPtrInput
	// Local proxy ID name.
	SrcName pulumi.StringPtrInput
	// Local proxy ID start.
	SrcStartIp pulumi.StringPtrInput
	// Local proxy ID subnet.
	SrcSubnet pulumi.StringPtrInput
}

func (VPNIPsecPhase2InterfaceLegacyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnipsecPhase2InterfaceLegacyState)(nil)).Elem()
}

type vpnipsecPhase2InterfaceLegacyArgs struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Local proxy ID type.
	DstAddrType *string `pulumi:"dstAddrType"`
	// Remote proxy ID IPv4 end.
	DstEndIp *string `pulumi:"dstEndIp"`
	// Remote proxy ID name.
	DstName *string `pulumi:"dstName"`
	// Remote proxy ID IPv4 start.
	DstStartIp *string `pulumi:"dstStartIp"`
	// Remote proxy ID IPv4 subnet.
	DstSubnet *string `pulumi:"dstSubnet"`
	// IPsec tunnel name.
	Name *string `pulumi:"name"`
	// Phase 1 determines the options required for phase 2.
	Phase1name string `pulumi:"phase1name"`
	// Phase2 proposal.
	Proposal *string `pulumi:"proposal"`
	// Local proxy ID type.
	SrcAddrType *string `pulumi:"srcAddrType"`
	// Local proxy ID end.
	SrcEndIp *string `pulumi:"srcEndIp"`
	// Local proxy ID name.
	SrcName *string `pulumi:"srcName"`
	// Local proxy ID start.
	SrcStartIp *string `pulumi:"srcStartIp"`
	// Local proxy ID subnet.
	SrcSubnet *string `pulumi:"srcSubnet"`
}

// The set of arguments for constructing a VPNIPsecPhase2InterfaceLegacy resource.
type VPNIPsecPhase2InterfaceLegacyArgs struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Local proxy ID type.
	DstAddrType pulumi.StringPtrInput
	// Remote proxy ID IPv4 end.
	DstEndIp pulumi.StringPtrInput
	// Remote proxy ID name.
	DstName pulumi.StringPtrInput
	// Remote proxy ID IPv4 start.
	DstStartIp pulumi.StringPtrInput
	// Remote proxy ID IPv4 subnet.
	DstSubnet pulumi.StringPtrInput
	// IPsec tunnel name.
	Name pulumi.StringPtrInput
	// Phase 1 determines the options required for phase 2.
	Phase1name pulumi.StringInput
	// Phase2 proposal.
	Proposal pulumi.StringPtrInput
	// Local proxy ID type.
	SrcAddrType pulumi.StringPtrInput
	// Local proxy ID end.
	SrcEndIp pulumi.StringPtrInput
	// Local proxy ID name.
	SrcName pulumi.StringPtrInput
	// Local proxy ID start.
	SrcStartIp pulumi.StringPtrInput
	// Local proxy ID subnet.
	SrcSubnet pulumi.StringPtrInput
}

func (VPNIPsecPhase2InterfaceLegacyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnipsecPhase2InterfaceLegacyArgs)(nil)).Elem()
}

type VPNIPsecPhase2InterfaceLegacyInput interface {
	pulumi.Input

	ToVPNIPsecPhase2InterfaceLegacyOutput() VPNIPsecPhase2InterfaceLegacyOutput
	ToVPNIPsecPhase2InterfaceLegacyOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyOutput
}

func (*VPNIPsecPhase2InterfaceLegacy) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (i *VPNIPsecPhase2InterfaceLegacy) ToVPNIPsecPhase2InterfaceLegacyOutput() VPNIPsecPhase2InterfaceLegacyOutput {
	return i.ToVPNIPsecPhase2InterfaceLegacyOutputWithContext(context.Background())
}

func (i *VPNIPsecPhase2InterfaceLegacy) ToVPNIPsecPhase2InterfaceLegacyOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNIPsecPhase2InterfaceLegacyOutput)
}

func (i *VPNIPsecPhase2InterfaceLegacy) ToVPNIPsecPhase2InterfaceLegacyPtrOutput() VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return i.ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(context.Background())
}

func (i *VPNIPsecPhase2InterfaceLegacy) ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNIPsecPhase2InterfaceLegacyPtrOutput)
}

type VPNIPsecPhase2InterfaceLegacyPtrInput interface {
	pulumi.Input

	ToVPNIPsecPhase2InterfaceLegacyPtrOutput() VPNIPsecPhase2InterfaceLegacyPtrOutput
	ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyPtrOutput
}

type vpnipsecPhase2InterfaceLegacyPtrType VPNIPsecPhase2InterfaceLegacyArgs

func (*vpnipsecPhase2InterfaceLegacyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (i *vpnipsecPhase2InterfaceLegacyPtrType) ToVPNIPsecPhase2InterfaceLegacyPtrOutput() VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return i.ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(context.Background())
}

func (i *vpnipsecPhase2InterfaceLegacyPtrType) ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNIPsecPhase2InterfaceLegacyPtrOutput)
}

// VPNIPsecPhase2InterfaceLegacyArrayInput is an input type that accepts VPNIPsecPhase2InterfaceLegacyArray and VPNIPsecPhase2InterfaceLegacyArrayOutput values.
// You can construct a concrete instance of `VPNIPsecPhase2InterfaceLegacyArrayInput` via:
//
//          VPNIPsecPhase2InterfaceLegacyArray{ VPNIPsecPhase2InterfaceLegacyArgs{...} }
type VPNIPsecPhase2InterfaceLegacyArrayInput interface {
	pulumi.Input

	ToVPNIPsecPhase2InterfaceLegacyArrayOutput() VPNIPsecPhase2InterfaceLegacyArrayOutput
	ToVPNIPsecPhase2InterfaceLegacyArrayOutputWithContext(context.Context) VPNIPsecPhase2InterfaceLegacyArrayOutput
}

type VPNIPsecPhase2InterfaceLegacyArray []VPNIPsecPhase2InterfaceLegacyInput

func (VPNIPsecPhase2InterfaceLegacyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (i VPNIPsecPhase2InterfaceLegacyArray) ToVPNIPsecPhase2InterfaceLegacyArrayOutput() VPNIPsecPhase2InterfaceLegacyArrayOutput {
	return i.ToVPNIPsecPhase2InterfaceLegacyArrayOutputWithContext(context.Background())
}

func (i VPNIPsecPhase2InterfaceLegacyArray) ToVPNIPsecPhase2InterfaceLegacyArrayOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNIPsecPhase2InterfaceLegacyArrayOutput)
}

// VPNIPsecPhase2InterfaceLegacyMapInput is an input type that accepts VPNIPsecPhase2InterfaceLegacyMap and VPNIPsecPhase2InterfaceLegacyMapOutput values.
// You can construct a concrete instance of `VPNIPsecPhase2InterfaceLegacyMapInput` via:
//
//          VPNIPsecPhase2InterfaceLegacyMap{ "key": VPNIPsecPhase2InterfaceLegacyArgs{...} }
type VPNIPsecPhase2InterfaceLegacyMapInput interface {
	pulumi.Input

	ToVPNIPsecPhase2InterfaceLegacyMapOutput() VPNIPsecPhase2InterfaceLegacyMapOutput
	ToVPNIPsecPhase2InterfaceLegacyMapOutputWithContext(context.Context) VPNIPsecPhase2InterfaceLegacyMapOutput
}

type VPNIPsecPhase2InterfaceLegacyMap map[string]VPNIPsecPhase2InterfaceLegacyInput

func (VPNIPsecPhase2InterfaceLegacyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (i VPNIPsecPhase2InterfaceLegacyMap) ToVPNIPsecPhase2InterfaceLegacyMapOutput() VPNIPsecPhase2InterfaceLegacyMapOutput {
	return i.ToVPNIPsecPhase2InterfaceLegacyMapOutputWithContext(context.Background())
}

func (i VPNIPsecPhase2InterfaceLegacyMap) ToVPNIPsecPhase2InterfaceLegacyMapOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNIPsecPhase2InterfaceLegacyMapOutput)
}

type VPNIPsecPhase2InterfaceLegacyOutput struct {
	*pulumi.OutputState
}

func (VPNIPsecPhase2InterfaceLegacyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (o VPNIPsecPhase2InterfaceLegacyOutput) ToVPNIPsecPhase2InterfaceLegacyOutput() VPNIPsecPhase2InterfaceLegacyOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyOutput) ToVPNIPsecPhase2InterfaceLegacyOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyOutput) ToVPNIPsecPhase2InterfaceLegacyPtrOutput() VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return o.ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(context.Background())
}

func (o VPNIPsecPhase2InterfaceLegacyOutput) ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return o.ApplyT(func(v VPNIPsecPhase2InterfaceLegacy) *VPNIPsecPhase2InterfaceLegacy {
		return &v
	}).(VPNIPsecPhase2InterfaceLegacyPtrOutput)
}

type VPNIPsecPhase2InterfaceLegacyPtrOutput struct {
	*pulumi.OutputState
}

func (VPNIPsecPhase2InterfaceLegacyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (o VPNIPsecPhase2InterfaceLegacyPtrOutput) ToVPNIPsecPhase2InterfaceLegacyPtrOutput() VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyPtrOutput) ToVPNIPsecPhase2InterfaceLegacyPtrOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyPtrOutput {
	return o
}

type VPNIPsecPhase2InterfaceLegacyArrayOutput struct{ *pulumi.OutputState }

func (VPNIPsecPhase2InterfaceLegacyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (o VPNIPsecPhase2InterfaceLegacyArrayOutput) ToVPNIPsecPhase2InterfaceLegacyArrayOutput() VPNIPsecPhase2InterfaceLegacyArrayOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyArrayOutput) ToVPNIPsecPhase2InterfaceLegacyArrayOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyArrayOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyArrayOutput) Index(i pulumi.IntInput) VPNIPsecPhase2InterfaceLegacyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VPNIPsecPhase2InterfaceLegacy {
		return vs[0].([]VPNIPsecPhase2InterfaceLegacy)[vs[1].(int)]
	}).(VPNIPsecPhase2InterfaceLegacyOutput)
}

type VPNIPsecPhase2InterfaceLegacyMapOutput struct{ *pulumi.OutputState }

func (VPNIPsecPhase2InterfaceLegacyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VPNIPsecPhase2InterfaceLegacy)(nil))
}

func (o VPNIPsecPhase2InterfaceLegacyMapOutput) ToVPNIPsecPhase2InterfaceLegacyMapOutput() VPNIPsecPhase2InterfaceLegacyMapOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyMapOutput) ToVPNIPsecPhase2InterfaceLegacyMapOutputWithContext(ctx context.Context) VPNIPsecPhase2InterfaceLegacyMapOutput {
	return o
}

func (o VPNIPsecPhase2InterfaceLegacyMapOutput) MapIndex(k pulumi.StringInput) VPNIPsecPhase2InterfaceLegacyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VPNIPsecPhase2InterfaceLegacy {
		return vs[0].(map[string]VPNIPsecPhase2InterfaceLegacy)[vs[1].(string)]
	}).(VPNIPsecPhase2InterfaceLegacyOutput)
}

func init() {
	pulumi.RegisterOutputType(VPNIPsecPhase2InterfaceLegacyOutput{})
	pulumi.RegisterOutputType(VPNIPsecPhase2InterfaceLegacyPtrOutput{})
	pulumi.RegisterOutputType(VPNIPsecPhase2InterfaceLegacyArrayOutput{})
	pulumi.RegisterOutputType(VPNIPsecPhase2InterfaceLegacyMapOutput{})
}

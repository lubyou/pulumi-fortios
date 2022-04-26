// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WAN optimization profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewWanoptProfile(ctx, "trname", &fortios.WanoptProfileArgs{
// 			Cifs: &WanoptProfileCifsArgs{
// 				ByteCaching:    pulumi.String("enable"),
// 				LogTraffic:     pulumi.String("enable"),
// 				Port:           pulumi.Int(445),
// 				PreferChunking: pulumi.String("fix"),
// 				SecureTunnel:   pulumi.String("disable"),
// 				Status:         pulumi.String("disable"),
// 				TunnelSharing:  pulumi.String("private"),
// 			},
// 			Comments: pulumi.String("test"),
// 			Ftp: &WanoptProfileFtpArgs{
// 				ByteCaching:    pulumi.String("enable"),
// 				LogTraffic:     pulumi.String("enable"),
// 				Port:           pulumi.Int(21),
// 				PreferChunking: pulumi.String("fix"),
// 				SecureTunnel:   pulumi.String("disable"),
// 				Status:         pulumi.String("disable"),
// 				TunnelSharing:  pulumi.String("private"),
// 			},
// 			Http: &WanoptProfileHttpArgs{
// 				ByteCaching:        pulumi.String("enable"),
// 				LogTraffic:         pulumi.String("enable"),
// 				Port:               pulumi.Int(80),
// 				PreferChunking:     pulumi.String("fix"),
// 				SecureTunnel:       pulumi.String("disable"),
// 				Ssl:                pulumi.String("disable"),
// 				SslPort:            pulumi.Int(443),
// 				Status:             pulumi.String("disable"),
// 				TunnelNonHttp:      pulumi.String("disable"),
// 				TunnelSharing:      pulumi.String("private"),
// 				UnknownHttpVersion: pulumi.String("tunnel"),
// 			},
// 			Mapi: &WanoptProfileMapiArgs{
// 				ByteCaching:   pulumi.String("enable"),
// 				LogTraffic:    pulumi.String("enable"),
// 				Port:          pulumi.Int(135),
// 				SecureTunnel:  pulumi.String("disable"),
// 				Status:        pulumi.String("disable"),
// 				TunnelSharing: pulumi.String("private"),
// 			},
// 			Tcp: &WanoptProfileTcpArgs{
// 				ByteCaching:    pulumi.String("disable"),
// 				ByteCachingOpt: pulumi.String("mem-only"),
// 				LogTraffic:     pulumi.String("enable"),
// 				Port:           pulumi.String("1-65535"),
// 				SecureTunnel:   pulumi.String("disable"),
// 				Ssl:            pulumi.String("disable"),
// 				SslPort:        pulumi.Int(443),
// 				Status:         pulumi.String("disable"),
// 				TunnelSharing:  pulumi.String("private"),
// 			},
// 			Transparent: pulumi.String("enable"),
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
// Wanopt Profile can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WanoptProfile struct {
	pulumi.CustomResourceState

	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringOutput `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs WanoptProfileCifsPtrOutput `pulumi:"cifs"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp WanoptProfileFtpPtrOutput `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http WanoptProfileHttpPtrOutput `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi WanoptProfileMapiPtrOutput `pulumi:"mapi"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp WanoptProfileTcpPtrOutput `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringOutput `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptProfile registers a new resource with the given unique name, arguments, and options.
func NewWanoptProfile(ctx *pulumi.Context,
	name string, args *WanoptProfileArgs, opts ...pulumi.ResourceOption) (*WanoptProfile, error) {
	if args == nil {
		args = &WanoptProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WanoptProfile
	err := ctx.RegisterResource("fortios:index/wanoptProfile:WanoptProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptProfile gets an existing WanoptProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptProfileState, opts ...pulumi.ResourceOption) (*WanoptProfile, error) {
	var resource WanoptProfile
	err := ctx.ReadResource("fortios:index/wanoptProfile:WanoptProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptProfile resources.
type wanoptProfileState struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup *string `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs *WanoptProfileCifs `pulumi:"cifs"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp *WanoptProfileFtp `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http *WanoptProfileHttp `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi *WanoptProfileMapi `pulumi:"mapi"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp *WanoptProfileTcp `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent *string `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptProfileState struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringPtrInput
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs WanoptProfileCifsPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp WanoptProfileFtpPtrInput
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http WanoptProfileHttpPtrInput
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi WanoptProfileMapiPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp WanoptProfileTcpPtrInput
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptProfileState)(nil)).Elem()
}

type wanoptProfileArgs struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup *string `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs *WanoptProfileCifs `pulumi:"cifs"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp *WanoptProfileFtp `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http *WanoptProfileHttp `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi *WanoptProfileMapi `pulumi:"mapi"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp *WanoptProfileTcp `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent *string `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptProfile resource.
type WanoptProfileArgs struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringPtrInput
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs WanoptProfileCifsPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp WanoptProfileFtpPtrInput
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http WanoptProfileHttpPtrInput
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi WanoptProfileMapiPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp WanoptProfileTcpPtrInput
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptProfileArgs)(nil)).Elem()
}

type WanoptProfileInput interface {
	pulumi.Input

	ToWanoptProfileOutput() WanoptProfileOutput
	ToWanoptProfileOutputWithContext(ctx context.Context) WanoptProfileOutput
}

func (*WanoptProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptProfile)(nil)).Elem()
}

func (i *WanoptProfile) ToWanoptProfileOutput() WanoptProfileOutput {
	return i.ToWanoptProfileOutputWithContext(context.Background())
}

func (i *WanoptProfile) ToWanoptProfileOutputWithContext(ctx context.Context) WanoptProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptProfileOutput)
}

// WanoptProfileArrayInput is an input type that accepts WanoptProfileArray and WanoptProfileArrayOutput values.
// You can construct a concrete instance of `WanoptProfileArrayInput` via:
//
//          WanoptProfileArray{ WanoptProfileArgs{...} }
type WanoptProfileArrayInput interface {
	pulumi.Input

	ToWanoptProfileArrayOutput() WanoptProfileArrayOutput
	ToWanoptProfileArrayOutputWithContext(context.Context) WanoptProfileArrayOutput
}

type WanoptProfileArray []WanoptProfileInput

func (WanoptProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptProfile)(nil)).Elem()
}

func (i WanoptProfileArray) ToWanoptProfileArrayOutput() WanoptProfileArrayOutput {
	return i.ToWanoptProfileArrayOutputWithContext(context.Background())
}

func (i WanoptProfileArray) ToWanoptProfileArrayOutputWithContext(ctx context.Context) WanoptProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptProfileArrayOutput)
}

// WanoptProfileMapInput is an input type that accepts WanoptProfileMap and WanoptProfileMapOutput values.
// You can construct a concrete instance of `WanoptProfileMapInput` via:
//
//          WanoptProfileMap{ "key": WanoptProfileArgs{...} }
type WanoptProfileMapInput interface {
	pulumi.Input

	ToWanoptProfileMapOutput() WanoptProfileMapOutput
	ToWanoptProfileMapOutputWithContext(context.Context) WanoptProfileMapOutput
}

type WanoptProfileMap map[string]WanoptProfileInput

func (WanoptProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptProfile)(nil)).Elem()
}

func (i WanoptProfileMap) ToWanoptProfileMapOutput() WanoptProfileMapOutput {
	return i.ToWanoptProfileMapOutputWithContext(context.Background())
}

func (i WanoptProfileMap) ToWanoptProfileMapOutputWithContext(ctx context.Context) WanoptProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptProfileMapOutput)
}

type WanoptProfileOutput struct{ *pulumi.OutputState }

func (WanoptProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptProfile)(nil)).Elem()
}

func (o WanoptProfileOutput) ToWanoptProfileOutput() WanoptProfileOutput {
	return o
}

func (o WanoptProfileOutput) ToWanoptProfileOutputWithContext(ctx context.Context) WanoptProfileOutput {
	return o
}

type WanoptProfileArrayOutput struct{ *pulumi.OutputState }

func (WanoptProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptProfile)(nil)).Elem()
}

func (o WanoptProfileArrayOutput) ToWanoptProfileArrayOutput() WanoptProfileArrayOutput {
	return o
}

func (o WanoptProfileArrayOutput) ToWanoptProfileArrayOutputWithContext(ctx context.Context) WanoptProfileArrayOutput {
	return o
}

func (o WanoptProfileArrayOutput) Index(i pulumi.IntInput) WanoptProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WanoptProfile {
		return vs[0].([]*WanoptProfile)[vs[1].(int)]
	}).(WanoptProfileOutput)
}

type WanoptProfileMapOutput struct{ *pulumi.OutputState }

func (WanoptProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptProfile)(nil)).Elem()
}

func (o WanoptProfileMapOutput) ToWanoptProfileMapOutput() WanoptProfileMapOutput {
	return o
}

func (o WanoptProfileMapOutput) ToWanoptProfileMapOutputWithContext(ctx context.Context) WanoptProfileMapOutput {
	return o
}

func (o WanoptProfileMapOutput) MapIndex(k pulumi.StringInput) WanoptProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WanoptProfile {
		return vs[0].(map[string]*WanoptProfile)[vs[1].(string)]
	}).(WanoptProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptProfileInput)(nil)).Elem(), &WanoptProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptProfileArrayInput)(nil)).Elem(), WanoptProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptProfileMapInput)(nil)).Elem(), WanoptProfileMap{})
	pulumi.RegisterOutputType(WanoptProfileOutput{})
	pulumi.RegisterOutputType(WanoptProfileArrayOutput{})
	pulumi.RegisterOutputType(WanoptProfileMapOutput{})
}

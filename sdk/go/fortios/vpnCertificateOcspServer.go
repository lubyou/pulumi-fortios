// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OCSP server configuration.
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
// 		_, err := fortios.NewVpnCertificateOcspServer(ctx, "trname", &fortios.VpnCertificateOcspServerArgs{
// 			Cert:          pulumi.String("ACCVRAIZ1"),
// 			SourceIp:      pulumi.String("0.0.0.0"),
// 			UnavailAction: pulumi.String("revoke"),
// 			Url:           pulumi.String("www.tetserv.com"),
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
// VpnCertificate OcspServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/vpnCertificateOcspServer:VpnCertificateOcspServer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type VpnCertificateOcspServer struct {
	pulumi.CustomResourceState

	// OCSP server certificate.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// OCSP server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringOutput `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringOutput `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringOutput `pulumi:"unavailAction"`
	// OCSP server URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnCertificateOcspServer registers a new resource with the given unique name, arguments, and options.
func NewVpnCertificateOcspServer(ctx *pulumi.Context,
	name string, args *VpnCertificateOcspServerArgs, opts ...pulumi.ResourceOption) (*VpnCertificateOcspServer, error) {
	if args == nil {
		args = &VpnCertificateOcspServerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnCertificateOcspServer
	err := ctx.RegisterResource("fortios:index/vpnCertificateOcspServer:VpnCertificateOcspServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCertificateOcspServer gets an existing VpnCertificateOcspServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCertificateOcspServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCertificateOcspServerState, opts ...pulumi.ResourceOption) (*VpnCertificateOcspServer, error) {
	var resource VpnCertificateOcspServer
	err := ctx.ReadResource("fortios:index/vpnCertificateOcspServer:VpnCertificateOcspServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCertificateOcspServer resources.
type vpnCertificateOcspServerState struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpnCertificateOcspServerState struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnCertificateOcspServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateOcspServerState)(nil)).Elem()
}

type vpnCertificateOcspServerArgs struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnCertificateOcspServer resource.
type VpnCertificateOcspServerArgs struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnCertificateOcspServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateOcspServerArgs)(nil)).Elem()
}

type VpnCertificateOcspServerInput interface {
	pulumi.Input

	ToVpnCertificateOcspServerOutput() VpnCertificateOcspServerOutput
	ToVpnCertificateOcspServerOutputWithContext(ctx context.Context) VpnCertificateOcspServerOutput
}

func (*VpnCertificateOcspServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateOcspServer)(nil)).Elem()
}

func (i *VpnCertificateOcspServer) ToVpnCertificateOcspServerOutput() VpnCertificateOcspServerOutput {
	return i.ToVpnCertificateOcspServerOutputWithContext(context.Background())
}

func (i *VpnCertificateOcspServer) ToVpnCertificateOcspServerOutputWithContext(ctx context.Context) VpnCertificateOcspServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateOcspServerOutput)
}

// VpnCertificateOcspServerArrayInput is an input type that accepts VpnCertificateOcspServerArray and VpnCertificateOcspServerArrayOutput values.
// You can construct a concrete instance of `VpnCertificateOcspServerArrayInput` via:
//
//          VpnCertificateOcspServerArray{ VpnCertificateOcspServerArgs{...} }
type VpnCertificateOcspServerArrayInput interface {
	pulumi.Input

	ToVpnCertificateOcspServerArrayOutput() VpnCertificateOcspServerArrayOutput
	ToVpnCertificateOcspServerArrayOutputWithContext(context.Context) VpnCertificateOcspServerArrayOutput
}

type VpnCertificateOcspServerArray []VpnCertificateOcspServerInput

func (VpnCertificateOcspServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateOcspServer)(nil)).Elem()
}

func (i VpnCertificateOcspServerArray) ToVpnCertificateOcspServerArrayOutput() VpnCertificateOcspServerArrayOutput {
	return i.ToVpnCertificateOcspServerArrayOutputWithContext(context.Background())
}

func (i VpnCertificateOcspServerArray) ToVpnCertificateOcspServerArrayOutputWithContext(ctx context.Context) VpnCertificateOcspServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateOcspServerArrayOutput)
}

// VpnCertificateOcspServerMapInput is an input type that accepts VpnCertificateOcspServerMap and VpnCertificateOcspServerMapOutput values.
// You can construct a concrete instance of `VpnCertificateOcspServerMapInput` via:
//
//          VpnCertificateOcspServerMap{ "key": VpnCertificateOcspServerArgs{...} }
type VpnCertificateOcspServerMapInput interface {
	pulumi.Input

	ToVpnCertificateOcspServerMapOutput() VpnCertificateOcspServerMapOutput
	ToVpnCertificateOcspServerMapOutputWithContext(context.Context) VpnCertificateOcspServerMapOutput
}

type VpnCertificateOcspServerMap map[string]VpnCertificateOcspServerInput

func (VpnCertificateOcspServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateOcspServer)(nil)).Elem()
}

func (i VpnCertificateOcspServerMap) ToVpnCertificateOcspServerMapOutput() VpnCertificateOcspServerMapOutput {
	return i.ToVpnCertificateOcspServerMapOutputWithContext(context.Background())
}

func (i VpnCertificateOcspServerMap) ToVpnCertificateOcspServerMapOutputWithContext(ctx context.Context) VpnCertificateOcspServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateOcspServerMapOutput)
}

type VpnCertificateOcspServerOutput struct{ *pulumi.OutputState }

func (VpnCertificateOcspServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateOcspServer)(nil)).Elem()
}

func (o VpnCertificateOcspServerOutput) ToVpnCertificateOcspServerOutput() VpnCertificateOcspServerOutput {
	return o
}

func (o VpnCertificateOcspServerOutput) ToVpnCertificateOcspServerOutputWithContext(ctx context.Context) VpnCertificateOcspServerOutput {
	return o
}

type VpnCertificateOcspServerArrayOutput struct{ *pulumi.OutputState }

func (VpnCertificateOcspServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateOcspServer)(nil)).Elem()
}

func (o VpnCertificateOcspServerArrayOutput) ToVpnCertificateOcspServerArrayOutput() VpnCertificateOcspServerArrayOutput {
	return o
}

func (o VpnCertificateOcspServerArrayOutput) ToVpnCertificateOcspServerArrayOutputWithContext(ctx context.Context) VpnCertificateOcspServerArrayOutput {
	return o
}

func (o VpnCertificateOcspServerArrayOutput) Index(i pulumi.IntInput) VpnCertificateOcspServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnCertificateOcspServer {
		return vs[0].([]*VpnCertificateOcspServer)[vs[1].(int)]
	}).(VpnCertificateOcspServerOutput)
}

type VpnCertificateOcspServerMapOutput struct{ *pulumi.OutputState }

func (VpnCertificateOcspServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateOcspServer)(nil)).Elem()
}

func (o VpnCertificateOcspServerMapOutput) ToVpnCertificateOcspServerMapOutput() VpnCertificateOcspServerMapOutput {
	return o
}

func (o VpnCertificateOcspServerMapOutput) ToVpnCertificateOcspServerMapOutputWithContext(ctx context.Context) VpnCertificateOcspServerMapOutput {
	return o
}

func (o VpnCertificateOcspServerMapOutput) MapIndex(k pulumi.StringInput) VpnCertificateOcspServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnCertificateOcspServer {
		return vs[0].(map[string]*VpnCertificateOcspServer)[vs[1].(string)]
	}).(VpnCertificateOcspServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateOcspServerInput)(nil)).Elem(), &VpnCertificateOcspServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateOcspServerArrayInput)(nil)).Elem(), VpnCertificateOcspServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateOcspServerMapInput)(nil)).Elem(), VpnCertificateOcspServerMap{})
	pulumi.RegisterOutputType(VpnCertificateOcspServerOutput{})
	pulumi.RegisterOutputType(VpnCertificateOcspServerArrayOutput{})
	pulumi.RegisterOutputType(VpnCertificateOcspServerMapOutput{})
}

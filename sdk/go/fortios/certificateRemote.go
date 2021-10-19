// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Remote certificate as a PEM file.
//
// ## Import
//
// Certificate Remote can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/certificateRemote:CertificateRemote labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type CertificateRemote struct {
	pulumi.CustomResourceState

	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// Remote certificate.
	Remote pulumi.StringOutput `pulumi:"remote"`
	// Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCertificateRemote registers a new resource with the given unique name, arguments, and options.
func NewCertificateRemote(ctx *pulumi.Context,
	name string, args *CertificateRemoteArgs, opts ...pulumi.ResourceOption) (*CertificateRemote, error) {
	if args == nil {
		args = &CertificateRemoteArgs{}
	}

	var resource CertificateRemote
	err := ctx.RegisterResource("fortios:index/certificateRemote:CertificateRemote", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateRemote gets an existing CertificateRemote resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateRemote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateRemoteState, opts ...pulumi.ResourceOption) (*CertificateRemote, error) {
	var resource CertificateRemote
	err := ctx.ReadResource("fortios:index/certificateRemote:CertificateRemote", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateRemote resources.
type certificateRemoteState struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CertificateRemoteState struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CertificateRemoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateRemoteState)(nil)).Elem()
}

type certificateRemoteArgs struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CertificateRemote resource.
type CertificateRemoteArgs struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CertificateRemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateRemoteArgs)(nil)).Elem()
}

type CertificateRemoteInput interface {
	pulumi.Input

	ToCertificateRemoteOutput() CertificateRemoteOutput
	ToCertificateRemoteOutputWithContext(ctx context.Context) CertificateRemoteOutput
}

func (*CertificateRemote) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateRemote)(nil))
}

func (i *CertificateRemote) ToCertificateRemoteOutput() CertificateRemoteOutput {
	return i.ToCertificateRemoteOutputWithContext(context.Background())
}

func (i *CertificateRemote) ToCertificateRemoteOutputWithContext(ctx context.Context) CertificateRemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateRemoteOutput)
}

func (i *CertificateRemote) ToCertificateRemotePtrOutput() CertificateRemotePtrOutput {
	return i.ToCertificateRemotePtrOutputWithContext(context.Background())
}

func (i *CertificateRemote) ToCertificateRemotePtrOutputWithContext(ctx context.Context) CertificateRemotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateRemotePtrOutput)
}

type CertificateRemotePtrInput interface {
	pulumi.Input

	ToCertificateRemotePtrOutput() CertificateRemotePtrOutput
	ToCertificateRemotePtrOutputWithContext(ctx context.Context) CertificateRemotePtrOutput
}

type certificateRemotePtrType CertificateRemoteArgs

func (*certificateRemotePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateRemote)(nil))
}

func (i *certificateRemotePtrType) ToCertificateRemotePtrOutput() CertificateRemotePtrOutput {
	return i.ToCertificateRemotePtrOutputWithContext(context.Background())
}

func (i *certificateRemotePtrType) ToCertificateRemotePtrOutputWithContext(ctx context.Context) CertificateRemotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateRemotePtrOutput)
}

// CertificateRemoteArrayInput is an input type that accepts CertificateRemoteArray and CertificateRemoteArrayOutput values.
// You can construct a concrete instance of `CertificateRemoteArrayInput` via:
//
//          CertificateRemoteArray{ CertificateRemoteArgs{...} }
type CertificateRemoteArrayInput interface {
	pulumi.Input

	ToCertificateRemoteArrayOutput() CertificateRemoteArrayOutput
	ToCertificateRemoteArrayOutputWithContext(context.Context) CertificateRemoteArrayOutput
}

type CertificateRemoteArray []CertificateRemoteInput

func (CertificateRemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CertificateRemote)(nil))
}

func (i CertificateRemoteArray) ToCertificateRemoteArrayOutput() CertificateRemoteArrayOutput {
	return i.ToCertificateRemoteArrayOutputWithContext(context.Background())
}

func (i CertificateRemoteArray) ToCertificateRemoteArrayOutputWithContext(ctx context.Context) CertificateRemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateRemoteArrayOutput)
}

// CertificateRemoteMapInput is an input type that accepts CertificateRemoteMap and CertificateRemoteMapOutput values.
// You can construct a concrete instance of `CertificateRemoteMapInput` via:
//
//          CertificateRemoteMap{ "key": CertificateRemoteArgs{...} }
type CertificateRemoteMapInput interface {
	pulumi.Input

	ToCertificateRemoteMapOutput() CertificateRemoteMapOutput
	ToCertificateRemoteMapOutputWithContext(context.Context) CertificateRemoteMapOutput
}

type CertificateRemoteMap map[string]CertificateRemoteInput

func (CertificateRemoteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CertificateRemote)(nil))
}

func (i CertificateRemoteMap) ToCertificateRemoteMapOutput() CertificateRemoteMapOutput {
	return i.ToCertificateRemoteMapOutputWithContext(context.Background())
}

func (i CertificateRemoteMap) ToCertificateRemoteMapOutputWithContext(ctx context.Context) CertificateRemoteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateRemoteMapOutput)
}

type CertificateRemoteOutput struct {
	*pulumi.OutputState
}

func (CertificateRemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateRemote)(nil))
}

func (o CertificateRemoteOutput) ToCertificateRemoteOutput() CertificateRemoteOutput {
	return o
}

func (o CertificateRemoteOutput) ToCertificateRemoteOutputWithContext(ctx context.Context) CertificateRemoteOutput {
	return o
}

func (o CertificateRemoteOutput) ToCertificateRemotePtrOutput() CertificateRemotePtrOutput {
	return o.ToCertificateRemotePtrOutputWithContext(context.Background())
}

func (o CertificateRemoteOutput) ToCertificateRemotePtrOutputWithContext(ctx context.Context) CertificateRemotePtrOutput {
	return o.ApplyT(func(v CertificateRemote) *CertificateRemote {
		return &v
	}).(CertificateRemotePtrOutput)
}

type CertificateRemotePtrOutput struct {
	*pulumi.OutputState
}

func (CertificateRemotePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateRemote)(nil))
}

func (o CertificateRemotePtrOutput) ToCertificateRemotePtrOutput() CertificateRemotePtrOutput {
	return o
}

func (o CertificateRemotePtrOutput) ToCertificateRemotePtrOutputWithContext(ctx context.Context) CertificateRemotePtrOutput {
	return o
}

type CertificateRemoteArrayOutput struct{ *pulumi.OutputState }

func (CertificateRemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateRemote)(nil))
}

func (o CertificateRemoteArrayOutput) ToCertificateRemoteArrayOutput() CertificateRemoteArrayOutput {
	return o
}

func (o CertificateRemoteArrayOutput) ToCertificateRemoteArrayOutputWithContext(ctx context.Context) CertificateRemoteArrayOutput {
	return o
}

func (o CertificateRemoteArrayOutput) Index(i pulumi.IntInput) CertificateRemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateRemote {
		return vs[0].([]CertificateRemote)[vs[1].(int)]
	}).(CertificateRemoteOutput)
}

type CertificateRemoteMapOutput struct{ *pulumi.OutputState }

func (CertificateRemoteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CertificateRemote)(nil))
}

func (o CertificateRemoteMapOutput) ToCertificateRemoteMapOutput() CertificateRemoteMapOutput {
	return o
}

func (o CertificateRemoteMapOutput) ToCertificateRemoteMapOutputWithContext(ctx context.Context) CertificateRemoteMapOutput {
	return o
}

func (o CertificateRemoteMapOutput) MapIndex(k pulumi.StringInput) CertificateRemoteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CertificateRemote {
		return vs[0].(map[string]CertificateRemote)[vs[1].(string)]
	}).(CertificateRemoteOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateRemoteOutput{})
	pulumi.RegisterOutputType(CertificateRemotePtrOutput{})
	pulumi.RegisterOutputType(CertificateRemoteArrayOutput{})
	pulumi.RegisterOutputType(CertificateRemoteMapOutput{})
}

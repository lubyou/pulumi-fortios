// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type VpnCertificateCa struct {
	pulumi.CustomResourceState

	AutoUpdateDays        pulumi.IntOutput       `pulumi:"autoUpdateDays"`
	AutoUpdateDaysWarning pulumi.IntOutput       `pulumi:"autoUpdateDaysWarning"`
	Ca                    pulumi.StringOutput    `pulumi:"ca"`
	CaIdentifier          pulumi.StringOutput    `pulumi:"caIdentifier"`
	EstUrl                pulumi.StringOutput    `pulumi:"estUrl"`
	LastUpdated           pulumi.IntOutput       `pulumi:"lastUpdated"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Obsolete              pulumi.StringOutput    `pulumi:"obsolete"`
	Range                 pulumi.StringOutput    `pulumi:"range"`
	ScepUrl               pulumi.StringOutput    `pulumi:"scepUrl"`
	Source                pulumi.StringOutput    `pulumi:"source"`
	SourceIp              pulumi.StringOutput    `pulumi:"sourceIp"`
	SslInspectionTrusted  pulumi.StringOutput    `pulumi:"sslInspectionTrusted"`
	Trusted               pulumi.StringOutput    `pulumi:"trusted"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnCertificateCa registers a new resource with the given unique name, arguments, and options.
func NewVpnCertificateCa(ctx *pulumi.Context,
	name string, args *VpnCertificateCaArgs, opts ...pulumi.ResourceOption) (*VpnCertificateCa, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ca == nil {
		return nil, errors.New("invalid value for required argument 'Ca'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnCertificateCa
	err := ctx.RegisterResource("fortios:index/vpnCertificateCa:VpnCertificateCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCertificateCa gets an existing VpnCertificateCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCertificateCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCertificateCaState, opts ...pulumi.ResourceOption) (*VpnCertificateCa, error) {
	var resource VpnCertificateCa
	err := ctx.ReadResource("fortios:index/vpnCertificateCa:VpnCertificateCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCertificateCa resources.
type vpnCertificateCaState struct {
	AutoUpdateDays        *int    `pulumi:"autoUpdateDays"`
	AutoUpdateDaysWarning *int    `pulumi:"autoUpdateDaysWarning"`
	Ca                    *string `pulumi:"ca"`
	CaIdentifier          *string `pulumi:"caIdentifier"`
	EstUrl                *string `pulumi:"estUrl"`
	LastUpdated           *int    `pulumi:"lastUpdated"`
	Name                  *string `pulumi:"name"`
	Obsolete              *string `pulumi:"obsolete"`
	Range                 *string `pulumi:"range"`
	ScepUrl               *string `pulumi:"scepUrl"`
	Source                *string `pulumi:"source"`
	SourceIp              *string `pulumi:"sourceIp"`
	SslInspectionTrusted  *string `pulumi:"sslInspectionTrusted"`
	Trusted               *string `pulumi:"trusted"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

type VpnCertificateCaState struct {
	AutoUpdateDays        pulumi.IntPtrInput
	AutoUpdateDaysWarning pulumi.IntPtrInput
	Ca                    pulumi.StringPtrInput
	CaIdentifier          pulumi.StringPtrInput
	EstUrl                pulumi.StringPtrInput
	LastUpdated           pulumi.IntPtrInput
	Name                  pulumi.StringPtrInput
	Obsolete              pulumi.StringPtrInput
	Range                 pulumi.StringPtrInput
	ScepUrl               pulumi.StringPtrInput
	Source                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslInspectionTrusted  pulumi.StringPtrInput
	Trusted               pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (VpnCertificateCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateCaState)(nil)).Elem()
}

type vpnCertificateCaArgs struct {
	AutoUpdateDays        *int    `pulumi:"autoUpdateDays"`
	AutoUpdateDaysWarning *int    `pulumi:"autoUpdateDaysWarning"`
	Ca                    string  `pulumi:"ca"`
	CaIdentifier          *string `pulumi:"caIdentifier"`
	EstUrl                *string `pulumi:"estUrl"`
	LastUpdated           *int    `pulumi:"lastUpdated"`
	Name                  *string `pulumi:"name"`
	Obsolete              *string `pulumi:"obsolete"`
	Range                 *string `pulumi:"range"`
	ScepUrl               *string `pulumi:"scepUrl"`
	Source                *string `pulumi:"source"`
	SourceIp              *string `pulumi:"sourceIp"`
	SslInspectionTrusted  *string `pulumi:"sslInspectionTrusted"`
	Trusted               *string `pulumi:"trusted"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnCertificateCa resource.
type VpnCertificateCaArgs struct {
	AutoUpdateDays        pulumi.IntPtrInput
	AutoUpdateDaysWarning pulumi.IntPtrInput
	Ca                    pulumi.StringInput
	CaIdentifier          pulumi.StringPtrInput
	EstUrl                pulumi.StringPtrInput
	LastUpdated           pulumi.IntPtrInput
	Name                  pulumi.StringPtrInput
	Obsolete              pulumi.StringPtrInput
	Range                 pulumi.StringPtrInput
	ScepUrl               pulumi.StringPtrInput
	Source                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslInspectionTrusted  pulumi.StringPtrInput
	Trusted               pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (VpnCertificateCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateCaArgs)(nil)).Elem()
}

type VpnCertificateCaInput interface {
	pulumi.Input

	ToVpnCertificateCaOutput() VpnCertificateCaOutput
	ToVpnCertificateCaOutputWithContext(ctx context.Context) VpnCertificateCaOutput
}

func (*VpnCertificateCa) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateCa)(nil)).Elem()
}

func (i *VpnCertificateCa) ToVpnCertificateCaOutput() VpnCertificateCaOutput {
	return i.ToVpnCertificateCaOutputWithContext(context.Background())
}

func (i *VpnCertificateCa) ToVpnCertificateCaOutputWithContext(ctx context.Context) VpnCertificateCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateCaOutput)
}

func (i *VpnCertificateCa) ToOutput(ctx context.Context) pulumix.Output[*VpnCertificateCa] {
	return pulumix.Output[*VpnCertificateCa]{
		OutputState: i.ToVpnCertificateCaOutputWithContext(ctx).OutputState,
	}
}

// VpnCertificateCaArrayInput is an input type that accepts VpnCertificateCaArray and VpnCertificateCaArrayOutput values.
// You can construct a concrete instance of `VpnCertificateCaArrayInput` via:
//
//	VpnCertificateCaArray{ VpnCertificateCaArgs{...} }
type VpnCertificateCaArrayInput interface {
	pulumi.Input

	ToVpnCertificateCaArrayOutput() VpnCertificateCaArrayOutput
	ToVpnCertificateCaArrayOutputWithContext(context.Context) VpnCertificateCaArrayOutput
}

type VpnCertificateCaArray []VpnCertificateCaInput

func (VpnCertificateCaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateCa)(nil)).Elem()
}

func (i VpnCertificateCaArray) ToVpnCertificateCaArrayOutput() VpnCertificateCaArrayOutput {
	return i.ToVpnCertificateCaArrayOutputWithContext(context.Background())
}

func (i VpnCertificateCaArray) ToVpnCertificateCaArrayOutputWithContext(ctx context.Context) VpnCertificateCaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateCaArrayOutput)
}

func (i VpnCertificateCaArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpnCertificateCa] {
	return pulumix.Output[[]*VpnCertificateCa]{
		OutputState: i.ToVpnCertificateCaArrayOutputWithContext(ctx).OutputState,
	}
}

// VpnCertificateCaMapInput is an input type that accepts VpnCertificateCaMap and VpnCertificateCaMapOutput values.
// You can construct a concrete instance of `VpnCertificateCaMapInput` via:
//
//	VpnCertificateCaMap{ "key": VpnCertificateCaArgs{...} }
type VpnCertificateCaMapInput interface {
	pulumi.Input

	ToVpnCertificateCaMapOutput() VpnCertificateCaMapOutput
	ToVpnCertificateCaMapOutputWithContext(context.Context) VpnCertificateCaMapOutput
}

type VpnCertificateCaMap map[string]VpnCertificateCaInput

func (VpnCertificateCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateCa)(nil)).Elem()
}

func (i VpnCertificateCaMap) ToVpnCertificateCaMapOutput() VpnCertificateCaMapOutput {
	return i.ToVpnCertificateCaMapOutputWithContext(context.Background())
}

func (i VpnCertificateCaMap) ToVpnCertificateCaMapOutputWithContext(ctx context.Context) VpnCertificateCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateCaMapOutput)
}

func (i VpnCertificateCaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnCertificateCa] {
	return pulumix.Output[map[string]*VpnCertificateCa]{
		OutputState: i.ToVpnCertificateCaMapOutputWithContext(ctx).OutputState,
	}
}

type VpnCertificateCaOutput struct{ *pulumi.OutputState }

func (VpnCertificateCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateCa)(nil)).Elem()
}

func (o VpnCertificateCaOutput) ToVpnCertificateCaOutput() VpnCertificateCaOutput {
	return o
}

func (o VpnCertificateCaOutput) ToVpnCertificateCaOutputWithContext(ctx context.Context) VpnCertificateCaOutput {
	return o
}

func (o VpnCertificateCaOutput) ToOutput(ctx context.Context) pulumix.Output[*VpnCertificateCa] {
	return pulumix.Output[*VpnCertificateCa]{
		OutputState: o.OutputState,
	}
}

func (o VpnCertificateCaOutput) AutoUpdateDays() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.IntOutput { return v.AutoUpdateDays }).(pulumi.IntOutput)
}

func (o VpnCertificateCaOutput) AutoUpdateDaysWarning() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.IntOutput { return v.AutoUpdateDaysWarning }).(pulumi.IntOutput)
}

func (o VpnCertificateCaOutput) Ca() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Ca }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) CaIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.CaIdentifier }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) EstUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.EstUrl }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

func (o VpnCertificateCaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) Obsolete() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Obsolete }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) SslInspectionTrusted() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.SslInspectionTrusted }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) Trusted() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringOutput { return v.Trusted }).(pulumi.StringOutput)
}

func (o VpnCertificateCaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCertificateCa) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnCertificateCaArrayOutput struct{ *pulumi.OutputState }

func (VpnCertificateCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateCa)(nil)).Elem()
}

func (o VpnCertificateCaArrayOutput) ToVpnCertificateCaArrayOutput() VpnCertificateCaArrayOutput {
	return o
}

func (o VpnCertificateCaArrayOutput) ToVpnCertificateCaArrayOutputWithContext(ctx context.Context) VpnCertificateCaArrayOutput {
	return o
}

func (o VpnCertificateCaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpnCertificateCa] {
	return pulumix.Output[[]*VpnCertificateCa]{
		OutputState: o.OutputState,
	}
}

func (o VpnCertificateCaArrayOutput) Index(i pulumi.IntInput) VpnCertificateCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnCertificateCa {
		return vs[0].([]*VpnCertificateCa)[vs[1].(int)]
	}).(VpnCertificateCaOutput)
}

type VpnCertificateCaMapOutput struct{ *pulumi.OutputState }

func (VpnCertificateCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateCa)(nil)).Elem()
}

func (o VpnCertificateCaMapOutput) ToVpnCertificateCaMapOutput() VpnCertificateCaMapOutput {
	return o
}

func (o VpnCertificateCaMapOutput) ToVpnCertificateCaMapOutputWithContext(ctx context.Context) VpnCertificateCaMapOutput {
	return o
}

func (o VpnCertificateCaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnCertificateCa] {
	return pulumix.Output[map[string]*VpnCertificateCa]{
		OutputState: o.OutputState,
	}
}

func (o VpnCertificateCaMapOutput) MapIndex(k pulumi.StringInput) VpnCertificateCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnCertificateCa {
		return vs[0].(map[string]*VpnCertificateCa)[vs[1].(string)]
	}).(VpnCertificateCaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateCaInput)(nil)).Elem(), &VpnCertificateCa{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateCaArrayInput)(nil)).Elem(), VpnCertificateCaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateCaMapInput)(nil)).Elem(), VpnCertificateCaMap{})
	pulumi.RegisterOutputType(VpnCertificateCaOutput{})
	pulumi.RegisterOutputType(VpnCertificateCaArrayOutput{})
	pulumi.RegisterOutputType(VpnCertificateCaMapOutput{})
}

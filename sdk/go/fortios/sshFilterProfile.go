// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SshFilterProfile struct {
	pulumi.CustomResourceState

	Block               pulumi.StringOutput                     `pulumi:"block"`
	DefaultCommandLog   pulumi.StringOutput                     `pulumi:"defaultCommandLog"`
	DynamicSortSubtable pulumi.StringPtrOutput                  `pulumi:"dynamicSortSubtable"`
	FileFilter          SshFilterProfileFileFilterOutput        `pulumi:"fileFilter"`
	Log                 pulumi.StringOutput                     `pulumi:"log"`
	Name                pulumi.StringOutput                     `pulumi:"name"`
	ShellCommands       SshFilterProfileShellCommandArrayOutput `pulumi:"shellCommands"`
	Vdomparam           pulumi.StringPtrOutput                  `pulumi:"vdomparam"`
}

// NewSshFilterProfile registers a new resource with the given unique name, arguments, and options.
func NewSshFilterProfile(ctx *pulumi.Context,
	name string, args *SshFilterProfileArgs, opts ...pulumi.ResourceOption) (*SshFilterProfile, error) {
	if args == nil {
		args = &SshFilterProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SshFilterProfile
	err := ctx.RegisterResource("fortios:index/sshFilterProfile:SshFilterProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshFilterProfile gets an existing SshFilterProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshFilterProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshFilterProfileState, opts ...pulumi.ResourceOption) (*SshFilterProfile, error) {
	var resource SshFilterProfile
	err := ctx.ReadResource("fortios:index/sshFilterProfile:SshFilterProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshFilterProfile resources.
type sshFilterProfileState struct {
	Block               *string                        `pulumi:"block"`
	DefaultCommandLog   *string                        `pulumi:"defaultCommandLog"`
	DynamicSortSubtable *string                        `pulumi:"dynamicSortSubtable"`
	FileFilter          *SshFilterProfileFileFilter    `pulumi:"fileFilter"`
	Log                 *string                        `pulumi:"log"`
	Name                *string                        `pulumi:"name"`
	ShellCommands       []SshFilterProfileShellCommand `pulumi:"shellCommands"`
	Vdomparam           *string                        `pulumi:"vdomparam"`
}

type SshFilterProfileState struct {
	Block               pulumi.StringPtrInput
	DefaultCommandLog   pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	FileFilter          SshFilterProfileFileFilterPtrInput
	Log                 pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ShellCommands       SshFilterProfileShellCommandArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (SshFilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshFilterProfileState)(nil)).Elem()
}

type sshFilterProfileArgs struct {
	Block               *string                        `pulumi:"block"`
	DefaultCommandLog   *string                        `pulumi:"defaultCommandLog"`
	DynamicSortSubtable *string                        `pulumi:"dynamicSortSubtable"`
	FileFilter          *SshFilterProfileFileFilter    `pulumi:"fileFilter"`
	Log                 *string                        `pulumi:"log"`
	Name                *string                        `pulumi:"name"`
	ShellCommands       []SshFilterProfileShellCommand `pulumi:"shellCommands"`
	Vdomparam           *string                        `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SshFilterProfile resource.
type SshFilterProfileArgs struct {
	Block               pulumi.StringPtrInput
	DefaultCommandLog   pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	FileFilter          SshFilterProfileFileFilterPtrInput
	Log                 pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ShellCommands       SshFilterProfileShellCommandArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (SshFilterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshFilterProfileArgs)(nil)).Elem()
}

type SshFilterProfileInput interface {
	pulumi.Input

	ToSshFilterProfileOutput() SshFilterProfileOutput
	ToSshFilterProfileOutputWithContext(ctx context.Context) SshFilterProfileOutput
}

func (*SshFilterProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**SshFilterProfile)(nil)).Elem()
}

func (i *SshFilterProfile) ToSshFilterProfileOutput() SshFilterProfileOutput {
	return i.ToSshFilterProfileOutputWithContext(context.Background())
}

func (i *SshFilterProfile) ToSshFilterProfileOutputWithContext(ctx context.Context) SshFilterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshFilterProfileOutput)
}

// SshFilterProfileArrayInput is an input type that accepts SshFilterProfileArray and SshFilterProfileArrayOutput values.
// You can construct a concrete instance of `SshFilterProfileArrayInput` via:
//
//	SshFilterProfileArray{ SshFilterProfileArgs{...} }
type SshFilterProfileArrayInput interface {
	pulumi.Input

	ToSshFilterProfileArrayOutput() SshFilterProfileArrayOutput
	ToSshFilterProfileArrayOutputWithContext(context.Context) SshFilterProfileArrayOutput
}

type SshFilterProfileArray []SshFilterProfileInput

func (SshFilterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshFilterProfile)(nil)).Elem()
}

func (i SshFilterProfileArray) ToSshFilterProfileArrayOutput() SshFilterProfileArrayOutput {
	return i.ToSshFilterProfileArrayOutputWithContext(context.Background())
}

func (i SshFilterProfileArray) ToSshFilterProfileArrayOutputWithContext(ctx context.Context) SshFilterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshFilterProfileArrayOutput)
}

// SshFilterProfileMapInput is an input type that accepts SshFilterProfileMap and SshFilterProfileMapOutput values.
// You can construct a concrete instance of `SshFilterProfileMapInput` via:
//
//	SshFilterProfileMap{ "key": SshFilterProfileArgs{...} }
type SshFilterProfileMapInput interface {
	pulumi.Input

	ToSshFilterProfileMapOutput() SshFilterProfileMapOutput
	ToSshFilterProfileMapOutputWithContext(context.Context) SshFilterProfileMapOutput
}

type SshFilterProfileMap map[string]SshFilterProfileInput

func (SshFilterProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshFilterProfile)(nil)).Elem()
}

func (i SshFilterProfileMap) ToSshFilterProfileMapOutput() SshFilterProfileMapOutput {
	return i.ToSshFilterProfileMapOutputWithContext(context.Background())
}

func (i SshFilterProfileMap) ToSshFilterProfileMapOutputWithContext(ctx context.Context) SshFilterProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshFilterProfileMapOutput)
}

type SshFilterProfileOutput struct{ *pulumi.OutputState }

func (SshFilterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshFilterProfile)(nil)).Elem()
}

func (o SshFilterProfileOutput) ToSshFilterProfileOutput() SshFilterProfileOutput {
	return o
}

func (o SshFilterProfileOutput) ToSshFilterProfileOutputWithContext(ctx context.Context) SshFilterProfileOutput {
	return o
}

func (o SshFilterProfileOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringOutput { return v.Block }).(pulumi.StringOutput)
}

func (o SshFilterProfileOutput) DefaultCommandLog() pulumi.StringOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringOutput { return v.DefaultCommandLog }).(pulumi.StringOutput)
}

func (o SshFilterProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SshFilterProfileOutput) FileFilter() SshFilterProfileFileFilterOutput {
	return o.ApplyT(func(v *SshFilterProfile) SshFilterProfileFileFilterOutput { return v.FileFilter }).(SshFilterProfileFileFilterOutput)
}

func (o SshFilterProfileOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringOutput { return v.Log }).(pulumi.StringOutput)
}

func (o SshFilterProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SshFilterProfileOutput) ShellCommands() SshFilterProfileShellCommandArrayOutput {
	return o.ApplyT(func(v *SshFilterProfile) SshFilterProfileShellCommandArrayOutput { return v.ShellCommands }).(SshFilterProfileShellCommandArrayOutput)
}

func (o SshFilterProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshFilterProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SshFilterProfileArrayOutput struct{ *pulumi.OutputState }

func (SshFilterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshFilterProfile)(nil)).Elem()
}

func (o SshFilterProfileArrayOutput) ToSshFilterProfileArrayOutput() SshFilterProfileArrayOutput {
	return o
}

func (o SshFilterProfileArrayOutput) ToSshFilterProfileArrayOutputWithContext(ctx context.Context) SshFilterProfileArrayOutput {
	return o
}

func (o SshFilterProfileArrayOutput) Index(i pulumi.IntInput) SshFilterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SshFilterProfile {
		return vs[0].([]*SshFilterProfile)[vs[1].(int)]
	}).(SshFilterProfileOutput)
}

type SshFilterProfileMapOutput struct{ *pulumi.OutputState }

func (SshFilterProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshFilterProfile)(nil)).Elem()
}

func (o SshFilterProfileMapOutput) ToSshFilterProfileMapOutput() SshFilterProfileMapOutput {
	return o
}

func (o SshFilterProfileMapOutput) ToSshFilterProfileMapOutputWithContext(ctx context.Context) SshFilterProfileMapOutput {
	return o
}

func (o SshFilterProfileMapOutput) MapIndex(k pulumi.StringInput) SshFilterProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SshFilterProfile {
		return vs[0].(map[string]*SshFilterProfile)[vs[1].(string)]
	}).(SshFilterProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SshFilterProfileInput)(nil)).Elem(), &SshFilterProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshFilterProfileArrayInput)(nil)).Elem(), SshFilterProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshFilterProfileMapInput)(nil)).Elem(), SshFilterProfileMap{})
	pulumi.RegisterOutputType(SshFilterProfileOutput{})
	pulumi.RegisterOutputType(SshFilterProfileArrayOutput{})
	pulumi.RegisterOutputType(SshFilterProfileMapOutput{})
}

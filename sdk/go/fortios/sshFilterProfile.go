// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSH filter profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSshFilterProfile(ctx, "trname", &fortios.SshFilterProfileArgs{
// 			Block:             pulumi.String("x11"),
// 			DefaultCommandLog: pulumi.String("enable"),
// 			Log:               pulumi.String("x11"),
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
// SshFilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/sshFilterProfile:SshFilterProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SshFilterProfile struct {
	pulumi.CustomResourceState

	// SSH blocking options.
	Block pulumi.StringOutput `pulumi:"block"`
	// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
	DefaultCommandLog pulumi.StringOutput `pulumi:"defaultCommandLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter SshFilterProfileFileFilterPtrOutput `pulumi:"fileFilter"`
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// File type name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SSH command filter. The structure of `shellCommands` block is documented below.
	ShellCommands SshFilterProfileShellCommandArrayOutput `pulumi:"shellCommands"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
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
	// SSH blocking options.
	Block *string `pulumi:"block"`
	// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
	DefaultCommandLog *string `pulumi:"defaultCommandLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *SshFilterProfileFileFilter `pulumi:"fileFilter"`
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// File type name.
	Name *string `pulumi:"name"`
	// SSH command filter. The structure of `shellCommands` block is documented below.
	ShellCommands []SshFilterProfileShellCommand `pulumi:"shellCommands"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SshFilterProfileState struct {
	// SSH blocking options.
	Block pulumi.StringPtrInput
	// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
	DefaultCommandLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter SshFilterProfileFileFilterPtrInput
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput
	// File type name.
	Name pulumi.StringPtrInput
	// SSH command filter. The structure of `shellCommands` block is documented below.
	ShellCommands SshFilterProfileShellCommandArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SshFilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshFilterProfileState)(nil)).Elem()
}

type sshFilterProfileArgs struct {
	// SSH blocking options.
	Block *string `pulumi:"block"`
	// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
	DefaultCommandLog *string `pulumi:"defaultCommandLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *SshFilterProfileFileFilter `pulumi:"fileFilter"`
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
	// File type name.
	Name *string `pulumi:"name"`
	// SSH command filter. The structure of `shellCommands` block is documented below.
	ShellCommands []SshFilterProfileShellCommand `pulumi:"shellCommands"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SshFilterProfile resource.
type SshFilterProfileArgs struct {
	// SSH blocking options.
	Block pulumi.StringPtrInput
	// Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
	DefaultCommandLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter SshFilterProfileFileFilterPtrInput
	// Enable/disable file filter logging. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput
	// File type name.
	Name pulumi.StringPtrInput
	// SSH command filter. The structure of `shellCommands` block is documented below.
	ShellCommands SshFilterProfileShellCommandArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          SshFilterProfileArray{ SshFilterProfileArgs{...} }
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
//          SshFilterProfileMap{ "key": SshFilterProfileArgs{...} }
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

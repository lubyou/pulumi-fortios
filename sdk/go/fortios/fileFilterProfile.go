// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure file-filter profiles. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// FileFilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/fileFilterProfile:FileFilterProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FileFilterProfile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// File type name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replacement message group
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// File filter rules. The structure of `rules` block is documented below.
	Rules FileFilterProfileRuleArrayOutput `pulumi:"rules"`
	// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
	ScanArchiveContents pulumi.StringOutput `pulumi:"scanArchiveContents"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFileFilterProfile registers a new resource with the given unique name, arguments, and options.
func NewFileFilterProfile(ctx *pulumi.Context,
	name string, args *FileFilterProfileArgs, opts ...pulumi.ResourceOption) (*FileFilterProfile, error) {
	if args == nil {
		args = &FileFilterProfileArgs{}
	}

	var resource FileFilterProfile
	err := ctx.RegisterResource("fortios:index/fileFilterProfile:FileFilterProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileFilterProfile gets an existing FileFilterProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileFilterProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileFilterProfileState, opts ...pulumi.ResourceOption) (*FileFilterProfile, error) {
	var resource FileFilterProfile
	err := ctx.ReadResource("fortios:index/fileFilterProfile:FileFilterProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileFilterProfile resources.
type fileFilterProfileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// File type name.
	Name *string `pulumi:"name"`
	// Replacement message group
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// File filter rules. The structure of `rules` block is documented below.
	Rules []FileFilterProfileRule `pulumi:"rules"`
	// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
	ScanArchiveContents *string `pulumi:"scanArchiveContents"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FileFilterProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// File type name.
	Name pulumi.StringPtrInput
	// Replacement message group
	ReplacemsgGroup pulumi.StringPtrInput
	// File filter rules. The structure of `rules` block is documented below.
	Rules FileFilterProfileRuleArrayInput
	// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
	ScanArchiveContents pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FileFilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileFilterProfileState)(nil)).Elem()
}

type fileFilterProfileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// File type name.
	Name *string `pulumi:"name"`
	// Replacement message group
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// File filter rules. The structure of `rules` block is documented below.
	Rules []FileFilterProfileRule `pulumi:"rules"`
	// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
	ScanArchiveContents *string `pulumi:"scanArchiveContents"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FileFilterProfile resource.
type FileFilterProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Enable/disable file-filter logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// File type name.
	Name pulumi.StringPtrInput
	// Replacement message group
	ReplacemsgGroup pulumi.StringPtrInput
	// File filter rules. The structure of `rules` block is documented below.
	Rules FileFilterProfileRuleArrayInput
	// Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
	ScanArchiveContents pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FileFilterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileFilterProfileArgs)(nil)).Elem()
}

type FileFilterProfileInput interface {
	pulumi.Input

	ToFileFilterProfileOutput() FileFilterProfileOutput
	ToFileFilterProfileOutputWithContext(ctx context.Context) FileFilterProfileOutput
}

func (*FileFilterProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*FileFilterProfile)(nil))
}

func (i *FileFilterProfile) ToFileFilterProfileOutput() FileFilterProfileOutput {
	return i.ToFileFilterProfileOutputWithContext(context.Background())
}

func (i *FileFilterProfile) ToFileFilterProfileOutputWithContext(ctx context.Context) FileFilterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileFilterProfileOutput)
}

func (i *FileFilterProfile) ToFileFilterProfilePtrOutput() FileFilterProfilePtrOutput {
	return i.ToFileFilterProfilePtrOutputWithContext(context.Background())
}

func (i *FileFilterProfile) ToFileFilterProfilePtrOutputWithContext(ctx context.Context) FileFilterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileFilterProfilePtrOutput)
}

type FileFilterProfilePtrInput interface {
	pulumi.Input

	ToFileFilterProfilePtrOutput() FileFilterProfilePtrOutput
	ToFileFilterProfilePtrOutputWithContext(ctx context.Context) FileFilterProfilePtrOutput
}

type fileFilterProfilePtrType FileFilterProfileArgs

func (*fileFilterProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileFilterProfile)(nil))
}

func (i *fileFilterProfilePtrType) ToFileFilterProfilePtrOutput() FileFilterProfilePtrOutput {
	return i.ToFileFilterProfilePtrOutputWithContext(context.Background())
}

func (i *fileFilterProfilePtrType) ToFileFilterProfilePtrOutputWithContext(ctx context.Context) FileFilterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileFilterProfilePtrOutput)
}

// FileFilterProfileArrayInput is an input type that accepts FileFilterProfileArray and FileFilterProfileArrayOutput values.
// You can construct a concrete instance of `FileFilterProfileArrayInput` via:
//
//          FileFilterProfileArray{ FileFilterProfileArgs{...} }
type FileFilterProfileArrayInput interface {
	pulumi.Input

	ToFileFilterProfileArrayOutput() FileFilterProfileArrayOutput
	ToFileFilterProfileArrayOutputWithContext(context.Context) FileFilterProfileArrayOutput
}

type FileFilterProfileArray []FileFilterProfileInput

func (FileFilterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FileFilterProfile)(nil))
}

func (i FileFilterProfileArray) ToFileFilterProfileArrayOutput() FileFilterProfileArrayOutput {
	return i.ToFileFilterProfileArrayOutputWithContext(context.Background())
}

func (i FileFilterProfileArray) ToFileFilterProfileArrayOutputWithContext(ctx context.Context) FileFilterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileFilterProfileArrayOutput)
}

// FileFilterProfileMapInput is an input type that accepts FileFilterProfileMap and FileFilterProfileMapOutput values.
// You can construct a concrete instance of `FileFilterProfileMapInput` via:
//
//          FileFilterProfileMap{ "key": FileFilterProfileArgs{...} }
type FileFilterProfileMapInput interface {
	pulumi.Input

	ToFileFilterProfileMapOutput() FileFilterProfileMapOutput
	ToFileFilterProfileMapOutputWithContext(context.Context) FileFilterProfileMapOutput
}

type FileFilterProfileMap map[string]FileFilterProfileInput

func (FileFilterProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FileFilterProfile)(nil))
}

func (i FileFilterProfileMap) ToFileFilterProfileMapOutput() FileFilterProfileMapOutput {
	return i.ToFileFilterProfileMapOutputWithContext(context.Background())
}

func (i FileFilterProfileMap) ToFileFilterProfileMapOutputWithContext(ctx context.Context) FileFilterProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileFilterProfileMapOutput)
}

type FileFilterProfileOutput struct {
	*pulumi.OutputState
}

func (FileFilterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileFilterProfile)(nil))
}

func (o FileFilterProfileOutput) ToFileFilterProfileOutput() FileFilterProfileOutput {
	return o
}

func (o FileFilterProfileOutput) ToFileFilterProfileOutputWithContext(ctx context.Context) FileFilterProfileOutput {
	return o
}

func (o FileFilterProfileOutput) ToFileFilterProfilePtrOutput() FileFilterProfilePtrOutput {
	return o.ToFileFilterProfilePtrOutputWithContext(context.Background())
}

func (o FileFilterProfileOutput) ToFileFilterProfilePtrOutputWithContext(ctx context.Context) FileFilterProfilePtrOutput {
	return o.ApplyT(func(v FileFilterProfile) *FileFilterProfile {
		return &v
	}).(FileFilterProfilePtrOutput)
}

type FileFilterProfilePtrOutput struct {
	*pulumi.OutputState
}

func (FileFilterProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileFilterProfile)(nil))
}

func (o FileFilterProfilePtrOutput) ToFileFilterProfilePtrOutput() FileFilterProfilePtrOutput {
	return o
}

func (o FileFilterProfilePtrOutput) ToFileFilterProfilePtrOutputWithContext(ctx context.Context) FileFilterProfilePtrOutput {
	return o
}

type FileFilterProfileArrayOutput struct{ *pulumi.OutputState }

func (FileFilterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileFilterProfile)(nil))
}

func (o FileFilterProfileArrayOutput) ToFileFilterProfileArrayOutput() FileFilterProfileArrayOutput {
	return o
}

func (o FileFilterProfileArrayOutput) ToFileFilterProfileArrayOutputWithContext(ctx context.Context) FileFilterProfileArrayOutput {
	return o
}

func (o FileFilterProfileArrayOutput) Index(i pulumi.IntInput) FileFilterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileFilterProfile {
		return vs[0].([]FileFilterProfile)[vs[1].(int)]
	}).(FileFilterProfileOutput)
}

type FileFilterProfileMapOutput struct{ *pulumi.OutputState }

func (FileFilterProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FileFilterProfile)(nil))
}

func (o FileFilterProfileMapOutput) ToFileFilterProfileMapOutput() FileFilterProfileMapOutput {
	return o
}

func (o FileFilterProfileMapOutput) ToFileFilterProfileMapOutputWithContext(ctx context.Context) FileFilterProfileMapOutput {
	return o
}

func (o FileFilterProfileMapOutput) MapIndex(k pulumi.StringInput) FileFilterProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FileFilterProfile {
		return vs[0].(map[string]FileFilterProfile)[vs[1].(string)]
	}).(FileFilterProfileOutput)
}

func init() {
	pulumi.RegisterOutputType(FileFilterProfileOutput{})
	pulumi.RegisterOutputType(FileFilterProfilePtrOutput{})
	pulumi.RegisterOutputType(FileFilterProfileArrayOutput{})
	pulumi.RegisterOutputType(FileFilterProfileMapOutput{})
}

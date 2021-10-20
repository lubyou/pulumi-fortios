// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure application control lists.
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
// 		_, err := fortios.NewApplicationList(ctx, "trname", &fortios.ApplicationListArgs{
// 			AppReplacemsg:            pulumi.String("enable"),
// 			DeepAppInspection:        pulumi.String("enable"),
// 			EnforceDefaultAppPort:    pulumi.String("disable"),
// 			ExtendedLog:              pulumi.String("disable"),
// 			Options:                  pulumi.String("allow-dns"),
// 			OtherApplicationAction:   pulumi.String("pass"),
// 			OtherApplicationLog:      pulumi.String("disable"),
// 			UnknownApplicationAction: pulumi.String("pass"),
// 			UnknownApplicationLog:    pulumi.String("disable"),
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
// Application List can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/applicationList:ApplicationList labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type ApplicationList struct {
	pulumi.CustomResourceState

	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringOutput `pulumi:"appReplacemsg"`
	// comments
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringOutput `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringOutput `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ApplicationListDefaultNetworkServiceArrayOutput `pulumi:"defaultNetworkServices"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringOutput `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries ApplicationListEntryArrayOutput `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringOutput `pulumi:"forceInclusionSslDiSigs"`
	// Parameter name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringOutput `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringOutput `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringOutput `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringOutput `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringOutput `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringOutput `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringOutput `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewApplicationList registers a new resource with the given unique name, arguments, and options.
func NewApplicationList(ctx *pulumi.Context,
	name string, args *ApplicationListArgs, opts ...pulumi.ResourceOption) (*ApplicationList, error) {
	if args == nil {
		args = &ApplicationListArgs{}
	}

	var resource ApplicationList
	err := ctx.RegisterResource("fortios:index/applicationList:ApplicationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationList gets an existing ApplicationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationListState, opts ...pulumi.ResourceOption) (*ApplicationList, error) {
	var resource ApplicationList
	err := ctx.ReadResource("fortios:index/applicationList:ApplicationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationList resources.
type applicationListState struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg *string `pulumi:"appReplacemsg"`
	// comments
	Comment *string `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices *string `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection *string `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices []ApplicationListDefaultNetworkService `pulumi:"defaultNetworkServices"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort *string `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries []ApplicationListEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs *string `pulumi:"forceInclusionSslDiSigs"`
	// Parameter name.
	Name *string `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options *string `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction *string `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog *string `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList *string `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList *string `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction *string `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog *string `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ApplicationListState struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringPtrInput
	// comments
	Comment pulumi.StringPtrInput
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringPtrInput
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringPtrInput
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ApplicationListDefaultNetworkServiceArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringPtrInput
	// Application list entries. The structure of `entries` block is documented below.
	Entries ApplicationListEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringPtrInput
	// Parameter name.
	Name pulumi.StringPtrInput
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringPtrInput
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringPtrInput
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringPtrInput
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ApplicationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationListState)(nil)).Elem()
}

type applicationListArgs struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg *string `pulumi:"appReplacemsg"`
	// comments
	Comment *string `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices *string `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection *string `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices []ApplicationListDefaultNetworkService `pulumi:"defaultNetworkServices"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort *string `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries []ApplicationListEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs *string `pulumi:"forceInclusionSslDiSigs"`
	// Parameter name.
	Name *string `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options *string `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction *string `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog *string `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList *string `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList *string `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction *string `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog *string `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ApplicationList resource.
type ApplicationListArgs struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringPtrInput
	// comments
	Comment pulumi.StringPtrInput
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringPtrInput
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringPtrInput
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ApplicationListDefaultNetworkServiceArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringPtrInput
	// Application list entries. The structure of `entries` block is documented below.
	Entries ApplicationListEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringPtrInput
	// Parameter name.
	Name pulumi.StringPtrInput
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringPtrInput
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringPtrInput
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringPtrInput
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ApplicationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationListArgs)(nil)).Elem()
}

type ApplicationListInput interface {
	pulumi.Input

	ToApplicationListOutput() ApplicationListOutput
	ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput
}

func (*ApplicationList) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationList)(nil))
}

func (i *ApplicationList) ToApplicationListOutput() ApplicationListOutput {
	return i.ToApplicationListOutputWithContext(context.Background())
}

func (i *ApplicationList) ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListOutput)
}

func (i *ApplicationList) ToApplicationListPtrOutput() ApplicationListPtrOutput {
	return i.ToApplicationListPtrOutputWithContext(context.Background())
}

func (i *ApplicationList) ToApplicationListPtrOutputWithContext(ctx context.Context) ApplicationListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListPtrOutput)
}

type ApplicationListPtrInput interface {
	pulumi.Input

	ToApplicationListPtrOutput() ApplicationListPtrOutput
	ToApplicationListPtrOutputWithContext(ctx context.Context) ApplicationListPtrOutput
}

type applicationListPtrType ApplicationListArgs

func (*applicationListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationList)(nil))
}

func (i *applicationListPtrType) ToApplicationListPtrOutput() ApplicationListPtrOutput {
	return i.ToApplicationListPtrOutputWithContext(context.Background())
}

func (i *applicationListPtrType) ToApplicationListPtrOutputWithContext(ctx context.Context) ApplicationListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListPtrOutput)
}

// ApplicationListArrayInput is an input type that accepts ApplicationListArray and ApplicationListArrayOutput values.
// You can construct a concrete instance of `ApplicationListArrayInput` via:
//
//          ApplicationListArray{ ApplicationListArgs{...} }
type ApplicationListArrayInput interface {
	pulumi.Input

	ToApplicationListArrayOutput() ApplicationListArrayOutput
	ToApplicationListArrayOutputWithContext(context.Context) ApplicationListArrayOutput
}

type ApplicationListArray []ApplicationListInput

func (ApplicationListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApplicationList)(nil))
}

func (i ApplicationListArray) ToApplicationListArrayOutput() ApplicationListArrayOutput {
	return i.ToApplicationListArrayOutputWithContext(context.Background())
}

func (i ApplicationListArray) ToApplicationListArrayOutputWithContext(ctx context.Context) ApplicationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListArrayOutput)
}

// ApplicationListMapInput is an input type that accepts ApplicationListMap and ApplicationListMapOutput values.
// You can construct a concrete instance of `ApplicationListMapInput` via:
//
//          ApplicationListMap{ "key": ApplicationListArgs{...} }
type ApplicationListMapInput interface {
	pulumi.Input

	ToApplicationListMapOutput() ApplicationListMapOutput
	ToApplicationListMapOutputWithContext(context.Context) ApplicationListMapOutput
}

type ApplicationListMap map[string]ApplicationListInput

func (ApplicationListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApplicationList)(nil))
}

func (i ApplicationListMap) ToApplicationListMapOutput() ApplicationListMapOutput {
	return i.ToApplicationListMapOutputWithContext(context.Background())
}

func (i ApplicationListMap) ToApplicationListMapOutputWithContext(ctx context.Context) ApplicationListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListMapOutput)
}

type ApplicationListOutput struct {
	*pulumi.OutputState
}

func (ApplicationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationList)(nil))
}

func (o ApplicationListOutput) ToApplicationListOutput() ApplicationListOutput {
	return o
}

func (o ApplicationListOutput) ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput {
	return o
}

func (o ApplicationListOutput) ToApplicationListPtrOutput() ApplicationListPtrOutput {
	return o.ToApplicationListPtrOutputWithContext(context.Background())
}

func (o ApplicationListOutput) ToApplicationListPtrOutputWithContext(ctx context.Context) ApplicationListPtrOutput {
	return o.ApplyT(func(v ApplicationList) *ApplicationList {
		return &v
	}).(ApplicationListPtrOutput)
}

type ApplicationListPtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationList)(nil))
}

func (o ApplicationListPtrOutput) ToApplicationListPtrOutput() ApplicationListPtrOutput {
	return o
}

func (o ApplicationListPtrOutput) ToApplicationListPtrOutputWithContext(ctx context.Context) ApplicationListPtrOutput {
	return o
}

type ApplicationListArrayOutput struct{ *pulumi.OutputState }

func (ApplicationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationList)(nil))
}

func (o ApplicationListArrayOutput) ToApplicationListArrayOutput() ApplicationListArrayOutput {
	return o
}

func (o ApplicationListArrayOutput) ToApplicationListArrayOutputWithContext(ctx context.Context) ApplicationListArrayOutput {
	return o
}

func (o ApplicationListArrayOutput) Index(i pulumi.IntInput) ApplicationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationList {
		return vs[0].([]ApplicationList)[vs[1].(int)]
	}).(ApplicationListOutput)
}

type ApplicationListMapOutput struct{ *pulumi.OutputState }

func (ApplicationListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationList)(nil))
}

func (o ApplicationListMapOutput) ToApplicationListMapOutput() ApplicationListMapOutput {
	return o
}

func (o ApplicationListMapOutput) ToApplicationListMapOutputWithContext(ctx context.Context) ApplicationListMapOutput {
	return o
}

func (o ApplicationListMapOutput) MapIndex(k pulumi.StringInput) ApplicationListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationList {
		return vs[0].(map[string]ApplicationList)[vs[1].(string)]
	}).(ApplicationListOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationListOutput{})
	pulumi.RegisterOutputType(ApplicationListPtrOutput{})
	pulumi.RegisterOutputType(ApplicationListArrayOutput{})
	pulumi.RegisterOutputType(ApplicationListMapOutput{})
}
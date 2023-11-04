// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type RouterIsis struct {
	pulumi.CustomResourceState

	AdjacencyCheck       pulumi.StringOutput                  `pulumi:"adjacencyCheck"`
	AdjacencyCheck6      pulumi.StringOutput                  `pulumi:"adjacencyCheck6"`
	AdvPassiveOnly       pulumi.StringOutput                  `pulumi:"advPassiveOnly"`
	AdvPassiveOnly6      pulumi.StringOutput                  `pulumi:"advPassiveOnly6"`
	AuthKeychainL1       pulumi.StringOutput                  `pulumi:"authKeychainL1"`
	AuthKeychainL2       pulumi.StringOutput                  `pulumi:"authKeychainL2"`
	AuthModeL1           pulumi.StringOutput                  `pulumi:"authModeL1"`
	AuthModeL2           pulumi.StringOutput                  `pulumi:"authModeL2"`
	AuthPasswordL1       pulumi.StringPtrOutput               `pulumi:"authPasswordL1"`
	AuthPasswordL2       pulumi.StringPtrOutput               `pulumi:"authPasswordL2"`
	AuthSendonlyL1       pulumi.StringOutput                  `pulumi:"authSendonlyL1"`
	AuthSendonlyL2       pulumi.StringOutput                  `pulumi:"authSendonlyL2"`
	DefaultOriginate     pulumi.StringOutput                  `pulumi:"defaultOriginate"`
	DefaultOriginate6    pulumi.StringOutput                  `pulumi:"defaultOriginate6"`
	DynamicHostname      pulumi.StringOutput                  `pulumi:"dynamicHostname"`
	DynamicSortSubtable  pulumi.StringPtrOutput               `pulumi:"dynamicSortSubtable"`
	GetAllTables         pulumi.StringPtrOutput               `pulumi:"getAllTables"`
	IgnoreLspErrors      pulumi.StringOutput                  `pulumi:"ignoreLspErrors"`
	IsType               pulumi.StringOutput                  `pulumi:"isType"`
	IsisInterfaces       RouterIsisIsisInterfaceArrayOutput   `pulumi:"isisInterfaces"`
	IsisNets             RouterIsisIsisNetArrayOutput         `pulumi:"isisNets"`
	LspGenIntervalL1     pulumi.IntOutput                     `pulumi:"lspGenIntervalL1"`
	LspGenIntervalL2     pulumi.IntOutput                     `pulumi:"lspGenIntervalL2"`
	LspRefreshInterval   pulumi.IntOutput                     `pulumi:"lspRefreshInterval"`
	MaxLspLifetime       pulumi.IntOutput                     `pulumi:"maxLspLifetime"`
	MetricStyle          pulumi.StringOutput                  `pulumi:"metricStyle"`
	OverloadBit          pulumi.StringOutput                  `pulumi:"overloadBit"`
	OverloadBitOnStartup pulumi.IntOutput                     `pulumi:"overloadBitOnStartup"`
	OverloadBitSuppress  pulumi.StringOutput                  `pulumi:"overloadBitSuppress"`
	Redistribute6L1      pulumi.StringOutput                  `pulumi:"redistribute6L1"`
	Redistribute6L1List  pulumi.StringOutput                  `pulumi:"redistribute6L1List"`
	Redistribute6L2      pulumi.StringOutput                  `pulumi:"redistribute6L2"`
	Redistribute6L2List  pulumi.StringOutput                  `pulumi:"redistribute6L2List"`
	Redistribute6s       RouterIsisRedistribute6ArrayOutput   `pulumi:"redistribute6s"`
	RedistributeL1       pulumi.StringOutput                  `pulumi:"redistributeL1"`
	RedistributeL1List   pulumi.StringOutput                  `pulumi:"redistributeL1List"`
	RedistributeL2       pulumi.StringOutput                  `pulumi:"redistributeL2"`
	RedistributeL2List   pulumi.StringOutput                  `pulumi:"redistributeL2List"`
	Redistributes        RouterIsisRedistributeArrayOutput    `pulumi:"redistributes"`
	SpfIntervalExpL1     pulumi.StringOutput                  `pulumi:"spfIntervalExpL1"`
	SpfIntervalExpL2     pulumi.StringOutput                  `pulumi:"spfIntervalExpL2"`
	SummaryAddress6s     RouterIsisSummaryAddress6ArrayOutput `pulumi:"summaryAddress6s"`
	SummaryAddresses     RouterIsisSummaryAddressArrayOutput  `pulumi:"summaryAddresses"`
	Vdomparam            pulumi.StringPtrOutput               `pulumi:"vdomparam"`
}

// NewRouterIsis registers a new resource with the given unique name, arguments, and options.
func NewRouterIsis(ctx *pulumi.Context,
	name string, args *RouterIsisArgs, opts ...pulumi.ResourceOption) (*RouterIsis, error) {
	if args == nil {
		args = &RouterIsisArgs{}
	}

	if args.AuthPasswordL1 != nil {
		args.AuthPasswordL1 = pulumi.ToSecret(args.AuthPasswordL1).(pulumi.StringPtrInput)
	}
	if args.AuthPasswordL2 != nil {
		args.AuthPasswordL2 = pulumi.ToSecret(args.AuthPasswordL2).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authPasswordL1",
		"authPasswordL2",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterIsis
	err := ctx.RegisterResource("fortios:index/routerIsis:RouterIsis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterIsis gets an existing RouterIsis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterIsis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterIsisState, opts ...pulumi.ResourceOption) (*RouterIsis, error) {
	var resource RouterIsis
	err := ctx.ReadResource("fortios:index/routerIsis:RouterIsis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterIsis resources.
type routerIsisState struct {
	AdjacencyCheck       *string                     `pulumi:"adjacencyCheck"`
	AdjacencyCheck6      *string                     `pulumi:"adjacencyCheck6"`
	AdvPassiveOnly       *string                     `pulumi:"advPassiveOnly"`
	AdvPassiveOnly6      *string                     `pulumi:"advPassiveOnly6"`
	AuthKeychainL1       *string                     `pulumi:"authKeychainL1"`
	AuthKeychainL2       *string                     `pulumi:"authKeychainL2"`
	AuthModeL1           *string                     `pulumi:"authModeL1"`
	AuthModeL2           *string                     `pulumi:"authModeL2"`
	AuthPasswordL1       *string                     `pulumi:"authPasswordL1"`
	AuthPasswordL2       *string                     `pulumi:"authPasswordL2"`
	AuthSendonlyL1       *string                     `pulumi:"authSendonlyL1"`
	AuthSendonlyL2       *string                     `pulumi:"authSendonlyL2"`
	DefaultOriginate     *string                     `pulumi:"defaultOriginate"`
	DefaultOriginate6    *string                     `pulumi:"defaultOriginate6"`
	DynamicHostname      *string                     `pulumi:"dynamicHostname"`
	DynamicSortSubtable  *string                     `pulumi:"dynamicSortSubtable"`
	GetAllTables         *string                     `pulumi:"getAllTables"`
	IgnoreLspErrors      *string                     `pulumi:"ignoreLspErrors"`
	IsType               *string                     `pulumi:"isType"`
	IsisInterfaces       []RouterIsisIsisInterface   `pulumi:"isisInterfaces"`
	IsisNets             []RouterIsisIsisNet         `pulumi:"isisNets"`
	LspGenIntervalL1     *int                        `pulumi:"lspGenIntervalL1"`
	LspGenIntervalL2     *int                        `pulumi:"lspGenIntervalL2"`
	LspRefreshInterval   *int                        `pulumi:"lspRefreshInterval"`
	MaxLspLifetime       *int                        `pulumi:"maxLspLifetime"`
	MetricStyle          *string                     `pulumi:"metricStyle"`
	OverloadBit          *string                     `pulumi:"overloadBit"`
	OverloadBitOnStartup *int                        `pulumi:"overloadBitOnStartup"`
	OverloadBitSuppress  *string                     `pulumi:"overloadBitSuppress"`
	Redistribute6L1      *string                     `pulumi:"redistribute6L1"`
	Redistribute6L1List  *string                     `pulumi:"redistribute6L1List"`
	Redistribute6L2      *string                     `pulumi:"redistribute6L2"`
	Redistribute6L2List  *string                     `pulumi:"redistribute6L2List"`
	Redistribute6s       []RouterIsisRedistribute6   `pulumi:"redistribute6s"`
	RedistributeL1       *string                     `pulumi:"redistributeL1"`
	RedistributeL1List   *string                     `pulumi:"redistributeL1List"`
	RedistributeL2       *string                     `pulumi:"redistributeL2"`
	RedistributeL2List   *string                     `pulumi:"redistributeL2List"`
	Redistributes        []RouterIsisRedistribute    `pulumi:"redistributes"`
	SpfIntervalExpL1     *string                     `pulumi:"spfIntervalExpL1"`
	SpfIntervalExpL2     *string                     `pulumi:"spfIntervalExpL2"`
	SummaryAddress6s     []RouterIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	SummaryAddresses     []RouterIsisSummaryAddress  `pulumi:"summaryAddresses"`
	Vdomparam            *string                     `pulumi:"vdomparam"`
}

type RouterIsisState struct {
	AdjacencyCheck       pulumi.StringPtrInput
	AdjacencyCheck6      pulumi.StringPtrInput
	AdvPassiveOnly       pulumi.StringPtrInput
	AdvPassiveOnly6      pulumi.StringPtrInput
	AuthKeychainL1       pulumi.StringPtrInput
	AuthKeychainL2       pulumi.StringPtrInput
	AuthModeL1           pulumi.StringPtrInput
	AuthModeL2           pulumi.StringPtrInput
	AuthPasswordL1       pulumi.StringPtrInput
	AuthPasswordL2       pulumi.StringPtrInput
	AuthSendonlyL1       pulumi.StringPtrInput
	AuthSendonlyL2       pulumi.StringPtrInput
	DefaultOriginate     pulumi.StringPtrInput
	DefaultOriginate6    pulumi.StringPtrInput
	DynamicHostname      pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	GetAllTables         pulumi.StringPtrInput
	IgnoreLspErrors      pulumi.StringPtrInput
	IsType               pulumi.StringPtrInput
	IsisInterfaces       RouterIsisIsisInterfaceArrayInput
	IsisNets             RouterIsisIsisNetArrayInput
	LspGenIntervalL1     pulumi.IntPtrInput
	LspGenIntervalL2     pulumi.IntPtrInput
	LspRefreshInterval   pulumi.IntPtrInput
	MaxLspLifetime       pulumi.IntPtrInput
	MetricStyle          pulumi.StringPtrInput
	OverloadBit          pulumi.StringPtrInput
	OverloadBitOnStartup pulumi.IntPtrInput
	OverloadBitSuppress  pulumi.StringPtrInput
	Redistribute6L1      pulumi.StringPtrInput
	Redistribute6L1List  pulumi.StringPtrInput
	Redistribute6L2      pulumi.StringPtrInput
	Redistribute6L2List  pulumi.StringPtrInput
	Redistribute6s       RouterIsisRedistribute6ArrayInput
	RedistributeL1       pulumi.StringPtrInput
	RedistributeL1List   pulumi.StringPtrInput
	RedistributeL2       pulumi.StringPtrInput
	RedistributeL2List   pulumi.StringPtrInput
	Redistributes        RouterIsisRedistributeArrayInput
	SpfIntervalExpL1     pulumi.StringPtrInput
	SpfIntervalExpL2     pulumi.StringPtrInput
	SummaryAddress6s     RouterIsisSummaryAddress6ArrayInput
	SummaryAddresses     RouterIsisSummaryAddressArrayInput
	Vdomparam            pulumi.StringPtrInput
}

func (RouterIsisState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerIsisState)(nil)).Elem()
}

type routerIsisArgs struct {
	AdjacencyCheck       *string                     `pulumi:"adjacencyCheck"`
	AdjacencyCheck6      *string                     `pulumi:"adjacencyCheck6"`
	AdvPassiveOnly       *string                     `pulumi:"advPassiveOnly"`
	AdvPassiveOnly6      *string                     `pulumi:"advPassiveOnly6"`
	AuthKeychainL1       *string                     `pulumi:"authKeychainL1"`
	AuthKeychainL2       *string                     `pulumi:"authKeychainL2"`
	AuthModeL1           *string                     `pulumi:"authModeL1"`
	AuthModeL2           *string                     `pulumi:"authModeL2"`
	AuthPasswordL1       *string                     `pulumi:"authPasswordL1"`
	AuthPasswordL2       *string                     `pulumi:"authPasswordL2"`
	AuthSendonlyL1       *string                     `pulumi:"authSendonlyL1"`
	AuthSendonlyL2       *string                     `pulumi:"authSendonlyL2"`
	DefaultOriginate     *string                     `pulumi:"defaultOriginate"`
	DefaultOriginate6    *string                     `pulumi:"defaultOriginate6"`
	DynamicHostname      *string                     `pulumi:"dynamicHostname"`
	DynamicSortSubtable  *string                     `pulumi:"dynamicSortSubtable"`
	GetAllTables         *string                     `pulumi:"getAllTables"`
	IgnoreLspErrors      *string                     `pulumi:"ignoreLspErrors"`
	IsType               *string                     `pulumi:"isType"`
	IsisInterfaces       []RouterIsisIsisInterface   `pulumi:"isisInterfaces"`
	IsisNets             []RouterIsisIsisNet         `pulumi:"isisNets"`
	LspGenIntervalL1     *int                        `pulumi:"lspGenIntervalL1"`
	LspGenIntervalL2     *int                        `pulumi:"lspGenIntervalL2"`
	LspRefreshInterval   *int                        `pulumi:"lspRefreshInterval"`
	MaxLspLifetime       *int                        `pulumi:"maxLspLifetime"`
	MetricStyle          *string                     `pulumi:"metricStyle"`
	OverloadBit          *string                     `pulumi:"overloadBit"`
	OverloadBitOnStartup *int                        `pulumi:"overloadBitOnStartup"`
	OverloadBitSuppress  *string                     `pulumi:"overloadBitSuppress"`
	Redistribute6L1      *string                     `pulumi:"redistribute6L1"`
	Redistribute6L1List  *string                     `pulumi:"redistribute6L1List"`
	Redistribute6L2      *string                     `pulumi:"redistribute6L2"`
	Redistribute6L2List  *string                     `pulumi:"redistribute6L2List"`
	Redistribute6s       []RouterIsisRedistribute6   `pulumi:"redistribute6s"`
	RedistributeL1       *string                     `pulumi:"redistributeL1"`
	RedistributeL1List   *string                     `pulumi:"redistributeL1List"`
	RedistributeL2       *string                     `pulumi:"redistributeL2"`
	RedistributeL2List   *string                     `pulumi:"redistributeL2List"`
	Redistributes        []RouterIsisRedistribute    `pulumi:"redistributes"`
	SpfIntervalExpL1     *string                     `pulumi:"spfIntervalExpL1"`
	SpfIntervalExpL2     *string                     `pulumi:"spfIntervalExpL2"`
	SummaryAddress6s     []RouterIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	SummaryAddresses     []RouterIsisSummaryAddress  `pulumi:"summaryAddresses"`
	Vdomparam            *string                     `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterIsis resource.
type RouterIsisArgs struct {
	AdjacencyCheck       pulumi.StringPtrInput
	AdjacencyCheck6      pulumi.StringPtrInput
	AdvPassiveOnly       pulumi.StringPtrInput
	AdvPassiveOnly6      pulumi.StringPtrInput
	AuthKeychainL1       pulumi.StringPtrInput
	AuthKeychainL2       pulumi.StringPtrInput
	AuthModeL1           pulumi.StringPtrInput
	AuthModeL2           pulumi.StringPtrInput
	AuthPasswordL1       pulumi.StringPtrInput
	AuthPasswordL2       pulumi.StringPtrInput
	AuthSendonlyL1       pulumi.StringPtrInput
	AuthSendonlyL2       pulumi.StringPtrInput
	DefaultOriginate     pulumi.StringPtrInput
	DefaultOriginate6    pulumi.StringPtrInput
	DynamicHostname      pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	GetAllTables         pulumi.StringPtrInput
	IgnoreLspErrors      pulumi.StringPtrInput
	IsType               pulumi.StringPtrInput
	IsisInterfaces       RouterIsisIsisInterfaceArrayInput
	IsisNets             RouterIsisIsisNetArrayInput
	LspGenIntervalL1     pulumi.IntPtrInput
	LspGenIntervalL2     pulumi.IntPtrInput
	LspRefreshInterval   pulumi.IntPtrInput
	MaxLspLifetime       pulumi.IntPtrInput
	MetricStyle          pulumi.StringPtrInput
	OverloadBit          pulumi.StringPtrInput
	OverloadBitOnStartup pulumi.IntPtrInput
	OverloadBitSuppress  pulumi.StringPtrInput
	Redistribute6L1      pulumi.StringPtrInput
	Redistribute6L1List  pulumi.StringPtrInput
	Redistribute6L2      pulumi.StringPtrInput
	Redistribute6L2List  pulumi.StringPtrInput
	Redistribute6s       RouterIsisRedistribute6ArrayInput
	RedistributeL1       pulumi.StringPtrInput
	RedistributeL1List   pulumi.StringPtrInput
	RedistributeL2       pulumi.StringPtrInput
	RedistributeL2List   pulumi.StringPtrInput
	Redistributes        RouterIsisRedistributeArrayInput
	SpfIntervalExpL1     pulumi.StringPtrInput
	SpfIntervalExpL2     pulumi.StringPtrInput
	SummaryAddress6s     RouterIsisSummaryAddress6ArrayInput
	SummaryAddresses     RouterIsisSummaryAddressArrayInput
	Vdomparam            pulumi.StringPtrInput
}

func (RouterIsisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerIsisArgs)(nil)).Elem()
}

type RouterIsisInput interface {
	pulumi.Input

	ToRouterIsisOutput() RouterIsisOutput
	ToRouterIsisOutputWithContext(ctx context.Context) RouterIsisOutput
}

func (*RouterIsis) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterIsis)(nil)).Elem()
}

func (i *RouterIsis) ToRouterIsisOutput() RouterIsisOutput {
	return i.ToRouterIsisOutputWithContext(context.Background())
}

func (i *RouterIsis) ToRouterIsisOutputWithContext(ctx context.Context) RouterIsisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterIsisOutput)
}

func (i *RouterIsis) ToOutput(ctx context.Context) pulumix.Output[*RouterIsis] {
	return pulumix.Output[*RouterIsis]{
		OutputState: i.ToRouterIsisOutputWithContext(ctx).OutputState,
	}
}

// RouterIsisArrayInput is an input type that accepts RouterIsisArray and RouterIsisArrayOutput values.
// You can construct a concrete instance of `RouterIsisArrayInput` via:
//
//	RouterIsisArray{ RouterIsisArgs{...} }
type RouterIsisArrayInput interface {
	pulumi.Input

	ToRouterIsisArrayOutput() RouterIsisArrayOutput
	ToRouterIsisArrayOutputWithContext(context.Context) RouterIsisArrayOutput
}

type RouterIsisArray []RouterIsisInput

func (RouterIsisArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterIsis)(nil)).Elem()
}

func (i RouterIsisArray) ToRouterIsisArrayOutput() RouterIsisArrayOutput {
	return i.ToRouterIsisArrayOutputWithContext(context.Background())
}

func (i RouterIsisArray) ToRouterIsisArrayOutputWithContext(ctx context.Context) RouterIsisArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterIsisArrayOutput)
}

func (i RouterIsisArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouterIsis] {
	return pulumix.Output[[]*RouterIsis]{
		OutputState: i.ToRouterIsisArrayOutputWithContext(ctx).OutputState,
	}
}

// RouterIsisMapInput is an input type that accepts RouterIsisMap and RouterIsisMapOutput values.
// You can construct a concrete instance of `RouterIsisMapInput` via:
//
//	RouterIsisMap{ "key": RouterIsisArgs{...} }
type RouterIsisMapInput interface {
	pulumi.Input

	ToRouterIsisMapOutput() RouterIsisMapOutput
	ToRouterIsisMapOutputWithContext(context.Context) RouterIsisMapOutput
}

type RouterIsisMap map[string]RouterIsisInput

func (RouterIsisMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterIsis)(nil)).Elem()
}

func (i RouterIsisMap) ToRouterIsisMapOutput() RouterIsisMapOutput {
	return i.ToRouterIsisMapOutputWithContext(context.Background())
}

func (i RouterIsisMap) ToRouterIsisMapOutputWithContext(ctx context.Context) RouterIsisMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterIsisMapOutput)
}

func (i RouterIsisMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterIsis] {
	return pulumix.Output[map[string]*RouterIsis]{
		OutputState: i.ToRouterIsisMapOutputWithContext(ctx).OutputState,
	}
}

type RouterIsisOutput struct{ *pulumi.OutputState }

func (RouterIsisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterIsis)(nil)).Elem()
}

func (o RouterIsisOutput) ToRouterIsisOutput() RouterIsisOutput {
	return o
}

func (o RouterIsisOutput) ToRouterIsisOutputWithContext(ctx context.Context) RouterIsisOutput {
	return o
}

func (o RouterIsisOutput) ToOutput(ctx context.Context) pulumix.Output[*RouterIsis] {
	return pulumix.Output[*RouterIsis]{
		OutputState: o.OutputState,
	}
}

func (o RouterIsisOutput) AdjacencyCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AdjacencyCheck }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AdjacencyCheck6() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AdjacencyCheck6 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AdvPassiveOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AdvPassiveOnly }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AdvPassiveOnly6() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AdvPassiveOnly6 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthKeychainL1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthKeychainL1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthKeychainL2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthKeychainL2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthModeL1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthModeL1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthModeL2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthModeL2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthPasswordL1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringPtrOutput { return v.AuthPasswordL1 }).(pulumi.StringPtrOutput)
}

func (o RouterIsisOutput) AuthPasswordL2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringPtrOutput { return v.AuthPasswordL2 }).(pulumi.StringPtrOutput)
}

func (o RouterIsisOutput) AuthSendonlyL1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthSendonlyL1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) AuthSendonlyL2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.AuthSendonlyL2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) DefaultOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.DefaultOriginate }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) DefaultOriginate6() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.DefaultOriginate6 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) DynamicHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.DynamicHostname }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterIsisOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterIsisOutput) IgnoreLspErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.IgnoreLspErrors }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) IsType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.IsType }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) IsisInterfaces() RouterIsisIsisInterfaceArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisIsisInterfaceArrayOutput { return v.IsisInterfaces }).(RouterIsisIsisInterfaceArrayOutput)
}

func (o RouterIsisOutput) IsisNets() RouterIsisIsisNetArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisIsisNetArrayOutput { return v.IsisNets }).(RouterIsisIsisNetArrayOutput)
}

func (o RouterIsisOutput) LspGenIntervalL1() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.IntOutput { return v.LspGenIntervalL1 }).(pulumi.IntOutput)
}

func (o RouterIsisOutput) LspGenIntervalL2() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.IntOutput { return v.LspGenIntervalL2 }).(pulumi.IntOutput)
}

func (o RouterIsisOutput) LspRefreshInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.IntOutput { return v.LspRefreshInterval }).(pulumi.IntOutput)
}

func (o RouterIsisOutput) MaxLspLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.IntOutput { return v.MaxLspLifetime }).(pulumi.IntOutput)
}

func (o RouterIsisOutput) MetricStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.MetricStyle }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) OverloadBit() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.OverloadBit }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) OverloadBitOnStartup() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.IntOutput { return v.OverloadBitOnStartup }).(pulumi.IntOutput)
}

func (o RouterIsisOutput) OverloadBitSuppress() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.OverloadBitSuppress }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistribute6L1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.Redistribute6L1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistribute6L1List() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.Redistribute6L1List }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistribute6L2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.Redistribute6L2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistribute6L2List() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.Redistribute6L2List }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistribute6s() RouterIsisRedistribute6ArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisRedistribute6ArrayOutput { return v.Redistribute6s }).(RouterIsisRedistribute6ArrayOutput)
}

func (o RouterIsisOutput) RedistributeL1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.RedistributeL1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) RedistributeL1List() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.RedistributeL1List }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) RedistributeL2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.RedistributeL2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) RedistributeL2List() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.RedistributeL2List }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) Redistributes() RouterIsisRedistributeArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisRedistributeArrayOutput { return v.Redistributes }).(RouterIsisRedistributeArrayOutput)
}

func (o RouterIsisOutput) SpfIntervalExpL1() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.SpfIntervalExpL1 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) SpfIntervalExpL2() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringOutput { return v.SpfIntervalExpL2 }).(pulumi.StringOutput)
}

func (o RouterIsisOutput) SummaryAddress6s() RouterIsisSummaryAddress6ArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisSummaryAddress6ArrayOutput { return v.SummaryAddress6s }).(RouterIsisSummaryAddress6ArrayOutput)
}

func (o RouterIsisOutput) SummaryAddresses() RouterIsisSummaryAddressArrayOutput {
	return o.ApplyT(func(v *RouterIsis) RouterIsisSummaryAddressArrayOutput { return v.SummaryAddresses }).(RouterIsisSummaryAddressArrayOutput)
}

func (o RouterIsisOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterIsis) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterIsisArrayOutput struct{ *pulumi.OutputState }

func (RouterIsisArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterIsis)(nil)).Elem()
}

func (o RouterIsisArrayOutput) ToRouterIsisArrayOutput() RouterIsisArrayOutput {
	return o
}

func (o RouterIsisArrayOutput) ToRouterIsisArrayOutputWithContext(ctx context.Context) RouterIsisArrayOutput {
	return o
}

func (o RouterIsisArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouterIsis] {
	return pulumix.Output[[]*RouterIsis]{
		OutputState: o.OutputState,
	}
}

func (o RouterIsisArrayOutput) Index(i pulumi.IntInput) RouterIsisOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterIsis {
		return vs[0].([]*RouterIsis)[vs[1].(int)]
	}).(RouterIsisOutput)
}

type RouterIsisMapOutput struct{ *pulumi.OutputState }

func (RouterIsisMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterIsis)(nil)).Elem()
}

func (o RouterIsisMapOutput) ToRouterIsisMapOutput() RouterIsisMapOutput {
	return o
}

func (o RouterIsisMapOutput) ToRouterIsisMapOutputWithContext(ctx context.Context) RouterIsisMapOutput {
	return o
}

func (o RouterIsisMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterIsis] {
	return pulumix.Output[map[string]*RouterIsis]{
		OutputState: o.OutputState,
	}
}

func (o RouterIsisMapOutput) MapIndex(k pulumi.StringInput) RouterIsisOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterIsis {
		return vs[0].(map[string]*RouterIsis)[vs[1].(string)]
	}).(RouterIsisOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterIsisInput)(nil)).Elem(), &RouterIsis{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterIsisArrayInput)(nil)).Elem(), RouterIsisArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterIsisMapInput)(nil)).Elem(), RouterIsisMap{})
	pulumi.RegisterOutputType(RouterIsisOutput{})
	pulumi.RegisterOutputType(RouterIsisArrayOutput{})
	pulumi.RegisterOutputType(RouterIsisMapOutput{})
}

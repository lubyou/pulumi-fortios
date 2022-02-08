// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IS-IS.
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
// 		_, err := fortios.NewRouterIsis(ctx, "trname", &fortios.RouterIsisArgs{
// 			AdjacencyCheck:     pulumi.String("disable"),
// 			AdjacencyCheck6:    pulumi.String("disable"),
// 			AdvPassiveOnly:     pulumi.String("disable"),
// 			AdvPassiveOnly6:    pulumi.String("disable"),
// 			AuthModeL1:         pulumi.String("password"),
// 			AuthModeL2:         pulumi.String("password"),
// 			AuthSendonlyL1:     pulumi.String("disable"),
// 			AuthSendonlyL2:     pulumi.String("disable"),
// 			DefaultOriginate:   pulumi.String("disable"),
// 			DefaultOriginate6:  pulumi.String("disable"),
// 			DynamicHostname:    pulumi.String("disable"),
// 			IgnoreLspErrors:    pulumi.String("disable"),
// 			IsType:             pulumi.String("level-1-2"),
// 			LspGenIntervalL1:   pulumi.Int(30),
// 			LspGenIntervalL2:   pulumi.Int(30),
// 			LspRefreshInterval: pulumi.Int(900),
// 			MaxLspLifetime:     pulumi.Int(1200),
// 			MetricStyle:        pulumi.String("narrow"),
// 			OverloadBit:        pulumi.String("disable"),
// 			Redistribute6L1:    pulumi.String("disable"),
// 			Redistribute6L2:    pulumi.String("disable"),
// 			RedistributeL1:     pulumi.String("disable"),
// 			RedistributeL2:     pulumi.String("disable"),
// 			SpfIntervalExpL1:   pulumi.String("500 50000"),
// 			SpfIntervalExpL2:   pulumi.String("500 50000"),
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
// Router Isis can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerIsis:RouterIsis labelname RouterIsis
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterIsis struct {
	pulumi.CustomResourceState

	// Enable/disable adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck pulumi.StringOutput `pulumi:"adjacencyCheck"`
	// Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck6 pulumi.StringOutput `pulumi:"adjacencyCheck6"`
	// Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly pulumi.StringOutput `pulumi:"advPassiveOnly"`
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly6 pulumi.StringOutput `pulumi:"advPassiveOnly6"`
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 pulumi.StringOutput `pulumi:"authKeychainL1"`
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 pulumi.StringOutput `pulumi:"authKeychainL2"`
	// Level 1 authentication mode. Valid values: `md5`, `password`.
	AuthModeL1 pulumi.StringOutput `pulumi:"authModeL1"`
	// Level 2 authentication mode. Valid values: `md5`, `password`.
	AuthModeL2 pulumi.StringOutput `pulumi:"authModeL2"`
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 pulumi.StringPtrOutput `pulumi:"authPasswordL1"`
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 pulumi.StringPtrOutput `pulumi:"authPasswordL2"`
	// Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL1 pulumi.StringOutput `pulumi:"authSendonlyL1"`
	// Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL2 pulumi.StringOutput `pulumi:"authSendonlyL2"`
	// Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
	DefaultOriginate pulumi.StringOutput `pulumi:"defaultOriginate"`
	// Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
	DefaultOriginate6 pulumi.StringOutput `pulumi:"defaultOriginate6"`
	// Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
	DynamicHostname pulumi.StringOutput `pulumi:"dynamicHostname"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
	IgnoreLspErrors pulumi.StringOutput `pulumi:"ignoreLspErrors"`
	// IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
	IsType pulumi.StringOutput `pulumi:"isType"`
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces RouterIsisIsisInterfaceArrayOutput `pulumi:"isisInterfaces"`
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets RouterIsisIsisNetArrayOutput `pulumi:"isisNets"`
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 pulumi.IntOutput `pulumi:"lspGenIntervalL1"`
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 pulumi.IntOutput `pulumi:"lspGenIntervalL2"`
	// LSP refresh time in seconds.
	LspRefreshInterval pulumi.IntOutput `pulumi:"lspRefreshInterval"`
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime pulumi.IntOutput `pulumi:"maxLspLifetime"`
	// Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
	MetricStyle pulumi.StringOutput `pulumi:"metricStyle"`
	// Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
	OverloadBit pulumi.StringOutput `pulumi:"overloadBit"`
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup pulumi.IntOutput `pulumi:"overloadBitOnStartup"`
	// Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
	OverloadBitSuppress pulumi.StringOutput `pulumi:"overloadBitSuppress"`
	// Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
	Redistribute6L1 pulumi.StringOutput `pulumi:"redistribute6L1"`
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List pulumi.StringOutput `pulumi:"redistribute6L1List"`
	// Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
	Redistribute6L2 pulumi.StringOutput `pulumi:"redistribute6L2"`
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List pulumi.StringOutput `pulumi:"redistribute6L2List"`
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s RouterIsisRedistribute6ArrayOutput `pulumi:"redistribute6s"`
	// Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
	RedistributeL1 pulumi.StringOutput `pulumi:"redistributeL1"`
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List pulumi.StringOutput `pulumi:"redistributeL1List"`
	// Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
	RedistributeL2 pulumi.StringOutput `pulumi:"redistributeL2"`
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List pulumi.StringOutput `pulumi:"redistributeL2List"`
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes RouterIsisRedistributeArrayOutput `pulumi:"redistributes"`
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 pulumi.StringOutput `pulumi:"spfIntervalExpL1"`
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 pulumi.StringOutput `pulumi:"spfIntervalExpL2"`
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s RouterIsisSummaryAddress6ArrayOutput `pulumi:"summaryAddress6s"`
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterIsisSummaryAddressArrayOutput `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterIsis registers a new resource with the given unique name, arguments, and options.
func NewRouterIsis(ctx *pulumi.Context,
	name string, args *RouterIsisArgs, opts ...pulumi.ResourceOption) (*RouterIsis, error) {
	if args == nil {
		args = &RouterIsisArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
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
	// Enable/disable adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck *string `pulumi:"adjacencyCheck"`
	// Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck6 *string `pulumi:"adjacencyCheck6"`
	// Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly *string `pulumi:"advPassiveOnly"`
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly6 *string `pulumi:"advPassiveOnly6"`
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 *string `pulumi:"authKeychainL1"`
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 *string `pulumi:"authKeychainL2"`
	// Level 1 authentication mode. Valid values: `md5`, `password`.
	AuthModeL1 *string `pulumi:"authModeL1"`
	// Level 2 authentication mode. Valid values: `md5`, `password`.
	AuthModeL2 *string `pulumi:"authModeL2"`
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 *string `pulumi:"authPasswordL1"`
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 *string `pulumi:"authPasswordL2"`
	// Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL1 *string `pulumi:"authSendonlyL1"`
	// Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL2 *string `pulumi:"authSendonlyL2"`
	// Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
	DefaultOriginate *string `pulumi:"defaultOriginate"`
	// Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
	DefaultOriginate6 *string `pulumi:"defaultOriginate6"`
	// Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
	DynamicHostname *string `pulumi:"dynamicHostname"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
	IgnoreLspErrors *string `pulumi:"ignoreLspErrors"`
	// IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
	IsType *string `pulumi:"isType"`
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces []RouterIsisIsisInterface `pulumi:"isisInterfaces"`
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets []RouterIsisIsisNet `pulumi:"isisNets"`
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 *int `pulumi:"lspGenIntervalL1"`
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 *int `pulumi:"lspGenIntervalL2"`
	// LSP refresh time in seconds.
	LspRefreshInterval *int `pulumi:"lspRefreshInterval"`
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime *int `pulumi:"maxLspLifetime"`
	// Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
	MetricStyle *string `pulumi:"metricStyle"`
	// Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
	OverloadBit *string `pulumi:"overloadBit"`
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup *int `pulumi:"overloadBitOnStartup"`
	// Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
	OverloadBitSuppress *string `pulumi:"overloadBitSuppress"`
	// Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
	Redistribute6L1 *string `pulumi:"redistribute6L1"`
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List *string `pulumi:"redistribute6L1List"`
	// Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
	Redistribute6L2 *string `pulumi:"redistribute6L2"`
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List *string `pulumi:"redistribute6L2List"`
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s []RouterIsisRedistribute6 `pulumi:"redistribute6s"`
	// Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
	RedistributeL1 *string `pulumi:"redistributeL1"`
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List *string `pulumi:"redistributeL1List"`
	// Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
	RedistributeL2 *string `pulumi:"redistributeL2"`
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List *string `pulumi:"redistributeL2List"`
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes []RouterIsisRedistribute `pulumi:"redistributes"`
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 *string `pulumi:"spfIntervalExpL1"`
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 *string `pulumi:"spfIntervalExpL2"`
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s []RouterIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []RouterIsisSummaryAddress `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterIsisState struct {
	// Enable/disable adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck pulumi.StringPtrInput
	// Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck6 pulumi.StringPtrInput
	// Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly pulumi.StringPtrInput
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly6 pulumi.StringPtrInput
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 pulumi.StringPtrInput
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 pulumi.StringPtrInput
	// Level 1 authentication mode. Valid values: `md5`, `password`.
	AuthModeL1 pulumi.StringPtrInput
	// Level 2 authentication mode. Valid values: `md5`, `password`.
	AuthModeL2 pulumi.StringPtrInput
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 pulumi.StringPtrInput
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 pulumi.StringPtrInput
	// Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL1 pulumi.StringPtrInput
	// Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL2 pulumi.StringPtrInput
	// Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
	DefaultOriginate pulumi.StringPtrInput
	// Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
	DefaultOriginate6 pulumi.StringPtrInput
	// Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
	DynamicHostname pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
	IgnoreLspErrors pulumi.StringPtrInput
	// IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
	IsType pulumi.StringPtrInput
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces RouterIsisIsisInterfaceArrayInput
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets RouterIsisIsisNetArrayInput
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 pulumi.IntPtrInput
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 pulumi.IntPtrInput
	// LSP refresh time in seconds.
	LspRefreshInterval pulumi.IntPtrInput
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime pulumi.IntPtrInput
	// Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
	MetricStyle pulumi.StringPtrInput
	// Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
	OverloadBit pulumi.StringPtrInput
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup pulumi.IntPtrInput
	// Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
	OverloadBitSuppress pulumi.StringPtrInput
	// Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
	Redistribute6L1 pulumi.StringPtrInput
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List pulumi.StringPtrInput
	// Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
	Redistribute6L2 pulumi.StringPtrInput
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List pulumi.StringPtrInput
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s RouterIsisRedistribute6ArrayInput
	// Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
	RedistributeL1 pulumi.StringPtrInput
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List pulumi.StringPtrInput
	// Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
	RedistributeL2 pulumi.StringPtrInput
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List pulumi.StringPtrInput
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes RouterIsisRedistributeArrayInput
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 pulumi.StringPtrInput
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 pulumi.StringPtrInput
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s RouterIsisSummaryAddress6ArrayInput
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterIsisSummaryAddressArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterIsisState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerIsisState)(nil)).Elem()
}

type routerIsisArgs struct {
	// Enable/disable adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck *string `pulumi:"adjacencyCheck"`
	// Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck6 *string `pulumi:"adjacencyCheck6"`
	// Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly *string `pulumi:"advPassiveOnly"`
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly6 *string `pulumi:"advPassiveOnly6"`
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 *string `pulumi:"authKeychainL1"`
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 *string `pulumi:"authKeychainL2"`
	// Level 1 authentication mode. Valid values: `md5`, `password`.
	AuthModeL1 *string `pulumi:"authModeL1"`
	// Level 2 authentication mode. Valid values: `md5`, `password`.
	AuthModeL2 *string `pulumi:"authModeL2"`
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 *string `pulumi:"authPasswordL1"`
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 *string `pulumi:"authPasswordL2"`
	// Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL1 *string `pulumi:"authSendonlyL1"`
	// Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL2 *string `pulumi:"authSendonlyL2"`
	// Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
	DefaultOriginate *string `pulumi:"defaultOriginate"`
	// Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
	DefaultOriginate6 *string `pulumi:"defaultOriginate6"`
	// Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
	DynamicHostname *string `pulumi:"dynamicHostname"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
	IgnoreLspErrors *string `pulumi:"ignoreLspErrors"`
	// IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
	IsType *string `pulumi:"isType"`
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces []RouterIsisIsisInterface `pulumi:"isisInterfaces"`
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets []RouterIsisIsisNet `pulumi:"isisNets"`
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 *int `pulumi:"lspGenIntervalL1"`
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 *int `pulumi:"lspGenIntervalL2"`
	// LSP refresh time in seconds.
	LspRefreshInterval *int `pulumi:"lspRefreshInterval"`
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime *int `pulumi:"maxLspLifetime"`
	// Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
	MetricStyle *string `pulumi:"metricStyle"`
	// Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
	OverloadBit *string `pulumi:"overloadBit"`
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup *int `pulumi:"overloadBitOnStartup"`
	// Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
	OverloadBitSuppress *string `pulumi:"overloadBitSuppress"`
	// Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
	Redistribute6L1 *string `pulumi:"redistribute6L1"`
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List *string `pulumi:"redistribute6L1List"`
	// Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
	Redistribute6L2 *string `pulumi:"redistribute6L2"`
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List *string `pulumi:"redistribute6L2List"`
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s []RouterIsisRedistribute6 `pulumi:"redistribute6s"`
	// Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
	RedistributeL1 *string `pulumi:"redistributeL1"`
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List *string `pulumi:"redistributeL1List"`
	// Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
	RedistributeL2 *string `pulumi:"redistributeL2"`
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List *string `pulumi:"redistributeL2List"`
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes []RouterIsisRedistribute `pulumi:"redistributes"`
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 *string `pulumi:"spfIntervalExpL1"`
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 *string `pulumi:"spfIntervalExpL2"`
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s []RouterIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []RouterIsisSummaryAddress `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterIsis resource.
type RouterIsisArgs struct {
	// Enable/disable adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck pulumi.StringPtrInput
	// Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
	AdjacencyCheck6 pulumi.StringPtrInput
	// Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly pulumi.StringPtrInput
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
	AdvPassiveOnly6 pulumi.StringPtrInput
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 pulumi.StringPtrInput
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 pulumi.StringPtrInput
	// Level 1 authentication mode. Valid values: `md5`, `password`.
	AuthModeL1 pulumi.StringPtrInput
	// Level 2 authentication mode. Valid values: `md5`, `password`.
	AuthModeL2 pulumi.StringPtrInput
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 pulumi.StringPtrInput
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 pulumi.StringPtrInput
	// Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL1 pulumi.StringPtrInput
	// Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
	AuthSendonlyL2 pulumi.StringPtrInput
	// Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
	DefaultOriginate pulumi.StringPtrInput
	// Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
	DefaultOriginate6 pulumi.StringPtrInput
	// Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
	DynamicHostname pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
	IgnoreLspErrors pulumi.StringPtrInput
	// IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
	IsType pulumi.StringPtrInput
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces RouterIsisIsisInterfaceArrayInput
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets RouterIsisIsisNetArrayInput
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 pulumi.IntPtrInput
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 pulumi.IntPtrInput
	// LSP refresh time in seconds.
	LspRefreshInterval pulumi.IntPtrInput
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime pulumi.IntPtrInput
	// Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
	MetricStyle pulumi.StringPtrInput
	// Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
	OverloadBit pulumi.StringPtrInput
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup pulumi.IntPtrInput
	// Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
	OverloadBitSuppress pulumi.StringPtrInput
	// Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
	Redistribute6L1 pulumi.StringPtrInput
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List pulumi.StringPtrInput
	// Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
	Redistribute6L2 pulumi.StringPtrInput
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List pulumi.StringPtrInput
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s RouterIsisRedistribute6ArrayInput
	// Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
	RedistributeL1 pulumi.StringPtrInput
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List pulumi.StringPtrInput
	// Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
	RedistributeL2 pulumi.StringPtrInput
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List pulumi.StringPtrInput
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes RouterIsisRedistributeArrayInput
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 pulumi.StringPtrInput
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 pulumi.StringPtrInput
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s RouterIsisSummaryAddress6ArrayInput
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterIsisSummaryAddressArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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

// RouterIsisArrayInput is an input type that accepts RouterIsisArray and RouterIsisArrayOutput values.
// You can construct a concrete instance of `RouterIsisArrayInput` via:
//
//          RouterIsisArray{ RouterIsisArgs{...} }
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

// RouterIsisMapInput is an input type that accepts RouterIsisMap and RouterIsisMapOutput values.
// You can construct a concrete instance of `RouterIsisMapInput` via:
//
//          RouterIsisMap{ "key": RouterIsisArgs{...} }
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

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router isis
func LookupRouterIsis(ctx *pulumi.Context, args *LookupRouterIsisArgs, opts ...pulumi.InvokeOption) (*LookupRouterIsisResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterIsisResult
	err := ctx.Invoke("fortios:index/getRouterIsis:GetRouterIsis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterIsis.
type LookupRouterIsisArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterIsis.
type LookupRouterIsisResult struct {
	// Enable/disable adjacency check.
	AdjacencyCheck string `pulumi:"adjacencyCheck"`
	// Enable/disable IPv6 adjacency check.
	AdjacencyCheck6 string `pulumi:"adjacencyCheck6"`
	// Enable/disable IS-IS advertisement of passive interfaces only.
	AdvPassiveOnly string `pulumi:"advPassiveOnly"`
	// Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
	AdvPassiveOnly6 string `pulumi:"advPassiveOnly6"`
	// Authentication key-chain for level 1 PDUs.
	AuthKeychainL1 string `pulumi:"authKeychainL1"`
	// Authentication key-chain for level 2 PDUs.
	AuthKeychainL2 string `pulumi:"authKeychainL2"`
	// Level 1 authentication mode.
	AuthModeL1 string `pulumi:"authModeL1"`
	// Level 2 authentication mode.
	AuthModeL2 string `pulumi:"authModeL2"`
	// Authentication password for level 1 PDUs.
	AuthPasswordL1 string `pulumi:"authPasswordL1"`
	// Authentication password for level 2 PDUs.
	AuthPasswordL2 string `pulumi:"authPasswordL2"`
	// Enable/disable level 1 authentication send-only.
	AuthSendonlyL1 string `pulumi:"authSendonlyL1"`
	// Enable/disable level 2 authentication send-only.
	AuthSendonlyL2 string `pulumi:"authSendonlyL2"`
	// Enable/disable distribution of default route information.
	DefaultOriginate string `pulumi:"defaultOriginate"`
	// Enable/disable distribution of default IPv6 route information.
	DefaultOriginate6 string `pulumi:"defaultOriginate6"`
	// Enable/disable dynamic hostname.
	DynamicHostname string `pulumi:"dynamicHostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable ignoring of LSP errors with bad checksums.
	IgnoreLspErrors string `pulumi:"ignoreLspErrors"`
	// IS type.
	IsType string `pulumi:"isType"`
	// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
	IsisInterfaces []GetRouterIsisIsisInterface `pulumi:"isisInterfaces"`
	// IS-IS net configuration. The structure of `isisNet` block is documented below.
	IsisNets []GetRouterIsisIsisNet `pulumi:"isisNets"`
	// Minimum interval for level 1 LSP regenerating.
	LspGenIntervalL1 int `pulumi:"lspGenIntervalL1"`
	// Minimum interval for level 2 LSP regenerating.
	LspGenIntervalL2 int `pulumi:"lspGenIntervalL2"`
	// LSP refresh time in seconds.
	LspRefreshInterval int `pulumi:"lspRefreshInterval"`
	// Maximum LSP lifetime in seconds.
	MaxLspLifetime int `pulumi:"maxLspLifetime"`
	// Use old-style (ISO 10589) or new-style packet formats
	MetricStyle string `pulumi:"metricStyle"`
	// Enable/disable signal other routers not to use us in SPF.
	OverloadBit string `pulumi:"overloadBit"`
	// Overload-bit only temporarily after reboot.
	OverloadBitOnStartup int `pulumi:"overloadBitOnStartup"`
	// Suppress overload-bit for the specific prefixes.
	OverloadBitSuppress string `pulumi:"overloadBitSuppress"`
	// Enable/disable redistribution of level 1 IPv6 routes into level 2.
	Redistribute6L1 string `pulumi:"redistribute6L1"`
	// Access-list for IPv6 route redistribution from l1 to l2.
	Redistribute6L1List string `pulumi:"redistribute6L1List"`
	// Enable/disable redistribution of level 2 IPv6 routes into level 1.
	Redistribute6L2 string `pulumi:"redistribute6L2"`
	// Access-list for IPv6 route redistribution from l2 to l1.
	Redistribute6L2List string `pulumi:"redistribute6L2List"`
	// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
	Redistribute6s []GetRouterIsisRedistribute6 `pulumi:"redistribute6s"`
	// Enable/disable redistribution of level 1 routes into level 2.
	RedistributeL1 string `pulumi:"redistributeL1"`
	// Access-list for route redistribution from l1 to l2.
	RedistributeL1List string `pulumi:"redistributeL1List"`
	// Enable/disable redistribution of level 2 routes into level 1.
	RedistributeL2 string `pulumi:"redistributeL2"`
	// Access-list for route redistribution from l2 to l1.
	RedistributeL2List string `pulumi:"redistributeL2List"`
	// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
	Redistributes []GetRouterIsisRedistribute `pulumi:"redistributes"`
	// Level 1 SPF calculation delay.
	SpfIntervalExpL1 string `pulumi:"spfIntervalExpL1"`
	// Level 2 SPF calculation delay.
	SpfIntervalExpL2 string `pulumi:"spfIntervalExpL2"`
	// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
	SummaryAddress6s []GetRouterIsisSummaryAddress6 `pulumi:"summaryAddress6s"`
	// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []GetRouterIsisSummaryAddress `pulumi:"summaryAddresses"`
	Vdomparam        *string                       `pulumi:"vdomparam"`
}

func LookupRouterIsisOutput(ctx *pulumi.Context, args LookupRouterIsisOutputArgs, opts ...pulumi.InvokeOption) LookupRouterIsisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterIsisResult, error) {
			args := v.(LookupRouterIsisArgs)
			r, err := LookupRouterIsis(ctx, &args, opts...)
			var s LookupRouterIsisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterIsisResultOutput)
}

// A collection of arguments for invoking GetRouterIsis.
type LookupRouterIsisOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterIsisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterIsisArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterIsis.
type LookupRouterIsisResultOutput struct{ *pulumi.OutputState }

func (LookupRouterIsisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterIsisResult)(nil)).Elem()
}

func (o LookupRouterIsisResultOutput) ToLookupRouterIsisResultOutput() LookupRouterIsisResultOutput {
	return o
}

func (o LookupRouterIsisResultOutput) ToLookupRouterIsisResultOutputWithContext(ctx context.Context) LookupRouterIsisResultOutput {
	return o
}

// Enable/disable adjacency check.
func (o LookupRouterIsisResultOutput) AdjacencyCheck() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AdjacencyCheck }).(pulumi.StringOutput)
}

// Enable/disable IPv6 adjacency check.
func (o LookupRouterIsisResultOutput) AdjacencyCheck6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AdjacencyCheck6 }).(pulumi.StringOutput)
}

// Enable/disable IS-IS advertisement of passive interfaces only.
func (o LookupRouterIsisResultOutput) AdvPassiveOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AdvPassiveOnly }).(pulumi.StringOutput)
}

// Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
func (o LookupRouterIsisResultOutput) AdvPassiveOnly6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AdvPassiveOnly6 }).(pulumi.StringOutput)
}

// Authentication key-chain for level 1 PDUs.
func (o LookupRouterIsisResultOutput) AuthKeychainL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthKeychainL1 }).(pulumi.StringOutput)
}

// Authentication key-chain for level 2 PDUs.
func (o LookupRouterIsisResultOutput) AuthKeychainL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthKeychainL2 }).(pulumi.StringOutput)
}

// Level 1 authentication mode.
func (o LookupRouterIsisResultOutput) AuthModeL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthModeL1 }).(pulumi.StringOutput)
}

// Level 2 authentication mode.
func (o LookupRouterIsisResultOutput) AuthModeL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthModeL2 }).(pulumi.StringOutput)
}

// Authentication password for level 1 PDUs.
func (o LookupRouterIsisResultOutput) AuthPasswordL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthPasswordL1 }).(pulumi.StringOutput)
}

// Authentication password for level 2 PDUs.
func (o LookupRouterIsisResultOutput) AuthPasswordL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthPasswordL2 }).(pulumi.StringOutput)
}

// Enable/disable level 1 authentication send-only.
func (o LookupRouterIsisResultOutput) AuthSendonlyL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthSendonlyL1 }).(pulumi.StringOutput)
}

// Enable/disable level 2 authentication send-only.
func (o LookupRouterIsisResultOutput) AuthSendonlyL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.AuthSendonlyL2 }).(pulumi.StringOutput)
}

// Enable/disable distribution of default route information.
func (o LookupRouterIsisResultOutput) DefaultOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.DefaultOriginate }).(pulumi.StringOutput)
}

// Enable/disable distribution of default IPv6 route information.
func (o LookupRouterIsisResultOutput) DefaultOriginate6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.DefaultOriginate6 }).(pulumi.StringOutput)
}

// Enable/disable dynamic hostname.
func (o LookupRouterIsisResultOutput) DynamicHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.DynamicHostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterIsisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable ignoring of LSP errors with bad checksums.
func (o LookupRouterIsisResultOutput) IgnoreLspErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.IgnoreLspErrors }).(pulumi.StringOutput)
}

// IS type.
func (o LookupRouterIsisResultOutput) IsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.IsType }).(pulumi.StringOutput)
}

// IS-IS interface configuration. The structure of `isisInterface` block is documented below.
func (o LookupRouterIsisResultOutput) IsisInterfaces() GetRouterIsisIsisInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisIsisInterface { return v.IsisInterfaces }).(GetRouterIsisIsisInterfaceArrayOutput)
}

// IS-IS net configuration. The structure of `isisNet` block is documented below.
func (o LookupRouterIsisResultOutput) IsisNets() GetRouterIsisIsisNetArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisIsisNet { return v.IsisNets }).(GetRouterIsisIsisNetArrayOutput)
}

// Minimum interval for level 1 LSP regenerating.
func (o LookupRouterIsisResultOutput) LspGenIntervalL1() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) int { return v.LspGenIntervalL1 }).(pulumi.IntOutput)
}

// Minimum interval for level 2 LSP regenerating.
func (o LookupRouterIsisResultOutput) LspGenIntervalL2() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) int { return v.LspGenIntervalL2 }).(pulumi.IntOutput)
}

// LSP refresh time in seconds.
func (o LookupRouterIsisResultOutput) LspRefreshInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) int { return v.LspRefreshInterval }).(pulumi.IntOutput)
}

// Maximum LSP lifetime in seconds.
func (o LookupRouterIsisResultOutput) MaxLspLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) int { return v.MaxLspLifetime }).(pulumi.IntOutput)
}

// Use old-style (ISO 10589) or new-style packet formats
func (o LookupRouterIsisResultOutput) MetricStyle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.MetricStyle }).(pulumi.StringOutput)
}

// Enable/disable signal other routers not to use us in SPF.
func (o LookupRouterIsisResultOutput) OverloadBit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.OverloadBit }).(pulumi.StringOutput)
}

// Overload-bit only temporarily after reboot.
func (o LookupRouterIsisResultOutput) OverloadBitOnStartup() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) int { return v.OverloadBitOnStartup }).(pulumi.IntOutput)
}

// Suppress overload-bit for the specific prefixes.
func (o LookupRouterIsisResultOutput) OverloadBitSuppress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.OverloadBitSuppress }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 1 IPv6 routes into level 2.
func (o LookupRouterIsisResultOutput) Redistribute6L1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.Redistribute6L1 }).(pulumi.StringOutput)
}

// Access-list for IPv6 route redistribution from l1 to l2.
func (o LookupRouterIsisResultOutput) Redistribute6L1List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.Redistribute6L1List }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 2 IPv6 routes into level 1.
func (o LookupRouterIsisResultOutput) Redistribute6L2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.Redistribute6L2 }).(pulumi.StringOutput)
}

// Access-list for IPv6 route redistribution from l2 to l1.
func (o LookupRouterIsisResultOutput) Redistribute6L2List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.Redistribute6L2List }).(pulumi.StringOutput)
}

// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
func (o LookupRouterIsisResultOutput) Redistribute6s() GetRouterIsisRedistribute6ArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisRedistribute6 { return v.Redistribute6s }).(GetRouterIsisRedistribute6ArrayOutput)
}

// Enable/disable redistribution of level 1 routes into level 2.
func (o LookupRouterIsisResultOutput) RedistributeL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.RedistributeL1 }).(pulumi.StringOutput)
}

// Access-list for route redistribution from l1 to l2.
func (o LookupRouterIsisResultOutput) RedistributeL1List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.RedistributeL1List }).(pulumi.StringOutput)
}

// Enable/disable redistribution of level 2 routes into level 1.
func (o LookupRouterIsisResultOutput) RedistributeL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.RedistributeL2 }).(pulumi.StringOutput)
}

// Access-list for route redistribution from l2 to l1.
func (o LookupRouterIsisResultOutput) RedistributeL2List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.RedistributeL2List }).(pulumi.StringOutput)
}

// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
func (o LookupRouterIsisResultOutput) Redistributes() GetRouterIsisRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisRedistribute { return v.Redistributes }).(GetRouterIsisRedistributeArrayOutput)
}

// Level 1 SPF calculation delay.
func (o LookupRouterIsisResultOutput) SpfIntervalExpL1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.SpfIntervalExpL1 }).(pulumi.StringOutput)
}

// Level 2 SPF calculation delay.
func (o LookupRouterIsisResultOutput) SpfIntervalExpL2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) string { return v.SpfIntervalExpL2 }).(pulumi.StringOutput)
}

// IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
func (o LookupRouterIsisResultOutput) SummaryAddress6s() GetRouterIsisSummaryAddress6ArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisSummaryAddress6 { return v.SummaryAddress6s }).(GetRouterIsisSummaryAddress6ArrayOutput)
}

// IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
func (o LookupRouterIsisResultOutput) SummaryAddresses() GetRouterIsisSummaryAddressArrayOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) []GetRouterIsisSummaryAddress { return v.SummaryAddresses }).(GetRouterIsisSummaryAddressArrayOutput)
}

func (o LookupRouterIsisResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterIsisResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterIsisResultOutput{})
}

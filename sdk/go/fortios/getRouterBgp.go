// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouterBgp(ctx *pulumi.Context, args *LookupRouterBgpArgs, opts ...pulumi.InvokeOption) (*LookupRouterBgpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterBgpResult
	err := ctx.Invoke("fortios:index/getRouterBgp:GetRouterBgp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterBgp.
type LookupRouterBgpArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterBgp.
type LookupRouterBgpResult struct {
	AdditionalPath                  string                          `pulumi:"additionalPath"`
	AdditionalPath6                 string                          `pulumi:"additionalPath6"`
	AdditionalPathSelect            int                             `pulumi:"additionalPathSelect"`
	AdditionalPathSelect6           int                             `pulumi:"additionalPathSelect6"`
	AdditionalPathSelectVpnv4       int                             `pulumi:"additionalPathSelectVpnv4"`
	AdditionalPathVpnv4             string                          `pulumi:"additionalPathVpnv4"`
	AdminDistances                  []GetRouterBgpAdminDistance     `pulumi:"adminDistances"`
	AggregateAddress6s              []GetRouterBgpAggregateAddress6 `pulumi:"aggregateAddress6s"`
	AggregateAddresses              []GetRouterBgpAggregateAddress  `pulumi:"aggregateAddresses"`
	AlwaysCompareMed                string                          `pulumi:"alwaysCompareMed"`
	As                              int                             `pulumi:"as"`
	BestpathAsPathIgnore            string                          `pulumi:"bestpathAsPathIgnore"`
	BestpathCmpConfedAspath         string                          `pulumi:"bestpathCmpConfedAspath"`
	BestpathCmpRouterid             string                          `pulumi:"bestpathCmpRouterid"`
	BestpathMedConfed               string                          `pulumi:"bestpathMedConfed"`
	BestpathMedMissingAsWorst       string                          `pulumi:"bestpathMedMissingAsWorst"`
	ClientToClientReflection        string                          `pulumi:"clientToClientReflection"`
	ClusterId                       string                          `pulumi:"clusterId"`
	ConfederationIdentifier         int                             `pulumi:"confederationIdentifier"`
	ConfederationPeers              []GetRouterBgpConfederationPeer `pulumi:"confederationPeers"`
	Dampening                       string                          `pulumi:"dampening"`
	DampeningMaxSuppressTime        int                             `pulumi:"dampeningMaxSuppressTime"`
	DampeningReachabilityHalfLife   int                             `pulumi:"dampeningReachabilityHalfLife"`
	DampeningReuse                  int                             `pulumi:"dampeningReuse"`
	DampeningRouteMap               string                          `pulumi:"dampeningRouteMap"`
	DampeningSuppress               int                             `pulumi:"dampeningSuppress"`
	DampeningUnreachabilityHalfLife int                             `pulumi:"dampeningUnreachabilityHalfLife"`
	DefaultLocalPreference          int                             `pulumi:"defaultLocalPreference"`
	DeterministicMed                string                          `pulumi:"deterministicMed"`
	DistanceExternal                int                             `pulumi:"distanceExternal"`
	DistanceInternal                int                             `pulumi:"distanceInternal"`
	DistanceLocal                   int                             `pulumi:"distanceLocal"`
	EbgpMultipath                   string                          `pulumi:"ebgpMultipath"`
	EnforceFirstAs                  string                          `pulumi:"enforceFirstAs"`
	FastExternalFailover            string                          `pulumi:"fastExternalFailover"`
	GracefulEndOnTimer              string                          `pulumi:"gracefulEndOnTimer"`
	GracefulRestart                 string                          `pulumi:"gracefulRestart"`
	GracefulRestartTime             int                             `pulumi:"gracefulRestartTime"`
	GracefulStalepathTime           int                             `pulumi:"gracefulStalepathTime"`
	GracefulUpdateDelay             int                             `pulumi:"gracefulUpdateDelay"`
	HoldtimeTimer                   int                             `pulumi:"holdtimeTimer"`
	IbgpMultipath                   string                          `pulumi:"ibgpMultipath"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string                       `pulumi:"id"`
	IgnoreOptionalCapability   string                       `pulumi:"ignoreOptionalCapability"`
	KeepaliveTimer             int                          `pulumi:"keepaliveTimer"`
	LogNeighbourChanges        string                       `pulumi:"logNeighbourChanges"`
	MultipathRecursiveDistance string                       `pulumi:"multipathRecursiveDistance"`
	NeighborGroups             []GetRouterBgpNeighborGroup  `pulumi:"neighborGroups"`
	NeighborRange6s            []GetRouterBgpNeighborRange6 `pulumi:"neighborRange6s"`
	NeighborRanges             []GetRouterBgpNeighborRange  `pulumi:"neighborRanges"`
	Neighbors                  []GetRouterBgpNeighbor       `pulumi:"neighbors"`
	Network6s                  []GetRouterBgpNetwork6       `pulumi:"network6s"`
	NetworkImportCheck         string                       `pulumi:"networkImportCheck"`
	Networks                   []GetRouterBgpNetwork        `pulumi:"networks"`
	RecursiveInheritPriority   string                       `pulumi:"recursiveInheritPriority"`
	RecursiveNextHop           string                       `pulumi:"recursiveNextHop"`
	Redistribute6s             []GetRouterBgpRedistribute6  `pulumi:"redistribute6s"`
	Redistributes              []GetRouterBgpRedistribute   `pulumi:"redistributes"`
	RouterId                   string                       `pulumi:"routerId"`
	ScanTime                   int                          `pulumi:"scanTime"`
	Synchronization            string                       `pulumi:"synchronization"`
	TagResolveMode             string                       `pulumi:"tagResolveMode"`
	Vdomparam                  *string                      `pulumi:"vdomparam"`
	Vrf6s                      []GetRouterBgpVrf6           `pulumi:"vrf6s"`
	VrfLeak6s                  []GetRouterBgpVrfLeak6       `pulumi:"vrfLeak6s"`
	VrfLeaks                   []GetRouterBgpVrfLeak        `pulumi:"vrfLeaks"`
	Vrves                      []GetRouterBgpVrf            `pulumi:"vrves"`
}

func LookupRouterBgpOutput(ctx *pulumi.Context, args LookupRouterBgpOutputArgs, opts ...pulumi.InvokeOption) LookupRouterBgpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterBgpResult, error) {
			args := v.(LookupRouterBgpArgs)
			r, err := LookupRouterBgp(ctx, &args, opts...)
			var s LookupRouterBgpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterBgpResultOutput)
}

// A collection of arguments for invoking GetRouterBgp.
type LookupRouterBgpOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterBgpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBgpArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterBgp.
type LookupRouterBgpResultOutput struct{ *pulumi.OutputState }

func (LookupRouterBgpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterBgpResult)(nil)).Elem()
}

func (o LookupRouterBgpResultOutput) ToLookupRouterBgpResultOutput() LookupRouterBgpResultOutput {
	return o
}

func (o LookupRouterBgpResultOutput) ToLookupRouterBgpResultOutputWithContext(ctx context.Context) LookupRouterBgpResultOutput {
	return o
}

func (o LookupRouterBgpResultOutput) AdditionalPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPath }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) AdditionalPath6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPath6 }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) AdditionalPathSelect() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelect }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) AdditionalPathSelect6() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelect6 }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) AdditionalPathSelectVpnv4() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.AdditionalPathSelectVpnv4 }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) AdditionalPathVpnv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AdditionalPathVpnv4 }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) AdminDistances() GetRouterBgpAdminDistanceArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAdminDistance { return v.AdminDistances }).(GetRouterBgpAdminDistanceArrayOutput)
}

func (o LookupRouterBgpResultOutput) AggregateAddress6s() GetRouterBgpAggregateAddress6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAggregateAddress6 { return v.AggregateAddress6s }).(GetRouterBgpAggregateAddress6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) AggregateAddresses() GetRouterBgpAggregateAddressArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpAggregateAddress { return v.AggregateAddresses }).(GetRouterBgpAggregateAddressArrayOutput)
}

func (o LookupRouterBgpResultOutput) AlwaysCompareMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.AlwaysCompareMed }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) As() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.As }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) BestpathAsPathIgnore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathAsPathIgnore }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) BestpathCmpConfedAspath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathCmpConfedAspath }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) BestpathCmpRouterid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathCmpRouterid }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) BestpathMedConfed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathMedConfed }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) BestpathMedMissingAsWorst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.BestpathMedMissingAsWorst }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) ClientToClientReflection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.ClientToClientReflection }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) ConfederationIdentifier() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.ConfederationIdentifier }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) ConfederationPeers() GetRouterBgpConfederationPeerArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpConfederationPeer { return v.ConfederationPeers }).(GetRouterBgpConfederationPeerArrayOutput)
}

func (o LookupRouterBgpResultOutput) Dampening() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Dampening }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) DampeningMaxSuppressTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningMaxSuppressTime }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DampeningReachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningReachabilityHalfLife }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DampeningReuse() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningReuse }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DampeningRouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.DampeningRouteMap }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) DampeningSuppress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningSuppress }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DampeningUnreachabilityHalfLife() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DampeningUnreachabilityHalfLife }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DefaultLocalPreference() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DefaultLocalPreference }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DeterministicMed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.DeterministicMed }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) DistanceExternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceExternal }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DistanceInternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceInternal }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) DistanceLocal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.DistanceLocal }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) EbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.EbgpMultipath }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) EnforceFirstAs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.EnforceFirstAs }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) FastExternalFailover() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.FastExternalFailover }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) GracefulEndOnTimer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.GracefulEndOnTimer }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) GracefulRestart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.GracefulRestart }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) GracefulRestartTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulRestartTime }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) GracefulStalepathTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulStalepathTime }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) GracefulUpdateDelay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.GracefulUpdateDelay }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) HoldtimeTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.HoldtimeTimer }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) IbgpMultipath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.IbgpMultipath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterBgpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) IgnoreOptionalCapability() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.IgnoreOptionalCapability }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) KeepaliveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.KeepaliveTimer }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) LogNeighbourChanges() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.LogNeighbourChanges }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) MultipathRecursiveDistance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.MultipathRecursiveDistance }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) NeighborGroups() GetRouterBgpNeighborGroupArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborGroup { return v.NeighborGroups }).(GetRouterBgpNeighborGroupArrayOutput)
}

func (o LookupRouterBgpResultOutput) NeighborRange6s() GetRouterBgpNeighborRange6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborRange6 { return v.NeighborRange6s }).(GetRouterBgpNeighborRange6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) NeighborRanges() GetRouterBgpNeighborRangeArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighborRange { return v.NeighborRanges }).(GetRouterBgpNeighborRangeArrayOutput)
}

func (o LookupRouterBgpResultOutput) Neighbors() GetRouterBgpNeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNeighbor { return v.Neighbors }).(GetRouterBgpNeighborArrayOutput)
}

func (o LookupRouterBgpResultOutput) Network6s() GetRouterBgpNetwork6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNetwork6 { return v.Network6s }).(GetRouterBgpNetwork6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) NetworkImportCheck() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.NetworkImportCheck }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) Networks() GetRouterBgpNetworkArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpNetwork { return v.Networks }).(GetRouterBgpNetworkArrayOutput)
}

func (o LookupRouterBgpResultOutput) RecursiveInheritPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RecursiveInheritPriority }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) RecursiveNextHop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RecursiveNextHop }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) Redistribute6s() GetRouterBgpRedistribute6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpRedistribute6 { return v.Redistribute6s }).(GetRouterBgpRedistribute6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) Redistributes() GetRouterBgpRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpRedistribute { return v.Redistributes }).(GetRouterBgpRedistributeArrayOutput)
}

func (o LookupRouterBgpResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.RouterId }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) ScanTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) int { return v.ScanTime }).(pulumi.IntOutput)
}

func (o LookupRouterBgpResultOutput) Synchronization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.Synchronization }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) TagResolveMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) string { return v.TagResolveMode }).(pulumi.StringOutput)
}

func (o LookupRouterBgpResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupRouterBgpResultOutput) Vrf6s() GetRouterBgpVrf6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrf6 { return v.Vrf6s }).(GetRouterBgpVrf6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) VrfLeak6s() GetRouterBgpVrfLeak6ArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrfLeak6 { return v.VrfLeak6s }).(GetRouterBgpVrfLeak6ArrayOutput)
}

func (o LookupRouterBgpResultOutput) VrfLeaks() GetRouterBgpVrfLeakArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrfLeak { return v.VrfLeaks }).(GetRouterBgpVrfLeakArrayOutput)
}

func (o LookupRouterBgpResultOutput) Vrves() GetRouterBgpVrfArrayOutput {
	return o.ApplyT(func(v LookupRouterBgpResult) []GetRouterBgpVrf { return v.Vrves }).(GetRouterBgpVrfArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterBgpResultOutput{})
}

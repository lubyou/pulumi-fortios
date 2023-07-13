// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemClusterSync(ctx *pulumi.Context, args *LookupSystemClusterSyncArgs, opts ...pulumi.InvokeOption) (*LookupSystemClusterSyncResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemClusterSyncResult
	err := ctx.Invoke("fortios:index/getSystemClusterSync:GetSystemClusterSync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemClusterSync.
type LookupSystemClusterSyncArgs struct {
	SyncId    int     `pulumi:"syncId"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemClusterSync.
type LookupSystemClusterSyncResult struct {
	DownIntfsBeforeSessSyncs []GetSystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	HbInterval               int                                           `pulumi:"hbInterval"`
	HbLostThreshold          int                                           `pulumi:"hbLostThreshold"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                                  `pulumi:"id"`
	IkeHeartbeatInterval    int                                     `pulumi:"ikeHeartbeatInterval"`
	IkeMonitor              string                                  `pulumi:"ikeMonitor"`
	IkeMonitorInterval      int                                     `pulumi:"ikeMonitorInterval"`
	IpsecTunnelSync         string                                  `pulumi:"ipsecTunnelSync"`
	Peerip                  string                                  `pulumi:"peerip"`
	Peervd                  string                                  `pulumi:"peervd"`
	SecondaryAddIpsecRoutes string                                  `pulumi:"secondaryAddIpsecRoutes"`
	SessionSyncFilters      []GetSystemClusterSyncSessionSyncFilter `pulumi:"sessionSyncFilters"`
	SlaveAddIkeRoutes       string                                  `pulumi:"slaveAddIkeRoutes"`
	SyncId                  int                                     `pulumi:"syncId"`
	Syncvds                 []GetSystemClusterSyncSyncvd            `pulumi:"syncvds"`
	Vdomparam               *string                                 `pulumi:"vdomparam"`
}

func LookupSystemClusterSyncOutput(ctx *pulumi.Context, args LookupSystemClusterSyncOutputArgs, opts ...pulumi.InvokeOption) LookupSystemClusterSyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemClusterSyncResult, error) {
			args := v.(LookupSystemClusterSyncArgs)
			r, err := LookupSystemClusterSync(ctx, &args, opts...)
			var s LookupSystemClusterSyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemClusterSyncResultOutput)
}

// A collection of arguments for invoking GetSystemClusterSync.
type LookupSystemClusterSyncOutputArgs struct {
	SyncId    pulumi.IntInput       `pulumi:"syncId"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemClusterSyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemClusterSyncArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemClusterSync.
type LookupSystemClusterSyncResultOutput struct{ *pulumi.OutputState }

func (LookupSystemClusterSyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemClusterSyncResult)(nil)).Elem()
}

func (o LookupSystemClusterSyncResultOutput) ToLookupSystemClusterSyncResultOutput() LookupSystemClusterSyncResultOutput {
	return o
}

func (o LookupSystemClusterSyncResultOutput) ToLookupSystemClusterSyncResultOutputWithContext(ctx context.Context) LookupSystemClusterSyncResultOutput {
	return o
}

func (o LookupSystemClusterSyncResultOutput) DownIntfsBeforeSessSyncs() GetSystemClusterSyncDownIntfsBeforeSessSyncArrayOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) []GetSystemClusterSyncDownIntfsBeforeSessSync {
		return v.DownIntfsBeforeSessSyncs
	}).(GetSystemClusterSyncDownIntfsBeforeSessSyncArrayOutput)
}

func (o LookupSystemClusterSyncResultOutput) HbInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) int { return v.HbInterval }).(pulumi.IntOutput)
}

func (o LookupSystemClusterSyncResultOutput) HbLostThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) int { return v.HbLostThreshold }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemClusterSyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) IkeHeartbeatInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) int { return v.IkeHeartbeatInterval }).(pulumi.IntOutput)
}

func (o LookupSystemClusterSyncResultOutput) IkeMonitor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.IkeMonitor }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) IkeMonitorInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) int { return v.IkeMonitorInterval }).(pulumi.IntOutput)
}

func (o LookupSystemClusterSyncResultOutput) IpsecTunnelSync() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.IpsecTunnelSync }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) Peerip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.Peerip }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) Peervd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.Peervd }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) SecondaryAddIpsecRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.SecondaryAddIpsecRoutes }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) SessionSyncFilters() GetSystemClusterSyncSessionSyncFilterArrayOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) []GetSystemClusterSyncSessionSyncFilter {
		return v.SessionSyncFilters
	}).(GetSystemClusterSyncSessionSyncFilterArrayOutput)
}

func (o LookupSystemClusterSyncResultOutput) SlaveAddIkeRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) string { return v.SlaveAddIkeRoutes }).(pulumi.StringOutput)
}

func (o LookupSystemClusterSyncResultOutput) SyncId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) int { return v.SyncId }).(pulumi.IntOutput)
}

func (o LookupSystemClusterSyncResultOutput) Syncvds() GetSystemClusterSyncSyncvdArrayOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) []GetSystemClusterSyncSyncvd { return v.Syncvds }).(GetSystemClusterSyncSyncvdArrayOutput)
}

func (o LookupSystemClusterSyncResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemClusterSyncResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemClusterSyncResultOutput{})
}

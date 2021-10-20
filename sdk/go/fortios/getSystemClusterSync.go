// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system clustersync
func LookupSystemClusterSync(ctx *pulumi.Context, args *LookupSystemClusterSyncArgs, opts ...pulumi.InvokeOption) (*LookupSystemClusterSyncResult, error) {
	var rv LookupSystemClusterSyncResult
	err := ctx.Invoke("fortios:index/getSystemClusterSync:GetSystemClusterSync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemClusterSync.
type LookupSystemClusterSyncArgs struct {
	// Specify the syncId of the desired system clustersync.
	SyncId int `pulumi:"syncId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemClusterSync.
type LookupSystemClusterSyncResult struct {
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs []GetSystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	// Heartbeat interval (1 - 10 sec).
	HbInterval int `pulumi:"hbInterval"`
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold int `pulumi:"hbLostThreshold"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable IPsec tunnel synchronization.
	IpsecTunnelSync string `pulumi:"ipsecTunnelSync"`
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip string `pulumi:"peerip"`
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd string `pulumi:"peervd"`
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter GetSystemClusterSyncSessionSyncFilter `pulumi:"sessionSyncFilter"`
	// Enable/disable IKE route announcement on the backup unit.
	SlaveAddIkeRoutes string `pulumi:"slaveAddIkeRoutes"`
	// Sync ID.
	SyncId int `pulumi:"syncId"`
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds   []GetSystemClusterSyncSyncvd `pulumi:"syncvds"`
	Vdomparam *string                      `pulumi:"vdomparam"`
}
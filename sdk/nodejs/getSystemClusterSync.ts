// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSystemClusterSync(args: GetSystemClusterSyncArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemClusterSyncResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemClusterSync:GetSystemClusterSync", {
        "syncId": args.syncId,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemClusterSync.
 */
export interface GetSystemClusterSyncArgs {
    syncId: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemClusterSync.
 */
export interface GetSystemClusterSyncResult {
    readonly downIntfsBeforeSessSyncs: outputs.GetSystemClusterSyncDownIntfsBeforeSessSync[];
    readonly hbInterval: number;
    readonly hbLostThreshold: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ikeHeartbeatInterval: number;
    readonly ikeMonitor: string;
    readonly ikeMonitorInterval: number;
    readonly ipsecTunnelSync: string;
    readonly peerip: string;
    readonly peervd: string;
    readonly secondaryAddIpsecRoutes: string;
    readonly sessionSyncFilters: outputs.GetSystemClusterSyncSessionSyncFilter[];
    readonly slaveAddIkeRoutes: string;
    readonly syncId: number;
    readonly syncvds: outputs.GetSystemClusterSyncSyncvd[];
    readonly vdomparam?: string;
}
export function getSystemClusterSyncOutput(args: GetSystemClusterSyncOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemClusterSyncResult> {
    return pulumi.output(args).apply((a: any) => getSystemClusterSync(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemClusterSync.
 */
export interface GetSystemClusterSyncOutputArgs {
    syncId: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

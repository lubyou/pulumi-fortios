// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSystemCentralManagement(args?: GetSystemCentralManagementArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemCentralManagementResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemCentralManagement:GetSystemCentralManagement", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemCentralManagement.
 */
export interface GetSystemCentralManagementArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemCentralManagement.
 */
export interface GetSystemCentralManagementResult {
    readonly allowMonitor: string;
    readonly allowPushConfiguration: string;
    readonly allowPushFirmware: string;
    readonly allowRemoteFirmwareUpgrade: string;
    readonly caCert: string;
    readonly encAlgorithm: string;
    readonly fmg: string;
    readonly fmgSourceIp: string;
    readonly fmgSourceIp6: string;
    readonly fmgUpdatePort: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeDefaultServers: string;
    readonly interface: string;
    readonly interfaceSelectMethod: string;
    readonly localCert: string;
    readonly mode: string;
    readonly scheduleConfigRestore: string;
    readonly scheduleScriptRestore: string;
    readonly serialNumber: string;
    readonly serverLists: outputs.GetSystemCentralManagementServerList[];
    readonly type: string;
    readonly vdom: string;
    readonly vdomparam?: string;
}
export function getSystemCentralManagementOutput(args?: GetSystemCentralManagementOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemCentralManagementResult> {
    return pulumi.output(args).apply((a: any) => getSystemCentralManagement(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemCentralManagement.
 */
export interface GetSystemCentralManagementOutputArgs {
    vdomparam?: pulumi.Input<string>;
}

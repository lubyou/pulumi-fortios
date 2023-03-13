// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemFm(args?: GetSystemFmArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemFmResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemFm:GetSystemFm", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemFm.
 */
export interface GetSystemFmArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemFm.
 */
export interface GetSystemFmResult {
    readonly autoBackup: string;
    readonly fosid: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ip: string;
    readonly ipsec: string;
    readonly scheduledConfigRestore: string;
    readonly status: string;
    readonly vdom: string;
    readonly vdomparam?: string;
}
export function getSystemFmOutput(args?: GetSystemFmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemFmResult> {
    return pulumi.output(args).apply((a: any) => getSystemFm(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemFm.
 */
export interface GetSystemFmOutputArgs {
    vdomparam?: pulumi.Input<string>;
}

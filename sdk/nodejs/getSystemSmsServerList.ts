// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemSmsServerList(args?: GetSystemSmsServerListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemSmsServerListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemSmsServerList:GetSystemSmsServerList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemSmsServerList.
 */
export interface GetSystemSmsServerListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemSmsServerList.
 */
export interface GetSystemSmsServerListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getSystemSmsServerListOutput(args?: GetSystemSmsServerListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemSmsServerListResult> {
    return pulumi.output(args).apply((a: any) => getSystemSmsServerList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemSmsServerList.
 */
export interface GetSystemSmsServerListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemDnsDatabaseList(args?: GetSystemDnsDatabaseListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemDnsDatabaseListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemDnsDatabaseList:GetSystemDnsDatabaseList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemDnsDatabaseList.
 */
export interface GetSystemDnsDatabaseListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemDnsDatabaseList.
 */
export interface GetSystemDnsDatabaseListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getSystemDnsDatabaseListOutput(args?: GetSystemDnsDatabaseListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemDnsDatabaseListResult> {
    return pulumi.output(args).apply((a: any) => getSystemDnsDatabaseList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemDnsDatabaseList.
 */
export interface GetSystemDnsDatabaseListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAutoupdatePushUpdate(args?: GetSystemAutoupdatePushUpdateArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutoupdatePushUpdateResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAutoupdatePushUpdate:GetSystemAutoupdatePushUpdate", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
 */
export interface GetSystemAutoupdatePushUpdateArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutoupdatePushUpdate.
 */
export interface GetSystemAutoupdatePushUpdateResult {
    readonly address: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly override: string;
    readonly port: number;
    readonly status: string;
    readonly vdomparam?: string;
}
export function getSystemAutoupdatePushUpdateOutput(args?: GetSystemAutoupdatePushUpdateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAutoupdatePushUpdateResult> {
    return pulumi.output(args).apply((a: any) => getSystemAutoupdatePushUpdate(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
 */
export interface GetSystemAutoupdatePushUpdateOutputArgs {
    vdomparam?: pulumi.Input<string>;
}

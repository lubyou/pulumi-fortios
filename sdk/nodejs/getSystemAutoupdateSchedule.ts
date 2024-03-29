// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAutoupdateSchedule(args?: GetSystemAutoupdateScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutoupdateScheduleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAutoupdateSchedule:GetSystemAutoupdateSchedule", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutoupdateSchedule.
 */
export interface GetSystemAutoupdateScheduleArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutoupdateSchedule.
 */
export interface GetSystemAutoupdateScheduleResult {
    readonly day: string;
    readonly frequency: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly status: string;
    readonly time: string;
    readonly vdomparam?: string;
}
export function getSystemAutoupdateScheduleOutput(args?: GetSystemAutoupdateScheduleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAutoupdateScheduleResult> {
    return pulumi.output(args).apply((a: any) => getSystemAutoupdateSchedule(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAutoupdateSchedule.
 */
export interface GetSystemAutoupdateScheduleOutputArgs {
    vdomparam?: pulumi.Input<string>;
}

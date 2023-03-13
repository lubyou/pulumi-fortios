// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemTosBasedPriority(args: GetSystemTosBasedPriorityArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemTosBasedPriorityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemTosBasedPriority:GetSystemTosBasedPriority", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemTosBasedPriority.
 */
export interface GetSystemTosBasedPriorityArgs {
    fosid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemTosBasedPriority.
 */
export interface GetSystemTosBasedPriorityResult {
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly priority: string;
    readonly tos: number;
    readonly vdomparam?: string;
}
export function getSystemTosBasedPriorityOutput(args: GetSystemTosBasedPriorityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemTosBasedPriorityResult> {
    return pulumi.output(args).apply((a: any) => getSystemTosBasedPriority(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemTosBasedPriority.
 */
export interface GetSystemTosBasedPriorityOutputArgs {
    fosid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

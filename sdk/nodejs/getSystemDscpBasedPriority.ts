// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemDscpBasedPriority(args: GetSystemDscpBasedPriorityArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemDscpBasedPriorityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemDscpBasedPriority:GetSystemDscpBasedPriority", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemDscpBasedPriority.
 */
export interface GetSystemDscpBasedPriorityArgs {
    fosid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemDscpBasedPriority.
 */
export interface GetSystemDscpBasedPriorityResult {
    readonly ds: number;
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly priority: string;
    readonly vdomparam?: string;
}
export function getSystemDscpBasedPriorityOutput(args: GetSystemDscpBasedPriorityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemDscpBasedPriorityResult> {
    return pulumi.output(args).apply((a: any) => getSystemDscpBasedPriority(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemDscpBasedPriority.
 */
export interface GetSystemDscpBasedPriorityOutputArgs {
    fosid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

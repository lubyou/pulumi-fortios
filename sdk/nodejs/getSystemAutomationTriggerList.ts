// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAutomationTriggerList(args?: GetSystemAutomationTriggerListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutomationTriggerListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAutomationTriggerList:GetSystemAutomationTriggerList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutomationTriggerList.
 */
export interface GetSystemAutomationTriggerListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutomationTriggerList.
 */
export interface GetSystemAutomationTriggerListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getSystemAutomationTriggerListOutput(args?: GetSystemAutomationTriggerListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAutomationTriggerListResult> {
    return pulumi.output(args).apply((a: any) => getSystemAutomationTriggerList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAutomationTriggerList.
 */
export interface GetSystemAutomationTriggerListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

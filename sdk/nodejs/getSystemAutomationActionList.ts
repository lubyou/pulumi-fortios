// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemAutomationAction`.
 */
export function getSystemAutomationActionList(args?: GetSystemAutomationActionListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutomationActionListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemAutomationActionList:GetSystemAutomationActionList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutomationActionList.
 */
export interface GetSystemAutomationActionListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutomationActionList.
 */
export interface GetSystemAutomationActionListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.SystemAutomationAction`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}

export function getSystemAutomationActionListOutput(args?: GetSystemAutomationActionListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAutomationActionListResult> {
    return pulumi.output(args).apply(a => getSystemAutomationActionList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAutomationActionList.
 */
export interface GetSystemAutomationActionListOutputArgs {
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

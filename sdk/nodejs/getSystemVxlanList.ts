// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemVxlan`.
 */
export function getSystemVxlanList(args?: GetSystemVxlanListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemVxlanListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemVxlanList:GetSystemVxlanList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemVxlanList.
 */
export interface GetSystemVxlanListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemVxlanList.
 */
export interface GetSystemVxlanListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.SystemVxlan`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}

export function getSystemVxlanListOutput(args?: GetSystemVxlanListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemVxlanListResult> {
    return pulumi.output(args).apply(a => getSystemVxlanList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemVxlanList.
 */
export interface GetSystemVxlanListOutputArgs {
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

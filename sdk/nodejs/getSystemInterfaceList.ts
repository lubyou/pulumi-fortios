// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemInterface`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const sample1 = fortios.GetSystemInterfaceList({
 *     filter: "name!=port1",
 * });
 * export const output1 = data.fortios_system_interfacelist.sample2.namelist;
 * ```
 */
export function getSystemInterfaceList(args?: GetSystemInterfaceListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemInterfaceListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemInterfaceList:GetSystemInterfaceList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemInterfaceList.
 */
export interface GetSystemInterfaceListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemInterfaceList.
 */
export interface GetSystemInterfaceListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.SystemInterface`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}

export function getSystemInterfaceListOutput(args?: GetSystemInterfaceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemInterfaceListResult> {
    return pulumi.output(args).apply(a => getSystemInterfaceList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemInterfaceList.
 */
export interface GetSystemInterfaceListOutputArgs {
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

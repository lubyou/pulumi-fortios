// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.RouterStatic`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const sample1 = fortios.GetRouterStaticList({
 *     filter: "seq_num>1",
 * });
 * export const output1 = sample1.then(sample1 => sample1.seqNumlists);
 * ```
 */
export function getRouterStaticList(args?: GetRouterStaticListArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterStaticListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getRouterStaticList:GetRouterStaticList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterStaticList.
 */
export interface GetRouterStaticListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterStaticList.
 */
export interface GetRouterStaticListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.RouterStatic`.
     */
    readonly seqNumlists: number[];
    readonly vdomparam?: string;
}

export function getRouterStaticListOutput(args?: GetRouterStaticListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterStaticListResult> {
    return pulumi.output(args).apply(a => getRouterStaticList(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterStaticList.
 */
export interface GetRouterStaticListOutputArgs {
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

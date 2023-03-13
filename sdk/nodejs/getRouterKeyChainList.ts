// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getRouterKeyChainList(args?: GetRouterKeyChainListArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterKeyChainListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterKeyChainList:GetRouterKeyChainList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterKeyChainList.
 */
export interface GetRouterKeyChainListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterKeyChainList.
 */
export interface GetRouterKeyChainListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getRouterKeyChainListOutput(args?: GetRouterKeyChainListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterKeyChainListResult> {
    return pulumi.output(args).apply((a: any) => getRouterKeyChainList(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterKeyChainList.
 */
export interface GetRouterKeyChainListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

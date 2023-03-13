// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getRouterCommunityListList(args?: GetRouterCommunityListListArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterCommunityListListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterCommunityListList:GetRouterCommunityListList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterCommunityListList.
 */
export interface GetRouterCommunityListListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterCommunityListList.
 */
export interface GetRouterCommunityListListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getRouterCommunityListListOutput(args?: GetRouterCommunityListListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterCommunityListListResult> {
    return pulumi.output(args).apply((a: any) => getRouterCommunityListList(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterCommunityListList.
 */
export interface GetRouterCommunityListListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

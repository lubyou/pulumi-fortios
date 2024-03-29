// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFirewallServiceCategoryList(args?: GetFirewallServiceCategoryListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallServiceCategoryListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallServiceCategoryList:GetFirewallServiceCategoryList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallServiceCategoryList.
 */
export interface GetFirewallServiceCategoryListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallServiceCategoryList.
 */
export interface GetFirewallServiceCategoryListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getFirewallServiceCategoryListOutput(args?: GetFirewallServiceCategoryListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallServiceCategoryListResult> {
    return pulumi.output(args).apply((a: any) => getFirewallServiceCategoryList(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallServiceCategoryList.
 */
export interface GetFirewallServiceCategoryListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

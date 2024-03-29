// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFirewallAddress6TemplateList(args?: GetFirewallAddress6TemplateListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallAddress6TemplateListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallAddress6TemplateList:GetFirewallAddress6TemplateList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallAddress6TemplateList.
 */
export interface GetFirewallAddress6TemplateListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallAddress6TemplateList.
 */
export interface GetFirewallAddress6TemplateListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namelists: string[];
    readonly vdomparam?: string;
}
export function getFirewallAddress6TemplateListOutput(args?: GetFirewallAddress6TemplateListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallAddress6TemplateListResult> {
    return pulumi.output(args).apply((a: any) => getFirewallAddress6TemplateList(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallAddress6TemplateList.
 */
export interface GetFirewallAddress6TemplateListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

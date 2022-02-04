// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall internetservicegroup
 */
export function getFirewallInternetServiceGroup(args: GetFirewallInternetServiceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallInternetServiceGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallInternetServiceGroup:GetFirewallInternetServiceGroup", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallInternetServiceGroup.
 */
export interface GetFirewallInternetServiceGroupArgs {
    /**
     * Specify the name of the desired firewall internetservicegroup.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallInternetServiceGroup.
 */
export interface GetFirewallInternetServiceGroupResult {
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * How this service may be used (source, destination or both).
     */
    readonly direction: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Internet Service group member. The structure of `member` block is documented below.
     */
    readonly members: outputs.GetFirewallInternetServiceGroupMember[];
    /**
     * Internet Service name.
     */
    readonly name: string;
    readonly vdomparam?: string;
}

export function getFirewallInternetServiceGroupOutput(args: GetFirewallInternetServiceGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallInternetServiceGroupResult> {
    return pulumi.output(args).apply(a => getFirewallInternetServiceGroup(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallInternetServiceGroup.
 */
export interface GetFirewallInternetServiceGroupOutputArgs {
    /**
     * Specify the name of the desired firewall internetservicegroup.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

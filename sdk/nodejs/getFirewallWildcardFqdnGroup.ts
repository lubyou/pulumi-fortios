// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewallwildcardfqdn group
 */
export function getFirewallWildcardFqdnGroup(args: GetFirewallWildcardFqdnGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallWildcardFqdnGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallWildcardFqdnGroup:GetFirewallWildcardFqdnGroup", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallWildcardFqdnGroup.
 */
export interface GetFirewallWildcardFqdnGroupArgs {
    /**
     * Specify the name of the desired firewallwildcardfqdn group.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallWildcardFqdnGroup.
 */
export interface GetFirewallWildcardFqdnGroupResult {
    /**
     * GUI icon color.
     */
    readonly color: number;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Address group members. The structure of `member` block is documented below.
     */
    readonly members: outputs.GetFirewallWildcardFqdnGroupMember[];
    /**
     * Address name.
     */
    readonly name: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable address visibility.
     */
    readonly visibility: string;
}

export function getFirewallWildcardFqdnGroupOutput(args: GetFirewallWildcardFqdnGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallWildcardFqdnGroupResult> {
    return pulumi.output(args).apply(a => getFirewallWildcardFqdnGroup(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallWildcardFqdnGroup.
 */
export interface GetFirewallWildcardFqdnGroupOutputArgs {
    /**
     * Specify the name of the desired firewallwildcardfqdn group.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

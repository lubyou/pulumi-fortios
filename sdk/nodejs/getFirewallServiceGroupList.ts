// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallServiceGroup`.
 */
export function getFirewallServiceGroupList(args?: GetFirewallServiceGroupListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallServiceGroupListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getFirewallServiceGroupList:GetFirewallServiceGroupList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallServiceGroupList.
 */
export interface GetFirewallServiceGroupListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallServiceGroupList.
 */
export interface GetFirewallServiceGroupListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallServiceGroup`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}

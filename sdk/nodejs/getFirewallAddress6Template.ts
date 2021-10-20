// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall address6template
 */
export function getFirewallAddress6Template(args: GetFirewallAddress6TemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallAddress6TemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getFirewallAddress6Template:GetFirewallAddress6Template", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallAddress6Template.
 */
export interface GetFirewallAddress6TemplateArgs {
    /**
     * Specify the name of the desired firewall address6template.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallAddress6Template.
 */
export interface GetFirewallAddress6TemplateResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IPv6 address prefix.
     */
    readonly ip6: string;
    /**
     * Subnet segment value name.
     */
    readonly name: string;
    /**
     * Number of IPv6 subnet segments.
     */
    readonly subnetSegmentCount: number;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    readonly subnetSegments: outputs.GetFirewallAddress6TemplateSubnetSegment[];
    readonly vdomparam?: string;
}
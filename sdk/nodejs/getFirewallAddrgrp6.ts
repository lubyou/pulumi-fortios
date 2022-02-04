// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall addrgrp6
 */
export function getFirewallAddrgrp6(args: GetFirewallAddrgrp6Args, opts?: pulumi.InvokeOptions): Promise<GetFirewallAddrgrp6Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallAddrgrp6:GetFirewallAddrgrp6", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallAddrgrp6.
 */
export interface GetFirewallAddrgrp6Args {
    /**
     * Specify the name of the desired firewall addrgrp6.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallAddrgrp6.
 */
export interface GetFirewallAddrgrp6Result {
    /**
     * Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
     */
    readonly color: number;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * Security Fabric global object setting.
     */
    readonly fabricObject: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Address objects contained within the group. The structure of `member` block is documented below.
     */
    readonly members: outputs.GetFirewallAddrgrp6Member[];
    /**
     * Tag name.
     */
    readonly name: string;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.GetFirewallAddrgrp6Tagging[];
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable address group6 visibility in the GUI.
     */
    readonly visibility: string;
}

export function getFirewallAddrgrp6Output(args: GetFirewallAddrgrp6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallAddrgrp6Result> {
    return pulumi.output(args).apply(a => getFirewallAddrgrp6(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallAddrgrp6.
 */
export interface GetFirewallAddrgrp6OutputArgs {
    /**
     * Specify the name of the desired firewall addrgrp6.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios router prefixlist
 */
export function getRouterPrefixList(args: GetRouterPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterPrefixListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getRouterPrefixList:GetRouterPrefixList", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterPrefixList.
 */
export interface GetRouterPrefixListArgs {
    /**
     * Specify the name of the desired router prefixlist.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterPrefixList.
 */
export interface GetRouterPrefixListResult {
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name.
     */
    readonly name: string;
    /**
     * IPv4 prefix list rule. The structure of `rule` block is documented below.
     */
    readonly rules: outputs.GetRouterPrefixListRule[];
    readonly vdomparam?: string;
}

export function getRouterPrefixListOutput(args: GetRouterPrefixListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterPrefixListResult> {
    return pulumi.output(args).apply(a => getRouterPrefixList(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterPrefixList.
 */
export interface GetRouterPrefixListOutputArgs {
    /**
     * Specify the name of the desired router prefixlist.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios router bfd6
 */
export function getRouterBfd6(args?: GetRouterBfd6Args, opts?: pulumi.InvokeOptions): Promise<GetRouterBfd6Result> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getRouterBfd6:GetRouterBfd6", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterBfd6.
 */
export interface GetRouterBfd6Args {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterBfd6.
 */
export interface GetRouterBfd6Result {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
     */
    readonly neighbors: outputs.GetRouterBfd6Neighbor[];
    readonly vdomparam?: string;
}

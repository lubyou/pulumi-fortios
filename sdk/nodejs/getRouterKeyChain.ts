// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios router keychain
 */
export function getRouterKeyChain(args: GetRouterKeyChainArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterKeyChainResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getRouterKeyChain:GetRouterKeyChain", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterKeyChain.
 */
export interface GetRouterKeyChainArgs {
    /**
     * Specify the name of the desired router keychain.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterKeyChain.
 */
export interface GetRouterKeyChainResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Configuration method to edit key settings. The structure of `key` block is documented below.
     */
    readonly keys: outputs.GetRouterKeyChainKey[];
    /**
     * Key-chain name.
     */
    readonly name: string;
    readonly vdomparam?: string;
}

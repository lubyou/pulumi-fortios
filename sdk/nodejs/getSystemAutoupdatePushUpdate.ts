// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios systemautoupdate pushupdate
 */
export function getSystemAutoupdatePushUpdate(args?: GetSystemAutoupdatePushUpdateArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutoupdatePushUpdateResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getSystemAutoupdatePushUpdate:GetSystemAutoupdatePushUpdate", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
 */
export interface GetSystemAutoupdatePushUpdateArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutoupdatePushUpdate.
 */
export interface GetSystemAutoupdatePushUpdateResult {
    /**
     * Push update override server.
     */
    readonly address: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable push update override server.
     */
    readonly override: string;
    /**
     * Push update override port. (Do not overlap with other service ports)
     */
    readonly port: number;
    /**
     * Enable/disable push updates.
     */
    readonly status: string;
    readonly vdomparam?: string;
}

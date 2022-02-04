// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system fm
 */
export function getSystemFm(args?: GetSystemFmArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemFmResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemFm:GetSystemFm", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemFm.
 */
export interface GetSystemFmArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemFm.
 */
export interface GetSystemFmResult {
    /**
     * Enable/disable automatic backup.
     */
    readonly autoBackup: string;
    /**
     * ID.
     */
    readonly fosid: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IP address.
     */
    readonly ip: string;
    /**
     * Enable/disable IPsec.
     */
    readonly ipsec: string;
    /**
     * Enable/disable scheduled configuration restore.
     */
    readonly scheduledConfigRestore: string;
    /**
     * Enable/disable FM.
     */
    readonly status: string;
    /**
     * VDOM.
     */
    readonly vdom: string;
    readonly vdomparam?: string;
}

export function getSystemFmOutput(args?: GetSystemFmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemFmResult> {
    return pulumi.output(args).apply(a => getSystemFm(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemFm.
 */
export interface GetSystemFmOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

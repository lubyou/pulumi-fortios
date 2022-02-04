// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system netflow
 */
export function getSystemNetflow(args?: GetSystemNetflowArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemNetflowResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemNetflow:GetSystemNetflow", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemNetflow.
 */
export interface GetSystemNetflowArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemNetflow.
 */
export interface GetSystemNetflowResult {
    /**
     * Timeout to report active flows (1 - 60 min, default = 30).
     */
    readonly activeFlowTimeout: number;
    /**
     * Collector IP.
     */
    readonly collectorIp: string;
    /**
     * NetFlow collector port number.
     */
    readonly collectorPort: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
     */
    readonly inactiveFlowTimeout: number;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly interface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly interfaceSelectMethod: string;
    /**
     * Source IP address for communication with the NetFlow agent.
     */
    readonly sourceIp: string;
    /**
     * Counter of flowset records before resending a template flowset record.
     */
    readonly templateTxCounter: number;
    /**
     * Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
     */
    readonly templateTxTimeout: number;
    readonly vdomparam?: string;
}

export function getSystemNetflowOutput(args?: GetSystemNetflowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemNetflowResult> {
    return pulumi.output(args).apply(a => getSystemNetflow(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemNetflow.
 */
export interface GetSystemNetflowOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

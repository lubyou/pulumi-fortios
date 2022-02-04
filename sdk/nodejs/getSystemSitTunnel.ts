// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system sittunnel
 */
export function getSystemSitTunnel(args: GetSystemSitTunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemSitTunnelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemSitTunnel:GetSystemSitTunnel", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemSitTunnel.
 */
export interface GetSystemSitTunnelArgs {
    /**
     * Specify the name of the desired system sittunnel.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemSitTunnel.
 */
export interface GetSystemSitTunnelResult {
    /**
     * Enable/disable tunnel ASIC offloading.
     */
    readonly autoAsicOffload: string;
    /**
     * Destination IP address of the tunnel.
     */
    readonly destination: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Interface name.
     */
    readonly interface: string;
    /**
     * IPv6 address of the tunnel.
     */
    readonly ip6: string;
    /**
     * Tunnel name.
     */
    readonly name: string;
    /**
     * Source IP address of the tunnel.
     */
    readonly source: string;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway.
     */
    readonly useSdwan: string;
    readonly vdomparam?: string;
}

export function getSystemSitTunnelOutput(args: GetSystemSitTunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemSitTunnelResult> {
    return pulumi.output(args).apply(a => getSystemSitTunnel(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemSitTunnel.
 */
export interface GetSystemSitTunnelOutputArgs {
    /**
     * Specify the name of the desired system sittunnel.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system csf
 */
export function getSystemCsf(args?: GetSystemCsfArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemCsfResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getSystemCsf:GetSystemCsf", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemCsf.
 */
export interface GetSystemCsfArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemCsf.
 */
export interface GetSystemCsfResult {
    /**
     * Accept connections with unknown certificates and ask admin for approval.
     */
    readonly acceptAuthByCert: string;
    /**
     * Authorization request type.
     */
    readonly authorizationRequestType: string;
    /**
     * Certificate.
     */
    readonly certificate: string;
    /**
     * Configuration sync mode.
     */
    readonly configurationSync: string;
    /**
     * Fabric device configuration. The structure of `fabricDevice` block is documented below.
     */
    readonly fabricDevices: outputs.GetSystemCsfFabricDevice[];
    /**
     * Fabric CMDB Object Unification
     */
    readonly fabricObjectUnification: string;
    /**
     * Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
     */
    readonly fixedKey: string;
    /**
     * Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
     */
    readonly groupName: string;
    /**
     * Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
     */
    readonly groupPassword: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
     */
    readonly managementIp: string;
    /**
     * Overriding port for management connection (Overrides admin port).
     */
    readonly managementPort: number;
    /**
     * SAML setting configuration synchronization.
     */
    readonly samlConfigurationSync: string;
    /**
     * Enable/disable Security Fabric.
     */
    readonly status: string;
    /**
     * Pre-authorized and blocked security fabric nodes. The structure of `trustedList` block is documented below.
     */
    readonly trustedLists: outputs.GetSystemCsfTrustedList[];
    /**
     * IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
     */
    readonly upstreamIp: string;
    /**
     * The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
     */
    readonly upstreamPort: number;
    readonly vdomparam?: string;
}

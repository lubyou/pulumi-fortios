// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Read/Update/Delete firewall object address for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.FortimanagerFirewallObjectAddress("test1", {
 *     associatedIntf: "any",
 *     comment: "test obj address",
 *     fqdn: "fqdn.google.com",
 *     type: "fqdn",
 * });
 * const test2 = new fortios.FortimanagerFirewallObjectAddress("test2", {
 *     allowRouting: "disable",
 *     associatedIntf: "any",
 *     comment: "test obj address",
 *     subnet: "2.2.2.0 255.255.255.0",
 *     type: "ipmask",
 * });
 * const test3 = new fortios.FortimanagerFirewallObjectAddress("test3", {
 *     associatedIntf: "any",
 *     comment: "test obj address",
 *     endIp: "2.2.2.100",
 *     startIp: "2.2.2.1",
 *     type: "iprange",
 * });
 * ```
 */
export class FortimanagerFirewallObjectAddress extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerFirewallObjectAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerFirewallObjectAddressState, opts?: pulumi.CustomResourceOptions): FortimanagerFirewallObjectAddress {
        return new FortimanagerFirewallObjectAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerFirewallObjectAddress:FortimanagerFirewallObjectAddress';

    /**
     * Returns true if the given object is an instance of FortimanagerFirewallObjectAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerFirewallObjectAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerFirewallObjectAddress.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable use of this address in the static route configuration. default is "disable".
     */
    public readonly allowRouting!: pulumi.Output<string | undefined>;
    /**
     * Network interface associated with address.
     */
    public readonly associatedIntf!: pulumi.Output<string | undefined>;
    /**
     * Comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    public readonly endIp!: pulumi.Output<string | undefined>;
    /**
     * Fully Qualified Domain Name address.
     */
    public readonly fqdn!: pulumi.Output<string | undefined>;
    /**
     * Address name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    public readonly startIp!: pulumi.Output<string | undefined>;
    /**
     * IPv4 address/mask
     */
    public readonly subnet!: pulumi.Output<string | undefined>;
    /**
     * Type of address, Enum: ["ipmask", "iprange", "fqdn"].
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerFirewallObjectAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerFirewallObjectAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerFirewallObjectAddressArgs | FortimanagerFirewallObjectAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerFirewallObjectAddressState | undefined;
            inputs["adom"] = state ? state.adom : undefined;
            inputs["allowRouting"] = state ? state.allowRouting : undefined;
            inputs["associatedIntf"] = state ? state.associatedIntf : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["endIp"] = state ? state.endIp : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["startIp"] = state ? state.startIp : undefined;
            inputs["subnet"] = state ? state.subnet : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FortimanagerFirewallObjectAddressArgs | undefined;
            inputs["adom"] = args ? args.adom : undefined;
            inputs["allowRouting"] = args ? args.allowRouting : undefined;
            inputs["associatedIntf"] = args ? args.associatedIntf : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["endIp"] = args ? args.endIp : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["startIp"] = args ? args.startIp : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FortimanagerFirewallObjectAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerFirewallObjectAddress resources.
 */
export interface FortimanagerFirewallObjectAddressState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Enable/disable use of this address in the static route configuration. default is "disable".
     */
    allowRouting?: pulumi.Input<string>;
    /**
     * Network interface associated with address.
     */
    associatedIntf?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name address.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * IPv4 address/mask
     */
    subnet?: pulumi.Input<string>;
    /**
     * Type of address, Enum: ["ipmask", "iprange", "fqdn"].
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerFirewallObjectAddress resource.
 */
export interface FortimanagerFirewallObjectAddressArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Enable/disable use of this address in the static route configuration. default is "disable".
     */
    allowRouting?: pulumi.Input<string>;
    /**
     * Network interface associated with address.
     */
    associatedIntf?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name address.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * IPv4 address/mask
     */
    subnet?: pulumi.Input<string>;
    /**
     * Type of address, Enum: ["ipmask", "iprange", "fqdn"].
     */
    type?: pulumi.Input<string>;
}
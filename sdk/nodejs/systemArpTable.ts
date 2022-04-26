// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure ARP table.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemArpTable("trname", {
 *     fosid: 11,
 *     interface: "port2",
 *     ip: "1.1.1.1",
 *     mac: "08:00:27:1c:a3:8b",
 * });
 * ```
 *
 * ## Import
 *
 * System ArpTable can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemArpTable:SystemArpTable labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemArpTable:SystemArpTable labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemArpTable extends pulumi.CustomResource {
    /**
     * Get an existing SystemArpTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemArpTableState, opts?: pulumi.CustomResourceOptions): SystemArpTable {
        return new SystemArpTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemArpTable:SystemArpTable';

    /**
     * Returns true if the given object is an instance of SystemArpTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemArpTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemArpTable.__pulumiType;
    }

    /**
     * Unique integer ID of the entry.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * MAC address.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemArpTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemArpTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemArpTableArgs | SystemArpTableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemArpTableState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemArpTableArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.mac === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mac'");
            }
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemArpTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemArpTable resources.
 */
export interface SystemArpTableState {
    /**
     * Unique integer ID of the entry.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemArpTable resource.
 */
export interface SystemArpTableArgs {
    /**
     * Unique integer ID of the entry.
     */
    fosid: pulumi.Input<number>;
    /**
     * Interface name.
     */
    interface: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

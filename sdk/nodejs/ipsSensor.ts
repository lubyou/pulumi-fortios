// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPS sensor.
 *
 * ## Import
 *
 * Ips Sensor can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/ipsSensor:IpsSensor labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class IpsSensor extends pulumi.CustomResource {
    /**
     * Get an existing IpsSensor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsSensorState, opts?: pulumi.CustomResourceOptions): IpsSensor {
        return new IpsSensor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ipsSensor:IpsSensor';

    /**
     * Returns true if the given object is an instance of IpsSensor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsSensor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsSensor.__pulumiType;
    }

    /**
     * Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
     */
    public readonly blockMaliciousUrl!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * IPS sensor filter. The structure of `entries` block is documented below.
     */
    public readonly entries!: pulumi.Output<outputs.IpsSensorEntry[] | undefined>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * IPS sensor filter. The structure of `filter` block is documented below.
     */
    public readonly filters!: pulumi.Output<outputs.IpsSensorFilter[] | undefined>;
    /**
     * Filter name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IPS override rule. The structure of `override` block is documented below.
     */
    public readonly overrides!: pulumi.Output<outputs.IpsSensorOverride[] | undefined>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
     */
    public readonly scanBotnetConnections!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IpsSensor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpsSensorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsSensorArgs | IpsSensorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsSensorState | undefined;
            inputs["blockMaliciousUrl"] = state ? state.blockMaliciousUrl : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["entries"] = state ? state.entries : undefined;
            inputs["extendedLog"] = state ? state.extendedLog : undefined;
            inputs["filters"] = state ? state.filters : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overrides"] = state ? state.overrides : undefined;
            inputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            inputs["scanBotnetConnections"] = state ? state.scanBotnetConnections : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IpsSensorArgs | undefined;
            inputs["blockMaliciousUrl"] = args ? args.blockMaliciousUrl : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["entries"] = args ? args.entries : undefined;
            inputs["extendedLog"] = args ? args.extendedLog : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overrides"] = args ? args.overrides : undefined;
            inputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            inputs["scanBotnetConnections"] = args ? args.scanBotnetConnections : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IpsSensor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsSensor resources.
 */
export interface IpsSensorState {
    /**
     * Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
     */
    blockMaliciousUrl?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * IPS sensor filter. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.IpsSensorEntry>[]>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * IPS sensor filter. The structure of `filter` block is documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.IpsSensorFilter>[]>;
    /**
     * Filter name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPS override rule. The structure of `override` block is documented below.
     */
    overrides?: pulumi.Input<pulumi.Input<inputs.IpsSensorOverride>[]>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
     */
    scanBotnetConnections?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsSensor resource.
 */
export interface IpsSensorArgs {
    /**
     * Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
     */
    blockMaliciousUrl?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * IPS sensor filter. The structure of `entries` block is documented below.
     */
    entries?: pulumi.Input<pulumi.Input<inputs.IpsSensorEntry>[]>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * IPS sensor filter. The structure of `filter` block is documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.IpsSensorFilter>[]>;
    /**
     * Filter name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPS override rule. The structure of `override` block is documented below.
     */
    overrides?: pulumi.Input<pulumi.Input<inputs.IpsSensorOverride>[]>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
     */
    scanBotnetConnections?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
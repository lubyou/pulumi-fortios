// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemZone("trname", {
 *     intrazone: "allow",
 * });
 * ```
 *
 * ## Import
 *
 * System Zone can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemZone:SystemZone labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemZone extends pulumi.CustomResource {
    /**
     * Get an existing SystemZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemZoneState, opts?: pulumi.CustomResourceOptions): SystemZone {
        return new SystemZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemZone:SystemZone';

    /**
     * Returns true if the given object is an instance of SystemZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemZone.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.SystemZoneInterface[] | undefined>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    public readonly intrazone!: pulumi.Output<string>;
    /**
     * Tag name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.SystemZoneTagging[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemZoneArgs | SystemZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemZoneState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["interfaces"] = state ? state.interfaces : undefined;
            inputs["intrazone"] = state ? state.intrazone : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["taggings"] = state ? state.taggings : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemZoneArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["interfaces"] = args ? args.interfaces : undefined;
            inputs["intrazone"] = args ? args.intrazone : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["taggings"] = args ? args.taggings : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemZone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemZone resources.
 */
export interface SystemZoneState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.SystemZoneInterface>[]>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    intrazone?: pulumi.Input<string>;
    /**
     * Tag name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.SystemZoneTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemZone resource.
 */
export interface SystemZoneArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.SystemZoneInterface>[]>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    intrazone?: pulumi.Input<string>;
    /**
     * Tag name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.SystemZoneTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
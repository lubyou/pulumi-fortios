// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure roaming consortium.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WirelessControllerHotspot20AnqpRoamingConsortium("trname", {});
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerHotspot20AnqpRoamingConsortium extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerHotspot20AnqpRoamingConsortium resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerHotspot20AnqpRoamingConsortiumState, opts?: pulumi.CustomResourceOptions): WirelessControllerHotspot20AnqpRoamingConsortium {
        return new WirelessControllerHotspot20AnqpRoamingConsortium(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium';

    /**
     * Returns true if the given object is an instance of WirelessControllerHotspot20AnqpRoamingConsortium.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerHotspot20AnqpRoamingConsortium {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerHotspot20AnqpRoamingConsortium.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Roaming consortium name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    public readonly oiLists!: pulumi.Output<outputs.WirelessControllerHotspot20AnqpRoamingConsortiumOiList[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerHotspot20AnqpRoamingConsortium resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerHotspot20AnqpRoamingConsortiumArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerHotspot20AnqpRoamingConsortiumArgs | WirelessControllerHotspot20AnqpRoamingConsortiumState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerHotspot20AnqpRoamingConsortiumState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oiLists"] = state ? state.oiLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerHotspot20AnqpRoamingConsortiumArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oiLists"] = args ? args.oiLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerHotspot20AnqpRoamingConsortium.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerHotspot20AnqpRoamingConsortium resources.
 */
export interface WirelessControllerHotspot20AnqpRoamingConsortiumState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Roaming consortium name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    oiLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20AnqpRoamingConsortiumOiList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerHotspot20AnqpRoamingConsortium resource.
 */
export interface WirelessControllerHotspot20AnqpRoamingConsortiumArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Roaming consortium name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    oiLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20AnqpRoamingConsortiumOiList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

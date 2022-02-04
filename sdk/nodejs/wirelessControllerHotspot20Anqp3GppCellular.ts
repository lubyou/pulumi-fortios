// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure 3GPP public land mobile network (PLMN).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WirelessControllerHotspot20Anqp3GppCellular("trname", {});
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 Anqp3GppCellular can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerHotspot20Anqp3GppCellular:WirelessControllerHotspot20Anqp3GppCellular labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerHotspot20Anqp3GppCellular extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerHotspot20Anqp3GppCellular resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerHotspot20Anqp3GppCellularState, opts?: pulumi.CustomResourceOptions): WirelessControllerHotspot20Anqp3GppCellular {
        return new WirelessControllerHotspot20Anqp3GppCellular(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerHotspot20Anqp3GppCellular:WirelessControllerHotspot20Anqp3GppCellular';

    /**
     * Returns true if the given object is an instance of WirelessControllerHotspot20Anqp3GppCellular.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerHotspot20Anqp3GppCellular {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerHotspot20Anqp3GppCellular.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Mobile Country Code and Mobile Network Code configuration. The structure of `mccMncList` block is documented below.
     */
    public readonly mccMncLists!: pulumi.Output<outputs.WirelessControllerHotspot20Anqp3GppCellularMccMncList[] | undefined>;
    /**
     * 3GPP PLMN name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerHotspot20Anqp3GppCellular resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerHotspot20Anqp3GppCellularArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerHotspot20Anqp3GppCellularArgs | WirelessControllerHotspot20Anqp3GppCellularState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerHotspot20Anqp3GppCellularState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["mccMncLists"] = state ? state.mccMncLists : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerHotspot20Anqp3GppCellularArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["mccMncLists"] = args ? args.mccMncLists : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerHotspot20Anqp3GppCellular.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerHotspot20Anqp3GppCellular resources.
 */
export interface WirelessControllerHotspot20Anqp3GppCellularState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Mobile Country Code and Mobile Network Code configuration. The structure of `mccMncList` block is documented below.
     */
    mccMncLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20Anqp3GppCellularMccMncList>[]>;
    /**
     * 3GPP PLMN name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerHotspot20Anqp3GppCellular resource.
 */
export interface WirelessControllerHotspot20Anqp3GppCellularArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Mobile Country Code and Mobile Network Code configuration. The structure of `mccMncList` block is documented below.
     */
    mccMncLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20Anqp3GppCellularMccMncList>[]>;
    /**
     * 3GPP PLMN name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch location services.
 *
 * ## Import
 *
 * SwitchController Location can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerLocation:SwitchControllerLocation labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerLocation extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerLocation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerLocationState, opts?: pulumi.CustomResourceOptions): SwitchControllerLocation {
        return new SwitchControllerLocation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerLocation:SwitchControllerLocation';

    /**
     * Returns true if the given object is an instance of SwitchControllerLocation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerLocation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerLocation.__pulumiType;
    }

    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    public readonly addressCivic!: pulumi.Output<outputs.SwitchControllerLocationAddressCivic | undefined>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    public readonly coordinates!: pulumi.Output<outputs.SwitchControllerLocationCoordinates | undefined>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    public readonly elinNumber!: pulumi.Output<outputs.SwitchControllerLocationElinNumber | undefined>;
    /**
     * Name (residence and office occupant).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerLocation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerLocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerLocationArgs | SwitchControllerLocationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerLocationState | undefined;
            inputs["addressCivic"] = state ? state.addressCivic : undefined;
            inputs["coordinates"] = state ? state.coordinates : undefined;
            inputs["elinNumber"] = state ? state.elinNumber : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerLocationArgs | undefined;
            inputs["addressCivic"] = args ? args.addressCivic : undefined;
            inputs["coordinates"] = args ? args.coordinates : undefined;
            inputs["elinNumber"] = args ? args.elinNumber : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerLocation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerLocation resources.
 */
export interface SwitchControllerLocationState {
    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    addressCivic?: pulumi.Input<inputs.SwitchControllerLocationAddressCivic>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    coordinates?: pulumi.Input<inputs.SwitchControllerLocationCoordinates>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    elinNumber?: pulumi.Input<inputs.SwitchControllerLocationElinNumber>;
    /**
     * Name (residence and office occupant).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerLocation resource.
 */
export interface SwitchControllerLocationArgs {
    /**
     * Configure location civic address. The structure of `addressCivic` block is documented below.
     */
    addressCivic?: pulumi.Input<inputs.SwitchControllerLocationAddressCivic>;
    /**
     * Configure location GPS coordinates. The structure of `coordinates` block is documented below.
     */
    coordinates?: pulumi.Input<inputs.SwitchControllerLocationCoordinates>;
    /**
     * Configure location ELIN number. The structure of `elinNumber` block is documented below.
     */
    elinNumber?: pulumi.Input<inputs.SwitchControllerLocationElinNumber>;
    /**
     * Name (residence and office occupant).
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
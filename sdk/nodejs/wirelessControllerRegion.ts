// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiAP regions (for floor plans and maps).
 *
 * ## Import
 *
 * WirelessController Region can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerRegion:WirelessControllerRegion labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerRegion extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerRegion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerRegionState, opts?: pulumi.CustomResourceOptions): WirelessControllerRegion {
        return new WirelessControllerRegion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerRegion:WirelessControllerRegion';

    /**
     * Returns true if the given object is an instance of WirelessControllerRegion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerRegion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerRegion.__pulumiType;
    }

    /**
     * Comments.
     */
    public readonly comments!: pulumi.Output<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    public readonly grayscale!: pulumi.Output<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    public readonly imageType!: pulumi.Output<string>;
    /**
     * FortiAP region name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Region image opacity (0 - 100).
     */
    public readonly opacity!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerRegion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerRegionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerRegionArgs | WirelessControllerRegionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerRegionState | undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["grayscale"] = state ? state.grayscale : undefined;
            inputs["imageType"] = state ? state.imageType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["opacity"] = state ? state.opacity : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerRegionArgs | undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["grayscale"] = args ? args.grayscale : undefined;
            inputs["imageType"] = args ? args.imageType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["opacity"] = args ? args.opacity : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WirelessControllerRegion.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerRegion resources.
 */
export interface WirelessControllerRegionState {
    /**
     * Comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    grayscale?: pulumi.Input<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * FortiAP region name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region image opacity (0 - 100).
     */
    opacity?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerRegion resource.
 */
export interface WirelessControllerRegionArgs {
    /**
     * Comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Region image grayscale. Valid values: `enable`, `disable`.
     */
    grayscale?: pulumi.Input<string>;
    /**
     * FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
     */
    imageType?: pulumi.Input<string>;
    /**
     * FortiAP region name.
     */
    name?: pulumi.Input<string>;
    /**
     * Region image opacity (0 - 100).
     */
    opacity?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

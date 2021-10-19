// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPS URL filter settings for IPv6.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WebfilterIpsUrlfilterSetting6("trname", {
 *     distance: 1,
 *     gateway6: "::",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter IpsUrlfilterSetting6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterIpsUrlfilterSetting6:WebfilterIpsUrlfilterSetting6 labelname WebfilterIpsUrlfilterSetting6
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebfilterIpsUrlfilterSetting6 extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterIpsUrlfilterSetting6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterIpsUrlfilterSetting6State, opts?: pulumi.CustomResourceOptions): WebfilterIpsUrlfilterSetting6 {
        return new WebfilterIpsUrlfilterSetting6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterIpsUrlfilterSetting6:WebfilterIpsUrlfilterSetting6';

    /**
     * Returns true if the given object is an instance of WebfilterIpsUrlfilterSetting6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterIpsUrlfilterSetting6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterIpsUrlfilterSetting6.__pulumiType;
    }

    /**
     * Interface for this route.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * Gateway IPv6 address for this route.
     */
    public readonly gateway6!: pulumi.Output<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
     */
    public readonly geoFilter!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebfilterIpsUrlfilterSetting6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebfilterIpsUrlfilterSetting6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterIpsUrlfilterSetting6Args | WebfilterIpsUrlfilterSetting6State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterIpsUrlfilterSetting6State | undefined;
            inputs["device"] = state ? state.device : undefined;
            inputs["distance"] = state ? state.distance : undefined;
            inputs["gateway6"] = state ? state.gateway6 : undefined;
            inputs["geoFilter"] = state ? state.geoFilter : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebfilterIpsUrlfilterSetting6Args | undefined;
            inputs["device"] = args ? args.device : undefined;
            inputs["distance"] = args ? args.distance : undefined;
            inputs["gateway6"] = args ? args.gateway6 : undefined;
            inputs["geoFilter"] = args ? args.geoFilter : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebfilterIpsUrlfilterSetting6.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterIpsUrlfilterSetting6 resources.
 */
export interface WebfilterIpsUrlfilterSetting6State {
    /**
     * Interface for this route.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Gateway IPv6 address for this route.
     */
    gateway6?: pulumi.Input<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
     */
    geoFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterIpsUrlfilterSetting6 resource.
 */
export interface WebfilterIpsUrlfilterSetting6Args {
    /**
     * Interface for this route.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255) for this route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Gateway IPv6 address for this route.
     */
    gateway6?: pulumi.Input<string>;
    /**
     * Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
     */
    geoFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

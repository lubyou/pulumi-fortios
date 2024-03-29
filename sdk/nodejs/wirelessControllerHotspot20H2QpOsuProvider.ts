// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerHotspot20H2QpOsuProvider extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerHotspot20H2QpOsuProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerHotspot20H2QpOsuProviderState, opts?: pulumi.CustomResourceOptions): WirelessControllerHotspot20H2QpOsuProvider {
        return new WirelessControllerHotspot20H2QpOsuProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerHotspot20H2QpOsuProvider:WirelessControllerHotspot20H2QpOsuProvider';

    /**
     * Returns true if the given object is an instance of WirelessControllerHotspot20H2QpOsuProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerHotspot20H2QpOsuProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerHotspot20H2QpOsuProvider.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly friendlyNames!: pulumi.Output<outputs.WirelessControllerHotspot20H2QpOsuProviderFriendlyName[] | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly icon!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly osuMethod!: pulumi.Output<string>;
    public readonly osuNai!: pulumi.Output<string>;
    public readonly serverUri!: pulumi.Output<string>;
    public readonly serviceDescriptions!: pulumi.Output<outputs.WirelessControllerHotspot20H2QpOsuProviderServiceDescription[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerHotspot20H2QpOsuProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerHotspot20H2QpOsuProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerHotspot20H2QpOsuProviderArgs | WirelessControllerHotspot20H2QpOsuProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerHotspot20H2QpOsuProviderState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["friendlyNames"] = state ? state.friendlyNames : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["icon"] = state ? state.icon : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["osuMethod"] = state ? state.osuMethod : undefined;
            resourceInputs["osuNai"] = state ? state.osuNai : undefined;
            resourceInputs["serverUri"] = state ? state.serverUri : undefined;
            resourceInputs["serviceDescriptions"] = state ? state.serviceDescriptions : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerHotspot20H2QpOsuProviderArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["friendlyNames"] = args ? args.friendlyNames : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["icon"] = args ? args.icon : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osuMethod"] = args ? args.osuMethod : undefined;
            resourceInputs["osuNai"] = args ? args.osuNai : undefined;
            resourceInputs["serverUri"] = args ? args.serverUri : undefined;
            resourceInputs["serviceDescriptions"] = args ? args.serviceDescriptions : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerHotspot20H2QpOsuProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerHotspot20H2QpOsuProvider resources.
 */
export interface WirelessControllerHotspot20H2QpOsuProviderState {
    dynamicSortSubtable?: pulumi.Input<string>;
    friendlyNames?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20H2QpOsuProviderFriendlyName>[]>;
    getAllTables?: pulumi.Input<string>;
    icon?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    osuMethod?: pulumi.Input<string>;
    osuNai?: pulumi.Input<string>;
    serverUri?: pulumi.Input<string>;
    serviceDescriptions?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20H2QpOsuProviderServiceDescription>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerHotspot20H2QpOsuProvider resource.
 */
export interface WirelessControllerHotspot20H2QpOsuProviderArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    friendlyNames?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20H2QpOsuProviderFriendlyName>[]>;
    getAllTables?: pulumi.Input<string>;
    icon?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    osuMethod?: pulumi.Input<string>;
    osuNai?: pulumi.Input<string>;
    serverUri?: pulumi.Input<string>;
    serviceDescriptions?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20H2QpOsuProviderServiceDescription>[]>;
    vdomparam?: pulumi.Input<string>;
}

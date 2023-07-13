// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerAccessControlList extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerAccessControlList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerAccessControlListState, opts?: pulumi.CustomResourceOptions): WirelessControllerAccessControlList {
        return new WirelessControllerAccessControlList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerAccessControlList:WirelessControllerAccessControlList';

    /**
     * Returns true if the given object is an instance of WirelessControllerAccessControlList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerAccessControlList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerAccessControlList.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly layer3Ipv4Rules!: pulumi.Output<outputs.WirelessControllerAccessControlListLayer3Ipv4Rule[] | undefined>;
    public readonly layer3Ipv6Rules!: pulumi.Output<outputs.WirelessControllerAccessControlListLayer3Ipv6Rule[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerAccessControlList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerAccessControlListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerAccessControlListArgs | WirelessControllerAccessControlListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerAccessControlListState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["layer3Ipv4Rules"] = state ? state.layer3Ipv4Rules : undefined;
            resourceInputs["layer3Ipv6Rules"] = state ? state.layer3Ipv6Rules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerAccessControlListArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["layer3Ipv4Rules"] = args ? args.layer3Ipv4Rules : undefined;
            resourceInputs["layer3Ipv6Rules"] = args ? args.layer3Ipv6Rules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerAccessControlList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerAccessControlList resources.
 */
export interface WirelessControllerAccessControlListState {
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    layer3Ipv4Rules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAccessControlListLayer3Ipv4Rule>[]>;
    layer3Ipv6Rules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAccessControlListLayer3Ipv6Rule>[]>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerAccessControlList resource.
 */
export interface WirelessControllerAccessControlListArgs {
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    layer3Ipv4Rules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAccessControlListLayer3Ipv4Rule>[]>;
    layer3Ipv6Rules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerAccessControlListLayer3Ipv6Rule>[]>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

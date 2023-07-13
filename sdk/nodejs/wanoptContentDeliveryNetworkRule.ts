// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WanoptContentDeliveryNetworkRule extends pulumi.CustomResource {
    /**
     * Get an existing WanoptContentDeliveryNetworkRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WanoptContentDeliveryNetworkRuleState, opts?: pulumi.CustomResourceOptions): WanoptContentDeliveryNetworkRule {
        return new WanoptContentDeliveryNetworkRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wanoptContentDeliveryNetworkRule:WanoptContentDeliveryNetworkRule';

    /**
     * Returns true if the given object is an instance of WanoptContentDeliveryNetworkRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WanoptContentDeliveryNetworkRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WanoptContentDeliveryNetworkRule.__pulumiType;
    }

    public readonly category!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly hostDomainNameSuffixes!: pulumi.Output<outputs.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly requestCacheControl!: pulumi.Output<string>;
    public readonly responseCacheControl!: pulumi.Output<string>;
    public readonly responseExpires!: pulumi.Output<string>;
    public readonly rules!: pulumi.Output<outputs.WanoptContentDeliveryNetworkRuleRule[] | undefined>;
    public readonly status!: pulumi.Output<string>;
    public readonly textResponseVcache!: pulumi.Output<string>;
    public readonly updateserver!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WanoptContentDeliveryNetworkRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WanoptContentDeliveryNetworkRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WanoptContentDeliveryNetworkRuleArgs | WanoptContentDeliveryNetworkRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WanoptContentDeliveryNetworkRuleState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["hostDomainNameSuffixes"] = state ? state.hostDomainNameSuffixes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["requestCacheControl"] = state ? state.requestCacheControl : undefined;
            resourceInputs["responseCacheControl"] = state ? state.responseCacheControl : undefined;
            resourceInputs["responseExpires"] = state ? state.responseExpires : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["textResponseVcache"] = state ? state.textResponseVcache : undefined;
            resourceInputs["updateserver"] = state ? state.updateserver : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WanoptContentDeliveryNetworkRuleArgs | undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["hostDomainNameSuffixes"] = args ? args.hostDomainNameSuffixes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["requestCacheControl"] = args ? args.requestCacheControl : undefined;
            resourceInputs["responseCacheControl"] = args ? args.responseCacheControl : undefined;
            resourceInputs["responseExpires"] = args ? args.responseExpires : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["textResponseVcache"] = args ? args.textResponseVcache : undefined;
            resourceInputs["updateserver"] = args ? args.updateserver : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WanoptContentDeliveryNetworkRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WanoptContentDeliveryNetworkRule resources.
 */
export interface WanoptContentDeliveryNetworkRuleState {
    category?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    hostDomainNameSuffixes?: pulumi.Input<pulumi.Input<inputs.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix>[]>;
    name?: pulumi.Input<string>;
    requestCacheControl?: pulumi.Input<string>;
    responseCacheControl?: pulumi.Input<string>;
    responseExpires?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.WanoptContentDeliveryNetworkRuleRule>[]>;
    status?: pulumi.Input<string>;
    textResponseVcache?: pulumi.Input<string>;
    updateserver?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WanoptContentDeliveryNetworkRule resource.
 */
export interface WanoptContentDeliveryNetworkRuleArgs {
    category?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    hostDomainNameSuffixes?: pulumi.Input<pulumi.Input<inputs.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix>[]>;
    name?: pulumi.Input<string>;
    requestCacheControl?: pulumi.Input<string>;
    responseCacheControl?: pulumi.Input<string>;
    responseExpires?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.WanoptContentDeliveryNetworkRuleRule>[]>;
    status?: pulumi.Input<string>;
    textResponseVcache?: pulumi.Input<string>;
    updateserver?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

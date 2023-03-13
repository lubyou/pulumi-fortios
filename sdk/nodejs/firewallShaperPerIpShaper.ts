// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallShaperPerIpShaper extends pulumi.CustomResource {
    /**
     * Get an existing FirewallShaperPerIpShaper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallShaperPerIpShaperState, opts?: pulumi.CustomResourceOptions): FirewallShaperPerIpShaper {
        return new FirewallShaperPerIpShaper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper';

    /**
     * Returns true if the given object is an instance of FirewallShaperPerIpShaper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallShaperPerIpShaper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallShaperPerIpShaper.__pulumiType;
    }

    public readonly bandwidthUnit!: pulumi.Output<string>;
    public readonly diffservForward!: pulumi.Output<string>;
    public readonly diffservReverse!: pulumi.Output<string>;
    public readonly diffservcodeForward!: pulumi.Output<string>;
    public readonly diffservcodeRev!: pulumi.Output<string>;
    public readonly maxBandwidth!: pulumi.Output<number>;
    public readonly maxConcurrentSession!: pulumi.Output<number>;
    public readonly maxConcurrentTcpSession!: pulumi.Output<number>;
    public readonly maxConcurrentUdpSession!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallShaperPerIpShaper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallShaperPerIpShaperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallShaperPerIpShaperArgs | FirewallShaperPerIpShaperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallShaperPerIpShaperState | undefined;
            resourceInputs["bandwidthUnit"] = state ? state.bandwidthUnit : undefined;
            resourceInputs["diffservForward"] = state ? state.diffservForward : undefined;
            resourceInputs["diffservReverse"] = state ? state.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = state ? state.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = state ? state.diffservcodeRev : undefined;
            resourceInputs["maxBandwidth"] = state ? state.maxBandwidth : undefined;
            resourceInputs["maxConcurrentSession"] = state ? state.maxConcurrentSession : undefined;
            resourceInputs["maxConcurrentTcpSession"] = state ? state.maxConcurrentTcpSession : undefined;
            resourceInputs["maxConcurrentUdpSession"] = state ? state.maxConcurrentUdpSession : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallShaperPerIpShaperArgs | undefined;
            resourceInputs["bandwidthUnit"] = args ? args.bandwidthUnit : undefined;
            resourceInputs["diffservForward"] = args ? args.diffservForward : undefined;
            resourceInputs["diffservReverse"] = args ? args.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = args ? args.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = args ? args.diffservcodeRev : undefined;
            resourceInputs["maxBandwidth"] = args ? args.maxBandwidth : undefined;
            resourceInputs["maxConcurrentSession"] = args ? args.maxConcurrentSession : undefined;
            resourceInputs["maxConcurrentTcpSession"] = args ? args.maxConcurrentTcpSession : undefined;
            resourceInputs["maxConcurrentUdpSession"] = args ? args.maxConcurrentUdpSession : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallShaperPerIpShaper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallShaperPerIpShaper resources.
 */
export interface FirewallShaperPerIpShaperState {
    bandwidthUnit?: pulumi.Input<string>;
    diffservForward?: pulumi.Input<string>;
    diffservReverse?: pulumi.Input<string>;
    diffservcodeForward?: pulumi.Input<string>;
    diffservcodeRev?: pulumi.Input<string>;
    maxBandwidth?: pulumi.Input<number>;
    maxConcurrentSession?: pulumi.Input<number>;
    maxConcurrentTcpSession?: pulumi.Input<number>;
    maxConcurrentUdpSession?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallShaperPerIpShaper resource.
 */
export interface FirewallShaperPerIpShaperArgs {
    bandwidthUnit?: pulumi.Input<string>;
    diffservForward?: pulumi.Input<string>;
    diffservReverse?: pulumi.Input<string>;
    diffservcodeForward?: pulumi.Input<string>;
    diffservcodeRev?: pulumi.Input<string>;
    maxBandwidth?: pulumi.Input<number>;
    maxConcurrentSession?: pulumi.Input<number>;
    maxConcurrentTcpSession?: pulumi.Input<number>;
    maxConcurrentUdpSession?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

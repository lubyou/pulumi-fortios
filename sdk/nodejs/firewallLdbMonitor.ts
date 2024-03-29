// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallLdbMonitor extends pulumi.CustomResource {
    /**
     * Get an existing FirewallLdbMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallLdbMonitorState, opts?: pulumi.CustomResourceOptions): FirewallLdbMonitor {
        return new FirewallLdbMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallLdbMonitor:FirewallLdbMonitor';

    /**
     * Returns true if the given object is an instance of FirewallLdbMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallLdbMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallLdbMonitor.__pulumiType;
    }

    public readonly dnsMatchIp!: pulumi.Output<string>;
    public readonly dnsProtocol!: pulumi.Output<string>;
    public readonly dnsRequestDomain!: pulumi.Output<string>;
    public readonly httpGet!: pulumi.Output<string>;
    public readonly httpMatch!: pulumi.Output<string>;
    public readonly httpMaxRedirects!: pulumi.Output<number>;
    public readonly interval!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number>;
    public readonly retry!: pulumi.Output<number>;
    public readonly srcIp!: pulumi.Output<string>;
    public readonly timeout!: pulumi.Output<number>;
    public readonly type!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallLdbMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallLdbMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallLdbMonitorArgs | FirewallLdbMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallLdbMonitorState | undefined;
            resourceInputs["dnsMatchIp"] = state ? state.dnsMatchIp : undefined;
            resourceInputs["dnsProtocol"] = state ? state.dnsProtocol : undefined;
            resourceInputs["dnsRequestDomain"] = state ? state.dnsRequestDomain : undefined;
            resourceInputs["httpGet"] = state ? state.httpGet : undefined;
            resourceInputs["httpMatch"] = state ? state.httpMatch : undefined;
            resourceInputs["httpMaxRedirects"] = state ? state.httpMaxRedirects : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["retry"] = state ? state.retry : undefined;
            resourceInputs["srcIp"] = state ? state.srcIp : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallLdbMonitorArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dnsMatchIp"] = args ? args.dnsMatchIp : undefined;
            resourceInputs["dnsProtocol"] = args ? args.dnsProtocol : undefined;
            resourceInputs["dnsRequestDomain"] = args ? args.dnsRequestDomain : undefined;
            resourceInputs["httpGet"] = args ? args.httpGet : undefined;
            resourceInputs["httpMatch"] = args ? args.httpMatch : undefined;
            resourceInputs["httpMaxRedirects"] = args ? args.httpMaxRedirects : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["retry"] = args ? args.retry : undefined;
            resourceInputs["srcIp"] = args ? args.srcIp : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallLdbMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallLdbMonitor resources.
 */
export interface FirewallLdbMonitorState {
    dnsMatchIp?: pulumi.Input<string>;
    dnsProtocol?: pulumi.Input<string>;
    dnsRequestDomain?: pulumi.Input<string>;
    httpGet?: pulumi.Input<string>;
    httpMatch?: pulumi.Input<string>;
    httpMaxRedirects?: pulumi.Input<number>;
    interval?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    retry?: pulumi.Input<number>;
    srcIp?: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallLdbMonitor resource.
 */
export interface FirewallLdbMonitorArgs {
    dnsMatchIp?: pulumi.Input<string>;
    dnsProtocol?: pulumi.Input<string>;
    dnsRequestDomain?: pulumi.Input<string>;
    httpGet?: pulumi.Input<string>;
    httpMatch?: pulumi.Input<string>;
    httpMaxRedirects?: pulumi.Input<number>;
    interval?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    retry?: pulumi.Input<number>;
    srcIp?: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
    type: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

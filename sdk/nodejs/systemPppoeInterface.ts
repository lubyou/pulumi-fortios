// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemPppoeInterface extends pulumi.CustomResource {
    /**
     * Get an existing SystemPppoeInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemPppoeInterfaceState, opts?: pulumi.CustomResourceOptions): SystemPppoeInterface {
        return new SystemPppoeInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemPppoeInterface:SystemPppoeInterface';

    /**
     * Returns true if the given object is an instance of SystemPppoeInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemPppoeInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemPppoeInterface.__pulumiType;
    }

    public readonly acName!: pulumi.Output<string>;
    public readonly authType!: pulumi.Output<string>;
    public readonly device!: pulumi.Output<string>;
    public readonly dialOnDemand!: pulumi.Output<string>;
    public readonly discRetryTimeout!: pulumi.Output<number>;
    public readonly idleTimeout!: pulumi.Output<number>;
    public readonly ipunnumbered!: pulumi.Output<string>;
    public readonly ipv6!: pulumi.Output<string>;
    public readonly lcpEchoInterval!: pulumi.Output<number>;
    public readonly lcpMaxEchoFails!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly padtRetryTimeout!: pulumi.Output<number>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly pppoeUnnumberedNegotiate!: pulumi.Output<string>;
    public readonly serviceName!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemPppoeInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemPppoeInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemPppoeInterfaceArgs | SystemPppoeInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemPppoeInterfaceState | undefined;
            resourceInputs["acName"] = state ? state.acName : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["dialOnDemand"] = state ? state.dialOnDemand : undefined;
            resourceInputs["discRetryTimeout"] = state ? state.discRetryTimeout : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["ipunnumbered"] = state ? state.ipunnumbered : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["lcpEchoInterval"] = state ? state.lcpEchoInterval : undefined;
            resourceInputs["lcpMaxEchoFails"] = state ? state.lcpMaxEchoFails : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["padtRetryTimeout"] = state ? state.padtRetryTimeout : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["pppoeUnnumberedNegotiate"] = state ? state.pppoeUnnumberedNegotiate : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemPppoeInterfaceArgs | undefined;
            if ((!args || args.device === undefined) && !opts.urn) {
                throw new Error("Missing required property 'device'");
            }
            resourceInputs["acName"] = args ? args.acName : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["dialOnDemand"] = args ? args.dialOnDemand : undefined;
            resourceInputs["discRetryTimeout"] = args ? args.discRetryTimeout : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["ipunnumbered"] = args ? args.ipunnumbered : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["lcpEchoInterval"] = args ? args.lcpEchoInterval : undefined;
            resourceInputs["lcpMaxEchoFails"] = args ? args.lcpMaxEchoFails : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["padtRetryTimeout"] = args ? args.padtRetryTimeout : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["pppoeUnnumberedNegotiate"] = args ? args.pppoeUnnumberedNegotiate : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemPppoeInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemPppoeInterface resources.
 */
export interface SystemPppoeInterfaceState {
    acName?: pulumi.Input<string>;
    authType?: pulumi.Input<string>;
    device?: pulumi.Input<string>;
    dialOnDemand?: pulumi.Input<string>;
    discRetryTimeout?: pulumi.Input<number>;
    idleTimeout?: pulumi.Input<number>;
    ipunnumbered?: pulumi.Input<string>;
    ipv6?: pulumi.Input<string>;
    lcpEchoInterval?: pulumi.Input<number>;
    lcpMaxEchoFails?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    padtRetryTimeout?: pulumi.Input<number>;
    password?: pulumi.Input<string>;
    pppoeUnnumberedNegotiate?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemPppoeInterface resource.
 */
export interface SystemPppoeInterfaceArgs {
    acName?: pulumi.Input<string>;
    authType?: pulumi.Input<string>;
    device: pulumi.Input<string>;
    dialOnDemand?: pulumi.Input<string>;
    discRetryTimeout?: pulumi.Input<number>;
    idleTimeout?: pulumi.Input<number>;
    ipunnumbered?: pulumi.Input<string>;
    ipv6?: pulumi.Input<string>;
    lcpEchoInterval?: pulumi.Input<number>;
    lcpMaxEchoFails?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    padtRetryTimeout?: pulumi.Input<number>;
    password?: pulumi.Input<string>;
    pppoeUnnumberedNegotiate?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemProbeResponse extends pulumi.CustomResource {
    /**
     * Get an existing SystemProbeResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemProbeResponseState, opts?: pulumi.CustomResourceOptions): SystemProbeResponse {
        return new SystemProbeResponse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemProbeResponse:SystemProbeResponse';

    /**
     * Returns true if the given object is an instance of SystemProbeResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemProbeResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemProbeResponse.__pulumiType;
    }

    public readonly httpProbeValue!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number>;
    public readonly securityMode!: pulumi.Output<string>;
    public readonly timeout!: pulumi.Output<number>;
    public readonly ttlMode!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemProbeResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemProbeResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemProbeResponseArgs | SystemProbeResponseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemProbeResponseState | undefined;
            resourceInputs["httpProbeValue"] = state ? state.httpProbeValue : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["securityMode"] = state ? state.securityMode : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["ttlMode"] = state ? state.ttlMode : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemProbeResponseArgs | undefined;
            resourceInputs["httpProbeValue"] = args ? args.httpProbeValue : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["securityMode"] = args ? args.securityMode : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["ttlMode"] = args ? args.ttlMode : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemProbeResponse.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemProbeResponse resources.
 */
export interface SystemProbeResponseState {
    httpProbeValue?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    securityMode?: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
    ttlMode?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemProbeResponse resource.
 */
export interface SystemProbeResponseArgs {
    httpProbeValue?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    securityMode?: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
    ttlMode?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class AuthenticationScheme extends pulumi.CustomResource {
    /**
     * Get an existing AuthenticationScheme resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthenticationSchemeState, opts?: pulumi.CustomResourceOptions): AuthenticationScheme {
        return new AuthenticationScheme(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/authenticationScheme:AuthenticationScheme';

    /**
     * Returns true if the given object is an instance of AuthenticationScheme.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthenticationScheme {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthenticationScheme.__pulumiType;
    }

    public readonly domainController!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fssoAgentForNtlm!: pulumi.Output<string>;
    public readonly fssoGuest!: pulumi.Output<string>;
    public readonly kerberosKeytab!: pulumi.Output<string>;
    public readonly method!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly negotiateNtlm!: pulumi.Output<string>;
    public readonly requireTfa!: pulumi.Output<string>;
    public readonly samlServer!: pulumi.Output<string>;
    public readonly samlTimeout!: pulumi.Output<number>;
    public readonly sshCa!: pulumi.Output<string>;
    public readonly userCert!: pulumi.Output<string>;
    public readonly userDatabases!: pulumi.Output<outputs.AuthenticationSchemeUserDatabase[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthenticationScheme resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthenticationSchemeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthenticationSchemeArgs | AuthenticationSchemeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthenticationSchemeState | undefined;
            resourceInputs["domainController"] = state ? state.domainController : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fssoAgentForNtlm"] = state ? state.fssoAgentForNtlm : undefined;
            resourceInputs["fssoGuest"] = state ? state.fssoGuest : undefined;
            resourceInputs["kerberosKeytab"] = state ? state.kerberosKeytab : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["negotiateNtlm"] = state ? state.negotiateNtlm : undefined;
            resourceInputs["requireTfa"] = state ? state.requireTfa : undefined;
            resourceInputs["samlServer"] = state ? state.samlServer : undefined;
            resourceInputs["samlTimeout"] = state ? state.samlTimeout : undefined;
            resourceInputs["sshCa"] = state ? state.sshCa : undefined;
            resourceInputs["userCert"] = state ? state.userCert : undefined;
            resourceInputs["userDatabases"] = state ? state.userDatabases : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AuthenticationSchemeArgs | undefined;
            if ((!args || args.method === undefined) && !opts.urn) {
                throw new Error("Missing required property 'method'");
            }
            resourceInputs["domainController"] = args ? args.domainController : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fssoAgentForNtlm"] = args ? args.fssoAgentForNtlm : undefined;
            resourceInputs["fssoGuest"] = args ? args.fssoGuest : undefined;
            resourceInputs["kerberosKeytab"] = args ? args.kerberosKeytab : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["negotiateNtlm"] = args ? args.negotiateNtlm : undefined;
            resourceInputs["requireTfa"] = args ? args.requireTfa : undefined;
            resourceInputs["samlServer"] = args ? args.samlServer : undefined;
            resourceInputs["samlTimeout"] = args ? args.samlTimeout : undefined;
            resourceInputs["sshCa"] = args ? args.sshCa : undefined;
            resourceInputs["userCert"] = args ? args.userCert : undefined;
            resourceInputs["userDatabases"] = args ? args.userDatabases : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthenticationScheme.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthenticationScheme resources.
 */
export interface AuthenticationSchemeState {
    domainController?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fssoAgentForNtlm?: pulumi.Input<string>;
    fssoGuest?: pulumi.Input<string>;
    kerberosKeytab?: pulumi.Input<string>;
    method?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    negotiateNtlm?: pulumi.Input<string>;
    requireTfa?: pulumi.Input<string>;
    samlServer?: pulumi.Input<string>;
    samlTimeout?: pulumi.Input<number>;
    sshCa?: pulumi.Input<string>;
    userCert?: pulumi.Input<string>;
    userDatabases?: pulumi.Input<pulumi.Input<inputs.AuthenticationSchemeUserDatabase>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthenticationScheme resource.
 */
export interface AuthenticationSchemeArgs {
    domainController?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fssoAgentForNtlm?: pulumi.Input<string>;
    fssoGuest?: pulumi.Input<string>;
    kerberosKeytab?: pulumi.Input<string>;
    method: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    negotiateNtlm?: pulumi.Input<string>;
    requireTfa?: pulumi.Input<string>;
    samlServer?: pulumi.Input<string>;
    samlTimeout?: pulumi.Input<number>;
    sshCa?: pulumi.Input<string>;
    userCert?: pulumi.Input<string>;
    userDatabases?: pulumi.Input<pulumi.Input<inputs.AuthenticationSchemeUserDatabase>[]>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SSH proxy local keys.
 *
 * ## Import
 *
 * FirewallSsh LocalKey can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSshLocalKey:FirewallSshLocalKey labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallSshLocalKey extends pulumi.CustomResource {
    /**
     * Get an existing FirewallSshLocalKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallSshLocalKeyState, opts?: pulumi.CustomResourceOptions): FirewallSshLocalKey {
        return new FirewallSshLocalKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallSshLocalKey:FirewallSshLocalKey';

    /**
     * Returns true if the given object is an instance of FirewallSshLocalKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallSshLocalKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallSshLocalKey.__pulumiType;
    }

    /**
     * SSH proxy local key name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for SSH private key.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * SSH proxy public key.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * SSH proxy local key source type. Valid values: `built-in`, `user`.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallSshLocalKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallSshLocalKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallSshLocalKeyArgs | FirewallSshLocalKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallSshLocalKeyState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["publicKey"] = state ? state.publicKey : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallSshLocalKeyArgs | undefined;
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["publicKey"] = args ? args.publicKey : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallSshLocalKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallSshLocalKey resources.
 */
export interface FirewallSshLocalKeyState {
    /**
     * SSH proxy local key name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for SSH private key.
     */
    password?: pulumi.Input<string>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * SSH proxy public key.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * SSH proxy local key source type. Valid values: `built-in`, `user`.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallSshLocalKey resource.
 */
export interface FirewallSshLocalKeyArgs {
    /**
     * SSH proxy local key name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for SSH private key.
     */
    password?: pulumi.Input<string>;
    /**
     * SSH proxy private key, encrypted with a password.
     */
    privateKey: pulumi.Input<string>;
    /**
     * SSH proxy public key.
     */
    publicKey: pulumi.Input<string>;
    /**
     * SSH proxy local key source type. Valid values: `built-in`, `user`.
     */
    source?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
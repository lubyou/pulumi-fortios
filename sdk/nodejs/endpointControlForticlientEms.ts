// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EndpointControlForticlientEms extends pulumi.CustomResource {
    /**
     * Get an existing EndpointControlForticlientEms resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointControlForticlientEmsState, opts?: pulumi.CustomResourceOptions): EndpointControlForticlientEms {
        return new EndpointControlForticlientEms(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/endpointControlForticlientEms:EndpointControlForticlientEms';

    /**
     * Returns true if the given object is an instance of EndpointControlForticlientEms.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointControlForticlientEms {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointControlForticlientEms.__pulumiType;
    }

    public readonly address!: pulumi.Output<string>;
    public readonly adminPassword!: pulumi.Output<string | undefined>;
    public readonly adminType!: pulumi.Output<string>;
    public readonly adminUsername!: pulumi.Output<string>;
    public readonly httpsPort!: pulumi.Output<number>;
    public readonly listenPort!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly restApiAuth!: pulumi.Output<string>;
    public readonly serialNumber!: pulumi.Output<string>;
    public readonly uploadPort!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EndpointControlForticlientEms resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointControlForticlientEmsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointControlForticlientEmsArgs | EndpointControlForticlientEmsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointControlForticlientEmsState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["adminType"] = state ? state.adminType : undefined;
            resourceInputs["adminUsername"] = state ? state.adminUsername : undefined;
            resourceInputs["httpsPort"] = state ? state.httpsPort : undefined;
            resourceInputs["listenPort"] = state ? state.listenPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["restApiAuth"] = state ? state.restApiAuth : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["uploadPort"] = state ? state.uploadPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EndpointControlForticlientEmsArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.adminUsername === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminUsername'");
            }
            if ((!args || args.serialNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serialNumber'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["adminPassword"] = args?.adminPassword ? pulumi.secret(args.adminPassword) : undefined;
            resourceInputs["adminType"] = args ? args.adminType : undefined;
            resourceInputs["adminUsername"] = args ? args.adminUsername : undefined;
            resourceInputs["httpsPort"] = args ? args.httpsPort : undefined;
            resourceInputs["listenPort"] = args ? args.listenPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["restApiAuth"] = args ? args.restApiAuth : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["uploadPort"] = args ? args.uploadPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EndpointControlForticlientEms.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointControlForticlientEms resources.
 */
export interface EndpointControlForticlientEmsState {
    address?: pulumi.Input<string>;
    adminPassword?: pulumi.Input<string>;
    adminType?: pulumi.Input<string>;
    adminUsername?: pulumi.Input<string>;
    httpsPort?: pulumi.Input<number>;
    listenPort?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    restApiAuth?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    uploadPort?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointControlForticlientEms resource.
 */
export interface EndpointControlForticlientEmsArgs {
    address: pulumi.Input<string>;
    adminPassword?: pulumi.Input<string>;
    adminType?: pulumi.Input<string>;
    adminUsername: pulumi.Input<string>;
    httpsPort?: pulumi.Input<number>;
    listenPort?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    restApiAuth?: pulumi.Input<string>;
    serialNumber: pulumi.Input<string>;
    uploadPort?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

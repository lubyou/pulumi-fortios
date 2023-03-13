// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EmailfilterFortishield extends pulumi.CustomResource {
    /**
     * Get an existing EmailfilterFortishield resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailfilterFortishieldState, opts?: pulumi.CustomResourceOptions): EmailfilterFortishield {
        return new EmailfilterFortishield(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/emailfilterFortishield:EmailfilterFortishield';

    /**
     * Returns true if the given object is an instance of EmailfilterFortishield.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailfilterFortishield {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailfilterFortishield.__pulumiType;
    }

    public readonly spamSubmitForce!: pulumi.Output<string>;
    public readonly spamSubmitSrv!: pulumi.Output<string>;
    public readonly spamSubmitTxt2htm!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailfilterFortishield resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EmailfilterFortishieldArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailfilterFortishieldArgs | EmailfilterFortishieldState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailfilterFortishieldState | undefined;
            resourceInputs["spamSubmitForce"] = state ? state.spamSubmitForce : undefined;
            resourceInputs["spamSubmitSrv"] = state ? state.spamSubmitSrv : undefined;
            resourceInputs["spamSubmitTxt2htm"] = state ? state.spamSubmitTxt2htm : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EmailfilterFortishieldArgs | undefined;
            resourceInputs["spamSubmitForce"] = args ? args.spamSubmitForce : undefined;
            resourceInputs["spamSubmitSrv"] = args ? args.spamSubmitSrv : undefined;
            resourceInputs["spamSubmitTxt2htm"] = args ? args.spamSubmitTxt2htm : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmailfilterFortishield.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailfilterFortishield resources.
 */
export interface EmailfilterFortishieldState {
    spamSubmitForce?: pulumi.Input<string>;
    spamSubmitSrv?: pulumi.Input<string>;
    spamSubmitTxt2htm?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailfilterFortishield resource.
 */
export interface EmailfilterFortishieldArgs {
    spamSubmitForce?: pulumi.Input<string>;
    spamSubmitSrv?: pulumi.Input<string>;
    spamSubmitTxt2htm?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

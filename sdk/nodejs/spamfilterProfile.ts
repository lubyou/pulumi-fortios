// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SpamfilterProfile extends pulumi.CustomResource {
    /**
     * Get an existing SpamfilterProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpamfilterProfileState, opts?: pulumi.CustomResourceOptions): SpamfilterProfile {
        return new SpamfilterProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/spamfilterProfile:SpamfilterProfile';

    /**
     * Returns true if the given object is an instance of SpamfilterProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpamfilterProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpamfilterProfile.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly external!: pulumi.Output<string>;
    public readonly flowBased!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly gmail!: pulumi.Output<outputs.SpamfilterProfileGmail>;
    public readonly imap!: pulumi.Output<outputs.SpamfilterProfileImap>;
    public readonly mapi!: pulumi.Output<outputs.SpamfilterProfileMapi>;
    public readonly msnHotmail!: pulumi.Output<outputs.SpamfilterProfileMsnHotmail>;
    public readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<string>;
    public readonly pop3!: pulumi.Output<outputs.SpamfilterProfilePop3>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly smtp!: pulumi.Output<outputs.SpamfilterProfileSmtp>;
    public readonly spamBwlTable!: pulumi.Output<number>;
    public readonly spamBwordTable!: pulumi.Output<number>;
    public readonly spamBwordThreshold!: pulumi.Output<number>;
    public readonly spamFiltering!: pulumi.Output<string>;
    public readonly spamIptrustTable!: pulumi.Output<number>;
    public readonly spamLog!: pulumi.Output<string>;
    public readonly spamLogFortiguardResponse!: pulumi.Output<string>;
    public readonly spamMheaderTable!: pulumi.Output<number>;
    public readonly spamRblTable!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly yahooMail!: pulumi.Output<outputs.SpamfilterProfileYahooMail>;

    /**
     * Create a SpamfilterProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SpamfilterProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpamfilterProfileArgs | SpamfilterProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpamfilterProfileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["flowBased"] = state ? state.flowBased : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gmail"] = state ? state.gmail : undefined;
            resourceInputs["imap"] = state ? state.imap : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["msnHotmail"] = state ? state.msnHotmail : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["pop3"] = state ? state.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["smtp"] = state ? state.smtp : undefined;
            resourceInputs["spamBwlTable"] = state ? state.spamBwlTable : undefined;
            resourceInputs["spamBwordTable"] = state ? state.spamBwordTable : undefined;
            resourceInputs["spamBwordThreshold"] = state ? state.spamBwordThreshold : undefined;
            resourceInputs["spamFiltering"] = state ? state.spamFiltering : undefined;
            resourceInputs["spamIptrustTable"] = state ? state.spamIptrustTable : undefined;
            resourceInputs["spamLog"] = state ? state.spamLog : undefined;
            resourceInputs["spamLogFortiguardResponse"] = state ? state.spamLogFortiguardResponse : undefined;
            resourceInputs["spamMheaderTable"] = state ? state.spamMheaderTable : undefined;
            resourceInputs["spamRblTable"] = state ? state.spamRblTable : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["yahooMail"] = state ? state.yahooMail : undefined;
        } else {
            const args = argsOrState as SpamfilterProfileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["flowBased"] = args ? args.flowBased : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gmail"] = args ? args.gmail : undefined;
            resourceInputs["imap"] = args ? args.imap : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["msnHotmail"] = args ? args.msnHotmail : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["pop3"] = args ? args.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["smtp"] = args ? args.smtp : undefined;
            resourceInputs["spamBwlTable"] = args ? args.spamBwlTable : undefined;
            resourceInputs["spamBwordTable"] = args ? args.spamBwordTable : undefined;
            resourceInputs["spamBwordThreshold"] = args ? args.spamBwordThreshold : undefined;
            resourceInputs["spamFiltering"] = args ? args.spamFiltering : undefined;
            resourceInputs["spamIptrustTable"] = args ? args.spamIptrustTable : undefined;
            resourceInputs["spamLog"] = args ? args.spamLog : undefined;
            resourceInputs["spamLogFortiguardResponse"] = args ? args.spamLogFortiguardResponse : undefined;
            resourceInputs["spamMheaderTable"] = args ? args.spamMheaderTable : undefined;
            resourceInputs["spamRblTable"] = args ? args.spamRblTable : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["yahooMail"] = args ? args.yahooMail : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SpamfilterProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpamfilterProfile resources.
 */
export interface SpamfilterProfileState {
    comment?: pulumi.Input<string>;
    external?: pulumi.Input<string>;
    flowBased?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    gmail?: pulumi.Input<inputs.SpamfilterProfileGmail>;
    imap?: pulumi.Input<inputs.SpamfilterProfileImap>;
    mapi?: pulumi.Input<inputs.SpamfilterProfileMapi>;
    msnHotmail?: pulumi.Input<inputs.SpamfilterProfileMsnHotmail>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    pop3?: pulumi.Input<inputs.SpamfilterProfilePop3>;
    replacemsgGroup?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.SpamfilterProfileSmtp>;
    spamBwlTable?: pulumi.Input<number>;
    spamBwordTable?: pulumi.Input<number>;
    spamBwordThreshold?: pulumi.Input<number>;
    spamFiltering?: pulumi.Input<string>;
    spamIptrustTable?: pulumi.Input<number>;
    spamLog?: pulumi.Input<string>;
    spamLogFortiguardResponse?: pulumi.Input<string>;
    spamMheaderTable?: pulumi.Input<number>;
    spamRblTable?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    yahooMail?: pulumi.Input<inputs.SpamfilterProfileYahooMail>;
}

/**
 * The set of arguments for constructing a SpamfilterProfile resource.
 */
export interface SpamfilterProfileArgs {
    comment?: pulumi.Input<string>;
    external?: pulumi.Input<string>;
    flowBased?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    gmail?: pulumi.Input<inputs.SpamfilterProfileGmail>;
    imap?: pulumi.Input<inputs.SpamfilterProfileImap>;
    mapi?: pulumi.Input<inputs.SpamfilterProfileMapi>;
    msnHotmail?: pulumi.Input<inputs.SpamfilterProfileMsnHotmail>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    pop3?: pulumi.Input<inputs.SpamfilterProfilePop3>;
    replacemsgGroup?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.SpamfilterProfileSmtp>;
    spamBwlTable?: pulumi.Input<number>;
    spamBwordTable?: pulumi.Input<number>;
    spamBwordThreshold?: pulumi.Input<number>;
    spamFiltering?: pulumi.Input<string>;
    spamIptrustTable?: pulumi.Input<number>;
    spamLog?: pulumi.Input<string>;
    spamLogFortiguardResponse?: pulumi.Input<string>;
    spamMheaderTable?: pulumi.Input<number>;
    spamRblTable?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    yahooMail?: pulumi.Input<inputs.SpamfilterProfileYahooMail>;
}

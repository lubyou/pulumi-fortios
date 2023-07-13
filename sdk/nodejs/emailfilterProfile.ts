// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class EmailfilterProfile extends pulumi.CustomResource {
    /**
     * Get an existing EmailfilterProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailfilterProfileState, opts?: pulumi.CustomResourceOptions): EmailfilterProfile {
        return new EmailfilterProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/emailfilterProfile:EmailfilterProfile';

    /**
     * Returns true if the given object is an instance of EmailfilterProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailfilterProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailfilterProfile.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly external!: pulumi.Output<string>;
    public readonly featureSet!: pulumi.Output<string>;
    public readonly fileFilter!: pulumi.Output<outputs.EmailfilterProfileFileFilter>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly gmail!: pulumi.Output<outputs.EmailfilterProfileGmail>;
    public readonly imap!: pulumi.Output<outputs.EmailfilterProfileImap>;
    public readonly mapi!: pulumi.Output<outputs.EmailfilterProfileMapi>;
    public readonly msnHotmail!: pulumi.Output<outputs.EmailfilterProfileMsnHotmail>;
    public readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<string>;
    public readonly otherWebmails!: pulumi.Output<outputs.EmailfilterProfileOtherWebmails>;
    public readonly pop3!: pulumi.Output<outputs.EmailfilterProfilePop3>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly smtp!: pulumi.Output<outputs.EmailfilterProfileSmtp>;
    public readonly spamBalTable!: pulumi.Output<number>;
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
    public readonly yahooMail!: pulumi.Output<outputs.EmailfilterProfileYahooMail>;

    /**
     * Create a EmailfilterProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EmailfilterProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailfilterProfileArgs | EmailfilterProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailfilterProfileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["fileFilter"] = state ? state.fileFilter : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gmail"] = state ? state.gmail : undefined;
            resourceInputs["imap"] = state ? state.imap : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["msnHotmail"] = state ? state.msnHotmail : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["otherWebmails"] = state ? state.otherWebmails : undefined;
            resourceInputs["pop3"] = state ? state.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["smtp"] = state ? state.smtp : undefined;
            resourceInputs["spamBalTable"] = state ? state.spamBalTable : undefined;
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
            const args = argsOrState as EmailfilterProfileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["fileFilter"] = args ? args.fileFilter : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gmail"] = args ? args.gmail : undefined;
            resourceInputs["imap"] = args ? args.imap : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["msnHotmail"] = args ? args.msnHotmail : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["otherWebmails"] = args ? args.otherWebmails : undefined;
            resourceInputs["pop3"] = args ? args.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["smtp"] = args ? args.smtp : undefined;
            resourceInputs["spamBalTable"] = args ? args.spamBalTable : undefined;
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
        super(EmailfilterProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailfilterProfile resources.
 */
export interface EmailfilterProfileState {
    comment?: pulumi.Input<string>;
    external?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    fileFilter?: pulumi.Input<inputs.EmailfilterProfileFileFilter>;
    getAllTables?: pulumi.Input<string>;
    gmail?: pulumi.Input<inputs.EmailfilterProfileGmail>;
    imap?: pulumi.Input<inputs.EmailfilterProfileImap>;
    mapi?: pulumi.Input<inputs.EmailfilterProfileMapi>;
    msnHotmail?: pulumi.Input<inputs.EmailfilterProfileMsnHotmail>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    otherWebmails?: pulumi.Input<inputs.EmailfilterProfileOtherWebmails>;
    pop3?: pulumi.Input<inputs.EmailfilterProfilePop3>;
    replacemsgGroup?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.EmailfilterProfileSmtp>;
    spamBalTable?: pulumi.Input<number>;
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
    yahooMail?: pulumi.Input<inputs.EmailfilterProfileYahooMail>;
}

/**
 * The set of arguments for constructing a EmailfilterProfile resource.
 */
export interface EmailfilterProfileArgs {
    comment?: pulumi.Input<string>;
    external?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    fileFilter?: pulumi.Input<inputs.EmailfilterProfileFileFilter>;
    getAllTables?: pulumi.Input<string>;
    gmail?: pulumi.Input<inputs.EmailfilterProfileGmail>;
    imap?: pulumi.Input<inputs.EmailfilterProfileImap>;
    mapi?: pulumi.Input<inputs.EmailfilterProfileMapi>;
    msnHotmail?: pulumi.Input<inputs.EmailfilterProfileMsnHotmail>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    otherWebmails?: pulumi.Input<inputs.EmailfilterProfileOtherWebmails>;
    pop3?: pulumi.Input<inputs.EmailfilterProfilePop3>;
    replacemsgGroup?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.EmailfilterProfileSmtp>;
    spamBalTable?: pulumi.Input<number>;
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
    yahooMail?: pulumi.Input<inputs.EmailfilterProfileYahooMail>;
}

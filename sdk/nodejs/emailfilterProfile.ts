// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Email Filter profiles. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Emailfilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/emailfilterProfile:EmailfilterProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable external Email inspection. Valid values: `enable`, `disable`.
     */
    public readonly external!: pulumi.Output<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    public readonly fileFilter!: pulumi.Output<outputs.EmailfilterProfileFileFilter | undefined>;
    /**
     * Gmail. The structure of `gmail` block is documented below.
     */
    public readonly gmail!: pulumi.Output<outputs.EmailfilterProfileGmail | undefined>;
    /**
     * IMAP. The structure of `imap` block is documented below.
     */
    public readonly imap!: pulumi.Output<outputs.EmailfilterProfileImap | undefined>;
    /**
     * MAPI. The structure of `mapi` block is documented below.
     */
    public readonly mapi!: pulumi.Output<outputs.EmailfilterProfileMapi | undefined>;
    /**
     * MSN Hotmail. The structure of `msnHotmail` block is documented below.
     */
    public readonly msnHotmail!: pulumi.Output<outputs.EmailfilterProfileMsnHotmail | undefined>;
    /**
     * File type name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Options.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Other supported webmails. The structure of `otherWebmails` block is documented below.
     */
    public readonly otherWebmails!: pulumi.Output<outputs.EmailfilterProfileOtherWebmails | undefined>;
    /**
     * POP3. The structure of `pop3` block is documented below.
     */
    public readonly pop3!: pulumi.Output<outputs.EmailfilterProfilePop3 | undefined>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * SMTP. The structure of `smtp` block is documented below.
     */
    public readonly smtp!: pulumi.Output<outputs.EmailfilterProfileSmtp | undefined>;
    /**
     * Anti-spam block/allow list table ID.
     */
    public readonly spamBalTable!: pulumi.Output<number>;
    /**
     * Anti-spam black/white list table ID.
     */
    public readonly spamBwlTable!: pulumi.Output<number>;
    /**
     * Anti-spam banned word table ID.
     */
    public readonly spamBwordTable!: pulumi.Output<number>;
    /**
     * Spam banned word threshold.
     */
    public readonly spamBwordThreshold!: pulumi.Output<number>;
    /**
     * Enable/disable spam filtering. Valid values: `enable`, `disable`.
     */
    public readonly spamFiltering!: pulumi.Output<string>;
    /**
     * Anti-spam IP trust table ID.
     */
    public readonly spamIptrustTable!: pulumi.Output<number>;
    /**
     * Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
     */
    public readonly spamLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
     */
    public readonly spamLogFortiguardResponse!: pulumi.Output<string>;
    /**
     * Anti-spam MIME header table ID.
     */
    public readonly spamMheaderTable!: pulumi.Output<number>;
    /**
     * Anti-spam DNSBL table ID.
     */
    public readonly spamRblTable!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Yahoo! Mail. The structure of `yahooMail` block is documented below.
     */
    public readonly yahooMail!: pulumi.Output<outputs.EmailfilterProfileYahooMail | undefined>;

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
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable external Email inspection. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.EmailfilterProfileFileFilter>;
    /**
     * Gmail. The structure of `gmail` block is documented below.
     */
    gmail?: pulumi.Input<inputs.EmailfilterProfileGmail>;
    /**
     * IMAP. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.EmailfilterProfileImap>;
    /**
     * MAPI. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.EmailfilterProfileMapi>;
    /**
     * MSN Hotmail. The structure of `msnHotmail` block is documented below.
     */
    msnHotmail?: pulumi.Input<inputs.EmailfilterProfileMsnHotmail>;
    /**
     * File type name.
     */
    name?: pulumi.Input<string>;
    /**
     * Options.
     */
    options?: pulumi.Input<string>;
    /**
     * Other supported webmails. The structure of `otherWebmails` block is documented below.
     */
    otherWebmails?: pulumi.Input<inputs.EmailfilterProfileOtherWebmails>;
    /**
     * POP3. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.EmailfilterProfilePop3>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * SMTP. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.EmailfilterProfileSmtp>;
    /**
     * Anti-spam block/allow list table ID.
     */
    spamBalTable?: pulumi.Input<number>;
    /**
     * Anti-spam black/white list table ID.
     */
    spamBwlTable?: pulumi.Input<number>;
    /**
     * Anti-spam banned word table ID.
     */
    spamBwordTable?: pulumi.Input<number>;
    /**
     * Spam banned word threshold.
     */
    spamBwordThreshold?: pulumi.Input<number>;
    /**
     * Enable/disable spam filtering. Valid values: `enable`, `disable`.
     */
    spamFiltering?: pulumi.Input<string>;
    /**
     * Anti-spam IP trust table ID.
     */
    spamIptrustTable?: pulumi.Input<number>;
    /**
     * Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
     */
    spamLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
     */
    spamLogFortiguardResponse?: pulumi.Input<string>;
    /**
     * Anti-spam MIME header table ID.
     */
    spamMheaderTable?: pulumi.Input<number>;
    /**
     * Anti-spam DNSBL table ID.
     */
    spamRblTable?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Yahoo! Mail. The structure of `yahooMail` block is documented below.
     */
    yahooMail?: pulumi.Input<inputs.EmailfilterProfileYahooMail>;
}

/**
 * The set of arguments for constructing a EmailfilterProfile resource.
 */
export interface EmailfilterProfileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable external Email inspection. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.EmailfilterProfileFileFilter>;
    /**
     * Gmail. The structure of `gmail` block is documented below.
     */
    gmail?: pulumi.Input<inputs.EmailfilterProfileGmail>;
    /**
     * IMAP. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.EmailfilterProfileImap>;
    /**
     * MAPI. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.EmailfilterProfileMapi>;
    /**
     * MSN Hotmail. The structure of `msnHotmail` block is documented below.
     */
    msnHotmail?: pulumi.Input<inputs.EmailfilterProfileMsnHotmail>;
    /**
     * File type name.
     */
    name?: pulumi.Input<string>;
    /**
     * Options.
     */
    options?: pulumi.Input<string>;
    /**
     * Other supported webmails. The structure of `otherWebmails` block is documented below.
     */
    otherWebmails?: pulumi.Input<inputs.EmailfilterProfileOtherWebmails>;
    /**
     * POP3. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.EmailfilterProfilePop3>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * SMTP. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.EmailfilterProfileSmtp>;
    /**
     * Anti-spam block/allow list table ID.
     */
    spamBalTable?: pulumi.Input<number>;
    /**
     * Anti-spam black/white list table ID.
     */
    spamBwlTable?: pulumi.Input<number>;
    /**
     * Anti-spam banned word table ID.
     */
    spamBwordTable?: pulumi.Input<number>;
    /**
     * Spam banned word threshold.
     */
    spamBwordThreshold?: pulumi.Input<number>;
    /**
     * Enable/disable spam filtering. Valid values: `enable`, `disable`.
     */
    spamFiltering?: pulumi.Input<string>;
    /**
     * Anti-spam IP trust table ID.
     */
    spamIptrustTable?: pulumi.Input<number>;
    /**
     * Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
     */
    spamLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
     */
    spamLogFortiguardResponse?: pulumi.Input<string>;
    /**
     * Anti-spam MIME header table ID.
     */
    spamMheaderTable?: pulumi.Input<number>;
    /**
     * Anti-spam DNSBL table ID.
     */
    spamRblTable?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Yahoo! Mail. The structure of `yahooMail` block is documented below.
     */
    yahooMail?: pulumi.Input<inputs.EmailfilterProfileYahooMail>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Email Filter profiles.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailfilterProfileState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["external"] = state ? state.external : undefined;
            inputs["featureSet"] = state ? state.featureSet : undefined;
            inputs["fileFilter"] = state ? state.fileFilter : undefined;
            inputs["gmail"] = state ? state.gmail : undefined;
            inputs["imap"] = state ? state.imap : undefined;
            inputs["mapi"] = state ? state.mapi : undefined;
            inputs["msnHotmail"] = state ? state.msnHotmail : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["otherWebmails"] = state ? state.otherWebmails : undefined;
            inputs["pop3"] = state ? state.pop3 : undefined;
            inputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            inputs["smtp"] = state ? state.smtp : undefined;
            inputs["spamBalTable"] = state ? state.spamBalTable : undefined;
            inputs["spamBwlTable"] = state ? state.spamBwlTable : undefined;
            inputs["spamBwordTable"] = state ? state.spamBwordTable : undefined;
            inputs["spamBwordThreshold"] = state ? state.spamBwordThreshold : undefined;
            inputs["spamFiltering"] = state ? state.spamFiltering : undefined;
            inputs["spamIptrustTable"] = state ? state.spamIptrustTable : undefined;
            inputs["spamLog"] = state ? state.spamLog : undefined;
            inputs["spamLogFortiguardResponse"] = state ? state.spamLogFortiguardResponse : undefined;
            inputs["spamMheaderTable"] = state ? state.spamMheaderTable : undefined;
            inputs["spamRblTable"] = state ? state.spamRblTable : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["yahooMail"] = state ? state.yahooMail : undefined;
        } else {
            const args = argsOrState as EmailfilterProfileArgs | undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["external"] = args ? args.external : undefined;
            inputs["featureSet"] = args ? args.featureSet : undefined;
            inputs["fileFilter"] = args ? args.fileFilter : undefined;
            inputs["gmail"] = args ? args.gmail : undefined;
            inputs["imap"] = args ? args.imap : undefined;
            inputs["mapi"] = args ? args.mapi : undefined;
            inputs["msnHotmail"] = args ? args.msnHotmail : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["otherWebmails"] = args ? args.otherWebmails : undefined;
            inputs["pop3"] = args ? args.pop3 : undefined;
            inputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            inputs["smtp"] = args ? args.smtp : undefined;
            inputs["spamBalTable"] = args ? args.spamBalTable : undefined;
            inputs["spamBwlTable"] = args ? args.spamBwlTable : undefined;
            inputs["spamBwordTable"] = args ? args.spamBwordTable : undefined;
            inputs["spamBwordThreshold"] = args ? args.spamBwordThreshold : undefined;
            inputs["spamFiltering"] = args ? args.spamFiltering : undefined;
            inputs["spamIptrustTable"] = args ? args.spamIptrustTable : undefined;
            inputs["spamLog"] = args ? args.spamLog : undefined;
            inputs["spamLogFortiguardResponse"] = args ? args.spamLogFortiguardResponse : undefined;
            inputs["spamMheaderTable"] = args ? args.spamMheaderTable : undefined;
            inputs["spamRblTable"] = args ? args.spamRblTable : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["yahooMail"] = args ? args.yahooMail : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EmailfilterProfile.__pulumiType, name, inputs, opts);
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
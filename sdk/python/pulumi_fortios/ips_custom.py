# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpsCustomArgs', 'IpsCustom']

@pulumi.input_type
class IpsCustomArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sig_name: Optional[pulumi.Input[str]] = None,
                 signature: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpsCustom resource.
        :param pulumi.Input[str] action: Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Applications to be protected. Blank for all applications.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] location: Protect client or server traffic.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] os: Operating system(s) that the signature protects. Blank for all operating systems.
        :param pulumi.Input[str] protocol: Protocol(s) that the signature scans. Blank for all protocols.
        :param pulumi.Input[int] rule_id: Signature ID.
        :param pulumi.Input[str] severity: Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        :param pulumi.Input[str] sig_name: Signature name.
        :param pulumi.Input[str] signature: Custom signature enclosed in single quotes.
        :param pulumi.Input[str] status: Enable/disable this signature. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tag: Signature tag.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if application is not None:
            pulumi.set(__self__, "application", application)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_packet is not None:
            pulumi.set(__self__, "log_packet", log_packet)
        if os is not None:
            pulumi.set(__self__, "os", os)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sig_name is not None:
            pulumi.set(__self__, "sig_name", sig_name)
        if signature is not None:
            pulumi.set(__self__, "signature", signature)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        Applications to be protected. Blank for all applications.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Protect client or server traffic.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @log_packet.setter
    def log_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_packet", value)

    @property
    @pulumi.getter
    def os(self) -> Optional[pulumi.Input[str]]:
        """
        Operating system(s) that the signature protects. Blank for all operating systems.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol(s) that the signature scans. Blank for all protocols.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        Signature ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="sigName")
    def sig_name(self) -> Optional[pulumi.Input[str]]:
        """
        Signature name.
        """
        return pulumi.get(self, "sig_name")

    @sig_name.setter
    def sig_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sig_name", value)

    @property
    @pulumi.getter
    def signature(self) -> Optional[pulumi.Input[str]]:
        """
        Custom signature enclosed in single quotes.
        """
        return pulumi.get(self, "signature")

    @signature.setter
    def signature(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this signature. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        Signature tag.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _IpsCustomState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sig_name: Optional[pulumi.Input[str]] = None,
                 signature: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpsCustom resources.
        :param pulumi.Input[str] action: Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Applications to be protected. Blank for all applications.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] location: Protect client or server traffic.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] os: Operating system(s) that the signature protects. Blank for all operating systems.
        :param pulumi.Input[str] protocol: Protocol(s) that the signature scans. Blank for all protocols.
        :param pulumi.Input[int] rule_id: Signature ID.
        :param pulumi.Input[str] severity: Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        :param pulumi.Input[str] sig_name: Signature name.
        :param pulumi.Input[str] signature: Custom signature enclosed in single quotes.
        :param pulumi.Input[str] status: Enable/disable this signature. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tag: Signature tag.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if application is not None:
            pulumi.set(__self__, "application", application)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_packet is not None:
            pulumi.set(__self__, "log_packet", log_packet)
        if os is not None:
            pulumi.set(__self__, "os", os)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if sig_name is not None:
            pulumi.set(__self__, "sig_name", sig_name)
        if signature is not None:
            pulumi.set(__self__, "signature", signature)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        Applications to be protected. Blank for all applications.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Protect client or server traffic.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @log_packet.setter
    def log_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_packet", value)

    @property
    @pulumi.getter
    def os(self) -> Optional[pulumi.Input[str]]:
        """
        Operating system(s) that the signature protects. Blank for all operating systems.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol(s) that the signature scans. Blank for all protocols.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        Signature ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="sigName")
    def sig_name(self) -> Optional[pulumi.Input[str]]:
        """
        Signature name.
        """
        return pulumi.get(self, "sig_name")

    @sig_name.setter
    def sig_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sig_name", value)

    @property
    @pulumi.getter
    def signature(self) -> Optional[pulumi.Input[str]]:
        """
        Custom signature enclosed in single quotes.
        """
        return pulumi.get(self, "signature")

    @signature.setter
    def signature(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this signature. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        Signature tag.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class IpsCustom(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sig_name: Optional[pulumi.Input[str]] = None,
                 signature: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPS custom signature.

        ## Import

        Ips Custom can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/ipsCustom:IpsCustom labelname {{tag}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/ipsCustom:IpsCustom labelname {{tag}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Applications to be protected. Blank for all applications.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] location: Protect client or server traffic.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] os: Operating system(s) that the signature protects. Blank for all operating systems.
        :param pulumi.Input[str] protocol: Protocol(s) that the signature scans. Blank for all protocols.
        :param pulumi.Input[int] rule_id: Signature ID.
        :param pulumi.Input[str] severity: Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        :param pulumi.Input[str] sig_name: Signature name.
        :param pulumi.Input[str] signature: Custom signature enclosed in single quotes.
        :param pulumi.Input[str] status: Enable/disable this signature. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tag: Signature tag.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IpsCustomArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPS custom signature.

        ## Import

        Ips Custom can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/ipsCustom:IpsCustom labelname {{tag}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/ipsCustom:IpsCustom labelname {{tag}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param IpsCustomArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpsCustomArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 sig_name: Optional[pulumi.Input[str]] = None,
                 signature: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpsCustomArgs.__new__(IpsCustomArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["application"] = application
            __props__.__dict__["comment"] = comment
            __props__.__dict__["location"] = location
            __props__.__dict__["log"] = log
            __props__.__dict__["log_packet"] = log_packet
            __props__.__dict__["os"] = os
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["rule_id"] = rule_id
            __props__.__dict__["severity"] = severity
            __props__.__dict__["sig_name"] = sig_name
            __props__.__dict__["signature"] = signature
            __props__.__dict__["status"] = status
            __props__.__dict__["tag"] = tag
            __props__.__dict__["vdomparam"] = vdomparam
        super(IpsCustom, __self__).__init__(
            'fortios:index/ipsCustom:IpsCustom',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            application: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            log: Optional[pulumi.Input[str]] = None,
            log_packet: Optional[pulumi.Input[str]] = None,
            os: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            rule_id: Optional[pulumi.Input[int]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            sig_name: Optional[pulumi.Input[str]] = None,
            signature: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'IpsCustom':
        """
        Get an existing IpsCustom resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Applications to be protected. Blank for all applications.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] location: Protect client or server traffic.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] os: Operating system(s) that the signature protects. Blank for all operating systems.
        :param pulumi.Input[str] protocol: Protocol(s) that the signature scans. Blank for all protocols.
        :param pulumi.Input[int] rule_id: Signature ID.
        :param pulumi.Input[str] severity: Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        :param pulumi.Input[str] sig_name: Signature name.
        :param pulumi.Input[str] signature: Custom signature enclosed in single quotes.
        :param pulumi.Input[str] status: Enable/disable this signature. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] tag: Signature tag.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpsCustomState.__new__(_IpsCustomState)

        __props__.__dict__["action"] = action
        __props__.__dict__["application"] = application
        __props__.__dict__["comment"] = comment
        __props__.__dict__["location"] = location
        __props__.__dict__["log"] = log
        __props__.__dict__["log_packet"] = log_packet
        __props__.__dict__["os"] = os
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["severity"] = severity
        __props__.__dict__["sig_name"] = sig_name
        __props__.__dict__["signature"] = signature
        __props__.__dict__["status"] = status
        __props__.__dict__["tag"] = tag
        __props__.__dict__["vdomparam"] = vdomparam
        return IpsCustom(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def application(self) -> pulumi.Output[str]:
        """
        Applications to be protected. Blank for all applications.
        """
        return pulumi.get(self, "application")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Protect client or server traffic.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def log(self) -> pulumi.Output[str]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> pulumi.Output[str]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @property
    @pulumi.getter
    def os(self) -> pulumi.Output[str]:
        """
        Operating system(s) that the signature protects. Blank for all operating systems.
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Protocol(s) that the signature scans. Blank for all protocols.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[int]:
        """
        Signature ID.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="sigName")
    def sig_name(self) -> pulumi.Output[str]:
        """
        Signature name.
        """
        return pulumi.get(self, "sig_name")

    @property
    @pulumi.getter
    def signature(self) -> pulumi.Output[str]:
        """
        Custom signature enclosed in single quotes.
        """
        return pulumi.get(self, "signature")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable this signature. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[str]:
        """
        Signature tag.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerFirewallObjectIppoolArgs', 'FortimanagerFirewallObjectIppool']

@pulumi.input_type
class FortimanagerFirewallObjectIppoolArgs:
    def __init__(__self__, *,
                 endip: pulumi.Input[str],
                 startip: pulumi.Input[str],
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FortimanagerFirewallObjectIppool resource.
        """
        pulumi.set(__self__, "endip", endip)
        pulumi.set(__self__, "startip", startip)
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_intf is not None:
            pulumi.set(__self__, "associated_intf", associated_intf)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: pulumi.Input[str]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: pulumi.Input[str]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_intf")

    @associated_intf.setter
    def associated_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_intf", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FortimanagerFirewallObjectIppoolState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FortimanagerFirewallObjectIppool resources.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_intf is not None:
            pulumi.set(__self__, "associated_intf", associated_intf)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if endip is not None:
            pulumi.set(__self__, "endip", endip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if startip is not None:
            pulumi.set(__self__, "startip", startip)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_intf")

    @associated_intf.setter
    def associated_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_intf", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def endip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def startip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FortimanagerFirewallObjectIppool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FortimanagerFirewallObjectIppool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FortimanagerFirewallObjectIppoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerFirewallObjectIppool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerFirewallObjectIppoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerFirewallObjectIppoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerFirewallObjectIppoolArgs.__new__(FortimanagerFirewallObjectIppoolArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["arp_intf"] = arp_intf
            __props__.__dict__["arp_reply"] = arp_reply
            __props__.__dict__["associated_intf"] = associated_intf
            __props__.__dict__["comment"] = comment
            if endip is None and not opts.urn:
                raise TypeError("Missing required property 'endip'")
            __props__.__dict__["endip"] = endip
            __props__.__dict__["name"] = name
            if startip is None and not opts.urn:
                raise TypeError("Missing required property 'startip'")
            __props__.__dict__["startip"] = startip
            __props__.__dict__["type"] = type
        super(FortimanagerFirewallObjectIppool, __self__).__init__(
            'fortios:index/fortimanagerFirewallObjectIppool:FortimanagerFirewallObjectIppool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            arp_intf: Optional[pulumi.Input[str]] = None,
            arp_reply: Optional[pulumi.Input[str]] = None,
            associated_intf: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            endip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            startip: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FortimanagerFirewallObjectIppool':
        """
        Get an existing FortimanagerFirewallObjectIppool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerFirewallObjectIppoolState.__new__(_FortimanagerFirewallObjectIppoolState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["arp_intf"] = arp_intf
        __props__.__dict__["arp_reply"] = arp_reply
        __props__.__dict__["associated_intf"] = associated_intf
        __props__.__dict__["comment"] = comment
        __props__.__dict__["endip"] = endip
        __props__.__dict__["name"] = name
        __props__.__dict__["startip"] = startip
        __props__.__dict__["type"] = type
        return FortimanagerFirewallObjectIppool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "arp_intf")

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "arp_reply")

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "associated_intf")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "startip")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallObjectVipArgs', 'FirewallObjectVip']

@pulumi.input_type
class FirewallObjectVipArgs:
    def __init__(__self__, *,
                 extip: pulumi.Input[str],
                 mappedips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 comment: Optional[pulumi.Input[str]] = None,
                 extintf: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallObjectVip resource.
        """
        pulumi.set(__self__, "extip", extip)
        pulumi.set(__self__, "mappedips", mappedips)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if extintf is not None:
            pulumi.set(__self__, "extintf", extintf)
        if extport is not None:
            pulumi.set(__self__, "extport", extport)
        if mappedport is not None:
            pulumi.set(__self__, "mappedport", mappedport)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portforward is not None:
            pulumi.set(__self__, "portforward", portforward)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def extip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "extip")

    @extip.setter
    def extip(self, value: pulumi.Input[str]):
        pulumi.set(self, "extip", value)

    @property
    @pulumi.getter
    def mappedips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "mappedips")

    @mappedips.setter
    def mappedips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "mappedips", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def extintf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extintf")

    @extintf.setter
    def extintf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extintf", value)

    @property
    @pulumi.getter
    def extport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extport")

    @extport.setter
    def extport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extport", value)

    @property
    @pulumi.getter
    def mappedport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mappedport")

    @mappedport.setter
    def mappedport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappedport", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def portforward(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portforward")

    @portforward.setter
    def portforward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portforward", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class _FirewallObjectVipState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 extintf: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 mappedips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallObjectVip resources.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if extintf is not None:
            pulumi.set(__self__, "extintf", extintf)
        if extip is not None:
            pulumi.set(__self__, "extip", extip)
        if extport is not None:
            pulumi.set(__self__, "extport", extport)
        if mappedips is not None:
            pulumi.set(__self__, "mappedips", mappedips)
        if mappedport is not None:
            pulumi.set(__self__, "mappedport", mappedport)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portforward is not None:
            pulumi.set(__self__, "portforward", portforward)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def extintf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extintf")

    @extintf.setter
    def extintf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extintf", value)

    @property
    @pulumi.getter
    def extip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extip")

    @extip.setter
    def extip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extip", value)

    @property
    @pulumi.getter
    def extport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extport")

    @extport.setter
    def extport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extport", value)

    @property
    @pulumi.getter
    def mappedips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "mappedips")

    @mappedips.setter
    def mappedips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "mappedips", value)

    @property
    @pulumi.getter
    def mappedport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mappedport")

    @mappedport.setter
    def mappedport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappedport", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def portforward(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portforward")

    @portforward.setter
    def portforward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portforward", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


class FirewallObjectVip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 extintf: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 mappedips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallObjectVip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallObjectVipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallObjectVip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallObjectVipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallObjectVipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 extintf: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 mappedips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallObjectVipArgs.__new__(FirewallObjectVipArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["extintf"] = extintf
            if extip is None and not opts.urn:
                raise TypeError("Missing required property 'extip'")
            __props__.__dict__["extip"] = extip
            __props__.__dict__["extport"] = extport
            if mappedips is None and not opts.urn:
                raise TypeError("Missing required property 'mappedips'")
            __props__.__dict__["mappedips"] = mappedips
            __props__.__dict__["mappedport"] = mappedport
            __props__.__dict__["name"] = name
            __props__.__dict__["portforward"] = portforward
            __props__.__dict__["protocol"] = protocol
        super(FirewallObjectVip, __self__).__init__(
            'fortios:index/firewallObjectVip:FirewallObjectVip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            extintf: Optional[pulumi.Input[str]] = None,
            extip: Optional[pulumi.Input[str]] = None,
            extport: Optional[pulumi.Input[str]] = None,
            mappedips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mappedport: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            portforward: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None) -> 'FirewallObjectVip':
        """
        Get an existing FirewallObjectVip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallObjectVipState.__new__(_FirewallObjectVipState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["extintf"] = extintf
        __props__.__dict__["extip"] = extip
        __props__.__dict__["extport"] = extport
        __props__.__dict__["mappedips"] = mappedips
        __props__.__dict__["mappedport"] = mappedport
        __props__.__dict__["name"] = name
        __props__.__dict__["portforward"] = portforward
        __props__.__dict__["protocol"] = protocol
        return FirewallObjectVip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def extintf(self) -> pulumi.Output[str]:
        return pulumi.get(self, "extintf")

    @property
    @pulumi.getter
    def extip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "extip")

    @property
    @pulumi.getter
    def extport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "extport")

    @property
    @pulumi.getter
    def mappedips(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "mappedips")

    @property
    @pulumi.getter
    def mappedport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mappedport")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def portforward(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portforward")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        return pulumi.get(self, "protocol")


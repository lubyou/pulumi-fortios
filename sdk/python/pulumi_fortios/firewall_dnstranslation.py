# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallDnstranslationArgs', 'FirewallDnstranslation']

@pulumi.input_type
class FirewallDnstranslationArgs:
    def __init__(__self__, *,
                 dst: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallDnstranslation resource.
        """
        if dst is not None:
            pulumi.set(__self__, "dst", dst)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if src is not None:
            pulumi.set(__self__, "src", src)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dst(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dst")

    @dst.setter
    def dst(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def src(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "src")

    @src.setter
    def src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallDnstranslationState:
    def __init__(__self__, *,
                 dst: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallDnstranslation resources.
        """
        if dst is not None:
            pulumi.set(__self__, "dst", dst)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if src is not None:
            pulumi.set(__self__, "src", src)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dst(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dst")

    @dst.setter
    def dst(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def src(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "src")

    @src.setter
    def src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallDnstranslation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallDnstranslation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallDnstranslationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallDnstranslation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallDnstranslationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallDnstranslationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallDnstranslationArgs.__new__(FirewallDnstranslationArgs)

            __props__.__dict__["dst"] = dst
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["netmask"] = netmask
            __props__.__dict__["src"] = src
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallDnstranslation, __self__).__init__(
            'fortios:index/firewallDnstranslation:FirewallDnstranslation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dst: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            src: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallDnstranslation':
        """
        Get an existing FirewallDnstranslation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallDnstranslationState.__new__(_FirewallDnstranslationState)

        __props__.__dict__["dst"] = dst
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["src"] = src
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallDnstranslation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dst(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[str]:
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def src(self) -> pulumi.Output[str]:
        return pulumi.get(self, "src")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


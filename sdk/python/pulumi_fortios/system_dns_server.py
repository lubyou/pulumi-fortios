# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemDnsServerArgs', 'SystemDnsServer']

@pulumi.input_type
class SystemDnsServerArgs:
    def __init__(__self__, *,
                 dnsfilter_profile: Optional[pulumi.Input[str]] = None,
                 doh: Optional[pulumi.Input[str]] = None,
                 doh3: Optional[pulumi.Input[str]] = None,
                 doq: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemDnsServer resource.
        """
        if dnsfilter_profile is not None:
            pulumi.set(__self__, "dnsfilter_profile", dnsfilter_profile)
        if doh is not None:
            pulumi.set(__self__, "doh", doh)
        if doh3 is not None:
            pulumi.set(__self__, "doh3", doh3)
        if doq is not None:
            pulumi.set(__self__, "doq", doq)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsfilterProfile")
    def dnsfilter_profile(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dnsfilter_profile")

    @dnsfilter_profile.setter
    def dnsfilter_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnsfilter_profile", value)

    @property
    @pulumi.getter
    def doh(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doh")

    @doh.setter
    def doh(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doh", value)

    @property
    @pulumi.getter
    def doh3(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doh3")

    @doh3.setter
    def doh3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doh3", value)

    @property
    @pulumi.getter
    def doq(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doq")

    @doq.setter
    def doq(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doq", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemDnsServerState:
    def __init__(__self__, *,
                 dnsfilter_profile: Optional[pulumi.Input[str]] = None,
                 doh: Optional[pulumi.Input[str]] = None,
                 doh3: Optional[pulumi.Input[str]] = None,
                 doq: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemDnsServer resources.
        """
        if dnsfilter_profile is not None:
            pulumi.set(__self__, "dnsfilter_profile", dnsfilter_profile)
        if doh is not None:
            pulumi.set(__self__, "doh", doh)
        if doh3 is not None:
            pulumi.set(__self__, "doh3", doh3)
        if doq is not None:
            pulumi.set(__self__, "doq", doq)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsfilterProfile")
    def dnsfilter_profile(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dnsfilter_profile")

    @dnsfilter_profile.setter
    def dnsfilter_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnsfilter_profile", value)

    @property
    @pulumi.getter
    def doh(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doh")

    @doh.setter
    def doh(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doh", value)

    @property
    @pulumi.getter
    def doh3(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doh3")

    @doh3.setter
    def doh3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doh3", value)

    @property
    @pulumi.getter
    def doq(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "doq")

    @doq.setter
    def doq(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "doq", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemDnsServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dnsfilter_profile: Optional[pulumi.Input[str]] = None,
                 doh: Optional[pulumi.Input[str]] = None,
                 doh3: Optional[pulumi.Input[str]] = None,
                 doq: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemDnsServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemDnsServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemDnsServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemDnsServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemDnsServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dnsfilter_profile: Optional[pulumi.Input[str]] = None,
                 doh: Optional[pulumi.Input[str]] = None,
                 doh3: Optional[pulumi.Input[str]] = None,
                 doq: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemDnsServerArgs.__new__(SystemDnsServerArgs)

            __props__.__dict__["dnsfilter_profile"] = dnsfilter_profile
            __props__.__dict__["doh"] = doh
            __props__.__dict__["doh3"] = doh3
            __props__.__dict__["doq"] = doq
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemDnsServer, __self__).__init__(
            'fortios:index/systemDnsServer:SystemDnsServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dnsfilter_profile: Optional[pulumi.Input[str]] = None,
            doh: Optional[pulumi.Input[str]] = None,
            doh3: Optional[pulumi.Input[str]] = None,
            doq: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemDnsServer':
        """
        Get an existing SystemDnsServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemDnsServerState.__new__(_SystemDnsServerState)

        __props__.__dict__["dnsfilter_profile"] = dnsfilter_profile
        __props__.__dict__["doh"] = doh
        __props__.__dict__["doh3"] = doh3
        __props__.__dict__["doq"] = doq
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemDnsServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsfilterProfile")
    def dnsfilter_profile(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dnsfilter_profile")

    @property
    @pulumi.getter
    def doh(self) -> pulumi.Output[str]:
        return pulumi.get(self, "doh")

    @property
    @pulumi.getter
    def doh3(self) -> pulumi.Output[str]:
        return pulumi.get(self, "doh3")

    @property
    @pulumi.getter
    def doq(self) -> pulumi.Output[str]:
        return pulumi.get(self, "doq")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


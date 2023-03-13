# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterIpsUrlfilterSetting6Args', 'WebfilterIpsUrlfilterSetting6']

@pulumi.input_type
class WebfilterIpsUrlfilterSetting6Args:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway6: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterIpsUrlfilterSetting6 resource.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if gateway6 is not None:
            pulumi.set(__self__, "gateway6", gateway6)
        if geo_filter is not None:
            pulumi.set(__self__, "geo_filter", geo_filter)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def gateway6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway6")

    @gateway6.setter
    def gateway6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway6", value)

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "geo_filter")

    @geo_filter.setter
    def geo_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geo_filter", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _WebfilterIpsUrlfilterSetting6State:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway6: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterIpsUrlfilterSetting6 resources.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if gateway6 is not None:
            pulumi.set(__self__, "gateway6", gateway6)
        if geo_filter is not None:
            pulumi.set(__self__, "geo_filter", geo_filter)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def gateway6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway6")

    @gateway6.setter
    def gateway6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway6", value)

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "geo_filter")

    @geo_filter.setter
    def geo_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geo_filter", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class WebfilterIpsUrlfilterSetting6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway6: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WebfilterIpsUrlfilterSetting6 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebfilterIpsUrlfilterSetting6Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WebfilterIpsUrlfilterSetting6 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WebfilterIpsUrlfilterSetting6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterIpsUrlfilterSetting6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway6: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebfilterIpsUrlfilterSetting6Args.__new__(WebfilterIpsUrlfilterSetting6Args)

            __props__.__dict__["device"] = device
            __props__.__dict__["distance"] = distance
            __props__.__dict__["gateway6"] = gateway6
            __props__.__dict__["geo_filter"] = geo_filter
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebfilterIpsUrlfilterSetting6, __self__).__init__(
            'fortios:index/webfilterIpsUrlfilterSetting6:WebfilterIpsUrlfilterSetting6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            distance: Optional[pulumi.Input[int]] = None,
            gateway6: Optional[pulumi.Input[str]] = None,
            geo_filter: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebfilterIpsUrlfilterSetting6':
        """
        Get an existing WebfilterIpsUrlfilterSetting6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterIpsUrlfilterSetting6State.__new__(_WebfilterIpsUrlfilterSetting6State)

        __props__.__dict__["device"] = device
        __props__.__dict__["distance"] = distance
        __props__.__dict__["gateway6"] = gateway6
        __props__.__dict__["geo_filter"] = geo_filter
        __props__.__dict__["vdomparam"] = vdomparam
        return WebfilterIpsUrlfilterSetting6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Output[int]:
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def gateway6(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway6")

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "geo_filter")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


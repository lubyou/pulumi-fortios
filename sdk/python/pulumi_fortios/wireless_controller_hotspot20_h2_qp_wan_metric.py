# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WirelessControllerHotspot20H2QpWanMetricArgs', 'WirelessControllerHotspot20H2QpWanMetric']

@pulumi.input_type
class WirelessControllerHotspot20H2QpWanMetricArgs:
    def __init__(__self__, *,
                 downlink_load: Optional[pulumi.Input[int]] = None,
                 downlink_speed: Optional[pulumi.Input[int]] = None,
                 link_at_capacity: Optional[pulumi.Input[str]] = None,
                 link_status: Optional[pulumi.Input[str]] = None,
                 load_measurement_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 symmetric_wan_link: Optional[pulumi.Input[str]] = None,
                 uplink_load: Optional[pulumi.Input[int]] = None,
                 uplink_speed: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerHotspot20H2QpWanMetric resource.
        :param pulumi.Input[int] downlink_load: Downlink load.
        :param pulumi.Input[int] downlink_speed: Downlink speed (in kilobits/s).
        :param pulumi.Input[str] link_at_capacity: Link at capacity. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] link_status: Link status. Valid values: `up`, `down`, `in-test`.
        :param pulumi.Input[int] load_measurement_duration: Load measurement duration (in tenths of a second).
        :param pulumi.Input[str] name: WAN metric name.
        :param pulumi.Input[str] symmetric_wan_link: WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        :param pulumi.Input[int] uplink_load: Uplink load.
        :param pulumi.Input[int] uplink_speed: Uplink speed (in kilobits/s).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if downlink_load is not None:
            pulumi.set(__self__, "downlink_load", downlink_load)
        if downlink_speed is not None:
            pulumi.set(__self__, "downlink_speed", downlink_speed)
        if link_at_capacity is not None:
            pulumi.set(__self__, "link_at_capacity", link_at_capacity)
        if link_status is not None:
            pulumi.set(__self__, "link_status", link_status)
        if load_measurement_duration is not None:
            pulumi.set(__self__, "load_measurement_duration", load_measurement_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if symmetric_wan_link is not None:
            pulumi.set(__self__, "symmetric_wan_link", symmetric_wan_link)
        if uplink_load is not None:
            pulumi.set(__self__, "uplink_load", uplink_load)
        if uplink_speed is not None:
            pulumi.set(__self__, "uplink_speed", uplink_speed)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downlinkLoad")
    def downlink_load(self) -> Optional[pulumi.Input[int]]:
        """
        Downlink load.
        """
        return pulumi.get(self, "downlink_load")

    @downlink_load.setter
    def downlink_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "downlink_load", value)

    @property
    @pulumi.getter(name="downlinkSpeed")
    def downlink_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Downlink speed (in kilobits/s).
        """
        return pulumi.get(self, "downlink_speed")

    @downlink_speed.setter
    def downlink_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "downlink_speed", value)

    @property
    @pulumi.getter(name="linkAtCapacity")
    def link_at_capacity(self) -> Optional[pulumi.Input[str]]:
        """
        Link at capacity. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "link_at_capacity")

    @link_at_capacity.setter
    def link_at_capacity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_at_capacity", value)

    @property
    @pulumi.getter(name="linkStatus")
    def link_status(self) -> Optional[pulumi.Input[str]]:
        """
        Link status. Valid values: `up`, `down`, `in-test`.
        """
        return pulumi.get(self, "link_status")

    @link_status.setter
    def link_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_status", value)

    @property
    @pulumi.getter(name="loadMeasurementDuration")
    def load_measurement_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Load measurement duration (in tenths of a second).
        """
        return pulumi.get(self, "load_measurement_duration")

    @load_measurement_duration.setter
    def load_measurement_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "load_measurement_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WAN metric name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="symmetricWanLink")
    def symmetric_wan_link(self) -> Optional[pulumi.Input[str]]:
        """
        WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        """
        return pulumi.get(self, "symmetric_wan_link")

    @symmetric_wan_link.setter
    def symmetric_wan_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "symmetric_wan_link", value)

    @property
    @pulumi.getter(name="uplinkLoad")
    def uplink_load(self) -> Optional[pulumi.Input[int]]:
        """
        Uplink load.
        """
        return pulumi.get(self, "uplink_load")

    @uplink_load.setter
    def uplink_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uplink_load", value)

    @property
    @pulumi.getter(name="uplinkSpeed")
    def uplink_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Uplink speed (in kilobits/s).
        """
        return pulumi.get(self, "uplink_speed")

    @uplink_speed.setter
    def uplink_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uplink_speed", value)

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
class _WirelessControllerHotspot20H2QpWanMetricState:
    def __init__(__self__, *,
                 downlink_load: Optional[pulumi.Input[int]] = None,
                 downlink_speed: Optional[pulumi.Input[int]] = None,
                 link_at_capacity: Optional[pulumi.Input[str]] = None,
                 link_status: Optional[pulumi.Input[str]] = None,
                 load_measurement_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 symmetric_wan_link: Optional[pulumi.Input[str]] = None,
                 uplink_load: Optional[pulumi.Input[int]] = None,
                 uplink_speed: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerHotspot20H2QpWanMetric resources.
        :param pulumi.Input[int] downlink_load: Downlink load.
        :param pulumi.Input[int] downlink_speed: Downlink speed (in kilobits/s).
        :param pulumi.Input[str] link_at_capacity: Link at capacity. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] link_status: Link status. Valid values: `up`, `down`, `in-test`.
        :param pulumi.Input[int] load_measurement_duration: Load measurement duration (in tenths of a second).
        :param pulumi.Input[str] name: WAN metric name.
        :param pulumi.Input[str] symmetric_wan_link: WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        :param pulumi.Input[int] uplink_load: Uplink load.
        :param pulumi.Input[int] uplink_speed: Uplink speed (in kilobits/s).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if downlink_load is not None:
            pulumi.set(__self__, "downlink_load", downlink_load)
        if downlink_speed is not None:
            pulumi.set(__self__, "downlink_speed", downlink_speed)
        if link_at_capacity is not None:
            pulumi.set(__self__, "link_at_capacity", link_at_capacity)
        if link_status is not None:
            pulumi.set(__self__, "link_status", link_status)
        if load_measurement_duration is not None:
            pulumi.set(__self__, "load_measurement_duration", load_measurement_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if symmetric_wan_link is not None:
            pulumi.set(__self__, "symmetric_wan_link", symmetric_wan_link)
        if uplink_load is not None:
            pulumi.set(__self__, "uplink_load", uplink_load)
        if uplink_speed is not None:
            pulumi.set(__self__, "uplink_speed", uplink_speed)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downlinkLoad")
    def downlink_load(self) -> Optional[pulumi.Input[int]]:
        """
        Downlink load.
        """
        return pulumi.get(self, "downlink_load")

    @downlink_load.setter
    def downlink_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "downlink_load", value)

    @property
    @pulumi.getter(name="downlinkSpeed")
    def downlink_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Downlink speed (in kilobits/s).
        """
        return pulumi.get(self, "downlink_speed")

    @downlink_speed.setter
    def downlink_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "downlink_speed", value)

    @property
    @pulumi.getter(name="linkAtCapacity")
    def link_at_capacity(self) -> Optional[pulumi.Input[str]]:
        """
        Link at capacity. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "link_at_capacity")

    @link_at_capacity.setter
    def link_at_capacity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_at_capacity", value)

    @property
    @pulumi.getter(name="linkStatus")
    def link_status(self) -> Optional[pulumi.Input[str]]:
        """
        Link status. Valid values: `up`, `down`, `in-test`.
        """
        return pulumi.get(self, "link_status")

    @link_status.setter
    def link_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_status", value)

    @property
    @pulumi.getter(name="loadMeasurementDuration")
    def load_measurement_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Load measurement duration (in tenths of a second).
        """
        return pulumi.get(self, "load_measurement_duration")

    @load_measurement_duration.setter
    def load_measurement_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "load_measurement_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WAN metric name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="symmetricWanLink")
    def symmetric_wan_link(self) -> Optional[pulumi.Input[str]]:
        """
        WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        """
        return pulumi.get(self, "symmetric_wan_link")

    @symmetric_wan_link.setter
    def symmetric_wan_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "symmetric_wan_link", value)

    @property
    @pulumi.getter(name="uplinkLoad")
    def uplink_load(self) -> Optional[pulumi.Input[int]]:
        """
        Uplink load.
        """
        return pulumi.get(self, "uplink_load")

    @uplink_load.setter
    def uplink_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uplink_load", value)

    @property
    @pulumi.getter(name="uplinkSpeed")
    def uplink_speed(self) -> Optional[pulumi.Input[int]]:
        """
        Uplink speed (in kilobits/s).
        """
        return pulumi.get(self, "uplink_speed")

    @uplink_speed.setter
    def uplink_speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uplink_speed", value)

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


class WirelessControllerHotspot20H2QpWanMetric(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 downlink_load: Optional[pulumi.Input[int]] = None,
                 downlink_speed: Optional[pulumi.Input[int]] = None,
                 link_at_capacity: Optional[pulumi.Input[str]] = None,
                 link_status: Optional[pulumi.Input[str]] = None,
                 load_measurement_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 symmetric_wan_link: Optional[pulumi.Input[str]] = None,
                 uplink_load: Optional[pulumi.Input[int]] = None,
                 uplink_speed: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WAN metrics.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerHotspot20H2QpWanMetric("trname",
            downlink_load=0,
            downlink_speed=2400,
            link_at_capacity="disable",
            link_status="up",
            load_measurement_duration=0,
            symmetric_wan_link="symmetric",
            uplink_load=0,
            uplink_speed=2400)
        ```

        ## Import

        WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerHotspot20H2QpWanMetric:WirelessControllerHotspot20H2QpWanMetric labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] downlink_load: Downlink load.
        :param pulumi.Input[int] downlink_speed: Downlink speed (in kilobits/s).
        :param pulumi.Input[str] link_at_capacity: Link at capacity. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] link_status: Link status. Valid values: `up`, `down`, `in-test`.
        :param pulumi.Input[int] load_measurement_duration: Load measurement duration (in tenths of a second).
        :param pulumi.Input[str] name: WAN metric name.
        :param pulumi.Input[str] symmetric_wan_link: WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        :param pulumi.Input[int] uplink_load: Uplink load.
        :param pulumi.Input[int] uplink_speed: Uplink speed (in kilobits/s).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerHotspot20H2QpWanMetricArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WAN metrics.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerHotspot20H2QpWanMetric("trname",
            downlink_load=0,
            downlink_speed=2400,
            link_at_capacity="disable",
            link_status="up",
            load_measurement_duration=0,
            symmetric_wan_link="symmetric",
            uplink_load=0,
            uplink_speed=2400)
        ```

        ## Import

        WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerHotspot20H2QpWanMetric:WirelessControllerHotspot20H2QpWanMetric labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerHotspot20H2QpWanMetricArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerHotspot20H2QpWanMetricArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 downlink_load: Optional[pulumi.Input[int]] = None,
                 downlink_speed: Optional[pulumi.Input[int]] = None,
                 link_at_capacity: Optional[pulumi.Input[str]] = None,
                 link_status: Optional[pulumi.Input[str]] = None,
                 load_measurement_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 symmetric_wan_link: Optional[pulumi.Input[str]] = None,
                 uplink_load: Optional[pulumi.Input[int]] = None,
                 uplink_speed: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerHotspot20H2QpWanMetricArgs.__new__(WirelessControllerHotspot20H2QpWanMetricArgs)

            __props__.__dict__["downlink_load"] = downlink_load
            __props__.__dict__["downlink_speed"] = downlink_speed
            __props__.__dict__["link_at_capacity"] = link_at_capacity
            __props__.__dict__["link_status"] = link_status
            __props__.__dict__["load_measurement_duration"] = load_measurement_duration
            __props__.__dict__["name"] = name
            __props__.__dict__["symmetric_wan_link"] = symmetric_wan_link
            __props__.__dict__["uplink_load"] = uplink_load
            __props__.__dict__["uplink_speed"] = uplink_speed
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerHotspot20H2QpWanMetric, __self__).__init__(
            'fortios:index/wirelessControllerHotspot20H2QpWanMetric:WirelessControllerHotspot20H2QpWanMetric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            downlink_load: Optional[pulumi.Input[int]] = None,
            downlink_speed: Optional[pulumi.Input[int]] = None,
            link_at_capacity: Optional[pulumi.Input[str]] = None,
            link_status: Optional[pulumi.Input[str]] = None,
            load_measurement_duration: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            symmetric_wan_link: Optional[pulumi.Input[str]] = None,
            uplink_load: Optional[pulumi.Input[int]] = None,
            uplink_speed: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerHotspot20H2QpWanMetric':
        """
        Get an existing WirelessControllerHotspot20H2QpWanMetric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] downlink_load: Downlink load.
        :param pulumi.Input[int] downlink_speed: Downlink speed (in kilobits/s).
        :param pulumi.Input[str] link_at_capacity: Link at capacity. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] link_status: Link status. Valid values: `up`, `down`, `in-test`.
        :param pulumi.Input[int] load_measurement_duration: Load measurement duration (in tenths of a second).
        :param pulumi.Input[str] name: WAN metric name.
        :param pulumi.Input[str] symmetric_wan_link: WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        :param pulumi.Input[int] uplink_load: Uplink load.
        :param pulumi.Input[int] uplink_speed: Uplink speed (in kilobits/s).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerHotspot20H2QpWanMetricState.__new__(_WirelessControllerHotspot20H2QpWanMetricState)

        __props__.__dict__["downlink_load"] = downlink_load
        __props__.__dict__["downlink_speed"] = downlink_speed
        __props__.__dict__["link_at_capacity"] = link_at_capacity
        __props__.__dict__["link_status"] = link_status
        __props__.__dict__["load_measurement_duration"] = load_measurement_duration
        __props__.__dict__["name"] = name
        __props__.__dict__["symmetric_wan_link"] = symmetric_wan_link
        __props__.__dict__["uplink_load"] = uplink_load
        __props__.__dict__["uplink_speed"] = uplink_speed
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerHotspot20H2QpWanMetric(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="downlinkLoad")
    def downlink_load(self) -> pulumi.Output[int]:
        """
        Downlink load.
        """
        return pulumi.get(self, "downlink_load")

    @property
    @pulumi.getter(name="downlinkSpeed")
    def downlink_speed(self) -> pulumi.Output[int]:
        """
        Downlink speed (in kilobits/s).
        """
        return pulumi.get(self, "downlink_speed")

    @property
    @pulumi.getter(name="linkAtCapacity")
    def link_at_capacity(self) -> pulumi.Output[str]:
        """
        Link at capacity. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "link_at_capacity")

    @property
    @pulumi.getter(name="linkStatus")
    def link_status(self) -> pulumi.Output[str]:
        """
        Link status. Valid values: `up`, `down`, `in-test`.
        """
        return pulumi.get(self, "link_status")

    @property
    @pulumi.getter(name="loadMeasurementDuration")
    def load_measurement_duration(self) -> pulumi.Output[int]:
        """
        Load measurement duration (in tenths of a second).
        """
        return pulumi.get(self, "load_measurement_duration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        WAN metric name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="symmetricWanLink")
    def symmetric_wan_link(self) -> pulumi.Output[str]:
        """
        WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
        """
        return pulumi.get(self, "symmetric_wan_link")

    @property
    @pulumi.getter(name="uplinkLoad")
    def uplink_load(self) -> pulumi.Output[int]:
        """
        Uplink load.
        """
        return pulumi.get(self, "uplink_load")

    @property
    @pulumi.getter(name="uplinkSpeed")
    def uplink_speed(self) -> pulumi.Output[int]:
        """
        Uplink speed (in kilobits/s).
        """
        return pulumi.get(self, "uplink_speed")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")


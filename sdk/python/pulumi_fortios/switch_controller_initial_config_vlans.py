# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerInitialConfigVlansArgs', 'SwitchControllerInitialConfigVlans']

@pulumi.input_type
class SwitchControllerInitialConfigVlansArgs:
    def __init__(__self__, *,
                 default_vlan: Optional[pulumi.Input[str]] = None,
                 nac: Optional[pulumi.Input[str]] = None,
                 nac_segment: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 rspan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video: Optional[pulumi.Input[str]] = None,
                 voice: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerInitialConfigVlans resource.
        :param pulumi.Input[str] default_vlan: Default VLAN (native) assigned to all switch ports upon discovery.
        :param pulumi.Input[str] nac: VLAN for NAC onboarding devices.
        :param pulumi.Input[str] nac_segment: VLAN for NAC segemnt primary interface.
        :param pulumi.Input[str] quarantine: VLAN for quarantined traffic.
        :param pulumi.Input[str] rspan: VLAN for RSPAN/ERSPAN mirrored traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] video: VLAN dedicated for video devices.
        :param pulumi.Input[str] voice: VLAN dedicated for voice devices.
        """
        if default_vlan is not None:
            pulumi.set(__self__, "default_vlan", default_vlan)
        if nac is not None:
            pulumi.set(__self__, "nac", nac)
        if nac_segment is not None:
            pulumi.set(__self__, "nac_segment", nac_segment)
        if quarantine is not None:
            pulumi.set(__self__, "quarantine", quarantine)
        if rspan is not None:
            pulumi.set(__self__, "rspan", rspan)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video is not None:
            pulumi.set(__self__, "video", video)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)

    @property
    @pulumi.getter(name="defaultVlan")
    def default_vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Default VLAN (native) assigned to all switch ports upon discovery.
        """
        return pulumi.get(self, "default_vlan")

    @default_vlan.setter
    def default_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_vlan", value)

    @property
    @pulumi.getter
    def nac(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for NAC onboarding devices.
        """
        return pulumi.get(self, "nac")

    @nac.setter
    def nac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac", value)

    @property
    @pulumi.getter(name="nacSegment")
    def nac_segment(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for NAC segemnt primary interface.
        """
        return pulumi.get(self, "nac_segment")

    @nac_segment.setter
    def nac_segment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac_segment", value)

    @property
    @pulumi.getter
    def quarantine(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for quarantined traffic.
        """
        return pulumi.get(self, "quarantine")

    @quarantine.setter
    def quarantine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quarantine", value)

    @property
    @pulumi.getter
    def rspan(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for RSPAN/ERSPAN mirrored traffic.
        """
        return pulumi.get(self, "rspan")

    @rspan.setter
    def rspan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rspan", value)

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

    @property
    @pulumi.getter
    def video(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN dedicated for video devices.
        """
        return pulumi.get(self, "video")

    @video.setter
    def video(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "video", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN dedicated for voice devices.
        """
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voice", value)


@pulumi.input_type
class _SwitchControllerInitialConfigVlansState:
    def __init__(__self__, *,
                 default_vlan: Optional[pulumi.Input[str]] = None,
                 nac: Optional[pulumi.Input[str]] = None,
                 nac_segment: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 rspan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video: Optional[pulumi.Input[str]] = None,
                 voice: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerInitialConfigVlans resources.
        :param pulumi.Input[str] default_vlan: Default VLAN (native) assigned to all switch ports upon discovery.
        :param pulumi.Input[str] nac: VLAN for NAC onboarding devices.
        :param pulumi.Input[str] nac_segment: VLAN for NAC segemnt primary interface.
        :param pulumi.Input[str] quarantine: VLAN for quarantined traffic.
        :param pulumi.Input[str] rspan: VLAN for RSPAN/ERSPAN mirrored traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] video: VLAN dedicated for video devices.
        :param pulumi.Input[str] voice: VLAN dedicated for voice devices.
        """
        if default_vlan is not None:
            pulumi.set(__self__, "default_vlan", default_vlan)
        if nac is not None:
            pulumi.set(__self__, "nac", nac)
        if nac_segment is not None:
            pulumi.set(__self__, "nac_segment", nac_segment)
        if quarantine is not None:
            pulumi.set(__self__, "quarantine", quarantine)
        if rspan is not None:
            pulumi.set(__self__, "rspan", rspan)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video is not None:
            pulumi.set(__self__, "video", video)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)

    @property
    @pulumi.getter(name="defaultVlan")
    def default_vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Default VLAN (native) assigned to all switch ports upon discovery.
        """
        return pulumi.get(self, "default_vlan")

    @default_vlan.setter
    def default_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_vlan", value)

    @property
    @pulumi.getter
    def nac(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for NAC onboarding devices.
        """
        return pulumi.get(self, "nac")

    @nac.setter
    def nac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac", value)

    @property
    @pulumi.getter(name="nacSegment")
    def nac_segment(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for NAC segemnt primary interface.
        """
        return pulumi.get(self, "nac_segment")

    @nac_segment.setter
    def nac_segment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac_segment", value)

    @property
    @pulumi.getter
    def quarantine(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for quarantined traffic.
        """
        return pulumi.get(self, "quarantine")

    @quarantine.setter
    def quarantine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quarantine", value)

    @property
    @pulumi.getter
    def rspan(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN for RSPAN/ERSPAN mirrored traffic.
        """
        return pulumi.get(self, "rspan")

    @rspan.setter
    def rspan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rspan", value)

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

    @property
    @pulumi.getter
    def video(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN dedicated for video devices.
        """
        return pulumi.get(self, "video")

    @video.setter
    def video(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "video", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN dedicated for voice devices.
        """
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voice", value)


class SwitchControllerInitialConfigVlans(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_vlan: Optional[pulumi.Input[str]] = None,
                 nac: Optional[pulumi.Input[str]] = None,
                 nac_segment: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 rspan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video: Optional[pulumi.Input[str]] = None,
                 voice: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure initial template for auto-generated VLAN interfaces. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        SwitchControllerInitialConfig Vlans can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerInitialConfigVlans:SwitchControllerInitialConfigVlans labelname SwitchControllerInitialConfigVlans
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_vlan: Default VLAN (native) assigned to all switch ports upon discovery.
        :param pulumi.Input[str] nac: VLAN for NAC onboarding devices.
        :param pulumi.Input[str] nac_segment: VLAN for NAC segemnt primary interface.
        :param pulumi.Input[str] quarantine: VLAN for quarantined traffic.
        :param pulumi.Input[str] rspan: VLAN for RSPAN/ERSPAN mirrored traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] video: VLAN dedicated for video devices.
        :param pulumi.Input[str] voice: VLAN dedicated for voice devices.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerInitialConfigVlansArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure initial template for auto-generated VLAN interfaces. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        SwitchControllerInitialConfig Vlans can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerInitialConfigVlans:SwitchControllerInitialConfigVlans labelname SwitchControllerInitialConfigVlans
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchControllerInitialConfigVlansArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerInitialConfigVlansArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_vlan: Optional[pulumi.Input[str]] = None,
                 nac: Optional[pulumi.Input[str]] = None,
                 nac_segment: Optional[pulumi.Input[str]] = None,
                 quarantine: Optional[pulumi.Input[str]] = None,
                 rspan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video: Optional[pulumi.Input[str]] = None,
                 voice: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchControllerInitialConfigVlansArgs.__new__(SwitchControllerInitialConfigVlansArgs)

            __props__.__dict__["default_vlan"] = default_vlan
            __props__.__dict__["nac"] = nac
            __props__.__dict__["nac_segment"] = nac_segment
            __props__.__dict__["quarantine"] = quarantine
            __props__.__dict__["rspan"] = rspan
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["video"] = video
            __props__.__dict__["voice"] = voice
        super(SwitchControllerInitialConfigVlans, __self__).__init__(
            'fortios:index/switchControllerInitialConfigVlans:SwitchControllerInitialConfigVlans',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_vlan: Optional[pulumi.Input[str]] = None,
            nac: Optional[pulumi.Input[str]] = None,
            nac_segment: Optional[pulumi.Input[str]] = None,
            quarantine: Optional[pulumi.Input[str]] = None,
            rspan: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            video: Optional[pulumi.Input[str]] = None,
            voice: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerInitialConfigVlans':
        """
        Get an existing SwitchControllerInitialConfigVlans resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_vlan: Default VLAN (native) assigned to all switch ports upon discovery.
        :param pulumi.Input[str] nac: VLAN for NAC onboarding devices.
        :param pulumi.Input[str] nac_segment: VLAN for NAC segemnt primary interface.
        :param pulumi.Input[str] quarantine: VLAN for quarantined traffic.
        :param pulumi.Input[str] rspan: VLAN for RSPAN/ERSPAN mirrored traffic.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] video: VLAN dedicated for video devices.
        :param pulumi.Input[str] voice: VLAN dedicated for voice devices.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerInitialConfigVlansState.__new__(_SwitchControllerInitialConfigVlansState)

        __props__.__dict__["default_vlan"] = default_vlan
        __props__.__dict__["nac"] = nac
        __props__.__dict__["nac_segment"] = nac_segment
        __props__.__dict__["quarantine"] = quarantine
        __props__.__dict__["rspan"] = rspan
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["video"] = video
        __props__.__dict__["voice"] = voice
        return SwitchControllerInitialConfigVlans(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultVlan")
    def default_vlan(self) -> pulumi.Output[str]:
        """
        Default VLAN (native) assigned to all switch ports upon discovery.
        """
        return pulumi.get(self, "default_vlan")

    @property
    @pulumi.getter
    def nac(self) -> pulumi.Output[str]:
        """
        VLAN for NAC onboarding devices.
        """
        return pulumi.get(self, "nac")

    @property
    @pulumi.getter(name="nacSegment")
    def nac_segment(self) -> pulumi.Output[str]:
        """
        VLAN for NAC segemnt primary interface.
        """
        return pulumi.get(self, "nac_segment")

    @property
    @pulumi.getter
    def quarantine(self) -> pulumi.Output[str]:
        """
        VLAN for quarantined traffic.
        """
        return pulumi.get(self, "quarantine")

    @property
    @pulumi.getter
    def rspan(self) -> pulumi.Output[str]:
        """
        VLAN for RSPAN/ERSPAN mirrored traffic.
        """
        return pulumi.get(self, "rspan")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def video(self) -> pulumi.Output[str]:
        """
        VLAN dedicated for video devices.
        """
        return pulumi.get(self, "video")

    @property
    @pulumi.getter
    def voice(self) -> pulumi.Output[str]:
        """
        VLAN dedicated for voice devices.
        """
        return pulumi.get(self, "voice")


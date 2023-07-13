# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SystemLldpNetworkPolicyArgs', 'SystemLldpNetworkPolicy']

@pulumi.input_type
class SystemLldpNetworkPolicyArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']] = None,
                 guest_voice_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']] = None,
                 streaming_video: Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']] = None,
                 video_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']] = None,
                 voice: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']] = None,
                 voice_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']] = None):
        """
        The set of arguments for constructing a SystemLldpNetworkPolicy resource.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if guest is not None:
            pulumi.set(__self__, "guest", guest)
        if guest_voice_signaling is not None:
            pulumi.set(__self__, "guest_voice_signaling", guest_voice_signaling)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if softphone is not None:
            pulumi.set(__self__, "softphone", softphone)
        if streaming_video is not None:
            pulumi.set(__self__, "streaming_video", streaming_video)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video_conferencing is not None:
            pulumi.set(__self__, "video_conferencing", video_conferencing)
        if video_signaling is not None:
            pulumi.set(__self__, "video_signaling", video_signaling)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)
        if voice_signaling is not None:
            pulumi.set(__self__, "voice_signaling", voice_signaling)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def guest(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']]:
        return pulumi.get(self, "guest")

    @guest.setter
    def guest(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']]):
        pulumi.set(self, "guest", value)

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]:
        return pulumi.get(self, "guest_voice_signaling")

    @guest_voice_signaling.setter
    def guest_voice_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]):
        pulumi.set(self, "guest_voice_signaling", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def softphone(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']]:
        return pulumi.get(self, "softphone")

    @softphone.setter
    def softphone(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']]):
        pulumi.set(self, "softphone", value)

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']]:
        return pulumi.get(self, "streaming_video")

    @streaming_video.setter
    def streaming_video(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']]):
        pulumi.set(self, "streaming_video", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']]:
        return pulumi.get(self, "video_conferencing")

    @video_conferencing.setter
    def video_conferencing(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']]):
        pulumi.set(self, "video_conferencing", value)

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']]:
        return pulumi.get(self, "video_signaling")

    @video_signaling.setter
    def video_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']]):
        pulumi.set(self, "video_signaling", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']]:
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']]):
        pulumi.set(self, "voice", value)

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']]:
        return pulumi.get(self, "voice_signaling")

    @voice_signaling.setter
    def voice_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']]):
        pulumi.set(self, "voice_signaling", value)


@pulumi.input_type
class _SystemLldpNetworkPolicyState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']] = None,
                 guest_voice_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']] = None,
                 streaming_video: Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']] = None,
                 video_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']] = None,
                 voice: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']] = None,
                 voice_signaling: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']] = None):
        """
        Input properties used for looking up and filtering SystemLldpNetworkPolicy resources.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if guest is not None:
            pulumi.set(__self__, "guest", guest)
        if guest_voice_signaling is not None:
            pulumi.set(__self__, "guest_voice_signaling", guest_voice_signaling)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if softphone is not None:
            pulumi.set(__self__, "softphone", softphone)
        if streaming_video is not None:
            pulumi.set(__self__, "streaming_video", streaming_video)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video_conferencing is not None:
            pulumi.set(__self__, "video_conferencing", video_conferencing)
        if video_signaling is not None:
            pulumi.set(__self__, "video_signaling", video_signaling)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)
        if voice_signaling is not None:
            pulumi.set(__self__, "voice_signaling", voice_signaling)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def guest(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']]:
        return pulumi.get(self, "guest")

    @guest.setter
    def guest(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestArgs']]):
        pulumi.set(self, "guest", value)

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]:
        return pulumi.get(self, "guest_voice_signaling")

    @guest_voice_signaling.setter
    def guest_voice_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]):
        pulumi.set(self, "guest_voice_signaling", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def softphone(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']]:
        return pulumi.get(self, "softphone")

    @softphone.setter
    def softphone(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicySoftphoneArgs']]):
        pulumi.set(self, "softphone", value)

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']]:
        return pulumi.get(self, "streaming_video")

    @streaming_video.setter
    def streaming_video(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyStreamingVideoArgs']]):
        pulumi.set(self, "streaming_video", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']]:
        return pulumi.get(self, "video_conferencing")

    @video_conferencing.setter
    def video_conferencing(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoConferencingArgs']]):
        pulumi.set(self, "video_conferencing", value)

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']]:
        return pulumi.get(self, "video_signaling")

    @video_signaling.setter
    def video_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVideoSignalingArgs']]):
        pulumi.set(self, "video_signaling", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']]:
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceArgs']]):
        pulumi.set(self, "voice", value)

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']]:
        return pulumi.get(self, "voice_signaling")

    @voice_signaling.setter
    def voice_signaling(self, value: Optional[pulumi.Input['SystemLldpNetworkPolicyVoiceSignalingArgs']]):
        pulumi.set(self, "voice_signaling", value)


class SystemLldpNetworkPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestArgs']]] = None,
                 guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicySoftphoneArgs']]] = None,
                 streaming_video: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyStreamingVideoArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoConferencingArgs']]] = None,
                 video_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoSignalingArgs']]] = None,
                 voice: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceArgs']]] = None,
                 voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceSignalingArgs']]] = None,
                 __props__=None):
        """
        Create a SystemLldpNetworkPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemLldpNetworkPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemLldpNetworkPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemLldpNetworkPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemLldpNetworkPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestArgs']]] = None,
                 guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicySoftphoneArgs']]] = None,
                 streaming_video: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyStreamingVideoArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoConferencingArgs']]] = None,
                 video_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoSignalingArgs']]] = None,
                 voice: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceArgs']]] = None,
                 voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceSignalingArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemLldpNetworkPolicyArgs.__new__(SystemLldpNetworkPolicyArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["guest"] = guest
            __props__.__dict__["guest_voice_signaling"] = guest_voice_signaling
            __props__.__dict__["name"] = name
            __props__.__dict__["softphone"] = softphone
            __props__.__dict__["streaming_video"] = streaming_video
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["video_conferencing"] = video_conferencing
            __props__.__dict__["video_signaling"] = video_signaling
            __props__.__dict__["voice"] = voice
            __props__.__dict__["voice_signaling"] = voice_signaling
        super(SystemLldpNetworkPolicy, __self__).__init__(
            'fortios:index/systemLldpNetworkPolicy:SystemLldpNetworkPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            guest: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestArgs']]] = None,
            guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyGuestVoiceSignalingArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            softphone: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicySoftphoneArgs']]] = None,
            streaming_video: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyStreamingVideoArgs']]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            video_conferencing: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoConferencingArgs']]] = None,
            video_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVideoSignalingArgs']]] = None,
            voice: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceArgs']]] = None,
            voice_signaling: Optional[pulumi.Input[pulumi.InputType['SystemLldpNetworkPolicyVoiceSignalingArgs']]] = None) -> 'SystemLldpNetworkPolicy':
        """
        Get an existing SystemLldpNetworkPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemLldpNetworkPolicyState.__new__(_SystemLldpNetworkPolicyState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["guest"] = guest
        __props__.__dict__["guest_voice_signaling"] = guest_voice_signaling
        __props__.__dict__["name"] = name
        __props__.__dict__["softphone"] = softphone
        __props__.__dict__["streaming_video"] = streaming_video
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["video_conferencing"] = video_conferencing
        __props__.__dict__["video_signaling"] = video_signaling
        __props__.__dict__["voice"] = voice
        __props__.__dict__["voice_signaling"] = voice_signaling
        return SystemLldpNetworkPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def guest(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyGuest']:
        return pulumi.get(self, "guest")

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyGuestVoiceSignaling']:
        return pulumi.get(self, "guest_voice_signaling")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def softphone(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicySoftphone']:
        return pulumi.get(self, "softphone")

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyStreamingVideo']:
        return pulumi.get(self, "streaming_video")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyVideoConferencing']:
        return pulumi.get(self, "video_conferencing")

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyVideoSignaling']:
        return pulumi.get(self, "video_signaling")

    @property
    @pulumi.getter
    def voice(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyVoice']:
        return pulumi.get(self, "voice")

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> pulumi.Output['outputs.SystemLldpNetworkPolicyVoiceSignaling']:
        return pulumi.get(self, "voice_signaling")


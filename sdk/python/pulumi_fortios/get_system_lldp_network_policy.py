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

__all__ = [
    'GetSystemLldpNetworkPolicyResult',
    'AwaitableGetSystemLldpNetworkPolicyResult',
    'get_system_lldp_network_policy',
    'get_system_lldp_network_policy_output',
]

@pulumi.output_type
class GetSystemLldpNetworkPolicyResult:
    """
    A collection of values returned by GetSystemLldpNetworkPolicy.
    """
    def __init__(__self__, comment=None, guest_voice_signalings=None, guests=None, id=None, name=None, softphones=None, streaming_videos=None, vdomparam=None, video_conferencings=None, video_signalings=None, voice_signalings=None, voices=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if guest_voice_signalings and not isinstance(guest_voice_signalings, list):
            raise TypeError("Expected argument 'guest_voice_signalings' to be a list")
        pulumi.set(__self__, "guest_voice_signalings", guest_voice_signalings)
        if guests and not isinstance(guests, list):
            raise TypeError("Expected argument 'guests' to be a list")
        pulumi.set(__self__, "guests", guests)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if softphones and not isinstance(softphones, list):
            raise TypeError("Expected argument 'softphones' to be a list")
        pulumi.set(__self__, "softphones", softphones)
        if streaming_videos and not isinstance(streaming_videos, list):
            raise TypeError("Expected argument 'streaming_videos' to be a list")
        pulumi.set(__self__, "streaming_videos", streaming_videos)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if video_conferencings and not isinstance(video_conferencings, list):
            raise TypeError("Expected argument 'video_conferencings' to be a list")
        pulumi.set(__self__, "video_conferencings", video_conferencings)
        if video_signalings and not isinstance(video_signalings, list):
            raise TypeError("Expected argument 'video_signalings' to be a list")
        pulumi.set(__self__, "video_signalings", video_signalings)
        if voice_signalings and not isinstance(voice_signalings, list):
            raise TypeError("Expected argument 'voice_signalings' to be a list")
        pulumi.set(__self__, "voice_signalings", voice_signalings)
        if voices and not isinstance(voices, list):
            raise TypeError("Expected argument 'voices' to be a list")
        pulumi.set(__self__, "voices", voices)

    @property
    @pulumi.getter
    def comment(self) -> str:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="guestVoiceSignalings")
    def guest_voice_signalings(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyGuestVoiceSignalingResult']:
        return pulumi.get(self, "guest_voice_signalings")

    @property
    @pulumi.getter
    def guests(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyGuestResult']:
        return pulumi.get(self, "guests")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def softphones(self) -> Sequence['outputs.GetSystemLldpNetworkPolicySoftphoneResult']:
        return pulumi.get(self, "softphones")

    @property
    @pulumi.getter(name="streamingVideos")
    def streaming_videos(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyStreamingVideoResult']:
        return pulumi.get(self, "streaming_videos")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="videoConferencings")
    def video_conferencings(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyVideoConferencingResult']:
        return pulumi.get(self, "video_conferencings")

    @property
    @pulumi.getter(name="videoSignalings")
    def video_signalings(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyVideoSignalingResult']:
        return pulumi.get(self, "video_signalings")

    @property
    @pulumi.getter(name="voiceSignalings")
    def voice_signalings(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyVoiceSignalingResult']:
        return pulumi.get(self, "voice_signalings")

    @property
    @pulumi.getter
    def voices(self) -> Sequence['outputs.GetSystemLldpNetworkPolicyVoiceResult']:
        return pulumi.get(self, "voices")


class AwaitableGetSystemLldpNetworkPolicyResult(GetSystemLldpNetworkPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemLldpNetworkPolicyResult(
            comment=self.comment,
            guest_voice_signalings=self.guest_voice_signalings,
            guests=self.guests,
            id=self.id,
            name=self.name,
            softphones=self.softphones,
            streaming_videos=self.streaming_videos,
            vdomparam=self.vdomparam,
            video_conferencings=self.video_conferencings,
            video_signalings=self.video_signalings,
            voice_signalings=self.voice_signalings,
            voices=self.voices)


def get_system_lldp_network_policy(name: Optional[str] = None,
                                   vdomparam: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemLldpNetworkPolicyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemLldpNetworkPolicy:GetSystemLldpNetworkPolicy', __args__, opts=opts, typ=GetSystemLldpNetworkPolicyResult).value

    return AwaitableGetSystemLldpNetworkPolicyResult(
        comment=pulumi.get(__ret__, 'comment'),
        guest_voice_signalings=pulumi.get(__ret__, 'guest_voice_signalings'),
        guests=pulumi.get(__ret__, 'guests'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        softphones=pulumi.get(__ret__, 'softphones'),
        streaming_videos=pulumi.get(__ret__, 'streaming_videos'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        video_conferencings=pulumi.get(__ret__, 'video_conferencings'),
        video_signalings=pulumi.get(__ret__, 'video_signalings'),
        voice_signalings=pulumi.get(__ret__, 'voice_signalings'),
        voices=pulumi.get(__ret__, 'voices'))


@_utilities.lift_output_func(get_system_lldp_network_policy)
def get_system_lldp_network_policy_output(name: Optional[pulumi.Input[str]] = None,
                                          vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemLldpNetworkPolicyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

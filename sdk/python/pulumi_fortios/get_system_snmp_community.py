# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemSnmpCommunityResult',
    'AwaitableGetSystemSnmpCommunityResult',
    'get_system_snmp_community',
]

@pulumi.output_type
class GetSystemSnmpCommunityResult:
    """
    A collection of values returned by GetSystemSnmpCommunity.
    """
    def __init__(__self__, events=None, fosid=None, hosts=None, hosts6s=None, id=None, name=None, query_v1_port=None, query_v1_status=None, query_v2c_port=None, query_v2c_status=None, status=None, trap_v1_lport=None, trap_v1_rport=None, trap_v1_status=None, trap_v2c_lport=None, trap_v2c_rport=None, trap_v2c_status=None, vdomparam=None):
        if events and not isinstance(events, str):
            raise TypeError("Expected argument 'events' to be a str")
        pulumi.set(__self__, "events", events)
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if hosts and not isinstance(hosts, list):
            raise TypeError("Expected argument 'hosts' to be a list")
        pulumi.set(__self__, "hosts", hosts)
        if hosts6s and not isinstance(hosts6s, list):
            raise TypeError("Expected argument 'hosts6s' to be a list")
        pulumi.set(__self__, "hosts6s", hosts6s)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query_v1_port and not isinstance(query_v1_port, int):
            raise TypeError("Expected argument 'query_v1_port' to be a int")
        pulumi.set(__self__, "query_v1_port", query_v1_port)
        if query_v1_status and not isinstance(query_v1_status, str):
            raise TypeError("Expected argument 'query_v1_status' to be a str")
        pulumi.set(__self__, "query_v1_status", query_v1_status)
        if query_v2c_port and not isinstance(query_v2c_port, int):
            raise TypeError("Expected argument 'query_v2c_port' to be a int")
        pulumi.set(__self__, "query_v2c_port", query_v2c_port)
        if query_v2c_status and not isinstance(query_v2c_status, str):
            raise TypeError("Expected argument 'query_v2c_status' to be a str")
        pulumi.set(__self__, "query_v2c_status", query_v2c_status)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if trap_v1_lport and not isinstance(trap_v1_lport, int):
            raise TypeError("Expected argument 'trap_v1_lport' to be a int")
        pulumi.set(__self__, "trap_v1_lport", trap_v1_lport)
        if trap_v1_rport and not isinstance(trap_v1_rport, int):
            raise TypeError("Expected argument 'trap_v1_rport' to be a int")
        pulumi.set(__self__, "trap_v1_rport", trap_v1_rport)
        if trap_v1_status and not isinstance(trap_v1_status, str):
            raise TypeError("Expected argument 'trap_v1_status' to be a str")
        pulumi.set(__self__, "trap_v1_status", trap_v1_status)
        if trap_v2c_lport and not isinstance(trap_v2c_lport, int):
            raise TypeError("Expected argument 'trap_v2c_lport' to be a int")
        pulumi.set(__self__, "trap_v2c_lport", trap_v2c_lport)
        if trap_v2c_rport and not isinstance(trap_v2c_rport, int):
            raise TypeError("Expected argument 'trap_v2c_rport' to be a int")
        pulumi.set(__self__, "trap_v2c_rport", trap_v2c_rport)
        if trap_v2c_status and not isinstance(trap_v2c_status, str):
            raise TypeError("Expected argument 'trap_v2c_status' to be a str")
        pulumi.set(__self__, "trap_v2c_status", trap_v2c_status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def events(self) -> str:
        """
        SNMP trap events.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Community ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def hosts(self) -> Sequence['outputs.GetSystemSnmpCommunityHostResult']:
        """
        Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def hosts6s(self) -> Sequence['outputs.GetSystemSnmpCommunityHosts6Result']:
        """
        Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
        """
        return pulumi.get(self, "hosts6s")

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
        """
        Community name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryV1Port")
    def query_v1_port(self) -> int:
        """
        SNMP v1 query port (default = 161).
        """
        return pulumi.get(self, "query_v1_port")

    @property
    @pulumi.getter(name="queryV1Status")
    def query_v1_status(self) -> str:
        """
        Enable/disable SNMP v1 queries.
        """
        return pulumi.get(self, "query_v1_status")

    @property
    @pulumi.getter(name="queryV2cPort")
    def query_v2c_port(self) -> int:
        """
        SNMP v2c query port (default = 161).
        """
        return pulumi.get(self, "query_v2c_port")

    @property
    @pulumi.getter(name="queryV2cStatus")
    def query_v2c_status(self) -> str:
        """
        Enable/disable SNMP v2c queries.
        """
        return pulumi.get(self, "query_v2c_status")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this SNMP community.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trapV1Lport")
    def trap_v1_lport(self) -> int:
        """
        SNMP v1 trap local port (default = 162).
        """
        return pulumi.get(self, "trap_v1_lport")

    @property
    @pulumi.getter(name="trapV1Rport")
    def trap_v1_rport(self) -> int:
        """
        SNMP v1 trap remote port (default = 162).
        """
        return pulumi.get(self, "trap_v1_rport")

    @property
    @pulumi.getter(name="trapV1Status")
    def trap_v1_status(self) -> str:
        """
        Enable/disable SNMP v1 traps.
        """
        return pulumi.get(self, "trap_v1_status")

    @property
    @pulumi.getter(name="trapV2cLport")
    def trap_v2c_lport(self) -> int:
        """
        SNMP v2c trap local port (default = 162).
        """
        return pulumi.get(self, "trap_v2c_lport")

    @property
    @pulumi.getter(name="trapV2cRport")
    def trap_v2c_rport(self) -> int:
        """
        SNMP v2c trap remote port (default = 162).
        """
        return pulumi.get(self, "trap_v2c_rport")

    @property
    @pulumi.getter(name="trapV2cStatus")
    def trap_v2c_status(self) -> str:
        """
        Enable/disable SNMP v2c traps.
        """
        return pulumi.get(self, "trap_v2c_status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemSnmpCommunityResult(GetSystemSnmpCommunityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemSnmpCommunityResult(
            events=self.events,
            fosid=self.fosid,
            hosts=self.hosts,
            hosts6s=self.hosts6s,
            id=self.id,
            name=self.name,
            query_v1_port=self.query_v1_port,
            query_v1_status=self.query_v1_status,
            query_v2c_port=self.query_v2c_port,
            query_v2c_status=self.query_v2c_status,
            status=self.status,
            trap_v1_lport=self.trap_v1_lport,
            trap_v1_rport=self.trap_v1_rport,
            trap_v1_status=self.trap_v1_status,
            trap_v2c_lport=self.trap_v2c_lport,
            trap_v2c_rport=self.trap_v2c_rport,
            trap_v2c_status=self.trap_v2c_status,
            vdomparam=self.vdomparam)


def get_system_snmp_community(fosid: Optional[int] = None,
                              vdomparam: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemSnmpCommunityResult:
    """
    Use this data source to get information on an fortios systemsnmp community


    :param int fosid: Specify the fosid of the desired systemsnmp community.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemSnmpCommunity:GetSystemSnmpCommunity', __args__, opts=opts, typ=GetSystemSnmpCommunityResult).value

    return AwaitableGetSystemSnmpCommunityResult(
        events=__ret__.events,
        fosid=__ret__.fosid,
        hosts=__ret__.hosts,
        hosts6s=__ret__.hosts6s,
        id=__ret__.id,
        name=__ret__.name,
        query_v1_port=__ret__.query_v1_port,
        query_v1_status=__ret__.query_v1_status,
        query_v2c_port=__ret__.query_v2c_port,
        query_v2c_status=__ret__.query_v2c_status,
        status=__ret__.status,
        trap_v1_lport=__ret__.trap_v1_lport,
        trap_v1_rport=__ret__.trap_v1_rport,
        trap_v1_status=__ret__.trap_v1_status,
        trap_v2c_lport=__ret__.trap_v2c_lport,
        trap_v2c_rport=__ret__.trap_v2c_rport,
        trap_v2c_status=__ret__.trap_v2c_status,
        vdomparam=__ret__.vdomparam)
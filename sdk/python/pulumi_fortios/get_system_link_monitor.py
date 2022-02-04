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
    'GetSystemLinkMonitorResult',
    'AwaitableGetSystemLinkMonitorResult',
    'get_system_link_monitor',
    'get_system_link_monitor_output',
]

@pulumi.output_type
class GetSystemLinkMonitorResult:
    """
    A collection of values returned by GetSystemLinkMonitor.
    """
    def __init__(__self__, addr_mode=None, class_id=None, diffservcode=None, fail_weight=None, failtime=None, gateway_ip=None, gateway_ip6=None, ha_priority=None, http_agent=None, http_get=None, http_match=None, id=None, interval=None, name=None, packet_size=None, password=None, port=None, probe_count=None, probe_timeout=None, protocol=None, recoverytime=None, routes=None, security_mode=None, server_config=None, server_lists=None, servers=None, service_detection=None, source_ip=None, source_ip6=None, srcintf=None, status=None, update_cascade_interface=None, update_policy_route=None, update_static_route=None, vdomparam=None):
        if addr_mode and not isinstance(addr_mode, str):
            raise TypeError("Expected argument 'addr_mode' to be a str")
        pulumi.set(__self__, "addr_mode", addr_mode)
        if class_id and not isinstance(class_id, int):
            raise TypeError("Expected argument 'class_id' to be a int")
        pulumi.set(__self__, "class_id", class_id)
        if diffservcode and not isinstance(diffservcode, str):
            raise TypeError("Expected argument 'diffservcode' to be a str")
        pulumi.set(__self__, "diffservcode", diffservcode)
        if fail_weight and not isinstance(fail_weight, int):
            raise TypeError("Expected argument 'fail_weight' to be a int")
        pulumi.set(__self__, "fail_weight", fail_weight)
        if failtime and not isinstance(failtime, int):
            raise TypeError("Expected argument 'failtime' to be a int")
        pulumi.set(__self__, "failtime", failtime)
        if gateway_ip and not isinstance(gateway_ip, str):
            raise TypeError("Expected argument 'gateway_ip' to be a str")
        pulumi.set(__self__, "gateway_ip", gateway_ip)
        if gateway_ip6 and not isinstance(gateway_ip6, str):
            raise TypeError("Expected argument 'gateway_ip6' to be a str")
        pulumi.set(__self__, "gateway_ip6", gateway_ip6)
        if ha_priority and not isinstance(ha_priority, int):
            raise TypeError("Expected argument 'ha_priority' to be a int")
        pulumi.set(__self__, "ha_priority", ha_priority)
        if http_agent and not isinstance(http_agent, str):
            raise TypeError("Expected argument 'http_agent' to be a str")
        pulumi.set(__self__, "http_agent", http_agent)
        if http_get and not isinstance(http_get, str):
            raise TypeError("Expected argument 'http_get' to be a str")
        pulumi.set(__self__, "http_get", http_get)
        if http_match and not isinstance(http_match, str):
            raise TypeError("Expected argument 'http_match' to be a str")
        pulumi.set(__self__, "http_match", http_match)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interval and not isinstance(interval, int):
            raise TypeError("Expected argument 'interval' to be a int")
        pulumi.set(__self__, "interval", interval)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if packet_size and not isinstance(packet_size, int):
            raise TypeError("Expected argument 'packet_size' to be a int")
        pulumi.set(__self__, "packet_size", packet_size)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if probe_count and not isinstance(probe_count, int):
            raise TypeError("Expected argument 'probe_count' to be a int")
        pulumi.set(__self__, "probe_count", probe_count)
        if probe_timeout and not isinstance(probe_timeout, int):
            raise TypeError("Expected argument 'probe_timeout' to be a int")
        pulumi.set(__self__, "probe_timeout", probe_timeout)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if recoverytime and not isinstance(recoverytime, int):
            raise TypeError("Expected argument 'recoverytime' to be a int")
        pulumi.set(__self__, "recoverytime", recoverytime)
        if routes and not isinstance(routes, list):
            raise TypeError("Expected argument 'routes' to be a list")
        pulumi.set(__self__, "routes", routes)
        if security_mode and not isinstance(security_mode, str):
            raise TypeError("Expected argument 'security_mode' to be a str")
        pulumi.set(__self__, "security_mode", security_mode)
        if server_config and not isinstance(server_config, str):
            raise TypeError("Expected argument 'server_config' to be a str")
        pulumi.set(__self__, "server_config", server_config)
        if server_lists and not isinstance(server_lists, list):
            raise TypeError("Expected argument 'server_lists' to be a list")
        pulumi.set(__self__, "server_lists", server_lists)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)
        if service_detection and not isinstance(service_detection, str):
            raise TypeError("Expected argument 'service_detection' to be a str")
        pulumi.set(__self__, "service_detection", service_detection)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 and not isinstance(source_ip6, str):
            raise TypeError("Expected argument 'source_ip6' to be a str")
        pulumi.set(__self__, "source_ip6", source_ip6)
        if srcintf and not isinstance(srcintf, str):
            raise TypeError("Expected argument 'srcintf' to be a str")
        pulumi.set(__self__, "srcintf", srcintf)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if update_cascade_interface and not isinstance(update_cascade_interface, str):
            raise TypeError("Expected argument 'update_cascade_interface' to be a str")
        pulumi.set(__self__, "update_cascade_interface", update_cascade_interface)
        if update_policy_route and not isinstance(update_policy_route, str):
            raise TypeError("Expected argument 'update_policy_route' to be a str")
        pulumi.set(__self__, "update_policy_route", update_policy_route)
        if update_static_route and not isinstance(update_static_route, str):
            raise TypeError("Expected argument 'update_static_route' to be a str")
        pulumi.set(__self__, "update_static_route", update_static_route)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> str:
        """
        Address mode (IPv4 or IPv6).
        """
        return pulumi.get(self, "addr_mode")

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> int:
        """
        Traffic class ID.
        """
        return pulumi.get(self, "class_id")

    @property
    @pulumi.getter
    def diffservcode(self) -> str:
        """
        Differentiated services code point (DSCP) in the IP header of the probe packet.
        """
        return pulumi.get(self, "diffservcode")

    @property
    @pulumi.getter(name="failWeight")
    def fail_weight(self) -> int:
        """
        Threshold weight to trigger link failure alert.
        """
        return pulumi.get(self, "fail_weight")

    @property
    @pulumi.getter
    def failtime(self) -> int:
        """
        Number of retry attempts before the server is considered down (1 - 10, default = 5)
        """
        return pulumi.get(self, "failtime")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> str:
        """
        Gateway IP address used to probe the server.
        """
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter(name="gatewayIp6")
    def gateway_ip6(self) -> str:
        """
        Gateway IPv6 address used to probe the server.
        """
        return pulumi.get(self, "gateway_ip6")

    @property
    @pulumi.getter(name="haPriority")
    def ha_priority(self) -> int:
        """
        HA election priority (1 - 50).
        """
        return pulumi.get(self, "ha_priority")

    @property
    @pulumi.getter(name="httpAgent")
    def http_agent(self) -> str:
        """
        String in the http-agent field in the HTTP header.
        """
        return pulumi.get(self, "http_agent")

    @property
    @pulumi.getter(name="httpGet")
    def http_get(self) -> str:
        """
        If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
        """
        return pulumi.get(self, "http_get")

    @property
    @pulumi.getter(name="httpMatch")
    def http_match(self) -> str:
        """
        String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
        """
        return pulumi.get(self, "http_match")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> int:
        """
        Detection interval (1 - 3600 sec, default = 5).
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Link monitor name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packetSize")
    def packet_size(self) -> int:
        """
        Packet size of a twamp test session,
        """
        return pulumi.get(self, "packet_size")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Twamp controller password in authentication mode
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Port number of the traffic to be used to monitor the server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="probeCount")
    def probe_count(self) -> int:
        """
        Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        """
        return pulumi.get(self, "probe_count")

    @property
    @pulumi.getter(name="probeTimeout")
    def probe_timeout(self) -> int:
        """
        Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).
        """
        return pulumi.get(self, "probe_timeout")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Protocols used to monitor the server.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def recoverytime(self) -> int:
        """
        Number of successful responses received before server is considered recovered (1 - 10, default = 5).
        """
        return pulumi.get(self, "recoverytime")

    @property
    @pulumi.getter
    def routes(self) -> Sequence['outputs.GetSystemLinkMonitorRouteResult']:
        """
        Subnet to monitor. The structure of `route` block is documented below.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> str:
        """
        Twamp controller security mode.
        """
        return pulumi.get(self, "security_mode")

    @property
    @pulumi.getter(name="serverConfig")
    def server_config(self) -> str:
        """
        Mode of server configuration.
        """
        return pulumi.get(self, "server_config")

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Sequence['outputs.GetSystemLinkMonitorServerListResult']:
        """
        Servers for link-monitor to monitor. The structure of `server_list` block is documented below.
        """
        return pulumi.get(self, "server_lists")

    @property
    @pulumi.getter
    def servers(self) -> Sequence['outputs.GetSystemLinkMonitorServerResult']:
        """
        IP address of the server(s) to be monitored. The structure of `server` block is documented below.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter(name="serviceDetection")
    def service_detection(self) -> str:
        """
        Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated.
        """
        return pulumi.get(self, "service_detection")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IP address used in packet to the server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> str:
        """
        Source IPv6 address used in packet to the server.
        """
        return pulumi.get(self, "source_ip6")

    @property
    @pulumi.getter
    def srcintf(self) -> str:
        """
        Interface that receives the traffic to be monitored.
        """
        return pulumi.get(self, "srcintf")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this link monitor.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateCascadeInterface")
    def update_cascade_interface(self) -> str:
        """
        Enable/disable update cascade interface.
        """
        return pulumi.get(self, "update_cascade_interface")

    @property
    @pulumi.getter(name="updatePolicyRoute")
    def update_policy_route(self) -> str:
        """
        Enable/disable updating the policy route.
        """
        return pulumi.get(self, "update_policy_route")

    @property
    @pulumi.getter(name="updateStaticRoute")
    def update_static_route(self) -> str:
        """
        Enable/disable updating the static route.
        """
        return pulumi.get(self, "update_static_route")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemLinkMonitorResult(GetSystemLinkMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemLinkMonitorResult(
            addr_mode=self.addr_mode,
            class_id=self.class_id,
            diffservcode=self.diffservcode,
            fail_weight=self.fail_weight,
            failtime=self.failtime,
            gateway_ip=self.gateway_ip,
            gateway_ip6=self.gateway_ip6,
            ha_priority=self.ha_priority,
            http_agent=self.http_agent,
            http_get=self.http_get,
            http_match=self.http_match,
            id=self.id,
            interval=self.interval,
            name=self.name,
            packet_size=self.packet_size,
            password=self.password,
            port=self.port,
            probe_count=self.probe_count,
            probe_timeout=self.probe_timeout,
            protocol=self.protocol,
            recoverytime=self.recoverytime,
            routes=self.routes,
            security_mode=self.security_mode,
            server_config=self.server_config,
            server_lists=self.server_lists,
            servers=self.servers,
            service_detection=self.service_detection,
            source_ip=self.source_ip,
            source_ip6=self.source_ip6,
            srcintf=self.srcintf,
            status=self.status,
            update_cascade_interface=self.update_cascade_interface,
            update_policy_route=self.update_policy_route,
            update_static_route=self.update_static_route,
            vdomparam=self.vdomparam)


def get_system_link_monitor(name: Optional[str] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemLinkMonitorResult:
    """
    Use this data source to get information on an fortios system linkmonitor


    :param str name: Specify the name of the desired system linkmonitor.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemLinkMonitor:GetSystemLinkMonitor', __args__, opts=opts, typ=GetSystemLinkMonitorResult).value

    return AwaitableGetSystemLinkMonitorResult(
        addr_mode=__ret__.addr_mode,
        class_id=__ret__.class_id,
        diffservcode=__ret__.diffservcode,
        fail_weight=__ret__.fail_weight,
        failtime=__ret__.failtime,
        gateway_ip=__ret__.gateway_ip,
        gateway_ip6=__ret__.gateway_ip6,
        ha_priority=__ret__.ha_priority,
        http_agent=__ret__.http_agent,
        http_get=__ret__.http_get,
        http_match=__ret__.http_match,
        id=__ret__.id,
        interval=__ret__.interval,
        name=__ret__.name,
        packet_size=__ret__.packet_size,
        password=__ret__.password,
        port=__ret__.port,
        probe_count=__ret__.probe_count,
        probe_timeout=__ret__.probe_timeout,
        protocol=__ret__.protocol,
        recoverytime=__ret__.recoverytime,
        routes=__ret__.routes,
        security_mode=__ret__.security_mode,
        server_config=__ret__.server_config,
        server_lists=__ret__.server_lists,
        servers=__ret__.servers,
        service_detection=__ret__.service_detection,
        source_ip=__ret__.source_ip,
        source_ip6=__ret__.source_ip6,
        srcintf=__ret__.srcintf,
        status=__ret__.status,
        update_cascade_interface=__ret__.update_cascade_interface,
        update_policy_route=__ret__.update_policy_route,
        update_static_route=__ret__.update_static_route,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_link_monitor)
def get_system_link_monitor_output(name: Optional[pulumi.Input[str]] = None,
                                   vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemLinkMonitorResult]:
    """
    Use this data source to get information on an fortios system linkmonitor


    :param str name: Specify the name of the desired system linkmonitor.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...

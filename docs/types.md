



## DatacenterWithConfig
A data center object that contains both metadata and configuration



```yaml
|
    kind: DatacenterWithConfig
    apiVersion: 1.0
    metadata:
      id: 332
      name: test-us03-chi-qts01-dc-3
      displayname: US03 DC
      ismaster: false
      ismaintenance: false
      type: metal_cloud
      createdtimestamp: "2023-03-06T11:12:20Z"
      updatedtimestamp: "2024-04-12T12:09:25Z"
      ishidden: false
      tags:
      - ""
    config:
      BSIMachinesSubnetIPv4CIDR: 172.18.38.16/29
      BSIVRRPListenIPv4: 172.18.38.22
      BSIMachineListenIPv4List:
      - 172.18.38.22
      BSIExternallyVisibleIPv4: 176.223.248.11
      repoURLRoot: http://repointegrationpublic.bigstepcloud.com
      repoURLRootQuarantineNetwork: http://repointegrationpublic.bigstepcloud.com
      SANRoutedSubnet: 100.96.0.0/16
      NTPServers:
      - 45.55.58.103
      DNSServers:
      - 1.1.1.1
      - 8.8.8.8
      KMS: 84.40.58.70:1688
      TFTPServerWANVRRPListenIPv4: 172.18.38.22
      dataLakeEnabled: false
      latitude: 45
      longitude: 51.509865
      serverRegisterUsingGeneratedIPMICredentialsEnabled: false
      serverRegisterUsingProvidedIPMICredentialsEnabled: true
      switchProvisioner:
        ASNRanges: []
        LAGRanges:
        - 11-19
        LANVLANRange: 200-299
        MLAGRanges:
        - 20-25
        SANVLANRange: 300-399
        SANVNIPrefix: 2
        VRFL3VNIPrefix: 9
        VRFVLANRanges: []
        WANVLANRange: 100-199
        WANVNIPrefix: 1
        allocateDefaultLANVLAN: false
        allocateDefaultSANVLAN: true
        allocateDefaultWANVLAN: true
        disableQuarantineNetwork: false
        leafSwitchesHaveMLAGPairs: false
        preventCleanupForVLANs: []
        preventCleanupForVLANsFromExternalConnectionUplinks: []
        preventCleanupForVRFs: []
        preventUsageOfVLANs:
        - 3204
        preventUsageOfVRFs: []
        quarantineVLANID: 5
        storageHasSeparateFabric: false
        type: EVPNVXLANL2Provisioner
        zeroTouchRegistrationEnabled: false
      enableTenantAccessToIPMI: true
      allowVLANOverrides: true
      allowNetworkProfiles: true
      enableServerRegistrationStartedByInBandDHCP: false
      extraInternalIPsPerSubnet: 0
      extraInternalIPsPerSANSubnet: 0
      serverRAIDConfigurationEnabled: true
      webProxy: null
      isKubernetesDeployment: false
      allowInstanceArrayFirmwarePolicies: true
      provisionUsingTheQuarantineNetwork: true
      enableDHCPRelaySecurityForQuarantineNetwork: false
      enableDHCPRelaySecurityForClientNetworks: false
      enableDHCPBMCMACAddressWhitelist: false
      dhcpBMCMACAddressWhitelist: []
      defaultServerCleanupPolicyID: 3
      defaultWANNetworkProfileID: 864
      defaultDeploymentMechanism: virtual_media
      defaultCleanupAndRegistrationMechanism: bmc
      NFSServer: 172.18.38.22
      Option82ToIPMapping:
        Eth1/1: 172.18.32.1
```



<hr />

<div class="dd">

<code>metadata</code>  <i><a href="#datacenter">Datacenter</a></i>

</div>
<div class="dt">

The datacenter part of the object

</div>

<hr />

<div class="dd">

<code>config</code>  <i><a href="#datacenterconfig">DatacenterConfig</a></i>

</div>
<div class="dt">

The datacenter configuration part of the object

</div>

<hr />





## Datacenter
Datacenter metadata

Appears in:


- <code><a href="#datacenterwithconfig">DatacenterWithConfig</a>.metadata</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of this datacenter.

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The name (label) of this datacenter. Once set it cannot be changed.

</div>

<hr />

<div class="dd">

<code>parentName</code>  <i>string</i>

</div>
<div class="dt">

The name (label) of the parent datacenter. This is useful in hierarchical setups where one datacenter needs to access it's parent's resources.

</div>

<hr />

<div class="dd">

<code>userId</code>  <i>int</i>

</div>
<div class="dt">

The owner of a datacenter.

</div>

<hr />

<div class="dd">

<code>displayName</code>  <i>string</i>

</div>
<div class="dt">

The display name of a data center. Can be changed.

</div>

<hr />

<div class="dd">

<code>isMaster</code>  <i>bool</i>

</div>
<div class="dt">

Deprecated.

</div>

<hr />

<div class="dd">

<code>isMaintenance</code>  <i>bool</i>

</div>
<div class="dt">

If set to true no new operations can happen on this datacenter.

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The datacenter type. Deprecated. Currently the only supported value is metal_cloud.


Valid values:


  - <code>metal_cloud</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the datacenter was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the datacenter was updated.

</div>

<hr />

<div class="dd">

<code>isHidden</code>  <i>bool</i>

</div>
<div class="dt">

If set the datacenter will not be visible in the UI

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

An array of tags (strings)

</div>

<hr />





## DatacenterConfig
DatacenterConfig - datacenter configuration

Appears in:


- <code><a href="#datacenterwithconfig">DatacenterWithConfig</a>.config</code>





<hr />

<div class="dd">

<code>BSIMachinesSubnetIPv4CIDR</code>  <i>string</i>

</div>
<div class="dt">

The ip address of the Global Controller. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIVRRPListenIPv4</code>  <i>string</i>

</div>
<div class="dt">

The ip address on which all datacenter agents listen for connections. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIMachineListenIPv4List</code>  <i>[]string</i>

</div>
<div class="dt">

Site Controller's secondary ip addresses. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIExternallyVisibleIPv4</code>  <i>string</i>

</div>
<div class="dt">

The agent's IP that is visible from the controller. Deprecated.

</div>

<hr />

<div class="dd">

<code>repoURLRoot</code>  <i>string</i>

</div>
<div class="dt">

The repository to use

</div>

<hr />

<div class="dd">

<code>repoURLRootQuarantineNetwork</code>  <i>string</i>

</div>
<div class="dt">

The repository to use during legacy (PXE) provisioning process. Same as repoURLRoot, with an IP address for the hostname, required in networks where DNS is not available.

</div>

<hr />

<div class="dd">

<code>SANRoutedSubnet</code>  <i>string</i>

</div>
<div class="dt">

The SAN subnet in CIDR format.

</div>

<hr />

<div class="dd">

<code>NTPServers</code>  <i>[]string</i>

</div>
<div class="dt">

IP addresses of NTP servers.

</div>

<hr />

<div class="dd">

<code>DNSServers</code>  <i>[]string</i>

</div>
<div class="dt">

IP addresses of DNS servers to be used in the DHCP response.

</div>

<hr />

<div class="dd">

<code>KMS</code>  <i>string</i>

</div>
<div class="dt">

Host (IP:port) of the Windows machine hosting the Key Management Service. Set to empty string to disable.

</div>

<hr />

<div class="dd">

<code>TFTPServerWANVRRPListenIPv4</code>  <i>string</i>

</div>
<div class="dt">

The IP of the Site Controller TFTP service used during the legacy (PXE) deployment process.

</div>

<hr />

<div class="dd">

<code>dataLakeEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, the datalake service is enabled in this environment. Deprecated


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>monitoringGraphitePlainTextSocketHost</code>  <i>string</i>

</div>
<div class="dt">

Graphite host (IPv4:port) for the plain text protocol socket. Set to empty string to disable. Deprecated

</div>

<hr />

<div class="dd">

<code>monitoringGraphiteRenderURLHost</code>  <i>string</i>

</div>
<div class="dt">

Graphite host (IPv4:port) for the HTTP Render URL API. Set to empty string to disable. Deprecated

</div>

<hr />

<div class="dd">

<code>latitude</code>  <i>float64</i>

</div>
<div class="dt">

The Datacenter's latitude. Use negative numbers for the south hemisphere

</div>

<hr />

<div class="dd">

<code>longitude</code>  <i>float64</i>

</div>
<div class="dt">

description: The data center's longitude: Use negative numbers for areas west of Greenwich (UK)

</div>

<hr />

<div class="dd">

<code>address</code>  <i>string</i>

</div>
<div class="dt">

The data center's address

</div>

<hr />

<div class="dd">

<code>serverRegisterUsingGeneratedIPMICredentialsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the system will configure a randomly generated username and password on the server's BMC(ILO/IDRAC etc.)


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>serverRegisterUsingProvidedIPMICredentialsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the system will ask for credentials during server registration.


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>switchProvisioner</code>  <i>map[string]interface{}</i>

</div>
<div class="dt">

The provisioner (fabric) to use when provisioning the network on switch devices


Valid values:


  - <code>VLAN</code>

  - <code>EVPNVXLANL2</code>

  - <code>VPLS</code>

  - <code>LAN</code>

  - <code>SDN</code>
</div>

<hr />

<div class="dd">

<code>enableTenantAccessToIPMI</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the tenants will receive credentials for accessing the server's BMC with a special user.

</div>

<hr />

<div class="dd">

<code>allowVLANOverrides</code>  <i>bool</i>

</div>
<div class="dt">

description: Allows the end-user to force a VLAN ID (or EPG in CISCO ACI environments). This enables the user to connect to pre-existing VLANs in the established infrastructure. WARNING: This enables a tenant to access unauthorized VLANs.

</div>

<hr />

<div class="dd">

<code>allowNetworkProfiles</code>  <i>bool</i>

</div>
<div class="dt">

Allows the usage of network profiles for customizing InstanceArray network connections.

</div>

<hr />

<div class="dd">

<code>enableServerRegistrationStartedByInBandDHCP</code>  <i>bool</i>

</div>
<div class="dt">

If set enables in-band triggered registration via the legacy (PXE) mechanism.

</div>

<hr />

<div class="dd">

<code>extraInternalIPsPerSubnet</code>  <i>int</i>

</div>
<div class="dt">

Extra ips to reserve on each subnet for WAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.

</div>

<hr />

<div class="dd">

<code>extraInternalIPsPerSANSubnet</code>  <i>int</i>

</div>
<div class="dt">

Extra ips to reserve on each subnet for SAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.

</div>

<hr />

<div class="dd">

<code>serverRAIDConfigurationEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If enabled RAID configurations are set on servers

</div>

<hr />

<div class="dd">

<code>webProxy</code>  <i><a href="#webproxy">WebProxy</a></i>

</div>
<div class="dt">

If configured the proxy will be used by all operations.

</div>

<hr />

<div class="dd">

<code>isKubernetesDeployment</code>  <i>bool</i>

</div>
<div class="dt">

Deprecated.

</div>

<hr />

<div class="dd">

<code>allowInstanceArrayFirmwarePolicies</code>  <i>bool</i>

</div>
<div class="dt">

If set it allows  the use of firmware policies. Note that for baselines to function this needs to be enabled.

</div>

<hr />

<div class="dd">

<code>provisionUsingTheQuarantineNetwork</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, during the legacy registration process (PXE) the system will configure special provisioning VLAN on server ports prior to performing the deployment

</div>

<hr />

<div class="dd">

<code>enableDHCPRelaySecurityForQuarantineNetwork</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, during the legacy registration process (PXE) the system will enforce DHCP option 82 security.

</div>

<hr />

<div class="dd">

<code>enableDHCPRelaySecurityForClientNetworks</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, the DHCP server will ignore requests that do not respect DHCP option 82 for regular networks.

</div>

<hr />

<div class="dd">

<code>enableDHCPBMCMACAddressWhitelist</code>  <i>bool</i>

</div>
<div class="dt">

If enabled, the DHCPBMCMACAddressWhitelist will be used to whitelist certain MAC addresses in order to ensure that only certain servers get registered during the ZTP process.

</div>

<hr />

<div class="dd">

<code>dhcpBMCMACAddressWhitelist</code>  <i>[]string</i>

</div>
<div class="dt">

The mac addresses of the servers that are to be allowed to be registered via ZTP. This is useful during initial testing.

</div>

<hr />

<div class="dd">

<code>defaultServerCleanupPolicyID</code>  <i>int</i>

</div>
<div class="dt">

If set the server cleanup policy will be the policy with the specified id instead of the default one (which is 0)

</div>

<hr />

<div class="dd">

<code>defaultWANNetworkProfileID</code>  <i>int</i>

</div>
<div class="dt">

If set, this will be the default network profile instead of no network profile.

</div>

<hr />

<div class="dd">

<code>defaultDeploymentMechanism</code>  <i>string</i>

</div>
<div class="dt">

Deployment mechanism used in case a server supports both Virtual Media and legacy (PXE).


Valid values:


  - <code>virtual_media</code>

  - <code>pxe</code>
</div>

<hr />

<div class="dd">

<code>defaultCleanupAndRegistrationMechanism</code>  <i>string</i>

</div>
<div class="dt">

The cleanup and register mechanism used in case a server supports both BMC-only and BDK mechanisms. Defaults to BMC.


Valid values:


  - <code>bmc</code>

  - <code>bdk</code>
</div>

<hr />

<div class="dd">

<code>NFSServer</code>  <i>string</i>

</div>
<div class="dt">

The NFS server to use for server OS deployment (the IP of the site controller as seen from the server's BMC). Should be an IP to avoid DNS resolutions.

</div>

<hr />

<div class="dd">

<code>Option82ToIPMapping</code>  <i>Option82ToIPMapping</i>

</div>
<div class="dt">

Can be used to set a mapping between Option82 and IPs that the DHCP server allocates to servers during registration.

</div>

<hr />





## WebProxy
Defines web proxy configuration

Appears in:


- <code><a href="#datacenterconfig">DatacenterConfig</a>.webProxy</code>





<hr />

<div class="dd">

<code>ip</code>  <i>string</i>

</div>
<div class="dt">

Ip fo the web proxy

</div>

<hr />

<div class="dd">

<code>port</code>  <i>int</i>

</div>
<div class="dt">

Port fo the web proxy

</div>

<hr />

<div class="dd">

<code>username</code>  <i>string</i>

</div>
<div class="dt">

Username of the web proxy

</div>

<hr />

<div class="dd">

<code>password</code>  <i>string</i>

</div>
<div class="dt">

Password to use for the web proxy

</div>

<hr />








## SubnetOOB
Subnet represents a subnet for OOB operations



```yaml
id: 10 # The id of the object
label: mysubnet # The label of the object

# # The Netmask to use.
# netmask: 255.255.255.192

# # The Prefix size in CIDR format. Must match the netmask
# size: 26

# # The start of the range
# rangeStart: 192.168.0.10

# # The end of the range.
# rangeEnd: 192.168.0.100
```



<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The type of the object


Valid values:


  - <code>ipv4</code>

  - <code>ipv6</code>
</div>

<hr />

<div class="dd">

<code>useForAutoAllocation</code>  <i>bool</i>

</div>
<div class="dt">

If set to `true` this subnet will be used for auto-allocation of IPs

</div>

<hr />

<div class="dd">

<code>forResourceType</code>  <i>string</i>

</div>
<div class="dt">

What type of resource to allocate this object for


Valid values:


  - <code>server</code>

  - <code>network_equipment</code>

  - <code>any</code>
</div>

<hr />

<div class="dd">

<code>blacklist</code>  <i>[]string</i>

</div>
<div class="dt">

description: Array of IPs that are to be skipped from the interval
examples:
  - value: ['192.168.0.10','192.168.0.22']


</div>

<hr />

<div class="dd">

<code>gatewayHex</code>  <i>string</i>

</div>
<div class="dt">

The Gateway in hexadecimal format. Readonly.

</div>

<hr />

<div class="dd">

<code>gateway</code>  <i>string</i>

</div>
<div class="dt">

description: The Gateway to use when allocating IPs from this subnet.
examples:
  -values:  '"192.168.0.1"'


</div>

<hr />

<div class="dd">

<code>netmaskHex</code>  <i>string</i>

</div>
<div class="dt">

The Netmask in hexadecimal format. Readonly.

</div>

<hr />

<div class="dd">

<code>netmask</code>  <i>string</i>

</div>
<div class="dt">

The Netmask to use.



Examples:


```yaml
netmask: 255.255.255.192
```


</div>

<hr />

<div class="dd">

<code>size</code>  <i>int</i>

</div>
<div class="dt">

The Prefix size in CIDR format. Must match the netmask



Examples:


```yaml
size: 26
```


</div>

<hr />

<div class="dd">

<code>rangeStartHex</code>  <i>string</i>

</div>
<div class="dt">

The start of the range in hexadecimal. Readonly.

</div>

<hr />

<div class="dd">

<code>rangeStart</code>  <i>string</i>

</div>
<div class="dt">

The start of the range



Examples:


```yaml
rangeStart: 192.168.0.10
```


</div>

<hr />

<div class="dd">

<code>rangeEndHex</code>  <i>string</i>

</div>
<div class="dt">

The end of the range in hexadecimal. Readonly

</div>

<hr />

<div class="dd">

<code>rangeEnd</code>  <i>string</i>

</div>
<div class="dt">

The end of the range.



Examples:


```yaml
rangeEnd: 192.168.0.100
```


</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

The data center in which this subnet is valid

</div>

<hr />








## Network
Network object describes an high level connection construct






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The Label of the object.

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

A unique string representing the network. Deprecated

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The network type


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The infrastructure on which this network is defined

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network was last updated.

</div>

<hr />

<div class="dd">

<code>LANAutoAllocateIPs</code>  <i>bool</i>

</div>
<div class="dt">

If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#networkoperation">NetworkOperation</a></i>

</div>
<div class="dt">

The operation object.

</div>

<hr />





## NetworkOperation
NetworkOperation object describes the change(s) required to be applied to a Network

Appears in:


- <code><a href="#network">Network</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The Label of the object.

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

A unique string representing the network. Deprecated

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The network type


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The infrastructure on which this network is defined

</div>

<hr />

<div class="dd">

<code>LANAutoAllocateIPs</code>  <i>bool</i>

</div>
<div class="dt">

If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.

</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The deploy type
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the change operation. Readonly.

</div>

<hr />








## NetworkProfile
NetworkProfile  A network profile modifies the default network configuration of an instance array when attached to a specific Network.






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the network profile

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the network profile

</div>

<hr />

<div class="dd">

<code>dc</code>  <i>string</i>

</div>
<div class="dt">

The label of the datacenter on which this network profile applies

</div>

<hr />

<div class="dd">

<code>networkProfileIsPublic</code>  <i>bool</i>

</div>
<div class="dt">

description: If set to true any of the users can use this network profile. If set to false the network profile needs to be
explicitly allowed in the user limits for a user to be able to use it


</div>

<hr />

<div class="dd">

<code>networkType</code>  <i>string</i>

</div>
<div class="dt">

Type of network to which this network profile can be applied


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>vlans</code>  <i>[]<a href="#networkprofilevlan">NetworkProfileVLAN</a></i>

</div>
<div class="dt">

VLAN (L2 network) entries

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network profile was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network profile was updated.

</div>

<hr />





## NetworkProfileVLAN
NetworkProfileVLAN object describes an overlay L2 network which usually translates into a VLAN and
associated VXLAN VNIs or EPGs depending on the vendor


Appears in:


- <code><a href="#networkprofile">NetworkProfile</a>.vlans</code>





<hr />

<div class="dd">

<code>vlanID</code>  <i>int</i>

</div>
<div class="dt">

description: the VLAN ID to use (on EVPNVXLAN provisioner this translates to a VLAN ID to be used as VTEP for VXLAN).
On VLAN provisioner the VLAN will be used end-to-end. On SDN provisioner with CISCO ACI it will be translated into a VLAN attached to an EPG.
If set to null it will be automatically allocated based on the settings in the Datacenter config object.


</div>

<hr />

<div class="dd">

<code>label</code>  <i>int</i>

</div>
<div class="dt">

the label of this VLAN entry

</div>

<hr />

<div class="dd">

<code>portMode</code>  <i>string</i>

</div>
<div class="dt">

Type of VLAN


Valid values:


  - <code>native</code>

  - <code>trunk</code>

  - <code>access</code>
</div>

<hr />

<div class="dd">

<code>provisionSubnetGateways</code>  <i>bool</i>

</div>
<div class="dt">

description: if true will provision subnet gateways on the switches. This depends on the provisioner and vendor: Symmetric IRB, VRRP-based redundant gateway configuration etc.

</div>

<hr />

<div class="dd">

<code>provisionVXLAN</code>  <i>bool</i>

</div>
<div class="dt">

if false it will not configure the vxlan tunnel. Equivalent to the VLAN provisioner.

</div>

<hr />

<div class="dd">

<code>extConnectionIDs</code>  <i>[]int</i>

</div>
<div class="dt">

If any external connection id is configured in the array the L2 network will be extended to that external connection point

</div>

<hr />

<div class="dd">

<code>subnetPools</code>  <i>[]<a href="#networkprofilesubnetpool">NetworkProfileSubnetPool</a></i>

</div>
<div class="dt">

Subnet pools to allocate IPs from

</div>

<hr />

<div class="dd">

<code>vlanAutoAllocationIndex</code>  <i>int</i>

</div>
<div class="dt">

If set to a non-null value this will help ensure that two different network profiles can get the same auto-allocated VLAN ID

</div>

<hr />

<div class="dd">

<code>vrfID</code>  <i>int</i>

</div>
<div class="dt">

The VRF ID to use. If set to null it will fall back to automatically allocated VRFs

</div>

<hr />





## NetworkProfileSubnetPool

Appears in:


- <code><a href="#networkprofilevlan">NetworkProfileVLAN</a>.subnetPools</code>





<hr />

<div class="dd">

<code>subnetPoolID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the subnet pool to use. Set to null for auto in which case the type of subnet pool will be used to pick a subnet pool

</div>

<hr />

<div class="dd">

<code>subnetPoolType</code>  <i>string</i>

</div>
<div class="dt">

description: The type of subnet pool to use
- ipv4
- ipv6


</div>

<hr />

<div class="dd">

<code>subnetPoolProvidesDefaultRoute</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this subnet pool will be used to set the default route

</div>

<hr />








## Infrastructure
description - The infrastructure object. All other client-side objects (InstanceArrays, DriveArrays) fall under it in the object hierarchy






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

The datacenter label on which this infrastructure resides

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An automatically generated string that represents this infrastructure. Deprecated.

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the user that is the owner of the infrastructure.

</div>

<hr />

<div class="dd">

<code>ownerEmail</code>  <i>string</i>

</div>
<div class="dt">

The email of the user that is the owner of the infrastructure. Readonly

</div>

<hr />

<div class="dd">

<code>touchUnixTime</code>  <i>string</i>

</div>
<div class="dt">

Internal.

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the infrastructure was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the infrastructure was last updated.

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the ongoing change operation. readonly

</div>

<hr />

<div class="dd">

<code>deployID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the current deploy operation (AFC group id). Readonly.

</div>

<hr />

<div class="dd">

<code>designIsLocked</code>  <i>bool</i>

</div>
<div class="dt">

Set if the infrastructure cannot be changed.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#infrastructureoperation">InfrastructureOperation</a></i>

</div>
<div class="dt">

The operation object

</div>

<hr />

<div class="dd">

<code></code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables at the infrastructure level. They will be pushed to all objects (such as operating systems) and workflows.

</div>

<hr />





## InfrastructureOperation
InfrastructureOperation - object with alternations to be applied

Appears in:


- <code><a href="#infrastructure">Infrastructure</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

The datacenter label on which this infrastructure resides

</div>

<hr />

<div class="dd">

<code>deployStatus</code>  <i>string</i>

</div>
<div class="dt">

The status of the ongoing deployment operation (if any)


Valid values:


  - <code>not_started</code>

  - <code>ongoing</code>

  - <code>finished</code>
</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The type of the current deploy operation (if any)
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An automatically generated string that represents this infrastructure. Deprecated.

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the user that is the owner of the infrastructure.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the infrastructure was last updated.

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the current change operation. Readonly.

</div>

<hr />

<div class="dd">

<code>deployID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the current deploy operation (AFC group id). Readonly.

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables at the infrastructure level. They will be pushed to all objects (such as operating systems) and workflows.

</div>

<hr />








## InstanceArray
InstanceArray object describes a collection of identical instances






<hr />

<div class="dd">

<code>instanceID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

description: |
		User editable DNS record that gets created for this instance array in the built-in DNS
		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
     as:  only "-", lowercase alphanumeric characters and not start with a number.
     Pattern:


</div>

<hr />

<div class="dd">

<code>bootMethod</code>  <i>string</i>

</div>
<div class="dt">

description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
values:
	- local_disks
 - pxe_iscsi


</div>

<hr />

<div class="dd">

<code>instanceCount</code>  <i>int</i>

</div>
<div class="dt">

The number of instances in this array.

</div>

<hr />

<div class="dd">

<code>ramGBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.

</div>

<hr />

<div class="dd">

<code>processorCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.

</div>

<hr />

<div class="dd">

<code>processorCoreMhz</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.

</div>

<hr />

<div class="dd">

<code>processorCoreCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).

</div>

<hr />

<div class="dd">

<code>diskCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.

</div>

<hr />

<div class="dd">

<code>diskSizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.

</div>

<hr />

<div class="dd">

<code>diskTypes</code>  <i>[]string</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.


Valid values:


  - <code>ssh</code>

  - <code>hdd</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The id of the infrastructure on which this infrastructure is created on.

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>interfaces</code>  <i>[]<a href="#instancearrayinterface">InstanceArrayInterface</a></i>

</div>
<div class="dt">

The instance array interfaces configuration

</div>

<hr />

<div class="dd">

<code>clusterID</code>  <i>int</i>

</div>
<div class="dt">

The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.

</div>

<hr />

<div class="dd">

<code>clusterRoleGroup</code>  <i>string</i>

</div>
<div class="dt">

If part of an app cluster this field will receive the role that this instance array has such as `master` or `worker` which is application specific.

</div>

<hr />

<div class="dd">

<code>firewallManaged</code>  <i>bool</i>

</div>
<div class="dt">

description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.


</div>

<hr />

<div class="dd">

<code>firewallRules</code>  <i>[]FirewallRule</i>

</div>
<div class="dt">

The list of firewall rules to configure

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

The operating system template to use.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#instancearrayoperation">InstanceArrayOperation</a></i>

</div>
<div class="dt">

Used when changing an instance array. It captures the operation that needs to happen on the instance array.

</div>

<hr />

<div class="dd">

<code>additionalWanIPv4</code>  <i>string</i>

</div>
<div class="dt">

Information about additional ips to be assigned to the WAN interfaces. Used internally.

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables and variable overrides to be pushed to the operating system deployment process.

</div>

<hr />

<div class="dd">

<code>firmwarePolicies</code>  <i>[]int</i>

</div>
<div class="dt">

Firmware policies to apply. Deprecated. Use baselines.

</div>

<hr />

<div class="dd">

<code>drive_array_id_boot</code>  <i>int</i>

</div>
<div class="dt">

When iSCSI boot is used this is the id of the drive array that will be the boot device.

</div>

<hr />





## InstanceArrayOperation
InstanceArrayOperation object describes the changes that will be applied to an instance array

Appears in:


- <code><a href="#instancearray">InstanceArray</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

description: |
		User editable DNS record that gets created for this instance array in the built-in DNS
		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
     as:  only "-", lowercase alphanumeric characters and not start with a number.
     Pattern:


</div>

<hr />

<div class="dd">

<code>bootMethod</code>  <i>string</i>

</div>
<div class="dt">

description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
values:
	- local_disks
 - pxe_iscsi


</div>

<hr />

<div class="dd">

<code>instanceCount</code>  <i>int</i>

</div>
<div class="dt">

The number of instances in this array.

</div>

<hr />

<div class="dd">

<code>ramGBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.

</div>

<hr />

<div class="dd">

<code>processorCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.

</div>

<hr />

<div class="dd">

<code>processorCoreMhz</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.

</div>

<hr />

<div class="dd">

<code>processorCoreCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).

</div>

<hr />

<div class="dd">

<code>diskCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.

</div>

<hr />

<div class="dd">

<code>diskSizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.

</div>

<hr />

<div class="dd">

<code>diskTypes</code>  <i>[]string</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.


Valid values:


  - <code>ssh</code>

  - <code>hdd</code>
</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>interfaces</code>  <i>[]<a href="#instancearrayinterfaceoperation">InstanceArrayInterfaceOperation</a></i>

</div>
<div class="dt">

The instance array interfaces configuration

</div>

<hr />

<div class="dd">

<code>clusterID</code>  <i>int</i>

</div>
<div class="dt">

The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.

</div>

<hr />

<div class="dd">

<code>clusterRoleGroup</code>  <i>string</i>

</div>
<div class="dt">

If part of an app cluster this field will receive the role that this instance array has such as master or worker which is application specific.

</div>

<hr />

<div class="dd">

<code>firewallManaged</code>  <i>bool</i>

</div>
<div class="dt">

description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.


</div>

<hr />

<div class="dd">

<code>firewallRules</code>  <i>[]FirewallRule</i>

</div>
<div class="dt">

The list of firewall rules to configure

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

The operating system template to use.

</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The deploy type, one of:
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>deployStatus</code>  <i>string</i>

</div>
<div class="dt">

The status of the deployment


Valid values:


  - <code>not_started</code>

  - <code>ongoing</code>

  - <code>finished</code>
</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the change operation. Readonly.

</div>

<hr />

<div class="dd">

<code>additionalWanIPv4</code>  <i>string</i>

</div>
<div class="dt">

Information about additional ips to be assigned to the WAN interfaces. Used internally.

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables and variable overrides to be pushed to the operating system deployment process.

</div>

<hr />

<div class="dd">

<code>firmwarePolicies</code>  <i>[]int</i>

</div>
<div class="dt">

Firmware policies to apply. Deprecated. Use baselines instead.

</div>

<hr />

<div class="dd">

<code>drive_array_id_boot</code>  <i>int</i>

</div>
<div class="dt">

When iSCSI boot is used this is the id of the drive array that will be the boot device.

</div>

<hr />





## InstanceArrayInterface
InstanceArrayInterface describes a network interface of the array.
It's properties will be applied to all InstanceInterfaces of the array's instances.


Appears in:


- <code><a href="#instancearray">InstanceArray</a>.interfaces</code>





<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the interface

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An unique string describing the interface.

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Interface ID. A unique id of the interface.

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>int</i>

</div>
<div class="dt">

The instance array to which this interface is associated to

</div>

<hr />

<div class="dd">

<code>networkID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.

</div>

<hr />

<div class="dd">

<code>LAGGIndexes</code>  <i>[]interface{}</i>

</div>
<div class="dt">

Used internally. Readonly.

</div>

<hr />

<div class="dd">

<code>index</code>  <i>int</i>

</div>
<div class="dt">

description: |
	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.


</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The creation date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The last update date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#instancearrayinterfaceoperation">InstanceArrayInterfaceOperation</a></i>

</div>
<div class="dt">

The operation object. Must be set to alter the configuration of an interface.

</div>

<hr />

<div class="dd">

<code>instance_array_interface_change_id</code>  <i>int</i>

</div>
<div class="dt">

Ongoing change ID. Readonly.

</div>

<hr />





## InstanceArrayInterfaceOperation
InstanceArrayInterfaceOperation describes changes to a network array interface

Appears in:


- <code><a href="#instancearrayoperation">InstanceArrayOperation</a>.interfaces</code>

- <code><a href="#instancearrayinterface">InstanceArrayInterface</a>.operation</code>





<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the interface

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An unique string describing the interface.

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Interface ID. A unique id of the interface.

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>int</i>

</div>
<div class="dt">

The instance array to which this interface is associated to

</div>

<hr />

<div class="dd">

<code>networkID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.

</div>

<hr />

<div class="dd">

<code>LAGGIndexes</code>  <i>[]interface{}</i>

</div>
<div class="dt">

Used internally. Readonly.

</div>

<hr />

<div class="dd">

<code>index</code>  <i>int</i>

</div>
<div class="dt">

description: |
	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.


</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The creation date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The last update date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

Ongoing change ID. Readonly.

</div>

<hr />








## DriveArray
DriveArray represents a collection of identical drives






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

When the drive array is used with images (such as for diskless boot) this field will keep the Volume Template ID of template to use.

</div>

<hr />

<div class="dd">

<code>storageType</code>  <i>string</i>

</div>
<div class="dt">

What type of storage to use. This string needs to match the type of storage exported by a storage pool

</div>

<hr />

<div class="dd">

<code>sizeMBytesDefault</code>  <i>int</i>

</div>
<div class="dt">

The size of the volume. Note that this value is used when creating the drive array but also when expanding the drive array (increasing the count).

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the instance array to which this drive array is connected. Can be 0 if not connected.

</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the infrastructure to which this drive array belongs.

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The ID of the infrastructure to which this drive array belongs.

</div>

<hr />

<div class="dd">

<code>count</code>  <i>int</i>

</div>
<div class="dt">

The number of drives in this drive array.

</div>

<hr />

<div class="dd">

<code>expandWithInstanceArray</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the number of drives in the drive array will expand if the attached instance array expands (instance count increases).

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#drivearrayoperation">DriveArrayOperation</a></i>

</div>
<div class="dt">

Operation object of the drive array

</div>

<hr />

<div class="dd">

<code>ioLimit</code>  <i>string</i>

</div>
<div class="dt">

Only valid for certain storage systems (such as Dell Unity).

</div>

<hr />

<div class="dd">

<code>storagePoolID</code>  <i>int</i>

</div>
<div class="dt">

If set the specified storage pool will be used to allocate the drive instead of automatic allocation.

</div>

<hr />

<div class="dd">

<code>affinity</code>  <i>string</i>

</div>
<div class="dt">

Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools


Valid values:


  - <code>same_storage</code>

  - <code>different_storage</code>
</div>

<hr />





## DriveArrayOperation
DriveArrayOperation defines changes to be applied to a DriveArray

Appears in:


- <code><a href="#drivearray">DriveArray</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

When the drive array is used with images (such as for diskless boot) this field will keep the Volume Template ID of template to use.

</div>

<hr />

<div class="dd">

<code>storageType</code>  <i>string</i>

</div>
<div class="dt">

What type of storage to use. This string needs to match the type of storage exported by a storage pool

</div>

<hr />

<div class="dd">

<code>sizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

The size of the volume. Note that this value is used when creating the drive array but also when expanding the drive array (increasing the count).

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>interface{}</i>

</div>
<div class="dt">

The ID of the instance array to which this drive array is connected. Can be 0 if not connected.

</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the infrastructure to which this drive array belongs.

</div>

<hr />

<div class="dd">

<code>count</code>  <i>int</i>

</div>
<div class="dt">

The ID of the infrastructure to which this drive array belongs.

</div>

<hr />

<div class="dd">

<code>expandWithInstanceArray</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the number of drives in the drive array will expand if the attached instance array expands (instance count increases).

</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The deploy type
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>deployStatus</code>  <i>string</i>

</div>
<div class="dt">

The status of the deployment


Valid values:


  - <code>not_started</code>

  - <code>ongoing</code>

  - <code>finished</code>
</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the change operation. Readonly.

</div>

<hr />

<div class="dd">

<code>ioLimit</code>  <i>string</i>

</div>
<div class="dt">

Only valid for certain storage systems (such as Dell Unity).

</div>

<hr />

<div class="dd">

<code>storagePoolID</code>  <i>int</i>

</div>
<div class="dt">

If set the specified storage pool will be used to allocate the drive instead of automatic allocation.

</div>

<hr />

<div class="dd">

<code>affinity</code>  <i>string</i>

</div>
<div class="dt">

Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools


Valid values:


  - <code>same_storage</code>

  - <code>different_storage</code>
</div>

<hr />








## Network
Network object describes an high level connection construct






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The Label of the object.

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

A unique string representing the network. Deprecated

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The network type


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The infrastructure on which this network is defined

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the network was last updated.

</div>

<hr />

<div class="dd">

<code>LANAutoAllocateIPs</code>  <i>bool</i>

</div>
<div class="dt">

If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#networkoperation">NetworkOperation</a></i>

</div>
<div class="dt">

The operation object.

</div>

<hr />





## NetworkOperation
NetworkOperation object describes the change(s) required to be applied to a Network

Appears in:


- <code><a href="#network">Network</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The Label of the object.

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

A unique string representing the network. Deprecated

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The network type


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The infrastructure on which this network is defined

</div>

<hr />

<div class="dd">

<code>LANAutoAllocateIPs</code>  <i>bool</i>

</div>
<div class="dt">

If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.

</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The deploy type
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the change operation. Readonly.

</div>

<hr />








## Variable
Variable struct defines a Variable type






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

The id of the owner (user object) of the object

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

The id of the user that is currently manipulating the object. Readonly.

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The name of the variable

</div>

<hr />

<div class="dd">

<code>usage</code>  <i>string</i>

</div>
<div class="dt">

The usage of a variable

</div>

<hr />

<div class="dd">

<code>json</code>  <i>string</i>

</div>
<div class="dt">

The content of the variable in json encoded format

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp of the variable creation date. Readonly

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp of the variable last update. Readonly

</div>

<hr />








## OSAsset
OSAsset struct defines a server type






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

Id of the owner of this object

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

Id of the user who created this object (may be different than the owner). Readonly.

</div>

<hr />

<div class="dd">

<code>fileName</code>  <i>string</i>

</div>
<div class="dt">

The file name as it will be seen in the final template

</div>

<hr />

<div class="dd">

<code>fileSizeBytes</code>  <i>int</i>

</div>
<div class="dt">

The size of the file

</div>

<hr />

<div class="dd">

<code>fileMime</code>  <i>string</i>

</div>
<div class="dt">

Type of file as a mime type. Both types are stored in the database but the text type will be interpreted for variables.


Valid values:


  - <code>application/octet-stream</code>

  - <code>text/plain</code>
</div>

<hr />

<div class="dd">

<code>templateType</code>  <i>string</i>

</div>
<div class="dt">

Type of asset. Required for text type files. It controls the type of template language used when interpreting the files. Simple - uses simple search and replace. Advanced - uses jinja2 template language.


Valid values:


  - <code>none</code>

  - <code>simple</code>

  - <code>advanced</code>
</div>

<hr />

<div class="dd">

<code>contentBase64</code>  <i>string</i>

</div>
<div class="dt">

the actual contents in base64 format.

</div>

<hr />

<div class="dd">

<code>contentSHA256Hex</code>  <i>string</i>

</div>
<div class="dt">

A hash of the content

</div>

<hr />

<div class="dd">

<code>usage</code>  <i>string</i>

</div>
<div class="dt">

description: The asset's intended usage
values:
	 - build_component


</div>

<hr />

<div class="dd">

<code>sourceURL</code>  <i>string</i>

</div>
<div class="dt">

If set then the content of this asset will not be stored in the database. Instead it will be pulled from this URL when needed.

</div>

<hr />

<div class="dd">

<code>variables</code>  <i>[]string</i>

</div>
<div class="dt">

A list of required custom variables for the asset to be generated. Automatically generated by parsing the file.

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

A list o tags

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the asset was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the asset was last updated.

</div>

<hr />








## OSTemplate
OSTemplate A template can be created based on a drive and it has the same characteristics and holds the same information as the parent drive.






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The display name of the object

</div>

<hr />

<div class="dd">

<code>sizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

If this is an image type os template, the size of the image

</div>

<hr />

<div class="dd">

<code>localDisk</code>  <i>bool</i>

</div>
<div class="dt">

If set to true if this template can be installed to local disks

</div>

<hr />

<div class="dd">

<code>isOsTemplate</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this is an unattended install type template. If set to false this will be an image type template.

</div>

<hr />

<div class="dd">

<code></code>  <i>bool</i>

</div>
<div class="dt">

If set to true this is a template for switches rather than servers.

</div>

<hr />

<div class="dd">

<code>isImageBuildRequired</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this is a template that requires a build using the image builder to package assets into the ISO

</div>

<hr />

<div class="dd">

<code>provisionViaOOB</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this template can be installed via the OOB network (Virtual Media)

</div>

<hr />

<div class="dd">

<code>bootMethods</code>  <i>string</i>

</div>
<div class="dt">

If set to true this template can be installed via the OOB network (Virtual Media)


Valid values:


  - <code>pxe_iscsi</code>

  - <code>local_drives</code>
</div>

<hr />

<div class="dd">

<code></code>  <i>string</i>

</div>
<div class="dt">

description: Defines how the template is initialized.
values:
 - "provisioner_os_bootstrap_dummy",
 - "provisioner_os_bootstrap_generic",
 - "provisioner_os_bootstrap_centos",
 - "provisioner_os_bootstrap_ubuntu",
 - "provisioner_os_cloudinit_prepare_centos",
 - "provisioner_os_cloudinit_prepare_rhel",
 - "provisioner_os_cloudinit_prepare_ubuntu",
 - "provisioner_os_cloudbaseinit_prepare_windows",
 - "provisioner_os_cloudinit_prepare_generic_unix",
 - "provisioner_os_bootstrap_generic_non_cloudinit


</div>

<hr />

<div class="dd">

<code>bootType</code>  <i>string</i>

</div>
<div class="dt">

Controls the type of boot method supported by this template


Valid values:


  - <code>legacy_only</code>

  - <code>uefi_only</code>

  - <code>hybrid</code>
</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

A description of the template

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the template was created. Readonly.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the template was updated. Readonly.

</div>

<hr />

<div class="dd">

<code>userID</code>  <i>int</i>

</div>
<div class="dt">

The user that owns this template

</div>

<hr />

<div class="dd">

<code>os</code>  <i>OperatingSystem</i>

</div>
<div class="dt">

The operating system details

</div>

<hr />

<div class="dd">

<code>repoURL</code>  <i>string</i>

</div>
<div class="dt">

In some cases additional files are required and these are typically pulled from this URL

</div>

<hr />

<div class="dd">

<code>deprecationStatus</code>  <i>string</i>

</div>
<div class="dt">

Controls if this template can be used for new provisioning operations or just expanding of existing clusters


Valid values:


  - <code>not_deprecated</code>

  - <code>deprecated_allow_expand</code>

  - <code>deprecated_deny_provision</code>
</div>

<hr />

<div class="dd">

<code>credentials</code>  <i><a href="#ostemplatecredentials">OSTemplateCredentials</a></i>

</div>
<div class="dt">

Credentials of the template

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

Tags

</div>

<hr />

<div class="dd">

<code>OSAssetIDBootloaderLocalInstall</code>  <i>int</i>

</div>
<div class="dt">

The ID of the bootloader asset to use during local install

</div>

<hr />

<div class="dd">

<code>OSAssetIDBootloaderOSBoot</code>  <i>int</i>

</div>
<div class="dt">

The ID of the bootloader asset to use during initial boot (such as iPXE)

</div>

<hr />

<div class="dd">

<code>variablesJSON</code>  <i>string</i>

</div>
<div class="dt">

Variables this template requires.

</div>

<hr />

<div class="dd">

<code>networkOS</code>  <i>NetworkOperatingSystem</i>

</div>
<div class="dt">

If this is an network operating system, here are the details

</div>

<hr />

<div class="dd">

<code></code>  <i>string</i>

</div>
<div class="dt">

Version of the operating system

</div>

<hr />

<div class="dd">

<code></code>  <i>string</i>

</div>
<div class="dt">

The method to determine if the OS is ready to use.


Valid values:


  - <code>wait_for_ssh</code>

  - <code>wait_for_signal_from_os</code>

  - <code>wait_for_power_off</code>

  - <code>wait_for_signal_from_os_and_wait_for_power_off</code>
</div>

<hr />





## OSTemplateCredentials
OSTemplateCredentials holds information needed to connect to an OS installed by an OSTemplate.

Appears in:


- <code><a href="#ostemplate">OSTemplate</a>.credentials</code>





<hr />

<div class="dd">

<code>initialUser</code>  <i>string</i>

</div>
<div class="dt">

The initial user

</div>

<hr />

<div class="dd">

<code>initialPasswordEncrypted</code>  <i>string</i>

</div>
<div class="dt">

The initial password. Readonly

</div>

<hr />

<div class="dd">

<code>initialPassword</code>  <i>string</i>

</div>
<div class="dt">

The initial password. Write-only

</div>

<hr />

<div class="dd">

<code>initialSSHPort</code>  <i>int</i>

</div>
<div class="dt">

The port to use for SSH

</div>

<hr />

<div class="dd">

<code>changePasswordAfterDeploy</code>  <i>bool</i>

</div>
<div class="dt">

If set, and if the template supports it the password will be changed after deploy (this password will be used only during deploy).

</div>

<hr />

<div class="dd">

<code>useAutogeneratedInitialPassword</code>  <i>bool</i>

</div>
<div class="dt">

If set, the password will be automatically generated

</div>

<hr />








## Secret
Secret struct defines a server type






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

The id of the owner of the object

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

The id of the user that created the object

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The name (label) of the secret

</div>

<hr />

<div class="dd">

<code>usage</code>  <i>string</i>

</div>
<div class="dt">

How this secret is to be used

</div>

<hr />

<div class="dd">

<code>base64</code>  <i>string</i>

</div>
<div class="dt">

The content of the secret. Readonly.

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the secret was created. Readonly.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the secret was updated. Readonly.

</div>

<hr />








## Server
Server represents a server in a datacenter. Many of these fields are readonly.
Servers should always be registered via the registration process but in certain circumstances such as brownfield setups
it may be required to create these entries by hand.







<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

id of the server

</div>

<hr />

<div class="dd">

<code>UUID</code>  <i>string</i>

</div>
<div class="dt">

UUID of the server. Readonly.

</div>

<hr />

<div class="dd">

<code>status</code>  <i>string</i>

</div>
<div class="dt">

Status of a server. Note that not all transitions between states are possible.


Valid values:


  - <code>registering</code>

  - <code>available</code>

  - <code>cleaning_required</code>

  - <code>cleaning</code>

  - <code>used</code>

  - <code>used_registering</code>

  - <code>unavailable</code>

  - <code>defective</code>

  - <code>removed_from_rack</code>

  - <code>decommissioned</code>

  - <code>updating_firmware</code>

  - <code>used_diagnostics</code>
</div>

<hr />

<div class="dd">

<code>serialNumber</code>  <i>string</i>

</div>
<div class="dt">

serial number

</div>

<hr />

<div class="dd">

<code>vendor</code>  <i>string</i>

</div>
<div class="dt">

Vendor. Must be one of the supported values.


Valid values:


  - <code>HP</code>

  - <code>HPE</code>

  - <code>Dell Inc.</code>

  - <code>Lenovo</code>

  - <code>Supermicro</code>

  - <code>BULL</code>

  - <code>QEMU</code>

  - <code>BSI</code>
</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

Datacenter where this datacenter is registered

</div>

<hr />

<div class="dd">

<code>networkTotalCapacityMbps</code>  <i>int</i>

</div>
<div class="dt">

Summed up value of total network capacities of all in-band network connections of the server.

</div>

<hr />

<div class="dd">

<code>bootType</code>  <i>string</i>

</div>
<div class="dt">

Type of boot this server uses


Valid values:


  - <code>classic</code>

  - <code>uefi</code>
</div>

<hr />

<div class="dd">

<code>powerStatus</code>  <i>string</i>

</div>
<div class="dt">

Server power status


Valid values:


  - <code>on</code>

  - <code>off</code>

  - <code>unknown</code>
</div>

<hr />

<div class="dd">

<code>processorName</code>  <i>string</i>

</div>
<div class="dt">

Server processor name

</div>

<hr />

<div class="dd">

<code>processorCoreCount</code>  <i>int</i>

</div>
<div class="dt">

Server total core count

</div>

<hr />

<div class="dd">

<code>processorCoreMhz</code>  <i>int</i>

</div>
<div class="dt">

Server maximum frequency

</div>

<hr />

<div class="dd">

<code>processorCount</code>  <i>int</i>

</div>
<div class="dt">

Server CPU count

</div>

<hr />

<div class="dd">

<code>processorThreads</code>  <i>int</i>

</div>
<div class="dt">

Server threads per core count

</div>

<hr />

<div class="dd">

<code>processorCPUMark</code>  <i>int</i>

</div>
<div class="dt">

Deprecated

</div>

<hr />

<div class="dd">

<code>ramGbytes</code>  <i>int</i>

</div>
<div class="dt">

Total amount of RAM in Gbytes

</div>

<hr />

<div class="dd">

<code>disks</code>  <i>[]<a href="#serverdisk">ServerDisk</a></i>

</div>
<div class="dt">

The list of disks

</div>

<hr />

<div class="dd">

<code>diskCount</code>  <i>int</i>

</div>
<div class="dt">

The count of disks

</div>

<hr />

<div class="dd">

<code>diskSizeMbytes</code>  <i>int</i>

</div>
<div class="dt">

If all disks are identical this is the size of the disk. Deprecated.

</div>

<hr />

<div class="dd">

<code>diskType</code>  <i>string</i>

</div>
<div class="dt">

If all disks are identical this type of disk. Deprecated.


Valid values:


  - <code>none</code>

  - <code>NVME</code>

  - <code>HDD</code>

  - <code>SSD</code>
</div>

<hr />

<div class="dd">

<code>rackName</code>  <i>string</i>

</div>
<div class="dt">

Name of rack (metadata)

</div>

<hr />

<div class="dd">

<code>rackPositionLowerUnit</code>  <i>string</i>

</div>
<div class="dt">

Rack position lower U (metadata)

</div>

<hr />

<div class="dd">

<code>rackPositionUpperUnit</code>  <i>string</i>

</div>
<div class="dt">

Rack position upper U (metadata)

</div>

<hr />

<div class="dd">

<code>rackID</code>  <i>string</i>

</div>
<div class="dt">

Rack ID (metadata)

</div>

<hr />

<div class="dd">

<code>chassisRackID</code>  <i>int</i>

</div>
<div class="dt">

ID of the chassis if any

</div>

<hr />

<div class="dd">

<code>inventoryId</code>  <i>string</i>

</div>
<div class="dt">

Inventory ID (metadata)

</div>

<hr />

<div class="dd">

<code>productName</code>  <i>string</i>

</div>
<div class="dt">

Name of the server as returned by the server itself

</div>

<hr />

<div class="dd">

<code>serverClass</code>  <i>string</i>

</div>
<div class="dt">

Class of server. Deprecated

</div>

<hr />

<div class="dd">

<code>serverTypeID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the server type. Note that the server types are normally automatically determined during registration. However this value can be changed by the client if needed although it is not recommended.

</div>

<hr />

<div class="dd">

<code>interfaces</code>  <i>[]<a href="#serverinterface">ServerInterface</a></i>

</div>
<div class="dt">

Descriptions of the server's interfaces

</div>

<hr />

<div class="dd">

<code>supportsOOBProvisioning</code>  <i>bool</i>

</div>
<div class="dt">

Set to true if the server can be deployed via the out-of-band network (currently via virtual media)

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

Tags (metadata)

</div>

<hr />

<div class="dd">

<code>IPMIChannel</code>  <i>int</i>

</div>
<div class="dt">

The IPMI channel to use when communicating on the IPMI protocol. Used in Legacy scenarios.

</div>

<hr />

<div class="dd">

<code>IPMIHostname</code>  <i>string</i>

</div>
<div class="dt">

The IPMI host  to use when communicating on the IPMI protocol. Used in Legacy scenarios.

</div>

<hr />

<div class="dd">

<code>IPMIUsername</code>  <i>string</i>

</div>
<div class="dt">

The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios.

</div>

<hr />

<div class="dd">

<code>IPMIPassword</code>  <i>string</i>

</div>
<div class="dt">

The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios. Use this field to set the password. Write-only.

</div>

<hr />

<div class="dd">

<code>IPMIPasswordEncrypted</code>  <i>string</i>

</div>
<div class="dt">

The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios. Readonly. Encrypted.

</div>

<hr />

<div class="dd">

<code>IPMICredentialsNeedUpdate</code>  <i>bool</i>

</div>
<div class="dt">

Internal field denoting that the credentials will be updated.

</div>

<hr />

<div class="dd">

<code>vendorSKU</code>  <i>string</i>

</div>
<div class="dt">

SKU ID returned by the server. Note that not all vendors return this.

</div>

<hr />

<div class="dd">

<code>comments</code>  <i>string</i>

</div>
<div class="dt">

Comments on the server (metadata)

</div>

<hr />

<div class="dd">

<code>BIOSInfoJson</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>CustomJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>infoJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>serverDetailsXML</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>instanceCustomJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>supportsSOL</code>  <i>bool</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>ILOResetTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>BootLastUpdateTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp value denoting the last known boot.

</div>

<hr />

<div class="dd">

<code>PowerStatusUpdateTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp value denoting the last power status update.

</div>

<hr />

<div class="dd">

<code>subnetOOBID</code>  <i>int</i>

</div>
<div class="dt">

What OOB subnet is used to allocate ips to the BMC of this server

</div>

<hr />

<div class="dd">

<code>subnetDHCPStatus</code>  <i>string</i>

</div>
<div class="dt">

DHCP server's behavior in relation to this server.


Valid values:


  - <code>quarantine</code>

  - <code>deny_requests</code>

  - <code>ansible_provision</code>

  - <code>allow_requests</code>
</div>

<hr />

<div class="dd">

<code>BMCMACAddress</code>  <i>string</i>

</div>
<div class="dt">

MAC address of the server's BMC

</div>

<hr />

<div class="dd">

<code>SNMPCommunityPaswordDCencrypted</code>  <i>string</i>

</div>
<div class="dt">

SNMP Community password encrypted. Internal.

</div>

<hr />

<div class="dd">

<code>MGMTNMPCommunityPasswordDCEncrypted</code>  <i>string</i>

</div>
<div class="dt">

SNMP Community password encrypted

</div>

<hr />

<div class="dd">

<code>MGMTSNMPVersion</code>  <i>int</i>

</div>
<div class="dt">

SNMP Community version

</div>

<hr />

<div class="dd">

<code>MGMTSNMPPort</code>  <i>int</i>

</div>
<div class="dt">

SNMP Community port

</div>

<hr />

<div class="dd">

<code>secureBootIsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If the Secure Boot is enabled

</div>

<hr />

<div class="dd">

<code>subnetOOBIndex</code>  <i>int</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>subnetIPMIVersion</code>  <i>string</i>

</div>
<div class="dt">

The version of the IPMI protocol when IPMI is used.

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the server record was created. Readonly.

</div>

<hr />

<div class="dd">

<code>lastCleanupStart</code>  <i>string</i>

</div>
<div class="dt">

Last cleanup timestamp

</div>

<hr />

<div class="dd">

<code>vendorInfoJSON</code>  <i>string</i>

</div>
<div class="dt">

Version information as returned by the server.

</div>

<hr />

<div class="dd">

<code>chipsetName</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>requiresReregiter</code>  <i>bool</i>

</div>
<div class="dt">

Set if the server is considered 'dirty' and needs re-registration.

</div>

<hr />

<div class="dd">

<code>GPUCount</code>  <i>int</i>

</div>
<div class="dt">

Count of GPUs

</div>

<hr />

<div class="dd">

<code>GPUModel</code>  <i>string</i>

</div>
<div class="dt">

Model of GPUs if all are identical

</div>

<hr />

<div class="dd">

<code>GPUVendor</code>  <i>string</i>

</div>
<div class="dt">

Vendor of GPUs if all are identical

</div>

<hr />

<div class="dd">

<code>allocationTimestamp</code>  <i>string</i>

</div>
<div class="dt">

When the server was allocated to an infrastructure.

</div>

<hr />

<div class="dd">

<code>DHCPRelaySecurityIsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If enabled, the DHCP server will use verify the DHCP option82 for added security in interpreting the DHCP packet. Used for diagnosing issues with the legacy protocol.

</div>

<hr />

<div class="dd">

<code>requiresManualCleaning</code>  <i>bool</i>

</div>
<div class="dt">

If true, some steps of the cleanup process could not be performed automatically.

</div>

<hr />

<div class="dd">

<code>cleanupInProgress</code>  <i>bool</i>

</div>
<div class="dt">

If true, a cleanup operation is ongoing.

</div>

<hr />

<div class="dd">

<code>serverSupportsVirtualMedia</code>  <i>bool</i>

</div>
<div class="dt">

If true, server supports virtual media-based deployment

</div>

<hr />

<div class="dd">

<code>serverIsInDiagnostics</code>  <i>bool</i>

</div>
<div class="dt">

If true, server is ongoing a diagnosis. Deprecated.

</div>

<hr />

<div class="dd">

<code>diskWipe</code>  <i>bool</i>

</div>
<div class="dt">

If true, server will be wiped.

</div>

<hr />

<div class="dd">

<code>metricsMetadataJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal.

</div>

<hr />

<div class="dd">

<code>agentID</code>  <i>int</i>

</div>
<div class="dt">

Id of the site controller agent managing this server.

</div>

<hr />

<div class="dd">

<code>DHCPPacketSniffingIsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the DHCP server will record a log of the dialog with the server. Useful for diagnosing issues with the legacy protocol.

</div>

<hr />

<div class="dd">

<code>BDKDebug</code>  <i>bool</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>keysJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>NICDetails</code>  <i>map[string]<a href="#servernicdetails">ServerNICDetails</a></i>

</div>
<div class="dt">

Details about the network interfaces

</div>

<hr />

<div class="dd">

<code>submodel</code>  <i>string</i>

</div>
<div class="dt">

Model details as returned by the server

</div>

<hr />





## ServerDisk
ServerDisk describes a disk

Appears in:


- <code><a href="#server">Server</a>.disks</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>model</code>  <i>string</i>

</div>
<div class="dt">

The model of the object

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The type of the disk


Valid values:


  - <code>NVME</code>

  - <code>SSD</code>

  - <code>HDD</code>
</div>

<hr />

<div class="dd">

<code>vendor</code>  <i>string</i>

</div>
<div class="dt">

The vendor of the disk

</div>

<hr />

<div class="dd">

<code>status</code>  <i>string</i>

</div>
<div class="dt">

The status of the disk

</div>

<hr />

<div class="dd">

<code>serial_number</code>  <i>string</i>

</div>
<div class="dt">

The serial number of the disk

</div>

<hr />

<div class="dd">

<code>sizeGB</code>  <i>int</i>

</div>
<div class="dt">

The size of the disk

</div>

<hr />





## ServerInterface
ServerInterface contains server connectivity information.

Appears in:


- <code><a href="#server">Server</a>.interfaces</code>





<hr />

<div class="dd">

<code>macAddress</code>  <i>string</i>

</div>
<div class="dt">

MAC address of the interface

</div>

<hr />





## ServerNICDetails

Appears in:


- <code><a href="#server">Server</a>.NICDetails</code>





<hr />

<div class="dd">

<code>networkEquipmentInterfaceLLDPInformation</code>  <i>string</i>

</div>
<div class="dt">

LLDP information for this interface

</div>

<hr />

<div class="dd">

<code>networkEquipmentInterfaceMACAddress</code>  <i>string</i>

</div>
<div class="dt">

MAC Address of the switch interface

</div>

<hr />

<div class="dd">

<code>switchPortID</code>  <i>string</i>

</div>
<div class="dt">

Switch ID of the switch to which this interface is connected to

</div>

<hr />

<div class="dd">

<code>switchPortDescription</code>  <i>string</i>

</div>
<div class="dt">

Switch port name

</div>

<hr />

<div class="dd">

<code>switchHostname</code>  <i>string</i>

</div>
<div class="dt">

Hostname of the switch

</div>

<hr />

<div class="dd">

<code>networkEquipmentDescription</code>  <i>string</i>

</div>
<div class="dt">

Description of the switch

</div>

<hr />

<div class="dd">

<code>switchVLANID</code>  <i>string</i>

</div>
<div class="dt">

VLAN ID allocated as native to this interface

</div>

<hr />

<div class="dd">

<code>switchInterfaceIndex</code>  <i>int</i>

</div>
<div class="dt">

Index of this interface on the server. This is determined based on the switch hostname and port description.

</div>

<hr />

<div class="dd">

<code>serverInterfaceMACAddress</code>  <i>string</i>

</div>
<div class="dt">

MAC Address of the server interface

</div>

<hr />

<div class="dd">

<code>serverInterfaceCapacityMBPs</code>  <i>int</i>

</div>
<div class="dt">

MAX Speed of the interface

</div>

<hr />








## SharedDrive
SharedDrive represents a drive that can be shared between instances






<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the shared drive

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

Unique string representing the shared drive

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Id of the shared drive

</div>

<hr />

<div class="dd">

<code>storageType</code>  <i>string</i>

</div>
<div class="dt">

Type of the shared drive


Valid values:


  - <code>iscsi_ssd</code>

  - <code>iscsi_hdd</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

ID of the infrastructure

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

Service status of the shared drive


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the SharedDrive was created.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the SharedDrive was updated.

</div>

<hr />

<div class="dd">

<code>sizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

Size of the drive

</div>

<hr />

<div class="dd">

<code>attachedInstaceArrays</code>  <i>[]int</i>

</div>
<div class="dt">

An array of the instance array ids attached to this drive

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#shareddriveoperation">SharedDriveOperation</a></i>

</div>
<div class="dt">

The operation object

</div>

<hr />

<div class="dd">

<code>credentials</code>  <i><a href="#shareddrivecredentials">SharedDriveCredentials</a></i>

</div>
<div class="dt">

Credentials of the shared drive

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The operation id

</div>

<hr />

<div class="dd">

<code>targetsJSON</code>  <i>string</i>

</div>
<div class="dt">

Details on the ISCSI or FC targets

</div>

<hr />

<div class="dd">

<code>ioLimit</code>  <i>string</i>

</div>
<div class="dt">

Used by certain storage types

</div>

<hr />

<div class="dd">

<code>wwn</code>  <i>string</i>

</div>
<div class="dt">

WWN of the drive as reported by the storage appliance

</div>

<hr />

<div class="dd">

<code>storagePoolID</code>  <i>int</i>

</div>
<div class="dt">

The storage pool id to use if set.

</div>

<hr />

<div class="dd">

<code>affinity</code>  <i>string</i>

</div>
<div class="dt">

Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools


Valid values:


  - <code>same_storage</code>

  - <code>different_storage</code>
</div>

<hr />





## SharedDriveCredentials
SharedDriveCredentials iscsi or other forms of connection details

Appears in:


- <code><a href="#shareddrive">SharedDrive</a>.credentials</code>





<hr />

<div class="dd">

<code>iscsi</code>  <i>ISCSI</i>

</div>
<div class="dt">

details

</div>

<hr />





## SharedDriveOperation
SharedDriveOperation represents an ongoing or new operation on a shared drive

Appears in:


- <code><a href="#shareddrive">SharedDrive</a>.operation</code>





<hr />

<div class="dd">

<code>deployStatus</code>  <i>string</i>

</div>
<div class="dt">

Deploy status


Valid values:


  - <code>not_started</code>

  - <code>ongoing</code>

  - <code>finished</code>
</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

Type of operation


Valid values:


  - <code>create</code>

  - <code>edit</code>

  - <code>delete</code>

  - <code>start</code>

  - <code>stop</code>
</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

Unique string describing this shared drive

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

ID of the shared drive

</div>

<hr />

<div class="dd">

<code>sizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

Size of the drive

</div>

<hr />

<div class="dd">

<code>storageType</code>  <i>string</i>

</div>
<div class="dt">

Type of the shared drive. Readonly.


Valid values:


  - <code>iscsi_ssd</code>

  - <code>iscsi_hdd</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

ID of the infrastructure

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

Status of the service


Valid values:


  - <code>active</code>

  - <code>ordered</code>

  - <code>deleted</code>

  - <code>suspended</code>

  - <code>stopped</code>
</div>

<hr />

<div class="dd">

<code>attachedInstanceArrays</code>  <i>[]int</i>

</div>
<div class="dt">

List of instance arrays to which this shared drive is attached

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

ID of the operation

</div>

<hr />

<div class="dd">

<code>ioLimit</code>  <i>string</i>

</div>
<div class="dt">

Used for certain SAN appliances

</div>

<hr />








## StageDefinition
StageDefinition Also called a workflow task contains a JavaScript file, HTTP request url and options, an AnsibleBundle or an API call template.






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

Id of the owner

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

Id of the user who created the task. Internal.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the task

</div>

<hr />

<div class="dd">

<code>iconAssetDataURI</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>title</code>  <i>string</i>

</div>
<div class="dt">

Title (Name) of the task

</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

Description of the task

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

Type of task


Valid values:


  - <code>AnsibleBundle</code>

  - <code>JavaScript</code>

  - <code>APICall</code>

  - <code>HTTPRequest</code>

  - <code>Copy</code>

  - <code>WorkflowReference</code>
</div>

<hr />

<div class="dd">

<code>variableNames</code>  <i>[]string</i>

</div>
<div class="dt">

list of variable names required when executing this task.

</div>

<hr />

<div class="dd">

<code>stageDefinition</code>  <i>interface{}</i>

</div>
<div class="dt">

More details depending on the type of task


Valid values:


  - <code>HTTPRequest</code>

  - <code>AnsibleBundle</code>

  - <code>WorkflowReference</code>

  - <code>SSHExec</code>

  - <code>Copy</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the stage definition record was created. Readonly.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the stage definition record was updated. Readonly.

</div>

<hr />

<div class="dd">

<code>context</code>  <i>string</i>

</div>
<div class="dt">

Internal. Readonly.


Valid values:


  - <code>local</code>

  - <code>global</code>
</div>

<hr />

<div class="dd">

<code>automaticallyAddedToPreDeploys</code>  <i>int</i>

</div>
<div class="dt">

If set to 1 it will be added to all infrastructures at the pre-deploy stage

</div>

<hr />

<div class="dd">

<code>automaticallyAddedToPostDeploys</code>  <i>int</i>

</div>
<div class="dt">

If set to 1 it will be added to all infrastructures at the post-deploy stage

</div>

<hr />





## HTTPRequest
HTTPRequest represents an HTTP request definition compatible with the standard Web Fetch API.






<hr />

<div class="dd">

<code>url</code>  <i>string</i>

</div>
<div class="dt">

URL to call

</div>

<hr />

<div class="dd">

<code>options</code>  <i><a href="#webfetchaapioptions">WebFetchAAPIOptions</a></i>

</div>
<div class="dt">

Options

</div>

<hr />

<div class="dd">

<code>url</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />





## WebFetchAAPIOptions
WebFetchAAPIOptions represents node-fetch options which is follows the Web API Fetch specification. See https://github.com/node-fetch/node-fetch

Appears in:


- <code><a href="#httprequest">HTTPRequest</a>.options</code>





<hr />

<div class="dd">

<code>method</code>  <i>string</i>

</div>
<div class="dt">

HTTP Method to use

</div>

<hr />

<div class="dd">

<code>redirect</code>  <i>string</i>

</div>
<div class="dt">

Defaults to 'follow'. Set to `manual` to extract redirect headers, `error` to reject redirect


Valid values:


  - <code>follow</code>

  - <code>manual</code>

  - <code>error</code>
</div>

<hr />

<div class="dd">

<code>follow</code>  <i>int</i>

</div>
<div class="dt">

Maximum redirect count. 0 to not follow redirect. Defaults to 20

</div>

<hr />

<div class="dd">

<code>compress</code>  <i>bool</i>

</div>
<div class="dt">

If se to true it will support compression

</div>

<hr />

<div class="dd">

<code>timeout</code>  <i>int</i>

</div>
<div class="dt">

Timeout setting in seconds

</div>

<hr />

<div class="dd">

<code>size</code>  <i>int</i>

</div>
<div class="dt">

Maximum response body size in bytes. 0 to disable (default)

</div>

<hr />

<div class="dd">

<code>headers</code>  <i><a href="#webfetchapirequestheaders">WebFetchAPIRequestHeaders</a></i>

</div>
<div class="dt">

Request headers

</div>

<hr />

<div class="dd">

<code>body</code>  <i>string</i>

</div>
<div class="dt">

Body of the request

</div>

<hr />

<div class="dd">

<code>bodyBase64</code>  <i>string</i>

</div>
<div class="dt">

Body of the requested encoded base64

</div>

<hr />





## WebFetchAPIRequestHeaders
WebFetchAPIRequestHeaders HTTP request headers. null means undefined (the default for most) so the header will not be included with the request.

Appears in:


- <code><a href="#webfetchaapioptions">WebFetchAAPIOptions</a>.headers</code>





<hr />

<div class="dd">

<code>accept</code>  <i>string</i>

</div>
<div class="dt">

accept header

</div>

<hr />

<div class="dd">

<code>userAgent</code>  <i>string</i>

</div>
<div class="dt">

user-agent header

</div>

<hr />

<div class="dd">

<code>contentType</code>  <i>string</i>

</div>
<div class="dt">

content-type header

</div>

<hr />

<div class="dd">

<code>cookie</code>  <i>string</i>

</div>
<div class="dt">

cookie header

</div>

<hr />

<div class="dd">

<code>authorization</code>  <i>string</i>

</div>
<div class="dt">

authorization header

</div>

<hr />

<div class="dd">

<code>proxyAuthorization</code>  <i>string</i>

</div>
<div class="dt">

proxy-authorization header

</div>

<hr />

<div class="dd">

<code>contentMD5</code>  <i>string</i>

</div>
<div class="dt">

content-md5 header

</div>

<hr />





## AnsibleBundle
AnsibleBundle contains an Ansible project as a single archive file, usually .zip






<hr />

<div class="dd">

<code>filename</code>  <i>string</i>

</div>
<div class="dt">

file name

</div>

<hr />

<div class="dd">

<code>contentsBase64</code>  <i>string</i>

</div>
<div class="dt">

Content in base64

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

internal

</div>

<hr />





## WorkflowReference
WorkflowReference points to a Workflow object via its workflow_id. To be used as a stage definition.






<hr />

<div class="dd">

<code>workflowId</code>  <i>int</i>

</div>
<div class="dt">

id of the workflow to reference

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

internal

</div>

<hr />





## SSHExec
SSHExec executes a command on a remote server using the SSH exec functionality (not through a shell).






<hr />

<div class="dd">

<code>command</code>  <i>string</i>

</div>
<div class="dt">

Command to execute

</div>

<hr />

<div class="dd">

<code>target</code>  <i><a href="#sshclientoptions">SSHClientOptions</a></i>

</div>
<div class="dt">

Details on how to connect to the target system

</div>

<hr />

<div class="dd">

<code>timeout</code>  <i>int</i>

</div>
<div class="dt">

Timeout in seconds

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />





## SSHClientOptions
SSHClientOptions defines an ssh cnnection such as the host, port, user, password, private keys, etc. All properties support template-like variables; for example, ${{instance_credentials_password}} may be used as value for the password property.

Appears in:


- <code><a href="#sshexec">SSHExec</a>.target</code>

- <code><a href="#scpresourcelocation">SCPResourceLocation</a>.target</code>





<hr />

<div class="dd">

<code>host</code>  <i>string</i>

</div>
<div class="dt">

Host

</div>

<hr />

<div class="dd">

<code>port</code>  <i>int</i>

</div>
<div class="dt">

Port

</div>

<hr />

<div class="dd">

<code>forceIPv4</code>  <i>bool</i>

</div>
<div class="dt">

Force the use of ipv4

</div>

<hr />

<div class="dd">

<code>forceIPv6</code>  <i>bool</i>

</div>
<div class="dt">

Force the use of ipv6

</div>

<hr />

<div class="dd">

<code>hostHash</code>  <i>string</i>

</div>
<div class="dt">

Hash of the host

</div>

<hr />

<div class="dd">

<code>hashedKey</code>  <i>string</i>

</div>
<div class="dt">

Hash key of the host

</div>

<hr />

<div class="dd">

<code>username</code>  <i>string</i>

</div>
<div class="dt">

Username

</div>

<hr />

<div class="dd">

<code>password</code>  <i>string</i>

</div>
<div class="dt">

Password

</div>

<hr />

<div class="dd">

<code>privateKey</code>  <i>string</i>

</div>
<div class="dt">

Private key to use

</div>

<hr />

<div class="dd">

<code>passphrase</code>  <i>string</i>

</div>
<div class="dt">

Private key passphrase to use

</div>

<hr />

<div class="dd">

<code>readyTimeout</code>  <i>int</i>

</div>
<div class="dt">

Timeout in seconds

</div>

<hr />

<div class="dd">

<code>strictVendor</code>  <i>bool</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>algorithms</code>  <i><a href="#sshalgorithms">SSHAlgorithms</a></i>

</div>
<div class="dt">

SSH Algorithms to use

</div>

<hr />

<div class="dd">

<code>compress</code>  <i>interface{}</i>

</div>
<div class="dt">

Internal

</div>

<hr />





## SSHAlgorithms
SSHAlgorithms defines algorithms that can be used during an ssh session

Appears in:


- <code><a href="#sshclientoptions">SSHClientOptions</a>.algorithms</code>





<hr />

<div class="dd">

<code>kex</code>  <i>[]string</i>

</div>
<div class="dt">

Kex

</div>

<hr />

<div class="dd">

<code>cipher</code>  <i>[]string</i>

</div>
<div class="dt">

Ciphers accepted

</div>

<hr />

<div class="dd">

<code>serverHostKey</code>  <i>[]string</i>

</div>
<div class="dt">

Host keys accepted

</div>

<hr />

<div class="dd">

<code>hmac</code>  <i>[]string</i>

</div>
<div class="dt">

HMAC accepted

</div>

<hr />

<div class="dd">

<code>compress</code>  <i>[]string</i>

</div>
<div class="dt">

Compress

</div>

<hr />





## Copy
Copy defines the source and destination of a SCP operation. The source may be of various types. SCP and HTTP requests are streamed so they are recommended as sources. The destination has to be a SCP resource.






<hr />

<div class="dd">

<code>source</code>  <i>interface{}</i>

</div>
<div class="dt">

the source of the file


Valid values:


  - <code>SCPResourceLocation</code>
</div>

<hr />

<div class="dd">

<code>destination</code>  <i><a href="#scpresourcelocation">SCPResourceLocation</a></i>

</div>
<div class="dt">

the destination

</div>

<hr />

<div class="dd">

<code>timeoutMinutes</code>  <i>int</i>

</div>
<div class="dt">

timeout in minutes

</div>

<hr />

<div class="dd">

<code>ifDestinationAlreadyExists</code>  <i>string</i>

</div>
<div class="dt">

What to do if file exists at the destination. Defaults to 'overwrite'


Valid values:


  - <code>error</code>

  - <code>overwrite</code>

  - <code>errorIfNotSameSize</code>

  - <code>overwriteIfNotSameSize</code>

  - <code>ignore</code>
</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

internal

</div>

<hr />





## SCPResourceLocation
SCPResourceLocation defines a file path and SSH client connection options for use with Secure Copy Protocol (SCP).

Appears in:


- <code><a href="#copy">Copy</a>.destination</code>





<hr />

<div class="dd">

<code>path</code>  <i>string</i>

</div>
<div class="dt">

path

</div>

<hr />

<div class="dd">

<code>target</code>  <i><a href="#sshclientoptions">SSHClientOptions</a></i>

</div>
<div class="dt">

SSH Options for the target

</div>

<hr />








## Workflow
Workflow struct defines a server type






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

Id of the owner

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

Internal.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label

</div>

<hr />

<div class="dd">

<code>usage</code>  <i>string</i>

</div>
<div class="dt">

Usage of this workflow


Valid values:


  - <code>infrastructure</code>

  - <code>switch_device</code>

  - <code>server</code>

  - <code>free_standing</code>

  - <code>storage_pool</code>

  - <code>user</code>

  - <code>os_template</code>

  - <code>global</code>
</div>

<hr />

<div class="dd">

<code>title</code>  <i>string</i>

</div>
<div class="dt">

Title (name) of this workflow

</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

Description of this workflow

</div>

<hr />

<div class="dd">

<code>isDeprecated</code>  <i>bool</i>

</div>
<div class="dt">

Set to true if this workflow is deprecated

</div>

<hr />

<div class="dd">

<code>assetDataURI</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the workflow record was created. Readonly.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the workflow record was updated. Readonly.

</div>

<hr />





## WorkflowStageDefinitionReference
WorkflowStageDefinitionReference defines where in a workflow a stage definition resides






<hr />

<div class="dd">

<code>workflowStageID</code>  <i>int</i>

</div>
<div class="dt">

Workflow stage id

</div>

<hr />

<div class="dd">

<code>workflowID</code>  <i>int</i>

</div>
<div class="dt">

Workflow id

</div>

<hr />

<div class="dd">

<code>stageDefinitionID</code>  <i>int</i>

</div>
<div class="dt">

Stage definition id

</div>

<hr />

<div class="dd">

<code>runLevel</code>  <i>int</i>

</div>
<div class="dt">

Run level in the workflow

</div>

<hr />





## WorkflowStageAssociation
WorkflowStageAssociation associations






<hr />

<div class="dd">

<code>infrastructureDeployStageID</code>  <i>int</i>

</div>
<div class="dt">

Infrastructure stage id

</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

Infrastructure id

</div>

<hr />

<div class="dd">

<code>definitionID</code>  <i>int</i>

</div>
<div class="dt">

Stage definition id

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

description: Type of position in the infrastructure:


Valid values:


  - <code>pre_deploy</code>

  - <code>post_deploy</code>
</div>

<hr />

<div class="dd">

<code>runLevel</code>  <i>int</i>

</div>
<div class="dt">

Run level (the depth in the tree where this stage resides).

</div>

<hr />

<div class="dd">

<code>lastOutput</code>  <i>string</i>

</div>
<div class="dt">

The output of the last execution of this task

</div>

<hr />








## SubnetPool
SubnetPool represents a pool of subnets






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Id of the subnetpool

</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

Label fo the Datacenter

</div>

<hr />

<div class="dd">

<code>networkEquipmentID</code>  <i>int</i>

</div>
<div class="dt">

ID fo the network equipment to which this subnet pool is associated

</div>

<hr />

<div class="dd">

<code>user</code>  <i>int</i>

</div>
<div class="dt">

Owner of this subent pool

</div>

<hr />

<div class="dd">

<code>prefix</code>  <i>string</i>

</div>
<div class="dt">

Prefix

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of this subnet pool

</div>

<hr />

<div class="dd">

<code>prefixHex</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>netmask</code>  <i>string</i>

</div>
<div class="dt">

Netmask

</div>

<hr />

<div class="dd">

<code>netmaskHex</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>size</code>  <i>int</i>

</div>
<div class="dt">

Size of the prefix

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

Type


Valid values:


  - <code>ipv4</code>

  - <code>ipv6</code>
</div>

<hr />

<div class="dd">

<code>routable</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this subnet pool will be used for subnets that are routed (usually to the internet)

</div>

<hr />

<div class="dd">

<code>destination</code>  <i>string</i>

</div>
<div class="dt">

description: What this subnet is intended for:


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>

  - <code>oob</code>

  - <code>quarantine</code>

  - <code>loopback</code>

  - <code>vtep</code>

  - <code>disabled</code>
</div>

<hr />

<div class="dd">

<code>currentUtilizationJSON</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>currentUtilizationLastUpdated</code>  <i>string</i>

</div>
<div class="dt">

Internal

</div>

<hr />

<div class="dd">

<code>manualAllocationOnly</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this subnet will not be used for automatic allocation and will only be used in network profiles.

</div>

<hr />





## SubnetPoolUtilization
SubnetPoolUtilization describes the current utilization of the subnet






<hr />

<div class="dd">

<code>availableSubnets</code>  <i>map[string]int</i>

</div>
<div class="dt">

Count of available addresses

</div>

<hr />

<div class="dd">

<code>allocatedSubnets</code>  <i>map[string]int</i>

</div>
<div class="dt">

Count of allocated addresses

</div>

<hr />

<div class="dd">

<code>availableUsableIps</code>  <i>string</i>

</div>
<div class="dt">

Amount of usable addresses (excludes broadcast)

</div>

<hr />

<div class="dd">

<code>allocatedUsableIps</code>  <i>string</i>

</div>
<div class="dt">

Amount of allocated from the usable set

</div>

<hr />

<div class="dd">

<code>availablePercentage</code>  <i>string</i>

</div>
<div class="dt">

Available usable (Percentage)

</div>

<hr />








## SwitchDevice
SwitchDevice Represents a switch installed in a datacenter.






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

id of the network equipment

</div>

<hr />

<div class="dd">

<code>identifierString</code>  <i>string</i>

</div>
<div class="dt">

Hostname (or unique label) of the network equipment

</div>

<hr />

<div class="dd">

<code>datacenterName</code>  <i>string</i>

</div>
<div class="dt">

Datacenter

</div>

<hr />

<div class="dd">

<code>provisionerType</code>  <i>string</i>

</div>
<div class="dt">

Type of provisioner. Read only


Valid values:


  - <code>vpls</code>

  - <code>vxlan</code>

  - <code>vlan</code>

  - <code>lan</code>

  - <code>sdn</code>

  - <code>evpnvxlanl2</code>
</div>

<hr />

<div class="dd">

<code>provisionerPosition</code>  <i>string</i>

</div>
<div class="dt">

Role of this equipment in the provisioner


Valid values:


  - <code>tor</code>

  - <code>north</code>

  - <code>leaf</code>

  - <code>spine</code>

  - <code>other</code>
</div>

<hr />

<div class="dd">

<code>driver</code>  <i>string</i>

</div>
<div class="dt">

Driver to use. Note that this list may change frequently.


Valid values:


  - <code>hp5800</code>

  - <code>hp5900</code>

  - <code>hp5950</code>

  - <code>dell_s4000</code>

  - <code>dell_s4048</code>

  - <code>dell_s6010</code>

  - <code>os_10</code>

  - <code>cumulus42</code>

  - <code>nexus9000</code>

  - <code>cisco_aci51</code>

  - <code>junos</code>

  - <code>junos18</code>

  - <code>sonic_enterprise</code>

  - <code>dummy</code>
</div>

<hr />

<div class="dd">

<code>managementUsername</code>  <i>string</i>

</div>
<div class="dt">

Username

</div>

<hr />

<div class="dd">

<code>managementPassword</code>  <i>string</i>

</div>
<div class="dt">

Password

</div>

<hr />

<div class="dd">

<code>managementAddress</code>  <i>string</i>

</div>
<div class="dt">

Address

</div>

<hr />

<div class="dd">

<code>managementPort</code>  <i>int</i>

</div>
<div class="dt">

Port

</div>

<hr />

<div class="dd">

<code>managementProtocol</code>  <i>string</i>

</div>
<div class="dt">

Deprecated.

</div>

<hr />

<div class="dd">

<code>managementAddressMask</code>  <i>string</i>

</div>
<div class="dt">

Netmask of the management address

</div>

<hr />

<div class="dd">

<code>managementAddressGateway</code>  <i>string</i>

</div>
<div class="dt">

Gateway of the management address

</div>

<hr />

<div class="dd">

<code>managementMACAddress</code>  <i>string</i>

</div>
<div class="dt">

MAC address of the management address

</div>

<hr />

<div class="dd">

<code>primaryWANIPv4SubnetPool</code>  <i>string</i>

</div>
<div class="dt">

When set it will automatically create an IPv4 subnet pool for WAN. Deprecated

</div>

<hr />

<div class="dd">

<code>primaryWANIPv4SubnetPrefixSize</code>  <i>int</i>

</div>
<div class="dt">

Size of the subnet pool to automatically create. Deprecated

</div>

<hr />

<div class="dd">

<code>primaryWANIPv6SubnetPool</code>  <i>string</i>

</div>
<div class="dt">

Label of an IPv6 subnet to use when creating the switch.

</div>

<hr />

<div class="dd">

<code>primaryWANIPv6SubnetPoolID</code>  <i>int</i>

</div>
<div class="dt">

ID of an IPv6 subnet to use when creating the switch.

</div>

<hr />

<div class="dd">

<code>primaryWANIPv6SubnetCIDR</code>  <i>string</i>

</div>
<div class="dt">

CIDR of the subnet to create a subnet pool with

</div>

<hr />

<div class="dd">

<code>primaryWANIPv6SubnetPrefixSize</code>  <i>int</i>

</div>
<div class="dt">

Size of the subnet

</div>

<hr />

<div class="dd">

<code>primarySANSubnetPool</code>  <i>string</i>

</div>
<div class="dt">

Label of the SAN subnet to use

</div>

<hr />

<div class="dd">

<code>primarySANSubnetPrefixSize</code>  <i>int</i>

</div>
<div class="dt">

Size of the san subnet to use

</div>

<hr />

<div class="dd">

<code>quarantineSubnetStart</code>  <i>string</i>

</div>
<div class="dt">

Only used for legacy operation. Start of the quarantine subnet to use during provisioning.

</div>

<hr />

<div class="dd">

<code>quarantineSubnetEnd</code>  <i>string</i>

</div>
<div class="dt">

Only used for legacy operation. End of the quarantine subnet to use during provisioning.

</div>

<hr />

<div class="dd">

<code>quarantineSubnetPrefixSize</code>  <i>int</i>

</div>
<div class="dt">

Only used for legacy operation. End of the quarantine subnet to use during provisioning.

</div>

<hr />

<div class="dd">

<code>quarantineSubnetGateway</code>  <i>string</i>

</div>
<div class="dt">

Only used for legacy operation. Gateway to use during provisioning.

</div>

<hr />

<div class="dd">

<code>description</code>  <i>string</i>

</div>
<div class="dt">

Description of the network equipment

</div>

<hr />

<div class="dd">

<code>country</code>  <i>string</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>city</code>  <i>string</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>netDatacenter</code>  <i>string</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>datacenterRoom</code>  <i>string</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>datacenterRack</code>  <i>string</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>rackPositionUpperUnit</code>  <i>int</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>rackPositionLowerUnit</code>  <i>int</i>

</div>
<div class="dt">

Location of the network equipment

</div>

<hr />

<div class="dd">

<code>serialNumber</code>  <i>string</i>

</div>
<div class="dt">

Serial number of the network equipment

</div>

<hr />

<div class="dd">

<code>chassisRackID</code>  <i>int</i>

</div>
<div class="dt">

If applicable the chassis id

</div>

<hr />

<div class="dd">

<code>TORLinkedID</code>  <i>int</i>

</div>
<div class="dt">

If applicable, the id of the linked equipment

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

Tags

</div>

<hr />

<div class="dd">

<code>requiresOSInstall</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the OS will be reinstalled upon reboot

</div>

<hr />

<div class="dd">

<code>isBorderDevice</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this network equipment can be used for external connections

</div>

<hr />

<div class="dd">

<code>isStorageSwitch</code>  <i>bool</i>

</div>
<div class="dt">

If set to true this network equipment can be used for SAN fabrics

</div>

<hr />

<div class="dd">

<code>isGateway</code>  <i>bool</i>

</div>
<div class="dt">

In certain situations (such as for the VPLS provisioner) a switch can be used as a router

</div>

<hr />

<div class="dd">

<code>networkTypesAllowed</code>  <i>[]string</i>

</div>
<div class="dt">

Types of networks allowed on this switch


Valid values:


  - <code>wan</code>

  - <code>lan</code>

  - <code>san</code>
</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

Template id of the NOS to be installed

</div>

<hr />

<div class="dd">

<code>LoopbackAddress</code>  <i>string</i>

</div>
<div class="dt">

Address of a loopback interface (if applicable)

</div>

<hr />

<div class="dd">

<code>VTEPAddress</code>  <i>string</i>

</div>
<div class="dt">

Address of a VTEP

</div>

<hr />

<div class="dd">

<code>ASN</code>  <i>int</i>

</div>
<div class="dt">

Address of an ASN

</div>

<hr />

<div class="dd">

<code>controllerID</code>  <i>int</i>

</div>
<div class="dt">

In the SDN provisioner the ID of the controller device

</div>

<hr />





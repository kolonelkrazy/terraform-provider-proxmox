---
layout: page
title: proxmox_virtual_environment_vm2
parent: Data Sources
subcategory: Virtual Environment
description: |-
  This is an experimental implementation of a Proxmox VM datasource using Plugin Framework.
---

# Data Source: proxmox_virtual_environment_vm2

!> **DO NOT USE**
This is an experimental implementation of a Proxmox VM datasource using Plugin Framework.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (Number) The unique identifier of the VM in the Proxmox cluster.
- `node_name` (String) The name of the node where the VM is provisioned.

### Optional

- `clone` (Attributes) The cloning configuration. (see [below for nested schema](#nestedatt--clone))
- `cpu` (Attributes) The CPU configuration. (see [below for nested schema](#nestedatt--cpu))
- `description` (String) The description of the VM.
- `name` (String) The name of the VM.
- `rng` (Attributes) The RNG (Random Number Generator) configuration. (see [below for nested schema](#nestedatt--rng))
- `tags` (Set of String) The tags assigned to the VM.
- `template` (Boolean) Whether the VM is a template.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vga` (Attributes) The VGA configuration. (see [below for nested schema](#nestedatt--vga))

<a id="nestedatt--clone"></a>
### Nested Schema for `clone`

Required:

- `id` (Number) The ID of the VM to clone.

Optional:

- `retries` (Number) The number of retries to perform when cloning the VM (default: 3).


<a id="nestedatt--cpu"></a>
### Nested Schema for `cpu`

Optional:

- `affinity` (String) List of host cores used to execute guest processes, for example: '0,5,8-11'
- `architecture` (String) The CPU architecture.
- `cores` (Number) The number of CPU cores per socket.
- `flags` (Set of String) Set of additional CPU flags.
- `hotplugged` (Number) The number of hotplugged vCPUs.
- `limit` (Number) Limit of CPU usage.
- `numa` (Boolean) Enable NUMA.
- `sockets` (Number) The number of CPU sockets.
- `type` (String) Emulated CPU type.
- `units` (Number) CPU weight for a VM


<a id="nestedatt--rng"></a>
### Nested Schema for `rng`

Optional:

- `max_bytes` (Number) Maximum bytes of entropy allowed to get injected into the guest every period.
- `period` (Number) Period in milliseconds to limit entropy injection to the guest.
- `source` (String) The entropy source for the RNG device.


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.


<a id="nestedatt--vga"></a>
### Nested Schema for `vga`

Optional:

- `clipboard` (String) Enable a specific clipboard.
- `memory` (Number) The VGA memory in megabytes (4-512 MB). Has no effect with serial display.
- `type` (String) The VGA type.

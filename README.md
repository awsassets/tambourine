# Tambourine

Kubelet replacement with Built in Linux extensions

# Development 

## Success: 

 - Install, Manage, and Observe a new `systemd` service from Kubernetes.
 - Can run in any namespace (however, the controllers must respect your choice)

## Stories

#### Install: 

`kubectl apply -f tambourinze.yaml`

```yaml 
spec:
  name: alice-zfs
  state: installed
```

`kubectl apply -f tambourine.yaml`

#### Manage:

```yaml
spec:
  command: start/stop/restart/reload
```

#### Observe:

`kubectl get tambourine`

```yaml 

status:
    systemdName: NetworkManager.service - Network Manager
    loaded: loaded (/usr/lib/systemd/system/NetworkManager.service; enabled; vendor preset>
    drop-in: /usr/lib/systemd/system/NetworkManager.service.d
    active: active (running) since Fri 2021-03-05 13:38:11 PST; 1 day 21h ago
    docs: man:NetworkManager(8)
    pid: 514 (NetworkManager)
    tasks: 3 (limit: 77070)
    memory: 17.5M
    cgroup: /system.slice/NetworkManager.service
```

# Componets of Tambourine 

### Examples of Tambourine CRs
 
 - CNI Toolchains
 - CRI
 - Storage Drivers
    - ZFS
 - Seccomp Configuration
 - AppArmor Configuration
 - Falco Driver
 - SELinux Configuration
 - IPTables Configuration
 - eBPF programs

### Kubernetes Components

```yaml 
// These are the convenience components that will live
// in Kubernetes. These will only matter if/when the 
// Tambourine service on the System is configured to
// respect these configurations.
```

 - Tambourine CRD
 - Tambourine Controller
    - One way Sync Tambourine CRs -> Bombshell socket
    
### Bombshell 

Taken from `io.proto`

```yaml
// The basic service that can be used for basic IO with tambourine.
//
// The bombshell service can be authenticated with TLS. However, by design
// bombshell will never perform authorization. The service, for lack of a
// better term, will be unintelligent and will accept whatever IO is sent
// to it.
``` 

 - gRPC (TLS) service with a Unix Domain Socket
    - basic IO to/from a persistent space (silo) on the host
 - read `/proc/tambourine` data and service `/proc/tambourine`

### Tambourine

```yaml 
Linux service that is managed with `systemd` and is 
required to be installed on the host alongside the kubelet
    NOTE: Should this be a kubelet feature?
    Note: Should we have a kubelet and tambourine management operator?
```

 - Read from Silo
 - Write Linux status to /proc (by design tambourine will never write to Silo)
 
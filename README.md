# Usage [![Build Status](https://travis-ci.org/vmware/docker-machine-driver-vmwareappcatalyst.svg?branch=master)](https://travis-ci.org/vmware/docker-machine-driver-vmwareappcatalyst)

To build and install, make sure you have a Go 1.5 environment and Docker Machine
greater than or equal to 0.5.0.  Then, just run:

```
$ make && make install
```

Note that the AppCatalyst driver uses [Photon OS](https://github.com/vmware/photon) as underlying OS, this requires a specific `libmachine` provisioner that is currently being reviewed as PR docker/machine#2344, make sure you have that support compiled in before testing the plugin.

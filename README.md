# ![Intel Device Plugins for Kubernetes](.intel-logo.png) Device Plugins for Kubernetes
[![Build Status](https://travis-ci.org/intel/intel-device-plugins-for-kubernetes.svg?branch=master)](https://travis-ci.org/intel/intel-device-plugins-for-kubernetes)
[![Go Report Card](https://goreportcard.com/badge/github.com/intel/intel-device-plugins-for-kubernetes)](https://goreportcard.com/report/github.com/intel/intel-device-plugins-for-kubernetes)
[![GoDoc](https://godoc.org/github.com/intel/intel-device-plugins-for-kubernetes/pkg/deviceplugin?status.svg)](https://godoc.org/github.com/intel/intel-device-plugins-for-kubernetes/pkg/deviceplugin)

## Table of Contents

* [About](#about)
* [Prerequisites](#prerequisites)
* [Plugins](#plugins)
    * [GPU device plugin](#gpu-device-plugin)
    * [FPGA device plugin](#fpga-device-plugin)
        * [device plugin](#device-plugin)
        * [admission controller](#admission-controller)
        * [CRI-O prestart hook](#cri-o-prestart-hook)
    * [QAT device plugin](#qat-device-plugin)
    * [VPU device plugin](#vpu-device-plugin)
* [Demos](#demos)
* [Developers](#developers)
* [Supported Kubernetes versions](#supported-kubernetes-versions)
* [Related code](#related-code)

## About

This repository contains a framework for developing plugins for the Kubernetes
[device plugins framework](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/),
along with a number of device plugin implementations utilising that framework.

## Prerequisites

Prerequisites for building and running these device plugins include:

- Appropriate hardware
- A fully configured [Kubernetes cluster]
- A working [Go environment], of at least version v1.13.

## Plugins

The below sections detail existing plugins developed using the framework.

### GPU device plugin

The [GPU device plugin](cmd/gpu_plugin/README.md) supports Intel
[GVT-d](https://github.com/intel/gvt-linux/wiki/GVTd_Setup_Guide) device passthrough
and acceleration using GPUs of the following hardware families:

- Integrated GPUs within Intel Core processors
- Intel Xeon processors
- Intel Visual Compute Accelerator (Intel VCA)

The demo subdirectory contains both a [GPU plugin demo video](demo#intel-gpu-device-plugin-demo-video)
as well as code for an OpenCL [FFT demo](demo/ubuntu-demo-opencl).

### FPGA device plugin

The [FPGA device plugin](cmd/fpga_plugin/README.md) supports FPGA passthrough for
the following hardware:

- Intel Arria 10
- Intel Stratix 10

The FPGA plugin comes as three parts.

- the [device plugin](#device-plugin)
- the [admissing controller](#admission-controller)
- the [CRIO-O prestart hook](#cri-o-prestart-hook)

Refer to each individual sub-components documentation for more details. Brief overviews
of the sub-components are below.

The demo subdirectory contains a [video](demo#intel-fpga-device-plugin-demo-video) showing deployment
and use of the FPGA plugin. Sources relating to the demo can be found in the
[opae-nlb-demo](demo/opae-nlb-demo) subdirectory.

#### device plugin

The [FPGA device plugin](cmd/fpga_plugin/README.md) is responsible for discovering and reporting FPGA
devices to `kubelet`.

#### admission controller

The [FPGA admission controller webhook](cmd/fpga_admissionwebhook/README.md) is responsible for performing
mapping from user-friendly function IDs to the Interface ID and Bitstream ID that are required for FPGA
programming.  It also implements access control by namespacing FPGA configuration information.

#### CRI-O prestart hook

The [FPGA prestart CRI-O hook](cmd/fpga_crihook/README.md) performs discovery of the requested FPGA
function bitstream and programs FPGA devices based on the environment variables in the workload
description.

### [QAT](https://01.org/intel-quick-assist-technology) device plugin

The [QAT plugin](cmd/qat_plugin/README.md) supports device plugin for Intel QAT adapters, and includes
code [showing deployment](cmd/qat_plugin/dpdkdrv) via [DPDK](https://doc.dpdk.org/guides/cryptodevs/qat.html).

The demo subdirectory includes details of both a
[QAT DPDK demo](demo#intel-quickassist-technology-device-plugin-with-dpdk-demo-video)
and a [QAT OpenSSL demo](demo#intel-quickassist-technology-device-plugin-openssl-demo-video).
Source for the OpenSSL demo can be found in the [relevant subdirectory](demo/openssl-qat-engine).

Details for integrating the QAT device plugin into [Kata Containers](https://katacontainers.io/)
can be found in the
[Kata Containers documentation repository](https://github.com/kata-containers/documentation/blob/master/use-cases/using-Intel-QAT-and-kata.md).

### VPU device plugin

The [VPU device plugin](cmd/vpu_plugin/README.md) supports Intel VCAC-A card
(https://www.intel.com/content/dam/www/public/us/en/documents/datasheets/media-analytics-vcac-a-accelerator-card-by-celestica-datasheet.pdf)
the card has:
- 1 Intel Core i3-7100U processor
- 12 MyriadX VPUs
- 8GB DDR4 memory

The demo subdirectory includes details of a OpenVINO deployment and use of the VPU plugin.
Sources can be found in [openvino-demo](demo/ubuntu-demo-openvino)

## Demos

The [demo subdirectory](demo/readme.md) contains a number of demonstrations for a variety of the
available plugins.

## Developers

For information on how to develop a new plugin using the framework, see the
[Developers Guide](DEVEL.md) and the code in the
[device plugins pkg directory](pkg/deviceplugin).

## Supported Kubernetes versions

Releases are made under the github [releases area](releases). Supported releases and
matching Kubernetes versions are listed below:

| Branch            | Kubernetes branch/version      |
|:------------------|:-------------------------------|
| release-0.15      | Kubernetes 1.15 branch v1.15.x |
| release-0.11      | Kubernetes 1.11 branch v1.11.x |

[Go environment]: https://golang.org/doc/install
[Kubernetes cluster]: https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/

## Related code

A related Intel SRIOV network device plugin can be found in [this repository](https://github.com/intel/sriov-network-device-plugin)

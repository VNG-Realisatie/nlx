---
id: setup-your-environment
title: Setup your environment
---


## Mac and Linux users

### Docker

You can download Docker for Linux and Mac OS [here](https://hub.docker.com?overlay=onboarding). Please note that you will have to create a free DockerHub account before you can download Docker.

## Windows users

We currently support Windows 10 64-bit: Pro, Enterprise, or Education (Build 15063 or later).

### Powershell

In order to successfully execute all commands in this guide make sure you are using `PowerShell` in administrator mode and not the `Command prompt`. To launch `PowerShell` in administrator mode, type `PowerShell` in the searchfield of the taskbar, you should find `Windows PowerShell`, right click on it and select `Run as Administrator`.

### Docker

You can download Docker for Windows [here](https://hub.docker.com?overlay=onboarding). Please note that you will have to create a free DockerHub account before you can download Docker. Docker requires hyper-V to be enabled and Docker will ask you to enable it on start-up if this does not happen you can enable it yourself by running following command in PowerShell.

```bash
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All
```

## Docker Compose

For setting up the testing environment locally we'll use Docker Compose.
You can install it by following the [installation steps of the official Docker website](https://docs.docker.com/compose/install/#install-compose).

### OpenSSL

You will need `OpenSSL` to create the TLS certificates you will need to run NLX. We advise you to install [Chocolatey](https://chocolatey.org/install), a package manager for Windows which will install `OpenSSL` for you. 

Once you installed Chocolatey you can install OpenSSL by running.

```bash
choco install openssl.light
```

Now close and reopen PowerShell and verify your OpenSSL installation by running

```bash
openssl version
```

if the installation was successful, OpenSSL will print its version number.


## Working directory

You can use any directory you want. Just make sure to update the example commands of all future commands accordingly.

First, let's navigate to our home directory.

```bash
cd ~
```

Now, let's create the `nlx-setup` directory.

```bash
mkdir nlx-setup
```
```bash
cd nlx-setup
```

> You can verify your current location by running

```bash
pwd
```

The output should be:
* For Mac & Windows: `/Users/<your-username>/nlx-setup`
* For Linux: `/home/<your-username>/nlx-setup`

All commands later in this guide assume you are located in this directory.


## In sum

Having prepared your environment will allow you to quickly move on during the next 
chapters. Now let's [retrieve a demo certificate](./retrieve-a-demo-certificate.md) 
so you are able to send traffic through the NLX network.

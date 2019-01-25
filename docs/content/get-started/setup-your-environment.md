---
title: "Part 1: Setup your environment"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Operating System

This documentation is written for Mac and Linux users. It assumes you have some experience with the terminal / shell. 

Windows 10 is not officially supported. However, users may be able to follow along using 
[ubuntu for windows](https://tutorials.ubuntu.com/tutorial/tutorial-ubuntu-on-windows).

## Docker

Make sure you have installed a recent version of [Docker](https://www.docker.com) installed.

## Working directory

All steps in the guide assume you're located in the `~/Users/<your-username>/nlx-setup` directory.
You can use any directory you want. Just make sure to update the example commands of all future commands accordingly

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

* For Mac: `/Users/<your-username>/nlx-setup`
* For Linux: `/home/<your-username>/nlx-setup`

All commands further down this guide assume you are located in this directory.

## Demo CA root certificate

You will need our demo CA's root certificate. It will be used to validate certificates of other organizations.
 
Download [the root certificate](https://certportal.demo.nlx.io/root.crt) file and save it as `root.crt` in the working directory described above.

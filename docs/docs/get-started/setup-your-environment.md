---
title: "Part 1: Setup your environment"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Operating System

This documentation is written for Mac and Linux users. It assumes you have some experience with the terminal / shell.

## Docker

Make sure you have a recent version of [Docker](https://www.docker.com) installed.

## Working directory

All steps in the guide assume you're located in the `~/nlx-setup` directory.
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

* For Mac: `/Users/<your-username>/nlx-setup`
* For Linux: `/home/<your-username>/nlx-setup`

All commands further down this guide assume you are located in this directory.

The Next step is to create some certificates [part 2](../create-certificates/).

---
id: introduction
title: Introduction
---

# Introduction

Note: The result of this guide is an NLX environment that connects with the NLX demo environment. This environment is not suitable for production purposes.

In this guide, we work towards offering an API via an [NLX inway](https://docs.nlx.io/understanding-the-basics/introduction#high-level-technical-overview). We request (and grant) access to that API and query that API with a client via an NLX outway. All NLX components are installed on a Kubernetes [(Haven)](https://haven.commonground.nl) cluster.


## Step by step guide

0. Preparation
1. Create namespace
2. Create certificates
3. Install PostgreSQL
4. Install NLX Management
5. Install NLX Inway
6. Install NLX Outway
7. Install and offer sample API
8. Access API through a client
---
id: enable-pricing
title: Enable tranparant pricing (experimental)
---

## Before you start

- Make sure you read our [try NLX](../try-nlx/introduction) guide
- Make sure that you have NLX management running
- Make sure that you have an inway and you are offering a service

## Introduction

This guide will explain how NLX can be used to implement a transparant pricing model for your services. You can create a subscription model, a pay per use model or a hybrid model. This feature is introduced, as a working example to aid the study, brainstorm, discussion and thinking process about cost of ownership of (public) data vs cost of use of (public) data. 

### Legally binding contract vs transparancy only

This feature is not intended as part of an automatic system where legally binding contracts are easily made. You, as a service owner, could use it to inform potential users about the costs but you would still need a separate contact in place to start billing. The time to do that, is before granting access to your service. 

## Setup the transaction-log database

Before you can use the finance page make sure to [enable transaction logs](./enable-transaction-logs).

Currently, we only support one database for all transaction logs.

## Enable the transaction-log

The transaction log is currently only used to generate the finance report.

To enable the transaction-log, start NLX Management with `--transaction-log-dsn` flag, or set the `TRANSACTION_LOG_DSN` environment variable.  

## Setup a price for a service

- Start NLX Management and navigate to your services
- Click on the service you want to add pricing for and click "Edit"
- Scroll down to "Pricing" (kosten), tick "This is a paid service" (Dit is een betaalde service) and add the desired pricing information
- After submitting the form you'll see the pricing information in the service detail pane

## Agree to the cost when requesting access to a service

- When you want to use a service and request access, a confirmation dialog is shown containing the revenue model entered by the service owner
- Confirmation and thereby accepting these costs does not represent a legally binding contact.

## Use NLX Management to create an export

- Use the main navigation to go to "Finances"
- Click the export button

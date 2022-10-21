/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { Then } from "@cucumber/cucumber";
import { By } from "selenium-webdriver";

Then(
  "{string} sees an outgoing transaction log entry for the service {string} of {string}",
  async function (
    this: CustomWorld,
    orgName: string,
    serviceName: string,
    serviceOrgName: string
  ) {
    await validateTransactionLogRow(this, {
      direction: "outgoing",
      orgName,
      serviceName,
      serviceOrgName,
    });
  }
);

Then(
  "{string} sees an incoming transaction log entry from {string} for the service {string}",
  async function (
    this: CustomWorld,
    orgName: string,
    requestingOrgName: string,
    serviceName: string
  ) {
    await validateTransactionLogRow(this, {
      direction: "incoming",
      orgName,
      requestingOrgName,
      serviceName,
    });
  }
);

Then(
  "{string} sees an outgoing delegation transaction log entry for the service {string} of {string} on behalf of {string}",
  async function (
    this: CustomWorld,
    orgName: string,
    serviceName: string,
    serviceOrgName: string,
    delegatorOrgName: string
  ) {
    await validateTransactionLogRow(this, {
      direction: "outgoing",
      orgName,
      serviceName,
      serviceOrgName,
      delegatorOrgName,
    });
  }
);

Then(
  "{string} sees an incoming delegation transaction log entry from {string} for the service {string} on behalf of {string}",
  async function (
    this: CustomWorld,
    orgName: string,
    requestingOrgName: string,
    serviceName: string,
    delegatorOrgName: string
  ) {
    await validateTransactionLogRow(this, {
      direction: "incoming",
      orgName,
      requestingOrgName,
      serviceName,
      delegatorOrgName,
    });
  }
);

async function validateTransactionLogRow(
  world: CustomWorld,
  args: {
    direction: "incoming" | "outgoing";
    orgName: string;
    serviceName: string;
    serviceOrgName?: string;
    requestingOrgName?: string;
    delegatorOrgName?: string;
  }
) {
  const serviceName = `${args.serviceName}-${world.id}`;

  const { driver } = world;

  const org = getOrgByName(args.orgName);

  await driver.get(`${org.management.url}/transaction-log`);

  // Retrieve one log row containing the unique service name for this scenario from transaction log table
  const logRow = await driver.findElement(
    By.xpath(`//tr[td//text()[contains(., '${serviceName}')]]`)
  );

  let directionText: string;
  let orgText: string;

  if (args.direction === "incoming") {
    directionText = "Binnenkomend van";
    orgText = args.requestingOrgName as string;
  } else {
    directionText = "Uitgaand naar";
    orgText = args.serviceOrgName as string;
  }

  // Check if direction text is correct
  await logRow.findElement(By.xpath(`.//td[.//text()='${directionText}']`));

  if (args.delegatorOrgName) {
    orgText = `${orgText} namens ${args.delegatorOrgName}`;
  }

  // Check if organization text is correct
  await logRow.findElement(By.xpath(`.//td[.//text()='${orgText}']`));
}

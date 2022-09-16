/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { When } from "@cucumber/cucumber";

When(
  "{string} opens a non-existing page",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(`${org.management.url}/arbitrary-path`);
  }
);

When(
  "{string} opens the login page",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(`${org.management.url}/login`);
  }
);

When(
  "{string} opens the directory page",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(`${org.management.url}/directory`);
  }
);

When(
  "{string} opens the inways page",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(`${org.management.url}/inways-and-outways/inways`);
  }
);

When(
  "{string} opens the inway detail page of the default inway",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(
      `${org.management.url}/inways-and-outways/inways/${org.defaultInway.name}`
    );
  }
);

When(
  "{string} opens the services page",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;
    const org = getOrgByName(orgName);
    await driver.get(`${org.management.url}/services`);
  }
);

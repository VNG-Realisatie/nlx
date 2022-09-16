/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then(
  "the service {string} of {string} is created",
  async function (this: CustomWorld, serviceName: string, orgName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const org = getOrgByName(orgName);

    const response = await org.apiClients.management?.managementGetService({
      name: serviceName,
    });

    assert.equal(response?.name, serviceName);
  }
);

Then(
  "the service {string} of {string} is no longer available",
  async function (this: CustomWorld, serviceName: string, orgName: string) {
    serviceName = `${serviceName}-${this.id}`;

    const org = getOrgByName(orgName);

    try {
      await org.apiClients.management?.managementGetService({
        name: serviceName,
      });
      throw new Error(
        "this code should not be triggered, since we expect the service to be removed"
      );
    } catch (error: any) {
      if (error.response.status !== 404) {
        throw new Error(
          `unexpected status code '${error.response.status}' while getting a service, expected 404: ${error}`
        );
      }
    }
  }
);

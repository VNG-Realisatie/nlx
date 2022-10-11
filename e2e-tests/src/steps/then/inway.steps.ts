/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { isManagementAPIProxyAddressForDirectoryEqualTo } from "../../utils/inway";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then(
  "the default inway of {string} is no longer the organization inway",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    const resp = await org.apiClients.management?.managementGetSettings();

    assert.equal(
      resp?.settings?.organizationInway,
      "",
      `organization ${orgName} still has an inway set: '${resp?.settings?.organizationInway}'`
    );

    await isManagementAPIProxyAddressForDirectoryEqualTo(org, "");
  }
);

Then(
  "the default inway of {string} is removed",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    try {
      const inway = await org.apiClients.management?.managementGetInway({
        name: org.defaultInway.name,
      });

      throw new Error(
        `this code should not be triggered, since we expect the inway to be removed: ${JSON.stringify(
          inway
        )}`
      );
    } catch (error: any) {
      if (!error.response) {
        throw error;
      }

      if (error.response.status !== 404) {
        throw new Error(
          `unexpected status code '${error.response.status}' while getting a inway, expected 404: ${error}`
        );
      }
    }
  }
);

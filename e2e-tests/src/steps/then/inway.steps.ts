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

    const resp =
      await org.apiClients.management?.managementServiceGetSettings();

    assert.equal(
      resp?.settings?.organizationInway,
      "",
      `organization ${orgName} still has an inway set: '${resp?.settings?.organizationInway}'`
    );

    await isManagementAPIProxyAddressForDirectoryEqualTo(org, "");
  }
);

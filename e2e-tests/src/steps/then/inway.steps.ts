import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { isInwayAddressInDirectory } from "../../utils/inway";
import { Then } from "@cucumber/cucumber";
import pWaitFor from "p-wait-for";
import { strict as assert } from "assert";

Then(
  "the default inway of {string} is no longer the organization inway",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    const resp = await org.apiClients.management?.managementGetSettings();

    assert.equal(
      resp?.organizationInway,
      "",
      `organization ${orgName} still has an inway set: '${resp?.organizationInway}'`
    );

    // wait until the inway is unset as organization inway in the directory
    await pWaitFor.default(async () => await isInwayAddressInDirectory(org), {
      interval: 200,
      timeout: 1000 * 21,
    });
  }
);
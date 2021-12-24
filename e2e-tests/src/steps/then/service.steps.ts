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

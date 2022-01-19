import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { ManagementOutway } from "../../../../management-ui/src/api/models";
import { Given } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Given(
  "{string} has the default Outway running",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    const response = await org.apiClients.management?.managementListOutways();

    assert.equal(
      response?.outways?.some(
        (outway: ManagementOutway) => outway.name === org.defaultOutway.name
      ),
      true
    );
  }
);

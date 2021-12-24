import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { ManagementInway } from "../../../../management-ui/src/api/models";
import { Given } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Given(
  "{string} has the default Inway running",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    const response = await org.apiClients.management?.managementListInways();

    assert.equal(
      response?.inways?.some(
        (inway: ManagementInway) => inway.name === org.defaultInwayName
      ),
      true
    );
  }
);

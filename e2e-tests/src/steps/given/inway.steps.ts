import { CustomWorld } from "../../support/custom-world";
import { Organization, getOrgByName } from "../../utils/organizations";
import { ManagementInway } from "../../../../management-ui/src/api/models";
import { env } from "../../utils/env";
import { Given } from "@cucumber/cucumber";
import fetch from "cross-fetch";
import pWaitFor from "p-wait-for";
import { strict as assert } from "assert";

const isInwayAddressInDirectory = async (
  org: Organization
): Promise<boolean> => {
  const res = await fetch(
    `${env.directoryUrl}/api/directory/organizations/${org.serialNumber}/inway`
  );

  assert.equal(res.status >= 400, false);

  const inway = await res.json();

  if (
    inway.address === undefined ||
    inway.address !== org.defaultInway.address
  ) {
    return Promise.resolve(false);
  }

  return Promise.resolve(true);
};

Given(
  "{string} has the default Inway running",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    const response = await org.apiClients.management?.managementListInways();

    assert.equal(
      response?.inways?.some(
        (inway: ManagementInway) => inway.name === org.defaultInway.name
      ),
      true
    );
  }
);

Given(
  "{string} has set its default Inway as organization Inway",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);

    await org.apiClients.management?.managementUpdateSettings({
      body: {
        organizationInway: org.defaultInway.name,
      },
    });

    // wait until the inway is set as organization inway in the directory
    await pWaitFor.default(async () => await isInwayAddressInDirectory(org), {
      interval: 200,
      timeout: 1000 * 21,
    });
  }
);

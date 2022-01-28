import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has accepted the Terms of Service",
  async function (this: CustomWorld, orgName: string) {
    const org = getOrgByName(orgName);
    await org.apiClients.management?.managementAcceptTermsOfService();
  }
);

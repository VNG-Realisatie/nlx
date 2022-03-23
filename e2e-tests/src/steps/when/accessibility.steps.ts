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

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import {
  authenticateUsingBasicAuth,
  authenticateUsingOIDC,
} from "../../utils/authenticate";
import { When } from "@cucumber/cucumber";

When(
  "{string} authenticates using the wrong credentials",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const org = getOrgByName(orgName);
    await driver.get(org.management.url);

    if (org.management.basicAuth) {
      await authenticateUsingBasicAuth(
        this,
        orgName,
        "arbirary@example.com",
        "incorrect"
      );
    } else {
      await authenticateUsingOIDC(
        this,
        orgName,
        "arbirary@example.com",
        "incorrect",
        false
      );
    }
  }
);

When(
  "{string} authenticates using the right credentials",
  async function (this: CustomWorld, orgName: string) {
    const { driver } = this;

    const org = getOrgByName(orgName);
    await driver.get(org.management.url);

    if (org.management.basicAuth) {
      await authenticateUsingBasicAuth(
        this,
        orgName,
        org.management.username,
        org.management.password
      );
    } else {
      await authenticateUsingOIDC(
        this,
        orgName,
        org.management.username,
        org.management.password,
        true
      );
    }
  }
);

import { CustomWorld } from "../../support/custom-world";
import { logout } from "../../utils/authenticate";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} is logged out of NLX Management",
  async function (this: CustomWorld, orgName: string) {
    await logout(this, orgName);
  }
);

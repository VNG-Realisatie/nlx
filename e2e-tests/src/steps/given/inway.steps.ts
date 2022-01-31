import { CustomWorld } from "../../support/custom-world";
import { hasDefaultInwayRunning, setDefaultInway } from "../../utils/inway";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has the default Inway running",
  async function (this: CustomWorld, orgName: string) {
    await hasDefaultInwayRunning(this, orgName);
  }
);

Given(
  "{string} has set its default Inway as organization Inway",
  async function (this: CustomWorld, orgName: string) {
    await setDefaultInway(this, orgName);
  }
);

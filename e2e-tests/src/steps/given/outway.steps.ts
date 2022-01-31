import { CustomWorld } from "../../support/custom-world";
import { hasDefaultOutwayRunning } from "../../utils/outway";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has the default Outway running",
  async function (this: CustomWorld, orgName: string) {
    await hasDefaultOutwayRunning(this, orgName);
  }
);

import { CustomWorld } from "../../support/custom-world";
import { hasOutwayRunning } from "../../utils/outway";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has the Outway {string} running",
  async function (this: CustomWorld, orgName: string, outwayName: string) {
    await hasOutwayRunning(this, orgName, outwayName);
  }
);

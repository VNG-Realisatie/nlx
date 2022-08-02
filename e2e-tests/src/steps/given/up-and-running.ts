import { CustomWorld } from "../../support/custom-world";
import { authenticate } from "../../utils/authenticate";
import { acceptToS } from "../../utils/tos";
import { getOrgByName } from "../../utils/organizations";
import {
  hasDefaultInwayRunning,
  setDefaultInwayAsOrganizationInway,
} from "../../utils/inway";
import { Given } from "@cucumber/cucumber";
import {hasOutwayRunning} from "../../utils/outway";

Given(
  "{string} is up and running",
  async function (this: CustomWorld, orgName: string) {
    await authenticate(this, orgName);
    await acceptToS(this, orgName);

    const org = getOrgByName(orgName);

    if (org.defaultInway.name !== "") {
      await hasDefaultInwayRunning(this, orgName);
      await setDefaultInwayAsOrganizationInway(this, orgName);
    }

    if (org.outways !== null){
      for (const outwayName in org.outways){
        await hasOutwayRunning(this, orgName, outwayName)
      }
    }
  }
);

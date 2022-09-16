/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { authenticate } from "../../utils/authenticate";
import { acceptToS } from "../../utils/tos";
import { getOrgByName } from "../../utils/organizations";
import {
  hasDefaultInwayRunning,
  setDefaultInwayAsOrganizationInway,
} from "../../utils/inway";
import { hasOutwayRunning } from "../../utils/outway";
import { default as logger } from "../../debug";
import { Given } from "@cucumber/cucumber";
const debug = logger("e2e-tests:up-and-running");

Given(
  "{string} is up and running",
  async function (this: CustomWorld, orgName: string) {
    await authenticate(this, orgName);
    debug(`authentication for ${orgName} succeeded`);
    await acceptToS(this, orgName);
    debug(`ToS accepted for ${orgName}`);

    const org = getOrgByName(orgName);

    if (org.defaultInway.name !== "") {
      await hasDefaultInwayRunning(this, orgName);
      debug(`default Inway is running for ${orgName}`);
      await setDefaultInwayAsOrganizationInway(this, orgName);
      debug(`set default Inway as organization Inway for ${orgName}`);
    }

    if (org.outways !== null) {
      for (const outwayName in org.outways) {
        debug(`waiting until Outway ${outwayName} is running for ${orgName}`);
        await hasOutwayRunning(this, orgName, outwayName);
        debug(`Outway ${outwayName} is running for ${orgName}`);
      }
    }
  }
);

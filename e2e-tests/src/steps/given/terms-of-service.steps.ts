/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { acceptToS } from "../../utils/tos";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has accepted the Terms of Service",
  async function (this: CustomWorld, orgName: string) {
    await acceptToS(this, orgName);
  }
);

/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { authenticate } from "../../utils/authenticate";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} is logged into NLX management",
  async function (this: CustomWorld, orgName: string) {
    await authenticate(this, orgName);
  }
);

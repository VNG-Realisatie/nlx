/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { getOrgByName } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import { default as logger } from "../debug";
const debug = logger("e2e-tests:tos");

export const acceptToS = async (world: CustomWorld, orgName: string) => {
  debug(`accepting ToS for ${orgName}`);
  const org = getOrgByName(orgName);
  await org.apiClients.management?.managementAcceptTermsOfService();
};

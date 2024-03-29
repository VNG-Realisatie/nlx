/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { getOrgByName, Outway, Outways } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import { ResponseError } from "../../../management-ui/src/api/runtime";
import pWaitFor from "p-wait-for";

export const hasOutwayRunning = async (
  world: CustomWorld,
  orgName: string,
  outwayName: string
) => {
  await pWaitFor.default(
    async () => {
      const outways = await getOutways(orgName);
      return !!outways[outwayName];
    },
    {
      interval: 200,
      timeout: 1000 * 11,
    }
  );
};

export const getOutways = async (orgName: string): Promise<Outways> => {
  const org = getOrgByName(orgName);

  try {
    const outwaysResponse =
      await org.apiClients.management?.managementServiceListOutways();

    const outways = outwaysResponse?.outways;
    outways?.forEach((outway) => {
      if (!outway.name) {
        return;
      }

      org.outways[`${outway.name}`].name = outway.name || "";
      org.outways[`${outway.name}`].publicKeyPem = outway.publicKeyPem || "";
      org.outways[`${outway.name}`].publicKeyFingerprint =
        outway.publicKeyFingerprint || "";
    });

    return org.outways;
  } catch (error) {
    console.log("error failed to list outways: ", error);
    const response = error as ResponseError;
    const responseAsText = await response.response.text();

    throw new Error(`failed to list outways: ${responseAsText}`);
  }
};

export const getOutwayByName = async (
  orgName: string,
  outwayName: string
): Promise<Outway> => {
  const outways = await getOutways(orgName);

  return outways[outwayName];
};

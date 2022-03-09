import { getOrgByName } from "./organizations";
import { ManagementOutway } from "../../../management-ui/src/api/models";
import { CustomWorld } from "../support/custom-world";
import { strict as assert } from "assert";

interface Outways {
  [name: string]: ManagementOutway;
}

export const hasOutwayRunning = async (
  world: CustomWorld,
  orgName: string,
  outwayName: string
) => {
  const outways = await getOutways(orgName);

  assert.equal(!!outways[outwayName], true);
};

export const getOutways = async (orgName: string): Promise<Outways> => {
  const org = getOrgByName(orgName);

  const response = await org.apiClients.management?.managementListOutways();

  const outways = {} as Outways;

  response?.outways?.forEach((o) => {
    outways[`${o.name}`] = o;
  });

  return outways;
};

export const getOutwayByName = async (
  orgName: string,
  outwayName: string
): Promise<ManagementOutway> => {
  const outways = await getOutways(orgName);

  return outways[outwayName];
};

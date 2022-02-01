import { getOrgByName } from "./organizations";
import { authenticate } from "./authenticate";
import { acceptToS } from "./tos";
import { ManagementOutway } from "../../../management-ui/src/api/models";
import { CustomWorld } from "../support/custom-world";
import { strict as assert } from "assert";

export const hasDefaultOutwayRunning = async (
  world: CustomWorld,
  orgName: string
) => {
  await authenticate(world, orgName);
  await acceptToS(world, orgName);
  const org = getOrgByName(orgName);

  const response = await org.apiClients.management?.managementListOutways();

  assert.equal(
    response?.outways?.some(
      (outway: ManagementOutway) => outway.name === org.defaultOutway.name
    ),
    true
  );
};

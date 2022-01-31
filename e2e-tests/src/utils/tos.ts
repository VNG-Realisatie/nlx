import { getOrgByName } from "./organizations";
import { CustomWorld } from "../support/custom-world";

export const acceptToS = async (world: CustomWorld, orgName: string) => {
  const org = getOrgByName(orgName);
  await org.apiClients.management?.managementAcceptTermsOfService();
};

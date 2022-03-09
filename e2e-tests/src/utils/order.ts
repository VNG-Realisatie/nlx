import { getOrgByName, Organization } from "./organizations";
import { getOutwayByName, getOutways } from "./outway";
import { CustomWorld } from "../support/custom-world";
import { default as logger } from "../debug";
import { ManagementOutway } from "../../../management-ui/src/api/models/ManagementOutway";
import dayjs from "dayjs";
const debug = logger("e2e-tests:order");

export const createOrder = async (
  world: CustomWorld,
  delegateeOrgName: string,
  delegateeOutwayName: string,
  orderReference: string,
  delegatorOrgName: string,
  serviceName: string,
  serviceProviderOrgName: string,
  delegatorOutwayName: string
) => {
  serviceName = `${serviceName}-${world.id}`;
  orderReference = `${orderReference}-${world.id}`;

  const delegator = getOrgByName(delegatorOrgName);
  const delegatee = getOrgByName(delegateeOrgName);
  const serviceProvider = getOrgByName(serviceProviderOrgName);

  debug(
    `creating an order for delegator '${delegatorOrgName} (${delegator.serialNumber})' with reference '${orderReference}' and service '${serviceName} by ${serviceProviderOrgName} (${serviceProvider.serialNumber})' delegated to '${delegateeOrgName} (${delegatee.serialNumber})'`
  );

  const delegatorOutway = await getOutwayByName(
    delegatorOrgName,
    delegatorOutwayName
  );

  const directoryServices =
    await delegator.apiClients.directory?.directoryListServices();

  const directoryService = directoryServices?.services?.find(
    (directoryService) =>
      directoryService.serviceName === serviceName &&
      directoryService?.organization?.serialNumber ===
        serviceProvider.serialNumber
  );

  const accessStateForService = directoryService?.accessStates?.find(
    (accessState) =>
      accessState?.accessProof?.publicKeyFingerprint ===
      delegatorOutway.publicKeyFingerprint
  );

  if (!accessStateForService || !accessStateForService.accessProof) {
    throw Error(`could not find access proof for service '${serviceName}'`);
  }

  const delegateeOutway = await getOutwayByName(
    delegateeOrgName,
    delegateeOutwayName
  );

  try {
    await delegator.apiClients.management?.managementCreateOutgoingOrder({
      body: {
        reference: orderReference,
        description: "arbitrary description",
        delegatee: delegatee.serialNumber,
        publicKeyPEM: delegateeOutway.publicKeyPEM,
        validFrom: dayjs().subtract(1, "day").toDate(),
        validUntil: dayjs().add(1, "day").toDate(),
        accessProofIds: [`${accessStateForService?.accessProof?.id}`],
      },
    });
  } catch (err) {
    const response = err as Response;
    const responseAsJson = await response.json();

    throw new Error(
      `failed to create outgoing order: ${responseAsJson.message}`
    );
  }

  debug(`created order with reference '${orderReference}'`);
};

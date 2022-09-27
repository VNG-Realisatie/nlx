/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { getOrgByName } from "./organizations";
import { getOutwayByName } from "./outway";
import { CustomWorld } from "../support/custom-world";
import { default as logger } from "../debug";
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
  delegatorOutwayName: string,
  validFrom: Date,
  validUntil: Date
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
        publicKeyPem: delegateeOutway.publicKeyPem,
        validFrom: validFrom,
        validUntil: validUntil,
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

export const revokeOrder = async (
  world: CustomWorld,
  delegatorOrgName: string,
  delegateeOrgName: string,
  orderReference: string
) => {
  orderReference = `${orderReference}-${world.id}`;

  const delegator = getOrgByName(delegatorOrgName);
  const delegatee = getOrgByName(delegateeOrgName);

  debug(
    `revoking the order with reference '${orderReference}' for delegator '${delegatorOrgName} (${delegator.serialNumber})' delegated to '${delegateeOrgName} (${delegatee.serialNumber})'`
  );

  try {
    await delegator.apiClients.management?.managementRevokeOutgoingOrder({
      delegatee: delegatee.serialNumber,
      reference: orderReference,
    });
  } catch (err) {
    const response = err as Response;
    const responseAsJson = await response.json();

    throw new Error(`failed to revoke order: ${responseAsJson.message}`);
  }

  debug(`revoked order with reference '${orderReference}'`);
};

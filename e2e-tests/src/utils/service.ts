/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { getOrgByName, Organization } from "./organizations";
import { getOutwayByName } from "./outway";
import { CustomWorld } from "../support/custom-world";
import {
  ManagementDirectoryService,
  ManagementIncomingAccessRequest,
} from "../../../management-ui/src/api/models";
import { default as logger } from "../debug";
import { isServiceKnownInServiceListOfOutway } from "../steps/when/outway.steps";
import pWaitFor from "p-wait-for";
import { strict as assert } from "assert";
const debug = logger("e2e-tests:service");

const isAccessRequestApprovedForService = async (
  uniqueServiceName: string,
  org: Organization,
  publicKeyFingerprint: string
): Promise<boolean> => {
  const directoryService = await getDirectoryService(uniqueServiceName, org);

  if (!directoryService || !directoryService.accessStates) {
    return Promise.resolve(false);
  }

  const accessState = directoryService.accessStates.find((accessState) => {
    return (
      accessState.accessRequest?.publicKeyFingerprint === publicKeyFingerprint
    );
  });

  if (
    !accessState ||
    !accessState.accessProof ||
    accessState.accessProof.revokedAt
  ) {
    return Promise.resolve(false);
  }

  return Promise.resolve(true);
};

const getDirectoryService = async (
  serviceName: string,
  org: Organization
): Promise<ManagementDirectoryService | undefined> => {
  const directoryServicesResponse =
    await org.apiClients.directory?.directoryListServices();

  return directoryServicesResponse?.services?.find(
    (service) => service.serviceName === serviceName
  );
};

const getIncomingAccessRequest = async (
  uniqueServiceName: string,
  org: Organization,
  serviceProvider: Organization,
  publicKeyFingerprint: string
): Promise<ManagementIncomingAccessRequest | undefined> => {
  const responseIncomingAccessRequests =
    await serviceProvider.apiClients.management?.managementListIncomingAccessRequests(
      {
        serviceName: uniqueServiceName,
      }
    );

  if (
    !responseIncomingAccessRequests ||
    !responseIncomingAccessRequests.accessRequests
  ) {
    return;
  }

  const incomingAccessRequest =
    responseIncomingAccessRequests.accessRequests.find(
      (accessRequest: ManagementIncomingAccessRequest) => {
        return (
          accessRequest.serviceName === uniqueServiceName &&
          accessRequest.organization?.serialNumber === org.serialNumber &&
          accessRequest.publicKeyFingerprint === publicKeyFingerprint
        );
      }
    );

  return incomingAccessRequest;
};

const isServicePresentInDirectory = async (
  serviceProvider: Organization,
  uniqueServiceName: string
): Promise<boolean> => {
  const result = await getServiceFromDirectory(
    serviceProvider,
    uniqueServiceName
  );
  return Promise.resolve(!!result);
};

const getServiceFromDirectory = async (
  serviceProvider: Organization,
  uniqueServiceName: string
): Promise<ManagementDirectoryService | undefined> => {
  try {
    return await serviceProvider.apiClients.directory?.directoryGetOrganizationService(
      {
        organizationSerialNumber: serviceProvider.serialNumber,
        serviceName: uniqueServiceName,
      }
    );
  } catch (error) {
    return;
  }
};

export const createService = async (
  world: CustomWorld,
  serviceName: string,
  serviceProviderOrgName: string
): Promise<string> => {
  debug(`creating service ${serviceName} for ${serviceProviderOrgName}`);

  const serviceProvider = getOrgByName(serviceProviderOrgName);
  const uniqueServiceName = `${serviceName}-${world.id}`;

  const createServiceResponse =
    await serviceProvider.apiClients.management?.managementCreateService({
      body: {
        name: uniqueServiceName,
        endpointUrl: "https://postman-echo.com",
        inways: [serviceProvider.defaultInway.name],
        internal: false,
      },
    });
  assert.equal(createServiceResponse?.name, uniqueServiceName);

  serviceProvider.createdItems[world.id].services.push(uniqueServiceName);

  debug(
    `successfully created service ${serviceName} for ${serviceProviderOrgName}`
  );

  debug(
    `waiting until ${serviceName} for ${serviceProviderOrgName} is present in the directory`
  );

  // wait until the service has been announced to the directory
  await pWaitFor.default(
    async () =>
      await isServicePresentInDirectory(serviceProvider, uniqueServiceName),
    {
      interval: 200,
      timeout: 1000 * 35,
    }
  );

  return uniqueServiceName;
};

export const getAccessToService = async (
  world: CustomWorld,
  serviceConsumerOrgName: string,
  serviceName: string,
  serviceProviderOrgName: string,
  outwayName: string
) => {
  debug(
    `${serviceConsumerOrgName} is requesting access to service ${serviceName} of ${serviceProviderOrgName}`
  );
  const serviceProvider = getOrgByName(serviceProviderOrgName);
  const serviceConsumer = getOrgByName(serviceConsumerOrgName);

  const uniqueServiceName = await createService(
    world,
    serviceName,
    serviceProviderOrgName
  );

  const outway = await getOutwayByName(serviceConsumerOrgName, outwayName);

  const url = `${outway.selfAddress}/${serviceProvider.serialNumber}/${uniqueServiceName}/get`;

  // wait until the Outway has had the time to update its internal services list
  await pWaitFor.default(
    async () => await isServiceKnownInServiceListOfOutway(url),
    {
      interval: 1000,
      timeout: 1000 * 90,
    }
  );

  // request access to new service
  const createAccessRequestResponse =
    await serviceConsumer.apiClients.management?.managementSendAccessRequest({
      organizationSerialNumber: serviceProvider.serialNumber,
      serviceName: uniqueServiceName,
      publicKeyPem: outway.publicKeyPem,
    });

  assert.equal(
    createAccessRequestResponse?.outgoingAccessRequest?.serviceName,
    uniqueServiceName
  );

  const incomingAccessRequest = await getIncomingAccessRequest(
    uniqueServiceName,
    serviceConsumer,
    serviceProvider,
    outway.publicKeyFingerprint || ""
  );

  assert.notEqual(incomingAccessRequest, undefined);

  await serviceProvider.apiClients.management?.managementApproveIncomingAccessRequest(
    {
      serviceName: uniqueServiceName,
      accessRequestId: incomingAccessRequest?.id as string,
    }
  );

  await serviceConsumer.apiClients.management?.managementSynchronizeOutgoingAccessRequests(
    {
      organizationSerialNumber: serviceProvider.serialNumber,
      serviceName: uniqueServiceName,
    }
  );

  await isAccessRequestApprovedForService(
    uniqueServiceName,
    serviceConsumer,
    outway.publicKeyFingerprint || ""
  );

  // Issue #1613: wait until the Inway has refreshed its internal list of services + access grants
  await new Promise((resolve) => {
    setTimeout(() => {
      resolve(null)
    }, 10 * 1000)
  })

  debug(
    `${serviceConsumerOrgName} has gotten access to service ${serviceName} of ${serviceProviderOrgName}`
  );
};

import { getOrgByName, Organization } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import {
  ManagementDirectoryService,
  ManagementIncomingAccessRequest,
} from "../../../management-ui/src/api/models";
import { default as logger } from "../debug";
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

const isIncomingAccessRequestPresent = async (
  uniqueServiceName: string,
  org: Organization,
  serviceProvider: Organization,
  publicKeyFingerprint: string
): Promise<boolean> => {
  const result = await getIncomingAccessRequest(
    uniqueServiceName,
    org,
    serviceProvider,
    publicKeyFingerprint
  );
  return Promise.resolve(!!result);
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

export const getAccessToService = async (
  world: CustomWorld,
  serviceConsumerOrgName: string,
  serviceName: string,
  serviceProviderOrgName: string
) => {
  debug(
    `${serviceConsumerOrgName} is requesting access to service ${serviceName} of ${serviceProviderOrgName}`
  );
  const serviceProvider = getOrgByName(serviceProviderOrgName);
  const serviceConsumer = getOrgByName(serviceConsumerOrgName);

  const uniqueServiceName = `${serviceName}-${world.id}`;

  const createServiceResponse =
    await serviceProvider.apiClients.management?.managementCreateService({
      body: {
        name: uniqueServiceName,
        endpointURL: "https://postman-echo.com",
        inways: [serviceProvider.defaultInway.name],
        internal: false,
      },
    });
  assert.equal(createServiceResponse?.name, uniqueServiceName);

  // retrieve public key fingerprint of the default Outway
  const outwaysResponse =
    await serviceConsumer.apiClients.management?.managementListOutways();
  assert.equal(!!outwaysResponse?.outways?.length, true);

  const defaultOutway = outwaysResponse?.outways?.find(
    (outway) => outway.name === serviceConsumer.defaultOutway.name
  );

  if (!defaultOutway) {
    throw new Error(
      `unable to find the default Outway for ${serviceConsumerOrgName}`
    );
  }

  // request access to new service
  const createAccessRequestResponse =
    await serviceConsumer.apiClients.management?.managementCreateAccessRequest({
      body: {
        organizationSerialNumber: serviceProvider.serialNumber,
        serviceName: uniqueServiceName,
        publicKeyPEM: defaultOutway?.publicKeyPEM,
      },
    });
  assert.equal(createAccessRequestResponse?.serviceName, uniqueServiceName);

  // wait until the other organization has received our access request
  await pWaitFor.default(
    async () =>
      await isIncomingAccessRequestPresent(
        uniqueServiceName,
        serviceConsumer,
        serviceProvider,
        defaultOutway.publicKeyFingerprint || ""
      ),
    {
      interval: 200,
      timeout: 1000 * 30,
    }
  );

  const incomingAccessRequest = await getIncomingAccessRequest(
    uniqueServiceName,
    serviceConsumer,
    serviceProvider,
    defaultOutway.publicKeyFingerprint || ""
  );

  assert.notEqual(incomingAccessRequest, undefined);

  await serviceProvider.apiClients.management?.managementApproveIncomingAccessRequest(
    {
      serviceName: uniqueServiceName,
      accessRequestID: incomingAccessRequest?.id as string,
    }
  );

  // wait until the other organization has retrieved the approval of the access request
  await pWaitFor.default(
    async () =>
      await isAccessRequestApprovedForService(
        uniqueServiceName,
        serviceConsumer,
        defaultOutway?.publicKeyFingerprint || ""
      ),
    {
      interval: 200,
      timeout: 1000 * 35, // TODO: we dont know how long it takes until an approval is being synced to the other organization
    }
  );

  debug(
    `${serviceConsumerOrgName} has gotten access to service ${serviceName} of ${serviceProviderOrgName}`
  );
};

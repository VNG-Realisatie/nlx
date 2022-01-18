import { CustomWorld } from "../../support/custom-world";
import { Organization, getOrgByName } from "../../utils/organizations";
import { authenticate } from "../../utils/authenticate";
import {
  ManagementDirectoryService,
  ManagementIncomingAccessRequest,
} from "../../../../management-ui/src/api/models";
import { Given } from "@cucumber/cucumber";
import pWaitFor from "p-wait-for";
import { strict as assert } from "assert";

const isAccessRequestApprovedForService = async (
  uniqueServiceName: string,
  org: Organization
): Promise<boolean> => {
  const directoryService = await getDirectoryService(uniqueServiceName, org);

  if (
    !directoryService ||
    !directoryService.latestAccessProof ||
    directoryService.latestAccessProof.revokedAt
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
  serviceProvider: Organization
): Promise<boolean> => {
  const result = await getIncomingAccessRequest(
    uniqueServiceName,
    org,
    serviceProvider
  );
  return Promise.resolve(!!result);
};

const getIncomingAccessRequest = async (
  uniqueServiceName: string,
  org: Organization,
  serviceProvider: Organization
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
      (accessRequest: ManagementIncomingAccessRequest) =>
        accessRequest.serviceName === uniqueServiceName &&
        accessRequest.organization?.serialNumber === org.serialNumber
    );

  return incomingAccessRequest;
};

Given(
  "{string} has access to {string} of {string}",
  async function (
    this: CustomWorld,
    serviceConsumerOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string
  ) {
    const serviceProvider = getOrgByName(serviceProviderOrgName);
    const serviceConsumer = getOrgByName(serviceConsumerOrgName);

    const uniqueServiceName = `${serviceName}-${this.id}`;

    // create the service for second org
    await authenticate(this, serviceProviderOrgName);

    const createServiceResponse =
      await serviceProvider.apiClients.management?.managementCreateService({
        body: {
          name: uniqueServiceName,
          endpointURL: "http://example.com",
          inways: [serviceProvider.defaultInwayName],
          internal: false,
        },
      });
    assert.equal(createServiceResponse?.name, uniqueServiceName);

    // request access to new service
    const createAccessRequestResponse =
      await serviceConsumer.apiClients.management?.managementCreateAccessRequest({
        body: {
          organizationSerialNumber: serviceProvider.serialNumber,
          serviceName: uniqueServiceName,
        },
      });
    assert.equal(createAccessRequestResponse?.serviceName, uniqueServiceName);

    // wait until the other organization has received our access request
    await pWaitFor.default(
      async () =>
        await isIncomingAccessRequestPresent(uniqueServiceName, serviceConsumer, serviceProvider),
      {
        interval: 200,
        timeout: 6500,
      }
    );

    const incomingAccessRequest = await getIncomingAccessRequest(
      uniqueServiceName,
      serviceConsumer,
      serviceProvider
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
        await isAccessRequestApprovedForService(uniqueServiceName, serviceConsumer),
      {
        interval: 200,
        timeout: 1000 * 35, // TODO: we dont know how long it takes until an approval is being synced to the other organization
      }
    );
  }
);

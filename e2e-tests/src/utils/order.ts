import { getOrgByName, Organization } from "./organizations";
import { CustomWorld } from "../support/custom-world";
import { default as logger } from "../debug";
import dayjs from "dayjs";
const debug = logger("e2e-tests:order");

export const createOrder = async (
  world: CustomWorld,
  delegateeOrgName: string,
  orderReference: string,
  delegatorOrgName: string,
  serviceName: string,
  serviceProviderOrgName: string
) => {
  serviceName = `${serviceName}-${world.id}`;
  orderReference = `${orderReference}-${world.id}`;

  const delegator = getOrgByName(delegatorOrgName);
  const delegatee = getOrgByName(delegateeOrgName);
  const serviceProvider = getOrgByName(serviceProviderOrgName);

  debug(
    `creating an order for delegator '${delegatorOrgName} (${delegator.serialNumber})' with reference '${orderReference}' and service '${serviceName} by ${serviceProviderOrgName} (${serviceProvider.serialNumber})' delegated to '${delegateeOrgName} (${delegatee.serialNumber})'`
  );

  const defaultOutwayPublicKeyPEM =
    await getDefaultOutwayPublicKeyPemForOrganization(delegateeOrgName);

  try {
    await delegator.apiClients.management?.managementCreateOutgoingOrder({
      body: {
        reference: orderReference,
        description: "arbitrary description",
        delegatee: delegatee.serialNumber,
        publicKeyPEM: defaultOutwayPublicKeyPEM,
        validFrom: dayjs().subtract(1, "day").toDate(),
        validUntil: dayjs().add(1, "day").toDate(),
        services: [
          {
            organization: {
              serialNumber: serviceProvider.serialNumber,
              name: serviceProviderOrgName,
            },
            service: serviceName,
          },
        ],
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

const getDefaultOutwayPublicKeyPemForOrganization = async (
  organizationName: string
): Promise<string | undefined> => {
  const organization: Organization = getOrgByName(organizationName);

  const outwaysResponse =
    await organization.apiClients.management?.managementListOutways();

  if (!outwaysResponse) {
    throw new Error(
      `unable to retrieve Outways for organization '${organizationName}'`
    );
  }

  const defaultOutway = outwaysResponse.outways?.find(
    (outway) => outway.name === organization.defaultOutway.name
  );

  return defaultOutway?.publicKeyPEM;
};

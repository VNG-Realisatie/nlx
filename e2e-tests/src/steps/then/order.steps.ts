import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import {
  ManagementOutgoingOrder,
  ManagementOrderService,
} from "../../../../management-ui/src/api/models";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";

Then(
  "an order of {string} with reference {string} for {string} with service {string} of {string} is created",
  async function (
    this: CustomWorld,
    delegatorOrgName: string,
    orderReference: string,
    delegateeOrgName: string,
    serviceName: string,
    orgName: string
  ) {
    serviceName = `${serviceName}-${this.id}`;
    orderReference = `${orderReference}-${this.id}`;

    const delegator = getOrgByName(delegatorOrgName);
    const delegatee = getOrgByName(delegateeOrgName);
    const org = getOrgByName(orgName);

    const orderResponse =
      await delegator.apiClients.management?.managementListOutgoingOrders();

    assert.equal(
      orderResponse?.orders?.some(
        (order: ManagementOutgoingOrder) =>
          order.reference === orderReference &&
          order.delegatee === delegatee.serialNumber &&
          order.revokedAt === undefined &&
          order.services?.some(
            (service: ManagementOrderService) =>
              service.service === serviceName &&
              service.organization?.serialNumber === org.serialNumber
          )
      ),
      true,
      `cannot find order, values used: orders: ${orderResponse?.orders}, serviceName: ${serviceName}, delegatee: ${delegatee.serialNumber}, org: ${org.serialNumber}`
    );
  }
);

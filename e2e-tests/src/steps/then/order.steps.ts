/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { getOrgByName } from "../../utils/organizations";
import {
  ManagementIncomingOrder,
  ManagementOutgoingOrder,
} from "../../../../management-ui/src/api/models";
import { default as logger } from "../../debug";
import { Then } from "@cucumber/cucumber";
import { strict as assert } from "assert";
const debug = logger("e2e-tests:order");

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
          order.delegatee?.serialNumber === delegatee.serialNumber &&
          order.revokedAt === undefined &&
          order.accessProofs?.some(
            (accessProof) =>
              accessProof.serviceName === serviceName &&
              accessProof.organization?.serialNumber === org.serialNumber
          )
      ),
      true,
      `cannot find order, values used: orders: ${orderResponse?.orders}, serviceName: ${serviceName}, delegatee: ${delegatee.serialNumber}, org: ${org.serialNumber}`
    );
  }
);

Then(
  "{string} has a revoked incoming order from {string} with reference {string}",
  async function (
    this: CustomWorld,
    delegateeOrgName: string,
    delegatorOrgName: string,
    orderReference: string
  ) {
    orderReference = `${orderReference}-${this.id}`;

    const delegator = getOrgByName(delegatorOrgName);
    const delegatee = getOrgByName(delegateeOrgName);

    await delegatee.apiClients.management?.managementSynchronizeOrders();

    const orderResponse =
      await delegatee.apiClients.management?.managementListIncomingOrders();

    debug(
      "incoming order",
      orderResponse?.orders?.find((order) => order.reference === orderReference)
    );

    assert.equal(
      orderResponse?.orders?.some(
        (order: ManagementIncomingOrder) =>
          order.reference === orderReference &&
          order.delegator?.serialNumber === delegator.serialNumber &&
          order.revokedAt !== undefined
      ),
      true,
      `cannot find incoming order, values used: reference: ${orderReference}, delegator: ${delegator.serialNumber}, org: ${delegatee.serialNumber}`
    );
  }
);

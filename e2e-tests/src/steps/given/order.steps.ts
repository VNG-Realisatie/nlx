/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import { CustomWorld } from "../../support/custom-world";
import { setDefaultInwayAsOrganizationInway } from "../../utils/inway";
import { createOrder, revokeOrder } from "../../utils/order";
import { authenticate } from "../../utils/authenticate";
import { Given } from "@cucumber/cucumber";
import dayjs from "dayjs";

Given(
  "{string} has an active order for Outway {string} with reference {string} from {string} for service {string} of {string} via Outway {string}",
  async function (
    this: CustomWorld,
    delegateeOrgName: string,
    delegateeOutwayName: string,
    orderReference: string,
    delegatorOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string,
    delegatorOutwayName: string
  ) {
    await authenticate(this, serviceProviderOrgName);
    await authenticate(this, delegatorOrgName);

    await setDefaultInwayAsOrganizationInway(this, serviceProviderOrgName);
    await setDefaultInwayAsOrganizationInway(this, delegatorOrgName);

    await createOrder(
      this,
      delegateeOrgName,
      delegateeOutwayName,
      orderReference,
      delegatorOrgName,
      serviceName,
      serviceProviderOrgName,
      delegatorOutwayName,
      dayjs().subtract(1, "day").toDate(),
      dayjs().add(1, "day").toDate()
    );
  }
);

Given(
  "{string} has a revoked order for Outway {string} with reference {string} from {string} for service {string} of {string} via Outway {string}",
  async function (
    this: CustomWorld,
    delegateeOrgName: string,
    delegateeOutwayName: string,
    orderReference: string,
    delegatorOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string,
    delegatorOutwayName: string
  ) {
    await authenticate(this, serviceProviderOrgName);
    await authenticate(this, delegatorOrgName);

    await setDefaultInwayAsOrganizationInway(this, serviceProviderOrgName);
    await setDefaultInwayAsOrganizationInway(this, delegatorOrgName);

    await createOrder(
      this,
      delegateeOrgName,
      delegateeOutwayName,
      orderReference,
      delegatorOrgName,
      serviceName,
      serviceProviderOrgName,
      delegatorOutwayName,
      dayjs().subtract(1, "day").toDate(),
      dayjs().add(1, "day").toDate()
    );

    await revokeOrder(this, delegatorOrgName, delegateeOrgName, orderReference);
  }
);

Given(
  "{string} has an expired order for Outway {string} with reference {string} from {string} for service {string} of {string} via Outway {string}",
  async function (
    this: CustomWorld,
    delegateeOrgName: string,
    delegateeOutwayName: string,
    orderReference: string,
    delegatorOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string,
    delegatorOutwayName: string
  ) {
    await authenticate(this, serviceProviderOrgName);
    await authenticate(this, delegatorOrgName);

    await setDefaultInwayAsOrganizationInway(this, serviceProviderOrgName);
    await setDefaultInwayAsOrganizationInway(this, delegatorOrgName);

    await createOrder(
      this,
      delegateeOrgName,
      delegateeOutwayName,
      orderReference,
      delegatorOrgName,
      serviceName,
      serviceProviderOrgName,
      delegatorOutwayName,
      dayjs().subtract(2, "day").toDate(),
      dayjs().subtract(1, "day").toDate()
    );
  }
);

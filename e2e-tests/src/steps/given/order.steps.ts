import { CustomWorld } from "../../support/custom-world";
import { setDefaultInwayAsOrganizationInway } from "../../utils/inway";
import { createOrder, revokeOrder } from "../../utils/order";
import { authenticate } from "../../utils/authenticate";
import { Given } from "@cucumber/cucumber";

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
      delegatorOutwayName
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
      delegatorOutwayName
    );

    await revokeOrder(this, delegatorOrgName, delegateeOrgName, orderReference);
  }
);

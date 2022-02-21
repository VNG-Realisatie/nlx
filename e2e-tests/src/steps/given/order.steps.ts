import { CustomWorld } from "../../support/custom-world";
import { setDefaultInwayAsOrganizationInway } from "../../utils/inway";
import { getAccessToService } from "../../utils/service";
import { createOrder } from "../../utils/order";
import { authenticate } from "../../utils/authenticate";
import { Given } from "@cucumber/cucumber";

Given(
  "{string} has an active order with reference {string} from {string} for service {string} of {string}",
  async function (
    this: CustomWorld,
    delegateeOrgName: string,
    orderReference: string,
    delegatorOrgName: string,
    serviceName: string,
    serviceProviderOrgName: string
  ) {
    await authenticate(this, serviceProviderOrgName);
    await authenticate(this, delegatorOrgName);

    await setDefaultInwayAsOrganizationInway(this, serviceProviderOrgName);
    await setDefaultInwayAsOrganizationInway(this, delegatorOrgName);

    await getAccessToService(
      this,
      delegatorOrgName,
      serviceName,
      serviceProviderOrgName
    );

    await createOrder(
      this,
      delegateeOrgName,
      orderReference,
      delegatorOrgName,
      serviceName,
      serviceProviderOrgName
    );
  }
);

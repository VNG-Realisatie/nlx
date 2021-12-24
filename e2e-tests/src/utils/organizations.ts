import { DirectoryApi, ManagementApi } from "../../../management-ui/src/api";
import { strict as assert } from "assert";

export interface Organization {
  serialNumber: string;
  defaultInwayName: string;
  management: {
    basicAuth: boolean;
    url: string;
    username: string;
    password: string;
  };
  apiClients: {
    management: ManagementApi | undefined;
    directory: DirectoryApi | undefined;
  };
}

interface Organizations {
  [key: string]: Organization;
}

const convertToBool = (
  input: string | undefined,
  defaultResult: boolean
): boolean => {
  if (typeof input === "undefined") {
    return defaultResult;
  }

  return input === "true";
};

export const organizations: Organizations = {
  "Gemeente Stijns": {
    serialNumber: "12345678901234567890",
    defaultInwayName:
      process.env.E2E_GEMEENTE_STIJNS_DEFAULT_INWAY_NAME || "Inway-01",
    management: {
      basicAuth: convertToBool(
        process.env.E2E_GEMEENTE_STIJNS_MANAGEMENT_BASIC_AUTH,
        false
      ),
      url:
        process.env.E2E_GEMEENTE_STIJNS_MANAGEMENT_URL ||
        "http://management.organization-a.nlx.local:3011",
      username:
        process.env.E2E_GEMEENTE_STIJNS_MANAGEMENT_USERNAME ||
        "admin@nlx.local",
      password:
        process.env.E2E_GEMEENTE_STIJNS_MANAGEMENT_PASSWORD || "development",
    },
    apiClients: {
      management: undefined,
      directory: undefined,
    },
  },
  RvRD: {
    serialNumber: "12345678901234567891",
    defaultInwayName: process.env.E2E_RVRD_DEFAULT_INWAY_NAME || "Inway-01",
    management: {
      basicAuth: convertToBool(
        process.env.E2E_RVRD_MANAGEMENT_BASIC_AUTH,
        true
      ),
      url:
        process.env.E2E_RVRD_MANAGEMENT_URL ||
        "http://management.organization-b.nlx.local:3021",
      username: process.env.E2E_RVRD_MANAGEMENT_USERNAME || "admin@nlx.local",
      password: process.env.E2E_RVRD_MANAGEMENT_PASSWORD || "development",
    },
    apiClients: {
      management: undefined,
      directory: undefined,
    },
  },
  "Vergunningsoftware BV": {
    serialNumber: "12345678901234567892",
    defaultInwayName:
      process.env.E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_INWAY_NAME || "Inway-01",
    management: {
      basicAuth: convertToBool(
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_BASIC_AUTH,
        false
      ),
      url: process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_URL || "TODO",
      username:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_USERNAME || "TODO",
      password:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_PASSWORD || "TODO",
    },
    apiClients: {
      management: undefined,
      directory: undefined,
    },
  },
};

export function getOrgByName(name: string): Organization {
  const org = organizations[name];
  assert.notEqual(org, undefined, `could not find org named '${name}'`);

  return org;
}

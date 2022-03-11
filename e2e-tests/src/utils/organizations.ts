import { DirectoryApi, ManagementApi } from "../../../management-ui/src/api";
import { strict as assert } from "assert";

export interface Organization {
  serialNumber: string;
  defaultInway: {
    name: string;
    address: string;
  };
  outways: Outways;
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

export interface Outway {
  name: string;
  publicKeyFingerprint: string;
  publicKeyPEM: string;
  selfAddress: string;
}

export interface Outways {
  [name: string]: Outway;
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
    defaultInway: {
      name: process.env.E2E_GEMEENTE_STIJNS_DEFAULT_INWAY_NAME || "Inway-01",
      address:
        process.env.E2E_GEMEENTE_STIJNS_DEFAULT_INWAY_ADDRESS ||
        "inway.organization-a.nlx.local:7913",
    },
    outways: {
      [process.env.E2E_GEMEENTE_STIJNS_OUTWAY_1_NAME ||
      "gemeente-stijns-nlx-outway"]: {
        name: "",
        selfAddress:
          process.env.E2E_GEMEENTE_STIJNS_OUTWAY_1_ADDRESS ||
          "http://127.0.0.1:7917",
        publicKeyFingerprint: "",
        publicKeyPEM: "",
      },
      [process.env.E2E_GEMEENTE_STIJNS_OUTWAY_2_NAME ||
      "gemeente-stijns-nlx-outway-2"]: {
        name: "",
        selfAddress:
          process.env.E2E_GEMEENTE_STIJNS_OUTWAY_2_ADDRESS ||
          "http://127.0.0.1:7947",
        publicKeyFingerprint: "",
        publicKeyPEM: "",
      },
    },
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
    defaultInway: {
      name: process.env.E2E_RVRD_DEFAULT_INWAY_NAME || "Inway-01",
      address:
        process.env.E2E_RVRD_DEFAULT_INWAY_ADDRESS ||
        "inway.organization-b.nlx.local:7923",
    },
    outways: {},
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
    defaultInway: {
      name: process.env.E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_INWAY_NAME || "",
      address:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_DEFAULT_INWAY_ADDRESS || "",
    },
    outways: {
      [process.env.E2E_VERGUNNINGSOFTWARE_BV_OUTWAY_1_NAME ||
      "vergunningsoftware-bv-nlx-outway"]: {
        name: "",
        selfAddress:
          process.env.E2E_VERGUNNINGSOFTWARE_BV_OUTWAY_1_ADDRESS ||
          "http://127.0.0.1:7937",
        publicKeyFingerprint: "",
        publicKeyPEM: "",
      },
    },
    management: {
      basicAuth: convertToBool(
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_BASIC_AUTH,
        true
      ),
      url:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_URL ||
        "http://management.organization-c.nlx.local:3031",
      username:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_USERNAME ||
        "admin@nlx.local",
      password:
        process.env.E2E_VERGUNNINGSOFTWARE_BV_MANAGEMENT_PASSWORD ||
        "development",
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

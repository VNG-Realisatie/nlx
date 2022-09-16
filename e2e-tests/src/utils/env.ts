/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

export interface Env {
  directoryUrl: string;
}

export const env: Env = {
  directoryUrl: process.env.E2E_DIRECTORY_URL || "http://127.0.0.1:7905",
};

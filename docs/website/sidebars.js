/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

module.exports = {
  docs: {
    "Understanding the basics": [
      "understanding-the-basics/introduction",
      "understanding-the-basics/security",
      "understanding-the-basics/product-vision",
    ],
    "Try NLX": [
      "try-nlx/introduction",
      "try-nlx/setup-your-environment",
      "try-nlx/retrieve-a-demo-certificate",
      {
        "With NLX Management (new)": [
          "try-nlx/management/introduction",
          "try-nlx/management/getting-up-and-running",
          "try-nlx/management/consume-an-api",
          "try-nlx/management/provide-an-api",
        ],
      },
      {
        "With config file": [
          "try-nlx/config-file/consume-an-api",
          "try-nlx/config-file/provide-an-api",
        ],
      },
    ],
    "Use NLX": [
      "use-nlx/request-a-production-cert",
      "use-nlx/enable-transaction-logs",
      "use-nlx/enable-finance",
      "use-nlx/setup-authorization",
      "use-nlx/new-releases",
    ],
    "Reference information": [
      "reference-information/service-configuration",
      "reference-information/transaction-log-headers",
      "reference-information/monitoring",
      "reference-information/outway-as-proxy",
    ],
    Support: ["support/contact", "support/common-errors"],
    Compliancy: [
      "compliancy/eif",
      "compliancy/eidas",
      "compliancy/accessibility",
      "compliancy/gdpr",
    ],
  },
};

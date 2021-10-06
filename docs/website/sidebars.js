/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

module.exports = {
  docs: [
    {
      type: 'category',
      label: 'Understanding the basics',
      items: [
        'understanding-the-basics/introduction',
        'understanding-the-basics/security',
        'understanding-the-basics/product-vision',
      ],
    },
    {
      type: 'category',
      label: 'Try NLX',
      collapsible: true,
      collapsed: true,
      items: [
        {
          type: 'category',
          label: 'Docker compose',
          collapsible: true,
          collapsed: true,
          items: [
            'try-nlx/docker/introduction',
            'try-nlx/docker/setup-your-environment',
            'try-nlx/docker/retrieve-a-demo-certificate',
            'try-nlx/docker/getting-up-and-running',
            'try-nlx/docker/provide-an-api',
          ],
        },
        {
          type: 'category',
          label: 'Helm',
          collapsible: true,
          collapsed: true,
          items: [
            'try-nlx/helm/introduction',
            'try-nlx/helm/0preparation',
            'try-nlx/helm/1createnamespace',
            'try-nlx/helm/2createcertificate',
            'try-nlx/helm/4nlxmanagement',
            'try-nlx/helm/5nlxinway',
            'try-nlx/helm/6nlxoutway',
            'try-nlx/helm/7sampleapi',
            'try-nlx/helm/8accessapi',
            'try-nlx/helm/finish',
          ],
        },
      ],
    },
    {
      type: 'category',
      label: 'NLX in production',
      collapsible: true,
      collapsed: true,
      items: [
        'use-nlx/introduction',
        'use-nlx/request-a-production-cert',
        'use-nlx/enable-transaction-logs',
        'use-nlx/enable-pricing',
        'use-nlx/setup-authorization',
        'use-nlx/new-releases',
      ],
    },
    {
      type: 'category',
      label: 'Reference information',
      collapsible: true,
      collapsed: true,
      items: [
        'reference-information/transaction-log-headers',
        'reference-information/monitoring',
        'reference-information/outway-as-proxy',
        'reference-information/environment-variables',
        'reference-information/ip-addresses',
        'reference-information/oidc',
      ],
    },
    {
      Support: ['support/common-errors'],
      Compliancy: [
        'compliancy/eif',
        'compliancy/eidas',
        'compliancy/accessibility',
        'compliancy/gdpr',
      ],
    },
  ],
};

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
          label: 'Docker Compose',
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
            'try-nlx/helm/preparation',
            'try-nlx/helm/create-namespace',
            'try-nlx/helm/create-certificate',
            'try-nlx/helm/postgresql',
            'try-nlx/helm/nlx-management',
            'try-nlx/helm/nlx-inway',
            'try-nlx/helm/nlx-outway',
            'try-nlx/helm/sample-api',
            'try-nlx/helm/access-api',
            'try-nlx/helm/transaction-log',
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
        'nlx-in-production/introduction',
        'nlx-in-production/request-a-production-cert',
        'nlx-in-production/setup-authorization',
        'nlx-in-production/new-releases',
      ],
    },
    {
      type: 'category',
      label: 'Reference information',
      collapsible: true,
      collapsed: true,
      items: [
        'reference-information/organization-identification',
        'reference-information/data-validation',
        'reference-information/transaction-log',
        'reference-information/transaction-log-headers',
        'reference-information/enable-pricing',
        'reference-information/monitoring',
        'reference-information/outway-as-proxy',
        'reference-information/environment-variables',
        'reference-information/ip-addresses',
        'reference-information/oidc',
        'reference-information/rewrite-base-url',
        'reference-information/add-headers-with-proxy',
        'reference-information/tls-between-inway-and-service',
        'reference-information/user-management',
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

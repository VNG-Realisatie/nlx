/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

module.exports = {
  title: 'Documentation',
  tagline: '',
  url: 'https://docs.nlx.io',
  baseUrl: '/',
  favicon: 'img/favicon.ico',
  organizationName: 'common-ground',
  projectName: 'nlx-docs',
  themeConfig: {
    navbar: {
      title: 'Documentation',
      logo: {
        alt: 'NLX logo',
        src: 'img/logo-light.svg',
        srcDark: 'img/logo.svg',
      },
      items: [
        {to: 'understanding-the-basics/introduction', label: 'Docs'},
        {href: 'https://directory.demo.nlx.io/', label: 'Directory'},
        {href: 'https://nlx.io', label: 'NLX'},
        {href: 'https://nlx.io/contact', label: 'Contact'},
      ],
    },
    footer: {
      style: 'light',
      links: [],
      logo: {
        alt: 'VNG Realisatie Logo',
        src: 'img/logo_vng_gs.svg',
      },
      copyright: `Copyright © ${new Date().getFullYear()} VNG Realisatie`,
    },
    algolia: {
      appId: 'OZCKSG4LL8',
      apiKey: 'e0b23250ffedb285810a41e0f2616eb0',
      indexName: 'docs-nlx',
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarCollapsible: false,
          path: '../docs',
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://gitlab.com/commonground/nlx/nlx/tree/master/docs/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  customFields: {
    startUrl: 'understanding-the-basics/introduction',
  },
};

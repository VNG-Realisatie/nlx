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
    sidebarCollapsible: false,
    navbar: {
      title: 'Documentation',
      logo: {
        alt: 'NLX logo',
        src: 'img/logo.svg',
      },
      items: [
        {href: 'https://nlx.io/about/', label: 'Over NLX', position: 'right'},
        {to: 'understanding-the-basics/introduction', label: 'Docs', position: 'right'},
        {href: 'https://directory.demo.nlx.io/', label: 'Directory', position: 'right'},
        {to: 'support/contact', label: 'Support', position: 'right'},
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
      apiKey: 'f5d0c017e70ffe180e05515a2869c1e4',
      indexName: 'nlx',
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
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

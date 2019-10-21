/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// See https://docusaurus.io/docs/site-config for all the possible
// site configuration options.

const siteConfig = {
  title: 'Documentation', // Title for your website.
  tagline: '',

  url: '', // Your website URL
  baseUrl: '/', // Base URL for your project */
  docsUrl: '',

  // For github.io type URLs, you would set the url and baseUrl like:
  //   url: 'https://facebook.github.io',
  //   baseUrl: '/test-site/',

  // Used for publishing and more
  projectName: 'nlx-docs',
  organizationName: 'common-ground',

  // For no header links in the top nav bar -> headerLinks: [],
  headerLinks: [
    {href: 'https://nlx.io/about/', label: 'Over NLX'},
    {doc: 'introduction', label: 'Docs'},
    {href: 'https://directory.demo.nlx.io/', label: 'Directory'},
    {doc: 'support', label: 'Support'},
  ],

  /* path to images for header/footer */
  headerIcon: 'img/logo.svg',
  favicon: 'img/favicon.ico',

  /* Colors for website */
  colors: {
    primaryColor: '#3d83fa',
    secondaryColor: '#FEBF24',
  },

  // This copyright info is used in /core/Footer.js and blog RSS/Atom feeds.
  copyright: `Copyright © ${new Date().getFullYear()} VNG Realisatie`,

  highlight: {
    // Highlight.js theme to use for syntax highlighting in code blocks.
    theme: 'atom-one-dark',
  },

  // Add custom scripts here that would be placed in <script> tags.
  scripts: [
    'https://buttons.github.io/buttons.js',
    'https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.0/clipboard.min.js',
    '/js/code-block-buttons.js',
  ],

  stylesheets: ['/css/code-block-buttons.css'],

  // On page navigation for the current documentation page.
  onPageNav: 'separate',
  // No .html extensions for paths.
  cleanUrl: true,

  // For sites with a sizable amount of content, set collapsible to true.
  // Expand/collapse the links and subcategories under categories.
  // docsSideNavCollapsible: true,

  // Show documentation's last contributor's name.
  // enableUpdateBy: true,

  // Show documentation's last update time.
  // enableUpdateTime: true,

  // You may provide arbitrary config keys to be used as needed by your
  // template. For example, if you need your repo's URL...
  //   repoUrl: 'https://github.com/facebook/test-site',

  editUrl: 'https://gitlab.com/commonground/nlx/nlx/tree/master/docs/docs/'
};

module.exports = siteConfig;

# Common UI

This repository contains common UI components for the NLX project.

## Install and use the components

You can import components from common UI one by one. Install the `common-ui` package by running:

```bash
$ npm install --save @nlxio/common-ui
```

Then you should be able to import components like this:

```js
import React from "react";
import ReactDOM from "react-dom";
import { Button } from "@nlxio/common-ui";

ReactDOM.render(<Button>Hello world!</Button>, document.getElementById("root"));
```

## Contributing

We use [Storybook](https://storybook.js.org/) to develop and test new components. To get Storybook up and running, use:

```bash
$ npm install
$ npm run storybook
```

The `package.json` defines a number of scripts that are used to build the project:

- `npm run build`: Builds all the JavaScript files using Babel.
- `npm run test`: Lints and tests all the JavaScript files.

## Building and publishing

To publish a new version of the library, first login to the NPM registry:

```bash
$ npm login
```

Then build a new version by running:

```bash
$ npm run build
```

You can bump the version in the `package.json` and `package-lock.json` by running:

```bash
$ npm version [ patch | minor | major ]
```

Finally publish a new version using:

```bash
$ npm publish
```

For more information about publishing NPM packages, consult the [official documentation](https://docs.npmjs.com/getting-started/publishing-npm-packages).

## Link common-ui package on your local machine

To use an unpublished version of `common-ui` in other projects on your local machine, you can use the `npm link` command.

1. First run the npm link command in the common-ui folder:

```bash
$ npm link
```

2. Then switch to the project where you would like to use the unpublished version of the `common-ui` project and run:

```bash
$ npm link @nlxio/common-ui
```

This last command created a symlink in the `node_modules` of the project to the local `common-ui` folder.

## Internationalization (i18n)

Common-ui components should be build language neutral and language independent. All labels and language specific strings need to be provided to common-ui component as a property by parent component. These parent components are project specific and not part of the common-ui library.

In November 2018 frontend team decided that NLX front-end applications should use [i18next](https://www.i18next.com/) library for internationalization.

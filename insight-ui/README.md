# NLX/Insight-ui

The insight-ui provides a human-readable view of users IRMA attributes exchanged between the government instances during the user interaction with the government.
The communication between government instances is performed over secured NLX network layer.

## Development

### NPM scripts

```bash
  # run using local backend (proxy to http://directory.dev.nlx.minikube:30080)
  # see NLX (root) readme how to setup local backend environment
  npm start
  # run in dev mode (proxy to https://directory.test.nlx.io/)
  npm run dev
  # run in acc mode (proxy to https://directory.acc.nlx.io/)
  npm run acc
  # run in demo mode (proxy to https://directory.demo.nlx.io/)
  npm run demo
  # run in production mode (api point fixed to https://directory.demo.nlx.io/)
  npm run prod
  # build production (api point fixed to https://directory.demo.nlx.io/)
  npm run build
  # build using dev mode
  npm run build:dev
  # run tests with jest (once)
  npm test
  # run tests in watch mode (TDD)
  npm run tdd
  # lint all js
  npm run lint
  # run webpack in debug mode
  npm run wp:dev
```

### Proxies

The proxies are defined in webpack/proxy.js. During development each request starting with `/api` will be proxied to environment specific proxy. The proxy server is based on --env property passed to webpack-dev-server. See NPM scripts (or package.json) for dev scripts available (start, dev, demo & acc). See root NLX readme file for available NLX environments.

In production mode the directory api point is set to static value https://directory.demo.nlx.io/api/directory/list-organizations

### Redux

All redux files are in `src/store` folder. The actions are dispatched by pages and by middleware (see picture). First action is dispatched from index.js. You can open action model (docs/redux-insight-ui.xml) using draw.io

![redux-actions](docs/redux-insight-ui.jpg)

### Project structure (folders)

- **dist:** builds are placed in this folder.
- **docker:** nginx configuration file for docker build
- **docs:** developers documentation. Contains redux action model used in this readme file
- **src:** app source code (react app)
  - `components`: generic components shared between pages
  - `layout`: main page layout component
  - `page`: page components. These usually load one or more components from components folder
  - `store`: redux store, actions, middleware and reducers
  - `styles`: CSS in JSS style definitions for all components
  - `utils`: utility functions for app and testing
- **static:** static files, like images, logos etc.
- **webpack:** webpack configuration modules. Main file (webpack.config.js) use these modules.

### Debugging webpack

Webpack can be debugged with chrome dev tools. Webpack is node.js app and debugging is similar to other node.js apps.

```bash
  # set debugging point in the script
  debugger

  # start webpack in inspect mode
  node --inspect-brk ./node_modules/webpack/bin/webpack.js
  # OR use available NPM script
  npm run wp:dev

  # open chrome inspect tools at chrome://inspect/
  # use ctl + c (twice!) to stop debugging

```

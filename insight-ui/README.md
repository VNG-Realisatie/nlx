# insight-ui

The insight-ui provides a human-readable view of information exchanged between the organizations. Insight-ui uses IRMA to identify the owner of a data-subject. The communication between organizations is performed over secure NLX network.

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

  # build production (api point fixed to https://directory.demo.nlx.io/)
  npm run build

  # run tests
  npm test
```

### Proxies

The proxies are configured in `src/setupProxy.js`.
During development each request starting with `/api` will be proxied to environment specific proxy.

The proxy server is based on the `PROXY` environment variable property passed to the start script. See NPM scripts (or package.json) for all available scripts.

### Redux

All redux files are in `src/store` folder. The actions are dispatched by pages and by middleware (see picture). First action is dispatched from index.js. You can open action model (docs/redux-insight-ui.xml) using draw.io

![redux-actions](docs/redux-insight-ui.jpg)

### Project structure (folders)

- **build:** builds are placed in this folder.
- **docker:** nginx configuration file for docker build
- **docs:** developers documentation. Contains redux action model used in this readme file
- **public:** static files, like images, logos etc.
- **src:** app source code (react app)
  - `components`: generic components shared between pages
  - `layout`: main page layout component
  - `page`: page components. These usually load one or more components from components folder
  - `store`: redux store, actions, middleware and reducers
  - `styles`: CSS in JSS style definitions for all components
  - `utils`: utility functions for app and testing

### Testing locally with IRMA app

When testing IRMA using a phone on the same WiFi network as your host machine you must setup a port-forward directly to the application you want to expose.

```bash
kubectl --namespace nlx-dev-rdw port-forward deployment/irma-api-server 2222:8080
socat tcp-listen:3333,fork tcp:127.0.0.1:2222
```

You can now let your phone connect to the IRMA api server of RDW on `your.host.machine.ip:3333`. See comment in src/store/middleware/mwIrma.js (line 72) for detailed instruction.

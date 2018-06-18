import React from 'react';
import ReactDOM from 'react-dom';

import './static/css/base.css';
import 'bootstrap/dist/js/bootstrap.js';

import App from './App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();

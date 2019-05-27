// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { applyMiddleware, createStore } from 'redux'
import createSagaMiddleware from 'redux-saga'
import { composeWithDevTools } from 'redux-devtools-extension'
import reducers from './reducers'
import rootSaga from './sagas'

const sagaMiddleware = createSagaMiddleware()

const store = createStore(
    reducers,
    composeWithDevTools(
      applyMiddleware(sagaMiddleware)
    )
)

sagaMiddleware.run(rootSaga)

export default store

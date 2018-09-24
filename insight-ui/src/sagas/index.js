import { all } from 'redux-saga/effects';
import { divaSaga } from 'diva-react';

import appSaga from './app-saga';
import sessionSaga from './session-saga';

export default function* rootSaga() {
  yield all([
    appSaga(),
    sessionSaga(),
    divaSaga(),
  ]);
}
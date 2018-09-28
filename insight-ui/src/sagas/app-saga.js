import { put, all, takeEvery } from 'redux-saga/effects';

import { types as divaTypes } from 'diva-react';
import { actions as sessionActions } from '../reducers/session-reducer';

export function* onDivaSessionCompleted(action) {
  if (action.serverStatus === 'DONE') {
    yield put(sessionActions.getSessionData());
  }
}

function* sagas() {
  yield all([
    takeEvery(divaTypes.SESSION_COMPLETED, onDivaSessionCompleted),
  ]);
}

export default sagas;
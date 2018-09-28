import { put, call, all, takeEvery } from 'redux-saga/effects';

import { types, actions } from '../reducers/session-reducer';
import service from '../services/session-service';

function* getSessionData(baseUrl) {
  try {
    const response = yield call(service.getSessionData, baseUrl);
    if (response.sessionId && response.attributes) {
      yield put(actions.sessionDataReceived(response.sessionId, response.attributes));
    } else {
      yield put(actions.getSessionDataFailed('Server Error', response));
    }
  } catch (error) {
    yield put(actions.getSessionDataFailed('Network Error', error.response));
  }
}

function* deauthenticate(baseUrl) {
  try {
    const response = yield call(service.deauthenticate, baseUrl);
    if (response) {
      yield put(actions.getSessionData());
    } else {
      yield put(actions.getSessionDataFailed('Server Error', response));
    }
  } catch (error) {
    yield put(actions.getSessionDataFailed('Network Error', error.response));
  }
}

function* sagas(baseUrl = '/api') {
  yield all([
    takeEvery(types.GET_DATA, getSessionData, baseUrl),
    takeEvery(types.DEAUTHENTICATE, deauthenticate, baseUrl),
  ]);
}

export default sagas;
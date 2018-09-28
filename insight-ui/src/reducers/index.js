import { combineReducers } from 'redux';
import { divaReducer } from 'diva-react';

import session from './session-reducer';

const rootReducer = combineReducers({
  session,
  divaReducer,
});

export default rootReducer;

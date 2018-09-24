export const types = {
    GET_DATA: 'SESSION/GET_DATA',
    GET_DATA_SUCCESS: 'SESSION/GET_DATA_SUCCESS',
    GET_DATA_FAILED: 'SESSION/GET_DATA_FAILED',
    DEAUTHENTICATE: 'SESSION/DEAUTHENTICATE',
  };
  
  export const initialState = {
    isFetching: false,
    sessionId: null,
    attributes: {},
  };
  
  export const actions = {
    getSessionData: () => ({
      type: types.GET_DATA,
    }),
    sessionDataReceived: (sessionId, attributes) => ({
      type: types.GET_DATA_SUCCESS,
      sessionId,
      attributes,
      receivedAt: Date.now(),
    }),
    getSessionDataFailed: (reason, response) => ({
      type: types.GET_DATA_FAILED,
      reason,
      response,
    }),
    deauthenticate: () => ({
      type: types.DEAUTHENTICATE,
    }),
  };
  
  export default (state = initialState, action) => {
    switch (action.type) {
      case types.GET_DATA:
        return {
          ...state,
          isFetching: true,
        };
      case types.GET_DATA_SUCCESS:
        return {
          ...state,
          isFetching: false,
          lastUpdated: action.receivedAt,
          sessionId: action.sessionId,
          attributes: action.attributes,
        };
      case types.GET_DATA_FAILED:
        return {
          ...state,
          isFetching: false,
          error: {
            reason: action.reason,
            response: action.response,
          },
        };
      default:
        return state;
    }
  };
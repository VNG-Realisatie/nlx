const createStore = (middleware, state = {}) => {
    const store = {
        getState: jest.fn(() => state),
        dispatch: jest.fn(),
    }
    const next = jest.fn()

    const invoke = (action) => middleware(store)(next)(action)

    return { store, next, invoke }
}

export { createStore }

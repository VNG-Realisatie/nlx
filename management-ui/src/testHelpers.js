// helper function, because Jest does not take care of chained promises.
// see https://github.com/facebook/jest/issues/2157#issuecomment-279171856
export const flushPromises = () =>
    new Promise((resolve) => setImmediate(resolve))

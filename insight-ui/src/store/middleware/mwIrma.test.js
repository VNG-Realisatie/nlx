import cfg from '../app.cfg'
import * as actionType from '../actions'

import mwIrma from './mwIrma'
import { createStore } from '../../utils/testing/reduxMock'

describe('mwIrma', () => {
    it(`call next on any action`, () => {
        const { next, invoke } = createStore(mwIrma)
        const action = { type: 'TEST_ACTION' }
        invoke(action)
        expect(next).toHaveBeenCalledWith(action)
    })

    it(`dispatch ${actionType.GET_ORGANIZATION_LOGS} on ${
        actionType.IRMA_GET_PROOF_OK
    } with payload`, () => {
        const action = {
            type: actionType.IRMA_GET_PROOF_OK,
            payload: 'TEST_PAYLOAD',
        }
        const state = {
            organization: {
                ...cfg.organization,
            },
        }
        const { invoke, store } = createStore(mwIrma, state)
        invoke(action)
        expect(store.dispatch).toHaveBeenCalledWith(
            expect.objectContaining({
                type: actionType.GET_ORGANIZATION_LOGS,
            }),
        )
    })
})

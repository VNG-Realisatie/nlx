// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import * as actionType from '../actions'

import mwOrganizations from './mwOrganizations'
import { createStore } from '../../utils/testing/reduxMock'

describe('mwOrganizations', () => {
    it(`call next on any action`, () => {
        const { next, invoke } = createStore(mwOrganizations)
        const action = { type: 'TEST_ACTION' }
        invoke(action)
        expect(next).toHaveBeenCalledWith(action)
    })

    it(`dispatch ${actionType.HIDE_LOADER} on ${
        actionType.GET_IRMA_ORGANIZATIONS_OK
    } with payload`, () => {
        const action = {
            type: actionType.GET_IRMA_ORGANIZATIONS_OK,
            payload: 'TEST_PAYLOAD',
        }
        const { invoke, store } = createStore(mwOrganizations)
        invoke(action)
        expect(store.dispatch).toHaveBeenCalledWith(
            expect.objectContaining({
                type: actionType.HIDE_LOADER,
            }),
        )
    })
})

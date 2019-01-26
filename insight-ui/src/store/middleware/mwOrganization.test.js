import * as actionType from '../actions'

import mwOrganization from './mwOrganization'
import { createStore } from '../../utils/testing/reduxMock'

describe('mwOrganization', () => {
    it(`call next on any action`, () => {
        const { next, invoke } = createStore(mwOrganization)
        const action = { type: 'TEST' }
        invoke(action)
        expect(next).toHaveBeenCalledWith(action)
    })

    it(`dispatch ${actionType.GET_QRCODE} on ${
        actionType.SELECT_ORGANIZATION
    } with same payload`, () => {
        const action = {
            type: actionType.SELECT_ORGANIZATION,
            payload: 'SELECT_ORGANIZATION_PAYLOAD',
        }
        const mwAction = {
            type: actionType.GET_QRCODE,
            payload: 'SELECT_ORGANIZATION_PAYLOAD',
        }
        const { invoke, store } = createStore(mwOrganization)
        invoke(action)
        expect(store.dispatch).toHaveBeenCalledWith(mwAction)
    })

    it(`dispatch ${actionType.GET_ORGANIZATION_LOGS_ERR} if ${
        actionType.GET_ORGANIZATION_LOGS
    } payload is MISSING`, () => {
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS,
        }
        const { invoke, store } = createStore(mwOrganization)
        invoke(action)
        expect(store.dispatch).toHaveBeenCalledWith(
            expect.objectContaining({
                type: actionType.GET_ORGANIZATION_LOGS_ERR,
            }),
        )
    })

    it(`dispatch ${actionType.GET_ORGANIZATION_LOGS_ERR} if ${
        actionType.GET_ORGANIZATION_LOGS
    } jwt is MISSING in payload`, () => {
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS,
            payload: {
                api: 'https://my-api',
            },
        }
        const { invoke, store } = createStore(mwOrganization)
        invoke(action)
        expect(store.dispatch).toHaveBeenCalledWith(
            expect.objectContaining({
                type: actionType.GET_ORGANIZATION_LOGS_ERR,
            }),
        )
    })
})

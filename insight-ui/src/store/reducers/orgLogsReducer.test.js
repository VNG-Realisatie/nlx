import cfg from '../app.cfg'
import * as actionType from '../actions'

import orgLogsReducer from './orgLogsReducer'

describe('orgLogsReducer', () => {
    it('returns inital state', () => {
        const state = orgLogsReducer(undefined, {})
        expect(state).toBe(cfg.organization.logs)
    })

    it("doesn't alter state on RANDOM action", () => {
        const initState = { prop: 'DONT_TOUCH_THIS' }
        const state = orgLogsReducer(initState, {
            type: actionType.GET_LANGUAGE,
        })
        expect(state).toBe(initState)
    })

    it('updates error object in the state', () => {
        const initState = { prop: 'DONT_TOUCH_THIS' }
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS_ERR,
            payload: {
                id: 404,
                msg: 'Page not found',
            },
        }
        const state = orgLogsReducer(initState, action)
        expect(state.error).toEqual(action.payload)
        // keep prop 'untouched'
        expect(state.prop).toEqual(initState.prop)
    })

    it(`resets items to [] on ${actionType.GET_ORGANIZATION_LOGS_ERR}`, () => {
        const initState = {
            prop: 'DONT_TOUCH_THIS',
            items: ['item1', 'item2', 'item3'],
        }
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS_ERR,
            payload: {
                id: 404,
                msg: 'Page not found',
            },
        }
        const state = orgLogsReducer(initState, action)
        // reset items array
        expect(state.items).toEqual([])
        // keep prop 'untouched'
        expect(state.prop).toEqual(initState.prop)
    })

    it(`adds payload to state on ${
        actionType.GET_ORGANIZATION_LOGS_OK
    }`, () => {
        const initState = {
            prop: 'DONT_TOUCH_THIS',
            jwt: 'OVERWRITE_OLD_BASE64_TOKEN',
            items: ['item1', 'item2', 'item3'],
        }
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS_OK,
            payload: {
                name: 'organization',
                api: 'https://myapi/fetch',
                jwt: 'BASE64ENCODEDJWT',
                items: [
                    { rec: true },
                    { rec: true },
                    { rec: false },
                    { rec: false },
                ],
                page: 0,
                rowCount: 10,
                rowsPerPage: 5,
            },
        }
        const state = orgLogsReducer(initState, action)
        // add new prop
        expect(state.name).toEqual(action.payload.name)
        // overwrite existing prop
        expect(state.jwt).toEqual(action.payload.jwt)
        expect(state.items).toEqual(action.payload.items)
        // keep prop 'untouched'
        expect(state.prop).toEqual(initState.prop)
    })

    it(`combines pageDef props on ${
        actionType.GET_ORGANIZATION_LOGS_OK
    }`, () => {
        const initState = cfg.organization.logs
        const action = {
            type: actionType.GET_ORGANIZATION_LOGS_OK,
            payload: {
                name: 'organization',
                api: 'https://myapi/fetch',
                jwt: 'BASE64ENCODEDJWT',
                items: [{ rec: true }],
                page: 1,
                rowCount: 10,
            },
        }

        const state = orgLogsReducer(initState, action)

        expect(state.pageDef.page).toEqual(action.payload.page)
        expect(state.pageDef.rowCount).toEqual(action.payload.rowCount)
        expect(state.pageDef.rowsPerPage).toEqual(initState.pageDef.rowsPerPage)
        expect(state.pageDef.rowsPerPageOptions).toEqual(
            initState.pageDef.rowsPerPageOptions,
        )
    })

    it(`resets to initial state on ${actionType.RESET_ORGANIZATION}`, () => {
        const initState = cfg.organization.logs
        const action = {
            type: actionType.RESET_ORGANIZATION,
        }

        const state = orgLogsReducer({ data: 'RANDOM_VALUE_IN_STATE' }, action)

        expect(state).toEqual(initState)
    })
})

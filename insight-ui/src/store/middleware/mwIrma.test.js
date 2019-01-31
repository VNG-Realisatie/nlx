import cfg from '../app.cfg'
import * as actionType from '../actions'

// import axios from 'axios'
// import MockAdapter from 'axios-mock-adapter'

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

    // it(`should NOT modify insight_irma_endpoint received from api (comment line 82)`, () => {
    //     const action = {
    //         type: actionType.GET_QRCODE,
    //         payload: {
    //             name: 'haarlem',
    //             insight_irma_endpoint: 'https://irma',
    //             insight_log_endpoint: 'https://insight-api.nl',
    //         },
    //     }

    //     const response = {
    //         type: 'GET_QRCODE_OK',
    //         payload: {
    //             name: 'haarlem',
    //             dataSubjects: {
    //                 burgerservicenummer: {
    //                     label: 'Burgerservicenummer',
    //                 },
    //                 kenteken: {
    //                     label: 'Kenteken',
    //                 },
    //             },
    //             qrCode:
    //                 '{"u":"https://irma/api/v2/verification/2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0","v":"2.0","vmax":"2.3","irmaqr":"disclosing"}',
    //             statusUrl:
    //                 'https://irma/api/v2/verification/2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0/status',
    //             proofUrl:
    //                 'https://irma/api/v2/verification/2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0/getproof',
    //             firstJWT: '2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0',
    //             inProgress: false,
    //         },
    //     }

    //     beforeAll(() => {
    //         let mock = new MockAdapter(axios)
    //         mock.onGet(
    //             `${action.payload.insight_log_endpoint}/getDataSubject`,
    //         ).reply(200, {
    //             dataSubjects: {
    //                 burgerservicenummer: {
    //                     label: 'Burgerservicenummer',
    //                 },
    //                 kenteken: {
    //                     label: 'Kenteken',
    //                 },
    //             },
    //         })

    //         mock.onPost(
    //             `${action.payload.insight_log_endpoint}/generateJWT`,
    //         ).reply(200, {
    //             data: '2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0',
    //         })

    //         mock.onPost(
    //             `${action.payload.insight_irma_endpoint}/api/v2/verification/`,
    //         ).reply(200, {
    //             u:
    //                 'https://irma/api/v2/verification/2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0',
    //             v: '2.0',
    //             vmax: '2.3',
    //             irmaqr: 'disclosing',
    //         })
    //     })

    //     const { invoke, store } = createStore(mwIrma)
    //     invoke(action)
    //     expect(store.dispatch).toHaveBeenCalledWith(response)
    // })
})

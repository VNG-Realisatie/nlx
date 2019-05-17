// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import cfg from '../app.cfg'
import * as actionType from '../actions'

import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'

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

    it(`does NOT alter insight_irma_endpoint => COMMENT app.cfg.localIp`, () => {
        const action = {
            type: actionType.GET_QRCODE,
            payload: {
                name: 'haarlem',
                insight_irma_endpoint: '',
                insight_log_endpoint: '',
            },
        }

        const qrCode =
            '{"u":"/api/v2/verification/","v":"2.0","vmax":"2.3","irmaqr":"disclosing"}'

        let mock = new MockAdapter(axios)

        mock.onGet()
            .reply(200, {
                dataSubjects: {
                    burgerservicenummer: {
                        label: 'Burgerservicenummer',
                    },
                    kenteken: {
                        label: 'Kenteken',
                    },
                },
            })
            .onPost('/generateJWT')
            .reply(200, '2ikuT1rt1CpMUQf8SmszjiDvxVcesV6Hl5QAstMycLz0', {
                'content-type': 'text/plain; charset=utf-8',
            })
            .onPost('/api/v2/verification/')
            .reply(200, {
                u: '',
                v: '2.0',
                vmax: '2.3',
                irmaqr: 'disclosing',
            })

        const { invoke, store } = createStore(mwIrma)

        return invoke(action).then(() => {
            expect(store.dispatch).toHaveBeenCalledWith(
                expect.objectContaining({
                    payload: expect.objectContaining({
                        qrCode,
                    }),
                }),
            )
        })
    })
})

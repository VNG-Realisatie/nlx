import React from 'react'
import { shallow } from 'enzyme'
import { MemoryRouter } from 'react-router-dom'
import { IrmaAuthPage } from './IrmaAuthPage'
import cfg from '../../store/app.cfg'

const testProps = {
    loading: true,
    info: cfg.organization.info,
    irma: cfg.organization.irma,
    qrCode: cfg.organization.irma.qrCode,
    loginInProgress: cfg.organization.irma.inProgress,
    jwt: cfg.organization.irma.jwt,
    error: cfg.organization.irma.error,
}

let component

beforeAll(() => {
    component = shallow(
        <MemoryRouter>
            <IrmaAuthPage {...testProps} />
        </MemoryRouter>,
    )
})

it('renders IrmaAuthPage component', () => {
    expect(component).toBeTruthy()
})

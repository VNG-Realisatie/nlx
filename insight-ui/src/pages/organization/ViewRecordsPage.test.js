import React from 'react'
import { shallow } from 'enzyme'

import { ViewRecordsPage } from './ViewRecordsPage'
import cfg from '../../store/app.cfg'

const testProps = {
    loading: true,
    colDef: cfg.organization.logs.colDef,
    logs: [],
    pageDef: cfg.organization.logs.pageDef,
    error: cfg.organization.logs.error,
    name: cfg.organization.logs.name,
    jwt: cfg.organization.logs.jwt,
    api: cfg.organization.logs.api,
}

let component

beforeAll(() => {
    component = shallow(<ViewRecordsPage {...testProps} />)
})

it('renders ViewRecordsPage component', () => {
    expect(component).toBeTruthy()
})

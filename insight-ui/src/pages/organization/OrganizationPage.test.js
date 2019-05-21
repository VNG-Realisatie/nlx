// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import { MemoryRouter } from 'react-router-dom'
import { OrganizationPage } from './OrganizationPage'
import cfg from '../../store/app.cfg'

const testProps = {
    loading: true,
    organization: cfg.organization,
    organizations: cfg.organizations.list,
}

let component

beforeAll(() => {
    component = shallow(
        <MemoryRouter>
            <OrganizationPage {...testProps} />
        </MemoryRouter>,
    )
    component.organization = {
        name: 'test',
    }
})

it('renders OrganizationPage component', () => {
    expect(component).toBeTruthy()
})

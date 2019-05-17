// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'

import OrganizationList from './OrganizationList'

const testProps = {
    organizations: [
        {
            name: 'demo-organization',
        },
        {
            name: 'brp',
            insight_irma_endpoint: 'irma-api.dev.brp.minikube',
            insight_log_endpoint: 'txlog-api.dev.brp.minikube',
        },
        {
            name: 'rdw',
            insight_irma_endpoint: 'irma-api.dev.rdw.minikube',
            insight_log_endpoint: 'txlog-api.dev.rdw.minikube',
        },
    ],
}

let component

beforeAll(() => {
    component = shallow(<OrganizationList {...testProps} />)
})

it('renders OrganizationList component with testProps', () => {
    expect(component).toBeTruthy()
})

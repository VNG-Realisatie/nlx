// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'

import LogModal from './LogModal'

const testProps = {
    data: {
        source_organization: 'denhaag',
        destination_organization: 'rdw',
        service_name: 'kentekenregister',
        'logrecord-id': '8c0oi0nngdnig',
        data: {
            'doelbinding-application-id': 'Parkeervergunningapplicatie',
            'doelbinding-data-elements': 'kenteken,burgerservicenummer',
            'doelbinding-process-id': 'Aanvragen van parkeervergunning',
            'doelbinding-subject-identifier': '123456987',
            'request-path': '/voertuigen',
        },
        created: '2018-09-19T09:31:03.413301Z',
    },
    open: false,
    closeModal: () => {
        return 'close modal'
    },
}

let component

beforeAll(() => {
    component = shallow(<LogModal {...testProps} />)
})

describe('<LogModal />', () => {
    it('creates LogModal component', () => {
        expect(component).toBeTruthy()
    })

    it('received source_organization from (test) props', () => {
        const props = component.props()
        expect(props.data['source_organization']).toEqual(
            testProps.data['source_organization'],
        )
    })
})

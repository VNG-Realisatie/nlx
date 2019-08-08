// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import DocumentationPage from './DocumentationPage'
import Documentation from "../../components/Documentation/Documentation";

describe('DocumentationPage', () => {
  let documentationEl

  beforeEach(() => {
    const wrapper = shallow(<DocumentationPage match={ { params: { serviceName: 'test-service', organizationName: 'test-organization' }}} />)
    documentationEl = wrapper.find(Documentation)
  })

  it('should show the Documentation component', () => {
    expect(documentationEl.exists()).toEqual(true)
  })

  it('should pass the service and organization name to the documentation component', () => {
    expect(documentationEl.prop('serviceName')).toEqual('test-service')
    expect(documentationEl.prop('organizationName')).toEqual('test-organization')
  })
})

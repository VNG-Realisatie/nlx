// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import Documentation from './Documentation'

it('renders without crashing', () => {
  expect(() => {
    shallow(<Documentation serviceName="test-service" organizationName="test-organization" />)
  }).not.toThrow()
})

describe('service is invalid / not specified', () => {
  it('should throw an exception', () => {
    expect(() => {
      shallow(<Documentation organizationName="test-organization" />)
    }).toThrow()
  })
})

describe('organization is invalid / not specified', () => {
  it('should throw an exception', () => {
    expect(() => {
      shallow(<Documentation serviceName="test-service" />)
    }).toThrow()
  })
})

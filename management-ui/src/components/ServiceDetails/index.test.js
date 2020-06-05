// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import ServiceDetails from './index'

const service = {
  name: 'name',
  internal: false,
  authorizationSettings: { mode: 'none' },
  inways: [],
}

describe('ServiceDetails', () => {
  beforeEach(() => {
    jest.useFakeTimers()
  })

  it('should display', () => {
    const { getByTestId, queryByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails service={service} />
      </Router>,
    )

    expect(getByTestId('service-published')).toHaveTextContent(
      'visible.svg' + 'Published in central directory', // eslint-disable-line no-useless-concat
    )
    expect(getByTestId('service-inways')).toHaveTextContent(
      'inway.svg' + 'Inways' + '0', // eslint-disable-line no-useless-concat
    )
    expect(queryByTestId('service-authorizations')).toBeNull()
  })

  it('should show hidden icon', () => {
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails
          service={{
            ...service,
            internal: true,
          }}
        />
      </Router>,
    )
    expect(getByTestId('service-published')).toHaveTextContent(
      'hidden.svg' + 'Not visible in central directory', // eslint-disable-line no-useless-concat
    )
  })

  it('should show an inway', async () => {
    const { getByTestId, queryByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails
          service={{
            ...service,
            inways: ['inway 1'],
          }}
        />
      </Router>,
    )
    expect(getByTestId('service-inways')).toHaveTextContent(
      'inway.svg' + 'Inways' + '1', // eslint-disable-line no-useless-concat
    )
    expect(queryByTestId('service-inways-list')).toBeNull()

    fireEvent.click(getByTestId('service-inways'))
    jest.runAllTimers()
    expect(getByTestId('service-inways-list')).toBeTruthy()
    expect(getByTestId('service-inway-0')).toHaveTextContent('inway 1')
  })

  it('should show a block for an empty list of inways', async () => {
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails service={service} />
      </Router>,
    )
    fireEvent.click(getByTestId('service-inways'))

    expect(getByTestId('service-no-inways')).toBeTruthy()
  })

  it('should show a block for an empty whitelist', async () => {
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails
          service={{
            ...service,
            authorizationSettings: {
              mode: 'whitelist',
              authorizations: [],
            },
          }}
        />
      </Router>,
    )
    expect(getByTestId('service-authorizations')).toHaveTextContent(
      'whitelist.svg' + 'Whitelisted organizations' + '0', // eslint-disable-line no-useless-concat
    )
    fireEvent.click(getByTestId('service-authorizations'))

    expect(getByTestId('service-no-authorizations')).toBeTruthy()
  })

  it('should show a whitelist', async () => {
    const { getByTestId, queryByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails
          service={{
            ...service,
            authorizationSettings: {
              mode: 'whitelist',
              authorizations: [{ organizationName: 'Organisatie A' }],
            },
          }}
        />
      </Router>,
    )
    expect(queryByTestId('service-authorizations-list')).toBeNull()
    expect(getByTestId('service-authorizations')).toHaveTextContent(
      'whitelist.svg' + 'Whitelisted organizations' + '1', // eslint-disable-line no-useless-concat
    )

    fireEvent.click(getByTestId('service-authorizations'))
    jest.runAllTimers()
    expect(getByTestId('service-authorizations-list')).toBeTruthy()
    expect(getByTestId('service-authorization-0')).toHaveTextContent(
      'Organisatie A',
    )
  })

  it('should call the removeHandler on remove', () => {
    const handleRemove = jest.fn()
    jest.spyOn(window, 'confirm').mockResolvedValue(true)
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetails
          removeHandler={handleRemove}
          service={{
            ...service,
            authorizationSettings: {
              mode: 'whitelist',
              authorizations: [{ organizationName: 'Organisatie A' }],
            },
          }}
        />
      </Router>,
    )

    fireEvent.click(getByTestId('remove-service'))
    expect(window.confirm).toHaveBeenCalled()
    expect(handleRemove).toBeCalled()
    expect(getByTestId('remove-success')).toBeTruthy()
  })
})

// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import ServiceDetailView from './index'

const service = {
  name: 'name',
  internal: false,
  inways: [],
}

describe('ServiceDetails', () => {
  it('should display', () => {
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={jest.fn()} />
      </Router>,
    )

    expect(getByTestId('service-published')).toHaveTextContent(
      'visible.svg' + 'Published in central directory', // eslint-disable-line no-useless-concat
    )

    expect(getByTestId('service-inways')).toHaveTextContent(
      'inway.svg' + 'Inways' + '0', // eslint-disable-line no-useless-concat
    )
  })

  it('should show hidden icon', () => {
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetailView
          service={{
            ...service,
            internal: true,
          }}
          removeHandler={jest.fn()}
        />
      </Router>,
    )
    expect(getByTestId('service-published')).toHaveTextContent(
      'hidden.svg' + 'Not visible in central directory', // eslint-disable-line no-useless-concat
    )
  })

  it('should call the removeHandler on remove', () => {
    const handleRemove = jest.fn()
    jest.spyOn(window, 'confirm').mockResolvedValue(true)
    const { getByTestId } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={handleRemove} />
      </Router>,
    )

    fireEvent.click(getByTestId('remove-service'))
    expect(window.confirm).toHaveBeenCalled()
    expect(handleRemove).toBeCalled()
  })
})

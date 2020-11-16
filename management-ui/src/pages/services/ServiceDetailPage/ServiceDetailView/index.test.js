// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import ServiceDetailView from './index'

// These all have their own tests
jest.mock('./InwaysSection', () => () => <div />)
jest.mock('./AccessRequestsSection', () => () => <div />)
jest.mock('./AccessGrantSection', () => () => <div />)

const service = {
  name: 'name',
  internal: false,
  inways: [],
}

describe('ServiceDetails', () => {
  it('should display', () => {
    const { queryByText } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={jest.fn()} />
      </Router>,
    )

    expect(queryByText('Published in central directory')).toBeInTheDocument()
  })

  it('should show hidden icon', () => {
    const { queryByText } = renderWithProviders(
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
    expect(queryByText('Not visible in central directory')).toBeInTheDocument()
  })

  it('should call the removeHandler on remove', () => {
    const handleRemove = jest.fn()
    jest.spyOn(window, 'confirm').mockResolvedValue(true)
    const { getByTitle } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={handleRemove} />
      </Router>,
    )

    fireEvent.click(getByTitle('Remove service'))
    expect(window.confirm).toHaveBeenCalled()
    expect(handleRemove).toBeCalled()
  })
})

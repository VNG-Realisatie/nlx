// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { act, fireEvent, render } from '@testing-library/react'
import { AUTHORIZATION_TYPE_WHITELIST } from '../../../vocabulary'
import { renderWithProviders } from '../../../test-utils'
import AddServiceForm from './index'

describe('with initial values', () => {
  it('should pre-fill the form fields with the initial values', () => {
    const { getByRole, getByLabelText } = renderWithProviders(
      <AddServiceForm
        initialValues={{
          // values copied from config-api/config/service.json
          name: 'my-service',
          endpointURL: 'my-service.test:8000',
          documentationURL: 'my-service.test:8000/docs',
          apiSpecificationURL: 'my-service.test:8000/openapi.json',
          internal: false,
          techSupportContact: 'tech@organization.test',
          publicSupportContact: 'public@organization.test',
          authorizationMode: AUTHORIZATION_TYPE_WHITELIST,
        }}
      />,
    )

    expect(getByLabelText('API naam').value).toBe('my-service')
    expect(getByLabelText('API endpoint URL').value).toBe(
      'my-service.test:8000',
    )
    expect(getByLabelText('API documentatie URL').value).toBe(
      'my-service.test:8000/docs',
    )
    expect(getByLabelText('API specificatie URL').value).toBe(
      'my-service.test:8000/openapi.json',
    )

    expect(getByLabelText('Publiceren in de centrale directory').value).toBe(
      'false',
    )

    expect(getByLabelText('Tech support email').value).toBe(
      'tech@organization.test',
    )
    expect(getByLabelText('Public support email').value).toBe(
      'public@organization.test',
    )

    expect(
      getByLabelText('Whitelist voor geauthorizeerde organisaties').checked,
    ).toBe(true)
    expect(getByLabelText('Alle organisaties toestaan').checked).toBe(false)

    expect(getByRole('button').textContent).toBe('Service toevoegen')
  })

  it('should allow configuring the submit button text', () => {
    const { getByRole } = render(<AddServiceForm submitButtonText="Opslaan" />)

    expect(getByRole('button').textContent).toBe('Opslaan')
  })
})

test('the form values of the onSubmitHandler', async () => {
  const onSubmitHandlerSpy = jest.fn()

  const { container, getByTestId, findByTestId, getByLabelText } = render(
    <AddServiceForm
      onSubmitHandler={onSubmitHandlerSpy}
      initialValues={{
        // values copied from config-api/config/service.json
        name: '',
        endpointURL: 'my-service.test:8000',
        documentationURL: 'my-service.test:8000/docs',
        apiSpecificationURL: 'my-service.test:8000/openapi.json',
        internal: false,
        techSupportContact: 'tech@organization.test',
        publicSupportContact: 'public@organization.test',
        authorizationMode: AUTHORIZATION_TYPE_WHITELIST,
      }}
    />,
  )

  // invalid form - name is missing
  const formElement = getByTestId('form')
  await act(async () => {
    fireEvent.submit(formElement)
  })

  // assert the validation feedback is shown
  const nameError = await findByTestId('name-error')
  expect(nameError).not.toBeNull()

  // fill-in required fields
  const nameField = getByLabelText('API naam')
  fireEvent.change(nameField, {
    target: { value: 'my-service' },
  })

  // re-submit the valid form
  await act(async () => {
    fireEvent.submit(formElement)
  })

  expect(
    container.querySelectorAll('p[class*="FieldValidationMessage"'),
  ).toHaveLength(0)

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    name: 'my-service',
    endpointURL: 'my-service.test:8000',
    documentationURL: 'my-service.test:8000/docs',
    apiSpecificationURL: 'my-service.test:8000/openapi.json',
    internal: false,
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    authorizationMode: AUTHORIZATION_TYPE_WHITELIST,
  })
})

// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { AUTHORIZATION_TYPE_WHITELIST } from '../../vocabulary'
import { renderWithProviders, act, fireEvent, waitFor } from '../../test-utils'
import ServiceForm from './index'

jest.mock('../FormikFocusError', () => () => <></>)

describe('with initial values', () => {
  it('should pre-fill the form fields with the initial values', async () => {
    const { getByLabelText, findByLabelText } = renderWithProviders(
      <ServiceForm
        initialValues={{
          // values copied from management-api/config/service.json
          name: 'my-service',
          endpointURL: 'my-service.test:8000',
          documentationURL: 'my-service.test:8000/docs',
          apiSpecificationURL: 'my-service.test:8000/openapi.json',
          internal: false,
          techSupportContact: 'tech@organization.test',
          publicSupportContact: 'public@organization.test',
          authorizationSettings: {
            mode: AUTHORIZATION_TYPE_WHITELIST,
          },
          inways: ['inway1'],
        }}
        submitButtonText="Submit"
        getInways={() => [{ name: 'inway1' }, { name: 'inway2' }]}
      />,
    )

    expect(getByLabelText('Service name').value).toBe('my-service')
    expect(getByLabelText('API endpoint URL').value).toBe(
      'my-service.test:8000',
    )
    expect(getByLabelText('API documentation URL').value).toBe(
      'my-service.test:8000/docs',
    )
    expect(getByLabelText('API specification URL').value).toBe(
      'my-service.test:8000/openapi.json',
    )

    expect(getByLabelText('Publish to central directory').value).toBe('true')

    expect(getByLabelText('Tech support email').value).toBe(
      'tech@organization.test',
    )
    expect(getByLabelText('Public support email').value).toBe(
      'public@organization.test',
    )

    expect(
      getByLabelText('Whitelist for authorized organizations').checked,
    ).toBe(true)
    expect(getByLabelText('Allow all organizations').checked).toBe(false)

    expect(await findByLabelText('inway1')).toHaveAttribute('checked')
    expect(await findByLabelText('inway2')).not.toHaveAttribute('checked')
  })

  it('should allow configuring the submit button text', async () => {
    const { findByRole } = renderWithProviders(
      <ServiceForm submitButtonText="Opslaan" allInways={() => []} />,
    )
    expect(await findByRole('button')).toHaveTextContent('Opslaan')
  })
})

test('the form values of the onSubmitHandler', async () => {
  const onSubmitHandlerSpy = jest.fn()

  const {
    container,
    getByTestId,
    findByTestId,
    getByLabelText,
  } = renderWithProviders(
    <ServiceForm
      submitButtonText="Submit"
      onSubmitHandler={onSubmitHandlerSpy}
      initialValues={{
        // values copied from management-api/config/service.json
        name: '',
        endpointURL: 'my-service.test:8000',
        documentationURL: 'my-service.test:8000/docs',
        apiSpecificationURL: 'my-service.test:8000/openapi.json',
        internal: false,
        techSupportContact: 'tech@organization.test',
        publicSupportContact: 'public@organization.test',
        authorizationSettings: {
          mode: AUTHORIZATION_TYPE_WHITELIST,
        },
      }}
      getInways={() => []}
    />,
  )

  // invalid form - name is missing
  const formElement = getByTestId('form')
  await act(async () => {
    fireEvent.submit(formElement)
  })

  // assert the validation feedback is shown
  const nameError = await findByTestId('error-name')
  expect(nameError).not.toBeNull()

  // fill-in required fields
  const nameField = getByLabelText('Service name')
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

  waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      name: 'my-service',
      endpointURL: 'my-service.test:8000',
      documentationURL: 'my-service.test:8000/docs',
      apiSpecificationURL: 'my-service.test:8000/openapi.json',
      internal: false,
      inways: [],
      techSupportContact: 'tech@organization.test',
      publicSupportContact: 'public@organization.test',
      authorizationSettings: {
        mode: AUTHORIZATION_TYPE_WHITELIST,
      },
    }),
  )
})

describe('when showing inways', () => {
  const initialValues = {
    // values copied from config-api/config/service.json
    name: 'my-service',
    endpointURL: 'my-service.test:8000',
    documentationURL: 'my-service.test:8000/docs',
    apiSpecificationURL: 'my-service.test:8000/openapi.json',
    internal: false,
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    authorizationSettings: {
      mode: AUTHORIZATION_TYPE_WHITELIST,
    },
    inways: [],
  }

  it('should show a warning when there are no inways registered', async () => {
    const { findByTestId } = renderWithProviders(
      <ServiceForm
        initialValues={{ ...initialValues, inways: [] }}
        submitButtonText="Submit"
        getInways={() => []}
      />,
    )
    expect(await findByTestId('inways-empty')).toBeTruthy()
    expect(
      await findByTestId('publishedInDirectory-warning'),
    ).toHaveTextContent('Service not yet accessible')
  })
  it('should show a warning when the service is published and no inways are selected', async () => {
    const { findByLabelText, findByTestId } = renderWithProviders(
      <ServiceForm
        initialValues={{ ...initialValues, inways: [], internal: false }}
        submitButtonText="Submit"
        getInways={() => [{ name: 'inway-one' }]}
      />,
    )
    expect(await findByLabelText('inway-one')).not.toHaveAttribute('checked')
    expect(
      await findByTestId('publishedInDirectory-warning'),
    ).toHaveTextContent('Service not yet accessible')
  })

  it('should not show a warning when the service is private and no inways are selected', async () => {
    const { findByLabelText, queryByTestId } = renderWithProviders(
      <ServiceForm
        initialValues={{ ...initialValues, inways: [], internal: true }}
        submitButtonText="Submit"
        getInways={() => [{ name: 'inway-one' }]}
      />,
    )
    expect(await findByLabelText('inway-one')).not.toHaveAttribute('checked')
    expect(queryByTestId('publishedInDirectory-warning')).toBeFalsy()
  })

  it('should save an inway selection', async () => {
    const onSubmitHandlerSpy = jest.fn()
    const {
      findByLabelText,
      getByLabelText,
      getByTestId,
    } = renderWithProviders(
      <ServiceForm
        onSubmitHandler={onSubmitHandlerSpy}
        initialValues={{ ...initialValues, inways: [] }}
        submitButtonText="Submit"
        getInways={() => [{ name: 'inway-one' }]}
      />,
    )

    await findByLabelText('inway-one')

    await fireEvent.click(getByLabelText('inway-one'))
    expect(getByLabelText('inway-one').checked).toEqual(true)

    await fireEvent.submit(getByTestId('form'))

    await waitFor(() => {
      expect(onSubmitHandlerSpy).toHaveBeenCalledTimes(1)
      expect(onSubmitHandlerSpy).toHaveBeenCalledWith(
        expect.objectContaining({ inways: ['inway-one'] }),
      )
    })
  })

  it('should be able to remove an inway from the selection', async () => {
    const onSubmitHandlerSpy = jest.fn()

    const {
      findByLabelText,
      findByTestId,
      getByLabelText,
      getByTestId,
    } = renderWithProviders(
      <ServiceForm
        onSubmitHandler={onSubmitHandlerSpy}
        initialValues={{ ...initialValues, inways: ['inway-one'] }}
        submitButtonText="Submit"
        getInways={() => [{ name: 'inway-one' }, { name: 'inway-two' }]}
      />,
    )

    await findByLabelText('inway-one')

    await fireEvent.click(getByLabelText('inway-one'))
    expect(getByLabelText('inway-one').checked).toEqual(false)
    expect(await findByTestId('publishedInDirectory-warning')).toBeTruthy()

    await fireEvent.submit(getByTestId('form'))

    await waitFor(() => {
      expect(onSubmitHandlerSpy).toHaveBeenCalledTimes(1)
      expect(onSubmitHandlerSpy).toBeCalledWith(
        expect.objectContaining({
          inways: [],
        }),
      )
    })
  })
})

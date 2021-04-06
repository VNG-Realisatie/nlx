// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { fireEvent, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { renderWithProviders } from '../../test-utils'
import { RootStore, StoreProvider } from '../../stores'
import { ManagementApi } from '../../api'
import ServiceForm, { checkStep } from './index'

jest.mock('../FormikFocusError', () => () => <></>)

describe('checkStep yup test', () => {
  it('returns expected results', () => {
    expect(checkStep(0.01, 0.03)).toBe(true)
    expect(checkStep(0.01, 0.0025)).toBe(false)
    expect(checkStep(0.25, 5.5)).toBe(true)
    expect(checkStep(5, 2)).toBe(false)
    expect(checkStep(5, 5)).toBe(true)
  })
})

describe('with initial values', () => {
  it('should pre-fill the form fields with the initial values', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest.fn().mockResolvedValue({
      inways: [{ name: 'inway1' }, { name: 'inway2' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { getByLabelText, findByLabelText } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          initialValues={{
            name: 'my-service',
            endpointURL: 'my-service.test:8000',
            documentationURL: 'my-service.test:8000/docs',
            apiSpecificationURL: 'my-service.test:8000/openapi.json',
            internal: false,
            techSupportContact: 'tech@organization.test',
            publicSupportContact: 'public@organization.test',
            inways: ['inway1'],
            oneTimeCosts: 0,
            monthlyCosts: 0,
            requestCosts: 0,
            isPaidService: false,
          }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
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

    expect(await findByLabelText('inway1')).toHaveAttribute('checked')
    expect(await findByLabelText('inway2')).not.toHaveAttribute('checked')
  })

  it('should allow configuring the submit button text', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { findByRole } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm submitButtonText="Opslaan" />
      </StoreProvider>,
    )
    expect(await findByRole('button')).toHaveTextContent('Opslaan')
  })
})

test('the form values of the onSubmitHandler', async () => {
  const onSubmitHandlerSpy = jest.fn()

  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest
    .fn()
    .mockResolvedValue({ inways: [] })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const { findByTestId, getByLabelText, getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <ServiceForm
        submitButtonText="Submit"
        onSubmitHandler={onSubmitHandlerSpy}
        initialValues={{
          name: '',
          endpointURL: 'my-service.test:8000',
          documentationURL: 'my-service.test:8000/docs',
          apiSpecificationURL: 'my-service.test:8000/openapi.json',
          internal: false,
          techSupportContact: 'tech@organization.test',
          publicSupportContact: 'public@organization.test',
          oneTimeCosts: 0,
          monthlyCosts: 0,
          requestCosts: 0,
          isPaidService: false,
        }}
      />
    </StoreProvider>,
  )

  // invalid form - name is missing
  userEvent.click(getByText('Submit'))

  const nameError = await findByTestId('error-name')
  expect(nameError).not.toBeNull()
  expect(onSubmitHandlerSpy).not.toHaveBeenCalled()

  // fill-in required fields
  userEvent.type(getByLabelText('Service name'), 'my-service')
  userEvent.click(getByText('Submit'))

  await waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      name: 'my-service',
      endpointURL: 'my-service.test:8000',
      documentationURL: 'my-service.test:8000/docs',
      apiSpecificationURL: 'my-service.test:8000/openapi.json',
      internal: false,
      inways: [],
      techSupportContact: 'tech@organization.test',
      publicSupportContact: 'public@organization.test',
      oneTimeCosts: 0,
      monthlyCosts: 0,
      requestCosts: 0,
      isPaidService: false,
    }),
  )
})

describe('when showing inways', () => {
  const initialValues = {
    name: 'my-service',
    endpointURL: 'my-service.test:8000',
    documentationURL: 'my-service.test:8000/docs',
    apiSpecificationURL: 'my-service.test:8000/openapi.json',
    internal: false,
    techSupportContact: 'tech@organization.test',
    publicSupportContact: 'public@organization.test',
    inways: [],
    oneTimeCosts: 0,
    monthlyCosts: 0,
    requestCosts: 0,
    isPaidService: false,
  }

  it('should show a warning when there are no inways registered', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { findByTestId } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          initialValues={{ ...initialValues, inways: [] }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )
    expect(await findByTestId('inways-empty')).toBeTruthy()
    expect(
      await findByTestId('publishedInDirectory-warning'),
    ).toHaveTextContent('Service not yet accessible')
  })

  it('should show a warning when the service is published and no inways are selected', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { findByLabelText, findByTestId } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          initialValues={{ ...initialValues, inways: [], internal: false }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )
    expect(await findByLabelText('inway-one')).not.toHaveAttribute('checked')
    expect(
      await findByTestId('publishedInDirectory-warning'),
    ).toHaveTextContent('Service not yet accessible')
  })

  it('should not show a warning when the service is private and no inways are selected', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { findByLabelText, queryByTestId } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          initialValues={{ ...initialValues, inways: [], internal: true }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )
    expect(await findByLabelText('inway-one')).not.toHaveAttribute('checked')
    expect(queryByTestId('publishedInDirectory-warning')).toBeFalsy()
  })

  it('should save an inway selection', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const onSubmitHandlerSpy = jest.fn()
    const {
      findByLabelText,
      getByLabelText,
      getByTestId,
    } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          onSubmitHandler={onSubmitHandlerSpy}
          initialValues={{ ...initialValues, inways: [] }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
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
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest.fn().mockResolvedValue({
      inways: [{ name: 'inway-one' }, { name: 'inway-two' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const onSubmitHandlerSpy = jest.fn()

    const {
      findByLabelText,
      findByTestId,
      getByLabelText,
      getByTestId,
    } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          onSubmitHandler={onSubmitHandlerSpy}
          initialValues={{ ...initialValues, inways: ['inway-one'] }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
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

  it('should clear costs when disabling finance', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const onSubmitHandlerSpy = jest.fn()
    const {
      findByLabelText,
      getByLabelText,
      getByTestId,
    } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          onSubmitHandler={onSubmitHandlerSpy}
          initialValues={{
            ...initialValues,
            oneTimeCosts: 10.5,
            monthlyCosts: 5,
            requestCosts: 1.25,
          }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )

    await findByLabelText('This is a paid service')

    await fireEvent.click(getByLabelText('This is a paid service'))

    await fireEvent.submit(getByTestId('form'))

    await waitFor(() => {
      expect(onSubmitHandlerSpy).toHaveBeenCalledTimes(1)
      expect(onSubmitHandlerSpy).toHaveBeenCalledWith(
        expect.objectContaining({
          oneTimeCosts: 0,
          monthlyCosts: 0,
          requestCosts: 0,
        }),
      )
    })
  })

  it('should save costs when finance was already enabled', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const onSubmitHandlerSpy = jest.fn()
    const { getByTestId } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          onSubmitHandler={onSubmitHandlerSpy}
          initialValues={{
            ...initialValues,
            inways: [],
            oneTimeCosts: 10.5,
            monthlyCosts: 5,
            requestCosts: 1.25,
          }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )

    await fireEvent.submit(getByTestId('form'))

    await waitFor(() => {
      expect(onSubmitHandlerSpy).toHaveBeenCalledTimes(1)
      expect(onSubmitHandlerSpy).toHaveBeenCalledWith(
        expect.objectContaining({
          oneTimeCosts: 10.5,
          monthlyCosts: 5,
          requestCosts: 1.25,
        }),
      )
    })
  })

  it('should save costs when finance is enabled', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListInways = jest
      .fn()
      .mockResolvedValue({ inways: [{ name: 'inway-one' }] })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const onSubmitHandlerSpy = jest.fn()
    const { getByLabelText, getByTestId } = renderWithProviders(
      <StoreProvider rootStore={rootStore}>
        <ServiceForm
          onSubmitHandler={onSubmitHandlerSpy}
          initialValues={{ ...initialValues }}
          submitButtonText="Submit"
        />
      </StoreProvider>,
    )

    await fireEvent.click(getByLabelText('This is a paid service'))

    const oneTime = getByLabelText('One time costs (in Euro)')
    const monthly = getByLabelText('Monthly costs (in Euro)')
    const request = getByLabelText('Cost per request (in Euro)')

    userEvent.clear(oneTime)
    userEvent.type(oneTime, '10.5')
    userEvent.clear(monthly)
    userEvent.type(monthly, '5')
    userEvent.clear(request)
    userEvent.type(request, '1.25')

    await fireEvent.submit(getByTestId('form'))

    await waitFor(() => {
      expect(onSubmitHandlerSpy).toHaveBeenCalledTimes(1)
      expect(onSubmitHandlerSpy).toHaveBeenCalledWith(
        expect.objectContaining({
          oneTimeCosts: 10.5,
          monthlyCosts: 5,
          requestCosts: 1.25,
        }),
      )
    })
  })
})

// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { act, fireEvent, render } from '@testing-library/react'
import * as Yup from 'yup'
import { ErrorMessage, Formik, Field } from 'formik'
import FormikFocusError from './index'

test('after submitting the form, the viewport scrolls to the first invalid input element', async () => {
  const scrollIntoView = jest.fn()
  window.HTMLElement.prototype.scrollIntoView = scrollIntoView

  const onSubmit = () => {}
  const validationSchema = Yup.object().shape({
    key: Yup.string().required('This field is required.'),
  })
  const { getByTestId } = render(
    <Formik
      initialValues={{ key: '' }}
      validationSchema={validationSchema}
      onSubmit={onSubmit}
    >
      {({ handleSubmit }) => (
        <form data-testid="form" onSubmit={handleSubmit}>
          <label>
            Key
            <Field>
              {({ field }) => (
                <input
                  name="key"
                  data-testid="key-input"
                  type="text"
                  {...field}
                />
              )}
            </Field>
            <ErrorMessage name="key">
              {(msg) => (
                <p data-testid="key-error" name="key">
                  {msg}
                </p>
              )}
            </ErrorMessage>
          </label>
          <FormikFocusError />
        </form>
      )}
    </Formik>,
  )

  await act(async () => {
    fireEvent.submit(getByTestId('form'))
  })

  const keyError = getByTestId('key-error')
  expect(keyError).toHaveTextContent('This field is required.')

  expect(scrollIntoView).toHaveBeenCalled()
})

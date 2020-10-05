// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, shape, string } from 'prop-types'
import { Field, Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import { Button, Fieldset, Label, Legend } from '@commonground/design-system'
import FormikFocusError from '../../../../components/FormikFocusError'
import usePromise from '../../../../hooks/use-promise'
import InwayRepository from '../../../../domain/inway-repository'
import {
  StyledForm,
  InwaysEmptyMessage,
  InwaysLoadingMessage,
} from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  organizationInway: '',
}

const Form = ({ initialValues, onSubmitHandler, getInways, ...props }) => {
  const { t } = useTranslation()
  const { isReady: inwaysIsReady, result: inways } = usePromise(getInways)

  const validationSchema = Yup.object().shape({
    organizationInway: Yup.string(),
  })

  return (
    <Formik
      initialValues={{
        ...DEFAULT_INITIAL_VALUES,
        ...initialValues,
      }}
      validationSchema={validationSchema}
      onSubmit={(values) => onSubmitHandler(values)}
    >
      {({ handleSubmit }) => (
        <StyledForm onSubmit={handleSubmit} data-testid="form" {...props}>
          <Fieldset>
            <Legend>{t('General settings')}</Legend>

            <Label htmlFor="organizationInway">{t('Organization inway')}</Label>

            <p>
              {t(
                'This inway is used to be able to retrieve & confirm access requests from other organizations.',
              )}
            </p>

            {!inwaysIsReady ? (
              <InwaysLoadingMessage />
            ) : !inways || inways.length === 0 ? (
              <InwaysEmptyMessage data-testid="no-inways-available">
                {t('There are no inways available.')}
              </InwaysEmptyMessage>
            ) : (
              <Field
                id="organizationInway"
                name="organizationInway"
                data-testid="organizationInway"
                as="select"
              >
                <option value="">{t('None')}</option>
                {inways.map((inway) => (
                  <option value={inway.name} key={inway.name}>
                    {inway.name}
                  </option>
                ))}
              </Field>
            )}
          </Fieldset>

          <Button type="submit">{t('Save settings')}</Button>

          <FormikFocusError />
        </StyledForm>
      )}
    </Formik>
  )
}

Form.propTypes = {
  onSubmitHandler: func,
  initialValues: shape({
    organizationInway: string,
  }),
  getInways: func,
}

Form.defaultProps = {
  onSubmitHandler: () => {},
  initialValues: DEFAULT_INITIAL_VALUES,
  getInways: InwayRepository.getAll,
}

export default Form

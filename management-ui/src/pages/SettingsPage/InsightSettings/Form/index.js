// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, shape, string } from 'prop-types'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import {
  Button,
  Fieldset,
  TextInput,
  Legend,
} from '@commonground/design-system'
import FormikFocusError from '../../../../components/FormikFocusError'
import { StyledForm } from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  irmaServerURL: '',
  insightAPIURL: '',
}

const Form = ({ initialValues, onSubmitHandler, ...props }) => {
  const { t } = useTranslation()

  const validationSchema = Yup.object().shape({
    irmaServerURL: Yup.string(),
    insightAPIURL: Yup.string(),
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
            <Legend>{t('Insight settings')}</Legend>

            <TextInput name="irmaServerURL" size="xl">
              {t('IRMA server URL')}
            </TextInput>

            <TextInput name="insightAPIURL" size="xl">
              {t('Insight API URL')}
            </TextInput>
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
    irmaServerURL: string,
    insightAPIURL: string,
  }),
}

Form.defaultProps = {
  onSubmitHandler: () => {},
  initialValues: DEFAULT_INITIAL_VALUES,
}

export default Form

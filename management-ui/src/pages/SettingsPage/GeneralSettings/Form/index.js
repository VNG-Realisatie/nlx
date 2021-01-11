// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, shape, string } from 'prop-types'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import { Button, Fieldset, Legend, Select } from '@commonground/design-system'

import FormikFocusError from '../../../../components/FormikFocusError'
import { useModalConfirm } from '../../../../components/ModalConfirm'
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
  const { isReady: inwaysIsReady, result: inways } = usePromise(getInways)
  const { t } = useTranslation()
  const [ModalConfirm, confirm] = useModalConfirm({
    okText: t('Save'),
    children: t(
      'By removing the organization inway it is no longer possible to process or receive access requests',
    ),
  })

  const validateOrganizationInwayAndSubmit = async (values) => {
    if (values.organizationInway) {
      onSubmitHandler(values)
      return
    }

    if (await confirm()) {
      onSubmitHandler(values)
    }
  }

  const validationSchema = Yup.object().shape({
    organizationInway: Yup.string(),
  })

  const selectInwayOptions = inways
    ? inways.map((inway) => ({
        value: inway.name,
        label: inway.name,
      }))
    : []

  const emptyOption = { value: '', label: t('None') }
  selectInwayOptions.unshift(emptyOption)

  return (
    <>
      <Formik
        initialValues={{
          ...DEFAULT_INITIAL_VALUES,
          ...initialValues,
        }}
        validationSchema={validationSchema}
        onSubmit={validateOrganizationInwayAndSubmit}
      >
        {({ handleSubmit }) => (
          <StyledForm onSubmit={handleSubmit} data-testid="form" {...props}>
            <Fieldset>
              <Legend>{t('General settings')}</Legend>
              {!inwaysIsReady ? (
                <InwaysLoadingMessage />
              ) : !inways || inways.length === 0 ? (
                <InwaysEmptyMessage>
                  {t('There are no inways available')}
                </InwaysEmptyMessage>
              ) : (
                <>
                  <Select
                    id="organizationInway"
                    name="organizationInway"
                    options={selectInwayOptions}
                  >
                    {t('Organization inway')}
                  </Select>
                  <small>
                    {t(
                      'This inway is used to be able to retrieve & confirm access requests from other organizations.',
                    )}
                  </small>
                </>
              )}
            </Fieldset>

            <Fieldset>
              <Button type="submit">{t('Save settings')}</Button>
            </Fieldset>

            <FormikFocusError />
          </StyledForm>
        )}
      </Formik>

      <ModalConfirm />
    </>
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

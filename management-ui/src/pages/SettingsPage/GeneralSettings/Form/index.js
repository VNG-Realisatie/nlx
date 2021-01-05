// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { func, shape, string } from 'prop-types'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import { Button, Fieldset, Legend, Select } from '@commonground/design-system'

import FormikFocusError from '../../../../components/FormikFocusError'
import ModalConfirm from '../../../../components/ModalConfirm'
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
  const [formValues, setFormValues] = useState({})
  const [showConfirm, setShowConfirm] = useState(false)
  const [confirmSave, setConfirmSave] = useState(null)

  const { t } = useTranslation()
  const { isReady: inwaysIsReady, result: inways } = usePromise(getInways)

  useEffect(() => {
    switch (confirmSave) {
      case true:
        onSubmitHandler(formValues)
        break

      case false:
      default:
        break
    }
  }, [confirmSave, formValues]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleChoice = (isConfirmed) => {
    setConfirmSave(isConfirmed)
    setShowConfirm(false)
  }

  const validateOrganizationInwayAndSubmit = (values) => {
    if (values.organizationInway) {
      onSubmitHandler(values)
      return
    }

    setFormValues(values)
    setShowConfirm(true)
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

      <ModalConfirm
        isVisible={showConfirm}
        handleClose={() => {
          setShowConfirm(false)
        }}
        onChoice={handleChoice}
        okText={t('Save')}
      >
        <p>
          {t(
            'By removing the organization inway it is no longer possible to process or receive access requests',
          )}
          .
        </p>
      </ModalConfirm>
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

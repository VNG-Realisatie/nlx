// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useEffect } from 'react'
import { useFormikContext } from 'formik'

const FormikFocusError = () => {
  const { isSubmitting, isValidating, errors } = useFormikContext()

  useEffect(() => {
    if (isSubmitting && !isValidating && Object.keys(errors).length) {
      const key = Object.keys(errors)[0]
      const selector = `[name="${key}"]`
      const errorElement = document.querySelector(selector)

      if (errorElement) {
        const scrollElement =
          errorElement.labels &&
          errorElement.labels.length &&
          errorElement.labels[0].scrollIntoView instanceof Function
            ? errorElement.labels[0]
            : errorElement
        if (scrollElement) scrollElement.scrollIntoView({ behavior: 'smooth' })
        errorElement.focus()
      }
    }
  }, [isSubmitting, isValidating, errors])

  return null
}

export default FormikFocusError

// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { useTranslation } from 'react-i18next'
import {
  Button,
  FieldLabel,
  Fieldset,
  TextInput,
} from '@commonground/design-system'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { func } from 'prop-types'
import UserContext from '../../user-context'
import UserRepositoryBasicAuth from '../../domain/user-repository-basic-auth'
import {
  Content,
  StyledAlert,
  StyledForm,
  StyledNLXManagementLogo,
  StyledSidebar,
  Wrapper,
} from './index.styles'

const LoginBasicAuthPage = ({ loginHandler, storeCredentials, logout }) => {
  const { t } = useTranslation()
  const { user } = useContext(UserContext)
  const [error, setError] = useState(null)

  const validationSchema = Yup.object().shape({
    email: Yup.string()
      .email(t('Invalid email'))
      .required(t('This field is required')),
    password: Yup.string().required(t('This field is required')),
  })

  const initialValues = {
    email: '',
    password: '',
  }

  const handleSubmit = async ({ email, password }) => {
    setError(false)

    try {
      await loginHandler(email, password)
      storeCredentials(email, password)
      window.location = '/'
    } catch (error) {
      setError(true)
    }
  }

  const handleBasicAuthLogout = () => {
    logout()
    window.location = '/'
  }

  return (
    <Wrapper>
      <StyledSidebar>
        <StyledNLXManagementLogo />
      </StyledSidebar>
      <Content>
        <h1>{t('Welcome')}</h1>

        {!user ? (
          <>
            <p>{t('Log in to continue')}</p>

            {error && (
              <StyledAlert data-testid="login-error-message" variant="error">
                {t('Invalid credentials.')}
              </StyledAlert>
            )}

            <Formik
              initialValues={{
                ...initialValues,
              }}
              validationSchema={validationSchema}
              onSubmit={handleSubmit}
            >
              <StyledForm>
                <Fieldset>
                  <TextInput
                    field="what"
                    name="email"
                    id="email"
                    size="m"
                    type="email"
                    autoComplete="email"
                    autoFocus
                  >
                    <FieldLabel label={t('Email address')} />
                  </TextInput>
                  <TextInput
                    name="password"
                    id="current-password"
                    size="m"
                    type="password"
                    autoComplete="current-password"
                  >
                    <FieldLabel label={t('Password')} />
                  </TextInput>
                </Fieldset>

                <Button type="submit">{t('Log in')}</Button>
              </StyledForm>
            </Formik>
          </>
        ) : (
          <Button
            data-testid="logout"
            type="button"
            onClick={handleBasicAuthLogout}
          >
            {t('Log out')}
          </Button>
        )}
      </Content>
    </Wrapper>
  )
}

LoginBasicAuthPage.propTypes = {
  loginHandler: func,
  storeCredentials: func,
  logout: func,
}

LoginBasicAuthPage.defaultProps = {
  loginHandler: UserRepositoryBasicAuth.login,
  storeCredentials: UserRepositoryBasicAuth.storeCredentials,
  logout: UserRepositoryBasicAuth.logout,
}

export default LoginBasicAuthPage

// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { Form, Formik } from 'formik'
import { Button, ErrorMessage, Field, Fieldset, Label, Legend } from '../Form'
import { func } from 'prop-types'

import {
    StyledButtonGroup,
    StyledFormGroup,
    StyledFormGroupColumn,
    StyledFormGroupColumnContainer,
} from '../ServiceForm/index.styles'

const LoginForm = ({ login }) => (
    <Formik
        initialValues={{ username: '', password: '' }}
        validate={(values) => {
            const errors = {}
            if (!values.username) {
                errors.username = 'Required'
            }
            return errors
        }}
        onSubmit={login}
    >
        <Form>
            <Legend>Inloggen</Legend>
            <Fieldset>
                <StyledFormGroupColumnContainer>
                    <StyledFormGroupColumn>
                        <StyledFormGroup>
                            <Label htmlFor="username">Gebruikersnaam</Label>
                            <Field type="username" name="username" />
                            <ErrorMessage name="username" component="div" />
                        </StyledFormGroup>
                    </StyledFormGroupColumn>
                </StyledFormGroupColumnContainer>
                <StyledFormGroupColumnContainer>
                    <StyledFormGroupColumn>
                        <StyledFormGroup>
                            <Label htmlFor="password">Wachtwoord</Label>
                            <Field type="password" name="password" />
                            <ErrorMessage name="password" component="div" />
                        </StyledFormGroup>
                    </StyledFormGroupColumn>
                </StyledFormGroupColumnContainer>

                <StyledButtonGroup>
                    <Button type="submit">Inloggen</Button>
                </StyledButtonGroup>
            </Fieldset>
        </Form>
    </Formik>
)

LoginForm.propTypes = {
    login: func.isRequired,
}

export default LoginForm

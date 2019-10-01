// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { Component, Fragment } from 'react'
import { Formik, Form, FieldArray } from 'formik'
import { func, object } from 'prop-types'
import * as Yup from 'yup'

import {
    Label,
    Legend,
    HelperMessage,
    Field,
    Fieldset,
    Button,
    ErrorMessage,
} from '../Form'

import {
    StyledFormGroup,
    StyledButtonGroup,
    StyledFormGroupColumnContainer,
    StyledFormGroupColumn,
    StyledDeletableField,
} from './index.styles.js'

import ConfirmButton from '../../components/ConfirmButton'

const defaultInitialValues = {
    name: '',
    endpointURL: '',
    documentationURL: '',
    apiSpecificationURL: '',
    internal: false,
    techSupportContact: '',
    publicSupportContact: '',
    authorizationSettings: {
        mode: 'whitelist',
        organizations: [],
    },
    inways: [],
}

export const validationSchema = Yup.object().shape({
    name: Yup.string().required(),
    endpointURL: Yup.string().required(),
    documentationURL: Yup.string().url(),
    apiSpecificationURL: Yup.string(),
    internal: Yup.boolean(),
    techSupportContact: Yup.string().email(),
    publicSupportContact: Yup.string().email(),
    authorizationSettings: Yup.object().shape({
        mode: Yup.string().oneOf(['whitelist', 'public']),
        organizations: Yup.array().of(Yup.string().required()),
    }),
})

class ServiceForm extends Component {
    constructor(props) {
        super(props)

        this.state = {
            initialValues: {
                ...defaultInitialValues,
                ...this.props.initialValues,
                authorizationSettings: {
                    ...defaultInitialValues.authorizationSettings,
                    ...(this.props.initialValues
                        ? this.props.initialValues.authorizationSettings
                        : {}),
                },
            },
        }

        this.onSubmit = this.onSubmit.bind(this)
    }

    onSubmit(values) {
        const castedValues = validationSchema.cast(values)
        this.props.onSubmit(castedValues)
    }

    render() {
        const { onDelete, initialValues } = this.props

        return (
            <Formik
                initialValues={this.state.initialValues}
                validationSchema={validationSchema}
                onSubmit={this.onSubmit}
            >
                {({ handleSubmit, touched, values, errors }) => (
                    <Form onSubmit={handleSubmit}>
                        <Legend>API</Legend>
                        <Fieldset>
                            <StyledFormGroup>
                                <Label htmlFor="name">Name</Label>
                                <Field name="name" readOnly={initialValues} />
                                {errors.name && touched.name ? (
                                    <ErrorMessage>{errors.name}</ErrorMessage>
                                ) : null}
                            </StyledFormGroup>

                            <StyledFormGroup>
                                <Label htmlFor="endpointURL">
                                    Endpoint URL
                                </Label>
                                <Field name="endpointURL" />
                                {errors.endpointURL && touched.endpointURL ? (
                                    <ErrorMessage>
                                        {errors.endpointURL}
                                    </ErrorMessage>
                                ) : null}
                            </StyledFormGroup>

                            <StyledFormGroupColumnContainer>
                                <StyledFormGroupColumn>
                                    <StyledFormGroup>
                                        <Label htmlFor="documentationURL">
                                            Documentation URL
                                        </Label>
                                        <Field name="documentationURL" />
                                        {errors.documentationURL &&
                                        touched.documentationURL ? (
                                            <ErrorMessage>
                                                {errors.documentationURL}
                                            </ErrorMessage>
                                        ) : null}
                                    </StyledFormGroup>
                                </StyledFormGroupColumn>

                                <StyledFormGroupColumn>
                                    <StyledFormGroup>
                                        <Label htmlFor="apiSpecificationURL">
                                            API Specification URL
                                        </Label>
                                        <Field name="apiSpecificationURL" />
                                        {errors.apiSpecificationURL &&
                                        touched.apiSpecificationURL ? (
                                            <ErrorMessage>
                                                {errors.apiSpecificationURL}
                                            </ErrorMessage>
                                        ) : null}
                                    </StyledFormGroup>
                                </StyledFormGroupColumn>
                            </StyledFormGroupColumnContainer>

                            <StyledFormGroup>
                                <Label htmlFor="internal">Internal</Label>
                                <Field component="select" name="internal">
                                    <option value="false">No</option>
                                    <option value="true">Yes</option>
                                </Field>
                                <HelperMessage>
                                    Internal services are not published in the
                                    central directory
                                </HelperMessage>
                                {errors.internal && touched.internal ? (
                                    <ErrorMessage>
                                        {errors.internal}
                                    </ErrorMessage>
                                ) : null}
                            </StyledFormGroup>
                        </Fieldset>

                        <Legend>Contact</Legend>
                        <Fieldset>
                            <StyledFormGroupColumnContainer>
                                <StyledFormGroupColumn>
                                    <StyledFormGroup>
                                        <Label htmlFor="techSupportContact">
                                            Tech support e-mail
                                        </Label>
                                        <Field name="techSupportContact" />
                                        {errors.techSupportContact &&
                                        touched.techSupportContact ? (
                                            <ErrorMessage>
                                                {errors.techSupportContact}
                                            </ErrorMessage>
                                        ) : null}
                                    </StyledFormGroup>
                                </StyledFormGroupColumn>

                                <StyledFormGroupColumn>
                                    <StyledFormGroup>
                                        <Label htmlFor="publicSupportContact">
                                            Public support e-mail
                                        </Label>
                                        <Field name="publicSupportContact" />
                                        {errors.publicSupportContact &&
                                        touched.publicSupportContact ? (
                                            <ErrorMessage>
                                                {errors.publicSupportContact}
                                            </ErrorMessage>
                                        ) : null}
                                    </StyledFormGroup>
                                </StyledFormGroupColumn>
                            </StyledFormGroupColumnContainer>
                        </Fieldset>

                        <Legend>Authorization</Legend>
                        <Fieldset>
                            <StyledFormGroup>
                                <Label htmlFor="authorizationSettings[mode]">
                                    Mode
                                </Label>
                                <Field
                                    component="select"
                                    name="authorizationSettings[mode]"
                                >
                                    <option value="whitelist">Whitelist</option>
                                    <option value="public">Public</option>
                                </Field>
                                <HelperMessage>
                                    Create a whitelist for authorized
                                    organizations or allow all organizations.
                                </HelperMessage>
                                {errors.authorizationSettings &&
                                errors.authorizationSettings.mode &&
                                touched.authorizationSettings &&
                                touched.authorizationSettings.mode ? (
                                    <ErrorMessage>
                                        {errors.authorizationSettings.mode}
                                    </ErrorMessage>
                                ) : null}
                            </StyledFormGroup>

                            {values.authorizationSettings.mode ===
                                'whitelist' && (
                                <StyledFormGroup>
                                    <Label htmlFor="authorizationSettings[mode]">
                                        Organizations
                                    </Label>
                                    <FieldArray
                                        name="authorizationSettings[organizations]"
                                        render={(arrayHelpers) => (
                                            <Fragment>
                                                {values.authorizationSettings.organizations.map(
                                                    (org, index) => (
                                                        <StyledDeletableField
                                                            key={index}
                                                        >
                                                            <Field
                                                                name={`authorizationSettings[organizations].${index}`}
                                                            />
                                                            <Button
                                                                secondary
                                                                type="button"
                                                                onClick={() =>
                                                                    arrayHelpers.remove(
                                                                        index,
                                                                    )
                                                                }
                                                                aria-label={`Delete organization ${org}`}
                                                            >
                                                                -
                                                            </Button>
                                                        </StyledDeletableField>
                                                    ),
                                                )}
                                                <Button
                                                    secondary
                                                    type="button"
                                                    onClick={() =>
                                                        arrayHelpers.push('')
                                                    }
                                                >
                                                    + Add organization
                                                </Button>
                                            </Fragment>
                                        )}
                                    />
                                    {errors.authorizationSettings &&
                                    errors.authorizationSettings
                                        .organizations &&
                                    touched.authorizationSettings &&
                                    touched.authorizationSettings
                                        .organizations ? (
                                        <ErrorMessage>
                                            {
                                                errors.authorizationSettings
                                                    .organizations
                                            }
                                        </ErrorMessage>
                                    ) : null}
                                </StyledFormGroup>
                            )}
                        </Fieldset>

                        <Legend>Inways</Legend>
                        <Fieldset>
                            <StyledFormGroup>
                                <FieldArray
                                    name="inways"
                                    render={(arrayHelpers) => (
                                        <Fragment>
                                            {values.inways.map(
                                                (inway, index) => (
                                                    <StyledDeletableField
                                                        key={index}
                                                    >
                                                        <Field
                                                            name={`inways.${index}`}
                                                        />
                                                        <Button
                                                            secondary
                                                            type="button"
                                                            onClick={() =>
                                                                arrayHelpers.remove(
                                                                    index,
                                                                )
                                                            }
                                                            aria-label={`Delete inway ${inway}`}
                                                        >
                                                            -
                                                        </Button>
                                                    </StyledDeletableField>
                                                ),
                                            )}
                                            <Button
                                                secondary
                                                type="button"
                                                onClick={() =>
                                                    arrayHelpers.push('')
                                                }
                                            >
                                                + Add inway
                                            </Button>
                                        </Fragment>
                                    )}
                                />
                            </StyledFormGroup>
                        </Fieldset>

                        <StyledButtonGroup>
                            <Button type="submit">Save</Button>
                            {initialValues && (
                                <ConfirmButton onConfirm={onDelete}>
                                    Delete
                                </ConfirmButton>
                            )}
                        </StyledButtonGroup>
                    </Form>
                )}
            </Formik>
        )
    }
}

ServiceForm.propTypes = {
    onSubmit: func.isRequired,
    initialValues: object,
}

export default ServiceForm

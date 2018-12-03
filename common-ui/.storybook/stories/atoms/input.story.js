import React, { Fragment } from 'react'
import { Formik, Field } from 'formik'
import Input from 'src/Input/Input'

const initialValues = {
    name1: '',
    name2: '',
    name3: '',
    name4: ''
}

const createExample = () => (
    <Fragment>
        <Field component={Input} name="name1" label="Full name" />
        <Field component={Input} name="name2" label="Full name" required />
        <Field component={Input} name="name4" label="Full name" required disabled />
    </Fragment>
)

export const input = (
    <Formik initialValues={initialValues} render={createExample}/>
)

export const info = `
This component has to be wrapped by [<Formik>](https://jaredpalmer.com/formik/docs/api/formik) and passed
as \`component\` prop with <Field>.
`

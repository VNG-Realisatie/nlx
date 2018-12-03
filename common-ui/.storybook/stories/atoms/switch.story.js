import React, {Fragment} from 'react'
import { Formik, Field } from 'formik'
import Switch from 'src/Switch/Switch'

const initialValues = {
    name1: '',
    name2: '',
    name3: '',
}

const createExample = () => (
    <Fragment>
        <Field component={Switch} name="name1" label="Stuur mij meldingen per e-mail" />
        <Field component={Switch} name="name2" label="Stuur mij meldingen per e-mail" required />
        <Field component={Switch} name="name3" label="Stuur mij meldingen per e-mail" required disabled />
    </Fragment>
)

export const switchStory = (
    <Formik initialValues={initialValues} render={createExample}/>
)
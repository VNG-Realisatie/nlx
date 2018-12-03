import React, {Fragment} from 'react'
import { Flex, Box } from '@rebass/grid'
import { Formik, Field } from 'formik'

import Checkbox from 'src/Checkbox/Checkbox'

const initialValues = {
    name1: '',
    name2: '',
    name3: '',
}

const createExample = () => (
    <Fragment>
        <Field component={Checkbox} name="name1" label="Send me notifications" />
        <Field component={Checkbox} name="name2" label="Send me notifications" required />
        <Field component={Checkbox} name="name3" label="Send me notifications" required disabled />
    </Fragment>
)

export const checkboxStory = (
    <Formik initialValues={initialValues} render={createExample}/>
)
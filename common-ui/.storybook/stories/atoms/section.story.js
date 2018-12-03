import React, {Fragment} from 'react'
import Section from 'src/Section/Section'

export const sectionStory = (
    <Fragment>
        <Section>
            {'This is a <Section />'}
        </Section>
        <Section backgroundColor="whitesmoke">
            {'This is a grey <Section />'}
        </Section>
    </Fragment>
)
